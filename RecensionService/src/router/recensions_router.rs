use rocket::response::content;
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{service::recensions_service, api_contracts::{self, create_recension::CreateRecension}};


#[get("/all")]
pub fn get_all_recensions() -> content::Json<String> {
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.get_all())).to_string());
}

#[get("/by-id/<id>")]
pub fn get_by_id(id :i32) -> content::Json<String> {
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.get_by_id(id))).to_string());
}

#[get("/by-movie-id/<id>")]
pub fn get_by_movie_id(id :i32) -> content::Json<String> {
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.get_by_movie_id(id))).to_string());
}

#[post("/create", data = "<recension>")]
pub fn create_recension(recension: Json<CreateRecension>) -> content::Json<String> {
  let mut recension_s = recensions_service::RecensionsService::new();

  return content::Json(Json(json!(recension_s.create_recension(recension.into_inner()))).to_string());
}