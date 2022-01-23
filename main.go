package main

import (
	"fmt"
	"os"
)

func getArgs() (string, string) {
	// handle args to choose api type
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./rice_again <domain> <api_type>")
		os.Exit(1)
	}
	domain := os.Args[1]
	apiType := os.Args[2]
	return domain, apiType
}

func generatePayloadList(apiType string) string {
	// generate payload_list from word_list
	// open wordlist for this apiType
	// read data
	// generate wordlist for this apiType using custom data
	// return string of payloads separated by newlines
	payloadList := ""
	if apiType == "graphql" || apiType == "gql" {
		payloadList = apiType + "_graphqlPayloadList"
	} else if apiType == "restful" || apiType == "rest" {
		payloadList = apiType + "_restfulPayloadList"
	} else {
		payloadList = apiType + "_error_payloadList"
		fmt.Println("SET 'api_type' to graphql['graphql', 'gql'] or restful['restful', 'rest]")

	}
	return payloadList
}

func flameApi(domain string, payloadList string) string {
	// spam requests for different methods depending on api type
	// and check which ones return a OK status code to enumerate api
	// execute(GET/POST domain+payload) for each payload
	// return apiMap of calls that returned OK status code
	apiMap := payloadList + "_apiMap[" + domain + "]"
	return apiMap
}

func formatMap(apiMap string) string {
	// then format print to screen and log
	output := apiMap + "_output"
	return output
}

func main() {
	domain, apiType := getArgs()

	payloadList := generatePayloadList(apiType)

	apiMap := flameApi(domain, payloadList)

	output := formatMap(apiMap)
	fmt.Printf("%s\n", output)
}
