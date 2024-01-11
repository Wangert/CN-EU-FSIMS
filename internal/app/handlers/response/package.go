package response

import (
	"CN-EU-FSIMS/internal/app/models/pack"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/warehouse"
)

type ResPackageReceiveRecords struct {
	Records []warehouse.PackageReceiveRecordInfo `json:"records"`
	Count   int64                                `json:"count"`
}

type ResPackageBatches struct {
	Records []pack.PackageBatchInfo `json:"records"`
	Count   int64                   `json:"count"`
}

type ResPackageWarehouseRecords struct {
	Records []warehouse.PackWarehouseInfo `json:"records"`
	Count   int64                         `json:"count"`
}

type ResPackageProductsRecords struct {
	Records []product.PackageProductInfo `json:"records"`
	Count   int64                        `json:"count"`
}

type ResEndPackage struct {
	Checkcode   string   `json:"checkcode"`
	ProductsNum []string `json:"products_num"`
	Count       int64    `json:"count"`
}
