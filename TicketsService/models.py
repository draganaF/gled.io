from sqlalchemy import Boolean, Column, Float ,Integer, String


from database import Base

class Ticket(Base):
  __tablename__ = "tickets"

  id = Column(Integer, primary_key=True, index=True)
  user_id = Column(Integer)
  projection_id = Column(Integer)
  price = Column(Float)
  seat = Column(String(16), index=True)
  is_reserved = Column(Boolean)
  is_bought = Column(Boolean)
  is_deleted = Column(Boolean)

