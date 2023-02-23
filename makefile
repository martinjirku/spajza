BUILD_PATH = $(CURDIR)/build
API_PATH = $(CURDIR)/api
DB_PATH = $(CURDIR)/db
BINARY_NAME = zasobar-api
GO = $(shell which go)
LIQUIBASE = $(shell which liquibase)
DBMATE = $(shell which dbmate)
GOINSTALL = $(GO) install
GOCLEAN = $(GO) clean
GOBUILD = $(GO) build


build: deps
	@cd $(API_PATH); env GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME).darwin.arm64 ./cmd/server/main.go
	@cd $(API_PATH); env GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME).linux.arm64 ./cmd/server/main.go

deps:
	@cd $(API_PATH); $(GOGET)

clean:
	@rm -rf $(BUILD_PATH)/$(BINARY_NAME)

db-up:
	@cd $(DB_PATH); $(DBMATE) up

dev:
	@cd $(API_PATH); $(GO) run ./cmd/server/main.go

generate:
	@cd $(API_PATH); mockery -r --with-expecter --dir="./usecase" --name="UserGateway"