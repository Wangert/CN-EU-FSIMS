package analysis

func RiskLevel(abnormalCount int) int {
	if abnormalCount == 0 {
		return 1
	} else if abnormalCount > 0 && abnormalCount <= 2 {
		return 2
	} else if abnormalCount > 2 && abnormalCount <= 5 {
		return 3
	} else if abnormalCount > 5 {
		return 4
	}

	return 0
}
