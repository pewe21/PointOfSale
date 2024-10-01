package dto

type CreateProductRequest struct {
	Name       string `json:"name"`
	SupplierId string `json:"supplier_id"`
	SKU        string `json:"sku"`
}

type UpdateProductRequest struct {
	Name       string `json:"name"`
	SupplierId string `json:"supplier_id"`
	SKU        string `json:"sku"`
}

type ProductData struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Supplier PSupplier `json:"supplier"`
}

type PSupplier struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
