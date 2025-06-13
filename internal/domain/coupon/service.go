package coupon

import (
	"github.com/jackc/pgx/v5"
)

type CouponService struct {
	tx pgx.Tx
}

func NewCouponService(tx pgx.Tx) *CouponService {
	return &CouponService{tx: tx}
}

func (s *CouponService) IssueCoupon(campaignId int32) (*Coupon, error) {
	return nil, nil
}
