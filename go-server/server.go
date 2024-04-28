package main

import (
	"html/template"
	"log"
	"net/http"
)

type Temps struct{
	index *template.Template
	summary *template.Template
	notemp *template.Template
}

func setupTemp() *Temps{
	temps := new(Temps)

	temps.notemp,_ = template.ParseFiles("templates/failedTemplate.html")
	
	index, err := template.ParseFiles("templates/index.html")
	if err != nil{
		index = temps.notemp
	}
	temps.index = index

  summary, err := template.ParseFiles("templates/summary.html")
	if err != nil{
		summary= temps.notemp
	}
	temps.summary = summary
	
	return temps
}

func handleIndex(w http.ResponseWriter, tmp *template.Template) {
	err := tmp.Execute(w,nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func handleSummary(w http.ResponseWriter, tmp *template.Template) {
	err := tmp.Execute(w,nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func main(){
	temps := setupTemp()
	http.HandleFunc("/",func(w http.ResponseWriter, rq *http.Request) {
		handleSummary(w,temps.index)
	})
	http.HandleFunc("/summary",func(w http.ResponseWriter, rq *http.Request) {
		handleSummary(w,temps.summary)
	})

	http.ListenAndServe("",nil)
}