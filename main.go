package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
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
	// db connection
	dsn := os.Getenv("MYSQL_SRC") //"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	// Log all SQL statements into the console
	db.Debug()

	// Use Gin to define routes
	r := gin.Default()
	v1 := r.Group("/v1")

	// register routes here
	restaurants := v1.Group("/restaurants")

	// CREATE API - Create a new restaurant
	restaurants.POST("/", func(c *gin.Context) {
		var data Restaurant
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		db.Create(&data)
		c.JSON(http.StatusCreated, data)
	})

	// GET/READ API - Get a restaurant by ID from the database
	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		var data Restaurant
		db.Where("id = ?", id).First(&data)
		c.JSON(http.StatusOK, data)

	})
	// GET/READ All API - Get all restaurants with page and limit parameters
	restaurants.GET("/", func(c *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if pagingData.Page < 1 {
			pagingData.Page = 1
		}
		if pagingData.Limit < 1 {
			pagingData.Limit = 10
		}
		var data []Restaurant
		db.Offset((pagingData.Page - 1) * pagingData.Limit).Order("id desc").Limit(pagingData.Limit).Find(&data)
		c.JSON(http.StatusOK, data)

	})
	// UPDATE API - Update restaurant's info by ID
	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		var data RestaurantUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		db.Where("id =?", id).Updates(&data)
		c.JSON(http.StatusOK, data)
	})

	// DELETE API Delete restaurant's info by ID from the database
	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		db.Where("id =?", id).Delete(&Restaurant{})
		c.JSON(http.StatusOK, gin.H{
			"message": "deleted",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
