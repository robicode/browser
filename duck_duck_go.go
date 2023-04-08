package browser

import (
	"regexp"
	"strings"
)

type DuckDuckGo struct {
	id        string
	name      string
	userAgent string
}

func newDuckDuckGo(userAgent string) *DuckDuckGo {
	return &DuckDuckGo{
		id:        "duckduckgo",
		name:      "DuckDuckGo",
		userAgent: userAgent,
	}
}

func (d *DuckDuckGo) ID() string {
	return d.id
}

func (d *DuckDuckGo) Name() string {
	return d.name
}

func (d *DuckDuckGo) Matches() bool {
	return strings.Contains(d.userAgent, "DuckDuckGo")
}

func (d *DuckDuckGo) FullVersion() string {
	v := regexp.MustCompile(`DuckDuckGo/([\d.]+)`).FindStringSubmatch(d.userAgent)
	if len(v) > 0 {
		return v[1]
	} else {
		return "0.0"
	}
}
