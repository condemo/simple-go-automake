package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

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

type AirData struct {
	RootMain string
}

type TailwindData struct{}

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

	// homeDir, err := os.UserHomeDir()
	checkErr(err)
	// shareDir := homeDir + "/.local/share/gomake"

	if *tailwind {
		cmd := exec.Command("tailwindcss", "init")
		cmd.Dir = "."
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	if *air {
		// cmd := exec.Command("air", "init")
		// cmd.Dir = *folder
		// if err := cmd.Run(); err != nil {
		// 	log.Fatal(err)
		// }

		// airTempPath := shareDir + "/air.tmpl"
		airPath := "./.air.toml"
		//
		// if _, err := os.Stat(airTempPath); err != nil {
		// 	err = os.Mkdir(shareDir, os.ModePerm)
		// 	checkErr(err)
		// 	f, err := os.Create(airTempPath)
		// 	checkErr(err)
		// 	fmt.Println("creando template")
		// 	f.Close()
		// }

		airFile, err := os.Create(airPath)
		checkErr(err)
		defer airFile.Close()

		airtempl, err := template.ParseFiles("templates/air.tmpl")
		checkErr(err)

		airtempl.Execute(airFile, AirData{RootMain: *binf})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
