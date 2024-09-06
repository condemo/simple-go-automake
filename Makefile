binary-name=gomake

build:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name} ./cmd/main.go
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 ./cmd/main.go

run: build
	@./bin/${binary-name}-linux

install: build
	@sudo cp ./bin/gomake /bin/gomake

arm-install: build
	@sudo cp ./bin/gomake-arm64 /bin/gomake

clean:
	@rm -rf ./bin/*
	@go clean
