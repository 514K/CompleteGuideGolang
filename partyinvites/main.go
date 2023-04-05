package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var responses = make([]*RSVP, 0, 10)
var templates = make(map[string]*template.Template, 3)

type RSVP struct {
	Name, Email, Phone string
	Accept             bool
}

func loadTemplate() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for i, item := range templateNames {
		t, err := template.ParseFiles("layout.html", item+".html")
		if err == nil {
			templates[item] = t
			fmt.Println("Loaded template ", i, item)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

type formData struct {
	*RSVP
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			RSVP:   &RSVP{},
			Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := RSVP{
			Name:   request.Form["name"][0],
			Email:  request.Form["email"][0],
			Phone:  request.Form["phone"][0],
			Accept: request.Form["accept"][0] == "true",
		}

		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone")
		}
		if len(errors) > 0 {
			templates["form"].Execute(writer, formData{
				RSVP:   &responseData,
				Errors: errors,
			})
		} else {
			responses = append(responses, &responseData)

			if responseData.Accept {
				templates["thanks"].Execute(writer, responseData.Name)
			} else {
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}
	}
}

func main() {
	loadTemplate()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("%v\n", "fuck")
	}
}
