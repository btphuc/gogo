package models

import "github.com/jinzhu/gorm"

type (
	//UserModel describes a usermodel type
	UserModel struct {
		gorm.Model
		Name     string `form:"name" json:"name"`
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}

	//TransformedUser represents a formatted user
	TransformedUser struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
