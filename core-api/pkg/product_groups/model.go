package product_groups

type ProductGroup struct {
	ID   uint64  `json:"id"`
	Name *string `json:"name"`
	Note *string `json:"note"`
}
