use rocket::response::content;
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{service::report_service, model::report::ReportRecension};

#[get("/all")]
pub fn get_all_reports() -> content::Json<String> {
  let mut report_s = report_service::ReportService::new();

  return content::Json(Json(json!(report_s.get_all())).to_string());
}

#[post("/create", data = "<report>")]
pub fn create_report(report: Json<ReportRecension>) -> content::Json<String> {
  let mut report_s = report_service::ReportService::new();

  return content::Json(Json(json!(report_s.create_report(report.into_inner()))).to_string());
}

#[delete("/delete/<id>")]
pub fn delete_report(id: i32) -> String {
  let mut report_s = report_service::ReportService::new();

  report_s.delete_recension(id);
  return "Uspjesno brisanje".to_string();
}