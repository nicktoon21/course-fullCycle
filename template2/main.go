package main

import (
	"html/template"
	"net/http"
	"strings"
)
type Curso struct {
	Nome string
	CargaHoraria int
}

type Cursos []Curso

func toUpper(texto string) string {
	return strings.ToUpper(texto)
}


func main(){
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"toUpper": toUpper})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 30},
			{"Python", 20},
		})
		
		if err != nil {
			panic(err)
		}

	})

	http.ListenAndServe(":8080", nil)
}