package fetcher

import (
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
)

type Fetcher struct {
	Box packr.Box
}

type File struct {
	Name     string
	Contents []byte
}

func (f Fetcher) GetFile(name string) ([]byte, error) {
	return f.Box.Find(name)
}

func (f Fetcher) Open(name string) (http.File, error) {
	return f.Box.Open(name)
}

func (f Fetcher) GetAll() ([]File, error) {
	var files []File
	err := f.Box.Walk(func(s string, file packd.File) error {
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		currFile := File{
			Name:     file.Name(),
			Contents: fileBytes,
		}
		files = append(files, currFile)
		return nil
	})
	if err != nil {
		return []File{}, err
	}
	return files, nil
}
