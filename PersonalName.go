package UniName

import "errors"

type PersonalName interface {
	Get(style NameStyle, variation NameStyleVariation, part NameStylePart, max_length int) string
	MayBeCloseRelative(who PersonalName) bool
	Gender(who PersonalName) Gender
	LooksEqual(name PersonalName) bool
}

func NewPersonalName(style NameStyle, name string) (PersonalName, error) {
	return nil, errors.New("unimplemented")
}
