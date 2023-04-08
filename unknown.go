package browser

import (
	"regexp"
	"strings"
)

type Unknown struct {
	id string
	ua string
}

var names = map[string]string{
	"QuickTime": "QuickTime",
	"CoreMedia": "Apple CoreMedia",
}

func newUnknown(userAgent string) *Unknown {
	return &Unknown{
		id: "unknown_browser",
		ua: userAgent,
	}
}

func (u *Unknown) ID() string {
	return u.id
}

func (u *Unknown) Name() string {
	return u.inferName()
}

func (u *Unknown) Matches() bool {
	return true
}

func (u *Unknown) inferName() string {
	for key, name := range names {
		if strings.Contains(u.ua, key) {
			return name
		}
	}

	return "Unknown Browser"
}

func (u *Unknown) FullVersion() string {
	matches := regexp.MustCompile(`(?:QuickTime)/([\d.]+)`).FindStringSubmatch(u.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	matches = regexp.MustCompile(`CoreMedia v([\d.]+)`).FindStringSubmatch(u.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
