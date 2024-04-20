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
		{Title: "Do you really need a Platform?", Author: "Mike", Body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Maecenas sed enim ut sem viverra aliquet eget sit. Id venenatis a condimentum vitae sapien pellentesque habitant. A cras semper auctor neque vitae tempus quam pellentesque. Nulla malesuada pellentesque elit eget gravida cum. Turpis in eu mi bibendum neque egestas. Vitae congue mauris rhoncus aenean vel. Suscipit adipiscing bibendum est ultricies integer quis auctor. Ultricies leo integer malesuada nunc vel risus commodo viverra. Amet cursus sit amet dictum sit amet justo. Sed id semper risus in hendrerit gravida rutrum quisque. Vulputate sapien nec sagittis aliquam malesuada bibendum arcu."},
		{Title: "What is a Platform?", Author: "Mike", Body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Maecenas sed enim ut sem viverra aliquet eget sit. Id venenatis a condimentum vitae sapien pellentesque habitant. A cras semper auctor neque vitae tempus quam pellentesque. Nulla malesuada pellentesque elit eget gravida cum. Turpis in eu mi bibendum neque egestas. Vitae congue mauris rhoncus aenean vel. Suscipit adipiscing bibendum est ultricies integer quis auctor. Ultricies leo integer malesuada nunc vel risus commodo viverra. Amet cursus sit amet dictum sit amet justo. Sed id semper risus in hendrerit gravida rutrum quisque. Vulputate sapien nec sagittis aliquam malesuada bibendum arcu."},
	}
	tmpl.Execute(w, data)
}
