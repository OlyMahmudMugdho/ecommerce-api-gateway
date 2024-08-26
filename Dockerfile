FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go  build -o bin/ecommerce-api-gateway cmd/main.go
CMD [ "./bin/ecommerce-api-gateway" ]