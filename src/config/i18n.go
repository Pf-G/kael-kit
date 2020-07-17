package config

import (
	"sync"
)

type _Locale struct {
	localeMaps map[string]_Config
}

var localeInstance *_Locale
var onceLocale sync.Once

func Locale() *_Locale {
	onceLocale.Do(func() {
		localeInstance = new(_Locale)
		//currentLocale := Config().Get("", "locale")
		//keys = Config().GetSectionKeys("i18n")
		//for _, key := range keys {
		//
		//}

	})
	return localeInstance
}