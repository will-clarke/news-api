deps:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	# npm install @openapitools/openapi-generator-cli # <-- needs Java unfortunately

generate:
	oapi-codegen -generate server,spec -package article api/openapi.yaml > article/server.gen.go
	oapi-codegen -generate types -package article api/openapi.yaml > article/types.gen.go
	oapi-codegen -generate client -package article api/openapi.yaml > article/client.gen.go

generate-docs:
	openapi-generator generate -i api/openapi.yaml -g html2 -o docs/

lint:
	golangci-lint run

test:
	go test ./...
