package api_client

func (client ApiClient) GetService(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/apisix/admin/services/"+id, nil)
}

func (client ApiClient) CreateService(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/apisix/admin/services/", &data)
}

func (client ApiClient) UpdateService(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/apisix/admin/services/"+id, &data)
}

func (client ApiClient) DeleteService(id string) (err error) {
	return client.Delete("/apisix/admin/services/" + id)
}
