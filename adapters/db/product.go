package db

import (
	"database/sql"

	"github.com/danvieira97/arq-hexagonal-projeto/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	DB *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.DB.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.DB.Prepare(`INSERT INTO products(id, name, price, status) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return product, nil
}

func (p *ProductDB) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.DB.Exec("UPDATE products SET name  = ?, price = ?, status = ? WHERE id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	defer p.DB.Close()
	return product, nil
}

func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.DB.QueryRow("SELECT id from products where id = ?", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}
