package controller

import (
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	Calcul    string
	Message   string
	Unité1    string
	Unité2    string
	Valeur1   string
	Resultat  string
	Direction string
}

var data = PageData{
	Calcul:    "Convertis avec assurance avec The convertisseur d'unités !",
	Message:   "Bienvenue 🎉",
	Unité1:    "Choisis une unité",
	Unité2:    "Choisis une unité",
	Valeur1:   "",
	Resultat:  "",
	Direction: "-->",
}

func Calcul(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("opération") != "" {
		switch r.FormValue("opération") {
		case "km":
			data.Unité1 = "Kilomètres"
			data.Unité2 = "Miles"
			data.Valeur1 = ""
			data.Resultat = ""
		case "kl":
			data.Unité1 = "Kilogrammes"
			data.Unité2 = "Livres"
			data.Valeur1 = ""
			data.Resultat = ""
		case "cf":
			data.Unité1 = "Celsius"
			data.Unité2 = "Fahrenheit"
			data.Valeur1 = ""
			data.Resultat = ""
		}
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("direction") != "" {
		switch data.Unité1 {
		case "Kilomètres":
			data.Unité1 = "Miles"
			data.Unité2 = "Kilomètres"
			data.Valeur1 = ""
			data.Resultat = ""
		case "Miles":
			data.Unité1 = "Kilomètres"
			data.Unité2 = "Miles"
			data.Valeur1 = ""
			data.Resultat = ""
		case "Kilogrammes":
			data.Unité1 = "Livres"
			data.Unité2 = "Kilogrammes"
			data.Valeur1 = ""
			data.Resultat = ""
		case "Livres":
			data.Unité1 = "Kilogrammes"
			data.Unité2 = "Livres"
			data.Valeur1 = ""
			data.Resultat = ""
		case "Celsius":
			data.Unité1 = "Fahrenheit"
			data.Unité2 = "Celsius"
			data.Valeur1 = ""
			data.Resultat = ""
		case "Fahrenheit":
			data.Unité1 = "Celsius"
			data.Unité2 = "Fahrenheit"
			data.Valeur1 = ""
			data.Resultat = ""
		}
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("nombre1") != "" {
		r, err := strconv.ParseFloat(r.FormValue("nombre1"), 64)
		if err != nil {
			r = 0
		}
		switch data.Unité1 {
		case "Kilomètres":
			data.Valeur1 = strconv.FormatFloat(r, 'f', 6, 64)
			result := r * 0.621371
			data.Resultat = strconv.FormatFloat(result, 'f', 6, 64)
			data.Calcul = data.Valeur1 + " " + data.Unité1 + " -> " + data.Resultat + " " + data.Unité2
		case "Miles":
			data.Valeur1 = strconv.FormatFloat(r, 'f', 6, 64)
			result := r / 0.621371
			data.Resultat = strconv.FormatFloat(result, 'f', 6, 64)
			data.Calcul = data.Valeur1 + " " + data.Unité1 + " -> " + data.Resultat + " " + data.Unité2
		case "Kilogrammes":
			data.Valeur1 = strconv.FormatFloat(r, 'f', 6, 64)
			result := r * 2.20462
			data.Resultat = strconv.FormatFloat(result, 'f', 6, 64)
			data.Calcul = data.Valeur1 + " " + data.Unité1 + " -> " + data.Resultat + " " + data.Unité2
		case "Livres":
			data.Valeur1 = strconv.FormatFloat(r, 'f', 6, 64)
			result := r / 2.20462
			data.Resultat = strconv.FormatFloat(result, 'f', 6, 64)
			data.Calcul = data.Valeur1 + " " + data.Unité1 + " -> " + data.Resultat + " " + data.Unité2
		case "Celsius":
			data.Valeur1 = strconv.FormatFloat(r, 'f', 6, 64)
			result := (r * 9 / 5) + 32
			data.Resultat = strconv.FormatFloat(result, 'f', 6, 64)
			data.Calcul = data.Valeur1 + " " + data.Unité1 + " -> " + data.Resultat + " " + data.Unité2
		case "Fahrenheit":
			data.Valeur1 = strconv.FormatFloat(r, 'f', 6, 64)
			result := (r - 32) * 5 / 9
			data.Resultat = strconv.FormatFloat(result, 'f', 6, 64)
			data.Calcul = data.Valeur1 + " " + data.Unité1 + " -> " + data.Resultat + " " + data.Unité2
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
