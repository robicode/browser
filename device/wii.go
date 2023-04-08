package device

import "strings"

type Wii struct {
	id        string
	name      string
	userAgent string
}

func NewWii(userAgent string) *Wii {
	return &Wii{
		id:        "wii",
		name:      "Nintendo Wii",
		userAgent: userAgent,
	}
}

func (w *Wii) ID() string {
	return w.id
}

func (w *Wii) Name() string {
	return w.name
}

func (w *Wii) Matches() bool {
	return strings.Contains(w.userAgent, "Nintendo Wii")
}
