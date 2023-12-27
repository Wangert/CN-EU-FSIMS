package premortem

import "gorm.io/gorm"

type PreSlaInfoRec struct {
	//宰前信息记录
	gorm.Model
	TimeRecordAt   string           `json:"time_record_at"`
	PreSlaInfoRec1 Gps              `gorm:"foreignKey:PreSlaInfoRecID; references:ID" json:"pre_sla_info_rec_1"` //Gps信息
	PreSlaInfoRec2 PreSlaDietManage `gorm:"foreignKey:PreSlaInfoRecID; references:ID" json:"pre_sla_info_rec_2"` //宰前牛饮食管理
	PreSlaInfoRec3 bool             `json:"pre_sla_info_rec_3"`                                                  //廋肉精含量
	PreSlaInfoRec4 string           `json:"pre_sla_info_rec_4"`                                                  //运载工具消毒时间
	PreSlaInfoRec5 float32          `json:"pre_sla_info_rec_5"`                                                  //消毒剂浓度
	PreSlaInfoRec6 PreSlaGerms      `gorm:"foreignKey:PreSlaInfoRecID; references:ID" json:"pre_sla_info_rec_6"` //待宰圈致病菌（沙门氏、STEC、单增）检测记录
	PreSlaInfoRec7 string           `json:"pre_sla_info_rec_7"`                                                  //待宰圈消毒清理记录
	PreSlaInfoRec8 PreSlaPicAndEn   `gorm:"foreignKey:PreSlaInfoRecID; references:ID" json:"pre_sla_info_rec_8"` //待宰动物视频，头部，皮毛采集
}
type Gps struct {
	gorm.Model
	PreSlaInfoRecID    uint   `json:"pre_sla_info_rec_id"`
	ChiledFreshIndexID uint   `json:"chiled_fresh_index_id"`
	Gps1               string `json:"gps_1"` //gps坐标
	Gps2               string `json:"gps_2"` //时间
}

type PreSlaGerms struct {
	gorm.Model
	PreSlaInfoRecID uint    `json:"pre_sla_info_rec_id"`
	PreSlaGerms1    float32 `json:"pre_sla_germs_1"` //沙门氏
	PreSlaGerms2    float32 `json:"pre_sla_germs_2"` //STEC
	PreSlaGerms3    float32 `json:"pre_sla_germs_3"` //单增
}

type PreSlaDietManage struct {
	//宰前饮食管理
	gorm.Model
	PreSlaInfoRecID   uint   `json:"pre_sla_info_rec_id"`
	PreSlaDietManage1 string `json:"pre_sla_diet_manage_1"` //牛进场时间
	PreSlaDietManage2 string `json:"pre_sla_diet_manage_2"` //牛停水时间
	PreSlaDietManage3 string `json:"pre_sla_diet_manage_3"` //牛禁食时间
	PreSlaDietManage4 string `json:"pre_sla_diet_manage_4"` //饮水时间
	PreSlaDietManage5 string `json:"pre_sla_diet_manage_5"` //送宰时间
}

type PreSlaPicAndEn struct {
	//类型后续调整，先做标记
	//宰前视频，图片，环境
	gorm.Model
	PreSlaInfoRecID uint   `json:"pre_sla_info_rec_id"`
	PreSlaPicAndEn1 string `json:"pre_sla_pic_and_en_1"` //待宰动物视频
	PreSlaPicAndEn2 string `json:"pre_sla_pic_and_en_2"` //待宰动物头部
	PreSlaPicAndEn3 string `json:"pre_sla_pic_and_en_3"` //待宰动物皮毛相片采集
	PreSlaPicAndEn4 string `json:"pre_sla_pic_and_en_4"` //照度
	PreSlaPicAndEn5 string `json:"pre_sla_pic_and_en_5"` //噪声
	PreSlaPicAndEn6 string `json:"pre_sla_pic_and_en_6"` //光照时间
}
