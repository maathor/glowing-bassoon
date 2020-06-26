package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	awsec2 "./awsec2"
	parseconf "./parseconf"
)

func main() {
	confs := parseconf.ParseConfigFile()
	lis := parseconf.ListEnvFromConfig(confs)
	_, envChosen := askFor("Please type %v and then press enter:\n", lis)
	err, env := parseconf.GetVarsForThisEnv(envChosen, confs)
	if nil != err {
		log.Error("Unable to find vars for specified environment")
	}
	_, tagChosen := askFor("Microservice Searched\n", nil)
	listPrivateIPs := awsec2.GetPrivateIpFromTag(tagChosen, env)
	for _, elem := range listPrivateIPs {
		fmt.Printf("IPs releaved %s\n", elem)
	}
}

func askFor(question string, okayResponses []string) (bool, string) {
	if okayResponses != nil {
		fmt.Printf(question, okayResponses)
	} else {
		fmt.Printf(question)
	}
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	if okayResponses != nil {
		if containsString(okayResponses, response) {
			return true, response
		} else {
			return askFor(question, okayResponses)
		}
	} else {
		return true, response
	}
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func searchingForIPs(tag string) {

}
