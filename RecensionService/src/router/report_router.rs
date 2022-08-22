use rocket::{response::content, request::{FromRequest, self}, Request, Outcome, http::{Status, RawStr}};
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{service::{report_service, auth_service::AuthService}, model::report::ReportRecension};
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

#[get("/all")]
pub fn get_all_reports(token: Token) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![0.to_string(), 1.to_string(), 2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };

  let mut report_s = report_service::ReportService::new();

  return content::Json(Json(json!(report_s.get_all().unwrap())).to_string());
}

#[post("/create", data = "<report>")]
pub fn create_report(token: Token, report: Json<ReportRecension>) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![0.to_string(), 1.to_string(), 2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };

  let mut report_s = report_service::ReportService::new();

  return content::Json(Json(json!(report_s.create_report(report.into_inner()))).to_string());
}

#[delete("/delete/<id>")]
pub fn delete_report(token: Token, id: i32) -> content::Json<String> {
  let auth_s = AuthService::new();

  let user = match auth_s.authorize(&token.0.to_string(), &vec![0.to_string(), 1.to_string(), 2.to_string()]) {
    Ok(user) => user,
    Err(err) => return rocket::response::content::Json(err)
  };

  let mut report_s = report_service::ReportService::new();

  report_s.delete_recension(id);
  return rocket::response::content::Json(String::from("report deleted"));
}