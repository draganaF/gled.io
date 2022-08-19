use postgres::{Client, Error, NoTls};

pub fn seed_db() -> Result<(), Error> {
  let mut client = Client::connect(
      "postgresql://postgres:root@localhost:5432/postgres",
      NoTls,
  )?;

  client.batch_execute("DROP TABLE IF EXISTS recensions")?;

  client.batch_execute(
      "
      CREATE TABLE IF NOT EXISTS recensions (
          id              SERIAL PRIMARY KEY,
          userId          INTEGER,
          movieId          INTEGER,
          deleted        BOOLEAN,
          message VARCHAR NOT NULL,
          score INTEGER
          )
  ",
  )?;

  client.execute(
      "INSERT INTO recensions (userId, movieId, deleted, message, score) VALUES ($1, $2, $3, $4, $5)",
      &[&3.to_owned(), &1.to_owned(), &false, &"Film je bio ocaj", &2.to_owned()],
  )?;

  client.execute(
    "INSERT INTO recensions (userId, movieId, deleted, message, score) VALUES ($1, $2, $3, $4, $5)",
    &[&3.to_owned(), &2.to_owned(), &false, &"Odlican", &5.to_owned()],
  )?;

  client.execute(
    "INSERT INTO recensions (userId, movieId, deleted, message, score) VALUES ($1, $2, $3, $4, $5)",
    &[&3.to_owned(), &3.to_owned(), &false, &"Moze to bolje", &3.to_owned()],
  )?;

  client.execute(
    "INSERT INTO recensions (userId, movieId, deleted, message, score) VALUES ($1, $2, $3, $4, $5)",
    &[&4.to_owned(), &1.to_owned(), &false, &"Jako kul", &5.to_owned()],
  )?;

  client.execute(
    "INSERT INTO recensions (userId, movieId, deleted, message, score) VALUES ($1, $2, $3, $4, $5)",
    &[&4.to_owned(), &2.to_owned(), &false, &"Valja!", &5.to_owned()],
  )?;

  client.batch_execute("DROP TABLE IF EXISTS reports")?;

  client.batch_execute(
      "
      CREATE TABLE IF NOT EXISTS reports (
          id              SERIAL PRIMARY KEY,
          userId          INTEGER,
          recensionId     INTEGER,
          deleted         BOOLEAN
          )
  ",
  )?;

  client.execute(
      "INSERT INTO reports (userId, recensionId, deleted) VALUES ($1, $2, $3)",
      &[&4.to_owned(), &1.to_owned(), &false],
  )?;

  client.execute(
    "INSERT INTO reports (userId, recensionId, deleted) VALUES ($1, $2, $3)",
    &[&3.to_owned(), &4.to_owned(), &false],
  )?;

  client.close()?;

  Ok(())
}