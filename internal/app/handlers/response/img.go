package response

import "CN-EU-FSIMS/internal/app/models"

type ResImgs struct {
	Imgs  []models.UploadImg `json:"imgs"`
	Count int64              `json:"count"`
}

type ResLymphNodes struct {
	Imgs  []models.LymphNode `json:"imgs"`
	Count int64              `json:"count"`
}
