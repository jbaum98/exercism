trait DivBy {
    fn div_by(self: Self, q: Self) -> bool;
}

impl DivBy for i32 {
    fn div_by(self: i32, q: i32) -> bool {
        self % q == 0
    }
}

pub fn is_leap_year(year: i32) -> bool {
    year.div_by(4) && (!year.div_by(100) || year.div_by(400))
}

