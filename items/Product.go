package items

type Product struct {
	ID  int    `json:"id"`
	SKU string `json:"sku"`
}

func NewProduct(id int, sku string) *Product {
	product := Product{}
	product.ID = id
	product.SKU = sku
	return &product
}
