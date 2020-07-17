package main

import (
	"flag"
	"fmt"
	"github.com/Pf-G/kael-kit/src/config"
)

func main() {
	confPath := flag.String("config", "", "")
	flag.Parse()
	config.InitConfigInstance(*confPath)
	fmt.Println(config.Config().GetSectionValues("server.ips"))
	fmt.Println(config.Config().Get("", "locale"))
}
