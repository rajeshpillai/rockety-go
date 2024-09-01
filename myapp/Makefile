BINARY_NAME=RocketyocketyApp
build:
	@go mod vendor
	@echo "Building Rockety..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Rockety built!" 

run: build
	@echo "Starting Rockety..."
	@./tmp/${BINARY_NAME} &
	@echo "Rockety started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME} 
	@echo "Cleaned!" 

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!" 

start: run 

stop:
	@echo "Stopping Rockety..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Rockety" 

restart: stop start 

