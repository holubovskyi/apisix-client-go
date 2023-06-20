package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Upstream struct {
	ID              *string                    `json:"id,omitempty"`
	Type            *string                    `json:"type"`
	ServiceName     *string                    `json:"service_name,omitempty"`
	DiscoveryType   *string                    `json:"discovery_type,omitempty"`
	Timeout         *TimeoutType               `json:"timeout,omitempty"`
	Name            *string                    `json:"name,omitempty"`
	Desc            *string                    `json:"desc,omitempty"`
	PassHost        *string                    `json:"pass_host,omitempty"`
	Scheme          *string                    `json:"scheme,omitempty"`
	Retries         *int64                     `json:"retries,omitempty"`
	RetryTimeout    *int64                     `json:"retry_timeout,omitempty"`
	Labels          map[string]string          `json:"labels"`
	UpstreamHost    *string                    `json:"upstream_host,omitempty"`
	HashOn          *string                    `json:"hash_on,omitempty"`
	Key             *string                    `json:"key,omitempty"`
	KeepalivePool   *UpstreamKeepAlivePoolType `json:"keepalive_pool,omitempty"`
	TLSClientCertID *string                    `json:"tls.client_cert_id,omitempty"`
	Checks          *UpstreamChecksType        `json:"checks,omitempty"`
	Nodes           *[]UpstreamNodeType        `json:"nodes,omitempty"`
}

type UpstreamUpdate struct {
	ID              *string                    `json:"id,omitempty"`
	Type            *string                    `json:"type"`
	ServiceName     *string                    `json:"service_name,omitempty"`
	DiscoveryType   *string                    `json:"discovery_type"`
	Timeout         *TimeoutType               `json:"timeout"`
	Name            *string                    `json:"name"`
	Desc            *string                    `json:"desc"`
	PassHost        *string                    `json:"pass_host"`
	Scheme          *string                    `json:"scheme"`
	Retries         *int64                     `json:"retries"`
	RetryTimeout    *int64                     `json:"retry_timeout"`
	Labels          map[string]string          `json:"labels"`
	UpstreamHost    *string                    `json:"upstream_host"`
	HashOn          *string                    `json:"hash_on"`
	Key             *string                    `json:"key"`
	KeepalivePool   *UpstreamKeepAlivePoolType `json:"keepalive_pool"`
	TLSClientCertID *string                    `json:"tls.client_cert_id"`
	Checks          *UpstreamChecksTypeUpdate  `json:"checks"`
	Nodes           *[]UpstreamNodeType        `json:"nodes,omitempty"`
}

type TimeoutType struct {
	Connect int64 `json:"connect"`
	Send    int64 `json:"send"`
	Read    int64 `json:"read"`
}

type UpstreamKeepAlivePoolType struct {
	Size        int64 `json:"size"`
	IdleTimeout int64 `json:"idle_timeout"`
	Requests    int64 `json:"requests"`
}

type UpstreamChecksType struct {
	Active  *UpstreamChecksActiveType  `json:"active,omitempty"`
	Passive *UpstreamChecksPassiveType `json:"passive,omitempty"`
}

type UpstreamChecksTypeUpdate struct {
	Active  *UpstreamChecksActiveTypeUpdate  `json:"active"`
	Passive *UpstreamChecksPassiveTypeUpdate `json:"passive"`
}

type UpstreamChecksActiveType struct {
	Type                   string                             `json:"type"`
	Timeout                int64                              `json:"timeout"`
	Concurrency            int64                              `json:"concurrency"`
	HTTPPath               string                             `json:"http_path"`
	Host                   *string                            `json:"host,omitempty"`
	Port                   *int64                             `json:"port,omitempty"`
	HTTPSVerifyCertificate bool                               `json:"https_verify_certificate"`
	ReqHeaders             []string                           `json:"req_headers,omitempty"`
	Healthy                *UpstreamChecksActiveHealthyType   `json:"healthy,omitempty"`
	Unhealthy              *UpstreamChecksActiveUnhealthyType `json:"unhealthy,omitempty"`
}

type UpstreamChecksActiveTypeUpdate struct {
	Type                   string                             `json:"type"`
	Timeout                int64                              `json:"timeout"`
	Concurrency            int64                              `json:"concurrency"`
	HTTPPath               string                             `json:"http_path"`
	Host                   *string                            `json:"host"`
	Port                   *int64                             `json:"port"`
	HTTPSVerifyCertificate bool                               `json:"https_verify_certificate"`
	ReqHeaders             []string                           `json:"req_headers"`
	Healthy                *UpstreamChecksActiveHealthyType   `json:"healthy"`
	Unhealthy              *UpstreamChecksActiveUnhealthyType `json:"unhealthy"`
}

type UpstreamChecksActiveHealthyType struct {
	Interval     int64   `json:"interval,omitempty"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	Successes    int64   `json:"successes,omitempty"`
}

type UpstreamChecksActiveUnhealthyType struct {
	Interval     int64   `json:"interval,omitempty"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	TCPFailures  int64   `json:"tcp_failures,omitempty"`
	Timeouts     int64   `json:"timeouts,omitempty"`
	HTTPFailures int64   `json:"http_failures,omitempty"`
}

type UpstreamChecksPassiveType struct {
	Healthy   *UpstreamChecksPassiveHealthyType   `json:"healthy,omitempty"`
	Unhealthy *UpstreamChecksPassiveUnhealthyType `json:"unhealthy,omitempty"`
}

type UpstreamChecksPassiveTypeUpdate struct {
	Healthy   *UpstreamChecksPassiveHealthyType   `json:"healthy"`
	Unhealthy *UpstreamChecksPassiveUnhealthyType `json:"unhealthy"`
}

type UpstreamChecksPassiveHealthyType struct {
	HTTPStatuses []int64 `json:"http_statuses"`
	Successes    int64   `json:"successes"`
}

type UpstreamChecksPassiveUnhealthyType struct {
	HTTPStatuses []int64 `json:"http_statuses"`
	TCPFailures  int64   `json:"tcp_failures"`
	Timeouts     int64   `json:"timeouts"`
	HTTPFailures int64   `json:"http_failures"`
}

type UpstreamNodeType struct {
	Host   string `json:"host"`
	Port   int64  `json:"port"`
	Weight int64  `json:"weight"`
}

type UpstreamAPIResponse struct {
	Key   string   `json:"key"`
	Value Upstream `json:"value"`
}

// GetUpstream - Return a specific upstream
func (c *ApiClient) GetUpstream(upstreamID string) (*Upstream, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apisix/admin/upstreams/%s", c.Endpoint, upstreamID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	getResponse := UpstreamAPIResponse{}
	err = json.Unmarshal(body, &getResponse)
	if err != nil {
		return nil, err
	}

	return &getResponse.Value, nil
}

func (c *ApiClient) CreateUpstream(upstream Upstream) (*Upstream, error) {
	rb, err := json.Marshal(upstream)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apisix/admin/upstreams/", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	creationResponse := UpstreamAPIResponse{}
	err = json.Unmarshal(body, &creationResponse)
	if err != nil {
		return nil, err
	}

	return &creationResponse.Value, nil
}

// UpdateUpstream - Updates an upstream
func (c *ApiClient) UpdateUpstream(upstreamID string, upstream UpstreamUpdate) (*Upstream, error) {
	rb, err := json.Marshal(upstream)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/apisix/admin/upstreams/%s", c.Endpoint, upstreamID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateResponse := UpstreamAPIResponse{}
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse.Value, nil
}

// DeleteUpstream - Deletes an upstream
func (c *ApiClient) DeleteUpstream(upstreamID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apisix/admin/upstreams/%s", c.Endpoint, upstreamID), nil)
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
