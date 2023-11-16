package utils

import (
	"encoding/csv"
	"os"
)

func ReadCSVFile(Path string) ([][]string, error) {
	file, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic("csv file close error!")
			return
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
