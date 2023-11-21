package utils

import (
	"bytes"
	"encoding/gob"
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
