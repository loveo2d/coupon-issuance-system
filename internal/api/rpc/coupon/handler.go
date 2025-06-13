package coupon_handler

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"
	couponpb "github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/coupon"
	"github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/coupon/couponconnect"
	coupon_issue "github.com/loveo2d/CouponIssuanceSystem/internal/app/coupon/issue"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/campaign"
	"github.com/loveo2d/CouponIssuanceSystem/internal/domain/coupon"
	"github.com/loveo2d/CouponIssuanceSystem/internal/infra/db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type couponServer struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) (string, http.Handler) {
	server := &couponServer{db: db}
	return couponconnect.NewCouponServiceHandler(server)
}

func (s *couponServer) IssueCoupon(ctx context.Context, req *connect.Request[couponpb.IssueCouponRequest]) (*connect.Response[couponpb.IssueCouponResponse], error) {
	uc := coupon_issue.New(
		s.db,
		func(db db.DB) campaign.Repository {
			return campaign.NewCampaignRepository(db)
		},
		func(db db.DB) coupon.Service {
			return coupon.NewCouponService(db)
		},
	)
	output, err := uc.Execute(ctx, coupon_issue.Input{
		CampaignId: req.Msg.CampaignId,
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&couponpb.IssueCouponResponse{
		CouponId:   output.CouponId,
		CampaignId: output.CampaignId,
		CouponCode: output.CouponCode,
		IssuedAt:   timestamppb.New(output.IssuedAt),
	})
	return res, nil
}
