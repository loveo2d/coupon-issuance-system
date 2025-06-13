package coupon

import (
	"github.com/jackc/pgx/v5"
)

type CouponRepository struct {
	tx pgx.Tx
}

func NewCouponRepository(tx pgx.Tx) *CouponRepository {
	return &CouponRepository{tx: tx}
}

func (r *CouponRepository) Create(coupon *Coupon) (*Coupon, error) {
	return nil, nil
}

func (r *CouponRepository) Get(couponID int32) (*Coupon, error) {
	return nil, nil
}
