package models

import (
	"github.com/arthurc0102/gin-auth/db"
	"github.com/gin-gonic/gin/binding"
	"github.com/leebenson/conform"
)

// User model
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" gorm:"not null;size:50;unique" conform:"trim"`
	Password  string `json:"password" gorm:"not null;size:60" conform:"trim"`
	FirstName string `json:"firstName" gorm:"not null;size:50" conform:"trim"`
	LastName  string `json:"lastName" gorm:"not null;size:50" conform:"trim"`
	Gender    string `json:"gender" gorm:"not null;size:1" conform:"trim"`
	Age       uint   `json:"age" gorm:"not null"`
}

// Save model
func (model *User) Save(check ...bool) error {
	if len(check) == 0 || !check[0] {
		model.Conform()
		if err := model.Validate(); err != nil {
			return err
		}
	}

	if db.Connection.NewRecord(model) {
		return db.Connection.Create(model).Error
	}

	return db.Connection.Save(model).Error
}

// Delete model
func (model *User) Delete() error {
	if db.Connection.NewRecord(model) {
		return nil
	}

	return db.Connection.Delete(model).Error
}

// Validate model
func (model *User) Validate() error {
	return binding.Validator.ValidateStruct(model)
}

// Conform model's string
func (model *User) Conform() {
	conform.Strings(model)
}
