package controller

import (
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	Calcul  string
	Message string
}

func Calcul(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nb1 := r.FormValue("nombre1")
		nb2 := r.FormValue("nombre2")
		f1, err := strconv.ParseFloat(nb1, 64)
		if err != nil {
			data := PageData{
				Calcul:  nb1 + " + " + nb2 + " = ERREUR",
				Message: "Entree invalide pour le premier nombre.",
			}
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, data)
			return
		}

		f2, err := strconv.ParseFloat(nb2, 64)
		if err != nil {
			data := PageData{
				Calcul:  nb1 + " + " + nb2 + " = ERREUR",
				Message: "Entree invalide pour le deuxième nombre.",
			}
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, data)
			return
		}
		if r.FormValue("opération") == "Add" {
			result := strconv.FormatFloat(f1+f2, 'f', -1, 64)
			data := PageData{
				Calcul:  nb1 + " + " + nb2 + " = " + result,
				Message: "Une super addition ça ! ➕",
			}
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, data)
			return
		} else if r.FormValue("opération") == "Sou" {
			result := strconv.FormatFloat(f1-f2, 'f', -1, 64)
			data := PageData{
				Calcul:  nb1 + " - " + nb2 + " = " + result,
				Message: "Une sacrée soustraction ça ! ➖",
			}
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, data)
			return
		} else if r.FormValue("opération") == "Mul" {
			result := strconv.FormatFloat(f1*f2, 'f', -1, 64)
			data := PageData{
				Calcul:  nb1 + " * " + nb2 + " = " + result,
				Message: "Une belle multiplication ça ! ✖️",
			}
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, data)
			return
		} else if r.FormValue("opération") == "Div" {
			if nb1 == "0" || nb2 == "0" {
				data := PageData{
					Calcul:  "Erreur : Division par zéro impossible !",
					Message: "Réessaie avec des nombres valides (pas de 0 en division) ⚠️",
				}
				tmpl := template.Must(template.ParseFiles("template/index.html"))
				tmpl.Execute(w, data)
				return
			} else {
				result := strconv.FormatFloat(f1/f2, 'f', -1, 64)
				data := PageData{
					Calcul:  nb1 + " / " + nb2 + " = " + result,
					Message: "Une omega super géniale division ça ! ➗",
				}
				tmpl := template.Must(template.ParseFiles("template/index.html"))
				tmpl.Execute(w, data)
				return
			}
		}
	}
	data := PageData{
		Calcul:  "Prépare ton calcul avec Das calculatrice !",
		Message: "Bienvenue 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
