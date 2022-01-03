deps:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen

generate:
	oapi-codegen -generate server,spec -package api pets.yaml > api/petstore-server.gen.go
	oapi-codegen -generate types -package api pets.yaml > api/petstore-types.gen.go
	oapi-codegen -generate client -package api pets.yaml > api/petstore-client.gen.go
