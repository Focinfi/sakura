package i18n

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"coding.net/focinfi/90s/config"

	"gopkg.in/yaml.v2"
)

// Locale for different locales
type Locale string

// Translator contains T to translates values
type Translator interface {
	T(key string, values ...interface{}) string
}

type translator struct {
	Locale Locale
}

func (t *translator) T(key string, values ...interface{}) string {
	return T(t.Locale, key, values...)
}

const (
	// EN for English
	EN Locale = "en"
	// ZhCN for simple Chinese
	ZhCN = "zh-CN"
)

// AllowedLocales for all supportted locales
var AllowedLocales = []Locale{EN, ZhCN}

// Translations for all translations
var Translations = map[Locale]map[string]string{}

// NewTranslator allocates and returns a new translator
func NewTranslator(locale Locale) Translator {
	return &translator{Locale: locale}
}

// T translates into the given local language with the given key
func T(locale Locale, key string, values ...interface{}) string {
	matched := false
	for _, l := range AllowedLocales {
		if l == locale {
			matched = true
			break
		}
	}

	if !matched {
		locale = EN
	}

	t := Translations[locale][key]
	if t == "" {
		return strings.Replace(key, "_", " ", -1)
	}

	return fmt.Sprintf(t, values...)
}

func init() {
	for _, l := range AllowedLocales {
		b, _ := ioutil.ReadFile(path.Join(config.ConfigRoot, "translations", string(l)+".yml"))
		t := map[string]string{}
		yaml.Unmarshal(b, &t)
		Translations[l] = t
	}
}
