APPVERSION=$(or $(shell git rev-parse HEAD 2>/dev/null),"unknown")
AWSBIN = bin/aws

GO = go
V  = 0
Q  = $(if $(filter 1,$V),,@)
M  = $(shell printf "\033[34;1m▶\033[0m")
OS = $(shell uname -s | tr A-Z a-z)

.PHONY: all
all:

all: build

build: aws

BUILDFLAGS = GOFLAGS=-mod=vendor
aws:
	$Q $(BUILDFLAGS) $(GO) build \
		-ldflags '-X main.BuildVersion=$(APPVERSION)' \
		-o $(AWSBIN) ./cmd/aws

# Misc
.PHONY: clean
clean: ; $(info $(M) cleaning...) @ ## Clean up everything
	$Q rm -f $(AWSBIN)
	$Q rm -f bin/example
