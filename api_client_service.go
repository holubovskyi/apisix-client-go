package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Service struct {
	ID              *string                 `json:"id,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Description     *string                 `json:"desc,omitempty"`
	EnableWebsocket *bool                   `json:"enable_websocket,omitempty"`
	Hosts           *[]string               `json:"hosts,omitempty"`
	Labels          *map[string]string      `json:"labels,omitempty"`
	Plugins         *map[string]interface{} `json:"plugins,omitempty"`
	UpstreamId      *string                 `json:"upstream_id,omitempty"`
}

type ServiceAPIResponse struct {
	Key   string  `json:"key"`
	Value Service `json:"value"`
}

// GetService- Returns a specific service
func (c *ApiClient) GetService(serviceID string) (*Service, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/services/%s", c.Endpoint, serviceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := ServiceAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateService - Creates a service
func (c *ApiClient) CreateService(service Service) (*Service, error) {
	rb, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apisix/admin/services/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := ServiceAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil
}

// UpdateService - Updates a service
func (c *ApiClient) UpdateService(serviceID string, service Service) (*Service, error) {
	rb, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/services/%s", c.Endpoint, serviceID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := ServiceAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteService - Deletes a service
func (c *ApiClient) DeleteService(serviceID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/services/%s", c.Endpoint, serviceID), nil)
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
