generate: 
	openapi-generator generate -i openapi-spec.yml -g go -o go --config config.yml