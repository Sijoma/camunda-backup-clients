generate: 
	openapi-generator generate -i api/openapi.yaml -g go --config config.yml
