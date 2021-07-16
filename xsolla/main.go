package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/catalog")
	if err != nil {
		panic(err)
	}
	defer db.Close()


	insert, err := db.Query("INSERT INTO items (sku, name, category, price) VALUES ('12jfdd2','PS5','console','49990')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	fmt.Println("Состыковочка")

	
	//Создание товара
//	r.POST("/", func(c *gin.Context) {})

	//Редактирование товара
	//r.PUT("/", func(c *gin.Context) {})

	//Удаление товара по его идентификатору или SKU
	//r.DELETE("/", func(c *gin.Context) {})

	//Получение информации о товаре по его идентификатору или SKU
	//r.GET("/", func(c *gin.Context) {})

	//Получение каталога товаров
	/*r.GET("/", func(c *gin.Context) {

	})*/

	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}