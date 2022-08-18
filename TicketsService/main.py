from fastapi import Depends, FastAPI, HTTPException
from sqlalchemy.orm import Session
import schema, models, repo, remote_calls

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

  if ticket.is_reserved == True:
    remote_calls.increment_users_reserved_tickets(ticket.user_id)
  else:
    if remote_calls.buy_tickets(ticket.user_id, ticket.price) == 200:
      remote_calls.increment_users_bougth_tickets(ticket.user_id)
    else:
      raise HTTPException(status_code=400, detail="Not enough resources on your ballans")

  saved_ticket = repo.create_ticket(db=db, ticket=ticket)
  
  return saved_ticket

@app.get("/api/tickets/users/{user_id}", response_model=list[schema.Ticket], status_code=200)
def get_tickets_from_user(user_id, db: Session = Depends(get_db)):
  tickets = repo.get_tickets_by_user_id(db=db, user_id=user_id)
  return tickets

@app.get("/api/tickets/projections/{projection_id}", response_model=list[schema.Ticket], status_code=200)
def get_tickets_for_projection(projection_id, db: Session = Depends(get_db)):
  tickets = repo.get_tickets_by_projection_id(db=db, projection_id=projection_id)
  return tickets

@app.get("/api/tickets/buy-reserved_ticket/{id}",response_model=schema.Ticket, status_code=200)
def buy_ticket(id, db:Session = Depends(get_db)):
  ticket = repo.buy_reserved_ticket(db=db, id=id)
  remote_calls.increment_users_bougth_tickets(user_id=ticket.user_id)

  return ticket

@app.get("/api/tickets/deleting-unbought-tickets/{projection_id}", status_code=200)
def delete_unbougth_tickets(projection_id, db: Session = Depends(get_db)):
  tickets = repo.find_reserved_tickets_users(db=db, projection_id=projection_id)

  users_id = {}
  for ticket in tickets:
    users_id.add(ticket["user_id"])

  for user_id in users_id:
    status_code = remote_calls.increment_users_negative_points(user_id=user_id)
    if status_code != 200:
      raise HTTPException(status_code=400, detail="Could not increment negative points")

  repo.delete_unbought_reserved_tickets_for_projection(db=db, projection_id=projection_id)
