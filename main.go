package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/spenserblack/gh-hacktoberfest/pkg/label"
)

var hackoberfestLabels = [2]label.Label{
	{Name: "hacktoberfest", Description: "Good for hacktoberfest participants", Color: "FF8800"},
	{Name: "hacktoberfest-accepted", Description: "Accepted hacktoberfest contributions", Color: "FFBB00"},
}

func main() {
	flag.Parse()
	repo, err := repository.Parse(repovar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
		os.Exit(1)
	}
	client, err := gh.RESTClient(nil)
	fmt.Println("Creating labels...")

	// NOTE: Refer to https://docs.github.com/en/rest/issues/labels#create-a-label
	for _, l := range hackoberfestLabels {
		body, _ := json.Marshal(l)
		response := label.Label{}
		err := client.Post(
			fmt.Sprintf("repos/%s/%s/labels", repo.Owner(), repo.Name()),
			bytes.NewReader(body),
			&response,
		)
		// TODO: Type assert to go-gh/pkg/api.HTTPError to check if it's an already exists error
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}
		fmt.Printf("Created label %s\n", response.Name)
	}
}

var repovar string

func init() {
	flag.StringVar(&repovar, "R", defaultRepo(), "repo to query")
}

func defaultRepo() string {
	repo, err := gh.CurrentRepository()
	if err != nil || repo == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", repo.Owner(), repo.Name())
}
