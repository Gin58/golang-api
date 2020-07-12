package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID          int    `gorm:"primary_key;not null"`
	ProductName string `gorm:"type:varchar(200);not null"`
	Memo        string `gorm:"type:varchar(400)"`
	Status      string `gorm:"type:char(2);not null"`
}

func getGormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Shopping"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&Product{})
	fmt.Println("db connected: ", &db)
	return db
}

func insertProduct(registerProduct *Product) {
	db := getGormConnect()
	db.Create(&registerProduct)
	defer db.Close()
}

func findAllProduct() []Product {
	db := getGormConnect()
	var products []Product

	db.Order("ID asc").Find(&products)
	defer db.Close()
	return products
}


func main() {
	var product = Product{
		ProductName: "テスト",
		Memo: "一旦テスト",
		Status: "01",
	}
	insertProduct(&product)
	resultProducts := findAllProduct()

	for i := range resultProducts {
		fmt.Printf("index: %d, 商品ID: %d, 商品名: %s, メモ: %s, ステータス: %s\n",
			i, resultProducts[i].ID, resultProducts[i].ProductName, resultProducts[i].Memo, resultProducts[i].Status)
	}
}