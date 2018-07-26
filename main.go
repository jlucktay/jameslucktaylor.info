package main

import (
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)

type templateParams struct {
	Notice string
	Name   string
}

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

	// fmt.Fprintln(w, strings.Title(os.Getenv("GREETING")), "!")

	params := templateParams{}

	if r.Method == "GET" {
		indexTemplate.Execute(w, params)
		return
	}

	// It's a POST request, so handle the form submission.

	name := r.FormValue("name")
	params.Name = name // Preserve the name field.
	if name == "" {
		name = "Anonymous Gopher"
	}

	if r.FormValue("message") == "" {
		w.WriteHeader(http.StatusBadRequest)

		params.Notice = "No message provided"
		indexTemplate.Execute(w, params)
		return
	}

	// TODO: save the message into a database.

	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", name)

	indexTemplate.Execute(w, params)
}
