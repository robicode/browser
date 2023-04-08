package browser

import (
	"regexp"
	"strings"
)

type Maxthon struct {
	id        string
	name      string
	userAgent string
}

func newMaxthon(userAgent string) *Maxthon {
	return &Maxthon{
		id:        "maxthon",
		name:      "Maxthon",
		userAgent: userAgent,
	}
}

func (m *Maxthon) ID() string {
	return m.id
}

func (m *Maxthon) Name() string {
	return m.name
}

func (m *Maxthon) Matches() bool {
	return strings.Contains(strings.ToLower(m.userAgent), strings.ToLower("Maxthon"))
}

func (m *Maxthon) FullVersion() string {
	matches := regexp.MustCompile(strings.ToLower(`(?:Maxthon)/([\d.]+)`)).FindStringSubmatch(strings.ToLower(m.userAgent))
	if len(matches) > 0 {
		return matches[1]
	} else {
		return "0.0"
	}
}
