package device

import "strings"

type IPodTouch struct {
	id        string
	name      string
	userAgent string
}

func NewIPodTouch(userAgent string) *IPodTouch {
	return &IPodTouch{
		id:        "ipod_touch",
		name:      "iPod Touch",
		userAgent: userAgent,
	}
}

func (i *IPodTouch) ID() string {
	return i.id
}

func (i *IPodTouch) Name() string {
	return i.name
}

func (i *IPodTouch) Matches() bool {
	return strings.Contains(i.userAgent, "iPod")
}
