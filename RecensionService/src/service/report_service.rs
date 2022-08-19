use crate::{repository, model::report::ReportRecension};


pub struct ReportService {
  repository: repository::report_repo::ReportRepo
}

impl ReportService {
  pub fn new() -> ReportService {
    ReportService {
      repository: repository::report_repo::ReportRepo::new()
    }
  }

  pub fn get_all(&mut self) -> Result<Vec<ReportRecension>, String> {
    match self.repository.get_all() {
      Some(reports) => Ok(reports),
      None => Err("No reports available.".to_string())
    }
  }

  pub fn get_by_id(&mut self, id: i32) -> Result<ReportRecension, String> {
    match self.repository.get_by_id(id) {
      Some(report) => Ok(report),
      None => Err(format!("Could not return report with ID {}", id))
    }
  } 
  pub fn create_report(&mut self, cr_report: ReportRecension) -> Result<ReportRecension, String> {
    match self.repository.create_report(cr_report) {
      Some(report) => Ok(report),
      None => Err("Could not create report".to_string())
    }
  }

  pub fn delete_recension(&mut self, id: i32) {
    self.repository.delete_report(id);
  }
}