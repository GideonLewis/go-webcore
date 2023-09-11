build:
	go build -ldflags "-s -w" -o ./tmp/server ./cmd/main.go

docker-dev:
	@docker-compose down
	@docker-compose up -d

run:
	@go get github.com/cosmtrek/air@latest
	@air -c .air.toml

swagger:
	@hash swag 2>/dev/null || GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag@latest
	@swag init -g cmd/main.go --parseDependency --parseInternal --parseDepth 2

