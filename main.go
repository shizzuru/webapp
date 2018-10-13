package main

import (
	"html/template"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func home(w http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.html", "templates/menu.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	log.Fatal(http.ListenAndServeTLS(":443", "tls/fullchain.pem", "tls/privkey.pem", nil))
}
