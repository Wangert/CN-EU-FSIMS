package pasture

import "gorm.io/gorm"

type Cow struct {
	gorm.Model
	TimeRecordAt string `json:"time_record_at"` //记录时间
	Cow1         string `json:"cow_1"`          //入场时间
	Cow2         string `json:"cow_2"`          //防疫记录
	Cow3         string `json:"cow_3"`          //用药记录
	Cow4         string `json:"cow_4"`          //饲料记录
	Cow5         string `json:"cow_5"`          //检疫证号
	Cow6         string `json:"cow_6"`          //牛主身份证号
	Cow7         string `json:"cow_7"`          //牛主地址
	Cow8         string `json:"cow_8"`          //毛色品种
	Cow9         string `json:"cow_9"`          //体重（记录）
	Cow10        string `json:"cow_10"`         //饮水记录
}
