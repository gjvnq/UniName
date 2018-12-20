package UniName

import (
	"strings"
	"unicode"
)

func firstUpperEachWord(str string) string {
	b := strings.Builder{}
	b.Grow(len(str))
	in_word := false
	for _, char := range str {
		if !unicode.IsLetter(char) {
			b.WriteRune(unicode.ToUpper(char))
			in_word = false
		} else if !in_word {
			b.WriteRune(unicode.ToUpper(char))
			in_word = true
		} else {
			b.WriteRune(unicode.ToLower(char))
		}
	}
	return b.String()
}

func trim_spaces_and_doubles(name string) string {
	b := strings.Builder{}
	b.Grow(len(name))
	in_space := true
	for _, char := range name {
		if in_space && unicode.IsSpace(char) {
			continue
		}
		in_space = unicode.IsSpace(char)
		b.WriteRune(char)
	}
	return b.String()
}

func str_joiner(joiner string, ignore_empty bool, list ...string) string {
	total_n := 0
	for _, part := range list {
		total_n += len(part)
	}
	b := strings.Builder{}
	b.Grow(total_n)
	for i, part := range list {
		if !(i != 0 && part == "" && !ignore_empty) {
			b.WriteString(joiner)
		}
		b.WriteString(part)
	}
	return b.String()
}
