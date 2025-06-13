package coupon_issue

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Input struct {
	CampaignId int32
}

type Output struct {
	CouponId   int64
	CouponCode string
	CampaignId int32
	IssuedAt   time.Time
}

type IssueCouponUC struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *IssueCouponUC {
	return &IssueCouponUC{db: db}
}

func (uc *IssueCouponUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
