package domain

type Product struct {
	Id         string   `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	Name       string   `json:"name" db:"name"`
	TypeId     string   `json:"type_id" db:"type_id"`
	Type       Type     `json:"type"`
	SupplierId string   `json:"supplier_id" db:"supplier_id"`
	Supplier   Supplier `json:"supplier"`
}
