package device

import "strings"

type PSP struct {
	id        string
	name      string
	userAgent string
}

func NewPSP(userAgent string) *PSP {
	return &PSP{
		id:        "psp",
		name:      "PlayStation Portable",
		userAgent: userAgent,
	}
}

func (p *PSP) ID() string {
	return p.id
}

func (p *PSP) Name() string {
	return p.name
}

func (p *PSP) Matches() bool {
	return strings.Contains(p.userAgent, "PlayStation Portable")
}
