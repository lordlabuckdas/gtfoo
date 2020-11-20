package main

import (
	"flag"
	"fmt"

	"github.com/lordlabuckdas/gtfoo/gtfobins"
)

// var gtfoSearch, lolbasSearch string
var gtfoSearch string

func init() {
	flag.StringVar(&gtfoSearch, "gtfo", "", "gtfobin search term")
	flag.StringVar(&gtfoSearch, "g", "", "gtfobin search term (shorthand flag)")
	// flag.StringVar(&lolbasSearch, "lolbas", "", "lolbas search term")
	// flag.StringVar(&lolbasSearch, "l", "", "lolbas search term (shorthand flag)")
	flag.Parse()
}

func main() {
	if len(gtfoSearch) == 0 {
		fmt.Print("No search query entered!")
	} else {
		gtfobins.GtfobinMain(gtfoSearch)
		fmt.Print("\n")
		fmt.Print(gtfobins.Sudo())
		fmt.Print(gtfobins.Shell())
		fmt.Print(gtfobins.Command())
		fmt.Print(gtfobins.ReverseShell())
		fmt.Print(gtfobins.NonInteractiveReverseShell())
		fmt.Print(gtfobins.BindShell())
		fmt.Print(gtfobins.NonInteractiveBindShell())
		fmt.Print(gtfobins.FileUpload())
		fmt.Print(gtfobins.FileDownload())
		fmt.Print(gtfobins.FileWrite())
		fmt.Print(gtfobins.FileRead())
		fmt.Print(gtfobins.LibraryLoad())
		fmt.Print(gtfobins.SUID())
		fmt.Print(gtfobins.Capabilities())
		fmt.Print(gtfobins.LimitedSuid())
	}
}
