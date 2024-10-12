# Variables
OUT_DIR ?= build
OUT_FILE := $(OUT_DIR)/libminequery.so
SRC_FILES := $(wildcard *.go)

# Build the dynamic library using Go
$(OUT_FILE): $(SRC_FILES)
	mkdir -p $(OUT_DIR)
	go build -o $(OUT_FILE) -buildmode=c-shared .

# Default target
all: $(OUT_FILE)

# Clean target to remove the built library
clean:
	rm -f $(OUT_FILE)

.PHONY: all clean

