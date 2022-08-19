
use postgres::{Client, NoTls};

use crate::model::recension::Recension;
use crate::api_contracts::create_recension::CreateRecension;

pub struct RecensionsRepo {
  client: Client 
}

impl RecensionsRepo {
  pub fn new() -> RecensionsRepo {
    RecensionsRepo {
      client: Client::connect("postgresql://postgres:root@localhost:5432/postgres",
      NoTls,).unwrap()
    }
  }

  pub fn get_all(&mut self) -> Option<Vec<Recension>> {
 
    let mut recensions: Vec<Recension> = vec![];

    let rows = self.client.query("SELECT id, userId, movieId, message, score, deleted FROM recensions WHERE deleted = false", &[]);

    if rows.is_err() {
      return None;
    }
    let unwraped_rows = rows.unwrap();

    for row in unwraped_rows {
      let id: i32 = row.get(0);
      let user_id: i32 = row.get(1);
      let movie_id: i32 = row.get(2);
      let message: String = row.get(3);
      let score: i32 = row.get(4);
      let deleted: bool = row.get(5);

      recensions.push(Recension 
        {
          id: (id),
          user_id: (user_id),
          movie_id: (movie_id),
          message: (message),
          score: (score),
          deleted: (deleted)
        });
    }

    Some(recensions)
  }

  pub fn get_by_id(&mut self, id :i32) -> Option<Recension> {
    let mut recension: Recension;
  
    let rows = self.client.query("SELECT id, userId, movieId, message, score, deleted FROM recensions WHERE deleted = false AND id = $1", &[&id]).unwrap();
  
    if rows.is_empty() {
      return None;
    }

    let row = rows.get(0).unwrap();

    let id: i32 = row.get(0);
    let user_id: i32 = row.get(1);
    let movie_id: i32 = row.get(2);
    let message: String = row.get(3);
    let score: i32 = row.get(4);
    let deleted: bool = row.get(5);
  
    Some(
      Recension{
        id: (id),
        user_id: (user_id),
        movie_id: (movie_id),
        message: (message),
        score: (score),
        deleted: (deleted)
    })
  
  }

  pub fn get_by_movie_id(&mut self, movie_id: i32) -> Option<Vec<Recension>> {
    let mut recensions: Vec<Recension> = vec![];

    let rows = self.client.query("SELECT id, userId, movieId, message, score, deleted FROM recensions WHERE deleted = false AND movieId = $1", &[&movie_id]).unwrap();

    if rows.is_empty() {
      return None;
    }
    let unwraped_rows = rows;

    for row in unwraped_rows {
      let id: i32 = row.get(0);
      let user_id: i32 = row.get(1);
      let movie_id: i32 = row.get(2);
      let message: String = row.get(3);
      let score: i32 = row.get(4);
      let deleted: bool = row.get(5);

      recensions.push(Recension 
        {
          id: (id),
          user_id: (user_id),
          movie_id: (movie_id),
          message: (message),
          score: (score),
          deleted: (deleted)
        });
    }

    Some(recensions)
  }

  pub fn create_recension(&mut self, recension: CreateRecension) -> Option<Recension> {
    
    let rows = self.client.query("SELECT id, userId, movieId, message, score, deleted FROM recensions WHERE userId = $1 AND movieId = $2", &[&recension.user_id, &recension.movie_id]);
    
    if !rows.unwrap().is_empty() {
      return None;
    }
    
    let rows_affected = self.client.execute(
      "INSERT INTO recensions (userId, movieId, deleted, message, score) VALUES ($1, $2, $3, $4, $5)", 
      &[&recension.user_id, &recension.movie_id, &false, &recension.message, &recension.score]).unwrap();
    
    if rows_affected == 0 {
      return None;
    }

    Some(
      Recension{
        id: 0,
        user_id: recension.user_id,
        movie_id: recension.movie_id,
        deleted: false,
        message: recension.message,
        score: recension.score
      }
    )
  }

  pub fn delete_recension(&mut self, id: i32) {
    self.client.execute(
      "UPDATE recensions SET deleted=true WHERE id = $1",
      &[&id],).unwrap();
  }
}

