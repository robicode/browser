package browser

import (
	"regexp"
	"strings"
)

type Firefox struct {
	id        string
	name      string
	userAgent string
}

func newFirefox(userAgent string) *Firefox {
	return &Firefox{
		id:        "firefox",
		name:      "Firefox",
		userAgent: userAgent,
	}
}

func (f *Firefox) ID() string {
	return f.id
}

func (f *Firefox) Name() string {
	return f.name
}

func (f *Firefox) Matches() bool {
	return strings.Contains(f.userAgent, "Firefox") || strings.Contains(f.userAgent, "FxiOS")
}

func (f *Firefox) FullVersion() string {
	v := regexp.MustCompile(`(?:Firefox|FxiOS)/([\d.]+)`).FindStringSubmatch(f.userAgent)
	if len(v) > 0 {
		return v[1]
	} else {
		return "0.0"
	}
}
