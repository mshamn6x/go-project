package router

import (
	"html/template"
	"log"
	"net/http"
	"new/test/project/ui/api"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func New() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/dashboard", DashboardHandler).Methods("GET")

	return r
}

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user_name")
	password := r.FormValue("password")

	loginAPI := api.API{}

	LoginResponse, err := loginAPI.Login(username, password)
	if err != nil {
		log.Println("router/LoginHandler failed to perform login ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(LoginResponse)

	session, _ := store.Get(r, "session-name")

	// Set some session values.
	session.Values["token"] = LoginResponse.Token
	session.Values["name"] = LoginResponse.Name
	session.Values["id"] = LoginResponse.ID

	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		log.Println("router/LoginHandler failed to save session ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	return
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")

	if session.Values["token"] == nil {
		http.Redirect(w, r, "/public/login.html", http.StatusSeeOther)
		return
	}

	templateData := make(map[string]string)

	templateData["Name"] = session.Values["name"].(string)

	templ, err := template.ParseFiles(`C:\Users\mshanm6x\Downloads\project\ui\views\dashboard.html`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ.Execute(w, templateData)
}
