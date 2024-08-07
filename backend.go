package Server

import (
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Welcome" { // check if it's us the right path
		http.Error(w, "404 not found", http.StatusNotFound) // send error
		return
	}
	if r.Method != "GET" { // accept only Get request
		http.Error(w, "Method is Not supported", http.StatusNotFound) //send error
		return
	}
	fmt.Fprintf(w, "Welcome")
}
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // check if there any error in connection
		fmt.Fprint(w, "ParseForm() err: %V", err)
	}
	fmt.Fprint(w, "POST request successful\n")
	name := r.FormValue("FullName")
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	checkForm(name, email, password, w)
}
func checkForm(name string, email string, password string, w http.ResponseWriter) {
	if email == "test@test.com" && password == "password" {
		fmt.Fprintf(w, "Welcome Mr.%s in ur website", name)
	}
}
