package coupon

import "time"

type Coupon struct {
	ID         int64
	Code       string
	CampaignId int32
	IssuedAt   time.Time
}
