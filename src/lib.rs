#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;
#[macro_use]
extern crate rocket_contrib;

#[cfg(test)] mod tests;

use std::path::Path;
use rocket::response::NamedFile;
use rocket_contrib::serve::StaticFiles;
use rocket_contrib::json::JsonValue;

#[catch(404)]
fn not_found() -> JsonValue {
    json!({
        "status": "error",
        "reason": "Resource was not found."
    })
}

#[get("/")]
fn index() -> Option<NamedFile> {
    NamedFile::open(Path::new("static/html/index.html")).ok()
}

#[get("/blog")]
fn blog() -> Option<NamedFile> {
    NamedFile::open(Path::new("static/html/under-construction.html")).ok()
}

pub fn rocket() -> rocket::Rocket {
    rocket::ignite()
        .mount("/", routes![index, blog])
        .mount("/static", StaticFiles::from("static"))
        .register(catchers![not_found])
}
