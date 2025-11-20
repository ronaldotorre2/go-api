package repository

import (
	"api-go/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.Id, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}
	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT id, name, price FROM product WHERE id=$1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var productObj model.Product
	err = query.QueryRow(id).Scan(&productObj.Id, &productObj.Name, &productObj.Price)
	if err != nil {
		if(err == sql.ErrNoRows) {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &productObj, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int,error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0,err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0,err
	}

	query.Close()

	return id,nil
}