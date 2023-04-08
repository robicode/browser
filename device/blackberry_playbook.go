package device

import "regexp"

type BlackberryPlaybook struct {
	id        string
	name      string
	userAgent string
}

func NewBlackberryPlaybook(userAgent string) *BlackberryPlaybook {
	return &BlackberryPlaybook{
		id:        "playbook",
		name:      "Blackberry Playbook",
		userAgent: userAgent,
	}
}

func (bbpb *BlackberryPlaybook) Name() string {
	return bbpb.name
}

func (bbpb *BlackberryPlaybook) ID() string {
	return bbpb.id
}

func (bbpb *BlackberryPlaybook) Matches() bool {
	re := regexp.MustCompile(`/PlayBook.*?RIM Tablet/`)
	return re.MatchString(bbpb.userAgent)
}
