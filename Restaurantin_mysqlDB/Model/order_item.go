package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order_item struct {
	gorm.Model
	Quantity      int    `json:"quantity" validate:"required"`
	Food_id       string `json:"food_id" validate:"required"`
	Table_id      string `json:"table_id" validate:"required"`
	Order_item_id string `gorm:"primaryKey;unique;" json:"order_item_id"`
}

type Orderprocess interface {
	Create_orderItem(item []Order_item) ([]Order_item, error)
	Vieworders() ([]Order_item, error)
	Filter(id string) (Order_item, error)
	Update_orderItem(item Order_item) (Order_item, error)
	Delete_orderItem(id string) (Order_item, error)
}

type OrderDB struct {
	db *gorm.DB
}

func Order_repo(db *gorm.DB) *OrderDB {

	return &OrderDB{db: db}
}

func (a *OrderDB) Create_orderItem(item []Order_item) ([]Order_item, error) {

	for i, _ := range item {

		item[i].Order_item_id = uuid.New().String()
	}

	err := a.db.Create(&item).Error

	if err != nil {

		return item, err
	}

	return item, nil
}

func (a *OrderDB) Vieworders() ([]Order_item, error) {

	var temp []Order_item
	err := a.db.Find(&temp).Error

	if err != nil {

		return temp, err
	}

	return temp, nil
}

func (a *OrderDB) Filter(id string) (Order_item, error) {

	var temp Order_item
	err := a.db.Where("order_item_id=?", id).Find(&temp).Error

	if err != nil {

		return temp, err
	}

	return temp, nil
}
func (a *OrderDB) Update_orderItem(item Order_item) (Order_item, error) {

	err := a.db.Save(&item).Error

	if err != nil {

		return item, err
	}

	return item, nil
}

func (a *OrderDB) Delete_orderItem(id string) (Order_item, error) {

	var temp Order_item

	err := a.db.Where("order_item_id=?", id).Delete(&temp).Error

	if err != nil {

		return temp, err
	}

	return temp, nil
}
