package main

import (
	"flag"

	"github.com/lordlabuckdas/gtfoo/gtfobins"
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
	gtfobins.GtfobinMain(gtfoSearch)
}
