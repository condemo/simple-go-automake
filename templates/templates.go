package templates

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
)

//go:embed make.tmpl
var fileStr string

type FileOps struct {
	BinName  string
	BinRoute string
	Arm      bool
	Test     bool
	Tailwind bool
	Templ    bool
	Air      bool
	GooseMig bool
}

func CreateMakeFile(d FileOps) {
	makeFile, err := os.Create("./Makefile")
	if err != nil {
		fmt.Println("error creating Makefile")
		os.Exit(1)
	}
	defer func() {
		if err := makeFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	templ := template.New("maketext")
	if _, err := templ.Parse(fileStr); err != nil {
		checkErr(err, "makefile templ parse failed")
	}
	if err := templ.ExecuteTemplate(makeFile, "maketext", d); err != nil {
		checkErr(err, "makefile templ execute failed")
	}
}

//go:embed air.tmpl
var s string

type AirData struct {
	RootMain string
}

func CreateAirFile(d AirData) {
	airPath := "./.air.toml"

	airFile, err := os.Create(airPath)
	checkErr(err, "airFile creation failed")
	defer func() {
		if err := airFile.Close(); err != nil {
			checkErr(err, "airFile closing failed")
		}
	}()

	airtempl := template.New("air")
	if _, err := airtempl.Parse(s); err != nil {
		checkErr(err, "air templ parse failed")
	}
	if err := airtempl.Execute(airFile, d); err != nil {
		checkErr(err, "air templ execute failed")
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//go:embed tailwind.tmpl
var t string

type TailwindData struct{}

func CreateTailwindFile(td TailwindData) {
	tailPath := "./tailwind.config.js"

	tailFile, err := os.Create(tailPath)
	checkErr(err, "tailwind file creation failed")
	defer func() {
		if err := tailFile.Close(); err != nil {
			checkErr(err, "tailFile closing failed")
		}
	}()

	tailwindTempl := template.New("tailwind")
	if _, err := tailwindTempl.Parse(t); err != nil {
		checkErr(err, "tailwind templ parse failed")
	}
	if err := tailwindTempl.Execute(tailFile, td); err != nil {
		checkErr(err, "tailwind templ execute failed")
	}
}

//go:embed mainFile.tmpl
var MainTempl string
