package campaign_create

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/campaign"
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

func (uc *CreateCampaignUC) Execute(ctx context.Context, input Input) (output *Output, err error /* defer 롤백 처리 시 이 변수(err)를 바라본다 */) {
	tx, err := uc.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		// 함수가 끝나고 반환 변수 err가 nil이 아니면 롤백
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	// tx 변수의 경우 interface 타입으로 이미 참조의 성격을 가지고 있어 참조 변환이 필요하지 않음
	campaignRepo := campaign.NewCampaignRepository(tx)
	campaignScheduleRepo := campaign.NewCampaignScheduleRepository(tx)

	campaignModel, errCampaign := campaignRepo.Create(ctx, &campaign.Campaign{
		Title:         input.Title,
		CouponRemains: input.CouponRemains,
		BeginAt:       input.BeginAt,
	})

	// 반환 변수에 err를 명시했으므로 defer에서 감지 가능
	if errCampaign != nil {
		return nil, errCampaign
	}

	_, errCampaignSchedule := campaignScheduleRepo.Create(ctx, &campaign.CampaignSchedule{
		CampaignId: campaignModel.ID,
		Status:     "PENDING",
		BeginAt:    campaignModel.BeginAt,
	})

	if errCampaignSchedule != nil {
		return nil, errCampaignSchedule
	}

	if err := tx.Commit(context.Background()); err != nil {
		// 커밋 자체가 실패할 수도 있다 -> 롤백
		return nil, err
	}

	return &Output{
		CampaignId:    campaignModel.ID,
		Title:         campaignModel.Title,
		CouponRemains: campaignModel.CouponRemains,
		BeginAt:       campaignModel.BeginAt,
	}, nil
}
