package UniName

const (
	ABREV_AS_MUCH_AS_RESONABLE = 0
	NO_ABREV                   = -1
)

type AbrevStyle string

const (
	ABREV_CLEAN  = "ABREV_CLEAN"
	ABREV_PERIOD = "ABREV_PERIOD"
)

type WesternName struct {
	Original      string
	OriginalStyle NameStyle
	Nick          string
	Prefix        string
	First         string
	FirstSocial   string // will replace the first name for trans people
	BeforeMiddle  string
	Middle        string
	BeforeLast    string
	Last          string
	Suffix        string
}

func (self WesternName) GetFisrtOrSocial() string {
	if self.FirstSocial != "" {
		return self.FirstSocial
	}
	return self.First
}
