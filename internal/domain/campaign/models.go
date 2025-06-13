package campaign

import "time"

type Campaign struct {
	ID            int32
	Title         string
	BeginAt       time.Time
	CouponRemains int32
}

type CampaignSchedule struct {
	ID         int32
	CampaignId int32
	Status     string
	BeginAt    time.Time
}
