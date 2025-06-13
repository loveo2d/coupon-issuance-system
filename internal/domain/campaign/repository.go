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
	query := "INSERT INTO campaigns (title, coupon_remains, begin_at) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(ctx, query, campaign.Title, campaign.CouponRemains, campaign.BeginAt).Scan(&campaign.ID)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (r *CampaignRepository) Get(ctx context.Context, campaignID int32) (*Campaign, error) {
	campaign := &Campaign{}
	query := "SELECT id, title, coupon_remains, begin_at FROM campaigns WHERE id = $1"
	err := r.db.QueryRow(ctx, query, campaignID).Scan(&campaign.ID, &campaign.Title, &campaign.CouponRemains, &campaign.BeginAt)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (r *CampaignRepository) GetWithLock(ctx context.Context, campaignID int32) (*Campaign, error) {
	campaign := &Campaign{}
	query := "SELECT id, title, coupon_remains, begin_at FROM campaigns WHERE id = $1 FOR UPDATE"
	err := r.db.QueryRow(ctx, query, campaignID).Scan(&campaign.ID, &campaign.Title, &campaign.CouponRemains, &campaign.BeginAt)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (r *CampaignRepository) Update(ctx context.Context, campaign *Campaign) (*Campaign, error) {
	query := "UPDATE campaigns SET title = $1, coupon_remains = $2, begin_at = $3 WHERE id = $4"
	_, err := r.db.Exec(ctx, query, campaign.Title, campaign.CouponRemains, campaign.BeginAt, campaign.ID)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (r *CampaignRepository) Delete(ctx context.Context, campaignID int32) error {
	query := "DELETE FROM campaigns WHERE id = $1"
	_, err := r.db.Exec(ctx, query, campaignID)
	if err != nil {
		return err
	}
	return nil
}

type CampaignScheduleRepository struct {
	db db.DB
}

func NewCampaignScheduleRepository(db db.DB) *CampaignScheduleRepository {
	return &CampaignScheduleRepository{db: db}
}

func (r *CampaignScheduleRepository) Create(ctx context.Context, campaignSchedule *CampaignSchedule) (*CampaignSchedule, error) {
	query := "INSERT INTO campaign_schedules (campaign_id, status, begin_at) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(ctx, query, campaignSchedule.CampaignId, campaignSchedule.Status, campaignSchedule.BeginAt).Scan(&campaignSchedule.ID)
	if err != nil {
		return nil, err
	}
	return campaignSchedule, nil
}

func (r *CampaignScheduleRepository) Get(ctx context.Context, campaignScheduleID int32) (*CampaignSchedule, error) {
	campaignSchedule := &CampaignSchedule{}
	query := "SELECT id, campaign_id, status, begin_at FROM campaign_schedules WHERE id = $1"
	err := r.db.QueryRow(ctx, query, campaignScheduleID).Scan(&campaignSchedule.ID, &campaignSchedule.CampaignId, &campaignSchedule.Status, &campaignSchedule.BeginAt)
	if err != nil {
		return nil, err
	}
	return campaignSchedule, nil
}
