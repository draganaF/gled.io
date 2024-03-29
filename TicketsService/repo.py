from sqlalchemy.orm import Session

import models
import schema

def get_ticket(db: Session, ticket_id: int):
  return db.query(models.Ticket).filter(models.Ticket.id == ticket_id, models.Ticket.is_deleted == False).first()

def get_tickets(db: Session):
  return db.query(models.Ticket).filter(models.Ticket.is_deleted == False).all()

def get_tickets_by_user_id(db: Session, user_id: int):
  return db.query(models.Ticket).filter(models.Ticket.user_id == user_id, models.Ticket.is_deleted == False).all()

def get_tickets_by_projection_id(db: Session, projection_id: int):
  return db.query(models.Ticket).filter(models.Ticket.projection_id == projection_id, models.Ticket.is_deleted == False).all()

def get_tickets_by_seat(db: Session, seat: str, projection_id: int):
  return db.query(models.Ticket).filter(models.Ticket.seat == seat, models.Ticket.is_deleted == False, models.Ticket.projection_id == projection_id).first()

def get_reserved_but_not_bought_tickets(db: Session):
  return db.query(models.Ticket).filter(models.Ticket.is_bought == False, models.Ticket.is_reserved == True, models.Ticket.is_deleted == False)

def create_ticket(db: Session, ticket: schema.CreateTicket):
  ticket_db = models.Ticket(**ticket.dict())
  db.add(ticket_db)
  db.commit()
  db.refresh(ticket_db)
  return ticket_db

def find_reserved_tickets_users(db: Session, projection_id: int):
  tickets = db.query(models.Ticket).filter(
    models.Ticket.is_bought == False, models.Ticket.is_reserved == True, models.Ticket.projection_id == projection_id)
  return tickets

def delete_unbought_reserved_tickets_for_projection(db: Session, projection_id: int):
  tickets = db.query(models.Ticket).filter(
    models.Ticket.is_bought == False, models.Ticket.is_reserved == True, models.Ticket.projection_id == projection_id, models.Ticket.is_deleted == False).all()

  for ticket in tickets:
    ticket.is_deleted = True
    db.add(ticket)
    db.commit()

def buy_reserved_ticket(db: Session, id: int):
  ticket = get_ticket(db, id)
  if not ticket:
    return None
  
  ticket.is_bought = True
  
  db.add(ticket)
  db.commit()
  
  return ticket