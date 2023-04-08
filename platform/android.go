package platform

import (
	"regexp"
	"strings"
)

type Android struct {
	id        string
	name      string
	userAgent string
}

func NewAndroid(userAgent string) *Android {
	return &Android{
		id:        "android",
		name:      "Android",
		userAgent: userAgent,
	}
}

func (a *Android) ID() string {
	return a.id
}

func (a *Android) Name() string {
	return a.name
}

func (a *Android) Matches() bool {
	return strings.Contains(a.userAgent, "Android") && !strings.Contains(a.userAgent, "KAIOS")
}

func (a *Android) Version() string {
	re := regexp.MustCompile(`/Android ([\d.]+)/`)

	matches := re.FindAllStringSubmatch(a.userAgent, 1)
	if len(matches) > 0 {
		return matches[0][1]
	}

	return ""
}
