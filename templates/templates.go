package templates

import (
	_ "embed"
	"html/template"
	"log"
	"os"
)

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
