package models

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// StoryArc ... bruh
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
