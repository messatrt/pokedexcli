package pokiapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageUrl *string) (LocationArea, error) {
	endpoint := "/location"
	fullurls := baseUrl + endpoint
	if pageUrl != nil {
		fullurls = *pageUrl
	}
	locations := LocationArea{}
	req, err := http.NewRequest("GET", fullurls, nil)
	if err != nil {
		return locations, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locations, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return locations, fmt.Errorf("bad status code : %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locations, err
	}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}
