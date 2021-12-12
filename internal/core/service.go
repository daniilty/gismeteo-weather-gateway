package core

import (
	"net/http"

	"golang.org/x/net/http2"
)

type Service interface {
	GetWeather() (string, error)
}

type ServiceImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewServiceImpl(baseURL string) *ServiceImpl {
	return &ServiceImpl{
		baseURL: baseURL,
		httpClient: &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
			},
		},
	}
}
