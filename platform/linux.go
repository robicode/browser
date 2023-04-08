package platform

import "strings"

type Linux struct {
	id        string
	name      string
	userAgent string
}

func NewLinux(userAgent string) *Linux {
	return &Linux{
		id:        "linux",
		name:      "Linux",
		userAgent: userAgent,
	}
}

func (l *Linux) ID() string {
	return l.id
}

func (l *Linux) Name() string {
	return l.name
}

func (l *Linux) Matches() bool {
	return strings.Contains(l.userAgent, "Linux")
}

func (l *Linux) Version() string {
	return "0"
}
