package handler

func NewstaticHandler(path string) staticHandler {
	return staticHandler{Path: path}
}
