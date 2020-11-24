package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lordlabuckdas/gtfoo/gtfobins"
)

// var gtfoSearch, lolbasSearch string
var (
	gtfoSearch  string
	versionFlag bool
)

const banner = 
`░█▀▀░▀█▀░█▀▀░█▀█░█▀█
░█░█░░█░░█▀▀░█░█░█░█
░▀▀▀░░▀░░▀░░░▀▀▀░▀▀▀

by 7phalange7 and lordlabuckdas`

func init() {
	flag.StringVar(&gtfoSearch, "gtfo", "", "gtfobin query")
	flag.StringVar(&gtfoSearch, "g", "", "gtfobin query (shorthand flag)")
	flag.BoolVar(&versionFlag, "version", false, "to print the version of gtfoo")
	flag.BoolVar(&versionFlag, "v", false, "to print the version of gtfoo (shorthand flag)")
	// flag.StringVar(&lolbasSearch, "lolbas", "", "lolbas search term")
	// flag.StringVar(&lolbasSearch, "l", "", "lolbas search term (shorthand flag)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n%s\n",banner)
		fmt.Fprintf(os.Stderr, "\nUsage: %s [options] [query]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  -h or --help\n\tto print usage\n")
		flag.VisitAll(func(f *flag.Flag) {
			if len(f.Name) == 1 {
				fmt.Fprintf(os.Stderr, "  -%v or ", f.Name)
			} else {
				fmt.Fprintf(os.Stderr, "--%v\n\t%v\n", f.Name, f.Usage)
			}
		})
		fmt.Fprintf(os.Stderr, "Example:\n  %s --version\n  %s -g vim\n", os.Args[0], os.Args[0])
	}
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Println("gtfoo v1.00")
		os.Exit(0)
	}
	if len(gtfoSearch) == 0 {
		fmt.Printf("No search query entered!\nTry `%s -h` for help\n", os.Args[0])
		os.Exit(1)
	}
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
