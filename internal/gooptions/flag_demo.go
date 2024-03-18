package gooptions

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func FlagUse() {
	var name string
	var help bool

	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅2", "帮助信息")
	flag.Usage = func() {
		_, err := fmt.Fprint(os.Stderr, "Usage: main [-Option] \n Options:")
		if err != nil {
			log.Fatal(err)
		}

		flag.PrintDefaults()
	}

	flag.Parse()
	if help {
		flag.Usage()
	}
}
