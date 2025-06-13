package campaign

import "github.com/jackc/pgx/v5"

type CampaignRepository struct {
	tx pgx.Tx
}

func NewCampaignRepository(tx pgx.Tx) *CampaignRepository {
	return &CampaignRepository{tx: tx}
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
	tx pgx.Tx
}

func NewCampaignScheduleRepository(tx pgx.Tx) *CampaignScheduleRepository {
	return &CampaignScheduleRepository{tx: tx}
}

func (r *CampaignScheduleRepository) Create(campaignSchedule *CampaignSchedule) (*CampaignSchedule, error) {
	return nil, nil
}

func (r *CampaignScheduleRepository) Get(campaignScheduleID int32) (*CampaignSchedule, error) {
	return nil, nil
}
