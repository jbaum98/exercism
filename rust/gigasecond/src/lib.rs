extern crate chrono;
use chrono::*;

pub fn after(birthdate: DateTime<UTC>) -> DateTime<UTC> {
    birthdate.checked_add(Duration::seconds(1000000000)).unwrap()
}
