package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type SSLCertificate struct {
	ID          string            `json:"id,omitempty"`
	Status      uint              `json:"status"`
	Certificate string            `json:"cert"`
	PrivateKey  string            `json:"key"`
	SNIs        []string          `json:"snis"`
	Type        string            `json:"type"`
	Labels      map[string]string `json:"labels,omitempty"`
}

type SSLCertificateAPIResponse struct {
	Key   string         `json:"key"`
	Value SSLCertificate `json:"value"`
}

type DeleteResponse struct {
	Key     string `json:"key"`
	Deleted string `json:"deleted"`
}

// GetSslCertificate - Returns a specifc certificate
func (c *ApiClient) GetSslCertificate(certificateID string) (*SSLCertificate, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/ssls/%s", c.Endpoint, certificateID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := SSLCertificateAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateSslCertificate - Create new certificate
func (c *ApiClient) CreateSslCertificate(sslCertificate SSLCertificate) (*SSLCertificate, error) {
	rb, err := json.Marshal(sslCertificate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apisix/admin/ssls/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := SSLCertificateAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil
}

// UpdateSslCertificate - Updates a certificate
func (c *ApiClient) UpdateSslCertificate(certificateID string, sslCertificate SSLCertificate) (*SSLCertificate, error) {
	rb, err := json.Marshal(sslCertificate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/ssls/%s", c.Endpoint, certificateID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := SSLCertificateAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteSslCertificate - Deletes a certificate
func (c *ApiClient) DeleteSslCertificate(certificateID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/ssls/%s", c.Endpoint, certificateID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	deleteResponse := DeleteResponse{}
	err = json.Unmarshal(body, &deleteResponse)
	if err != nil {
		return err
	}

	if deleteResponse.Deleted != "1" {
		return errors.New(string(body))
	}

	return nil
}
