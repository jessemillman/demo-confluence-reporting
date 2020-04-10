# # confluence-reporting

Is a demonstration golang application to pull reporting information out of confluence.

## Supportet Features

- Gets History & Version information from all of your spaces
- Cane enumerate _all_ spaces, or an _individual_ defined space.

This is a demonstration application & therefore extremely limited in scope. If you want additional features, please feel free to raise an issue, or have a crack yourself & open a PR.

## Installation

### Docker
It's easiest to use the docker image, so after cloning this repo run

```
docker build --tag confluence-reporting:latest .
```
This will build the image, which you can then use with (dont forget to also have the correct environment variables set):

```
Docker run --name cr confluence-reporting ./main -spaceKey="example"
```

### Source
1. Ensure you have at least go1.12+ installed by running `go version` & updating if necessary
2. Grab a copy of the confluence api wrapper by running `go get github.com/jessemillman/confluence-go-api`
3. Ensure you have the required environment varibales set

## Usage

### Simple example

To return information of all content, in all spaces:
`confluence-reporting -allSpaces=true`

To return information on a certain space:
`confluence-reporting -spaceKey="EXAMPLE"`

## Code Documentation

Code is somewhat documented in line, but this is not meant to be production.

## Contribution

Please feel free to contribute.