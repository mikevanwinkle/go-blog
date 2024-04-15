package handlers

import (
	"mikevanwinkle/blog/models"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("recevied request method: %s, path: %s", r.Method, r.URL)
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := map[string]any{}
	data["Posts"] = []models.Post{
		{Title: "one", Author: "Mike", Body: "Test"},
		{Title: "two", Author: "Mike", Body: "Test2"},
	}
	tmpl.Execute(w, data)
}
