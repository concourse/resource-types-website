package sitegenerator

import (
	"html/template"
	"path"
	"runtime"
)

func load(templateName string) *template.Template {
	_, filename, _, _ := runtime.Caller(0)

	funcMap := template.FuncMap{"lastindex": lastindex}

	return template.Must(template.New(templateName).Funcs(funcMap).ParseFiles(
		path.Join(path.Dir(filename), "templates/page_layout.html"),
		path.Join(path.Dir(filename), "templates", templateName),
	))
}

func lastindex(slice []string) int {
	return len(slice) - 1
}
