package postgresql

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
	// customer storage.CustomerRepoI
	// user     storage.UserRepoI
	brand    storage.BrandRepoI
	product  storage.ProductRepoI
	category storage.CategoryRepoI
	stock storage.StockRepoI
	// order    storage.OrderRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pgpool,
		product:  NewProductRepo(pgpool),
		category: NewCategoryRepo(pgpool),
		brand:    NewBrandRepo(pgpool),
		stock:    NewStockRepo(pgpool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

// func (s *Store) Customer() storage.CustomerRepoI {
// 	if s.customer == nil {
// 		s.customer = NewCustomerRepo(s.db)
// 	}

// 	return s.customer
// }

// func (s *Store) User() storage.UserRepoI {
// 	if s.user == nil {
// 		s.user = NewUserRepo(s.db)
// 	}

// 	return s.user
// }

func (s *Store) Brand() storage.BrandRepoI {
	if s.brand == nil {
		s.brand = NewBrandRepo(s.db)
	}

	return s.brand
}
func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Stock() storage.StockRepoI {
	if s.stock == nil {
		s.stock = NewStockRepo(s.db)
	}

	return s.stock
}

// func (s *Store) Order() storage.OrderRepoI {
// 	if s.order == nil {
// 		s.order = NewOrderRepo(s.db)
// 	}

// 	return s.order
// }
