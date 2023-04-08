package browser

import (
	"regexp"
	"strings"
)

type Electron struct {
	id        string
	name      string
	userAgent string
}

func newElectron(userAgent string) *Electron {
	return &Electron{
		id:        "electron",
		name:      "Electron",
		userAgent: userAgent,
	}
}

func (e *Electron) ID() string {
	return e.id
}

func (e *Electron) Name() string {
	return e.name
}

func (e *Electron) Matches() bool {
	return strings.Contains(e.userAgent, "Electron")
}

func (e *Electron) FullVersion() string {
	v := regexp.MustCompile(`Electron/([\d.]+)`).FindStringSubmatch(e.userAgent)
	if len(v) > 0 {
		return v[1]
	} else {
		return "0.0"
	}
}
