package coupon

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type Service interface {
	IssueCoupon(ctx context.Context, campaignId int32) (*Coupon, error)
}
type CouponService struct {
	db db.DB
}

func NewCouponService(db db.DB) *CouponService {
	return &CouponService{db: db}
}

func (s *CouponService) IssueCoupon(ctx context.Context, campaignId int32) (*Coupon, error) {
	couponRepo := NewCouponRepository(s.db)

	couponModel, errCoupon := couponRepo.Create(ctx, &Coupon{
		Code:       generateCouponCode(),
		CampaignId: campaignId,
		IssuedAt:   time.Now(),
	})
	if errCoupon != nil {
		return nil, errCoupon
	}
	return couponModel, nil
}

func generateCouponCode() string {
	const (
		startHangul = 0xAC00 // '가'
		endHangul   = 0xD7A3 // '힣'
	)
	randomHangulCodePoint := rune(rand.Intn(endHangul-startHangul+1) + startHangul)
	randomNumberPart := rand.Intn(1_000_000_000)
	return fmt.Sprintf("%c%09d", randomHangulCodePoint, randomNumberPart)
}
