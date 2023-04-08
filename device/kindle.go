package device

import "strings"

type Kindle struct {
	id        string
	name      string
	userAgent string
}

func NewKindle(userAgent string) *Kindle {
	return &Kindle{
		id:        "kindle",
		name:      "Kindle",
		userAgent: userAgent,
	}
}

func (k *Kindle) ID() string {
	return k.id
}

func (k *Kindle) Name() string {
	return k.name
}

func (k *Kindle) Matches() bool {
	return strings.Contains(k.userAgent, "Kindle")
}
