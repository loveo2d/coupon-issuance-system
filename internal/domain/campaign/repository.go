package campaign

import (
	"context"

	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type CampaignRepository struct {
	db db.DB
}

func NewCampaignRepository(db db.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (r *CampaignRepository) Create(ctx context.Context, campaign *Campaign) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) Get(ctx context.Context, campaignID int32) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) GetWithLock(ctx context.Context, campaignID int32) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) Update(ctx context.Context, campaign *Campaign) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) Delete(ctx context.Context, campaignID int32) error {
	return nil
}

type CampaignScheduleRepository struct {
	db db.DB
}

func NewCampaignScheduleRepository(db db.DB) *CampaignScheduleRepository {
	return &CampaignScheduleRepository{db: db}
}

func (r *CampaignScheduleRepository) Create(ctx context.Context, campaignSchedule *CampaignSchedule) (*CampaignSchedule, error) {
	return nil, nil
}

func (r *CampaignScheduleRepository) Get(ctx context.Context, campaignScheduleID int32) (*CampaignSchedule, error) {
	return nil, nil
}
