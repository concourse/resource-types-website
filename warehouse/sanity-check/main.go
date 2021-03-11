package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/resource"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

func printColor(colour *color.Color, strToPrint ...string) {
	if _, err := colour.Println(strToPrint); err != nil {
		fmt.Println(strToPrint)
	}
}

func main() {

	ftchr := fetcher.Fetcher{
		Box: os.DirFS("../resource-types"),
	}

	files, err := ftchr.GetAll()
	if err != nil {
		panic(err)
	}
	red := color.New(color.FgRed, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)

	errors := []string{}
	warning := []string{}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, fileBytes := range files {
		if strings.Contains(fileBytes.Name, ".yml") {
			failed := true
			var currResource resource.Resource

			err = yaml.UnmarshalStrict(fileBytes.Contents, &currResource)
			if err != nil {
				errors = append(errors, fileBytes.Name+red.Sprintf(" is not a valid resource type file"))
				printColor(red, "checked: ", fileBytes.Name)
				continue
			}

			res, err := client.Get(currResource.URL)
			if err != nil {
				errors = append(errors, "In: "+fileBytes.Name+","+red.Sprintf(currResource.URL+" Get error "+err.Error()))
				printColor(red, "checked: ", fileBytes.Name)
				continue
			}
			if res.StatusCode == http.StatusMovedPermanently {
				movedLocation, _ := res.Location()
				errors = append(errors, "In: "+fileBytes.Name+","+red.Sprintf(" Repo "+currResource.URL+" moved ----> "+movedLocation.String()))
				failed = false
			}

			if res.StatusCode == http.StatusNotFound {
				errors = append(errors, "In: "+fileBytes.Name+","+red.Sprintf(" Repo "+currResource.URL+" Not Found"))
				failed = false
			}

			res, err = http.Get("https://hub.docker.com/v2/repositories/" + currResource.Image)
			if err != nil {
				panic(err)
			}
			if res.StatusCode == http.StatusNotFound {
				if len(strings.Split(currResource.Image, "/")) > 2 {
					warning = append(warning, "In: "+fileBytes.Name+", "+yellow.Sprintf("Skipped image check for non docker hub image "+currResource.Image))
				} else {
					errors = append(errors, "In: "+fileBytes.Name+","+red.Sprintf(" docker image "+currResource.Image+" deosn't exist"))
					failed = false
				}
			}
			if failed {
				printColor(green, "checked: ", fileBytes.Name)
			} else {
				printColor(red, "checked: ", fileBytes.Name)
			}
		}
	}
	if len(warning) > 0 {
		printColor(yellow, "Warning:")
		for _, w := range warning {
			fmt.Println(w)
		}
	}
	if len(errors) > 0 {
		printColor(red, "Error:")
		for _, problem := range errors {
			fmt.Println(problem)
		}
		os.Exit(1)
	}

	printColor(green, "Everything seems as clean as it could be!!, push it to the cloud")
}
