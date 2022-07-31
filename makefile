BUILD_PATH = $(CURDIR)/build
API_PATH = $(CURDIR)/api
DB_PATH = $(CURDIR)/db
BINARY_NAME = zasobar-api
GO = $(shell which go)
LIQUIBASE = $(shell which liquibase)
GOINSTALL = $(GO) install
GOCLEAN = $(GO) clean
GOBUILD = $(GO) build


build: deps
	@cd $(API_PATH); $(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) ./cmd/main.go

deps:
	@cd $(API_PATH); $(GOGET)

clean:
	@rm -rf $(BUILD_PATH)/$(BINARY_NAME)

db-update:
	@cd $(DB_PATH); $(LIQUIBASE) --changelog-file=./000_changelog.xml update

dev:
	cd $(API_PATH); $(GO) run ./cmd/main.go