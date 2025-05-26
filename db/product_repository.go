package db

import (
	"database/sql"
	"log"
	"storefront/items"
)

func GetAllProducts() ([]items.Product, error) {
	db := GetConnection()

	rows, err := db.Query("select id, sku from product")

	defer func(db *sql.DB) {
		if err != nil {
			log.Println("Error selecting all products: ", err)
		}
	}(db)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db: ", err)
		}
	}(db)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

	return scanProductRows(rows)
}

func GetProductBySKU(sku string) ([]items.Product, error) {
	db := GetConnection()

	rows, err := db.Query("select id, sku from product where sku = ?", sku)

	if err != nil {
		log.Println("Error selecting product by sku: ", err)
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection: ", err)
		}
	}(db)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

	return scanProductRows(rows)
}

func GetProductByID(id int) ([]items.Product, error) {
	db := GetConnection()

	rows, err := db.Query("select id, sku from product where id = ?", id)

	if err != nil {
		log.Println("Error selecting product by id: ", err)
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection: ", err)
		}
	}(db)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

	return scanProductRows(rows)
}

func GetProductsByName(name string) ([]items.Product, error) {
	db := GetConnection()

	rows, err := db.Query(`select p.id, p.sku from product p, inventory i where i.sku = p.sku and 
								 lower(i.name) like lower('%' || ? || '%')`, name)

	if err != nil {
		log.Println("Error selecting product by name", err)
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection", err)
		}
	}(db)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println("Error closing rows", err)
		}
	}(rows)

	return scanProductRows(rows)
}

func getProductSKU(id int) (string, error) {
	db := GetConnection()
	rows, err := db.Query("select id, sku from product where id = ?", id)
	if err != nil {
		log.Println("Error selecting product by id: ", err)
		return "", err
	}
	item, err := scanProductRows(rows)
	if err != nil {
		log.Println("Error scanning product by id: ", err)
		return "", err
	}
	if item == nil {
		return "", nil
	}
	return item[0].SKU, err
}

func InsertProduct(item items.Product) error {
	db := GetConnection()

	_, err := db.Exec("insert into product (sku) values (?)", item.SKU)

	if err != nil {
		log.Println("Error inserting new product record: ", err)
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection: ", err)
		}
	}(db)

	return nil
}

func SellProduct(item items.Product) error {
	db := GetConnection()
	var err error
	if item.SKU == "" {
		item.SKU, err = getProductSKU(item.ID)
	}

	if err != nil {
		log.Println("Error getting product SKU: ", err)
		return err
	}

	_, err = db.Exec("delete from product where id = ?", item.ID)

	if err != nil {
		log.Println("Error deleting product at point of sale: ", err)
		return err
	}

	_, err = db.Exec("insert into sold_products (id, sku) values (?,?)", item.ID, item.SKU)

	if err != nil {
		log.Println("Error inserting sold product record: ", err)
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection: ", err)
		}
	}(db)

	return nil
}

// Extracts the ID and SKU from the rows and returns them as a slice of products
func scanProductRows(rows *sql.Rows) ([]items.Product, error) {
	var products []items.Product

	for rows.Next() {
		var item items.Product
		if err := rows.Scan(&item.ID, &item.SKU); err != nil {
			log.Println("Error scanning product: ", err)
			return nil, err
		}
		products = append(products, item)
	}
	return products, nil
}
