// This is a Tentative matcher for Playstation 5 browser. I don't have one
// so I cannot test it, and the Ruby gem does not yet have a ps5 matcher.
package device

import "strings"

type Playstation5 struct {
	id        string
	name      string
	userAgent string
}

func NewPlaystation5(userAgent string) *Playstation3 {
	return &Playstation3{
		id:        "ps5",
		name:      "Playstation 5",
		userAgent: userAgent,
	}
}

func (p *Playstation5) ID() string {
	return p.id
}

func (p *Playstation5) Name() string {
	return p.name
}

func (p *Playstation5) Matches() bool {
	return strings.Contains(p.userAgent, "PLAYSTATION 5")
}
