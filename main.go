package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` //tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("MYSQL_SRC")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open")
	}

	// Create a new restaurant
	newRestaurant := Restaurant{Name: "American Food", Addr: "2 Story Road"}
	if err := db.Create(&newRestaurant).Error; err != nil {
		log.Println("Failed to create")
	}
	log.Println("New Id", newRestaurant.Id)

	// Get a restaurant by ID from the database
	var myRestaurant Restaurant
	if err := db.Where("id = ?", 1).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	}
	log.Println(myRestaurant)

	// Update restaurant's info by ID from the database
	newName := "Korean Food"
	updateData := RestaurantUpdate{Name: &newName}
	if err := db.Where("id = ?", 1).Updates(&updateData).Error; err != nil {
		log.Println(err)
	}
	log.Println("Updated Restaurant:", myRestaurant)

	// Delete restaurant's info by ID from the database
	if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 7).Delete(nil).Error; err != nil {
		log.Println(err)
	}

}
