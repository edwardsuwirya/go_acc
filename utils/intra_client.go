package utils

import (
	"io"
	"net/http"
)

type IntraClient struct {
	http.Client
}

func (ic *IntraClient) intraDo(req *http.Request) (*http.Response, error) {
	return ic.Do(req)
}

func (ic *IntraClient) IntraGet(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return ic.intraDo(request)
}

func (ic *IntraClient) IntraPost(url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return ic.intraDo(request)
}
