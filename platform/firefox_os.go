package platform

import (
	"regexp"
	"strings"
)

type FirefoxOS struct {
	id        string
	name      string
	userAgent string
}

func NewFirefoxOS(userAgent string) *FirefoxOS {
	return &FirefoxOS{
		id:        "firefox_os",
		name:      "Firefox OS",
		userAgent: userAgent,
	}
}

func (f *FirefoxOS) ID() string {
	return f.id
}

func (f *FirefoxOS) Name() string {
	return f.name
}

func (f *FirefoxOS) Matches() bool {
	re := regexp.MustCompile(`(Android|Linux|BlackBerry|Windows|Mac)`)
	return !re.MatchString(f.userAgent) && strings.Contains(f.userAgent, "Firefox")
}

func (f *FirefoxOS) Version() string {
	return "0"
}
