package handlers

import (
	"exercise3/handlers/template"
	"exercise3/models"
	"net/http"
)

// StoryHandler ... burh
func StoryHandler(storyArcs map[string]models.StoryArc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path[1:]

		if url == "" {
			url = "intro"
		}

		storyArc, ok := storyArcs[url]
		if !ok {
			http.NotFound(w, r)
			return
		}

		template.WriteStoryTemplate(w, storyArc)
	}
}
