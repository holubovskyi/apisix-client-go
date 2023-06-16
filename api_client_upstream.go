package api_client

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"strings"
)

type Upstream struct {
	ID              string                    `json:"id,omitempty"`
	Type            string                    `json:"type"`
	ServiceName     string                    `json:"service_name"`
	DiscoveryType   string                    `json:"discovery_type"`
	Timeout         TimeoutType               `json:"timeout"`
	Name            string                    `json:"name"`
	Desc            string                    `json:"desc"`
	PassHost        string                    `json:"pass_host"`
	Scheme          string                    `json:"scheme"`
	Retries         uint                      `json:"retries"`
	RetryTimeout    uint                      `json:"retry_timeout"`
	Labels          map[string]string         `json:"labels,omitempty"`
	UpstreamHost    string                    `json:"upstream_host"`
	HashOn          string                    `json:"hash_on"`
	Key             string                    `json:"key"`
	KeepalivePool   UpstreamKeepAlivePoolType `json:"keepalive_pool"`
	TLSClientCertID string                    `json:"tls.client_cert_id"`
	Checks          UpstreamChecksType        `json:"checks"`
	Nodes           []UpstreamNodeType        `json:"nodes"`
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
	Active  UpstreamChecksActiveType  `json:"active"`
	Passive UpstreamChecksPassiveType `json:"passive"`
}

type UpstreamChecksActiveType struct {
	Type                   string                            `json:"type"`
	Timeout                uint                              `json:"timeout"`
	Concurrency            uint                              `json:"concurrency"`
	HTTPPath               string                            `json:"http_path"`
	Host                   string                            `json:"host"`
	Port                   uint                              `json:"port"`
	HTTPSVerifyCertificate bool                              `json:"https_verify_certificate"`
	ReqHeaders             []string                          `json:"req_headers"`
	Healthy                UpstreamChecksActiveHealthyType   `json:"healthy"`
	Unhealthy              UpstreamChecksActiveUnhealthyType `json:"unhealthy"`
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
	Healthy   UpstreamChecksPassiveHealthyType   `json:"healthy"`
	Unhealthy UpstreamChecksPassiveUnhealthyType `json:"unhealthy"`
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
	Host   string `tfsdk:"host"`
	Port   uint   `tfsdk:"port"`
	Weight uint   `tfsdk:"weight"`
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
