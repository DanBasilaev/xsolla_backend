package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type item struct {
	id int
	sku string
	name string
	category string
	price int
}

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/catalog")
	if err != nil {
		panic(err)
	}
	defer db.Close()


	/*insert, err := db.Query("INSERT INTO items (sku, name, category, price) VALUES ('12jfdd2','PS5','console','49990')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	fmt.Println("Состыковочка")*/


	//Создание товара
	r.POST("/addItem", func(c *gin.Context) {
		sku := c.Query("sku")
		name := c.Query("name")
		category := c.Query("category")
		price := c.Query("price")

		insert, err := db.Query("INSERT INTO items (sku, name, category, price) VALUES (?, ?, ?, ?)", sku, name, category, price)
		if err != nil {
			panic(err)
		}

		defer insert.Close()

		get_id, err := db.Query("SELECT id FROM items where sku=?", sku)
		if err != nil {
			panic(err)
		}

		get_id.Next()
		var id int
		err = get_id.Scan(&id)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"id":id})

		})

	//Редактирование товара
	//r.PUT("/", func(c *gin.Context) {})

	//Удаление товара по его идентификатору или SKU
	//r.DELETE("/", func(c *gin.Context) {})

	//Получение информации о товаре по его идентификатору или SKU
	r.GET("/getItem", func(c *gin.Context) {
		id := c.Query("id")
		get_item, err := db.Query("SELECT * FROM items where id=?", id)
		if err != nil {
			panic(err)
		}

		get_item.Next()
		var itm item
		err = get_item.Scan(&itm.id, &itm.sku, &itm.name, &itm.category, &itm.price)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"id":itm.id, "sku":itm.sku, "name":itm.name, "category":itm.category, "price":itm.price})

	})

	//Получение каталога товаров
	/*r.GET("/", func(c *gin.Context) {

	})*/

	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}