package items

type Inventory struct {
	Stock      int     `json:"stock"`
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	Rentable   bool    `json:"rentable"`
	SalesPrice float64 `json:"salesPrice"`
	RentPrice  float64 `json:"rentPrice"`
}

func NewInventory(stock int, name string, sku string, rentable bool, salesprice float64, rentprice float64) *Inventory {
	inventory := Inventory{}
	inventory.Stock = stock
	inventory.Name = name
	inventory.SKU = sku
	inventory.Rentable = rentable
	inventory.SalesPrice = salesprice
	inventory.RentPrice = rentprice
	return &inventory
}
