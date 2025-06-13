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
	query := "INSERT INTO coupons (code, campaign_id, issued_at) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(ctx, query, coupon.Code, coupon.CampaignId, coupon.IssuedAt).Scan(&coupon.ID)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (r *CouponRepository) Get(ctx context.Context, couponID int32) (*Coupon, error) {
	coupon := &Coupon{}
	query := "SELECT id, code, campaign_id, issued_at FROM coupons WHERE id = $1"
	err := r.db.QueryRow(ctx, query, couponID).Scan(&coupon.ID, &coupon.Code, &coupon.CampaignId, &coupon.IssuedAt)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}
