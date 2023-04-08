package device

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type Samsung struct {
	id        string
	matches   []string
	names     map[string]string
	name      string
	userAgent string
}

const samsungRegexp string = `/\(Linux.*?; Android.*?; (SAMSUNG )?(SM-[A-Z0-9]+).*?\)/i`

func NewSamsung(userAgent string, dataPath string) *Samsung {
	names, err := loadNames(dataPath)
	if err != nil {
		log.Println("error[Device:Samsung]:", err)
		return nil
	}

	return &Samsung{
		id:        "samsung",
		name:      "Samsung",
		names:     names,
		userAgent: userAgent,
	}
}

func (s *Samsung) ID() string {
	return s.id
}

// loadNames loads a list of known Samsung devices from disk.
func loadNames(dataPath string) (map[string]string, error) {
	names := make(map[string]string)
	path := filepath.Join(dataPath, "samsung.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		return names, err
	}

	err = yaml.Unmarshal(data, &names)
	if err != nil {
		return names, err
	}

	return names, nil
}

func (s *Samsung) Name() string {
	if len(s.matches) > 1 {
		code := s.matches[2]
		name := s.names[code]

		if strings.TrimSpace(name) == "" {
			return s.name + " " + code
		}

		return s.name + " " + name
	}
	return s.name
}

func (s *Samsung) Matches() bool {
	return regexp.MustCompile(samsungRegexp).MatchString(s.userAgent)
}
