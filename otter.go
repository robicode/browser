package browser

import (
	"regexp"
	"strings"
)

type Otter struct {
	id   string
	name string
	ua   string
}

func newOtter(userAgent string) *Otter {
	return &Otter{
		id:   "otter",
		name: "Otter",
		ua:   userAgent,
	}
}

func (o *Otter) ID() string {
	return o.id
}

func (o *Otter) Name() string {
	return o.name
}

func (o *Otter) Matches() bool {
	return strings.Contains(o.ua, "Otter")
}

func (o *Otter) FullVersion() string {
	matches := regexp.MustCompile(`Otter/([\d.]+)`).FindStringSubmatch(o.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
