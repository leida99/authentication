package handlers

import (
    "database/sql"
    "html/template"
    "net/http"
    "myapp/database"
    "myapp/models"
    "myapp/utils"
    "log"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl, _ := template.ParseFiles("views/templates/register.html")
        tmpl.Execute(w, nil)
    } else if r.Method == http.MethodPost {
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")

        hashedPassword, err := utils.HashPassword(password)
        if err != nil {
            log.Println(err)
            http.Error(w, "Server error, unable to create your account.", 500)
            return
        }

        user := models.User{Username: username, Email: email, Password: hashedPassword}

        _, err = database.DB.Exec("INSERT INTO users(username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
        if err != nil {
            log.Println(err)
            http.Error(w, "User already exists.", 400)
            return
        }

        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl, _ := template.ParseFiles("views/templates/login.html")
        tmpl.Execute(w, nil)
    } else if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        var user models.User
        row := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username=?", username)
        err := row.Scan(&user.ID, &user.Username, &user.Password)
        if err == sql.ErrNoRows || !utils.CheckPasswordHash(password, user.Password) {
            http.Error(w, "Invalid username or password", 401)
            return
        }

        // User authenticated successfully, redirect to welcome page
        http.Redirect(w, r, "/welcome?username="+username, http.StatusSeeOther)
    }
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    tmpl, _ := template.ParseFiles("views/templates/welcome.html")
    tmpl.Execute(w, struct{ Username string }{Username: username})
}
