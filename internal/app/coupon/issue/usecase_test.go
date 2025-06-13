package coupon_issue

import (
	"context"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/campaign"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/coupon"
	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
)

type mockCampaignRepo struct {
	mock.Mock
}

func (m *mockCampaignRepo) Create(ctx context.Context, c *campaign.Campaign) (*campaign.Campaign, error) {
	args := m.Called(ctx, c)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (m *mockCampaignRepo) Get(ctx context.Context, id int32) (*campaign.Campaign, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (m *mockCampaignRepo) GetWithLock(ctx context.Context, id int32) (*campaign.Campaign, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (m *mockCampaignRepo) Update(ctx context.Context, c *campaign.Campaign) (*campaign.Campaign, error) {
	args := m.Called(ctx, c)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (m *mockCampaignRepo) Delete(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type mockCouponService struct {
	mock.Mock
}

func (m *mockCouponService) IssueCoupon(ctx context.Context, campaignId int32) (*coupon.Coupon, error) {
	args := m.Called(ctx, campaignId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*coupon.Coupon), args.Error(1)
}

func Test_IssueCouponUC_Execute(t *testing.T) {
	ctx := context.Background()
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockPool.Close()

	t.Run("실패 - 쿠폰 재고 부족", func(t *testing.T) {
		mockRepo := new(mockCampaignRepo)
		mockCampaign := &campaign.Campaign{ID: 1, CouponRemains: 0, BeginAt: time.Now().Add(-time.Hour)}
		mockRepo.On("GetWithLock", ctx, int32(1)).Return(mockCampaign, nil).Once()
		fakeCampaignRepoFactory := func(q db.DB) campaign.Repository { return mockRepo }
		fakeCouponServiceFactory := func(q db.DB) coupon.Service { return nil }
		mockPool.ExpectBegin()
		mockPool.ExpectRollback()
		uc := New(mockPool, fakeCampaignRepoFactory, fakeCouponServiceFactory)
		_, err := uc.Execute(ctx, Input{CampaignId: 1})
		assert.Error(t, err)
		assert.Equal(t, "잔여 쿠폰 없음", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("실패 - 캠페인 시작일 이전", func(t *testing.T) {
		mockRepo := new(mockCampaignRepo)
		mockCampaign := &campaign.Campaign{ID: 1, CouponRemains: 10, BeginAt: time.Now().Add(time.Hour)}
		mockRepo.On("GetWithLock", ctx, int32(1)).Return(mockCampaign, nil).Once()
		fakeCampaignRepoFactory := func(q db.DB) campaign.Repository { return mockRepo }
		fakeCouponServiceFactory := func(q db.DB) coupon.Service { return nil }
		mockPool.ExpectBegin()
		mockPool.ExpectRollback()
		uc := New(mockPool, fakeCampaignRepoFactory, fakeCouponServiceFactory)
		_, err := uc.Execute(ctx, Input{CampaignId: 1})
		assert.Error(t, err)
		assert.Equal(t, "캠페인이 아직 시작되지 않음", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("성공", func(t *testing.T) {
		mockRepo := new(mockCampaignRepo)
		mockSvc := new(mockCouponService)
		mockCampaign := &campaign.Campaign{ID: 1, CouponRemains: 10, BeginAt: time.Now().Add(-time.Hour)}
		mockRepo.On("GetWithLock", ctx, int32(1)).Return(mockCampaign, nil).Once()
		updatedCampaign := &campaign.Campaign{ID: 1, CouponRemains: 9, BeginAt: mockCampaign.BeginAt}
		mockRepo.On("Update", ctx, mock.Anything).Return(updatedCampaign, nil).Once()
		issuedCoupon := &coupon.Coupon{ID: 100, Code: "가123456789", CampaignId: 1, IssuedAt: time.Now()}
		mockSvc.On("IssueCoupon", ctx, int32(1)).Return(issuedCoupon, nil).Once()
		fakeCampaignRepoFactory := func(q db.DB) campaign.Repository { return mockRepo }
		fakeCouponServiceFactory := func(q db.DB) coupon.Service { return mockSvc }
		mockPool.ExpectBegin()
		mockPool.ExpectCommit()
		uc := New(mockPool, fakeCampaignRepoFactory, fakeCouponServiceFactory)
		output, err := uc.Execute(ctx, Input{CampaignId: 1})
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, int64(100), output.CouponId)
		mockRepo.AssertExpectations(t)
		mockSvc.AssertExpectations(t)
	})

	assert.NoError(t, mockPool.ExpectationsWereMet())
}
