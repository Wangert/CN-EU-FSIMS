package response

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/product"
)

type ResFoodchains struct {
	Foodchains     []models.Foodchain `json:"foodchains"`
	TotalCount     int64              `json:"total_count"`
	CompletedCount int64              `json:"completed_count"`
}

type ResPidInfo struct {
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Address     string `json:"address"`
	HouseNumber string `json:"house_number"`
}

type ResProductsInfo struct {
	CowsInfo              []product.CowInfo              `json:"cows_info"`
	SlaughterProductsInfo []product.SlaughterProductInfo `json:"slaughter_products_info"`
	PackageProductsInfo   []product.PackageProductInfo   `json:"package_products_info"`
	ColdchainInfo         *ColdchainInfo                 `json:"coldchain_info"`
}

type ColdchainInfo struct {
	TVNumber   string `json:"tv_number"`
	MallNumber string `json:"mall_number"`
	MallName   string `json:"mall_name"`
}
