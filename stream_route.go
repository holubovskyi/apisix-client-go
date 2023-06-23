package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type StreamRoute struct {
	ID         *string `json:"id,omitempty"`
	UpstreamId *string `json:"upstream_id,omitempty"`
	RemoteAddr *string `json:"remote_addr,omitempty"`
	ServerAddr *string `json:"server_addr,omitempty"`
	ServerPort *int64  `json:"server_port,omitempty"`
	SNI        *string `json:"sni,omitempty"`
}

type StreamRouteAPIResponse struct {
	Key   string      `json:"key"`
	Value StreamRoute `json:"value"`
}

// GetStreamRoute - Returns a specific stream route
func (c *ApiClient) GetStreamRoute(routeID string) (*StreamRoute, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/stream_routes/%s", c.Endpoint, routeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := StreamRouteAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateStreamRoute - Creates a steam route
func (c *ApiClient) CreateStreamRoute(route StreamRoute) (*StreamRoute, error) {
	rb, err := json.Marshal(route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apisix/admin/stream_routes/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := StreamRouteAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil
}

// UpdateStreamRoute - Updates a stream route
func (c *ApiClient) UpdateStreamRoute(routeID string, route StreamRoute) (*StreamRoute, error) {
	rb, err := json.Marshal(route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/stream_routes/%s", c.Endpoint, routeID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := StreamRouteAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteStreamRoute - Deletes a stream route
func (c *ApiClient) DeleteStreamRoute(routeID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/stream_routes/%s", c.Endpoint, routeID), nil)
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
