package main

import (
	"backend/database"
	"backend/database/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := database.Connect()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/products", func(c echo.Context) error {
		var products []models.Product
		db.Find(&products)
		return c.JSON(http.StatusOK, products)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		id := c.Param("id")
		var product models.Product

		if result := db.First(&product, id); result.Error != nil {
			return c.String(http.StatusNotFound, "Database Error")
		}

		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		product := new(models.Product)
		if err := c.Bind(product); err != nil {
			return err
		}
		db.Create(&product)
		return c.String(http.StatusCreated, "Created")
	})

	e.GET("/payments", func(c echo.Context) error {
		var payments []models.Payment
		db.Find(&payments)
		return c.JSON(http.StatusOK, payments)
	})

	e.POST("/payment", func(c echo.Context) error {
		payment := new(models.Payment)
		if err := c.Bind(payment); err != nil {
			return err
		}
		db.Create(&payment)
		return c.String(http.StatusCreated, "Created")
	})

	e.GET("/orders", func(c echo.Context) error {
		var orders []models.Order
		db.Find(&orders)
		return c.JSON(http.StatusOK, orders)
	})

	e.POST("/order", func(c echo.Context) error {
		order := new(models.Order)
		if err := c.Bind(order); err != nil {
			return err
		}
		db.Create(&order)
		return c.String(http.StatusCreated, "Created")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
