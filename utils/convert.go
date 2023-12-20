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
