package Model

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

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

	err := s.DB.Where("ID=?", id).Delete(dum).Error

	if err != nil {

		return dum, err
	}

	return dum, nil
}
