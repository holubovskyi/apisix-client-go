package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type GlobalRule struct {
	ID      *string                 `json:"id"`
	Plugins *map[string]interface{} `json:"plugins"`
}

type GlobalRuleResponse struct {
	Key   string     `json:"key"`
	Value GlobalRule `json:"value"`
}

// GetGlobalRule - Returns a specific global rule
func (c *ApiClient) GetGlobalRule(ruleID string) (*GlobalRule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/global_rules/%s", c.Endpoint, ruleID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := GlobalRuleResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// CreateGlobalRule - Creates a new global rule
func (c *ApiClient) CreateGlobalRule(ruleID string, rule GlobalRule) (*GlobalRule, error) {
	rb, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/global_rules/%s", c.Endpoint, ruleID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := GlobalRuleResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil

}

// UpdateGlobalRule - Updates a global rule
func (c *ApiClient) UpdateGlobalRule(ruleID string, rule GlobalRule) (*GlobalRule, error) {
	rb, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/global_rules/%s", c.Endpoint, ruleID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := GlobalRuleResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil

}

// DeleteGlobalRule - Deletes a global rule
func (c *ApiClient) DeleteGlobalRule(ruleID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/global_rules/%s", c.Endpoint, ruleID), nil)
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
