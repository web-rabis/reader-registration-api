swagger:
	swag init --parseInternal --parseVendor --parseDependency --parseDepth 3 -g ./cmd/apiserver/main.go -o ./api
