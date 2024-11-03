BINARY_NAME=server.exe
OUTPUT_DIR=build

local: build
	@${OUTPUT_DIR}/${BINARY_NAME} -config=local

prod: build
	@${OUTPUT_DIR}/${BINARY_NAME} -config=prod

build:
	@go build -o ${OUTPUT_DIR}/${BINARY_NAME} cmd/main.go