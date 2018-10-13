package riminder

import (
	"encoding/json"
	"reflect"

	"github.com/Xalrandion/go-riminder-api/riminder/errors"
	"gopkg.in/resty.v1"
)

type clientw struct {
	client      *resty.Client
	baseHeaders map[string]string
	baseURL     string
	apiKey      string
}

func newClientw(apiKey, baseURL string) *clientw {
	c := new(clientw)
	c.apiKey = apiKey
	c.baseURL = baseURL
	c.baseHeaders = map[string]string{
		"X-API-KEY": apiKey,
	}
	c.client = resty.New()
	c.client.SetHeaders(c.baseHeaders)
	c.client.SetHostURL(baseURL)
	return c
}

func (c *clientw) SetBaseURL(url string) {
	c.baseURL = url
	c.client.SetHostURL(c.baseURL)
}

func responseProcesser(resp *resty.Response, respContainer interface{}) (interface{}, error) {
	if !resp.IsSuccess() {
		return nil, errors.NewResponseError(resp.StatusCode(), resp.Status(), resp.String())
	}
	respFinal := reflect.New(reflect.TypeOf(respContainer)).Interface()
	err := json.Unmarshal(resp.Body(), &respFinal)
	if err != nil {
		return nil, err
	}

	return respFinal, nil
}

func (c *clientw) Get(endpoint string, queryParams map[string]string, respContainer interface{}) (interface{}, error) {

	resp, err := c.client.R().SetQueryParams(queryParams).Get(endpoint)
	if err != nil {
		return nil, err
	}
	return responseProcesser(resp, respContainer)
}
