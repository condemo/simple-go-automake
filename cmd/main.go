package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type FileOps struct {
	BinName  string
	Arm      bool
	Test     bool
	Tailwind bool
	Templ    bool
}

var fileStr string = `binary-name={{ .BinName }}

build:{{ if .Templ }} templ-gen{{ end }}
{{"\t"}}@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe ./cmd/main.go
{{"\t"}}@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/main.go
{{"\t"}}@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/main.go

run: build
{{"\t"}}@./bin/${binary-name}-linux
{{if .Arm}}
arm-build:
{{"\t"}}@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 ./cmd/main.go

arm-run: arm-build
{{"\t"}}@./bin/${binary-name}-arm64{{end}}
{{ if .Test }}
test:
{{"\t"}}@go test cmd/main.go
{{ end }}
clean:
{{"\t"}}@rm -rf ./bin/*
{{"\t"}}@go clean
{{ if .Tailwind }}
css-build:
{{"\t"}}@tailwindcss -i ./static/css/input.css -o ./static/css/style.css

css-watch:
{{"\t"}}@tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch{{ end }}
{{ if .Templ }}
templ-gen:
{{"\t"}}@templ generate

templ-watch:
{{"\t"}}@templ generate --watch{{ end }}`

func main() {
	binName := flag.String("n", "default", "binary-name")
	folder := flag.String("d", ".", "directory name")
	arm := flag.Bool("arm", false, "enable arm build")
	test := flag.Bool("t", false, "enable test")
	tailwind := flag.Bool("tail", false, "enable tailwind")
	tem := flag.Bool("templ", false, "enable templ")
	flag.Parse()

	data := FileOps{
		BinName:  *binName,
		Arm:      *arm,
		Test:     *test,
		Tailwind: *tailwind,
		Templ:    *tem,
	}

	makeFile, err := os.Create(*folder + "/Makefile")
	if err != nil {
		fmt.Println("error creating Makefile")
		os.Exit(1)
	}
	defer makeFile.Close()

	templ := template.New("maketext")
	templ.Parse(fileStr)
	templ.ExecuteTemplate(makeFile, "maketext", data)
}
