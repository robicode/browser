package platform

type Unknown struct {
	id        string
	name      string
	userAgent string
}

func NewUnknown(userAgent string) *Unknown {
	return &Unknown{
		name:      "Unknown",
		id:        "unknown_platform",
		userAgent: userAgent,
	}
}

func (u *Unknown) Name() string {
	return u.name
}

func (u *Unknown) ID() string {
	return u.id
}

func (u *Unknown) Matches() bool {
	return true
}

func (u *Unknown) Version() string {
	return "0"
}
