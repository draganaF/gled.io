use serde_derive::{Serialize, Deserialize};

#[derive(Serialize, Clone, Deserialize)]
pub struct CreateRecension {
  pub message: String,
  pub score: i32,
  pub user_id: i32,
  pub movie_id: i32
}