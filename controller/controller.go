package controller

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Question    string
	NbQuestion  int
	Finquizz    bool
	Score       int
	Debutquizz  bool
	Finquestion bool
	BonneRep    string
	Reponses    []string
}

var question []string = []string{
	"Quelle est la capitale de la France ?",
	"Combien de continents y a-t-il sur Terre ?",
	"Quel est l'élément chimique dont le symbole est 'O' ?",
	"Lesquels de ces aliments n'est pas un fruit ?",
	"Est-ce que le cannibalisme est légal en France ?",

	"Les claquettes chaussettes, sont-elles socialement acceptables ?",
	"Quelle est la couleur du cheval blanc d'Henri IV ?",
	"Combien de côté a un triangle ?",
	"As-tu faim ?",
	"Aurais-je une bonne note ?",
}
var reponses [][]string = [][]string{
	{"Paris", "Londres", "Berlin", "Madrid"},
	{"5", "6", "7", "8"},
	{"Oxygène", "Or", "Osmium", "Oganesson"},
	{"Pomme", "Carotte", "Banane", "Orange"},
	{"Oui", "Non", "Seulement le dimanche"},
	{"Oui j'adore", "Non c'est pas fou", "Seulement avec les chaussettes à formes"},
	{"Blanc", "Noir", "Marron", "Gris"},
	{"3", "4", "5", "6"},
	{"Pas trop", "Pas du tout", "je mangerais bien mon père"},
	{"20/20", "Eclaté au sol", "Passable"},
}

var bonnesReponses []string = []string{
	"Paris",
	"7",
	"Oxygène",
	"Carotte",
	"Non",
	"Non c'est pas fou",
	"Blanc",
	"3",
	"Pas trop",
	"20/20",
}

var data = PageData{
	Question:    "Bienvenue 🎉Convertis avec assurance avec The convertisseur d'unités !",
	Finquizz:    false,
	Score:       0,
	Debutquizz:  false,
	Finquestion: false,
	BonneRep:    "",
	Reponses:    []string{},
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("recommencer") == "oui" {
		data = PageData{
			Question:    "Bienvenue 🎉Un quizz super crazy t'attend !",
			Finquizz:    false,
			Score:       0,
			Debutquizz:  false,
			Finquestion: false,
			BonneRep:    "",
			Reponses:    []string{},
		}
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("commencer") == "oui" {
		data.Debutquizz = true
		data.Question = question[0]
		data.Reponses = reponses[0]
		data.BonneRep = bonnesReponses[0]
		data.NbQuestion = 1
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("réponse") != "" {
		if data.BonneRep == r.FormValue("réponse") {
			data.Score++
		}
		data.Finquestion = true
		if data.NbQuestion == 9 {
			data.Finquizz = true
			data.Question = "Quiz Terminé !"
		} else {
			data.NbQuestion++
		}
	}
	if r.FormValue("suivant") == "oui" {
		data.Finquestion = false
		data.BonneRep = bonnesReponses[data.NbQuestion]
		data.Question = question[data.NbQuestion]
		data.Reponses = reponses[data.NbQuestion]
		data.BonneRep = bonnesReponses[data.NbQuestion]
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
