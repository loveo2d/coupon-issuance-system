package coupon_issue

type Input struct {
	CampaignID int64 `json:"campaign_id"`
}

type Output struct {
	CouponCode string `json:"coupon_code"`
}

type IssueCouponUC struct {
}

func New() *IssueCouponUC {
	return &IssueCouponUC{}
}

func (uc *IssueCouponUC) Execute(input Input) (*Output, error) {
	return &Output{}, nil
}
