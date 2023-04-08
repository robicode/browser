package browser

import (
	"strings"
)

var (
	blankBotName   string = "Blank Bot"
	genericBotName string = "Generic Bot"
)

type EmptyUserAgentMatcher struct {
	name string
	bot  Bot
}

type KnownBotsMatcher struct {
	name string
	bot  Bot
}

type KeywordBotMatcher struct {
	bot      Bot
	keywords []string

	name string
}

func NewEmptyUserAgentBotMatcher(_bot Bot) *EmptyUserAgentMatcher {
	return &EmptyUserAgentMatcher{
		name: blankBotName,
		bot:  _bot,
	}
}

func (e *EmptyUserAgentMatcher) Name() string {
	return e.name
}

func (e *EmptyUserAgentMatcher) Check(userAgent string) bool {
	if userAgent == "" {
		e.bot.Name = blankBotName
		return true
	}

	return false
}

func NewKnownBotMatcher(_bot Bot) *KnownBotsMatcher {
	return &KnownBotsMatcher{
		bot: _bot,
	}
}

func (k *KnownBotsMatcher) Check(userAgent string) bool {
	for key, value := range k.bot.bots {
		if strings.Contains(userAgent, key) {
			k.name = value
			return true
		}
	}

	return false
}

func (k *KnownBotsMatcher) Name() string {
	return k.name
}

func NewKeywordBotMatcher(_bot Bot) *KeywordBotMatcher {
	keywords := []string{
		"crawl",
		"fetch",
		"search",
		"monitoring",
		"spider",
		"bot",
	}

	return &KeywordBotMatcher{
		bot:      _bot,
		keywords: keywords,
	}
}

func (k *KeywordBotMatcher) Check(userAgent string) bool {
	for _, value := range k.keywords {
		if strings.Contains(userAgent, value) {
			k.bot.Name = genericBotName
			return true
		}
	}

	return false
}

func (k *KeywordBotMatcher) Name() string {
	return k.name
}
