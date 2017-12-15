package models

import "github.com/jinzhu/gorm"

type (
	//UserModel describes a usermodel type
	UserModel struct {
		gorm.Model
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	//TransformedUser represents a formatted user
	TransformedUser struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
