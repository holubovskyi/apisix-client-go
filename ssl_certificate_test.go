package api_client

import (
	"encoding/json"
	"testing"
)

func TestSSLCertificateClientMarshaling(t *testing.T) {
	ca := "test-ca-pem"
	depth := int64(3)
	cert := SSLCertificate{
		Client: &SSLClient{CA: &ca, Depth: &depth},
	}

	body, err := json.Marshal(cert)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var result map[string]map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	client, ok := result["client"]
	if !ok {
		t.Fatalf("missing client field")
	}
	if client["ca"] != ca {
		t.Errorf("expected ca %q, got %q", ca, client["ca"])
	}
	if client["depth"] != float64(depth) {
		t.Errorf("expected depth %d, got %v", depth, client["depth"])
	}
}

func TestSSLCertificateClientOmittedWhenNil(t *testing.T) {
	cert := SSLCertificate{}

	body, err := json.Marshal(cert)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if _, ok := result["client"]; ok {
		t.Errorf("expected client field to be omitted, but it was present")
	}
}
