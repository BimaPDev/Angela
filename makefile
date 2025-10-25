# Makefile (put at repo root)
SHELL := /bin/bash

# --- Config ---
IP ?= $(shell ipconfig getifaddr en0)
API_PORT ?= 8080
API_BASE := http://$(IP):$(API_PORT)

CLIENT_DIR := client
SERVER_DIR := server

# Adjust if your venv path differs:
PYTHON_BIN ?= $(SERVER_DIR)/python/.venv/bin/python

# --- Phony targets ---
.PHONY: help dev client/start server/start client/install server/env stop

help:
	@echo "make dev              # start Go API and Expo (both, with cleanup on Ctrl-C)"
	@echo "make server/start     # start Go API only"
	@echo "make client/start     # start Expo only"
	@echo "make client/install   # install client deps"
	@echo "make server/env       # echo server env (debug)"
	@echo "IP=$(IP)  API=$(API_BASE)"

# Start both processes and wire Ctrl-C to kill both.
dev:
	@echo "API: $(API_BASE)"
	@echo "Starting server and client..."
	@set -euo pipefail; \
	trap 'echo; echo Stopping...; kill 0' INT TERM EXIT; \
	( cd $(SERVER_DIR) && \
	  PYTHON_BIN="$(abspath $(PYTHON_BIN))" \
	  go run ./cmd/api ) & \
	( cd $(CLIENT_DIR) && \
	  EXPO_PUBLIC_API_BASE="$(API_BASE)" \
	  npx expo start ) & \
	wait

# Start Expo app only (uses detected IP)
client/start:
	cd $(CLIENT_DIR) && EXPO_PUBLIC_API_BASE="$(API_BASE)" npx expo start

# Start Go API only (uses your venv Python)
server/start:
	cd $(SERVER_DIR) && PYTHON_BIN="$(abspath $(PYTHON_BIN))" go run ./cmd/api

# One-time installs for the client
client/install:
	cd $(CLIENT_DIR) && \
	npm i && \
	npx expo install expo-camera && \
	npx expo install expo-file-system && \
	npm i @react-native-async-storage/async-storage

# Show the resolved env used by the server
server/env:
	@echo "PYTHON_BIN=$(abspath $(PYTHON_BIN))"
	@echo "API will bind on :$(API_PORT)"
	@which go || true
