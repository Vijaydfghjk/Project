package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Foodname string `json:"foodname" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Foodid   string `gorm:"primaryKey;unique;" json:"foodid"`
}

type Foodprocess interface {
	Add_food(f Food) (Food, error)
	Removefood(id string) (Food, error)
	Viewfoods() ([]Food, error)
}

type FoodDB struct {
	db *gorm.DB
}

func Food_repo(db *gorm.DB) *FoodDB {

	return &FoodDB{db: db}
}

func (a *FoodDB) Add_food(f Food) (Food, error) {

	f.Foodid = uuid.New().String()

	err := a.db.Create(&f).Error

	if err != nil {

		return f, err
	}

	return f, nil
}

func (a *FoodDB) Removefood(id string) (Food, error) {

	var temp Food
	err := a.db.Where("foodid=?", id).Delete(&temp).Error

	if err != nil {

		return temp, err
	}

	return temp, nil
}

func (a *FoodDB) Viewfoods() ([]Food, error) {

	var temp []Food
	err := a.db.Find(&temp).Error

	if err != nil {

		return temp, err
	}

	return temp, nil
}
