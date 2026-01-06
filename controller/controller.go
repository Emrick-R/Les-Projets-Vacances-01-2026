package controller

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

type PageData struct {
	Message     string
	Debutjeu    bool
	Finjeu      bool
	Score       int
	PremiereRep bool
	DerniereRep string
	Aide        string
	Nbmystere   int
	Historique  []string
	Abandon     bool
}

var data = PageData{
	Message:     "Bienvenue 🎉 Le goofy ahhh nombre mystère se cache !",
	Debutjeu:    false,
	Finjeu:      false,
	Score:       0,
	PremiereRep: false,
	DerniereRep: "",
	Aide:        "",
	Nbmystere:   0,
	Historique:  []string{},
	Abandon:     false,
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("commencer") == "oui" {
		data.Debutjeu = true
		data.PremiereRep = true
		data.Message = "Maintenant saisit un nombre entre 0 et 1000 !"
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}
	if r.FormValue("tentative") != "" {
		if data.PremiereRep {
			data.Nbmystere = rand.Intn(1000) + 1
			data.PremiereRep = false
		}
		data.DerniereRep = r.FormValue("tentative")
		data.Score++
		tentative := r.FormValue("tentative")
		tentativeInt, _ := strconv.Atoi(tentative)
		switch {
		case tentative == "abandonner":
			data.Finjeu = true
			data.Abandon = true
			data.Message = "Tu as abandonné le jeu. Is that a skill issue ?!"
			data.DerniereRep = strconv.Itoa(data.Nbmystere) + ". Mieux vaut réessayer !"
		case tentativeInt < data.Nbmystere:
			data.Aide = "C'est plus grand !"
		case tentativeInt > data.Nbmystere:
			data.Aide = "C'est plus petit !"
		case tentativeInt == data.Nbmystere:
			data.Finjeu = true
			data.DerniereRep = tentative
		}
		data.Historique = append(data.Historique, tentative+" ("+data.Aide+")")
	}
	if r.FormValue("recommencer") == "oui" {
		data = PageData{
			Message:     "Bienvenue 🎉 Le goofy ahhh nombre mystère se cache !",
			Debutjeu:    true,
			Finjeu:      false,
			Score:       0,
			PremiereRep: true,
			DerniereRep: "",
			Aide:        "",
			Nbmystere:   0,
			Abandon:     false,
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
