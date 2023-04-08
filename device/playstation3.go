package device

import "strings"

type Playstation3 struct {
	id        string
	name      string
	userAgent string
}

func NewPlaystation3(userAgent string) *Playstation3 {
	return &Playstation3{
		id:        "ps3",
		name:      "Playstation 3",
		userAgent: userAgent,
	}
}

func (p *Playstation3) ID() string {
	return p.id
}

func (p *Playstation3) Name() string {
	return p.name
}

func (p *Playstation3) Matches() bool {
	return strings.Contains(p.userAgent, "PLAYSTATION 3")
}
