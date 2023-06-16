package api_client

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strings"
)

type Upstream struct {
	ID              string                     `json:"id,omitempty"`
	Type            string                     `json:"type"`
	ServiceName     string                     `json:"service_name,omitempty"`
	DiscoveryType   string                     `json:"discovery_type,omitempty"`
	Timeout         *TimeoutType               `json:"timeout,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Desc            string                     `json:"desc,omitempty"`
	PassHost        string                     `json:"pass_host,omitempty"`
	Scheme          string                     `json:"scheme,omitempty"`
	Retries         *uint                       `json:"retries,omitempty"`
	RetryTimeout    *uint                       `json:"retry_timeout,omitempty"`
	Labels          map[string]string          `json:"labels"`
	UpstreamHost    string                     `json:"upstream_host,omitempty"`
	HashOn          string                     `json:"hash_on,omitempty"`
	Key             string                     `json:"key,omitempty"`
	KeepalivePool   *UpstreamKeepAlivePoolType `json:"keepalive_pool,omitempty"`
	TLSClientCertID string                     `json:"tls.client_cert_id,omitempty"`
	Checks          *UpstreamChecksType        `json:"checks,omitempty"`
	Nodes           *[]UpstreamNodeType        `json:"nodes,omitempty"`
}

type TimeoutType struct {
	Connect uint `json:"connect"`
	Send    uint `json:"send"`
	Read    uint `json:"read"`
}

type UpstreamKeepAlivePoolType struct {
	Size        uint `json:"size"`
	IdleTimeout uint `json:"idle_timeout"`
	Requests    uint `json:"requests"`
}

type UpstreamChecksType struct {
	Active  *UpstreamChecksActiveType  `json:"active,omitempty"`
	Passive *UpstreamChecksPassiveType `json:"passive,omitempty"`
}

type UpstreamChecksActiveType struct {
	Type                   string                             `json:"type"`
	Timeout                uint                               `json:"timeout"`
	Concurrency            uint                               `json:"concurrency"`
	HTTPPath               string                             `json:"http_path"`
	Host                   string                             `json:"host,omitempty"`
	Port                   uint                               `json:"port,omitempty"`
	HTTPSVerifyCertificate bool                               `json:"https_verify_certificate"`
	ReqHeaders             []string                           `json:"req_headers,omitempty"`
	Healthy                *UpstreamChecksActiveHealthyType   `json:"healthy"`
	Unhealthy              *UpstreamChecksActiveUnhealthyType `json:"unhealthy"`
}

type UpstreamChecksActiveHealthyType struct {
	Interval     uint   `json:"interval"`
	HTTPStatuses []uint `json:"http_statuses"`
	Successes    uint   `json:"successes"`
}

type UpstreamChecksActiveUnhealthyType struct {
	Interval     uint   `json:"interval"`
	HTTPStatuses []uint `json:"http_statuses"`
	TCPFailures  uint   `json:"tcp_failures"`
	Timeouts     uint   `json:"timeouts"`
	HTTPFailures uint   `json:"http_failures"`
}

type UpstreamChecksPassiveType struct {
	Healthy   *UpstreamChecksPassiveHealthyType   `json:"healthy"`
	Unhealthy *UpstreamChecksPassiveUnhealthyType `json:"unhealthy"`
}

type UpstreamChecksPassiveHealthyType struct {
	HTTPStatuses []uint `json:"http_statuses"`
	Successes    uint   `json:"successes"`
}

type UpstreamChecksPassiveUnhealthyType struct {
	HTTPStatuses []uint `json:"http_statuses"`
	TCPFailures  uint   `json:"tcp_failures"`
	Timeouts     uint   `json:"timeouts"`
	HTTPFailures uint   `json:"http_failures"`
}

type UpstreamNodeType struct {
	Host   string `json:"host"`
	Port   uint   `json:"port"`
	Weight uint   `json:"weight"`
}

// func (client ApiClient) GetUpstream(id string) (map[string]interface{}, error) {
// 	return client.RunObject("GET", "/apisix/admin/upstreams/"+id, nil)
// }

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

//	func (client ApiClient) CreateUpstream(data map[string]interface{}) (map[string]interface{}, error) {
//		return client.RunObject("POST", "/apisix/admin/upstreams/", &data)
//	}
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

func (client ApiClient) UpdateUpstream(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/apisix/admin/upstreams/"+id, &data)
}

func (client ApiClient) DeleteUpstream(id string) (err error) {
	return client.Delete("/apisix/admin/upstreams/" + id)
}
