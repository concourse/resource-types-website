package indexhandler

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"sync"

	"github.com/concourse/dutyfree/fetcher"
)

type templateFs struct {
	assetIDs map[string]string
	assetsL  sync.Mutex
	fetcher  fetcher.Fetcher
}

func TemplateFunctions(fetchr fetcher.Fetcher) template.FuncMap {
	tfs := &templateFs{
		assetIDs: map[string]string{},
		fetcher:  fetchr,
	}
	return template.FuncMap{
		"asset": tfs.asset,
	}
}

func (fs *templateFs) asset(asset string) (string, error) {
	fs.assetsL.Lock()
	defer fs.assetsL.Unlock()

	id, found := fs.assetIDs[asset]
	if !found {
		hash := md5.New()

		contents, err := fs.fetcher.GetFile(asset)
		if err != nil {
			return "", err
		}

		_, err = hash.Write(contents)
		if err != nil {
			return "", err
		}

		id = fmt.Sprintf("%x", hash.Sum(nil))
	}

	return fmt.Sprintf("/public/%s?id=%s", asset, id), nil
}
