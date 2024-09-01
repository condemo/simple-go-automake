package main

import (
	"flag"

	"github.com/condemo/simple-go-automake/templates"
)

// TODO: Ya que se está llegando a cierta complejidad,
// conviene empezar a separar la movidas
// por ejemplo: mover el string a un tmpl file y la funcionalidad
// al modulo templates/

func main() {
	binName := flag.String("n", "default", "binary-name")
	binf := flag.String("b", "./cmd/main.go", "route to main go file")
	arm := flag.Bool("arm", false, "enable arm build")
	test := flag.Bool("t", false, "enable test")
	tailwind := flag.Bool("tail", false, "enable tailwind")
	tem := flag.Bool("templ", false, "enable templ")
	air := flag.Bool("air", false, "enable air")
	flag.Parse()

	data := templates.FileOps{
		BinName:  *binName,
		BinRoute: *binf,
		Arm:      *arm,
		Test:     *test,
		Tailwind: *tailwind,
		Templ:    *tem,
		Air:      *air,
	}
	templates.CreateMakeFile(data)

	if *tailwind {
		// TODO: Añadir flags para configurar tailwind, daisyui ...
		td := templates.TailwindData{}
		templates.CreateTailwindFile(td)
	}

	if *air {
		ad := templates.AirData{RootMain: *binf}
		templates.CreateAirFile(ad)
	}
}
