package main

import (
<<<<<<< HEAD
=======
	"flag"
>>>>>>> master
	"fmt"
	"log"

	"github.com/gocarina/gocsv"
	confluence "github.com/jessemillman/confluence-go-api"
	"github.com/jessemillman/demo-confluence-reporting/common"
	"github.com/jessemillman/demo-confluence-reporting/config"

	"github.com/pkg/errors"
)

func main() {

	// initialize configuration
	conf, confError := config.Initialize()
	if confError != nil {
		log.Fatal(errors.Wrap(confError, "Exiting due to configuration error"))
	}

	// initialize variables
	var spaces *confluence.AllSpaces // the spaces to process
	reports := []common.ReportLine{} // what we'll report on
	expand := []string{              // options to expand (confluence api)
		"version", // gets version information
		"space",   // gets parent space information
	}

	// initialize a new confluence api instance
	api, err := confluence.NewAPI(conf.ConfluenceURL,
		conf.UserName,
		conf.APIKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error with Confluence API"))
	}

<<<<<<< HEAD
	// get all spaces if requested
	if conf.QueryAllSpaces {
		spaces, err = api.GetAllSpaces(confluence.AllSpacesQuery{
			Status: "current",
			Type:   "global",
			Limit:  500,
		})
		if err != nil {
			log.Fatal(errors.Wrap(err, "Error with Querying All Spaces"))
=======
		reports := []SimpleReport{} // what we'll report on
		expand := []string{         // options to expand (confluence api)
			"version", // gets version information
			"space",   // gets parent space information
>>>>>>> master
		}
	} else { // check only the requested spaceKey
		spaces, err = api.GetAllSpaces(confluence.AllSpacesQuery{
			Status:   "current",
			Type:     "global",
			SpaceKey: conf.SpaceKey,
			Limit:    500,
		})
		if err != nil {
			log.Fatal(errors.Wrap(err, "Error with Querying Space"))
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
			fmt.Println("Processing Page ", v.Title)

			h, err := api.GetHistory(v.ID)
			if err != nil {
				log.Fatal(err)
			}
<<<<<<< HEAD
			r := common.ReportLine{
				ID:            v.ID,
				Type:          v.Type,
				Status:        v.Status,
				Title:         v.Title,
				Version:       v.Version.Number,
				Space:         v.Space.Key,
				LastUpdated:   h.LastUpdated.When,
				LastUpdatedBy: h.LastUpdated.By.DisplayName,
				Latest:        h.Latest,
				CreatedBy:     h.CreatedBy.DisplayName,
				CreatedDate:   h.CreatedDate,
=======

			for _, v := range res.Results {
				//fmt.Printf("%+v\n", v)
				fmt.Println("Processing Page ", v.Title)

				h, err := api.GetHistory(v.ID)
				if err != nil {
					log.Fatal(err)
				}
				r := SimpleReport{
					ID:            v.ID,
					Type:          v.Type,
					Status:        v.Status,
					Title:         v.Title,
					Version:       v.Version.Number,
					Space:         v.Space.Key,
					LastUpdated:   h.LastUpdated.When,
					LastUpdatedBy: h.LastUpdated.By.DisplayName,
					Latest:        h.Latest,
					CreatedBy:     h.CreatedBy.DisplayName,
					CreatedDate:   h.CreatedDate,
				}
				reports = append(reports, r)
>>>>>>> master
			}
			reports = append(reports, r)
		}
<<<<<<< HEAD
=======
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
>>>>>>> master
	}
	common.FileWriter(reports, conf.ReportType)

}
