package device

import "strings"

type Xbox360 struct {
	id        string
	name      string
	userAgent string
}

func NewXbox360(userAgent string) *Xbox360 {
	return &Xbox360{
		id:        "xbox_360",
		name:      "Xbox 360",
		userAgent: userAgent,
	}
}

func (x *Xbox360) ID() string {
	return x.id
}

func (x *Xbox360) Name() string {
	return x.name
}

func (x *Xbox360) Matches() bool {
	return strings.Contains(x.userAgent, "Xbox")
}
