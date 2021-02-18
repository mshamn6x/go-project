package db

import (
	"log"
	"new/test/project/api/model"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Get(int) (*model.User, error)
	Insert(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(int) (int, error)
	Login(string, string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}

}

func (u *UserDao) Get(id int) (*model.User, error) {
	user := &model.User{}
	result := db.First(user, id)
	return user, result.Error
}

func (u *UserDao) Insert(data *model.User) (*model.User, error) {
	result := db.Create(data)
	return data, result.Error
}

func (u *UserDao) Update(data *model.User) (*model.User, error) {
	result := db.Save(&data)
	data.Password = ""
	return data, result.Error
}

func (u *UserDao) Delete(id int) (int, error) {
	user := &model.User{}
	result := db.Delete(&user, id)
	return int(user.ID), result.Error
}

// get collections
func (u *UserDao) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	result := db.Find(&users)
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}

	return users, result.Error
}

func (u *UserDao) Login(user string, password string) (*model.User, error) {
	usermodel := &model.User{}
	result := db.Where("email = ? or mobile = ?", user, user).First(&usermodel)
	if result.Error != nil {
		log.Println("Given Username is not matched", result.Error)
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(usermodel.Password), []byte(password))

	if err != nil {
		log.Println("Given PAssword is not matched", err)
		return nil, err
	}
	usermodel.Password = ""
	return usermodel, nil

}
