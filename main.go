package main

import (
	"html/template"
	"net/http"
	"os"

	"google.golang.org/appengine" // Required external App Engine library
)

var (
	htmlFuncs = template.FuncMap{
		"comment": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	templates = template.Must(template.New("index.gohtml").Funcs(htmlFuncs).ParseFiles("index.gohtml", "favicon.html", "github.html", "opengraph.html", "roboto.css", "index.css"))
)

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Redirects all invalid/unhandled URLs to the root homepage
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	addHeaders(w)

	params := struct {
		Greeting string
	}{
		os.Getenv("GREETING"),
	}

	if r.Method == http.MethodGet {
		templates.Execute(w, params)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func addHeaders(w http.ResponseWriter) {
	headers := map[string]string{
		// Caching
		"Cache-Control": "no-cache, no-store, must-revalidate",
		"Pragma":        "no-cache",

		// Security
		"Strict-Transport-Security": "max-age=300; includeSubDomains", // Set up STS with 5 minute max-age
		"X-Content-Type-Options":    "nosniff",                        // Opt-out of MIME type sniffing
		"X-Frame-Options":           "DENY",                           // Clickjacking defense
		"X-XSS-Protection":          "1; mode=block",                  // Ensure the browser's XSS filter is enabled
	}

	for k, v := range headers {
		w.Header().Add(k, v)
	}
}
