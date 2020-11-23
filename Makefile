VERSION := $(shell git describe --tags --abbrev=0)
GIT_SHA := $(shell git rev-parse --short HEAD)
DATE := $(shell date +%s)
MODULE := urlshortener
GO = GO111MODULE=on CGO_ENABLED=0 go
GO_FLAGS := -mod=vendor -tags production -installsuffix cgo


LDFLAGS += -X $(MODULE)/pkg.Timestamp=$(DATE)
LDFLAGS += -X $(MODULE)/pkg.Version=$(VERSION)
LDFLAGS += -X $(MODULE)/pkg.GitSHA=$(GIT_SHA)

CMD_NAMES := $(foreach pb, $(wildcard ./cmd/*), $(pb)_cmd)

build: 
	$(GO) build $(GO_FLAGS) -o bin/urlshortener -ldflags "$(LDFLAGS)" ./cmd/service

# test::
# 	GO111MODULE=on go test -mod vendor -v ./cmd/clients

clean:
	@rm bin/*

.PHONY: gen tools
gen:
	wire ./cmd/service
	mockgen -source=./internal/backend.go -destination=./mock/service_backend.go -package=mock

tools:
	GO111MODULE=off go get -u github.com/google/wire/...
	# GO111MODULE=off go get -u github.com/golang/mock/mockgen