package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
)

func SerializeStructToBytes(i interface{}) []byte {
	// create buffer
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(i)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func SerializeStructToString(i interface{}) string {
	// create buffer
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(i)

	if err != nil {
		log.Panic(err)
	}

	return result.String()
}

// struct to map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func FlattenMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range m {
		if nestedMap, ok := value.(map[string]interface{}); ok { // 判断值是否为嵌套的map类型
			nestedResult := FlattenMap(nestedMap)              // 递归调用自身处理嵌套的map
			for nestedKey, nestedValue := range nestedResult { // 合并结果到最外层的map
				result[nestedKey] = nestedValue
			}
		} else {
			result[key] = value
		}
	}

	return result
}

func StrArrToStr(strArr []string) string {
	result := ""

	for _, s := range strArr {
		result = result + "[" + s + "]" + ";"
	}

	return result
}
