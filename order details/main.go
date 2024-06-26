//Order Details Management

package main

import (
	"Order/management/Mysql/mysql"
	"database/sql"
	"fmt"
	"net/http"

	//its a third party package having its own server for fast routing
	//Gin takes advantage of httprouter, which is faster than the operations to be executed by net/http package.
	"github.com/gin-gonic/gin"
)

type OrderDetails struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Brand string  `json:"brand"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

func init() {
	dataBase, err := sql.Open("mysql", mysql.DataBaseURL)
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	err = dataBase.Ping()
	if err != nil {
		fmt.Println("Database ping error:", err)
		return
	}
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templet/*.html")
	authentication := router.Group("/")
	authentication.Use()
	{
		authentication.GET("/viewcart", ViewCart)
		authentication.GET("/viewOrderDetails", viewOrderDetails)
	}
	http.ListenAndServe(":3333", router)
}
func ViewCart(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "viewCart.html", nil)
}
func viewOrderDetails(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "viewCart.html", nil)
}
