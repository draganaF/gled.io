use crate::{repository, model::recension::Recension, api_contracts::create_recension::CreateRecension};

pub struct RecensionsService {
  repository: repository::recensions_repo::RecensionsRepo
}

impl RecensionsService {
  pub fn new() -> RecensionsService {
    RecensionsService { repository:     repository::recensions_repo::RecensionsRepo::new()
    }
  }

  pub fn get_all(&mut self) -> Result<Vec<Recension>, String> {
    match self.repository.get_all() {
      Some(recensions) => Ok(recensions),
      None => Err("No recensions available.".to_string())
    }
  }

  pub fn get_by_id(&mut self, id: i32) -> Result<Recension, String> {
    match self.repository.get_by_id(id) {
      Some(recension) => Ok(recension),
      None => Err(format!("Could not return recension with ID {}", id))
    }
  }

  pub fn get_by_movie_id(&mut self, movie_id: i32) -> Result<Vec<Recension>, String> {
    match self.repository.get_by_movie_id(movie_id) {
      Some(recensios) => Ok(recensios),
      None => Err(format!("No recensions with movie ID {}", movie_id))
    }
  }
 
  pub fn create_recension(&mut self, cr_recension: CreateRecension) -> Result<Recension, String> {
    match self.repository.create_recension(cr_recension) {
      Some(recension) => Ok(recension),
      None => Err("Could not create recension".to_string())
    }
  }

  pub fn delete_recension(&mut self, id: i32) {
    self.repository.delete_recension(id);
  }
}