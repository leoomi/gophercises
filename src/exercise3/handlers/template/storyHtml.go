package template

import (
	"exercise3/models"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// StoryTemplate ... bru
func WriteStoryTemplate(w http.ResponseWriter, storyArc models.StoryArc) {
	var tmpl, err = template.ParseFiles("html/story.html")

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	tmpl.Execute(w, storyArc)
}
