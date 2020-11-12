package gtfobins

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

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
