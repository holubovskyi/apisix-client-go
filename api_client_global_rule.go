package api_client

func (client ApiClient) GetGlobalRule(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/apisix/admin/global_rules/"+id, nil)
}

func (client ApiClient) CreateGlobalRule(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/apisix/admin/global_rules/"+id, &data)
}

func (client ApiClient) UpdateGlobalRule(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/apisix/admin/global_rules/"+id, &data)
}

func (client ApiClient) DeleteGlobalRule(id string) (err error) {
	return client.Delete("/apisix/admin/global_rules/" + id)
}
