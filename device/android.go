package device

import (
	"regexp"
	"strings"
)

type Android struct {
	id        string
	name      string
	userAgent string
}

// NewAndroid creates a new Android device matcher.
// In keeping with the Ruby gem, we set the id to unknown
func NewAndroid(userAgent string) *Android {
	return &Android{
		id:        "unknown",
		name:      getName(userAgent),
		userAgent: userAgent,
	}
}

// getName parses the device name from the userAgent.
func getName(userAgent string) string {
	re := regexp.MustCompile(`/\(Linux.*?; Android.*?; ([-_a-z0-9 ]+) Build[^)]+\)/i`)
	match := re.FindString(userAgent)
	if strings.TrimSpace(match) != "" {
		return match
	}

	return "Unknown"
}

func (a *Android) ID() string {
	return a.id
}

func (a *Android) Name() string {
	return a.name
}

// Matches returns true if this is an Android device.
func (a *Android) Matches() bool {
	return strings.Contains(a.userAgent, "Android")
}
