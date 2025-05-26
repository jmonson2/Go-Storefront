package db

import (
	"database/sql"
	"fmt"
	"log"
	"storefront/items"
)

func GetAllInventory() ([]items.Inventory, error) {
	db := GetConnection()

	rows, err := db.Query("SELECT sku, name, stock, price FROM inventory")

	if err != nil {
		log.Println("Error selecting all inventory: ", err)
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
			log.Println("Error selecting all inventory: ", err)
		}
	}(rows)

	return scanInventoryRows(rows)
}

func GetInventoryBySKU(sku string) ([]items.Inventory, error) {
	db := GetConnection()

	rows, err := db.Query("SELECT sku, name, stock, price FROM inventory WHERE sku = ?", sku)

	if err != nil {
		log.Println("Error selecting inventory by sku: ", err)
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

	return scanInventoryRows(rows)
}

func InsertInventory(item items.Inventory) error {
	db := GetConnection()

	_, err := db.Exec("insert into inventory(sku, name, price, stock) values(?,?,?, 0)", item.SKU, item.Name, item.SalesPrice)
	if err != nil {
		log.Println("Error inserting new inventory record: ", err)
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

func scanInventoryRows(rows *sql.Rows) ([]items.Inventory, error) {
	var inventory []items.Inventory

	for rows.Next() {
		var item items.Inventory
		if err := rows.Scan(&item.SKU, &item.Name, &item.Stock, &item.SalesPrice); err != nil {
			return nil, fmt.Errorf("Failed to scan inventory row: %w", err)
		}
		inventory = append(inventory, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during row iteration: %w", err)
	}

	return inventory, nil
}
