package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	products, err := selectAllProducts(db)
	if err!= nil {
    panic(err)
  }
	// products[0].Name= "Macbook"
	// products[0].Price = 7585.65

	// err = updateProduct(db, &products[0])
	// if err!= nil {
  //   panic(err)
  // }
	err = deleteProduct(db, products[len(products) -1 ].ID)
	if err!= nil {
    panic(err)
  }
	for _, product := range products {
    fmt.Printf("O produto: %v, possui o valor de R$ %.2f\n", product.Name, product.Price)
  }
	// product, err := selectProductForId(db, "59667c71-8ef6-4cc4-a4de-b2fa8149cfe1")
	// if err!= nil {
  //   panic(err)
  // }
	// fmt.Printf("O produto: %v, possui o valor de R$ %.2f\n", product.Name, product.Price)
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare(`INSERT INTO products (id, name, price) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func selectAllProducts(db *sql.DB) ([]Product, error){
	rows, err := db.Query(`SELECT id, name, price FROM products`)
  if err!= nil {
    return nil, err 
  }
  defer rows.Close()

	var products []Product
  for rows.Next() {
    var p Product

    err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err!= nil {
			return nil, err
		}

		products = append(products, p)
  }
	return products, nil
}

func selectProductForId(db *sql.DB, id string) (Product, error) {
	stmt, err := db.Prepare(`SELECT id, name, price FROM products WHERE id =?`)
  if err!= nil {
    panic(err)
  }
  defer stmt.Close()

	var product Product
  err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err!= nil {
    return product , err
  }
  
  return product, nil
}

func updateProduct(db  *sql.DB, product *Product) error {
	stmt, err := db.Prepare(`UPDATE products SET name = ? , price = ? WHERE id = ?`)
	if err!= nil {
    return err
  }
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err!= nil {
    return err
  }
	return nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare(`DELETE FROM products WHERE id =?`)
  if err!= nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(id)
  if err!= nil {
    return err
  }
  return nil
}