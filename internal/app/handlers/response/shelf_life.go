package response

import "CN-EU-FSIMS/internal/app/models"

type ResSpectralData struct {
	Spectras []models.Spectra `json:"spectras"`
	Count    int64            `json:"count"`
}
