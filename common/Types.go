package common

// Configuration defines the settings used to run the program
type Configuration struct {
	SpaceKey       string `json:"spaceKey"`
	QueryAllSpaces bool   `json:"queryAllSpaces"`
	ConfluenceURL  string `json:"subdomain"`
	UserName       string `json:"userName"`
	APIKey         string `json:"apiKey"`
	ReportType     string `json:"type"`
}

// ReportLine is used to generate the csv & json output
type ReportLine struct {
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
