package main

import (
	"fmt"
	"html/template"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		usage("undefined output directory")
	}
	outputDir := os.Args[1]

	f, err := os.Create(path.Join(outputDir, "index.html"))
	if err != nil {
		usage("output directory cannot be found")
	}

	defer f.Close()

	tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))

	err = tmpl.Execute(f, nil)

	if err != nil {
		fmt.Println("Cannot write index.html")
		os.Exit(1)
	}
	os.Exit(0)
}

func usage(errorMsg string) {
	fmt.Fprintln(os.Stderr, errorMsg)
	fmt.Fprintf(os.Stderr, "usage: %s <output-directory>\n", os.Args[0])
	os.Exit(1)
}
