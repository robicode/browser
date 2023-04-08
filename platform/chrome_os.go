package platform

import (
	"regexp"
	"strings"
)

type ChromeOS struct {
	id        string
	name      string
	userAgent string
}

func NewChromeOS(userAgent string) *ChromeOS {
	return &ChromeOS{
		id:        "chrome_os",
		name:      "Chrome OS",
		userAgent: userAgent,
	}
}

func (c *ChromeOS) ID() string {
	return c.id
}

func (c *ChromeOS) Name() string {
	return c.name
}

func (c *ChromeOS) Matches() bool {
	return strings.Contains(c.userAgent, "CrOS")
}

func (c *ChromeOS) Version() string {
	re := regexp.MustCompile(`/CrOS(?: x86_64)? ([\d.]+)/`)

	matches := re.FindAllStringSubmatch(c.userAgent, 1)
	if len(matches) > 1 {
		return matches[0][1]
	}

	return ""
}
