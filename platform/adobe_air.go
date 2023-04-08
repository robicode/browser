package platform

import (
	"regexp"
	"strings"
)

type AdobeAir struct {
	id        string
	name      string
	userAgent string
}

func NewAdobeAir(userAgent string) *AdobeAir {
	return &AdobeAir{
		id:        "adobe_air",
		name:      "Adobe AIR",
		userAgent: userAgent,
	}
}

func (a *AdobeAir) ID() string {
	return a.id
}

func (a *AdobeAir) Name() string {
	return a.name
}

func (a *AdobeAir) Version() string {
	re := regexp.MustCompile(`AdobeAIR/([\d.]+)`)

	matches := re.FindAllStringSubmatch(a.userAgent, 1)
	if len(matches) > 0 {
		return matches[0][1]
	}

	return ""
}

func (a *AdobeAir) Matches() bool {
	return strings.Contains(a.userAgent, "AdobeAIR")
}
