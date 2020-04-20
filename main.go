package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	confluence "github.com/jessemillman/confluence-go-api"
)

func main() {

	// args
	spaceKey := flag.String("spaceKey", "", "Space to process")
	allSpacesFlag := flag.Bool("allSpaces", false, "Should we process all spaces?")
	flag.Parse()

	// environment variables
	subdomain, domainExists := os.LookupEnv("CONFLUENCE_SUBDOMAIN")
	userName, usernameExists := os.LookupEnv("CONFLUENCE_USERNAME")
	apiKey, keyExists := os.LookupEnv("CONFLUENCE_KEY")

	if domainExists && usernameExists && keyExists {
		// variables
		apiURL := fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", subdomain)
		var spaces *confluence.AllSpaces // the spaces to process

		reports := []SimpleReport{} // what we'll report on
		expand := []string{         // options to expand (confluence api)
			"version", // gets version information
			"space",   // gets parent space information
		}

		// valiate input
		if *spaceKey == "" && *allSpacesFlag == false {
			flag.PrintDefaults()
			os.Exit(1)
		}

		// initialize a new api instance
		api, err := confluence.NewAPI(apiURL,
			userName,
			apiKey)
		if err != nil {
			log.Fatal(err)
		}

		// get all spaces if requested
		if *allSpacesFlag {
			spaces, err = api.GetAllSpaces(confluence.AllSpacesQuery{
				Status: "current",
				Type:   "global",
				Limit:  500,
			})
			if err != nil {
				log.Fatal(err)
			}
		} else { // check only the requested spaceKey
			spaces, err = api.GetAllSpaces(confluence.AllSpacesQuery{
				Status:   "current",
				Type:     "global",
				SpaceKey: *spaceKey,
				Limit:    500,
			})
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, s := range spaces.Results {
			// get contents of space
			fmt.Println("Processing Space ", s.Key)

			res, err := api.GetContent(confluence.ContentQuery{
				SpaceKey: s.Key,
				Type:     "page",
				OrderBy:  "history.createdDate desc",
				Expand:   expand,
				Limit:    1000,
			})
			if err != nil {
				log.Fatal(err)
			}

			for _, v := range res.Results {
				//fmt.Printf("%+v\n", v)
				fmt.Println("Processing Page ", v.Title)

				h, err := api.GetHistory(v.ID)
				if err != nil {
					log.Fatal(err)
				}
				r := SimpleReport{
					ID:          v.ID,
					Type:        v.Type,
					Status:      v.Status,
					Title:       v.Title,
					Version:     v.Version.Number,
					Space:       v.Space.Key,
					LastUpdated: h.LastUpdated.When,
					Latest:      h.Latest,
					CreatedBy:   h.CreatedBy.DisplayName,
					CreatedDate: h.CreatedDate,
				}
				reports = append(reports, r)
			}
		}
		b, err := gocsv.MarshalString(&reports)

		file, err := os.Create("/output/results.csv")
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(b)
			fmt.Println("Done")
		}
		file.Close()

	} else {
		os.Exit(1)
	}
}
