package coupon_issue

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/campaign"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/coupon"
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

func (uc *IssueCouponUC) Execute(ctx context.Context, input Input) (output *Output, err error) {
	tx, err := uc.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	campaignRepo := campaign.NewCampaignRepository(tx)

	// 비관적 락을 걸면서 캠페인 조회
	campaignModel, errCampaign := campaignRepo.GetWithLock(ctx, input.CampaignId)
	if errCampaign != nil {
		return nil, errCampaign
	}

	if campaignModel.BeginAt.After(time.Now()) {
		return nil, errors.New("캠페인이 아직 시작되지 않음")
	}

	if campaignModel.CouponRemains <= 0 {
		return nil, errors.New("잔여 쿠폰 없음")
	}

	campaignModel.CouponRemains--
	_, errCampaignUpdate := campaignRepo.Update(ctx, campaignModel)
	if errCampaignUpdate != nil {
		return nil, errCampaignUpdate
	}

	couponService := coupon.NewCouponService(tx)
	couponModel, errCoupon := couponService.IssueCoupon(ctx, input.CampaignId)
	if errCoupon != nil {
		return nil, errCoupon
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	output = &Output{
		CouponId:   couponModel.ID,
		CouponCode: couponModel.Code,
		CampaignId: couponModel.CampaignId,
		IssuedAt:   couponModel.IssuedAt,
	}
	return output, nil
}
