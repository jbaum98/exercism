use std::char;

pub fn reply(s: &str) -> String {
    if is_silent(s) {
        "Fine. Be that way!".into()
    } else if is_question(s) {
        "Sure.".into()
    } else if is_shouting(s) {
        "Whoa, chill out!".into()
    } else {
        "Whatever.".into()
    }
}

fn is_question(s: &str) -> bool {
    s.ends_with("?")
}

fn is_shouting(s: &str) -> bool {
    s.chars().filter(|c| char::is_alphabetic(*c)).all(char::is_uppercase)
}

fn is_silent(s: &str) -> bool {
    s == ""
}
