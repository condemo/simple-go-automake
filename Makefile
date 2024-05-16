binary-name=gomake

build:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name} ./cmd/main.go

run: build
	@./bin/${binary-name}-linux

install:
	@sudo cp ./bin/gomake /bin/gomake

clean:
	@rm -rf ./bin/*
	@go clean
