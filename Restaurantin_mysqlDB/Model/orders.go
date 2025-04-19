package model

import "gorm.io/gorm"

type OrderSummary struct {
	gorm.Model
	FoodName string `json:"foodname" gorm:"column:foodname"`
	Price    int    `json:"price" gorm:"column:price"`
	Quantity int    `json:"quantity" gorm:"column:quantity"`
}

func CreateOrdersView(db *gorm.DB) error {

	dropQuery := "DROP VIEW IF EXISTS orders;"

	if err := db.Exec(dropQuery).Error; err != nil {
		return err
	}

	querry := `

	   create view orders as 

	   select 
	      foods.foodname,
		  foods.price,
		  order_items.quantity,
		  tables.number_of_guest, 
		  tables.table_number
          from foods
          inner join order_items 
		  on order_items.food_id = foods.foodid
          inner join tables 
		  on tables.table_id = order_items.table_id;

    `

	db.Exec(querry)
	return nil
}
