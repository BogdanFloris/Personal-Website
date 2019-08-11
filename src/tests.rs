use std::fs::File;
use std::io::Read;

use rocket::local::Client;
use rocket::http::Status;

use super::rocket;

fn test_query_file<T> (path: &str, file: T, status: Status)
    where T: Into<Option<&'static str>>
{
    let client = Client::new(rocket()).unwrap();
    let mut response = client.get(path).dispatch();
    assert_eq!(response.status(), status);

    let body_data = response.body().and_then(|body| body.into_bytes());
    if let Some(filename) = file.into() {
        let expected_data = read_file_content(filename);
        assert!(body_data.map_or(false, |s| s == expected_data));
    }
}

fn read_file_content(path: &str) -> Vec<u8> {
    let mut fp = File::open(&path).expect(&format!("Can't open {}", path));
    let mut file_content = vec![];

    fp.read_to_end(&mut file_content).expect(&format!("Reading {} failed.", path));
    file_content
}

#[test]
fn test_index_html() {
    test_query_file("/", "public/html/index.html", Status::Ok);
    test_query_file("/?v=1", "public/html/index.html", Status::Ok);
    test_query_file("/?this=should&be=ignored", "public/html/index.html", Status::Ok);
}

#[test]
fn test_invalid_path() {
    test_query_file("/thou_shalt_not_exist", None, Status::NotFound);
    test_query_file("/thou/shalt/not/exist", None, Status::NotFound);
    test_query_file("/thou/shalt/not/exist?a=b&c=d", None, Status::NotFound);
}