package models

type Stock struct {
	StockId   int    `json:"stock_id"`
	ProductId int    `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type StockPrimaryKey struct {
	StockId int `json:"stock_id"`
}

// type StockCategory struct {
// 	StockId int `json:"stock_id"`
// }

type CreateStock struct {
	Quantity  int `json:"quantity"`
	ProductId int    `json:"product_id"`
}

type UpdateStock struct {
	StockId   int    `json:"stock_id"`
	Quantity  int `json:"quantity"`
	ProductId int    `json:"product_id"`
}

type GetListStockRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStockResponse struct {
	Count  int `json:"count"`
	Stocks []*Stock
}
