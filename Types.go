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
