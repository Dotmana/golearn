package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name, Address, City, Zip string
	Region                   Region
}

type Region struct {
	Region string
}

type Hotels []Hotel

var tpl *template.Template

func main() {
	tpl := template.Must(template.ParseFiles("test.gohtml"))
	south := Region{
		"Southern",
	}
	north := Region{
		"Northern",
	}

	h := Hotels{
		Hotel{
			Name:    "Hotel California",
			Address: "Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "77074",
			Region:  south,
		},
		Hotel{
			Name:    "Waldorf Astoria",
			Address: "Park Lane",
			City:    "San Diego",
			Zip:     "72312",
			Region:  north,
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}
