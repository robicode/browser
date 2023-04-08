package browser

import (
	"regexp"
	"strings"
)

type UCBrowser struct {
	id   string
	name string
	ua   string
}

func newUCBrowser(userAgent string) *UCBrowser {
	return &UCBrowser{
		id:   "uc_browser",
		name: "UCBrowser",
		ua:   userAgent,
	}
}

func (u *UCBrowser) ID() string {
	return u.id
}

func (u *UCBrowser) Name() string {
	return u.name
}

func (u *UCBrowser) Matches() bool {
	return strings.Contains(u.ua, "UCBrowser")
}

func (u *UCBrowser) FullVersion() string {
	matches := regexp.MustCompile(`UCBrowser/([\d.]+)`).FindStringSubmatch(u.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
