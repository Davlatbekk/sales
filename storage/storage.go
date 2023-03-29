package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Category() CategoryRepoI
	Brand() BrandRepoI
	Stock() StockRepoI
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (int, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (int, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
}

type BrandRepoI interface {
	Create(ctx context.Context, req *models.CreateBrand) (int, error)
	GetByID(ctx context.Context, req *models.BrandPrimaryKey) (*models.Brand, error)
	GetList(ctx context.Context, req *models.GetListBrandRequest) (resp *models.GetListBrandResponse, err error)
	Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateBrand) (int64, error)
}

type StockRepoI interface {
	CreateStock(ctx context.Context, req *models.CreateStock) (int, error)
	GetByIDStock(ctx context.Context, req *models.StockPrimaryKey) (*models.Stock, error)
	GetListStock(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error)
	DeleteStock(ctx context.Context, req *models.StockPrimaryKey) (int64, error)
	UpdateStock(ctx context.Context, req *models.UpdateStock) (int64, error)
}
