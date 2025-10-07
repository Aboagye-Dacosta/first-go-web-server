package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//http
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/hello", handleSayHello)
	http.HandleFunc("/contact", handleContact)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
