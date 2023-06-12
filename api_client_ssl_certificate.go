package api_client

import (
	"encoding/json"
	//	"errors"
	"fmt"
	"net/http"
	// "strings"
)

type SSLCertificate struct {
	ID          string            `json:"id"`
	Certificate string            `json:"cert"`
	PrivateKey  string            `json:"key"`
	SNIs        []string          `json:"snis"`
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

	sslCertificate := SSLCertificate{}
	err = json.Unmarshal(body, &sslCertificate)
	if err != nil {
		return nil, err
	}

	return &sslCertificate, nil
}

func (client ApiClient) CreateSslCertificate(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/apisix/admin/ssls/", &data)
}

func (client ApiClient) UpdateSslCertificate(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/apisix/admin/ssls/"+id, &data)
}

func (client ApiClient) DeleteSslCertificate(id string) (err error) {
	return client.Delete("/apisix/admin/ssls/" + id)
}
