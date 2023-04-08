package browser

import (
	"regexp"
	"strings"
)

type HuaweiBrowser struct {
	id        string
	name      string
	userAgent string
}

func newHuaweiBrowser(userAgent string) *HuaweiBrowser {
	return &HuaweiBrowser{
		id:        "huawei_browser",
		name:      "Huawei Browser",
		userAgent: userAgent,
	}
}

func (h *HuaweiBrowser) ID() string {
	return h.id
}

func (h *HuaweiBrowser) Name() string {
	return h.name
}

func (h *HuaweiBrowser) FullVersion() string {
	if len(regexp.MustCompile(`(?:HuaweiBrowser)/([\d.]+)`).FindStringSubmatch(h.userAgent)) > 0 {
		return regexp.MustCompile(`(?:HuaweiBrowser)/([\d.]+)`).FindStringSubmatch(h.userAgent)[1]
	} else {
		return "0.0"
	}
}

func (h *HuaweiBrowser) Matches() bool {
	return strings.Contains(strings.ToLower(h.userAgent), strings.ToLower("HuaweiBrowser"))
}
