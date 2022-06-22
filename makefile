BUILD_PATH = $(CURDIR)/build
API_PATH = $(CURDIR)/api
BINARY_NAME = zasobar-api
GO = $(shell which go)
GOINSTALL = $(GO) install
GOCLEAN = $(GO) clean
GOBUILD = $(GO) build



build: deps
	@cd $(API_PATH); $(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) ./server.go

deps:
	@cd $(API_PATH); $(GOGET)

clean:
	@rm -rf $(BUILD_PATH)/$(BINARY_NAME)

