package platform

import (
	"regexp"
	"strings"
)

type KaiOS struct {
	id        string
	name      string
	userAgent string
}

func NewKaiOS(userAgent string) *KaiOS {
	return &KaiOS{
		id:        "kai_os",
		name:      "KaiOS",
		userAgent: userAgent,
	}
}

func (k *KaiOS) ID() string {
	return k.id
}

func (k *KaiOS) Name() string {
	return k.name
}

func (k *KaiOS) Matches() bool {
	return strings.Contains(k.userAgent, "KAIOS")
}

func (k *KaiOS) Version() string {
	re := regexp.MustCompile(`KAIOS/([\d.]+)`)

	matches := re.FindAllStringSubmatch(k.userAgent, -1)
	if len(matches) > 0 {
		return matches[0][1]
	}

	return ""
}
