package platform

import (
	"regexp"
	"strings"

	"github.com/robicode/browser/util"
)

type Mac struct {
	id        string
	userAgent string
}

func NewMac(userAgent string) *Mac {
	return &Mac{
		id:        "mac",
		userAgent: userAgent,
	}
}

func (m *Mac) ID() string {
	return m.id
}

func (m *Mac) Matches() bool {
	return strings.Contains(m.userAgent, "Mac")
}

func (m *Mac) Version() string {
	re := regexp.MustCompile(`/Mac OS X\s*([0-9_.]+)?/`)
	match := re.FindString(m.userAgent)

	if match != "" {
		return strings.ReplaceAll(match, "_", ".")
	}

	return ""
}

func (m *Mac) Name() string {
	if m.MacV(">= 10.12") {
		return "macOS"
	} else {
		return "Mac OS X"
	}
}

func (m *Mac) MacV(expected string) bool {
	return util.DetectVersion(m.Version(), expected)
}
