GO    := go
GODEP := godep $(GO)
pkgs   = $(shell $(GO) list ./... | grep -v /vendor/)

ifdef DEBUG
	bindata_flags = -debug
endif

all: assets test install

assets:
	@echo ">> writing assets"
	@$(GO) get -u github.com/jteeuwen/go-bindata/...
	@go-bindata $(bindata_flags) -pkg ui -o ui/bindata.go ui/assets/...

test:
	@echo ">> running tests"
	@$(GODEP) test $(pkgs)

install:
	@echo ">> installing packages"
	@$(GODEP) install $(pkgs)

.PHONY: all assets test install
