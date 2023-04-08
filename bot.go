package browser

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Bot struct {
	bots          map[string]string
	userAgent     string
	matchers      []botMatcher
	exceptions    []string
	matched       botMatcher
	isBot         bool
	searchEngine  bool
	searchEngines map[string]string

	Name string
}

type botMatcher interface {
	Check(userAgent string) bool
	Name() string
}

// NewBot creates a properly initialized Bot struct
func NewBot(dataPath, userAgent string) (*Bot, error) {
	var bot Bot

	agent := strings.ToLower(strings.TrimSpace(userAgent))

	bots, err := loadBots(dataPath)
	if err != nil {
		return nil, err
	}

	exceptions, err := loadBotExceptions(dataPath)
	if err != nil {
		return nil, err
	}

	engines, err := loadSearchEngines(dataPath)
	if err != nil {
		return nil, err
	}

	emptyBotMatcher := NewEmptyUserAgentBotMatcher(bot)
	knownBotsMatcher := NewKnownBotMatcher(bot)
	keywordBotMatcher := NewKeywordBotMatcher(bot)

	matchers := []botMatcher{
		emptyBotMatcher,
		knownBotsMatcher,
		keywordBotMatcher,
	}

	bot.bots = bots
	bot.Name = ""
	bot.matchers = matchers
	bot.userAgent = agent
	bot.exceptions = exceptions
	bot.searchEngines = engines

	return &bot, nil
}

func (b *Bot) IsSearchEngine() bool {
	return b.searchEngine
}

func (b *Bot) IsBot() bool {
	return b.detectBot()
}

func (b *Bot) Why() interface{} {
	return b.matched
}

// loadBots loads the list of known bots from YAML and
// populates the bots in the Bot struct
func loadBots(dataPath string) (map[string]string, error) {
	bots := make(map[string]string)

	yfile := filepath.Join(dataPath, "bots.yml")
	data, err := os.ReadFile(yfile)
	if err != nil {
		return bots, err
	}

	err = yaml.Unmarshal(data, &bots)
	if err != nil {
		return bots, err
	}

	return bots, nil
}

// loadBotExceptions reads a list of bot exceptions from disk
func loadBotExceptions(dataPath string) ([]string, error) {
	var exceptions []string

	exceptionData, err := os.ReadFile(filepath.Join(dataPath, "bot_exceptions.yml"))
	if err != nil {
		return exceptions, err
	}

	err = yaml.Unmarshal(exceptionData, &exceptions)
	if err != nil {
		return exceptions, err
	}

	return exceptions, nil
}

// loadSearchEngines reads a list of known search engines from disk
func loadSearchEngines(dataPath string) (map[string]string, error) {
	engines := make(map[string]string)

	seData, err := os.ReadFile(filepath.Join(dataPath, "search_engines.yml"))
	if err != nil {
		return engines, err
	}

	err = yaml.Unmarshal(seData, &engines)
	if err != nil {
		return engines, err
	}

	return engines, nil
}

// detectBot runs all configured matchers and returns true if
// the UserAgent was found to be a bot AND is not in the exceptions
// list. Otherwise, returns false.
func (b *Bot) detectBot() bool {
	for _, matcher := range b.matchers {
		if matcher.Check(b.userAgent) && !b.isExcepted(b.userAgent) {
			b.matched = matcher
			b.isBot = true
			b.Name = matcher.Name()

			b.detectSearchEngine()

			return true
		}
	}

	return false
}

// isExcepted checks to see if the given UserAgent is in the exceptions list.
func (b *Bot) isExcepted(userAgent string) bool {
	for _, exception := range b.exceptions {
		if strings.Contains(userAgent, exception) {
			return true
		}
	}

	return false
}

// Checks to see if the bot's userAgent is a known search engine
// and sets values on the bot appropriately.
func (b *Bot) detectSearchEngine() {
	for key, value := range b.searchEngines {
		if strings.Contains(b.userAgent, key) {
			b.searchEngine = true
			b.Name = value
		}
	}
}
