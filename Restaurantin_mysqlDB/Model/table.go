package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Table_number    string `json:"table_number" validate:"required"`
	Number_of_guest int    `json:"number_of_guest" validate:"required"`
	Table_id        string `gorm:"primaryKey;unique;" json:"table_id"`
}

type TableProcess interface {
	Createtable(t Table) (Table, error)
	Get_All() ([]Table, error)
	Viewtable(id string) (Table, error)
	Update_table(t Table) (Table, error)
	Delete_table(id string) (Table, error)
}

type Table_Db struct {
	db *gorm.DB
}

func Table_Repo(db *gorm.DB) *Table_Db {

	return &Table_Db{db: db}
}

func (a *Table_Db) Createtable(t Table) (Table, error) {

	t.Table_id = uuid.New().String()

	err := a.db.Create(&t).Error

	if err != nil {

		return t, err
	}

	return t, nil
}

func (a *Table_Db) Get_All() ([]Table, error) {

	var temp []Table
	err := a.db.Find(&temp).Error

	if err != nil {

		return temp, err
	}
	return temp, err
}

func (a *Table_Db) Viewtable(id string) (Table, error) {

	var temp Table
	err := a.db.Where("table_number=?", id).Find(&temp).Error
	if err != nil {

		return temp, err
	}

	return temp, nil
}

func (a *Table_Db) Update_table(t Table) (Table, error) {

	err := a.db.Save(&t).Error

	if err != nil {

		return t, err
	}
	return t, nil
}

func (a *Table_Db) Delete_table(id string) (Table, error) {

	var temp Table

	err := a.db.Where("table_number=?", id).Unscoped().Delete(&temp).Error

	if err != nil {

		return temp, err
	}

	return temp, nil
}
