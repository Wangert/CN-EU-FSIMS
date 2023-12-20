package response

import "CN-EU-FSIMS/internal/app/models/pasture"

type ResFeedingRecords struct {
	FeedingBatches []pasture.FeedingBatchInfo `json:"feeding_batches"`
	Count          int64                      `json:"count"`
}

type ResEndFeeding struct {
	Checkcode string   `json:"checkcode"`
	CowsNum   []string `json:"cows_num"`
	Count     int64    `json:"count"`
}
