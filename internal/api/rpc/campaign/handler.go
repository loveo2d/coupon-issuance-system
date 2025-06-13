package campaign_handler

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"
	campaignpb "github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/campaign"
	"github.com/loveo2d/CouponIssuanceSystem/internal/api/proto/campaign/campaignconnect"
	campaign_create "github.com/loveo2d/CouponIssuanceSystem/internal/app/campaign/create"
	campaign_get "github.com/loveo2d/CouponIssuanceSystem/internal/app/campaign/get"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type campaignServer struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) (string, http.Handler) {
	server := &campaignServer{db: db}
	return campaignconnect.NewCampaignServiceHandler(server)
}

func (s *campaignServer) CreateCampaign(ctx context.Context, req *connect.Request[campaignpb.CreateCampaignRequest]) (*connect.Response[campaignpb.CreateCampaignResponse], error) {
	uc := campaign_create.New(s.db)

	output, err := uc.Execute(campaign_create.Input{
		Title:         req.Msg.Title,
		CouponRemains: req.Msg.CouponRemains,
		BeginAt:       req.Msg.BeginAt.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&campaignpb.CreateCampaignResponse{
		CampaignId:    output.CampaignId,
		Title:         output.Title,
		CouponRemains: output.CouponRemains,
		BeginAt:       timestamppb.New(output.BeginAt),
	})
	return res, nil
}

func (s *campaignServer) GetCampaign(ctx context.Context, req *connect.Request[campaignpb.GetCampaignRequest]) (*connect.Response[campaignpb.GetCampaignResponse], error) {
	uc := campaign_get.New(s.db)
	output, err := uc.Execute(campaign_get.Input{
		CampaignId: req.Msg.CampaignId,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&campaignpb.GetCampaignResponse{
		CampaignId:    output.CampaignId,
		Title:         output.Title,
		CouponRemains: output.CouponRemains,
		BeginAt:       timestamppb.New(output.BeginAt),
	})
	return res, nil
}
