package api_client

func (client ApiClient) GetConsumer(username string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/apisix/admin/consumers/"+username, nil)
}

func (client ApiClient) CreateConsumer(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/apisix/admin/consumers/", &data)
}

func (client ApiClient) UpdateConsumer(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/apisix/admin/consumers/", &data)
}

func (client ApiClient) DeleteConsumer(username string) (err error) {
	return client.Delete("/apisix/admin/consumers/" + username)
}
