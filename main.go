package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	// func ParseFiles(filenames ...string) (*Template, error)
	// Parse the template file and create a new Template instance
	var err error
	tpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template files: %v", err)// log for os.Stderr ,Fatalf typically used for logging fatal errors that require the program to stop immediately
	}

	// Register the indexHandler function to handle requests to the root path ("/")
	http.HandleFunc("/", indexHandler)

	// Start the HTTP server on port 8080
	err = http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Execute the template and send the rendered HTML to the client
	// func (t *Template) Execute(wr io.Writer, data interface{}) error
	err := tpl.Execute(w, nil)
	if err != nil {
		// If there's an error executing the template, send a 500 Internal Server Error response
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}
