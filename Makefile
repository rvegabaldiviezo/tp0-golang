# Variables
CLIENT_MAIN=client/main.go
SERVER_MAIN=server/main.go
CLIENT_BIN=client_app
SERVER_BIN=server_app

# Build individuales
build-client:
	go build -o $(CLIENT_BIN) $(CLIENT_MAIN)

build-server:
	go build -o $(SERVER_BIN) $(SERVER_MAIN)

# Ejecutar
run-client: build-client
	./$(CLIENT_BIN)

run-server: build-server
	./$(SERVER_BIN)

# Limpiar binarios y logs
clean:
	rm -f $(CLIENT_BIN) $(SERVER_BIN)

clean-logs:
	find . -type f -name "*.log" -print -delete

clean-all: clean clean-logs

.PHONY: build-client build-server run-client run-server clean clean-logs clean-all
