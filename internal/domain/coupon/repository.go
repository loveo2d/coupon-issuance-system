package coupon

import (
	"context"

	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type CouponRepository struct {
	db db.DB
}

func NewCouponRepository(db db.DB) *CouponRepository {
	return &CouponRepository{db: db}
}

func (r *CouponRepository) Create(ctx context.Context, coupon *Coupon) (*Coupon, error) {
	return nil, nil
}

func (r *CouponRepository) Get(ctx context.Context, couponID int32) (*Coupon, error) {
	return nil, nil
}
