# Makefile

# Variables
GO ?= go

# Default target to run when `make` is called without arguments
.PHONY: all
all: test

# Run all tests with verbose output
.PHONY: test_unit
test_unit:
	$(GO) test -v ./internal/tests/unit/...