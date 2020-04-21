module github.com/jessemillman/demo-confluence-reporting

go 1.12

require (
	github.com/jessemillman/confluence-go-api v1.0.5
	github.com/jessemillman/demo-confluence-reporting/common v0.0.0-00010101000000-000000000000
	github.com/jessemillman/demo-confluence-reporting/config v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)

replace github.com/jessemillman/demo-confluence-reporting/config => ./config

replace github.com/jessemillman/demo-confluence-reporting/common => ./common
