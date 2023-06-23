package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Consumer struct {
	Username    *string                 `json:"username,omitempty"`
	Description *string                 `json:"desc,omitempty"`
	Labels      *map[string]string      `json:"labels,omitempty"`
	Plugins     *map[string]interface{} `json:"plugins,omitempty"`
	GroupId     *string                 `json:"group_id,omitempty"`
}

type ConsumerAPIResponse struct {
	Key   string   `json:"key"`
	Value Consumer `json:"value"`
}

// GetConsumer - Returns a specific consumer
func (c *ApiClient) GetConsumer(consumerName string) (*Consumer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/consumers/%s", c.Endpoint, consumerName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := ConsumerAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateConsumer - Creates a consumer
func (c *ApiClient) CreateConsumer(consumer Consumer) (*Consumer, error) {
	rb, err := json.Marshal(consumer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/consumers/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := ConsumerAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil

}

// UpdateConsumer - Updates a consumer
func (c *ApiClient) UpdateConsumer(consumer Consumer) (*Consumer, error) {
	rb, err := json.Marshal(consumer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/consumers/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := ConsumerAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteConsumer - Deletes a consumer
func (c *ApiClient) DeleteConsumer(consumerName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/consumers/%s", c.Endpoint, consumerName), nil)
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
