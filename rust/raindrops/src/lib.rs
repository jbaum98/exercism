trait DivBy {
    fn div_by(self: Self, q: Self) -> bool;
}

impl DivBy for i32 {
    fn div_by(self: i32, q: i32) -> bool {
        self % q == 0
    }
}

pub fn raindrops(x: i32) -> String {
    let factors = (x.div_by(3), x.div_by(5), x.div_by(7));
    match factors {
        (false, false, false) => format!("{}", x),
        (true, false, false) => "Pling".to_string(),
        (false, true, false) => "Plang".to_string(),
        (false, false, true) => "Plong".to_string(),
        (true, true, false) => "PlingPlang".to_string(),
        (true, false, true) => "PlingPlong".to_string(),
        (false, true, true) => "PlangPlong".to_string(),
        (true, true, true) => "PlingPlangPlong".to_string(),
    }
}
