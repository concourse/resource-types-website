package main

import (
	"fmt"
	"os"
	"path"

	"github.com/concourse/dutyfree/sitegenerator"
	"github.com/otiai10/copy"
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
		fmt.Println("Cannot write index.html", err)
		os.Exit(1)
	}

	err = os.MkdirAll(path.Join(outputDir, "resources"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		usage("resources folder cannot be created")
	}

	for _, resource := range resources {
		fileName := resource.Identifier

		resourceHTML, err := os.Create(path.Join(outputDir, "resources", fmt.Sprintf("%s.html", fileName)))
		if err != nil {
			fmt.Println(err)
			usage("resource page cannot be generated")
			continue
		}

		rp := sitegenerator.NewResourcePage("sitegenerator", resource)
		err = rp.Generate(resourceHTML)

		if err != nil {
			usage(fmt.Sprintf("resource page %s cannot be generated", resourceHTML.Name()))
		}

		resourceHTML.Close()
	}

	copy.Copy("static", path.Join(outputDir, "static"))

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
