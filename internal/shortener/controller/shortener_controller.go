package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/tudemaha/toode/internal/shortener/service"
)

func ShortenerSolverHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Not Found", http.StatusNotFound)
		}

		getShorten(&w, r)
	}
}

func getShorten(w *http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	// paths[2] is the shorten url
	if len(paths) != 2 || paths[1] == "" {
		http.Redirect(*w, r, os.Getenv("DEFAULT_REDIRECT"), http.StatusSeeOther)
	}

	url, err := service.GetRedirection(paths[1])
	if err != nil {
		http.Redirect(*w, r, os.Getenv("DEFAULT_REDIRECT"), http.StatusSeeOther)
	}

	http.Redirect(*w, r, url, http.StatusSeeOther)
}
