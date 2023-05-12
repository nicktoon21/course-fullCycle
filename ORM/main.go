package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Category struct {
// 	ID int `gorm:"primaryKey"`
// 	Name string
// 	gorm.Model
// }

// type Product struct {
// 	ID int `gorm:"primaryKey"`
// 	Name string
// 	Price float64
// 	CategoryID int
// 	Category Category
// 	SerialNumber SerialNumber
// 	gorm.Model
// }
type Category struct {
	ID int `gorm:"primaryKey"`
	Name string
	Products []Product `gorm:"many2many:products_categories;"`
}



type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

// type SerialNumber struct {
// 	ID int `gorm:"primaryKey"`
// 	Number string
// 	ProductID int 
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); if err!= nil {
    panic(err)
  }
	// products := []Product{
	// 	{Name: "Placa de vídeo", Price: 4000},
	// 	{Name: "Monitor", Price: 600},
	// 	{Name: "Mouse", Price: 55},
	// 	{Name: "Teclado", Price: 335},
	// }


	//  db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	 // db.AutoMigrate(&Product{}, &Category{},)
	// db.Create(&products)
	// select
	// db.First(&products, 6)
	// db.First(&products, "name = ?", "Mouse")

	// var product []Product
	// db.Delete(&product, 1)
	// db.Find(&product)
	// for _, product := range product {
	// 	fmt.Println(product)
	// }

	//fmt.Println(products)


	// //create category
	// 	category := Category{
	// 		Name: "Eletronicos",
	// 	}
	// 	db.Create(&category)

		// //create product
		// product := Product{
    //   Name: "Monitor",
    //   Price: 2200,
    //   CategoryID: 1,
    // }
    // db.Create(&product)

	//Create a product
	// db.Create(&SerialNumber{
	// 	Number: "1234564",
	// 	ProductID: 1,
	// })

	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	// if err!= nil {
  //   panic(err)
  // }

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
  //   fmt.Println(product.Name)
  // }

	// Many to many
	// db.AutoMigrate(&Product{}, &Category{})
	
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)
	// category2 := Category{Name: "Eletronicos"}
	// db.Create(&category2)

	// // create product
	// db.Create(&Product{
	// 	Name: "Panela",
	// 	Price: 99.0,
	// 	Categories: []Category{category, category2},
	// })

	var categories []Category
  err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err!= nil {
    panic(err)
  }
	for _, category := range categories {
		for _, product := range category.Products {
      fmt.Printf("Produto: %v - Categorias: %v.\n" , category.Name, product.Name)
    }
  }
}