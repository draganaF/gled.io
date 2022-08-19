use std::string;

use reqwest::StatusCode;

use crate::model::auth::AuthenticatedUser;



pub struct AuthService {
  client: reqwest::blocking::Client
}

impl AuthService {
    
  pub fn new() -> AuthService {
    AuthService {
      client: reqwest::blocking::Client::new()
    }
  }

  pub fn authorize(&self, token: &String, roles: &Vec<String>) -> Result<AuthenticatedUser, String> {

    let response = self.client.get("http://localhost:8083/api/authorize")
      .header("Authorization", token)
      .send()
      .unwrap();
      
    match response.status() {
      StatusCode::OK => {},
      _ => return Err("You are not authorized".to_string())
    };

    let authenticated_user = response.json::<AuthenticatedUser>().unwrap();
    
    if roles.contains(&authenticated_user.Role) {
      Ok(authenticated_user)
    } else {
      Err("You are not authorized.".to_string())
    }
  }

}