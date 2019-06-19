package sitegenerator

import (
	"html/template"
	"path"
	"runtime"
)

func load(templateName string) *template.Template {
	_, filename, _, _ := runtime.Caller(0)

	return template.Must(template.ParseFiles(
		path.Join(path.Dir(filename), "templates/page_layout.html"),
		path.Join(path.Dir(filename), "templates", templateName),
	))
}
