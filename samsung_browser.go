package browser

import (
	"regexp"
	"strings"
)

type SamsungBrowser struct {
	id   string
	name string
	ua   string
}

func newSamsungBrowser(userAgent string) *SamsungBrowser {
	return &SamsungBrowser{
		id:   "samsung_browser",
		name: "Samsung Browser",
		ua:   userAgent,
	}
}

func (s *SamsungBrowser) ID() string {
	return s.id
}

func (s *SamsungBrowser) Name() string {
	return s.name
}

func (s *SamsungBrowser) Matches() bool {
	return strings.Contains(s.ua, "SamsungBrowser")
}

func (s *SamsungBrowser) FullVersion() string {
	matches := regexp.MustCompile(`SamsungBrowser/([\d.]+)`).FindStringSubmatch(s.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
