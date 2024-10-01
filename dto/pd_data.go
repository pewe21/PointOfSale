package dto

type ProductxDto struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Supplier Supplierx `json:"supplier"` // Brand sebagai objek
	//Category Categoryx `json:"category"` // Category sebagai objek
}

// Model untuk Brand
type Supplierx struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Model untuk Category
//type Categoryx struct {
//	ID   int    `json:"id"`
//	Name string `json:"name"`
//}
