package api_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "time"
)

type AddHeadersRoundtripper struct {
	Headers http.Header
	Nested  http.RoundTripper
}

func (h AddHeadersRoundtripper) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, vs := range h.Headers {
		for _, v := range vs {
			r.Header.Add(k, v)
		}
	}
	return h.Nested.RoundTrip(r)
}

type ApiClient struct {
	Endpoint   string
	HTTPClient *http.Client
	APIKey     string
}

func NewClient(endpoint, apiKey *string) (*ApiClient, error) {

	if endpoint == nil {
		return nil, fmt.Errorf("the value of the endpoint is not provided")
	}

	if apiKey == nil {
		return nil, fmt.Errorf("the value of the API Key is not provided")
	}

	apiClient := http.DefaultClient
	headers := make(http.Header, 0)
	headers.Add("X-API-KEY", *apiKey)

	apiClient.Transport = AddHeadersRoundtripper{
		Headers: headers,
		Nested:  http.DefaultTransport,
	}

	c := ApiClient{
		HTTPClient: apiClient,
		Endpoint:   *endpoint,
		APIKey:     *apiKey,
	}

	return &c, nil
}

func (c *ApiClient) doRequest(req *http.Request) ([]byte, error) {
	key := c.APIKey

	req.Header.Set("X-API-KEY", key)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// if status code >= 400
	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
