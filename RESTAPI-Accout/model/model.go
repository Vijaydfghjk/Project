package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Adharr string `json:"adharr"`
}

type Process interface {
	Createaccount(mine Account) (Account, error)
	Viewaccount() ([]Account, error)
	Getaccount(id int) (Account, error)
	Updateaccount(mine Account) (Account, error)
	Deleteaccount(id int) (Account, error)
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
