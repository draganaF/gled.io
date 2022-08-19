use serde_derive::{Serialize, Deserialize};

#[derive(Serialize, Clone, Deserialize)]
pub struct ReportRecension {
  pub id: i32,
  pub user_id: i32,
  pub recension_id: i32,
  pub deleted: bool
}