use serde_derive::{Serialize, Deserialize};

#[derive(Serialize, Clone, Deserialize)]
pub struct ReturnScore {
  pub Score: i32,
}