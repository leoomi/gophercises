package handlers

import (
	"exercise3/handlers/template"
	"exercise3/models"
	"net/http"
)

// StoryHandler ... burh
func StoryHandler(storyArcs []models.StoryArc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		storyArc := models.StoryArc{}
		template.WriteStoryTemplate(w, storyArc)
	}
}
