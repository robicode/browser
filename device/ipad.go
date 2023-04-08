package device

import "strings"

type IPad struct {
	name      string
	id        string
	userAgent string
}

func NewIPad(userAgent string) *IPad {
	return &IPad{
		id:        "ipad",
		name:      "iPad",
		userAgent: userAgent,
	}
}

func (i *IPad) ID() string {
	return i.id
}

func (i *IPad) Name() string {
	return i.name
}

func (i *IPad) Matches() bool {
	return strings.Contains(i.userAgent, "iPad")
}
