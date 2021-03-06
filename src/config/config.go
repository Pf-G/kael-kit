package config

import (
	"fmt"
	"github.com/Pf-G/kael-kit/src/share"
	"gopkg.in/ini.v1"
	"sync"
)

type _Config struct {
	file    *ini.File
	runPath string
}

var (
	configInstance *_Config
	once           sync.Once
)

func InitConfigInstance(path string, runPath string) *_Config {
	once.Do(func() {
		if path == "" {
			path = share.GetDefaultConfigPath()
		}
		if runPath == "" {
			runPath = share.GetRunPath()
		}
		configInstance = new(_Config)
		cfg, err := ini.ShadowLoad(path)
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			panic("Fail to read file")
		}
		configInstance.file = cfg
		configInstance.runPath = runPath
	})
	return configInstance
}

func Config() *_Config {
	if configInstance == nil {
		InitConfigInstance(share.GetDefaultConfigPath(), share.GetRunPath())
	}
	return configInstance
}

func (c _Config) LoadConfFromFile(path string) *_Config {
	cfg, err := ini.ShadowLoad(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		panic("Fail to read file")
	}
	c.file = cfg
	return &c
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
