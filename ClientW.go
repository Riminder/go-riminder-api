package riminder

import (
	"encoding/json"

	"github.com/Xalrandion/go-riminder-api/riminder/errors"
	"gopkg.in/resty.v1"
)

// clientw is a wraper for resty library.
type clientw struct {
	client      *resty.Client
	baseHeaders map[string]string
	baseURL     string
	apiKey      string
}

// newClientw create a clientw
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

// responseProcesser transform a response to a response container.
func responseProcesser(resp *resty.Response, respContainer interface{}) error {
	if !resp.IsSuccess() {
		return errors.NewResponseError(resp.StatusCode(), resp.Status(), resp.String())
	}
	err := json.Unmarshal(resp.Body(), respContainer)
	if err != nil {
		return err
	}

	return nil
}

func (c *clientw) Get(endpoint string, queryParams map[string]string, respContainer interface{}) error {

	resp, err := c.client.R().SetQueryParams(queryParams).Get(endpoint)
	if err != nil {
		return err
	}
	return responseProcesser(resp, respContainer)
}

func (c *clientw) Post(endpoint string, bodyParams map[string]interface{}, respContainer interface{}) error {
	resp, err := c.client.R().SetBody(bodyParams).Post(endpoint)
	if err != nil {
		return err
	}
	return responseProcesser(resp, respContainer)
}

func (c *clientw) PostFile(endpoint, filepath string, multipartFields map[string]string, respContainer interface{}) error {
	fileMap := map[string]string{"file": filepath}
	resp, err := c.client.R().SetFiles(fileMap).SetFormData(multipartFields).Post(endpoint)
	if err != nil {
		return err
	}
	return responseProcesser(resp, respContainer)
}

func (c *clientw) Patch(endpoint string, bodyParams map[string]interface{}, respContainer interface{}) error {
	resp, err := c.client.R().SetBody(bodyParams).Patch(endpoint)
	if err != nil {
		return err
	}
	return responseProcesser(resp, respContainer)
}
