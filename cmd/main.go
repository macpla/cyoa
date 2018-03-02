package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/macpla/cyoa"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").ParseGlob("index.html"))
}

func main() {

	b, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(b)
	story := cyoa.New(r)
	chapterFn := func(chap cyoa.Chapter) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := tmpl.ExecuteTemplate(w, "index.html", chap)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	mux := http.NewServeMux()
	for key, chapter := range story {
		path := "/" + key
		mux.HandleFunc(path, chapterFn(chapter))
	}

	mux.HandleFunc("/", chapterFn(story["intro"]))
	fmt.Println("Listen and Serve on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func intro(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", "Choose your own advanture")
	if err != nil {
		log.Fatal(err)
	}
}
