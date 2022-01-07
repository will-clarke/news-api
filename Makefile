deps:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	# npm install @openapitools/openapi-generator-cli # <-- needs Java unfortunately

generate:
	oapi-codegen -generate server,spec -package article api/article.yaml > article/server.gen.go
	oapi-codegen -generate types -package article api/article.yaml > article/types.gen.go
	oapi-codegen -generate client -package article api/article.yaml > article/client.gen.go

	oapi-codegen -generate server,spec -package feed api/feed.yaml > feed/server.gen.go
	oapi-codegen -generate types -package feed api/feed.yaml > feed/types.gen.go
	oapi-codegen -generate client -package feed api/feed.yaml > feed/client.gen.go

generate-docs:
	openapi-generator generate -i api/article.yaml -g html2 -o docs/article/
	openapi-generator generate -i api/feed.yaml -g html2 -o docs/feed/

lint:
	golangci-lint run

test:
	go test ./...

run:
	go run .
