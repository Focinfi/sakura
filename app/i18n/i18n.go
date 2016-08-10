package i18n

import (
	"io/ioutil"
	"path"
	"strings"

	"coding.net/focinfi/90s/config"

	"gopkg.in/yaml.v2"
)

// Locale for different locales
type Locale string

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

// T translates into the given local language with the given key
func T(key string, locale Locale) string {
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

	return t
}

func init() {
	for _, l := range AllowedLocales {
		b, _ := ioutil.ReadFile(path.Join(config.ConfigRoot, string(l)+".yml"))
		t := map[string]string{}
		yaml.Unmarshal(b, &t)
		Translations[l] = t
	}
}
