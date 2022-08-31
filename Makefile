VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git rev-parse --short HEAD)

GOBIN = $(shell go env GOPATH)/bin

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=somechain \
	-X github.com/cosmos/cosmos-sdk/version.AppName=somechain \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

export GO111MODULE=on

###############################################################################
###                                  Build                                  ###
###############################################################################
.PHONY: build

install:
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/somechaind		

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

###############################################################################
###                                  Release                                ###
###############################################################################

release:
	@echo "dsd"