package api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
//	"time"
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

func GetCl(apiKey string, endpoint string) (*ApiClient, error) {
	apiClient := http.DefaultClient
	headers := make(http.Header, 0)
	headers.Add("X-API-KEY", apiKey)
	apiClient.Transport = AddHeadersRoundtripper{
		Headers: headers,
		Nested:  http.DefaultTransport,
	}

	c := ApiClient{
		Endpoint:   endpoint,
		HTTPClient: http.DefaultClient,
	}

	return &c, nil
}

func NewClient(endpoint, apiKey *string) (*ApiClient, error) {

	if endpoint == nil {
		return nil, fmt.Errorf("The value of the endpoint is not provided")
	}

	if apiKey == nil {
		return nil, fmt.Errorf("The value of the API Key is not provided")
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

func parseHttpResult(res *http.Response, body []byte) (int, []byte, error) {
	log.Printf("[DEBUG] Got response: %#v", res)
	log.Printf("[DEBUG] Got statuscode: %#v", res.StatusCode)
	log.Printf("[DEBUG] Got body: %v", string(body))

	var result map[string]interface{}
	//var result interface{}
	err := json.Unmarshal(body, &result)

	if err != nil {
		return res.StatusCode, []byte{}, err
	}

	if res.StatusCode >= 400 {

		errorMessage := "No message"
		//if result["error_msg"] != nil {
		//	errorMessage = result["error_msg"].(string)
		//}

		return res.StatusCode, []byte(errorMessage), fmt.Errorf("can't make request, cause: %v", errorMessage)
	}

	//node := result["node"].(map[string]interface{})
	//value := node["value"].(map[string]interface{})
	value := result["value"].(map[string]interface{})
	item, err := json.Marshal(value)
	return res.StatusCode, item, err
}

func (client ApiClient) Get(path string) (int, []byte, error) {
	url := client.Endpoint + path
	res, err := client.HTTPClient.Get(url)

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Post(path string, jsonBytes []byte) (int, []byte, error) {
	apiUrl := client.Endpoint + path

	log.Printf("[DEBUG] SEND POST -> %v ->  %v", path, string(jsonBytes))
	res, err := client.HTTPClient.Post(apiUrl, "application/json; charset=utf-8", bytes.NewReader(jsonBytes))

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Patch(path string, jsonBytes []byte) (int, []byte, error) {
	apiUrl := client.Endpoint + path

	log.Printf("[DEBUG] PATCH SEND %v -> %v", apiUrl, string(jsonBytes))
	req, err := http.NewRequest("PATCH", apiUrl, bytes.NewReader(jsonBytes))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.HTTPClient.Do(req)

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Put(path string, jsonBytes []byte) (int, []byte, error) {
	apiUrl := client.Endpoint + path

	log.Printf("[DEBUG] SEND PUT to %v -> %v", apiUrl, string(jsonBytes))
	req, err := http.NewRequest("PUT", apiUrl, bytes.NewReader(jsonBytes))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.HTTPClient.Do(req)

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Delete(path string) error {
	apiUrl := client.Endpoint + path

	req, err := http.NewRequest("DELETE", apiUrl, nil)
	if err != nil {
		return err
	}
	res, err := client.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	return err
}

func (client ApiClient) RunObject(method string, url string, data *map[string]interface{}) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	var statusCode int
	var body []byte
	var err error
	switch method {
	case "GET":
		statusCode, body, err = client.Get(url)
	case "POST":
		b, errA := json.Marshal(*data)
		if errA == nil {
			statusCode, body, err = client.Post(url, b)
		}
		err = errA
	case "PUT":
		b, errA := json.Marshal(*data)
		if errA == nil {
			statusCode, body, err = client.Put(url, b)
		}
		err = errA

	case "PATCH":
		b, errA := json.Marshal(*data)
		if errA == nil {
			statusCode, body, err = client.Patch(url, b)
		}
		err = errA
	}

	if err != nil {
		return response, err
	}

	if statusCode >= 400 {
		return response, fmt.Errorf("got error: %v", string(body))
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (c *ApiClient) doRequest(req *http.Request, apiKey *string) ([]byte, error) {
	key := c.APIKey

	if apiKey != nil {
		key = *apiKey
	}

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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
