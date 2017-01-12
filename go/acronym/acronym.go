package acronym

import "io"
import "strings"
import "bytes"
import "unicode"

const testVersion = 2

func Abbreviate(input string) string {
	var buffer bytes.Buffer

	words := strings.FieldsFunc(input, func(r rune) bool {
		switch r {
		case ' ', '-':
			return true
		default:
			return false
		}
	})

	for _, w := range words {
		seenLower := false
		var tempBuf bytes.Buffer
		for i, r := range w {
			if i == 0 {
				r = unicode.ToUpper(r)
				buffer.WriteRune(r)
			} else if unicode.IsUpper(r) {
				tempBuf.WriteRune(r)
			} else if !seenLower && unicode.IsLower(r) {
				seenLower = true
			}
		}
		if seenLower {
			io.Copy(&buffer, &tempBuf)
		}
	}

	return buffer.String()
}
