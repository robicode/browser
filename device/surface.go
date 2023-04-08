package device

import "strings"

type Surface struct {
	id        string
	name      string
	userAgent string
}

func NewSurface(userAgent string) *Surface {
	return &Surface{
		id:        "surface",
		name:      "Microsoft Surface",
		userAgent: userAgent,
	}
}

func (s *Surface) ID() string {
	return s.id
}

func (s *Surface) Name() string {
	return s.name
}

func (s *Surface) Matches() bool {
	// TODO Include platform.IsWindowsRT()
	return strings.Contains(s.userAgent, "Touch")
}
