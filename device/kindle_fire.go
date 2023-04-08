package device

import (
	"regexp"
)

type KindleFire struct {
	id        string
	name      string
	userAgent string
}

func NewKindleFire(userAgent string) *KindleFire {
	return &KindleFire{
		id:        "kindle_fire",
		name:      "Kindle Fire",
		userAgent: userAgent,
	}
}

func (k *KindleFire) ID() string {
	return k.id
}

func (k *KindleFire) Name() string {
	return k.name
}

func (k *KindleFire) Matches() bool {
	return regexp.MustCompile(`/Kindle Fire|KFTT/`).MatchString(k.userAgent)
}
