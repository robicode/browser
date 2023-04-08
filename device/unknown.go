package device

type Unknown struct {
	id        string
	name      string
	userAgent string
}

func NewUnknown(userAgent string) *Unknown {
	return &Unknown{
		id:        "unknown_device",
		name:      "Unknown",
		userAgent: userAgent,
	}
}

func (u *Unknown) ID() string {
	return u.id
}

func (u *Unknown) Name() string {
	return u.name
}

func (u *Unknown) Matches() bool {
	return true
}
