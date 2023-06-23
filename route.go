package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Route struct {
	ID              *string                 `json:"id,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Description     *string                 `json:"desc,omitempty"`
	URI             *string                 `json:"uri,omitempty"`
	URIS            *[]string               `json:"uris,omitempty"`
	Host            *string                 `json:"host,omitempty"`
	Hosts           *[]string               `json:"hosts,omitempty"`
	RemoteAddr      *string                 `json:"remote_addr,omitempty"`
	RemoteAddrs     *[]string               `json:"remote_addrs,omitempty"`
	Methods         *[]string               `json:"methods,omitempty"`
	Priority        *int64                  `json:"priority,omitempty"`
	Vars            *[]interface{}          `json:"vars,omitempty"`
	FilterFunc      *string                 `json:"filter_func,omitempty"`
	Plugins         *map[string]interface{} `json:"plugins,omitempty"`
	Script          *string                 `json:"script,omitempty"`
	UpstreamId      *string                 `json:"upstream_id,omitempty"`
	ServiceId       *string                 `json:"service_id,omitempty"`
	PluginConfigId  *string                 `json:"plugin_config_id,omitempty"`
	Labels          *map[string]string      `json:"labels,omitempty"`
	Timeout         *TimeoutType            `json:"timeout,omitempty"`
	EnableWebsocket *bool                   `json:"enable_websocket,omitempty"`
	Status          *int64                  `json:"status,omitempty"`
}

type RouteAPIResponse struct {
	Key   string `json:"key"`
	Value Route  `json:"value"`
}

// GetRoute - Returns a specific route
func (c *ApiClient) GetRoute(routeID string) (*Route, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/routes/%s", c.Endpoint, routeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := RouteAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateRoute - Creates a route
func (c *ApiClient) CreateRoute(route Route) (*Route, error) {
	rb, err := json.Marshal(route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apisix/admin/routes/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := RouteAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil
}

// UpdateRoute - Updates a route
func (c *ApiClient) UpdateRoute(routeID string, route Route) (*Route, error) {
	rb, err := json.Marshal(route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/routes/%s", c.Endpoint, routeID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := RouteAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteRoute - Deletes a route
func (c *ApiClient) DeleteRoute(routeID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/routes/%s", c.Endpoint, routeID), nil)
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
