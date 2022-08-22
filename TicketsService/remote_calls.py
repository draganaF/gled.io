import requests

def authorize(token, roles):
  url = "http://localhost:8083/api/authorize"
  headers = {"Authorization": token}
  response = requests.get(url, headers=headers)

  if response.status_code != 200:
    return response.status_code

  role = int(response.json()["Role"])
  if role not in roles:
    return 403

  return 200
  

def increment_users_negative_points(user_id):
  url = "http://localhost:8083/api/users/increment-negative-points/" + str(user_id)
  response = requests.get(url)
  return response.status_code

def increment_users_bougth_tickets(user_id):
  url = "http://localhost:8083/api/users/increment-bought-tickets/" + str(user_id)
  response = requests.get(url)
  return response.status_code

def increment_users_reserved_tickets(user_id):
  url = "http://localhost:8083/api/users/increment-reserved-tickets/" + str(user_id)
  response = requests.get(url)
  return response.status_code

def buy_tickets(user_id, price):
  url = "http://localhost:8083/api/users/buy-tickets/" + str(user_id)
  response = requests.post(url, data={"Money": -1*price, "UserId": user_id})
  return response.status_code