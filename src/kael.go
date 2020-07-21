package main

import (
	"flag"
	"fmt"
	"github.com/Pf-G/kael-kit/src/config"
	"github.com/Pf-G/kael-kit/src/share"
	"github.com/briandowns/spinner"
	"time"
)

var (
	Locale  = config.Locale
	LocaleE = config.LocaleE
)

func main() {
	confPath := flag.String("config", "", "config file path")
	runPath := flag.String("path", share.GetRunPath(), "")
	flag.Parse()
	config.InitConfigInstance(*confPath, *runPath)

	fmt.Println(LocaleE("en", "hello world, what a nice day", "world", "nice"))
	fmt.Println(Locale("hello world, what a nice day", "world", "nice"))
	fmt.Println(Locale("nice"))

	s := spinner.New(spinner.CharSets[38], 1 * time.Second) // Build our new spinner
	s.Start()                                                    // Restart the spinner
	time.Sleep(10 * time.Second)
	s.Stop()
}
