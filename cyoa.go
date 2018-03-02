package cyoa

import (
	"encoding/json"
	"io"
	"log"

	"github.com/pkg/errors"
)

// Chapter type represents single advanture of a story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Story type aggregate all possible advantures together with it's outcomes
type Story map[string]Chapter

// New returns your new adventure story
func New(r io.Reader) Story {
	var plot Story
	dec := json.NewDecoder(r)
	err := dec.Decode(&plot)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Couldn't parse json"))
	}
	return plot
}
