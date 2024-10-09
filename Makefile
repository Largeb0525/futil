BINARY=futil
INSTALL_DIR=/usr/local/bin

.PHONY: all build install uninstall

all: install

build:
	@echo "Building $(BINARY)..."
	go build -o $(BINARY)

install: build
	@echo "Installing $(BINARY) to $(INSTALL_DIR)..."
	sudo mv $(BINARY) $(INSTALL_DIR)

uninstall:
	@echo "Removing $(INSTALL_DIR)/$(BINARY)..."
	sudo rm -f $(INSTALL_DIR)/$(BINARY)
