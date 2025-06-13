package coupon_handler

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	couponpb "github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/coupon"
	"github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/coupon/couponconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type couponServer struct{}

func New() (string, http.Handler) {
	server := &couponServer{}
	return couponconnect.NewCouponServiceHandler(server)
}

func (s *couponServer) IssueCoupon(ctx context.Context, req *connect.Request[couponpb.IssueCouponRequest]) (*connect.Response[couponpb.IssueCouponResponse], error) {
	res := connect.NewResponse(&couponpb.IssueCouponResponse{
		CouponId:   101,
		CampaignId: 1,
		CouponCode: "가123456789",
		IssuedAt:   timestamppb.Now(),
	})
	return res, nil
}
