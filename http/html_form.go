package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Name    string
	Email   string
	Subject string
	Message string
	Success bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("./html/forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			_ = tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
			Success: true,
		}

		// do something with details
		_ = details

		err := tmpl.Execute(w, details)
		if err != nil {
			panic(fmt.Sprintf("ERROR Executing template: %v", err))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Sprintf("ERROR when ListenAndServe : %v", err))
	}
}
