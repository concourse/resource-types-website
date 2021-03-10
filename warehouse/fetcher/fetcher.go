package fetcher

import (
	"io/fs"
	"net/http"
)

type Fetcher struct {
	Box fs.FS
}

type File struct {
	Name     string
	Contents []byte
}

func (f Fetcher) GetFile(name string) ([]byte, error) {
	return fs.ReadFile(f.Box, name)
}

func (f Fetcher) Open(name string) (http.File, error) {
	return http.FS(f.Box).Open(name)
}

func (f Fetcher) GetAll() ([]File, error) {
	var files []File
	err := fs.WalkDir(f.Box, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		fileBytes, err := fs.ReadFile(f.Box, path)
		if err != nil {
			return err
		}
		currFile := File{
			Name:     d.Name(),
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
