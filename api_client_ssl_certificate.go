package api_client

import (
	"encoding/json"
	"strings"
	//	"errors"
	"fmt"
	"net/http"
	// "strings"
)

type SSLCertificate struct {
	ID          string            `json:"id"`
	Status      uint              `json:"status"`
	Certificate string            `json:"cert"`
	PrivateKey  string            `json:"key"`
	SNIs        []string          `json:"snis"`
	Type        string            `json:"type"`
	Labels      map[string]string `json:"labels"`
}

// func (client ApiClient) GetSslCertificate(id string) (map[string]interface{}, error) {
// 	return client.RunObject("GET", "/apisix/admin/ssls/"+id, nil)
// }

func (c *ApiClient) GetSslCertificate(certificateID string, apiKey *string) (*SSLCertificate, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/ssls/%s", c.Endpoint, certificateID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, apiKey)
	if err != nil {
		return nil, err
	}

	certificate := SSLCertificate{}
	err = json.Unmarshal(body, &certificate)
	if err != nil {
		return nil, err
	}

	return &certificate, nil
}

// func (client ApiClient) CreateSslCertificate(data map[string]interface{}) (map[string]interface{}, error) {
// 	return client.RunObject("POST", "/apisix/admin/ssls/", &data)
// }

func (c *ApiClient) CreateSslCertificate(sslCertificate SSLCertificate, apiKey *string) (*SSLCertificate, error) {
	rb, err := json.Marshal(sslCertificate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s//apisix/admin/ssls/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, apiKey)
	if err != nil {
		return nil, err
	}

	certificate := SSLCertificate{}
	err = json.Unmarshal(body, &certificate)
	if err != nil {
		return nil, err
	}

	return &certificate, nil
}

func (client ApiClient) UpdateSslCertificate(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/apisix/admin/ssls/"+id, &data)
}

func (client ApiClient) DeleteSslCertificate(id string) (err error) {
	return client.Delete("/apisix/admin/ssls/" + id)
}
