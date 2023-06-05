package api_client

func (client ApiClient) GetPluginMetadataLogFormat(pluginName string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/apisix/admin/plugin_metadata/"+pluginName, nil)
}

func (client ApiClient) CreatePluginMetadataLogFormat(pluginName string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/apisix/admin/plugin_metadata/"+pluginName, &data)
}

func (client ApiClient) UpdatePluginMetadataLogFormat(pluginName string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/apisix/admin/plugin_metadata/"+pluginName, &data)
}

func (client ApiClient) DeletePluginMetadataLogFormat(pluginName string) (err error) {
	return client.Delete("/apisix/admin/plugin_metadata/" + pluginName)
}
