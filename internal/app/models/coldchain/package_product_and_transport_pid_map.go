package coldchain

type PackageProductAndTransportPIDMap struct {
	PackageProductNumber string `gorm:"primarykey" json:"package_product_number"`
	TransportPID         string `gorm:"not null; type:varchar(256)" json:"transport_pid"`
}
