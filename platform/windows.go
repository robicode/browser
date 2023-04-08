package platform

import (
	"regexp"
	"strings"
)

type Windows struct {
	id        string
	name      string
	userAgent string
}

func NewWindows(userAgent string) *Windows {
	return &Windows{
		id:        "windows",
		name:      "Windows",
		userAgent: userAgent,
	}
}

func (w *Windows) ID() string {
	return w.id
}

func (w *Windows) Name() string {
	return w.name
}

func (w *Windows) Matches() bool {
	return strings.Contains(w.userAgent, "Windows")
}

func (w *Windows) Version() string {
	exp := regexp.MustCompile(`Windows NT\s*([0-9_.]+)?`)

	matches := exp.FindStringSubmatch(w.userAgent)
	if len(matches) > 0 {
		return matches[0]
	}

	return "0"
}
