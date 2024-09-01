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
}

func CreateMakeFile(d FileOps) {
	makeFile, err := os.Create("./Makefile")
	if err != nil {
		fmt.Println("error creating Makefile")
		os.Exit(1)
	}
	defer makeFile.Close()

	templ := template.New("maketext")
	templ.Parse(fileStr)
	templ.ExecuteTemplate(makeFile, "maketext", d)
}

//go:embed air.tmpl
var s string

type AirData struct {
	RootMain string
}

func MakeAirFile(d AirData) {
	airPath := "./.air.toml"

	airFile, err := os.Create(airPath)
	checkErr(err)
	defer airFile.Close()

	airtempl := template.New("air")
	airtempl.Parse(s)
	airtempl.Execute(airFile, d)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed tailwind.tmpl
var t string

type TailwindData struct{}

func MakeTailwindFile() {
	// TODO: Modificar el template para cargar los datos dinámicos

	// TODO: Seguir con la implementación para automatizar tailwind
}
