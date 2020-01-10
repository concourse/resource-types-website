package indexhandler

import (
	"html/template"
	"net/http"

	"github.com/concourse/dutyfree/fetcher"
)

type Handler struct {
	template *template.Template
}

func NewHandler(fetchr fetcher.Fetcher) (Handler, error) {

	fns := TemplateFunctions(fetchr)

	src, err := fetchr.GetFile("index.html")
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
