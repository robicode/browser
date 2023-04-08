package browser

import (
	"regexp"
	"strings"
)

type Facebook struct {
	id        string
	name      string
	userAgent string
}

func newFacebook(userAgent string) *Facebook {
	return &Facebook{
		id:        "facebook",
		name:      "Facebook",
		userAgent: userAgent,
	}
}

func (f *Facebook) ID() string {
	return f.id
}

func (f *Facebook) Name() string {
	return f.name
}

func (f *Facebook) Matches() bool {
	return strings.Contains(f.userAgent, "FBAV") || strings.Contains(f.userAgent, "FBAN")
}

func (f *Facebook) FullVersion() string {
	v1 := regexp.MustCompile(`FBAV/([\d.]+)`).FindStringSubmatch(f.userAgent)
	v2 := regexp.MustCompile(``).FindStringSubmatch(f.userAgent)

	if len(v1) > 0 {
		return v1[1]
	} else if len(v2) > 0 {
		return v2[1]
	} else {
		return "0.0"
	}
}
