package routes

import (
	"net/http"

	"github.com/Aaryan376/to-do-list-go/controllers"
)

func SetupRoutes() http.Handler {
	r := http.NewServeMux()

	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

	r.HandleFunc("/", controllers.IndexHandler)
	r.HandleFunc("/create", controllers.CreateHandler)
	r.HandleFunc("/update", controllers.UpdateHandler)
	r.HandleFunc("/delete", controllers.DeleteHandler)

	return r
}
