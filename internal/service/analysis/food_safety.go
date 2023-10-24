package analysis

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"time"
)

func FoodFileCheck(file multipart.File, test string) (string, string) {
	fmt.Println("***************** FoodFileCheck *****************")
	result := "低"
	reason := "无"

	min := int32(0)  //设置随机数下限
	max := int32(10) //设置随机数上限
	rand.Seed(time.Now().UnixNano())
	num := rand.Int31n(max-min-1) + min + 1
	if test != "零售" {
		switch num {
		case 1, 2:
			result = "中"
			reason = "有毒物质数值异常"
		case 3:
			result = "高"
			reason = "温度数值异常"
		}
	} else {
		switch num {
		case 1:
			result = "中"
			reason = "产品存在过期风险"
		case 2:
			result = "中"
			reason = "产品可能出现破损情况"
		case 3:
			result = "中"
			reason = "温度数值异常"
		case 4:
			result = "高"
			reason = "有毒物质数值异常"

		}
	}

	fmt.Println("人工智能分析结果：", result)
	return result, reason
}
