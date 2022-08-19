use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct AuthenticatedUser {
  pub id: String,
  pub role: String
}