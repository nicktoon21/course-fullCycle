package main

import (
	"html/template"
	"os"
)
type Curso struct {
	Nome string
	CargaHoraria int
}

func main(){
	curso := Curso{"Go", 40}
	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n"))
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	} 
}