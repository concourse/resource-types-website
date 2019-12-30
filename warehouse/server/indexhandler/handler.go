package indexhandler

import (
	"html/template"
	"net/http"

	"github.com/gobuffalo/packr"
)

type Handler struct {
	template *template.Template
}

func NewHandler(path string) (Handler, error) {

	fns := TemplateFunctions(path)

	box := packr.NewBox(path)

	src, err := box.Find("index.html")
	if err != nil {
		return Handler{}, err
	}

	t, err := template.New("index").Funcs(fns).Parse(string(src))
	if err != nil {
		return Handler{}, err
	}

	return Handler{
		template: t,
	}, nil
}
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.template.Execute(w, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
