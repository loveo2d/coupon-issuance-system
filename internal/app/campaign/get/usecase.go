package campaign_get

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Input struct {
	CampaignID int32
}

type Output struct {
	CampaignId    int32
	Title         string
	CouponRemains int32
	BeginAt       time.Time
}

type GetCampaignUC struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *GetCampaignUC {
	return &GetCampaignUC{
		db: db,
	}
}

func (uc *GetCampaignUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
