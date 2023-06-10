package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khoaphungnguyen/food_delivery/component/appcontext"
	"github.com/khoaphungnguyen/food_delivery/middleware"
	"github.com/khoaphungnguyen/food_delivery/module/restaurants/transport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}
	// db connection
	dsn := os.Getenv("MYSQL_SRC") //"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	// Log all SQL statements into the console
	db.Debug()

	appContext := appcontext.NewAppContext(db)
	// Use Gin to define routes
	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")

	// register routes here
	restaurants := v1.Group("/restaurants")

	// CREATE API - Create a new restaurant
	restaurants.POST("/", ginrestaurant.CreateRestaurants(appContext))

	// GET API - Get a restaurant by ID from the database
	restaurants.GET("/:id", ginrestaurant.ListByIdRestaurants(appContext))

	// GET All API - Get all restaurants with page and limit parameters
	restaurants.GET("/", ginrestaurant.ListRestaurants(appContext))

	// UPDATE API - Update restaurant's info by ID
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurants(appContext))

	// DELETE API Delete restaurant's info by ID from the database
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurants(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
