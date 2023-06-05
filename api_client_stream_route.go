package api_client

func (client ApiClient) GetStreamRoute(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/apisix/admin/stream_routes/"+id, nil)
}

func (client ApiClient) CreateStreamRoute(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/apisix/admin/stream_routes/", &data)
}

func (client ApiClient) UpdateStreamRoute(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/apisix/admin/stream_routes/"+id, &data)
}

func (client ApiClient) DeleteStreamRoute(id string) (err error) {
	return client.Delete("/apisix/admin/stream_routes/" + id)
}
