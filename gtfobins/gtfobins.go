package gtfobins

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

var exploit gtfoStruct

func getExploit(name string) gtfoStruct {
	url := fmt.Sprintf(baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data : %s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data : %s\n", err)
	}

	var gtfoExploit gtfoStruct
	err = yaml.Unmarshal(body, &gtfoExploit)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}
	return gtfoExploit
}

//returns formatted exploit
func exploitFormatter(funcArray []function, funcName string) string {
	funcString := COLOR_CYAN + funcName + COLOR_RESET + "\n\n" + funcDesc[funcName] + "\n\n"

	for _, s := range funcArray {

		if s.Description != "" {
			funcString += "Description :\n" + s.Description + "\n"
		}
		if s.Code != "" {
			funcString += "Code :\n" + s.Code + "\n"
		}
		funcString += "\n"
	}
	return funcString
}

//Shell - returns formatted Shell exploit
func Shell() string {
	if exploit.Functions.Shell != nil {
		return exploitFormatter(exploit.Functions.Shell, SHELL)
	}
	return ""
}

//Command - returns formatted Command exploit
func Command() string {
	if exploit.Functions.Command != nil {
		return exploitFormatter(exploit.Functions.Command, COMMAND)
	}
	return ""
}

//ReverseShell - returns formatted ReverseShell exploit
func ReverseShell() string {
	if exploit.Functions.ReverseShell != nil {
		return exploitFormatter(exploit.Functions.ReverseShell, R_SHELL)
	}
	return ""
}

//NonInteractiveReverseShell - returns formatted NonInteractiveReverseShell exploit
func NonInteractiveReverseShell() string {
	if exploit.Functions.NonInteractiveReverseShell != nil {
		return exploitFormatter(exploit.Functions.NonInteractiveReverseShell, NI_R_SHELL)
	}
	return ""
}

//BindShell - returns formatted BindShell exploit
func BindShell() string {
	if exploit.Functions.BindShell != nil {
		return exploitFormatter(exploit.Functions.BindShell, B_SHELL)
	}
	return ""
}

//NonInteractiveBindShell - returns formatted NonInteractiveBindShell exploit
func NonInteractiveBindShell() string {
	if exploit.Functions.NonInteractiveBindShell != nil {
		return exploitFormatter(exploit.Functions.NonInteractiveBindShell, NI_B_SHELL)
	}
	return ""
}

//FileUpload - returns formatted FileUpload exploit
func FileUpload() string {
	if exploit.Functions.FileUpload != nil {
		return exploitFormatter(exploit.Functions.FileUpload, F_UPLOAD)
	}
	return ""
}

//FileDownload - returns formatted FileDownload exploit
func FileDownload() string {
	if exploit.Functions.FileDownload != nil {
		return exploitFormatter(exploit.Functions.FileDownload, F_DOWNLOAD)
	}
	return ""
}

//FileWrite - returns formatted FileWrite exploit
func FileWrite() string {
	if exploit.Functions.FileWrite != nil {
		return exploitFormatter(exploit.Functions.FileWrite, F_WRITE)
	}
	return ""
}

//FileRead - returns formatted FileRead exploit
func FileRead() string {
	if exploit.Functions.FileRead != nil {
		return exploitFormatter(exploit.Functions.FileRead, F_READ)
	}
	return ""
}

//LibraryLoad - returns formatted LibraryLoad exploit
func LibraryLoad() string {
	if exploit.Functions.LibraryLoad != nil {
		return exploitFormatter(exploit.Functions.LibraryLoad, LIBRARY_LOAD)
	}
	return ""
}

//Suid  - returns formatted Suid  exploit
func Suid() string {
	if exploit.Functions.Suid != nil {
		return exploitFormatter(exploit.Functions.Suid, SUID)
	}
	return ""
}

//Sudo - returns formatted Sudo exploit
func Sudo() string {
	if exploit.Functions.Sudo != nil {
		return exploitFormatter(exploit.Functions.Sudo, SUDO)
	}
	return ""
}

//Capabilities - returns formatted Capabilities exploit
func Capabilities() string {
	if exploit.Functions.Capabilities != nil {
		return exploitFormatter(exploit.Functions.Capabilities, CAPABILITIES)
	}
	return ""
}

//LimitedSuid - returns formatted LimitedSuid exploit
func LimitedSuid() string {
	if exploit.Functions.LimitedSuid != nil {
		return exploitFormatter(exploit.Functions.LimitedSuid, LIM_SUID)
	}
	return ""
}

//GtfobinMain - main function of gtfobins package
func GtfobinMain(name string) {
	exploit = getExploit(name)
}
