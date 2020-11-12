package gtfobins

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func init() {
	logFileName := "logs.txt"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Could not create log file: %s", err)
	}

	infoLogger = log.New(logFile, colorYellow+"[INFO] "+colorReset, log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(logFile, colorRed+"[ERROR] "+colorReset, log.Ldate|log.Ltime|log.Lshortfile)
}

func getExploit(name string) gtfoStruct {
	url := fmt.Sprintf(baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		errorLogger.Fatalf("Error fetching data : %s\n", err)
	} else {
		infoLogger.Printf("Fetching data from %s", url)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorLogger.Fatalf("Error reading data : %s\n", err)
	} else {
		infoLogger.Printf("Reading data from GTFObins for %s", name)
	}

	var gtfoExploit gtfoStruct
	err = yaml.Unmarshal(body, &gtfoExploit)
	if err != nil {
		errorLogger.Fatalf("Error parsing YAML file: %s\n", err)
	} else {
		infoLogger.Printf("Parsing %s data", name)
	}

	return gtfoExploit
}

func GtfobinMain(name string) {
	exploit := getExploit(name)
	//right now only the sudo function is printed
	fmt.Printf("\n\n%s\n", funcDesc["Sudo"])

	for i := 0; i < len(exploit.Functions.Sudo); i++ {
		fmt.Printf("\n")
		if exploit.Functions.Sudo[i].Description != "" {
			fmt.Printf("Description : %s\n", exploit.Functions.Sudo[i].Description)
		}
		if exploit.Functions.Sudo[i].Code != "" {
			fmt.Printf("Code : %s\n", exploit.Functions.Sudo[i].Code)
		}
		fmt.Printf("\n")
	}
}
