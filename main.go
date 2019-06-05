package main

import (
	"fmt"
	"os"
	"path"

	"github.com/concourse/dutyfree/sitegenerator"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) != 3 {
		usage("undefined output directory")
	}
	outputDir := os.Args[1]
	resourcesPath := os.Args[2]

	indexHTML, err := os.Create(path.Join(outputDir, "index.html"))
	if err != nil {
		usage("output directory cannot be found")
	}

	defer indexHTML.Close()

	resources, err := resourceReader(resourcesPath)

	if err != nil {
		usage(err.Error())
	}

	indexPage := sitegenerator.NewIndexPage("sitegenerator", resources)
	err = indexPage.Generate(indexHTML)

	if err != nil {
		fmt.Println("Cannot write index.html")
		os.Exit(1)
	}
	os.Exit(0)
}

func usage(errorMsg string) {
	fmt.Fprintln(os.Stderr, errorMsg)
	fmt.Fprintf(os.Stderr, "usage: %s <output-directory> <resource-file>\n", os.Args[0])
	os.Exit(1)
}

func resourceReader(resourcesPath string) ([]sitegenerator.Resource, error) {
	resourceFile, err := os.Open(resourcesPath)

	if err != nil {
		usage("cannot read resources file")
	}

	decoder := yaml.NewDecoder(resourceFile)

	var resources []sitegenerator.Resource

	err = decoder.Decode(&resources)

	if err != nil {
		return nil, fmt.Errorf("cannot decode resources yaml: %s", err)
	}

	return resources, nil
}
