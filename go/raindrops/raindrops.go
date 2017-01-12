package raindrops

import "fmt"

const testVersion = 2

func divides(p, q int) bool {
	return q%p == 0
}

func Convert(n int) string {
	var out string

	if divides(3, n) {
		out += "Pling"
	}
	if divides(5, n) {
		out += "Plang"
	}
	if divides(7, n) {
		out += "Plong"
	}

	if out == "" {
		return fmt.Sprintf("%d", n)
	}

	return out
}
