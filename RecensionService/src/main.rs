#![feature(proc_macro_hygiene, decl_macro)]

use std::env;
use dotenv::dotenv;
use rocket::{Config, config::Environment};
use rocket::http::Method;
use rocket_cors::{AllowedOrigins, CorsOptions};


#[macro_use] extern crate rocket;
extern crate base64;
extern crate dotenv;

mod model;
mod api_contracts;
mod repository;
mod service;
mod utils;
mod router;

fn main() {
    println!("UBICU SE!");

    
    std::thread::spawn(|| {
        utils::database_driver::seed_db();
    }).join().expect("Thread panicked");
    
    dotenv().ok();

    let cfg = Config::build(Environment::Development)
        .address(env::var("HOST").unwrap())
        .port(env::var("PORT").unwrap().parse().unwrap())   
        .unwrap();

    let cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .allowed_methods(
            vec![Method::Get, Method::Post, Method::Put, Method::Delete, Method::Patch]
                .into_iter()
                .map(From::from)
                .collect(),
        )
        .allow_credentials(true);
    

    rocket::custom(cfg)
        .attach(cors.to_cors().unwrap())
        .mount("/api/recensions",  routes![
            router::recensions_router::get_all_recensions,
            router::recensions_router::get_by_id,
            router::recensions_router::get_by_movie_id,
            router::recensions_router::create_recension,
            router::recensions_router::delete_recension
        ])
        .mount("/api/reports", routes![
            router::report_router::get_all_reports,
            router::report_router::create_report,
            router::report_router::delete_report
        ])
        .launch();

        

}