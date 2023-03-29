package models

type Product struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	BrandId     int     `json:"brand_id"`
	CategoryId  int     `json:"category_id"`
	ModelYear   int     `json:"model_year"`
	ListPrice   float64 `json:"list_price"`
}
type ProductPrimaryKey struct {
	ProductId int `json:"product_id"`
}

type CreateProduct struct {
	ProductName string  `json:"product_name"`
	BrandId     int     `json:"brand_id"`
	CategoryId  int     `json:"category_id"`
	ModelYear   int     `json:"model_year"`
	ListPrice   float64 `json:"list_price"`
}

type UpdateProduct struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	BrandId     int     `json:"brand_id"`
	CategoryId  int     `json:"category_id"`
	ModelYear   int     `json:"model_year"`
	ListPrice   float64 `json:"list_price"`
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"authors"`
}
