package browser

import (
	"regexp"
	"strings"
)

type Opera struct {
	id        string
	name      string
	userAgent string
}

func newOpera(userAgent string) *Opera {
	return &Opera{
		id:        "opera",
		name:      "Opera",
		userAgent: userAgent,
	}
}

func (o *Opera) ID() string {
	return o.id
}

func (o *Opera) Name() string {
	return o.name
}

func (o *Opera) Matches() bool {
	return strings.Contains(o.userAgent, "Opera") || strings.Contains(o.userAgent, "OPR")
}

func (o *Opera) FullVersion() string {
	matches := regexp.MustCompile(`OPR/([\d.]+)`).FindStringSubmatch(o.userAgent)
	if len(matches) > 0 {
		return matches[1]
	}

	matches = regexp.MustCompile(`Version/([\d.]+)`).FindStringSubmatch(o.userAgent)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
