BINARY_NAME=./dist/SumahoNikki
PROJECT_MAIN=./cmd/sumaho_nikki/main.go

build:
	@echo "Building the go app and putting it in ${BINARY_NAME}"
	@go mod tidy
	@go build -o ${BINARY_NAME} ${PROJECT_MAIN}
	@echo "Built the go app"

start:
	@go mod tidy
	@go build -o ${BINARY_NAME} ${PROJECT_MAIN}
	@./dist/SumahoNikki

clean:
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned dist! :)"
