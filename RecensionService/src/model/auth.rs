use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct AuthenticatedUser {
  pub Name: String,
  pub Role: String
}