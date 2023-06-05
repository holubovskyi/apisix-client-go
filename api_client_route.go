package api_client

func (client ApiClient) GetRoute(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/apisix/admin/routes/"+id, nil)
}

func (client ApiClient) CreateRoute(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/apisix/admin/routes/", &data)
}

func (client ApiClient) UpdateRoute(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/apisix/admin/routes/"+id, &data)
}

func (client ApiClient) DeleteRoute(id string) (err error) {
	return client.Delete("/apisix/admin/routes/" + id)
}
