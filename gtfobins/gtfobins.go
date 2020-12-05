package gtfobins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

var exploit gtfoStruct

func getExploit(name string) gtfoStruct {
	url := baseURL + name + ".md"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(colorRed + "Error fetching data:" + colorReset, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println(colorRed + "Queried binary not available!" + colorReset)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(colorRed + "Error reading data:" + colorReset, err)
		os.Exit(1)
	}

	var gtfoExploit gtfoStruct
	err = yaml.Unmarshal(body, &gtfoExploit)
	if err != nil {
		fmt.Println(colorRed + "Error parsing YAML file:" + colorReset, err)
		os.Exit(1)
	}
	return gtfoExploit
}

// returns formatted exploit
func exploitFormatter(funcArray []function, funcName string) string {
	funcString := colorCyan + funcName + colorReset + "\n" + colorGreen + funcDesc[funcName] + colorReset + "\n"

	for i, s := range funcArray {

		if s.Description !=  {
			funcString += colorWhite + fmt.Sprintf("[%d] Description:\n%s", i+1, s.Description) + colorReset + "\n"
		}
		if s.Code !=  {
			funcString += colorYellow + "Code:" + s.Code + colorReset + "\n"
		}
	}
	return funcString
}

// Shell - returns formatted Shell exploit
func Shell() string {
	if exploit.Functions.Shell != nil {
		return exploitFormatter(exploit.Functions.Shell, shell)
	}
	return 
}

// Command - returns formatted Command exploit
func Command() string {
	if exploit.Functions.Command != nil {
		return exploitFormatter(exploit.Functions.Command, cmd)
	}
	return 
}

// ReverseShell - returns formatted ReverseShell exploit
func ReverseShell() string {
	if exploit.Functions.ReverseShell != nil {
		return exploitFormatter(exploit.Functions.ReverseShell, revShell)
	}
	return 
}

// NonInteractiveReverseShell - returns formatted NonInteractiveReverseShell exploit
func NonInteractiveReverseShell() string {
	if exploit.Functions.NonInteractiveReverseShell != nil {
		return exploitFormatter(exploit.Functions.NonInteractiveReverseShell, nonIntRevShell)
	}
	return 
}

// BindShell - returns formatted BindShell exploit
func BindShell() string {
	if exploit.Functions.BindShell != nil {
		return exploitFormatter(exploit.Functions.BindShell, bindShell)
	}
	return 
}

//NonInteractiveBindShell - returns formatted NonInteractiveBindShell exploit
func NonInteractiveBindShell() string {
	if exploit.Functions.NonInteractiveBindShell != nil {
		return exploitFormatter(exploit.Functions.NonInteractiveBindShell, nonIntBindShell)
	}
	return 
}

// FileUpload - returns formatted FileUpload exploit
func FileUpload() string {
	if exploit.Functions.FileUpload != nil {
		return exploitFormatter(exploit.Functions.FileUpload, fileUpload)
	}
	return 
}

// FileDownload - returns formatted FileDownload exploit
func FileDownload() string {
	if exploit.Functions.FileDownload != nil {
		return exploitFormatter(exploit.Functions.FileDownload, fileDownload)
	}
	return 
}

// FileWrite - returns formatted FileWrite exploit
func FileWrite() string {
	if exploit.Functions.FileWrite != nil {
		return exploitFormatter(exploit.Functions.FileWrite, fileWrite)
	}
	return 
}

// FileRead - returns formatted FileRead exploit
func FileRead() string {
	if exploit.Functions.FileRead != nil {
		return exploitFormatter(exploit.Functions.FileRead, fileRead)
	}
	return 
}

// LibraryLoad - returns formatted LibraryLoad exploit
func LibraryLoad() string {
	if exploit.Functions.LibraryLoad != nil {
		return exploitFormatter(exploit.Functions.LibraryLoad, libLoad)
	}
	return 
}

// SUID - returns formatted Suid  exploit
func SUID() string {
	if exploit.Functions.SUID != nil {
		return exploitFormatter(exploit.Functions.SUID, suid)
	}
	return 
}

// Sudo - returns formatted Sudo exploit
func Sudo() string {
	if exploit.Functions.Sudo != nil {
		return exploitFormatter(exploit.Functions.Sudo, sudo)
	}
	return 
}

// Capabilities - returns formatted Capabilities exploit
func Capabilities() string {
	if exploit.Functions.Capabilities != nil {
		return exploitFormatter(exploit.Functions.Capabilities, capab)
	}
	return 
}

// LimitedSuid - returns formatted LimitedSuid exploit
func LimitedSuid() string {
	if exploit.Functions.LimitedSuid != nil {
		return exploitFormatter(exploit.Functions.LimitedSuid, limSUID)
	}
	return 
}

// GtfobinMain - main function of gtfobins package
func GtfobinMain(name string) {
	exploit = getExploit(name)
}
