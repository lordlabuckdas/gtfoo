package main

import (
	"flag"
	"fmt"

	"github.com/lordlabuckdas/gtfoo/gtfobins"
	"github.com/lordlabuckdas/gtfoo/lolbas"
)

var gtfoSearch, lolbasSearch string

func init() {
	flag.StringVar(&gtfoSearch, "gtfo", "", "gtfobin search term")
	flag.StringVar(&gtfoSearch, "g", "", "gtfobin search term (shorthand flag)")
	flag.StringVar(&lolbasSearch, "lolbas", "", "lolbas search term")
	flag.StringVar(&lolbasSearch, "l", "", "lolbas search term (shorthand flag)")
	flag.Parse()
}

func main() {
	fmt.Println(gtfobins.Greet(gtfoSearch))
	fmt.Println(lolbas.Hello(lolbasSearch))
}
