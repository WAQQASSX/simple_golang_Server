package main

import (
	"fmt"
	"log"
	"net/http"
	"Server"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // this variable store where the files
	http.Handle("/", fileServer)                        // Create the root file
	http.HandleFunc("/form", Server.FormHandler)               //Create file inside root
	http.HandleFunc("/Welcome", Server.WelcomeHandler)         //Create file inside the root

	fmt.Printf("Server connected to Port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
