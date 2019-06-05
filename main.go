package main

import (
	"fmt"
	"html/template"
	"os"
	"path"

	yaml "gopkg.in/yaml.v2"
)

type Resource struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

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

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	resourceFile, err := os.Open(resourcesPath)

	if err != nil {
		usage("cannot read resources file")
	}

	decoder := yaml.NewDecoder(resourceFile)

	var resources []Resource

	decoder.Decode(&resources)

	err = tmpl.Execute(indexHTML, resources)

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
