package main

import (
    "log"
    "net/http"
    "myapp/database"
    "myapp/handlers"
    "html/template"
)

func main() {
    database.InitDB()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, _ := template.ParseFiles("views/templates/index.html")
        tmpl.Execute(w, nil)
    })
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/welcome", handlers.WelcomeHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
