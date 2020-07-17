package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
)

type _Config struct {
	file *ini.File
}

var configInstance *_Config
var once sync.Once

func InitConfigInstance(path string) *_Config {
	once.Do(func() {
		configInstance = new(_Config)
		configInstance.LoadConfFromFile(path)
	})
	return configInstance
}

func Config() *_Config {
	if configInstance == nil{
		panic("configInstance is nil")
	}
	return configInstance
}

func (c _Config) LoadConfFromFile(path string) {
	cfg, err := ini.ShadowLoad(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		panic("Fail to read file")
	}
	c.file = cfg
	configInstance.file = cfg
}

func (c _Config) Get(section string, key string) *ini.Key {
	return c.file.Section(section).Key(key)
}

func (c _Config) HasValve(section string, key string) bool {
	return c.file.Section(section).HasKey(key)
}

func (c _Config) GetKeys(section string) []*ini.Key {
	return c.file.Section(section).Keys()
}

func (c _Config) GetSectionKeys(section string) []string {
	return c.file.Section(section).KeyStrings()
}

func (c _Config) GetSection(sectionName string) *ini.Section {
	section, _ := c.file.GetSection(sectionName)
	return section
}

func (c _Config) GetSections(sectionName string) []*ini.Section {
	return c.file.Sections()
}

func (c _Config) GetSectionNames(sectionName string) []string {
	return c.file.SectionStrings()
}

func (c _Config) GetSectionValues(section string) []string {
	var values []string
	keys := c.GetSectionKeys(section)
	for _, key := range keys {
		values = append(values, c.Get(section, key).Value())
	}
	return values
}
