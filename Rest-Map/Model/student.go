package Model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Student struct {
	gorm.Model           //`gorm:""json:name`
	Name          string `json:"name"`
	Place         string `json:"place"`
	Contactnumber string `json:"contactnumber"`
	DOB           string `json:"dob"`
}

type Repository interface {
	Createlist(student Student) (Student, error)
	Getall() ([]Student, error)
	GetbyID(id int) (Student, error)
	Update(boy Student) (Student, error)
	Delete(id int) (Student, error)
	Createnewuser(Re Register) (Register, error)
	Loginuser(Lo Login) (Register, error)
}

type Reposit struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Reposit {

	return &Reposit{DB: db}

}

func (s *Reposit) Createlist(student Student) (Student, error) {

	//s.DB.NewRecord(s)
	err := s.DB.Create(&student).Error

	if err != nil {

		return Student{}, err
	}

	return Student{}, nil
}

func (s *Reposit) Getall() ([]Student, error) {

	var person []Student

	err := s.DB.Find(&person).Error

	if err != nil {

		return person, err
	}

	return person, nil
}
func (s *Reposit) GetbyID(id int) (Student, error) {

	var dum Student

	err := s.DB.Where("ID=?", id).Find(&dum).Error

	if err != nil {

		return dum, err
	}

	return dum, nil
}

func (s *Reposit) Update(boy Student) (Student, error) {

	err := s.DB.Save(&boy).Error

	if err != nil {

		return boy, err
	}
	return boy, nil
}

func (s *Reposit) Delete(id int) (Student, error) {

	var dum Student

	err := s.DB.Where("ID=?", id).Delete(&dum).Error

	if err != nil {

		return dum, err
	}

	return dum, nil
}

func (s *Reposit) Createnewuser(Re Register) (Register, error) {

	temp := Register{

		Name:  Re.Name,
		Email: Re.Email,
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(Re.Password), bcrypt.MinCost)

	if err != nil {

		return temp, err
	}

	temp.Password = string(pwd)
	err = s.DB.Create(&temp).Error

	if err != nil {

		return temp, err
	}
	return temp, nil
}

func (s *Reposit) Loginuser(Lo Login) (Register, error) {

	var person Register

	err := s.DB.Where("email = ?", Lo.Email).Find(&person).Error

	if err != nil {

		return person, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(Lo.Password))

	if err != nil {

		return person, errors.New("invalid credentials")
	}

	return person, nil
}
