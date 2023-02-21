package core

import (
	"bytes"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

type FileList shell.LsLink

type API struct {
	sh *shell.Shell
}

func NewAPI() *API {
	sh := shell.NewShell("localhost:5001")
	return &API{
		sh: sh,
	}
}

func (a *API) Upload(data []byte) (string, error) {
	hash, err := a.sh.Add(bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (a *API) Download(hash string) ([]byte, error) {
	rc, err := a.sh.Cat(hash)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	b, err2 := io.ReadAll(rc)
	if err2 != nil {
		return nil, err2
	}
	return b, nil

}

func (a *API) GetReadCloser(hash string) (io.ReadCloser, error) {
	rc, err := a.sh.Cat(hash)
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func (a *API) List(hash string) ([]*FileList, error) {
	ll, err := a.sh.List(hash)
	if err != nil {
		return nil, err
	}
	res := make([]*FileList, 0, 10)
	for _, v := range ll {
		l := FileList(*v)
		res = append(res, &l)
	}
	return res, nil
}
