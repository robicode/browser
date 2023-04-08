package browser

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AcceptLanguage struct {
	languages map[string]string
}

type Language struct {
	Code    string
	Full    string
	Region  string
	Quality float32
}

func NewAcceptLanguage(acceptLanguage, dataPath string) (*AcceptLanguage, error) {
	languages, err := loadLanguages(filepath.Join(dataPath, "languages.yml"))
	if err != nil {
		return nil, err
	}

	return &AcceptLanguage{
		languages: languages,
	}, nil
}

// loadLanguages loads the list of code to region mappings from a YAML file.
func loadLanguages(filename string) (map[string]string, error) {
	languages := make(map[string]string)

	data, err := os.ReadFile(filename)
	if err != nil {
		return languages, err
	}

	err = yaml.Unmarshal(data, &languages)
	if err != nil {
		return languages, err
	}

	return languages, nil
}
