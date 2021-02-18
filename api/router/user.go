package router

import (
	"encoding/json"
	"log"
	"net/http"
	"new/test/project/api/auth"
	"new/test/project/api/db"
	"new/test/project/api/model"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Token string
	*model.User
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Unable to Parse json request data")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var db_connector db.User
	db_connector = db.NewUserDao()

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println("Invalid password")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = string(bytes)

	data, err := db_connector.Insert(&user)
	if err != nil {
		log.Println("Unable to create user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data.Password = ""
	json_data, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json_data)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id_var := mux.Vars(r)
	value, _ := strconv.Atoi(id_var["id"])
	var db_connector db.User
	db_connector = db.NewUserDao()
	data, err := db_connector.Get(value)
	if err != nil {
		log.Println("Unable to get user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data.Password = ""
	json_data, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var db_connector db.User
	db_connector = db.NewUserDao()
	data, err := db_connector.GetAll()
	if err != nil {
		log.Println("Unable to get all users")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json_data, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Unable to Parse json request data")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var db_connector db.User
	db_connector = db.NewUserDao()
	idVar := mux.Vars(r)
	value, _ := strconv.Atoi(idVar["id"])
	user.ID = uint(value)
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println("Invalid password")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = string(bytes)

	data, err := db_connector.Update(&user)
	if err != nil {
		log.Println("Unable to update user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json_data, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idVar := mux.Vars(r)
	value, _ := strconv.Atoi(idVar["id"])
	var db_connector db.User
	db_connector = db.NewUserDao()
	id, err := db_connector.Delete(value)
	if err != nil {
		log.Println("Unable to delete user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(strconv.Itoa(id)))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user_name")
	password := r.FormValue("password")

	user := db.NewUserDao()
	res, err := user.Login(username, password)
	if err != nil {
		log.Println("Unable to login user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := auth.New().CreateToken(res)
	if err != nil {
		log.Println("Unable to create a token. Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json_data, _ := json.Marshal(LoginResponse{Token: token, User: res})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data)
}

func MiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		uri := r.RequestURI
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if uri == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		bearerToken := r.Header.Get("Authorization")
		jwtToken := strings.Replace(bearerToken, "Bearer ", "", 1)
		log.Println(jwtToken)
		isValid, err := auth.New().ValidateToken(jwtToken)
		if !isValid || err != nil {
			log.Println("InValid Token")
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
