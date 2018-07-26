package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"google.golang.org/appengine" // Required external App Engine library
)

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	params := struct {
		Greeting string
	}{
		fmt.Sprintf("%s!", strings.Title(os.Getenv("GREETING"))),
	}

	if r.Method == http.MethodGet {
		indexTemplate.Execute(w, params)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
