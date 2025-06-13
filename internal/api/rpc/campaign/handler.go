package campaign_handler

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	campaignpb "github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/campaign"
	"github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/campaign/campaignconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type campaignServer struct{}

func New() (string, http.Handler) {
	server := &campaignServer{}
	return campaignconnect.NewCampaignServiceHandler(server)
}

func (s *campaignServer) CreateCampaign(ctx context.Context, req *connect.Request[campaignpb.CreateCampaignRequest]) (*connect.Response[campaignpb.CreateCampaignResponse], error) {
	res := connect.NewResponse(&campaignpb.CreateCampaignResponse{
		CampaignId:    1,
		Title:         "캠페인 테스트",
		CouponRemains: 10,
		BeginAt:       timestamppb.Now(),
	})
	return res, nil
}

func (s *campaignServer) GetCampaign(ctx context.Context, req *connect.Request[campaignpb.GetCampaignRequest]) (*connect.Response[campaignpb.GetCampaignResponse], error) {
	var couponRemains int32 = 0
	res := connect.NewResponse(&campaignpb.GetCampaignResponse{
		CampaignId:    1,
		Title:         "캠페인 테스트",
		CouponRemains: &couponRemains,
		BeginAt:       timestamppb.Now(),
	})
	return res, nil
}
