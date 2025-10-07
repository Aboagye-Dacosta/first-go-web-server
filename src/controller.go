package main

import (
	"fmt"
	"log"
	"net/http"
)

type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func handleSayHello(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == http.MethodGet {
		w.Write([]byte("Hello world"))
	}
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to validate form", http.StatusInternalServerError)
		return
	}

	contact := ContactForm{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	log.Print(contact)
	// For demo: just print and confirm
	fmt.Printf("Received message: %+v\n", contact)
	fmt.Fprintf(w, "Thanks %s! Your message has been received.", contact.Name)
}
