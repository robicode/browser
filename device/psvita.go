package device

import "strings"

type PSVita struct {
	id        string
	name      string
	userAgent string
}

func NewPSVita(userAgent string) *PSVita {
	return &PSVita{
		id:        "psvita",
		name:      "PlayStation Vita",
		userAgent: userAgent,
	}
}

func (p *PSVita) ID() string {
	return p.id
}

func (p *PSVita) Name() string {
	return p.name
}

func (p *PSVita) Matches() bool {
	return strings.Contains(p.userAgent, "PlayStation Vita")
}
