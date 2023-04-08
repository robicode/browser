package device

import "strings"

type Switch struct {
	id        string
	name      string
	userAgent string
}

func NewSwitch(userAgent string) *Switch {
	return &Switch{
		id:        "switch",
		name:      "Nintendo Switch",
		userAgent: userAgent,
	}
}

func (s *Switch) ID() string {
	return s.id
}

func (s *Switch) Name() string {
	return s.name
}

func (s *Switch) Matches() bool {
	return strings.Contains(s.userAgent, "Nintendo Switch")
}
