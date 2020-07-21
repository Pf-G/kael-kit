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

func (c _Locale) setLocale(locale string) _Locale {
	c.locale = locale
	return c
}

func initLocaleFromINI(){
	onceLocale.Do(func() {
		currentLocale := Config().Get("", "locale").Value()
		localeInstance = new(_Locale)
		localeInstance.locale = currentLocale
		items := Config().GetKeys(configSection)
		localeInstance.localeMaps = make(map[string]*_Config)
		for _, item := range items {
			langPath := Config().runPath + Config().Get(configSection, item.Value()).Name()
			localeInstance.localeMaps[item.Name()] = new(_Config).LoadConfFromFile(langPath)
		}
	})
}

func Locale(txtKey string, values ...interface{}) string {
	initLocaleFromINI()
	result := txtKey
	if _, ok := localeInstance.localeMaps[localeInstance.locale]; ok {
		result = localeInstance.localeMaps[localeInstance.locale].Get("", txtKey).Value()
	}
	for key, value := range values {
		values[key] = Locale(value.(string))
	}
	return fmt.Sprintf(result, values...)
}

func LocaleE(locale string, txtKey string, values ...interface{}) string {
	initLocaleFromINI()
	if locale == "" {
		return Locale(txtKey, values...)
	}
	result := txtKey
	if _, ok := localeInstance.localeMaps[locale]; ok {
		result = localeInstance.localeMaps[locale].Get("", txtKey).Value()
	}
	for key, value := range values {
		values[key] = LocaleE(locale, value.(string))
	}
	return fmt.Sprintf(result, values...)
}