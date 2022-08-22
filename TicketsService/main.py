from fastapi import Depends, FastAPI, HTTPException, Request
from sqlalchemy.orm import Session
import schema, models, repo, remote_calls
from fastapi.middleware.cors import CORSMiddleware
from database import SessionLocal, engine

models.Base.metadata.create_all(bind=engine)

app = FastAPI()
origins = [
    "http://localhost:8082",
    "http://localhost:8083",
    "http://localhost:8086",
    "http://localhost:8080",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/api/tickets/{id}", status_code=200)
def get_ticket(reques: Request, id, db: Session = Depends(get_db)):
  token = reques.headers.get("Authorization")
  status_code = remote_calls.authorize(token, [0, 1, 2])
  if status_code == 400:
    raise HTTPException(status_code=400, detail="No header")
  if status_code == 401:
    raise HTTPException(status_code=401, detail="You are unauthorized")
  if status_code == 403:
    raise HTTPException(status_code=403, detail="You dont have permission")

  ticket = repo.get_ticket(db=db, ticket_id=id)
  if ticket == None:
    return {"status_code": 404, "message": "Not found"}

  return ticket

@app.get("/api/tickets", response_model=list[schema.Ticket], status_code=200)
def get_all_tickets(request: Request, db: Session = Depends(get_db)):
  token = request.headers.get("Authorization")
  status_code = remote_calls.authorize(token, [0, 1, 2])
  if status_code == 400:
    raise HTTPException(status_code=400, detail="No header")
  if status_code == 401:
    raise HTTPException(status_code=401, detail="You are unauthorized")
  if status_code == 403:
    raise HTTPException(status_code=403, detail="You dont have permission")

  tickets = repo.get_tickets(db)
  if tickets == None:
    raise HTTPException(status_code=404, detail="Ticket not found")
  return tickets

@app.post("/api/tickets",response_model=schema.Ticket, status_code=201)
def create_ticket(request: Request, ticket: schema.CreateTicket, db: Session = Depends(get_db)):

  token = request.headers.get("Authorization")
  status_code = remote_calls.authorize(token, [0, 1, 2])
  if status_code == 400:
    raise HTTPException(status_code=400, detail="No header")
  if status_code == 401:
    raise HTTPException(status_code=401, detail="You are unauthorized")
  if status_code == 403:
    raise HTTPException(status_code=403, detail="You dont have permission")
  print("Tu sam jeebs")
  db_ticket = repo.get_tickets_by_seat(db, seat=ticket.seat, projection_id=ticket.projection_id)
  if db_ticket:
    raise HTTPException(status_code=400, detail="Seat is taken")

  if ticket.is_reserved == True:
    remote_calls.increment_users_reserved_tickets(ticket.user_id)
  else:
    if ticket.user_id != -1:
      if remote_calls.buy_tickets(ticket.user_id, ticket.price) == 200:
        print("U price ssam")
        remote_calls.increment_users_bougth_tickets(ticket.user_id)
      else:
        raise HTTPException(status_code=400, detail="Not enough resource on your ballans")

  saved_ticket = repo.create_ticket(db=db, ticket=ticket)

  return saved_ticket

@app.get("/api/tickets/users/{user_id}", response_model=list[schema.Ticket], status_code=200)
def get_tickets_from_user(request: Request, user_id, db: Session = Depends(get_db)):
  token = request.headers.get("Authorization")
  status_code = remote_calls.authorize(token, [0, 1, 2])
  if status_code == 400:
    raise HTTPException(status_code=400, detail="No header")
  if status_code == 401:
    raise HTTPException(status_code=401, detail="You are unauthorized")
  if status_code == 403:
    raise HTTPException(status_code=403, detail="You dont have permission")

  tickets = repo.get_tickets_by_user_id(db=db, user_id=user_id)
  return tickets

@app.get("/api/tickets/projections/{projection_id}", response_model=list[schema.Ticket], status_code=200)
def get_tickets_for_projection(request: Request, projection_id, db: Session = Depends(get_db)):
  token = request.headers.get("Authorization")
  status_code = remote_calls.authorize(token, [0, 1, 2])
  if status_code == 400:
    raise HTTPException(status_code=400, detail="No header")
  if status_code == 401:
    raise HTTPException(status_code=401, detail="You are unauthorized")
  if status_code == 403:
    raise HTTPException(status_code=403, detail="You dont have permission")

  tickets = repo.get_tickets_by_projection_id(db=db, projection_id=projection_id)
  return tickets

@app.get("/api/tickets/buy-reserved-ticket/{id}",response_model=schema.Ticket, status_code=200)
def buy_ticket(request: Request, id, db:Session = Depends(get_db)):
  token = request.headers.get("Authorization")
  status_code = remote_calls.authorize(token, [0, 1, 2])
  if status_code == 400:
    raise HTTPException(status_code=400, detail="No header")
  if status_code == 401:
    raise HTTPException(status_code=401, detail="You are unauthorized")
  if status_code == 403:
    raise HTTPException(status_code=403, detail="You dont have permission")

  ticket = repo.buy_reserved_ticket(db=db, id=id)
  remote_calls.increment_users_bougth_tickets(user_id=ticket.user_id)

  return ticket

@app.get("/api/tickets/deleting-unbought-tickets/{projection_id}", status_code=200)
def delete_unbougth_tickets(projection_id, db: Session = Depends(get_db)):
  tickets = repo.find_reserved_tickets_users(db=db, projection_id=projection_id)

  users_id = set()
  for ticket in tickets:
    users_id.add(ticket.user_id)

  for user_id in users_id:
    status_code = remote_calls.increment_users_negative_points(user_id=user_id)
    if status_code != 200:
      raise HTTPException(status_code=400, detail="Could not increment negative points")

  repo.delete_unbought_reserved_tickets_for_projection(db=db, projection_id=projection_id)
