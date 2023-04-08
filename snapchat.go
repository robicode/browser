package browser

import (
	"regexp"
	"strings"
)

type Snapchat struct {
	id   string
	name string
	ua   string
}

func newSnapchat(userAgent string) *Snapchat {
	return &Snapchat{
		id:   "snapchat",
		name: "Snapchat",
		ua:   userAgent,
	}
}

func (s *Snapchat) ID() string {
	return s.id
}

func (s *Snapchat) Name() string {
	return s.name
}

func (s *Snapchat) Matches() bool {
	return strings.Contains(s.ua, "Snapchat")
}

func (s *Snapchat) FullVersion() string {
	matches := regexp.MustCompile(`Snapchat( ?|/)([\d.]+)`).FindStringSubmatch(s.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
