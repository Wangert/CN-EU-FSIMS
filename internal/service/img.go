package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/imgquery"
	"context"
)

func GetAllImgs() ([]models.UploadImg, int64, error) {
	q := imgquery.UploadImg
	ims, err := q.WithContext(context.Background()).Find()
	if err != nil {
		return []models.UploadImg{}, 0, err
	}

	count, err := q.WithContext(context.Background()).Count()

	imgs := make([]models.UploadImg, count)
	for i, im := range ims {
		imgs[i] = models.UploadImg(*im)
	}

	return imgs, count, nil
}
