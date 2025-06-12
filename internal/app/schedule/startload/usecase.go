package schedule_startload

type Input struct {
	CampaignID int32 `json:"campaign_id"`
}

type Output struct {
}

type StartLoadUC struct {
}

func New() *StartLoadUC {
	return &StartLoadUC{}
}

func (uc *StartLoadUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
