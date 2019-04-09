SHELL=/bin/bash -o pipefail

.PHONY: sdk
sdk:
		curl https://raw.githubusercontent.com/ory/hydra/master/docs/api.swagger.json > ./api.swagger.json
		rm -rf ./swagger/*
		java -jar ./swagger-codegen-cli-2.2.3.jar generate -i ./api.swagger.json -l go -o ./swagger

		git checkout HEAD -- swagger/configuration.go
		git checkout HEAD -- swagger/api_client.go
