package browser

import (
	"regexp"
	"strings"
)

type MicroMessenger struct {
	id        string
	name      string
	userAgent string
}

func newMicroMessenger(userAgent string) *MicroMessenger {
	return &MicroMessenger{
		id:        "micro_messenger",
		name:      "MicroMessenger",
		userAgent: userAgent,
	}
}

func (mm *MicroMessenger) ID() string {
	return mm.id
}

func (mm *MicroMessenger) Name() string {
	return mm.name
}

func (mm *MicroMessenger) Matches() bool {
	return strings.Contains(strings.ToLower(mm.userAgent), strings.ToLower("MicroMessenger"))
}

func (mm *MicroMessenger) FullVersion() string {
	matches := regexp.MustCompile(`(?:MicroMessenger)/([\d.]+)`).FindStringSubmatch(mm.userAgent)
	if len(matches) > 0 {
		return matches[1]
	} else {
		return "0.0"
	}
}
