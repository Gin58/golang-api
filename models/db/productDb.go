package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	entity "../entity"
)

func open() *gorm.DB {
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
	db.AutoMigrate(&entity.Product{})
	fmt.Println("db connected: ", &db)
	return db
}

func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := open()
	db.Order("ID asc").Find(&products)
	defer db.Close()
	return products
}

func FindProduct(productID int) []entity.Product {
	product := []entity.Product{}

	db := open()
	db.First(&product, productID)
	defer db.Close()
	return product
}

func InsertProduct(registerProduct *entity.Product) {
	db := open()
	db.Create(&registerProduct)
	defer db.Close()
}

func UpdateStateProduct(productID int, productState int) {
	product := []entity.Product{}

	db := open()
	db.Model(&product).Where("ID = ?", productID).Update("State", productState)
	defer db.Close()
}

func DeleteProduct(productID int) {
	product := []entity.Product{}

	db := open()
	db.Delete(&product, productID)
	defer db.Close()
}
