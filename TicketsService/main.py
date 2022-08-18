from fastapi import Depends, FastAPI, HTTPException, status
from sqlalchemy.orm import Session
from fastapi.responses import JSONResponse
import schema, models, repo

from database import SessionLocal, engine

models.Base.metadata.create_all(bind=engine)

app = FastAPI()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/api/tickets/{id}", response_model=schema.Ticket, status_code=200)
def get_ticket(id, db: Session = Depends(get_db)):
  ticket = repo.get_ticket(db=db, ticket_id=id)
  return ticket

@app.get("/api/tickets", response_model=list[schema.Ticket], status_code=200)
def get_all_tickets(db: Session = Depends(get_db)):
  tickets = repo.get_tickets(db)
  return tickets

@app.post("/api/tickets",response_model=schema.Ticket, status_code=201)
def create_ticket(ticket: schema.CreateTicket, db: Session = Depends(get_db)):
  db_ticket = repo.get_tickets_by_seat(db, seat=ticket.seat)
  if db_ticket:
    raise HTTPException(status_code=400, detail="Seat is taken")

  saved_ticket = repo.create_ticket(db=db, ticket=ticket)
  return saved_ticket

@app.get("/api/tickets/users/{user_id}", response_model=list[schema.Ticket])
def get_tickets_from_user(user_id, db: Session = Depends(get_db)):
  tickets = repo.get_tickets_by_user_id(db=db, user_id=user_id)
  return tickets

@app.get("/api/tickets/projections/{projection_id}", response_model=list[schema.Ticket])
def get_tickets_for_projection(projection_id, db: Session = Depends(get_db)):
  tickets = repo.get_tickets_by_projection_id(db=db, projection_id=projection_id)
  return tickets

@app.get("/api/tickets/buy-reserved_ticket/{id}")
def buy_ticket(id, db:Session = Depends(get_db)):
  ticket = repo.buy_reserved_ticket(db=db, id=id)
  return ticket

@app.get("/api/tickets/deleting-unbought-tickets/{projection_id}", status_code=200)
def delete_unbougth_tickets(projection_id, db: Session = Depends(get_db)):
  repo.delete_unbought_reserved_tickets_for_projection(db=db, projection_id=projection_id)
