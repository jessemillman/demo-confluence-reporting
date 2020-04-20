package main

import confluence "github.com/jessemillman/confluence-go-api"

// Report is the aggregate used for the report we want
type Report struct {
	ID          string                 `json:"id,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Status      string                 `json:"status,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Version     confluence.Version     `json:"version"`
	Space       confluence.Space       `json:"space"`
	LastUpdated confluence.LastUpdated `json:"lastUpdated"`
	Latest      bool                   `json:"latest"`
	CreatedBy   confluence.User        `json:"createdBy"`
	CreatedDate string                 `json:"createdDate"`
}

type SimpleReport struct {
	ID            string `csv:"id"`
	Type          string `csv:"type"`
	Status        string `csv:"status"`
	Title         string `csv:"title"`
	Version       int    `csv:"version"`
	Space         string `csv:"space"`
	LastUpdated   string `csv:"lastUpdated"`
	LastUpdatedBy string `csv:"lastUpdatedBy"`
	Latest        bool   `csv:"latest"`
	CreatedBy     string `csv:"createdBy"`
	CreatedDate   string `csv:"createdDate"`
}
