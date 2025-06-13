package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const databaseURL = "postgresql://postgres:postgres@localhost:5432/coupon_issuance_system"

func NewDB() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), databaseURL)
}
