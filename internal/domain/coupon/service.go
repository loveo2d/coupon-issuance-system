package coupon

import (
	"context"

	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type CouponService struct {
	db db.DB
}

func NewCouponService(db db.DB) *CouponService {
	return &CouponService{db: db}
}

func (s *CouponService) IssueCoupon(ctx context.Context, campaignId int32) (*Coupon, error) {
	return nil, nil
}
