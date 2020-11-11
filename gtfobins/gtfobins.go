package gtfobins

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lordlabuckdas/gtfoo/utils"
	"gopkg.in/yaml.v2"
)

func getExploit(name string) gtfoStruct {

	url := baseURL + name + ".md"

	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Printf("Error fetching data : %s\n", err1)
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		fmt.Printf("Error reading data : %s\n", err2)
	}

	var gtfoExploit gtfoStruct
	err3 := yaml.Unmarshal(body, &gtfoExploit)
	if err3 != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err3)
	}

	return gtfoExploit

}

func Greet(name string) string {
	return fmt.Sprintf("%sHi, %v%s", utils.ColorBlue, name, utils.ColorReset)
}

func GtfobinMain(name string) {

	exploit := getExploit(name)

	
	//right now only the sudo function is printed
	fmt.Printf("\n\n%s\n",funcDesc["Sudo"])

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
