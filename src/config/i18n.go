package config

import (
	"fmt"
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

func Locale(langKey string, values ...interface{}) string {
	currentLocale := Config().Get("", "locale").Value()
	onceLocale.Do(func() {
		localeInstance = new(_Locale)
		localeInstance.locale = currentLocale
		items := Config().GetKeys(configSection)
		localeInstance.localeMaps = make(map[string]*_Config)
		for _, item := range items {
			langPath := Config().runPath + Config().Get(configSection, item.Value()).Name()
			localeInstance.localeMaps[item.Name()] = new(_Config).LoadConfFromFile(langPath)
		}
	})
	result := langKey
	if _, ok := localeInstance.localeMaps[localeInstance.locale]; ok {
		result = localeInstance.localeMaps[localeInstance.locale].Get("", langKey).Value()
	}
	return fmt.Sprintf(result, values...)
}
