package response

import "CN-EU-FSIMS/internal/app/models/coldchain"

type ResTransportBatches struct {
	Records []coldchain.TransportBatchInfo `json:"records"`
	Count   int64                          `json:"count"`
}
