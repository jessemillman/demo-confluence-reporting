package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func initialize() (configuration, error) {
	// args
	spaceKey := flag.String("spaceKey", "", "The Key of the space to process")
	reportType := flag.String("reportType", "csv", "Type of report to generate (csv or json) - csv is default")
	allSpacesFlag := flag.Bool("allSpaces", false, "Should we process all spaces?")
	flag.Parse()

	// validate the input
	if *spaceKey == "" && *reportType == "" && *allSpacesFlag == false {
		err := errors.New("Error validating input, try adding a report type or space key")
		log.Print(err)
		flag.PrintDefaults()
	}

	// validate the report type
	if !reportValidator(*reportType) {
		err := errors.New(fmt.Sprint("Error validating report type, try csv or json. Input was: ", *reportType))
		log.Print(err)
		flag.PrintDefaults()
	}

	// environment variables
	subdomain, domainExists := os.LookupEnv("CONFLUENCE_SUBDOMAIN")
	userName, usernameExists := os.LookupEnv("CONFLUENCE_USERNAME")
	apiKey, keyExists := os.LookupEnv("CONFLUENCE_KEY")

	if domainExists && usernameExists && keyExists {
		return configuration{
			SpaceKey:       *spaceKey,
			QueryAllSpaces: *allSpacesFlag,
			ConfluenceURL:  fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", subdomain),
			UserName:       userName,
			APIKey:         apiKey,
			ReportType:     *reportType,
		}, nil
	}
	return configuration{}, errors.New("Error generating configuration")
}

func reportValidator(r string) bool {
	switch r {
	case "csv":
		return true
	case "json":
		return true
	default:
		return false
	}
}
