package device

import "strings"

type XboxOne struct {
	id        string
	name      string
	userAgent string
}

func NewXboxOne(userAgent string) *Xbox360 {
	return &Xbox360{
		id:        "xbox_one",
		name:      "Xbox One",
		userAgent: userAgent,
	}
}

func (x *XboxOne) ID() string {
	return x.id
}

func (x *XboxOne) Name() string {
	return x.name
}

func (x *XboxOne) Matches() bool {
	return strings.Contains(x.userAgent, "Xbox One")
}
