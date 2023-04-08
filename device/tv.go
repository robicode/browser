package device

import "regexp"

type TV struct {
	id        string
	name      string
	userAgent string
}

func NewTV(userAgent string) *TV {
	return &TV{
		id:        "tv",
		name:      "TV",
		userAgent: userAgent,
	}
}

func (t *TV) ID() string {
	return t.id
}

func (t *TV) Name() string {
	return t.name
}

func (t *TV) Matches() bool {
	re := regexp.MustCompile(`/(\btv|Android.*?ADT-1|Nexus Player)/i`)

	return re.MatchString(t.userAgent)
}
