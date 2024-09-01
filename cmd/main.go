package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/condemo/simple-go-automake/templates"
)

// TODO: Ya que se está llegando a cierta complejidad,
// conviene empezar a separar la movidas
// por ejemplo: mover el string a un tmpl file y la funcionalidad
// al modulo templates/

type FileOps struct {
	BinName  string
	BinRoute string
	Arm      bool
	Test     bool
	Tailwind bool
	Templ    bool
	Air      bool
}

var fileStr string = `binary-name={{ .BinName }}

build:{{ if .Templ }} templ-build{{ end }}
{{"\t"}}@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe {{ .BinRoute }}
{{"\t"}}@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux {{ .BinRoute }}
{{"\t"}}@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin {{ .BinRoute }}

run: build
{{"\t"}}@./bin/${binary-name}-linux
{{if .Arm}}
arm-build:
{{"\t"}}@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 {{ .BinRoute }}

arm-run: arm-build
{{"\t"}}@./bin/${binary-name}-arm64{{end}}
{{ if .Test }}
test:
{{"\t"}}@go test {{ .BinRoute }}
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
templ-build:
{{"\t"}}@templ generate

templ-watch:
{{"\t"}}@templ generate --watch{{ end }}`

func main() {
	binName := flag.String("n", "default", "binary-name")
	binf := flag.String("b", "./cmd/main.go", "route to main go file")
	arm := flag.Bool("arm", false, "enable arm build")
	test := flag.Bool("t", false, "enable test")
	tailwind := flag.Bool("tail", false, "enable tailwind")
	tem := flag.Bool("templ", false, "enable templ")
	air := flag.Bool("air", false, "enable air")
	flag.Parse()

	data := FileOps{
		BinName:  *binName,
		BinRoute: *binf,
		Arm:      *arm,
		Test:     *test,
		Tailwind: *tailwind,
		Templ:    *tem,
		Air:      *air,
	}

	makeFile, err := os.Create("./Makefile")
	if err != nil {
		fmt.Println("error creating Makefile")
		os.Exit(1)
	}
	defer makeFile.Close()

	templ := template.New("maketext")
	templ.Parse(fileStr)
	templ.ExecuteTemplate(makeFile, "maketext", data)

	checkErr(err)

	if *tailwind {
		// TODO: Cambiar esto y hacer que el archivo se cree usando template
		// no hay necesidad de tirar el comando de tailwind
		cmd := exec.Command("tailwindcss", "init")
		cmd.Dir = "."
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		templates.MakeTailwindFile()
	}

	if *air {
		ad := templates.AirData{RootMain: *binf}
		templates.MakeAirFile(ad)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
