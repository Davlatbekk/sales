package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type stockRepo struct {
	db *pgxpool.Pool
}

func NewStockRepo(db *pgxpool.Pool) *stockRepo {
	return &stockRepo{
		db: db,
	}
}

func (r *stockRepo) CreateStock(ctx context.Context, req *models.CreateStock) (int, error) {
	var (
		query string
		id    int
	)

	// get last id
	query = `
	SELECT
		stock_id
	FROM Stocks
	ORDER BY stock_id  DESC
	LIMIT 1
`
	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO Stocks(
			stock_id, 
			product_id,
			quantity

	)
	VALUES ($1, $2, $3)`

	_, err = r.db.Exec(ctx, query,
		id+1,
		req.Quantity,
	)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (r *stockRepo) GetByIDStock(ctx context.Context, req *models.StockPrimaryKey) (*models.Stock, error) {

	var (
		query string
		Stock models.Stock
	)

	query = `
		SELECT
			stock_id, 
			product_id,
			quantity
		FROM stocks
		WHERE stock_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.StockId).Scan(
		&Stock.StockId,
		&Stock.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return &Stock, nil
}

func (r *stockRepo) GetListStock(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error) {

	resp = &models.GetListStockResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			stock_id, 
			product_id,
			quantity
		FROM Stocks
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Stock models.Stock
		err = rows.Scan(
			&resp.Count,
			&Stock.StockId,
			&Stock.Quantity,
		)
		if err != nil {
			return nil, err
		}

		resp.Stocks = append(resp.Stocks, &Stock)
	}

	return resp, nil
}

func (r *stockRepo) UpdateStock(ctx context.Context, req *models.UpdateStock) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		stocks
		SET
			stock_id = :stock_id, 
			product_id = :product_id
			quantity = :quantity
		WHERE stock_id = :stock_id
	`

	params = map[string]interface{}{
		"stock_id":   req.StockId,
		"product_id": req.ProductId,
		"quantity": req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *stockRepo) DeleteStock(ctx context.Context, req *models.StockPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM stocks
		WHERE stock_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.StockId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
