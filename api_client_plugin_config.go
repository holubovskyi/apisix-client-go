package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type PluginConfig struct {
	ID          *string                 `json:"id"`
	Description *string                 `json:"desc,omitempty"`
	Labels      *map[string]string      `json:"labels,omitempty"`
	Plugins     *map[string]interface{} `json:"plugins"`
}

type PluginConfigAPIResponse struct {
	Key   string       `json:"key"`
	Value PluginConfig `json:"value"`
}

// GetPluginConfig - Returns a plugin config
func (c *ApiClient) GetPluginConfig(configID string) (*PluginConfig, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/plugin_configs/%s", c.Endpoint, configID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := PluginConfigAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreatePluginConfig - Creates a new plugin config
func (c *ApiClient) CreatePluginConfig(configID string, config PluginConfig) (*PluginConfig, error) {
	rb, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/plugin_configs/%s", c.Endpoint, configID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := PluginConfigAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// UpdatePluginConfig - Updates a plugin config
func (c *ApiClient) UpdatePluginConfig(configID string, config PluginConfig) (*PluginConfig, error) {
	rb, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/plugin_configs/%s", c.Endpoint, configID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := PluginConfigAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeletePluginConfig - Deletes a plugin config
func (c *ApiClient) DeletePluginConfig(configID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/plugin_configs/%s", c.Endpoint, configID), nil)
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
