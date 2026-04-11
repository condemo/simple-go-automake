package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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

	createFiles(*mainPath)
}

// TODO:
func createFiles(mainPath string) {
	dir, _ := filepath.Split(mainPath)
	_, err := exec.Command("mkdir", dir).Output()
	if err != nil {
		log.Fatalf("mkdir failed: %s", err)
	}
	fmt.Println("main folder created")

	f, err := os.Create(mainPath)
	if err != nil {
		log.Fatalf("main.go creation failed: %s", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("main.go closing failed: %s", err)
		}
	}()

	if _, err := f.WriteString(templates.MainTempl); err != nil {
		log.Fatalf("write to main.go failed: %s", err)
	}
	fmt.Println("main.go created")
}
