package main

import (
	"flag"
	"fmt"
	"github.com/Pf-G/kael-kit/src/config"
)

func main() {
	confPath := flag.String("config", "", "config file path")
	runPath := flag.String("path", "", "")
	flag.Parse()
	config.InitConfigInstance(*confPath, *runPath)
	fmt.Println(config.Config().GetSectionValues("server.ips"))
	fmt.Println(config.Config().Get("", "locale"))

	fmt.Print(config.Locale("hello, world", config.Locale("nice")))
}
