binary-name=gomake
build:
		@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name} ./cmd/main.go

run:build
		@./bin/${binary-name}-linux

clean:
		@rm -rf ./bin/*
		@go clean
