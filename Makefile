BINARY=futil
INSTALL_DIR=/usr/local/bin
GIT_TAG=$(shell git describe --tags)

.PHONY: all build install uninstall

all: install

build:
	@echo "Building $(BINARY)..."
	go build -ldflags "-X 'github.com/Largeb0525/futil/cmd.version=$(GIT_TAG)'" -o $(BINARY)

install: build
	@echo "Installing $(BINARY) to $(INSTALL_DIR)..."
	sudo mv $(BINARY) $(INSTALL_DIR)

uninstall:
	@echo "Removing $(INSTALL_DIR)/$(BINARY)..."
	sudo rm -f $(INSTALL_DIR)/$(BINARY)
