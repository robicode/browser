package device

import "strings"

type IPhone struct {
	id        string
	name      string
	userAgent string
}

func NewIPhone(userAgent string) *IPhone {
	return &IPhone{
		id:        "iphone",
		name:      "iPhone",
		userAgent: userAgent,
	}
}

func (i *IPhone) ID() string {
	return i.id
}

func (i IPhone) Name() string {
	return i.name
}

func (i *IPhone) Matches() bool {
	return strings.Contains(i.userAgent, "iPhone")
}
