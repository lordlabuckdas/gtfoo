package main

import (
	"flag"
	"fmt"

	"github.com/lordlabuckdas/gtfoo/gtfobins"
	"github.com/lordlabuckdas/gtfoo/lolbas"
)

var gtfoSearch, lolbasSearch *string

func init() {
	gtfoSearch = flag.String("g", "", "enter binary to be searched in gtfobins")
	lolbasSearch = flag.String("l", "", "enter binary to be searched in lolbas")
	flag.Parse()
}

func main() {
	fmt.Println(gtfobins.Greet(*gtfoSearch))
	fmt.Println(lolbas.Hello(*lolbasSearch))
}
