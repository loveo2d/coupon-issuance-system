package campaign_create

type Input struct {
}

type Output struct {
}

type CreateCampaignUC struct {
}

func New() *CreateCampaignUC {
	return &CreateCampaignUC{}
}

func (uc *CreateCampaignUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
