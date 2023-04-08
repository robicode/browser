package browser

import (
	"regexp"
	"strings"
)

type Nokia struct {
	id        string
	name      string
	userAgent string
}

func newNokia(userAgent string) *Nokia {
	return &Nokia{
		id:        "nokia",
		name:      "Nokia S40 Ovi Browser",
		userAgent: userAgent,
	}
}

func (n *Nokia) ID() string {
	return n.id
}

func (n *Nokia) Name() string {
	return n.name
}

func (n *Nokia) Matches() bool {
	return strings.Contains(n.userAgent, "S40OviBrowser")
}

func (n *Nokia) FullVersion() string {
	matches := regexp.MustCompile(`S40OviBrowser/([\d.]+)`).FindStringSubmatch(n.userAgent)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
