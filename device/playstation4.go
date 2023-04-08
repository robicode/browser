package device

import "strings"

type Playstation4 struct {
	id        string
	name      string
	userAgent string
}

func NewPlaystation4(userAgent string) *Playstation4 {
	return &Playstation4{
		id:        "ps4",
		name:      "Playstation 4",
		userAgent: userAgent,
	}
}

func (p *Playstation4) ID() string {
	return p.id
}

func (p *Playstation4) Name() string {
	return p.name
}

func (p *Playstation4) Matches() bool {
	return strings.Contains(p.userAgent, "PLAYSTATION 4")
}
