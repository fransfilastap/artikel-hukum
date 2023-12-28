package main

import (
	"flag"
)

func main() {
	_ = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()

}
