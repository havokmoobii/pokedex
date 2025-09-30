package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//AreaInformation -
func (c *Client) AreaInformation(areaName string) (RespShallowArea, error) {
	url := baseURL + "/location-area/" + areaName
	
	cachedResp, exists := c.pokeapiCache.Get(url)

	if exists {
		areaResp := RespShallowArea{}
		err := json.Unmarshal(cachedResp, &areaResp)
		if err != nil {
			return RespShallowArea{}, err
		}
		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowArea{}, err
	}

	c.pokeapiCache.Add(url, dat)

	areaResp := RespShallowArea{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespShallowArea{}, err
	}
	
	return areaResp, nil
}