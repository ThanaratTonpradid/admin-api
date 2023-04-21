BINARY_FILE = mini-api

clean:
	@rm -rf $(BINARY_FILE) coverage.out

build:
	@go build
	@ls -alh $(BINARY_FILE)

cover:
	@go test ./... -cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out

@docs:
	@swag init --instanceName $(NAME) --ot go -d ./internal/$(NAME),./lib -g ./route/swagger.route.go -o ./internal/$(NAME)/doc
	@swag fmt -g ./internal/$(NAME)/route/swagger.route.go

models:
	go run cmd/generate/generate.go

infra-up:
	docker-compose -p mini -f .development/docker-compose.yml up -d

infra-down:
	docker-compose -p mini -f .development/docker-compose.yml down

docs:
	@make @docs NAME=API

api: docs
	@go run main.go api || true
