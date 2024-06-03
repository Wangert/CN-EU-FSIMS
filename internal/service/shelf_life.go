package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/imgquery"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/golang/glog"
)

const T = 4

func GetAllSpectralData() ([]models.Spectra, int64, error) {
	q := imgquery.Spectra
	sd, err := q.WithContext(context.Background()).Find()
	if err != nil {
		return []models.Spectra{}, 0, err
	}

	count, err := q.WithContext(context.Background()).Count()

	sda := make([]models.Spectra, count)
	for i, s := range sd {
		sda[i] = models.Spectra(*s)
	}

	return sda, count, nil
}

func GetOneSpectralData() (string, error) {
	q := imgquery.Spectra
	s, err := q.WithContext(context.Background()).First()
	if err != nil {
		return "", err
	}

	return s.Data, nil
}

func ShelfLifeForecast(data string) (float64, error) {
	glog.Info("data:", data)
	//将json字符串进行解析
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 0.0, err
	}

	//开始计算
	_, x := findClosestKey(result, 826.63)
	_, y := findClosestKey(result, 938.11)
	_, z := findClosestKey(result, 1270.93)
	_, k := findClosestKey(result, 1319.24)

	lam := GetLambda(T)
	u := GetMuMax(T)
	no, ok := GetNo(x, y, z, k)
	if !ok {
		return 0.0, nil
	}

	part1 := (9.763 - no) / u * math.E
	part2 := math.Log(-math.Log((7.34-no)/(9.763-no))) - 1
	SL := lam - part1*part2
	return math.Round(SL), nil
}

func GetLambda(t float64) float64 {
	return math.Pow(0.01381*t+0.12459, -2)
}

func GetMuMax(t float64) float64 {
	return math.Pow(0.01684*t+0.14215, -2)
}

func GetNo(x, y, z, k interface{}) (float64, bool) {
	//return 6.017 - 0.027*x + 0.025*y + 0.029*z - 0.027*k
	xf, ok1 := transform(x)
	yf, ok2 := transform(y)
	zf, ok3 := transform(z)
	kf, ok4 := transform(k)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return 0.0, false
	}
	return 6.017 - 0.027*xf + 0.025*yf + 0.029*zf - 0.027*kf, true
}

func transform(p interface{}) (float64, bool) {
	strx, ok := p.(string)
	if !ok {
		fmt.Println("Value is not a string")
		return 0.0, false
	}
	floatValue, err := strconv.ParseFloat(strx, 64) // 64表示解析为float64
	if err != nil {
		fmt.Printf("Error parsing string to float64: %v\n", err)
		return 0.0, false
	}
	return floatValue, true
}

func findClosestKey(m map[string]interface{}, target float64) (string, interface{}) {
	var closestKey string
	var closestValue interface{}
	minDiff := math.MaxInt32

	for k, v := range m {
		f, err := strconv.ParseFloat(k, 64)
		if err != nil {
			fmt.Println("String to float64:", err)
			return "", nil
		}

		diff := int(math.Abs(f - target))
		if diff < minDiff {
			minDiff = diff
			closestKey = k
			closestValue = v
		}
	}

	return closestKey, closestValue
}
