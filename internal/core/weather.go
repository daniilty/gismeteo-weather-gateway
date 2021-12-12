package core

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/daniilty/gismeteo-weather-gateway/internal/requests"
	"golang.org/x/net/html"
)

func (s *ServiceImpl) GetWeather() (string, error) {
	const (
		minLen                  = 1
		unitTempXpath           = `//div[@class="information__content__temperature"]`
		weatherDescriptionXpath = `//div[@class="information__content__additional information__content__additional_first"]/div[@class="information__content__additional__item"]`
	)

	body := requests.GetEmptyBody()

	req, err := http.NewRequest(http.MethodGet, s.baseURL, body)
	if err != nil {
		return "", fmt.Errorf("create HTTP request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do HTTP request: %w", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parse response body: %w", err)
	}

	unitTemp, err := htmlquery.QueryAll(doc, unitTempXpath)
	if err != nil {
		return "", fmt.Errorf("query xpath: %s: %w", unitTempXpath, err)
	}

	weatherDescription, err := htmlquery.QueryAll(doc, weatherDescriptionXpath)
	if err != nil {
		return "", fmt.Errorf("query xpath: %s: %w", weatherDescriptionXpath, err)
	}

	if len(unitTemp) < minLen {
		return "", fmt.Errorf("no temperature")
	}

	if len(weatherDescription) < minLen {
		return "", fmt.Errorf("no weather description")
	}

	description := strings.ReplaceAll(weatherDescription[0].LastChild.Data, "\n", "")
	description = strings.ReplaceAll(description, "\t", "")

	temp := strings.ReplaceAll(unitTemp[0].LastChild.Data, "\n", "")
	temp = strings.ReplaceAll(temp, "\t", "")

	weather := temp + " " + description

	return weather, nil
}
