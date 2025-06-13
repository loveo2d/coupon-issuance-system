package campaign_create

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Input struct {
	Title         string
	CouponRemains int32
	BeginAt       time.Time
}

type Output struct {
	CampaignId    int32
	Title         string
	CouponRemains int32
	BeginAt       time.Time
}

type CreateCampaignUC struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *CreateCampaignUC {
	return &CreateCampaignUC{db: db}
}

func (uc *CreateCampaignUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
