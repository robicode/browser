package browser

import (
	"regexp"
	"strings"
)

type MiuiBrowser struct {
	id        string
	name      string
	userAgent string
}

func newMiuiBrowser(userAgent string) *MiuiBrowser {
	return &MiuiBrowser{
		id:        "miui_browser",
		name:      "Miui Browser",
		userAgent: userAgent,
	}
}

func (mb *MiuiBrowser) ID() string {
	return mb.id
}

func (mb *MiuiBrowser) Name() string {
	return mb.name
}

func (mb *MiuiBrowser) Matches() bool {
	return strings.Contains(mb.userAgent, "MiuiBrowser")
}

func (mb *MiuiBrowser) FullVersion() string {
	matches := regexp.MustCompile(`MiuiBrowser/([\d.]+)`).FindStringSubmatch(mb.userAgent)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
