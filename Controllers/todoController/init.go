package todoController

import (
	"github.com/jinzhu/gorm"
	"github.com/user/gogo/models"
)

var db *gorm.DB

func init() {
	// khoi tao bien chua loi
	var err error

	db, err = gorm.Open("mysql", "root:123456159@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("ket noi vao db k dc nha")
	}

	// migrate the schema

	db.AutoMigrate(&models.TodoModel{})
}
