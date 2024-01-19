package slaughter

import (
	"time"

	"gorm.io/gorm"
)

type PreCoolShop struct {
	//预冷间
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"`  //记录时间
	PreCoolShop1 float64   `json:"pre_cool_shop_1"` //预冷间温度
	PreCoolShop2 float64   `json:"pre_cool_shop_2"` //预冷间湿度
	PreCoolShop3 float64   `json:"pre_cool_shop_3"` //副产物温度
}

type SlaShop struct {
	//屠宰车间
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"` //记录时间
	SlaShop1     float64   `json:"sla_shop_1"`     //紫外灭菌
	SlaShop2     float64   `json:"sla_shop_2"`     //臭氧
	SlaShop3     float64   `json:"sla_shop_3"`     //臭氧残留
	SlaShop4     float64   `json:"sla_shop_4"`     //湿度
	SlaShop5     float64   `json:"sla_shop_5"`     //温度
	SlaShop6     float64   `json:"sla_shop_6"`     //预冷间消毒记录
	SlaShop7     float64   `json:"sla_shop_7"`     //氯含量
	SlaShop8     float64   `json:"sla_shop_8"`     //工作服 功率
	SlaShop9     float64   `json:"sla_shop_9"`     //工作服 时间
	SlaShop10    string    `json:"sla_shop_10"`    //消毒记录 方式
	SlaShop11    string    `json:"sla_shop_11"`    //消毒记录 浓度
	SlaShop12    string    `json:"sla_shop_12"`    //消毒记录 班次
	SlaShop13    string    `json:"sla_shop_13"`    //消毒记录 器具
	SlaShop14    string    `json:"sla_shop_14"`    //消毒记录 环境
}

type DivShop struct {
	//分割车间
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"` //记录时间
	DivShop1     float64   `json:"div_shop_1"`     //紫外灭菌
	DivShop2     float64   `json:"div_shop_2"`     //臭氧
	DivShop3     float64   `json:"div_shop_3"`     //臭氧残留
	DivShop4     float64   `json:"div_shop_4"`     //湿度
	DivShop5     float64   `json:"div_shop_5"`     //温度
	DivShop6     float64   `json:"div_shop_6"`     //预冷间消毒记录
	DivShop7     float64   `json:"div_shop_7"`     //氯含量
	DivShop8     float64   `json:"div_shop_8"`     //工作服 功率
	DivShop9     float64   `json:"div_shop_9"`     //工作服 时间
	DivShop10    string    `json:"div_shop_10"`    //消毒记录 方式
	DivShop11    string    `json:"div_shop_11"`    //消毒记录 浓度
	DivShop12    string    `json:"div_shop_12"`    //消毒记录 班次
	DivShop13    string    `json:"div_shop_13"`    //消毒记录 器具
	DivShop14    string    `json:"div_shop_14"`    //消毒记录 环境
}

type AcidShop struct {
	//排酸车间
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"` //记录时间
	AcidShop1    float64   `json:"acid_shop_1"`    //紫外灭菌
	AcidShop2    float64   `json:"acid_shop_2"`    //臭氧
	AcidShop3    float64   `json:"acid_shop_3"`    //臭氧残留
	AcidShop4    float64   `json:"acid_shop_4"`    //湿度
	AcidShop5    float64   `json:"acid_shop_5"`    //温度
	AcidShop6    float64   `json:"acid_shop_6"`    //预冷间消毒记录
	AcidShop7    float64   `json:"acid_shop_7"`    //氯含量
	AcidShop8    float64   `json:"acid_shop_8"`    //工作服 功率
	AcidShop9    float64   `json:"acid_shop_9"`    //工作服 时间
	AcidShop10   string    `json:"acid_shop_10"`   //消毒记录 方式
	AcidShop11   string    `json:"acid_shop_11"`   //消毒记录 浓度
	AcidShop12   string    `json:"acid_shop_12"`   //消毒记录 班次
	AcidShop13   string    `json:"acid_shop_13"`   //消毒记录 器具
	AcidShop14   string    `json:"acid_shop_14"`   //消毒记录 环境
}

type FroShop struct {
	//冷冻库
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"` //记录时间
	FroShop1     float64   `json:"fro_shop_1"`     //紫外灭菌
	FroShop2     float64   `json:"fro_shop_2"`     //臭氧
	FroShop3     float64   `json:"fro_shop_3"`     //臭氧残留
	FroShop4     float64   `json:"fro_shop_4"`     //湿度
	FroShop5     float64   `json:"fro_shop_5"`     //温度
	FroShop6     float64   `json:"fro_shop_6"`     //预冷间消毒记录
	FroShop7     float64   `json:"fro_shop_7"`     //氯含量
	FroShop8     float64   `json:"fro_shop_8"`     //工作服 功率
	FroShop9     float64   `json:"fro_shop_9"`     //工作服 时间
	FroShop10    string    `json:"fro_shop_10"`    //消毒记录 方式
	FroShop11    string    `json:"fro_shop_11"`    //消毒记录 浓度
	FroShop12    string    `json:"fro_shop_12"`    //消毒记录 班次
	FroShop13    string    `json:"fro_shop_13"`    //消毒记录 器具
	FroShop14    string    `json:"fro_shop_14"`    //消毒记录 环境
}

type PackShop struct {
	//包装车间
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"` //记录时间
	PackShop1    float64   `json:"pack_shop_1"`    //紫外灭菌
	PackShop2    float64   `json:"pack_shop_2"`    //臭氧
	PackShop3    float64   `json:"pack_shop_3"`    //臭氧残留
	PackShop4    float64   `json:"pack_shop_4"`    //湿度
	PackShop5    float64   `json:"pack_shop_5"`    //温度
	PackShop6    float64   `json:"pack_shop_6"`    //预冷间消毒记录
	PackShop7    float64   `json:"pack_shop_7"`    //氯含量
	PackShop8    float64   `json:"pack_shop_8"`    //工作服 功率
	PackShop9    float64   `json:"pack_shop_9"`    //工作服 时间
	PackShop10   string    `json:"pack_shop_10"`   //消毒记录 方式
	PackShop11   string    `json:"pack_shop_11"`   //消毒记录 浓度
	PackShop12   string    `json:"pack_shop_12"`   //消毒记录 班次
	PackShop13   string    `json:"pack_shop_13"`   //消毒记录 器具
	PackShop14   string    `json:"pack_shop_14"`   //消毒记录 环境
}

type PreCoolShopInfo struct {
	//预冷间
	TimeRecordAt string  `json:"time_record_at"`  //记录时间
	PreCoolShop1 float64 `json:"pre_cool_shop_1"` //预冷间温度
	PreCoolShop2 float64 `json:"pre_cool_shop_2"` //预冷间湿度
	PreCoolShop3 float64 `json:"pre_cool_shop_3"` //副产物温度
}

type SlaShopInfo struct {
	//屠宰车间
	TimeRecordAt string  `json:"time_record_at"` //记录时间
	SlaShop1     float64 `json:"sla_shop_1"`     //紫外灭菌
	SlaShop2     float64 `json:"sla_shop_2"`     //臭氧
	SlaShop3     float64 `json:"sla_shop_3"`     //臭氧残留
	SlaShop4     float64 `json:"sla_shop_4"`     //湿度
	SlaShop5     float64 `json:"sla_shop_5"`     //温度
	SlaShop6     float64 `json:"sla_shop_6"`     //预冷间消毒记录
	SlaShop7     float64 `json:"sla_shop_7"`     //氯含量
	SlaShop8     float64 `json:"sla_shop_8"`     //工作服 功率
	SlaShop9     float64 `json:"sla_shop_9"`     //工作服 时间
	SlaShop10    string  `json:"sla_shop_10"`    //消毒记录 方式
	SlaShop11    string  `json:"sla_shop_11"`    //消毒记录 浓度
	SlaShop12    string  `json:"sla_shop_12"`    //消毒记录 班次
	SlaShop13    string  `json:"sla_shop_13"`    //消毒记录 器具
	SlaShop14    string  `json:"sla_shop_14"`    //消毒记录 环境
}

type DivShopInfo struct {
	//分割车间
	TimeRecordAt string  `json:"time_record_at"` //记录时间
	DivShop1     float64 `json:"div_shop_1"`     //紫外灭菌
	DivShop2     float64 `json:"div_shop_2"`     //臭氧
	DivShop3     float64 `json:"div_shop_3"`     //臭氧残留
	DivShop4     float64 `json:"div_shop_4"`     //湿度
	DivShop5     float64 `json:"div_shop_5"`     //温度
	DivShop6     float64 `json:"div_shop_6"`     //预冷间消毒记录
	DivShop7     float64 `json:"div_shop_7"`     //氯含量
	DivShop8     float64 `json:"div_shop_8"`     //工作服 功率
	DivShop9     float64 `json:"div_shop_9"`     //工作服 时间
	DivShop10    string  `json:"div_shop_10"`    //消毒记录 方式
	DivShop11    string  `json:"div_shop_11"`    //消毒记录 浓度
	DivShop12    string  `json:"div_shop_12"`    //消毒记录 班次
	DivShop13    string  `json:"div_shop_13"`    //消毒记录 器具
	DivShop14    string  `json:"div_shop_14"`    //消毒记录 环境
}

type AcidShopInfo struct {
	//排酸车间
	TimeRecordAt string  `json:"time_record_at"` //记录时间
	AcidShop1    float64 `json:"acid_shop_1"`    //紫外灭菌
	AcidShop2    float64 `json:"acid_shop_2"`    //臭氧
	AcidShop3    float64 `json:"acid_shop_3"`    //臭氧残留
	AcidShop4    float64 `json:"acid_shop_4"`    //湿度
	AcidShop5    float64 `json:"acid_shop_5"`    //温度
	AcidShop6    float64 `json:"acid_shop_6"`    //预冷间消毒记录
	AcidShop7    float64 `json:"acid_shop_7"`    //氯含量
	AcidShop8    float64 `json:"acid_shop_8"`    //工作服 功率
	AcidShop9    float64 `json:"acid_shop_9"`    //工作服 时间
	AcidShop10   string  `json:"acid_shop_10"`   //消毒记录 方式
	AcidShop11   string  `json:"acid_shop_11"`   //消毒记录 浓度
	AcidShop12   string  `json:"acid_shop_12"`   //消毒记录 班次
	AcidShop13   string  `json:"acid_shop_13"`   //消毒记录 器具
	AcidShop14   string  `json:"acid_shop_14"`   //消毒记录 环境
}

type FroShopInfo struct {
	//冷冻库
	TimeRecordAt string  `json:"time_record_at"` //记录时间
	FroShop1     float64 `json:"fro_shop_1"`     //紫外灭菌
	FroShop2     float64 `json:"fro_shop_2"`     //臭氧
	FroShop3     float64 `json:"fro_shop_3"`     //臭氧残留
	FroShop4     float64 `json:"fro_shop_4"`     //湿度
	FroShop5     float64 `json:"fro_shop_5"`     //温度
	FroShop6     float64 `json:"fro_shop_6"`     //预冷间消毒记录
	FroShop7     float64 `json:"fro_shop_7"`     //氯含量
	FroShop8     float64 `json:"fro_shop_8"`     //工作服 功率
	FroShop9     float64 `json:"fro_shop_9"`     //工作服 时间
	FroShop10    string  `json:"fro_shop_10"`    //消毒记录 方式
	FroShop11    string  `json:"fro_shop_11"`    //消毒记录 浓度
	FroShop12    string  `json:"fro_shop_12"`    //消毒记录 班次
	FroShop13    string  `json:"fro_shop_13"`    //消毒记录 器具
	FroShop14    string  `json:"fro_shop_14"`    //消毒记录 环境
}

type PackShopInfo struct {
	//包装车间
	TimeRecordAt string  `json:"time_record_at"` //记录时间
	PackShop1    float64 `json:"pack_shop_1"`    //紫外灭菌
	PackShop2    float64 `json:"pack_shop_2"`    //臭氧
	PackShop3    float64 `json:"pack_shop_3"`    //臭氧残留
	PackShop4    float64 `json:"pack_shop_4"`    //湿度
	PackShop5    float64 `json:"pack_shop_5"`    //温度
	PackShop6    float64 `json:"pack_shop_6"`    //预冷间消毒记录
	PackShop7    float64 `json:"pack_shop_7"`    //氯含量
	PackShop8    float64 `json:"pack_shop_8"`    //工作服 功率
	PackShop9    float64 `json:"pack_shop_9"`    //工作服 时间
	PackShop10   string  `json:"pack_shop_10"`   //消毒记录 方式
	PackShop11   string  `json:"pack_shop_11"`   //消毒记录 浓度
	PackShop12   string  `json:"pack_shop_12"`   //消毒记录 班次
	PackShop13   string  `json:"pack_shop_13"`   //消毒记录 器具
	PackShop14   string  `json:"pack_shop_14"`   //消毒记录 环境
}

func ToPreCoolShopInfo(shop *PreCoolShop) PreCoolShopInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := shop.TimeRecordAt.In(loc)
	return PreCoolShopInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		PreCoolShop1: shop.PreCoolShop1,
		PreCoolShop2: shop.PreCoolShop2,
		PreCoolShop3: shop.PreCoolShop3,
	}
}

func ToSlaShopInfo(shop *SlaShop) SlaShopInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := shop.TimeRecordAt.In(loc)
	return SlaShopInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		SlaShop1:     shop.SlaShop1,
		SlaShop2:     shop.SlaShop2,
		SlaShop3:     shop.SlaShop3,
		SlaShop4:     shop.SlaShop4,
		SlaShop5:     shop.SlaShop5,
		SlaShop6:     shop.SlaShop6,
		SlaShop7:     shop.SlaShop7,
		SlaShop8:     shop.SlaShop8,
		SlaShop9:     shop.SlaShop9,
		SlaShop10:    shop.SlaShop10,
		SlaShop11:    shop.SlaShop11,
		SlaShop12:    shop.SlaShop12,
		SlaShop13:    shop.SlaShop13,
		SlaShop14:    shop.SlaShop14,
	}
}

func ToDivShopInfo(shop *DivShop) DivShopInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := shop.TimeRecordAt.In(loc)
	return DivShopInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		DivShop1:     shop.DivShop1,
		DivShop2:     shop.DivShop2,
		DivShop3:     shop.DivShop3,
		DivShop4:     shop.DivShop4,
		DivShop5:     shop.DivShop5,
		DivShop6:     shop.DivShop6,
		DivShop7:     shop.DivShop7,
		DivShop8:     shop.DivShop8,
		DivShop9:     shop.DivShop9,
		DivShop10:    shop.DivShop10,
		DivShop11:    shop.DivShop11,
		DivShop12:    shop.DivShop12,
		DivShop13:    shop.DivShop13,
		DivShop14:    shop.DivShop14,
	}
}

func ToAcidShopInfo(shop *AcidShop) AcidShopInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := shop.TimeRecordAt.In(loc)
	return AcidShopInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		AcidShop1:    shop.AcidShop1,
		AcidShop2:    shop.AcidShop2,
		AcidShop3:    shop.AcidShop3,
		AcidShop4:    shop.AcidShop4,
		AcidShop5:    shop.AcidShop5,
		AcidShop6:    shop.AcidShop6,
		AcidShop7:    shop.AcidShop7,
		AcidShop8:    shop.AcidShop8,
		AcidShop9:    shop.AcidShop9,
		AcidShop10:   shop.AcidShop10,
		AcidShop11:   shop.AcidShop11,
		AcidShop12:   shop.AcidShop12,
		AcidShop13:   shop.AcidShop13,
		AcidShop14:   shop.AcidShop14,
	}
}

func ToFroShopInfo(shop *FroShop) FroShopInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := shop.TimeRecordAt.In(loc)
	return FroShopInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		FroShop1:     shop.FroShop1,
		FroShop2:     shop.FroShop2,
		FroShop3:     shop.FroShop3,
		FroShop4:     shop.FroShop4,
		FroShop5:     shop.FroShop5,
		FroShop6:     shop.FroShop6,
		FroShop7:     shop.FroShop7,
		FroShop8:     shop.FroShop8,
		FroShop9:     shop.FroShop9,
		FroShop10:    shop.FroShop10,
		FroShop11:    shop.FroShop11,
		FroShop12:    shop.FroShop12,
		FroShop13:    shop.FroShop13,
		FroShop14:    shop.FroShop14,
	}
}
