package browser

import (
	"regexp"
	"strings"
)

type Sputnik struct {
	id   string
	name string
	ua   string
}

func newSputnik(userAgent string) *Sputnik {
	return &Sputnik{
		id:   "sputnik",
		name: "Sputnik",
		ua:   userAgent,
	}
}

func (s *Sputnik) ID() string {
	return s.id
}

func (s *Sputnik) Name() string {
	return s.name
}

func (s *Sputnik) Matches() bool {
	return strings.Contains(s.ua, "SputnikBrowser")
}

func (s *Sputnik) FullVersion() string {
	matches := regexp.MustCompile(`SputnikBrowser/([\d.]+)`).FindStringSubmatch(s.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
