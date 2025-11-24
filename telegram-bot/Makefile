APP_NAME=telegram-bot
BIN_DIR=bin

.PHONY: run build tidy clean

run:
	go run ./...

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
