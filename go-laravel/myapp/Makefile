BINARY_NAME=myapp.exe

## build: builds all binaries
build:
	@go build -o tmp/${BINARY_NAME} .
	@echo Celeritas built!

run: build
	@echo Staring Celeritas...
	@.\tmp\${BINARY_NAME} &
	@echo Celeritas started!

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
	
stop:
	@echo "Starting the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Celeritas

restart: stop start