deps:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	# npm install @openapitools/openapi-generator-cli # <-- needs Java unfortunately

generate:
	oapi-codegen -generate server,spec -package model api/openapi.yaml > model/server.gen.go
	oapi-codegen -generate types -package model api/openapi.yaml > model/types.gen.go
	oapi-codegen -generate client -package model api/openapi.yaml > model/client.gen.go

generate-docs:
	openapi-generator generate -i api/article.yaml -g html2 -o docs/article/
	openapi-generator generate -i api/feed.yaml -g html2 -o docs/feed/

lint:
	golangci-lint run

test:
	go test ./...

run:
	go run .
