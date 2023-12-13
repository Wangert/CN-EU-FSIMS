package response

import "time"

type ResLogs struct {
	TimeStamp time.Time `json:"timestamp"`
	UUID      string    `json:"uuid"`
	Account   string    `json:"account"`
	Type      int       `json:"type"`
	Action    string    `json:"action"`
}
