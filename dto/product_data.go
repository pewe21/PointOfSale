package dto

type ProductxDto struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	SKU      string    `json:"sku"`
	Supplier Supplierx `json:"supplier"`
	Brand    Brandx    `json:"brand"`
}

// Model untuk Brand
type Supplierx struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Brandx struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required"`
	SupplierId string `json:"supplier_id" validate:"required,uuid"`
	BrandId    string `json:"brand_id" validate:"required,uuid"`
	SKU        string `json:"sku" validate:"required"`
}

type UpdateProductRequest struct {
	Name       string `json:"name" validate:"required"`
	SupplierId string `json:"supplier_id" validate:"required,uuid"`
	BrandId    string `json:"brand_id" validate:"required,uuid"`
	SKU        string `json:"sku" validate:"required"`
}
