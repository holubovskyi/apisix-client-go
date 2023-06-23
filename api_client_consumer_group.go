package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ConsumerGroup struct {
	ID          *string                 `json:"id"`
	Description *string                 `json:"desc"`
	Labels      *map[string]string      `json:"labels"`
	Plugins     *map[string]interface{} `json:"plugins"`
}

type ConsumerGroupAPIResponse struct {
	Key   string        `json:"key"`
	Value ConsumerGroup `json:"value"`
}

// GetConsumerGroup - Returns a consumer group
func (c *ApiClient) GetConsumerGroup(groupID string) (*ConsumerGroup, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/consumer_groups/%s", c.Endpoint, groupID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := ConsumerGroupAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateConsumerGroup - Creates a new consumer group
func (c *ApiClient) CreateConsumerGroup(groupID string, group ConsumerGroup) (*ConsumerGroup, error) {
	rb, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/consumer_groups/%s", c.Endpoint, groupID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := ConsumerGroupAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// UpdateConsumerGroup - Updates a consumer group
func (c *ApiClient) UpdateConsumerGroup(groupID string, group ConsumerGroup) (*ConsumerGroup, error) {
	rb, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/consumer_groups/%s", c.Endpoint, groupID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := ConsumerGroupAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteConsumerGroup - Deletes a consumer group
func (c *ApiClient) DeleteConsumerGroup(groupID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/consumer_groups/%s", c.Endpoint, groupID), nil)
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
