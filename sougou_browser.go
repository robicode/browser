package browser

import (
	"regexp"
	"strings"
)

type SougouBrowser struct {
	id   string
	name string
	ua   string
}

func newSougouBrowser(userAgent string) *SougouBrowser {
	return &SougouBrowser{
		id:   "sougou_browser",
		name: "Sougou Browser",
		ua:   userAgent,
	}
}

func (s *SougouBrowser) ID() string {
	return s.id
}

func (s *SougouBrowser) Name() string {
	return s.name
}

func (s *SougouBrowser) Matches() bool {
	matches, err := regexp.MatchString(`\bSE\b`, s.ua)
	if err != nil {
		return false
	}

	return strings.Contains(strings.ToLower(s.ua), "SogouMobileBrowser") || matches
}

func (s *SougouBrowser) FullVersion() string {
	matches := regexp.MustCompile(`(?:SogouMobileBrowser)/([\d.]+)`).FindStringSubmatch(s.ua)
	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}
