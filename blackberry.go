package browser

import (
	"regexp"
	"strings"
)

type Blackberry struct {
	id   string
	name string

	userAgent string
}

func newBlackberry(userAgent string) *Blackberry {
	return &Blackberry{
		name:      "BlackBerry",
		id:        "blackberry",
		userAgent: userAgent,
	}
}

func (bb *Blackberry) Name() string {
	return bb.name
}

func (bb *Blackberry) ID() string {
	return bb.id
}

func (bb *Blackberry) Matches() bool {
	if strings.Contains(bb.userAgent, "BlackBerry") || strings.Contains(bb.userAgent, "BB10") {
		return true
	}

	return false
}

func (bb *Blackberry) FullVersion() string {
	re1 := regexp.MustCompile(`BlackBerry[\da-z]+/([\d.]+)`)
	re2 := regexp.MustCompile(`Version/([\d.]+)`)

	matches := re1.FindStringSubmatch(bb.userAgent)
	if len(matches) > 0 {
		return matches[1]
	}

	matches = re2.FindStringSubmatch(bb.userAgent)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
