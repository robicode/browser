package browser

import (
	"regexp"
	"strings"
)

type GoogleSearchApp struct {
	id        string
	name      string
	userAgent string
	browser   *Browser
}

func newGoogleSearchApp(userAgent string, b *Browser) *GoogleSearchApp {
	return &GoogleSearchApp{
		id:        "google_search_app",
		name:      "Google Search App",
		userAgent: userAgent,
		browser:   b,
	}
}

func (g *GoogleSearchApp) ID() string {
	return g.id
}

func (g *GoogleSearchApp) Name() string {
	return g.name
}

func (g *GoogleSearchApp) Matches() bool {
	return strings.Contains(g.userAgent, "GSA")
}

func (g *GoogleSearchApp) FullVersion() string {
	v := regexp.MustCompile(`GSA/([\d.]+\d)`).FindStringSubmatch(g.userAgent)
	if len(v) > 0 {
		return v[1]
	} else {
		c := newChromeMatcher(g.userAgent, g.browser)
		return c.FullVersion()
	}
}
