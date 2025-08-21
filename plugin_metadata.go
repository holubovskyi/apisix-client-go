package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type PluginMetadata struct {
	Id       *string                 `json:"-"`
	Metadata *map[string]interface{} `json:"-"`
}

type PluginMetadataAPIResponse struct {
	Key   string         `json:"key"`
	Value PluginMetadata `json:"value"`
}

// normalizeJSON ensures consistent JSON key ordering
func normalizeJSON(data map[string]interface{}) map[string]interface{} {
	normalized := make(map[string]interface{})
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := data[key]
		switch v := value.(type) {
		case map[string]interface{}:
			normalized[key] = normalizeJSON(v)
		case []interface{}:
			normalizedArray := make([]interface{}, len(v))
			for i, item := range v {
				if itemMap, ok := item.(map[string]interface{}); ok {
					normalizedArray[i] = normalizeJSON(itemMap)
				} else {
					normalizedArray[i] = item
				}
			}
			normalized[key] = normalizedArray
		default:
			normalized[key] = value
		}
	}
	return normalized
}

// MarshalJSON implements custom JSON marshaling - passes user input directly
func (pm PluginMetadata) MarshalJSON() ([]byte, error) {
	result := make(map[string]interface{})

	fmt.Printf("DEBUG MarshalJSON: pm.Id = %v\n", pm.Id)
	fmt.Printf("DEBUG MarshalJSON: pm.Metadata = %v\n", pm.Metadata)

	// Add ID if present
	if pm.Id != nil {
		result["id"] = *pm.Id
		fmt.Printf("DEBUG MarshalJSON: Added ID: %s\n", *pm.Id)
	}

	// Add metadata content directly as provided by user
	if pm.Metadata != nil && len(*pm.Metadata) > 0 {
		normalizedMetadata := normalizeJSON(*pm.Metadata)
		fmt.Printf("DEBUG MarshalJSON: Metadata to add: %+v\n", normalizedMetadata)

		// Add all metadata fields directly to the result
		for key, value := range normalizedMetadata {
			result[key] = value
			fmt.Printf("DEBUG MarshalJSON: Added field %s: %+v\n", key, value)
		}
	} else {
		fmt.Printf("DEBUG MarshalJSON: No metadata to add\n")
	}

	fmt.Printf("DEBUG MarshalJSON: Final result: %+v\n", result)
	finalJSON, err := json.Marshal(result)
	fmt.Printf("DEBUG MarshalJSON: Final JSON: %s\n", string(finalJSON))
	return finalJSON, err
}

// UnmarshalJSON implements custom JSON unmarshaling
func (pm *PluginMetadata) UnmarshalJSON(data []byte) error {
	temp := make(map[string]interface{})
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	metadata := make(map[string]interface{})

	// Extract the plugin ID
	if idValue, exists := temp["id"]; exists {
		if idStr, ok := idValue.(string); ok {
			pm.Id = &idStr
		}
	}

	// Extract all fields except system fields as metadata
	systemFields := map[string]bool{
		"id":  true,
		"key": true,
	}

	for key, value := range temp {
		if !systemFields[key] {
			metadata[key] = value
		}
	}

	if len(metadata) > 0 {
		normalizedMetadata := normalizeJSON(metadata)
		pm.Metadata = &normalizedMetadata
	}

	return nil
}

// CreatePluginMetadata - creates a new plugin metadata
func (c *ApiClient) CreatePluginMetadata(Id string, metadata PluginMetadata) (*PluginMetadata, error) {
	// Ensure plugin name is set
	metadata.Id = &Id

	// Marshal with custom method
	rb, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	// Debug: Log what we're sending to APISIX
	fmt.Printf("DEBUG: Sending to APISIX for plugin %s: %s\n", Id, string(rb))

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/plugin_metadata/%s", c.Endpoint, Id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// Debug: Log APISIX response
	fmt.Printf("DEBUG: APISIX response: %s\n", string(body))

	return &metadata, nil
}

// GetPluginMetadata - retrieves a plugin metadata
func (c *ApiClient) GetPluginMetadata(Id string) (*PluginMetadata, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/plugin_metadata/%s", c.Endpoint, Id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := PluginMetadataAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

// UpdatePluginMetadata - updates an existing plugin metadata
func (c *ApiClient) UpdatePluginMetadata(Id string, metadata PluginMetadata) (*PluginMetadata, error) {
	return c.CreatePluginMetadata(Id, metadata)
}

// DeletePluginMetadata - deletes a plugin metadata
func (c *ApiClient) DeletePluginMetadata(Id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/plugin_metadata/%s", c.Endpoint, Id), nil)
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
