codegen:
	go run github.com/ogen-go/ogen/cmd/ogen --target ./internal/generated/scheme/ -package scheme --clean ./api/openapi.json
	protoc -I proto proto/clients/dialog.proto --go_out=plugins=grpc:proto/