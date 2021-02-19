package db

import (
	"fmt"
	"new/test/project/api/constants"
	"new/test/project/api/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(postgres.Open(constants.DbName), &gorm.Config{})
	if err != nil {
		fmt.Println("Unable to connect to database", err)
		os.Exit(1)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Unable to creste user table", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&model.CPU{})
	if err != nil {
		fmt.Println("Unable to creste CPU table", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&model.RAM{})
	if err != nil {
		fmt.Println("Unable to creste RAM table", err)
		os.Exit(1)
	}
}
