package browser

import (
	"regexp"
	"strings"
)

type Instagram struct {
	id        string
	name      string
	userAgent string
}

func newInstagram(userAgent string) *Instagram {
	return &Instagram{
		id:        "instagram",
		name:      "Instagram",
		userAgent: userAgent,
	}
}

func (i *Instagram) ID() string {
	return i.id
}

func (i *Instagram) Name() string {
	return i.name
}

func (i *Instagram) Matches() bool {
	return strings.Contains(i.userAgent, "Instagram")
}

func (i *Instagram) FullVersion() string {
	v := regexp.MustCompile(`Instagram[ /]([\d.]+)`).FindStringSubmatch(i.userAgent)
	if len(v) > 0 {
		return v[1]
	} else {
		return "0.0"
	}
}
