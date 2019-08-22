package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) get(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM schema_user.products WHERE id=$1", p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) update(db *sql.DB) error {
	_, err := db.Exec("UPDATE schema_user.products SET name=$1, price=$2 WHERE id=$3",
		p.Name, p.Price, p.ID)
	return err
}

func (p *product) delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM schema_user.products WHERE id=$1", p.ID)
	return err
}

func (p *product) create(db *sql.DB) error {
	return db.QueryRow("INSERT INTO schema_user.products(name, price) VALUES($1, $2) RETURNING id", p.Name, p.Price).Scan(&p.ID)
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query("SELECT id, name, price FROM schema_user.products LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products:=[]product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}


