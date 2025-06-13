package coupon

import (
	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type CouponRepository struct {
	db db.DB
}

func NewCouponRepository(db db.DB) *CouponRepository {
	return &CouponRepository{db: db}
}

func (r *CouponRepository) Create(coupon *Coupon) (*Coupon, error) {
	return nil, nil
}

func (r *CouponRepository) Get(couponID int32) (*Coupon, error) {
	return nil, nil
}
