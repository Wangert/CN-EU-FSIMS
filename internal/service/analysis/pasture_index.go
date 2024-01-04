package analysis

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/utils"
)

// 判断牧场饲料重金属危害
func JudgeHarmForPastureFeedHeavyMetal(record *pasture.PastureFeedHeavyMetal) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 3)
	for k, v := range PastureFeedHeavyMetalBounds {
		if v <= indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	return abnormalList, len(abnormalList), nil
}

// 判断牧场饲料真菌毒素危害
func JudgeHarmForPastureFeedFungus(record *pasture.PastureFeedMycotoxins) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 3)
	for k, v := range PastureFeedMycotoxinsBounds {
		if v <= indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	return abnormalList, len(abnormalList), nil
}
