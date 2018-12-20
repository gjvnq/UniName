package UniName

import (
	"regexp"
	"strings"
)

const (
	// STYLE_BRAZIL_OFFICIAL never abreviates
	STYLE_BRAZIL_OFFICIAL = "STYLE_BRAZIL_OFFICIAL"
	STYLE_BRAZIL_COMMON   = "STYLE_BRAZIL_COMMON"
)

var bra_prefix *regexp.Regexp = regexp.MustCompile("(?i)(Professor(a|ª)|Doutor(a|ª)|Senhor(a|ª)|Mestre(a|ª)|Prof(.)(a|ª)(.)|Dr(.)(a|ª)(.)|Sr(.)(a|ª)(.)|Dom)")
var bra_suffix *regexp.Regexp = regexp.MustCompile("(?i)(Jr(.)|Júnior(.)|Net(o|a))")
var bra_connector *regexp.Regexp = regexp.MustCompile("(?i)(de|d(a|o)(s|)|e)")

func brazil_to_title(name string) string {
	b := strings.Builder{}
	b.Grow(len(name))

	l := strings.Split(name, " ")
	for i, part := range l {
		if i != 0 {
			b.WriteRune(' ')
		}
		got := bra_connector.FindString(part)
		if got == part {
			b.WriteString(strings.ToLower(part))
		} else {
			b.WriteString(firstUpperEachWord(part))
		}
	}

	return b.String()
}

func brazil_fill(original string) WesternName {
	return WesternName{}
}

func brazil_full_suffix(value string) string {
	str := strings.Replace(strings.ToLower(trim_spaces_and_doubles(value)), ".", "", -1)

	if str == "jr" || str == "juinor" || str == "júnior" {
		return "Júnior"
	}

	return value
}

func brazil_full_name(name WesternName, style NameStyle) string {
	if style == STYLE_BRAZIL_OFFICIAL {
		ans := str_joiner(" ", false,
			name.First,
			name.BeforeMiddle,
			name.Middle,
			name.BeforeLast,
			name.Last,
			brazil_full_suffix(name.Suffix))
		return strings.ToUpper(trim_spaces_and_doubles(ans))
	}
	if style == STYLE_BRAZIL_COMMON {
		ans := str_joiner(" ", true,
			name.Prefix,
			name.GetFisrtOrSocial(),
			name.BeforeMiddle,
			name.Middle,
			name.BeforeLast,
			name.Last,
			name.Suffix)
		if name.Nick != "" {
			ans = "(" + name.Nick + ") " + ans
		}
		return brazil_to_title(trim_spaces_and_doubles(ans))
	}

	return name.Original
}

func brazil_abrev(name WesternName, style AbrevStyle, max_len int) string {
	return name.Original
}
