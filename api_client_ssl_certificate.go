package api_client

func (client ApiClient) GetSslCertificate(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/ssls/"+id, nil)
}

func (client ApiClient) CreateSslCertificate(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/ssls/", &data)
}

func (client ApiClient) UpdateSslCertificate(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/ssls/"+id, &data)
}

func (client ApiClient) DeleteSslCertificate(id string) (err error) {
	return client.Delete("/ssls/" + id)
}
