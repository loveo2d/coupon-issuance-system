package campaign_get

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/campaign"
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

func (uc *GetCampaignUC) Execute(input Input) (output *Output, err error) {
	tx, err := uc.db.BeginTx(context.Background(), pgx.TxOptions{AccessMode: pgx.ReadOnly})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	campaignRepo := campaign.NewCampaignRepository(tx)

	campaignModel, errCampaign := campaignRepo.Get(input.CampaignId)
	if errCampaign != nil {
		return nil, errCampaign
	}

	output = &Output{
		CampaignId:    campaignModel.ID,
		Title:         campaignModel.Title,
		CouponRemains: &campaignModel.CouponRemains,
		BeginAt:       campaignModel.BeginAt,
	}
	return output, nil
}
