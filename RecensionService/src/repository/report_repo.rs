use postgres::{Client, NoTls};
use serde::__private::de;

use crate::model::report::ReportRecension;

pub struct ReportRepo {
  client: Client
}

impl ReportRepo {
  pub fn new() -> ReportRepo {
    ReportRepo {
      client: Client::connect("postgresql://postgres:root@localhost:5432/postgres",
      NoTls,).unwrap()
    }
  }

  pub fn get_all(&mut self) -> Option<Vec<ReportRecension>> {
    let mut reports: Vec<ReportRecension> = vec![];

    let rows = self.client.query("SELECT id, userId, recensionId, deleted FROM reports WHERE deleted = false", &[]);

    if rows.is_err() {
      return None;
    }
    let unwraped_rows = rows.unwrap();

    for row in unwraped_rows {
      let id: i32 = row.get(0);
      let user_id: i32 = row.get(1);
      let recension_id: i32 = row.get(2);
      let deleted: bool = row.get(3);

      reports.push(ReportRecension { 
        id: (id), 
        user_id: (user_id), 
        recension_id: (recension_id), 
        deleted: (deleted) 
      });
    }

    Some(reports)
  }

  pub fn get_by_id(&mut self, id:i32) -> Option<ReportRecension> {

  let mut recension: ReportRecension;
  
    let rows = self.client.query("SELECT id, userId, recensionId, deleted FROM reports WHERE deleted = false AND id = $1", &[&id]);
  
    if rows.is_err() {
      return None;
    }

    let binding = rows.unwrap();
    let row = binding.get(0).unwrap();

    let id: i32 = row.get(0);
    let user_id: i32 = row.get(1);
    let recension_id: i32 = row.get(2);
    let deleted: bool = row.get(3);
  
    Some(
      ReportRecension{
        id: (id),
        user_id: (user_id),
        recension_id: (recension_id),
        deleted: (deleted)
    })
  
  }

  pub fn create_report(&mut self, cr_report: ReportRecension) -> Option<ReportRecension> {
    let rows = self.client.query("SELECT id, userId, recensionId, deleted FROM reports WHERE userId = $1 AND recensionId = $2", &[&cr_report.user_id, &cr_report.recension_id]).unwrap();

    if !rows.is_empty() {
      return None;
    }
    
    let rows_affected = self.client.execute(
      "INSERT INTO reports (userId, recensionId, deleted) VALUES ($1, $2, $3)", 
      &[&cr_report.user_id, &cr_report.recension_id, &false]).unwrap();
    
    if rows_affected == 0 {
      return None;
    }

    Some(
      ReportRecension{
        id: 0,
        user_id: cr_report.user_id,
        recension_id: cr_report.recension_id,
        deleted: false
      }
    )
  }

  pub fn delete_report(&mut self, id: i32) {
    self.client.execute(
      "UPDATE reports SET deleted=true WHERE id = $1",
      &[&id],).unwrap();
  }
}