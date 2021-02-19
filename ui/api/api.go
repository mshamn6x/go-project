package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"new/test/project/ui/constants"

	"github.com/jinzhu/gorm"
)

type API struct {
}

type User struct {
	gorm.Model
	// Id       int
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Mobile   string `gorm:"uniqueIndex"`
	Password string
}

type LoginResponse struct {
	Token string
	*User
}

func (api *API) Login(username, password string) (LoginResponse, error) {

	data := url.Values{
		"user_name": {username},
		"password":  {password},
	}

	resp, err := http.PostForm(constants.APIServer+constants.Login, data)
	if err != nil {
		log.Println("Error in postform ", err)
		return LoginResponse{}, err
	}

	var res LoginResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Println("Error in response decoder ", err)
		return LoginResponse{}, err
	}
	return res, nil
}

func GetUsers(jwtToken string) ([]User, error) {

	users := []User{}
	req, err := http.NewRequest("GET", constants.APIServer+"/users", nil)
	if err != nil {
		log.Println("Error in NewRequest ", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in Do request ", err)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		log.Println("Error in response decoder ", err)
		return nil, err
	}

	return users, nil
}
