package main

import (
	"fmt"
	"os"
	"strings"
)

var binName string

func createStr() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("binary-name=%s\n", binName))

	b.WriteString("build:\n")
	b.WriteString("\t\t@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/main.go\n")
	b.WriteString("\t\t@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-windows.exe ./cmd/main.go\n")
	b.WriteString("\t\t@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/main.go\n")

	b.WriteString("\n")

	b.WriteString("run:build\n")
	b.WriteString("\t\t@./bin/${binary-name}-linux\n")

	b.WriteString("\n")

	b.WriteString("clean:\n")
	b.WriteString("\t\t@rm -rf ./bin/*\n\t\t@go clean")

	return b.String()
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("error: bad args: <binName> <folder>")
		os.Exit(1)
	}

	binName = os.Args[1]
	folder := os.Args[2]

	makeFile, err := os.Create(folder + "/Makefile")
	if err != nil {
		fmt.Println("error creating Makefile")
		os.Exit(1)
	}
	defer makeFile.Close()

	str := createStr()
	_, err = makeFile.WriteString(str)
	if err != nil {
		fmt.Println("error: writing fail")
		os.Exit(1)
	}
}
