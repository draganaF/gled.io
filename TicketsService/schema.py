from pydantic import BaseModel

class BaseTicket(BaseModel):
  user_id: int
  projection_id: int
  price: float
  seat: str
  is_reserved: bool
  is_bought: bool
  is_deleted: bool

class CreateTicket(BaseTicket):
  pass

class Ticket(BaseTicket):
  id: int

  class Config: 
    orm_mode = True
    