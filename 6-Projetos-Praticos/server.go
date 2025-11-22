package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{Name: "pen", Price: 1.50},
	{Name: "notebook", Price: 3.00},
	{Name: "eraser", Price: 0.50},
	{Name: "ruler", Price: 2.00},
}

func getProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProductbyName(c echo.Context) error {
	productName := c.Param("Name")

	for _, p := range products {

		if p.Name == productName {
			return c.JSON(http.StatusOK, p)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
}

func helloWord(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.GET("/products", getProduct)
	e.GET("/products/:Name", getProductbyName)
	e.GET("/", helloWord)

	e.Logger.Fatal(e.Start(":8080"))

}
