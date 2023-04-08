package browser

import (
	"regexp"
	"strings"
)

type PhantomJS struct {
	id   string
	name string
	ua   string
}

func newPhantomJS(userAgent string) *PhantomJS {
	return &PhantomJS{
		id:   "phantom_js",
		name: "PhantomJS",
		ua:   userAgent,
	}
}

func (o *PhantomJS) ID() string {
	return o.id
}

func (o *PhantomJS) Name() string {
	return o.name
}

func (o *PhantomJS) Matches() bool {
	return strings.Contains(o.ua, "PhantomJS")
}

func (o *PhantomJS) FullVersion() string {
	matches := regexp.MustCompile(`PhantomJS/([\d.]+)`).FindStringSubmatch(o.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
