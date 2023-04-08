package browser

import "regexp"

type Edge struct {
	id        string
	name      string
	userAgent string
}

func newEdge(userAgent string) *Edge {
	return &Edge{
		id:        "edge",
		name:      "Microsoft Edge",
		userAgent: userAgent,
	}
}

func (e *Edge) ID() string {
	return e.id
}

func (e *Edge) Name() string {
	return e.name
}

func (e *Edge) Matches() bool {
	return regexp.MustCompile(`((?:Edge|Edg|EdgiOS|EdgA)/[\d.]+|Trident/8)`).MatchString(e.userAgent)
}

func (e *Edge) FullVersion() string {
	v := regexp.MustCompile(`(?:Edge|Edg|EdgiOS|EdgA)/([\d.]+)`).FindStringSubmatch(e.userAgent)
	if len(v) > 0 {
		return v[1]
	} else {
		return "0.0"
	}
}

func (e *Edge) IsChromeBased() bool {
	return e.Matches() && regexp.MustCompile(`\bEdg\b`).MatchString(e.userAgent)
}
