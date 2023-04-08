package browser

import (
	"regexp"
	"strings"
)

// GOogle Chrome(ium?) Detection

type Chrome struct {
	id        string
	name      string
	userAgent string
	browser   *Browser
}

// NewChromeMatcher creates a new Chrome matcher.
func newChromeMatcher(userAgent string, b *Browser) *Chrome {
	return &Chrome{
		id:        "chrome",
		name:      "Chrome",
		userAgent: userAgent,
		browser:   b,
	}
}

// Matcher name
func (c *Chrome) Name() string {
	return c.name
}

// Matcher ID
func (c *Chrome) ID() string {
	return c.id
}

// Matches returns true if the userAgent is Chrome.
func (c *Chrome) Matches() bool {
	re := regexp.MustCompile(`Chrome|CriOS`)

	if len(re.FindAllStringSubmatch(c.userAgent, -1)) > 0 {
		if containsAny(c.userAgent, "PhantomJS", "FxiOS", "ArchiveBot") {
			return false
		}

		if !c.browser.IsOpera() &&
			!c.browser.IsEdge() &&
			!c.browser.IsDuckDuckGo() &&
			!c.browser.IsYandex() &&
			!c.browser.IsSputnik() &&
			!c.browser.IsSamsungBrowser() &&
			!c.browser.IsHuaweiBrowser() &&
			!c.browser.IsMiuiBrowser() &&
			!c.browser.IsMaxthon() &&
			!c.browser.IsQQ() &&
			!c.browser.IsSougouBrowser() &&
			!c.browser.IsGoogleSearchApp() {
			return true
		}
	}

	return false
}

// FullVersion returns the full version of the browser.
func (c *Chrome) FullVersion() string {
	sources := []string{
		`Chrome/([\d.]+)`,
		`CriOS/([\d.]+)`,
		`Safari/([\d.]+)`,
		`AppleWebKit/([\d.]+)`,
	}

	expressions := []regexp.Regexp{}

	for _, expression := range sources {
		expressions = append(expressions, *regexp.MustCompile(expression))
	}

	for _, re := range expressions {
		matches := re.FindStringSubmatch(c.userAgent)
		if len(matches) > 0 {
			if strings.TrimSpace(matches[1]) != "" {
				return matches[1]
			}
		}
	}

	return ""
}
