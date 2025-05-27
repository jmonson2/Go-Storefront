# Go Storefront 

This project is a Go-based API server built using the Gin web framework. It provides basic functionality of a storefront.

## Features

-   **Product Management:**
    -   List all products
    -   Retrieve products by SKU, ID, or name
    -   Add new products
    -   Sell products
-   **Inventory Management:**
    -   List all inventory items
    -   Retrieve inventory by SKU
    -   Add items to inventory

## Prerequisites

-   Go
-   Gin Web Framework (`github.com/gin-gonic/gin`)
-   SQLite

## API Endpoints

The API server runs on `localhost:8080` by default.

### Product Endpoints

-   **GET `/products`**: Get a list of all products.
    -   Handler: `GetAllProducts`
-   **GET `/products/sku`**: Get products filtered by SKU.
    -   Query Parameters: (e.g., `?sku=SKU123`)
    -   Handler: `GetProductsBySKU`
-   **GET `/products/id`**: Get a product by its ID.
    -   Query Parameters: (e.g., `?id=1`)
    -   Handler: `GetProductByID`
-   **GET `/products/name`**: Get products filtered by name.
    -   Query Parameters: (e.g., `?name=ProductName`)
    -   Handler: `GetProductsByName`
-   **POST `/products/add`**: Add a new product.
    -   Request Body:
    ```json
    { 
        "id":  int,
        "sku": string 
    }
-   Handler: `AddProduct`
-   **POST `/products/sell`**: Record a product sale.
    -   Request Body:
    ```json 
    { 
        "id": string
    } 
-   Handler: `SellProduct`

### Inventory Endpoints

-   **GET `/inventory`**: Get a list of all inventory items.
    -   Handler: `GetAllInventory`
-   **GET `/inventory/sku`**: Get inventory items filtered by SKU.
    -   Query Parameters: (e.g., `?sku=SKU12345`)
    -   Handler: `GetInventoryBySKU`
-   **POST `/inventory/add`**: Add items to the inventory.
    -   Request Body:
    ```json
    {
        "sku":   string,
        "name":  string,
        "price": float,
        "stock": int 
    }
    
-   Handler: `AddInventory`
