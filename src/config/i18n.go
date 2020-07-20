package config

import (
	"sync"
)

const configSection = "i18n"

type _Locale struct {
	locale     string
	localeMaps map[string]*_Config
}

var (
	localeInstance *_Locale
	onceLocale     sync.Once
)

func Locale(lang string) string {
	currentLocale := Config().Get("", "locale")
	onceLocale.Do(func() {
		localeInstance = new(_Locale)
		localeInstance.locale = currentLocale.String()
		keys := Config().GetSectionKeys(configSection)
		localeInstance.localeMaps = make(map[string]*_Config)
		for _, key := range keys {
			langPath := Config().runPath +
				Config().Get(configSection, currentLocale.String()).String()
			localeInstance.localeMaps[key] = new(_Config).LoadConfFromFile(langPath)
		}
	})
	return localeInstance.localeMaps[currentLocale.String()].Get("", lang).String()
}
