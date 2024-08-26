run:
	@go run cmd/main.go
build:
	@ GOOS=linux GOARCH=amd64 go  build -o bin/ecommerce-api-gateway cmd/main.go
docker-build:
	@docker build -t olymahmudmugdho/ecommerce-api-gateway .