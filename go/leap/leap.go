// Package leap contains functions concerning leap years
package leap

const testVersion = 2

func divides(q int, p int) bool {
	return p%q == 0
}

func IsLeapYear(year int) bool {
	return divides(4, year) && (!divides(100, year) || divides(400, year))
}
