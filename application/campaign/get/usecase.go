package campaign_get

type Input struct {
	CampaignID int64 `json:"campaign_id"`
}

type Output struct {
}

type GetCampaignUC struct {
}

func New() *GetCampaignUC {
	return &GetCampaignUC{}
}

func (uc *GetCampaignUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
