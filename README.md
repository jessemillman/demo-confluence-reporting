# # confluence-reporting

Is a demonstration golang application to pull reporting information out of confluence.

## Supported Features

- Gets History & Version information from all of pages in your spaces as JSON or CSV.
- Can enumerate _all_ spaces, or an _individual_ defined space.

This is a demonstration application & therefore extremely limited in scope. If you want additional features, please feel free to raise an issue, or have a crack yourself & open a PR.

## Installation
### Prerequisites
1. Please ensure you've got an API key for confluence. You can generate one [here](https://id.atlassian.com/manage-profile/security/api-tokens).

### Docker
It's easiest to use the docker image, so after cloning this repo run

```
docker build --tag demo-confluence-reporting:latest .
```
This will build the image, which you can then use with (dont forget to also have the correct environment variables set):

```
Docker run --name cr demo-confluence-reporting ./main -spaceKey=example -reportType=csv
```

### Source
1. Ensure you have at least go1.12+ installed by running `go version` & updating if necessary
2. Grab a copy of the confluence api wrapper by running `go get github.com/jessemillman/confluence-go-api`
3. Ensure you have the required environment varibales set (i.e. run `export CONFLUENCE_SUBDOMAIN=XXXX`, etc)
4. Build the application with `go build all`
5. Run the application with `./demo-confluence-reporting --spaceKey=TECH --reportType=csv` or similar.

## Usage

### Simple example

To return information of all content, in all spaces as a CSV (this writes to):
`demo-confluence-reporting --allSpaces=true --reportType=csv`

To return information on a certain space with the key TEST in JSON:
`demo-confluence-reporting -spaceKey=TEST --reportType=json`

> NOTE: Output is written to a file `/output/result.csv` or `/output/result.json` because i map that to /tmp/ in docker with `-v /tmp:/output/`. 

## Code Documentation

Code is somewhat documented in line, but this is not meant to be production.

## Contribution

Please feel free to contribute.