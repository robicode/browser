package browser

import (
	"regexp"
	"strings"
)

type Safari struct {
	id      string
	name    string
	ua      string
	browser *Browser
}

func newSafari(userAgent string, b *Browser) *Safari {
	return &Safari{
		id:      "safari",
		name:    "Safari",
		ua:      userAgent,
		browser: b,
	}
}

func (s *Safari) ID() string {
	return s.id
}

func (s *Safari) Name() string {
	return s.name
}

func (s *Safari) Matches() bool {
	return strings.Contains(s.ua, "Safari") && !strings.Contains(s.ua, "PhantomJS") &&
		!strings.Contains(s.ua, "FxiOS") &&
		!s.browser.IsEdge() &&
		!s.browser.IsChrome() &&
		!s.browser.IsOpera() &&
		!s.browser.IsSamsungBrowser() &&
		!s.browser.IsHuaweiBrowser() &&
		!s.browser.IsMiuiBrowser() &&
		!s.browser.IsDuckDuckGo() &&
		!s.browser.IsYandex() &&
		!s.browser.IsSputnik() &&
		!s.browser.IsMaxthon() &&
		!s.browser.IsQQ() &&
		!s.browser.IsAlipay() &&
		!s.browser.IsSougouBrowser() &&
		!s.browser.IsGoogleSearchApp()
}

func (s *Safari) FullVersion() string {
	matches := regexp.MustCompile(``).FindStringSubmatch(s.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
