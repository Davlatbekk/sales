package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, req *models.CreateProduct) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		SELECT
			product_id
		FROM products
		ORDER BY product_id  DESC
		LIMIT 1
	`

	err := r.db.QueryRow(ctx, query).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO products(
			product_id, 
			product_name, 
			brand_id,
			category_id,
			model_year,
			list_price
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = r.db.Exec(ctx, query,
		id+1,
		req.ProductName,
		req.BrandId,
		req.CategoryId,
		req.ModelYear,
		req.ListPrice,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query   string
		product models.Product
	)

	query = `
		SELECT
			product_id, 
			product_name, 
			brand_id,
			category_id,
			model_year,
			list_price
		FROM products
		WHERE product_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.ProductId).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.BrandId,
		&product.CategoryId,
		&product.ModelYear,
		&product.ListPrice,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {

	resp = &models.GetListProductResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			product_id, 
			product_name, 
			brand_id,
			category_id,
			model_year,
			list_price
		FROM products
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
		var product models.Product
		err = rows.Scan(
			&resp.Count,
			&product.ProductId,
			&product.ProductName,
			&product.BrandId,
			&product.CategoryId,
			&product.ModelYear,
			&product.ListPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product)
	}

	return resp, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		products
		SET
			product_id = :product_id, 
			product_name = :product_name, 
			brand_id = :brand_id,
			category_id = :category_id,
			model_year = :model_year,
			list_price = :list_price
		WHERE product_id = :product_id
	`

	params = map[string]interface{}{
		"product_id":   req.ProductId,
		"product_name": req.ProductName,
		"brand_id":     req.BrandId,
		"category_id":  req.CategoryId,
		"model_year":   req.ModelYear,
		"list_price":   req.ListPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM products
		WHERE product_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.ProductId)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
