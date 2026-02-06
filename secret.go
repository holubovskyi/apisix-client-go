package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Secret interface {
	IsSecret()
}

type SecretManager string

const (
	Vault SecretManager = "vault"
	AWS   SecretManager = "aws"
	GCP   SecretManager = "gcp"
)

type VaultSecret struct {
	Uri       *string `json:"uri"`
	Prefix    *string `json:"prefix"`
	Token     *string `json:"token"`
	Namespace *string `json:"namespace,omitempty"`
}

func (s *VaultSecret) IsSecret() {}

type AWSSecret struct {
	AccessKeyId     *string `json:"access_key_id"`
	SecretAccessKey *string `json:"secret_access_key"`
	SessionToken    *string `json:"session_token,omitempty"`
	Region          *string `json:"region,omitempty"`
	EndpointUrl     *string `json:"endpoint_url,omitempty"`
}

func (s *AWSSecret) IsSecret() {}

type GCPSecret struct {
	AuthConfig *AuthConfigType `json:"auth_config,omitempty"`
	AuthFile   *string         `json:"auth_file,omitempty"`
	SslVerify  *bool           `json:"ssl_verify,omitempty"`
}

func (s *GCPSecret) IsSecret() {}

type AuthConfigType struct {
	ClientEmail *string `json:"client_email"`
	PrivateKey  *string `json:"private_key"`
	ProjectId   *string `json:"project_id"`
	TokenUri    *string `json:"token_uri,omitempty"`
	EntriesUri  *string `json:"entries_uri,omitempty"`
	Scope       *string `json:"scope,omitempty"`
}

type SecretAPIResponse struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

func SecretFactory(secretManager SecretManager) (Secret, error) {
	switch secretManager {
	case Vault:
		return &VaultSecret{}, nil
	case AWS:
		return &AWSSecret{}, nil
	case GCP:
		return &GCPSecret{}, nil
	default:
		return nil, fmt.Errorf("unsupported secret manager: %s", secretManager)
	}
}

// GetSecret - Returns a specific secret
func (c *ApiClient) GetSecret(secretManager SecretManager, secretID string) (Secret, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/secrets/%s/%s", c.Endpoint, secretManager, secretID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := SecretAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	gotSecret, err := SecretFactory(secretManager)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(getResponse.Value, gotSecret)
	if err != nil {
		return nil, err
	}

	return gotSecret, nil
}

// CreateSecret - Create a secret
func (c *ApiClient) CreateSecret(secretManager SecretManager, secret Secret) (Secret, error) {
	rb, err := json.Marshal(secret)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apisix/admin/secrets/%s", c.Endpoint, secretManager), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := SecretAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	createdSecret, err := SecretFactory(secretManager)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(creationResponse.Value, createdSecret)
	if err != nil {
		return nil, err
	}

	return createdSecret, nil
}

// UpdateSecret - Updates a Secret
func (c *ApiClient) UpdateSecret(secretManager SecretManager, secretID string, secret Secret) (Secret, error) {
	rb, err := json.Marshal(secret)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/apisix/admin/secrets/%s/%s", c.Endpoint, secretManager, secretID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := SecretAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	updatedSecret, err := SecretFactory(secretManager)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(updateResponse.Value, updatedSecret)
	if err != nil {
		return nil, err
	}

	return updatedSecret, nil
}

// DeleteSecret - Deletes an secret
func (c *ApiClient) DeleteSecret(secretManager SecretManager, secretID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/secrets/%s/%s", c.Endpoint, secretManager, secretID), nil)
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
