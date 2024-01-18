package analysis

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/utils"
	"fmt"
)

// 判断牧场饲料重金属危害
func JudgeHarmForPastureFeedHeavyMetal(record *pasture.PastureFeedHeavyMetal) ([]string, int, error) {
	//glog.Info("test_record:", record)
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)
	//glog.Info("test_indexMap:", indexMap)
	abnormalList := make([]string, 0)
	for k, v := range PastureFeedHeavyMetalBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	return abnormalList, len(abnormalList), nil
}

// 判断牧场饲料真菌毒素危害
func JudgeHarmForPastureFeedMycotoxins(record *pasture.PastureFeedMycotoxins) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PastureFeedMycotoxinsBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	return abnormalList, len(abnormalList), nil
}

// 判断牧场水质危害
func JudgeHarmForPastureWaterQuality(record *pasture.PastureWaterRecord) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PastureWaterQualityUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	for k, v := range PastureWaterQualityLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	return abnormalList, len(abnormalList), nil
}

// 判断牧场缓冲区危害
func JudgeHarmForPastureBuffer(record *pasture.PastureBuffer) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PastureBufferUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断牧场场区危害
func JudgeHarmForPastureArea(record *pasture.PastureArea) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PastureAreaUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断牧场牛舍危害
func JudgeHarmForPastureCowHouse(record *pasture.CowHouse) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PastureCowHouseUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断牧场基本环境危害
func JudgeHarmForPastureBasicEnvironment(record *pasture.PastureBasicEnvironment) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PastureBasicEnvironmentUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	for k, v := range PastureBasicEnvironmentLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断牧场垫料危害
func JudgeHarmForPasturePadding(record *pasture.PasturePaddingRequire) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range PasturePaddingUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断屠宰场水质危害
func JudgeHarmForSlaughterWaterQuality(record *slaughter.SlaughterWaterQualityMon) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range SlaughterWaterQualityUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	for k, v := range SlaughterWaterQualityLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断屠宰过程危害
func JudgeHarmForSlaughterProcedure(record *slaughter.SlaughterProcedureMonitoringData) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range SlaughterProcedureSensorDataUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	for k, v := range SlaughterProcedureSensorDataLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断屠宰场预冷车间危害
func JudgeHarmForSlaughterPreCoolShop(record *slaughter.PreCoolShop) ([]string, int, error) {
	fmt.Println("000000000000000000000000000000000000000000000000000000")
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	fmt.Println("111111111111111111111111111111111111111111111111111")
	for k, v := range SlaughterPreCoolShopUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	fmt.Println("22222222222222222222222222222222222222222222222222")

	for k, v := range SlaughterPreCoolShopLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断屠宰场屠宰车间危害
func JudgeHarmForSlaughterSlaShop(record *slaughter.SlaShop) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}

	indexMap = utils.FlattenMap(recordMap)
	abnormalList := make([]string, 0)
	for k, v := range SlaughterSlaShopUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	for k, v := range SlaughterSlaShopLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}
	return abnormalList, len(abnormalList), nil
}

// 判断屠宰场分割车间危害
func JudgeHarmForSlaughterDivShop(record *slaughter.DivShop) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range SlaughterDivShopUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	for k, v := range SlaughterDivShopLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断屠宰场排酸车间危害
func JudgeHarmForSlaughterAcidShop(record *slaughter.AcidShop) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range SlaughterAcidShopUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	for k, v := range SlaughterAcidShopLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}

// 判断屠宰场冷冻车间危害
func JudgeHarmForSlaughterFroShop(record *slaughter.FroShop) ([]string, int, error) {
	indexMap := map[string]interface{}{}
	recordMap, err := utils.StructToMap(record)
	if err != nil {
		return nil, 0, err
	}
	indexMap = utils.FlattenMap(recordMap)

	abnormalList := make([]string, 0)
	for k, v := range SlaughterFroShopUpperBounds {
		if v < indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	for k, v := range SlaughterFroShopLowerBounds {
		if v > indexMap[k].(float64) {
			abnormalList = append(abnormalList, k)
		}
	}

	return abnormalList, len(abnormalList), nil
}
