use rocket::{response::content, request::{FromRequest, self}, Request, Outcome, http::{Status, RawStr}};
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{service::{recensions_service, auth_service::AuthService}, api_contracts::{self, create_recension::CreateRecension}};

#[derive(Debug)]
pub struct Token(String);

#[derive(Debug)]
pub enum ApiTokenError {
  Missing,
}

impl<'a, 'r> FromRequest<'a, 'r> for Token {
  type Error = ApiTokenError;

  fn from_request(request: &'a Request<'r>) -> request::Outcome<Self, Self::Error> {
    let token = request.headers().get_one("Authorization");
    match token {
      Some(token) => { Outcome::Success(Token(token.to_string())) },
      None => Outcome::Failure((Status::Unauthorized, ApiTokenError::Missing)),
    }
  }
}

#[get("/all", format = "application/json")]
pub fn get_all_recensions(token: Token) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };
  
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.get_all())).to_string());
}

#[get("/by-id/<id>", format = "application/json")]
pub fn get_by_id(token: Token, id :i32) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };
  
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.get_by_id(id))).to_string());
}

#[get("/by-movie-id/<id>")]
pub fn get_by_movie_id(token: Token, id :i32) -> content::Json<String> {
  let auth_s = AuthService::new();
  println!("{}", &token.0.to_string());
  println!("Sta se desava");
  let user = match auth_s.authorize(&token.0.to_string(), &vec![0.to_string(), 1.to_string(), 2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };

  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.get_by_movie_id(id))).to_string());
}

#[post("/create", data = "<recension>")]
pub fn create_recension(token: Token, recension: Json<CreateRecension>) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![0.to_string(), 2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.create_recension(recension.into_inner()))).to_string());
}

#[delete("/delete/<id>")]
pub fn delete_recension(token: Token, id: i32) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![0.to_string(), 2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };
  let mut recension_s = recensions_service::RecensionsService::new();
  recension_s.delete_recension(id);

  return rocket::response::content::Json(String::from("recension deleted"));
}