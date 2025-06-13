package campaign

import (
	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type CampaignRepository struct {
	db db.DB
}

func NewCampaignRepository(db db.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (r *CampaignRepository) Create(campaign *Campaign) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) Get(campaignID int32) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) GetWithLock(campaignID int32) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) Update(campaign *Campaign) (*Campaign, error) {
	return nil, nil
}

func (r *CampaignRepository) Delete(campaignID int32) error {
	return nil
}

type CampaignScheduleRepository struct {
	db db.DB
}

func NewCampaignScheduleRepository(db db.DB) *CampaignScheduleRepository {
	return &CampaignScheduleRepository{db: db}
}

func (r *CampaignScheduleRepository) Create(campaignSchedule *CampaignSchedule) (*CampaignSchedule, error) {
	return nil, nil
}

func (r *CampaignScheduleRepository) Get(campaignScheduleID int32) (*CampaignSchedule, error) {
	return nil, nil
}
