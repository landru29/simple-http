GO := go
BUILD := go build
DOCKER := docker
IMAGE_NAME := simple-http
EXEC := simple-http
SOURCES := $(shell find . -name '*.go' -not -path 'vendor')

$(EXEC): $(SOURCES)
	$(BUILD)

package: $(EXEC)
	$(DOCKER) build -t $(IMAGE_NAME) .