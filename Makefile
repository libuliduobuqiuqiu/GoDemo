BINARY_NAME = godemo
EXCUTE_PATH = $(shell pwd)
VERSION = $(shell git describe --tags)
BUILD_TIME = $(shell date -u "+%Y-%m-%d %H:%M:%S UTC")

define build
	@echo "Building $(BINARY_NAME)..."
	$(GO_BUILD_ENV) go build -ldflags="-X 'godemo/cmd.version=$(VERSION)' -X 'godemo/cmd.buildTime=$(BUILD_TIME)'" -o godemo main.go
	@echo "Build completed! Excutable path is $(EXCUTE_PATH)/$(BINARY_NAME)"
endef

wsl: GO_BUILD_ENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
wsl:
	$(call build)

mac: GO_BUILD_ENV = CGO_ENABLED=0 GOOS=darwin GOARCH=arm64
mac:
	$(call build)
