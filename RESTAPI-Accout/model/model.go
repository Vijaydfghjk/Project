package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Adharr string `json:"adharr"`
}

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Process interface {
	Createaccount(mine Account) (Account, error)
	Viewaccount() ([]Account, error)
	Getaccount(id int) (Account, error)
	Updateaccount(mine Account) (Account, error)
	Deleteaccount(id int) (Account, error)
	RegisterUser(user User) (User, error)
	LoginUser(email, password string) (User, error)
}

type Storage struct {
	db *gorm.DB
}

func Newmodel(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Createaccount(mine Account) (Account, error) {

	err := s.db.Create(&mine).Error

	if err != nil {

		return mine, err
	}

	return mine, nil
}

func (s *Storage) Viewaccount() ([]Account, error) {

	var temp []Account
	err := s.db.Find(&temp).Error

	if err != nil {

		return temp, err
	}
	return temp, nil
}

func (s *Storage) Getaccount(id int) (Account, error) {

	var temp Account

	err := s.db.Where("ID=?", id).Find(&temp).Error

	if err != nil {

		return temp, err
	}
	return temp, nil
}

func (s *Storage) Updateaccount(mine Account) (Account, error) {

	err := s.db.Save(&mine).Error

	if err != nil {

		return mine, err
	}

	return mine, nil
}

func (s *Storage) Deleteaccount(id int) (Account, error) {

	var temp Account

	err := s.db.Where("ID=?", id).Delete(&temp).Error

	if err != nil {

		return temp, err
	}
	return temp, nil
}

func (s *Storage) RegisterUser(user User) (User, error) {

	person := User{
		Name:  user.Name,
		Email: user.Email,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return person, err
	}

	person.Password = string(passwordHash)

	err = s.db.Create(&person).Error

	if err != nil {

		return person, err
	}

	return person, nil
}

func (s *Storage) LoginUser(email, password string) (User, error) {

	var person User

	err := s.db.Where("email = ?", email).Find(&person).Error

	if err != nil {

		return person, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))

	if err != nil {

		return person, errors.New("invalid credentials")
	}

	return person, nil
}
