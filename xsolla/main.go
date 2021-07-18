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



	//Создание товара
	r.POST("/addItem", func(c *gin.Context) {
		sku := c.Query("sku")
		name := c.Query("name")
		category := c.Query("category")
		price := c.Query("price")

		insert, err := db.Query("INSERT INTO items (sku, name, category, price) VALUES (?, ?, ?, ?)", sku, name, category, price)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}

		defer insert.Close()

		get_id, err := db.Query("SELECT id FROM items where sku=?", sku)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}

		get_id.Next()
		var id int
		err = get_id.Scan(&id)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Empty res"})
			panic(err)
		}

		c.JSON(200, gin.H{"id":id})
		defer get_id.Close()

		})

	//Редактирование товара
	r.PUT("/upItem", func(c *gin.Context) {
		id := c.Query("id")
		get_item, err := db.Query("SELECT id, sku, name, category, price FROM items where id=?", id)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}

		get_item.Next()
		var itm item
		err = get_item.Scan(&itm.id, &itm.sku, &itm.name, &itm.category, &itm.price)
		if err != nil {
			c.JSON(500, gin.H{"massage": "there is no such id"})
			panic(err)
		}

		name := c.Query("name")
		category := c.Query("category")
		price := c.Query("price")

		update, err := db.Query("UPDATE items SET name=?, category=?,price=? where id=?", name, category, price, id)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}

		defer update.Close()

		//c.JSON(200, gin.H{"id":itm.id, "sku":itm.sku, "name":itm.name, "category":itm.category, "price":itm.price})
		defer get_item.Close()

	})

	//Удаление товара по его идентификатору или SKU
	r.DELETE("/delItem", func(c *gin.Context) {
		id := c.Query("id")
		get_id, err := db.Query("SELECT id FROM items where id=?", id)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}
		get_id.Next()
		err = get_id.Scan(&id)
		if err != nil {
			c.JSON(500, gin.H{"massage": "there is no such id"})
			panic(err)
		}
		defer get_id.Close()


		del_item, err := db.Query("DELETE FROM items where id=?", id)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}

		defer del_item.Close()

	})

	//Получение информации о товаре по его идентификатору или SKU
	r.GET("/getItem", func(c *gin.Context) {
		id := c.Query("id")
		get_item, err := db.Query("SELECT id, sku, name, category, price FROM items where id=?", id)
		if err != nil {
			c.JSON(404, gin.H{"massage": "Query error"})
			panic(err)
		}

		get_item.Next()
		var itm item
		err = get_item.Scan(&itm.id, &itm.sku, &itm.name, &itm.category, &itm.price)
		if err != nil {
			c.JSON(500, gin.H{"massage": "there is no such id"})
			panic(err)
		}

		c.JSON(200, gin.H{"id":itm.id, "sku":itm.sku, "name":itm.name, "category":itm.category, "price":itm.price})
		defer get_item.Close()

	})

	//Получение каталога товаров
	r.GET("/getAll", func(c *gin.Context) {
		get_item, err := db.Query("SELECT * FROM items")
		if err != nil {
			panic(err)
		}

		for get_item.Next(){
			var itm item
			err = get_item.Scan(&itm.id, &itm.sku, &itm.name, &itm.category, &itm.price)
			if err != nil {
				//c.JSON(500, gin.H{"massage": "there is no such id"})
				panic(err)
			}

			c.JSON(200, gin.H{"id":itm.id, "sku":itm.sku, "name":itm.name, "category":itm.category, "price":itm.price})
		}

		defer get_item.Close()

	})

	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}