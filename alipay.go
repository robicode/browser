package browser

import (
	"regexp"
	"strings"
)

type AlipayMatcher struct {
	id   string
	name string

	userAgent string
}

// NewAlipayMatcher returns a new matcher instance.
func newAlipayMatcher(userAgent string) *AlipayMatcher {
	return &AlipayMatcher{
		userAgent: userAgent,
		id:        "alipay",
		name:      "Alipay",
	}
}

// Matches returns true if this client is AlipayClient
func (a *AlipayMatcher) Matches() bool {
	return strings.Contains(a.userAgent, "AlipayClient")
}

// FullVersion returns the full version of the client.
func (a *AlipayMatcher) FullVersion() string {
	re := regexp.MustCompile(`(?:AlipayClient)/([\\d.]+)`)

	matches := re.FindStringSubmatch(a.userAgent)

	if len(matches) > 0 {
		return matches[1]
	}

	return "0.0"
}

// Name returns the friendly name of the browser
func (a *AlipayMatcher) Name() string {
	return a.name
}

// ID returns a code-friendly ID for the browser
func (a *AlipayMatcher) ID() string {
	return a.id
}
