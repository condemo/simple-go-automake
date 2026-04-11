package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/condemo/simple-go-automake/templates"
)

// TODO: Ya que se está llegando a cierta complejidad,
// conviene empezar a separar la movidas
// por ejemplo: mover el string a un tmpl file y la funcionalidad
// al modulo templates/

func main() {
	binName := flag.String("n", "default", "binary-name")
	mainPath := flag.String("b", "./cmd/main.go", "route to main go file")
	arm := flag.Bool("arm", false, "enable arm build")
	test := flag.Bool("t", false, "enable test")
	tailwind := flag.Bool("tail", false, "enable tailwind")
	tem := flag.Bool("templ", false, "enable templ")
	air := flag.Bool("air", false, "enable air")
	gooseMig := flag.Bool("goose", false, "enable goose migrations")
	flag.Parse()

	data := templates.FileOps{
		BinName:  *binName,
		BinRoute: *mainPath,
		Arm:      *arm,
		Test:     *test,
		Tailwind: *tailwind,
		Templ:    *tem,
		Air:      *air,
		GooseMig: *gooseMig,
	}
	templates.CreateMakeFile(data)

	if *tailwind {
		// TODO: Añadir flags para configurar tailwind, daisyui ...
		td := templates.TailwindData{}
		templates.CreateTailwindFile(td)
	}

	if *air {
		ad := templates.AirData{RootMain: *mainPath}
		templates.CreateAirFile(ad)
	}

	if err := createFiles(*mainPath); err != nil {
		log.Fatalf("error creating main.go file and folder: %s", err)
	}
}

// TODO:
func createFiles(mainPath string) error {
	dir, _ := filepath.Split(mainPath)
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}
	fmt.Println("main folder created")

	f, err := os.Create(mainPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("main.go closing failed: %s", err)
		}
	}()

	if _, err := f.WriteString(templates.MainTempl); err != nil {
		return err
	}
	fmt.Println("main.go created")
	return nil
}
