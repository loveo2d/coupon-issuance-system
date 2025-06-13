package campaign_get

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Input struct {
	CampaignId int32
}

type Output struct {
	CampaignId    int32
	Title         string
	CouponRemains *int32 // proto 정의에서 optional이어서 참조를 전달한다. 핸들러는 이것이 참조인지 값인지 신경을 쓰지 않는다.
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
