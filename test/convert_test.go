package test

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/utils"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJsonToMap(t *testing.T) {
	data := "{\"497.90\": \"448.33\", \"500.34\": \"414.00\"}"
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("result:", result)
	return
}

func TestStructToMap(t *testing.T) {
	heavyMetal := pasture.PastureFeedHeavyMetal{
		TimeRecordAt: time.Now(),
		HouseNumber:  "shshshhs",
		PastureFeedAs: pasture.PastureFeedAs{
			PastureFeedHeavyMetalID: nil,
			As1:                     1,
			As2:                     2,
			As3:                     3,
			As4:                     4,
			As5:                     5,
			As6:                     6,
			As7:                     7,
			As8:                     8,
		},
		PastureFeedPb: pasture.PastureFeedPb{
			PastureFeedHeavyMetalID: nil,
			Pb1:                     1,
			Pb2:                     2,
			Pb3:                     3,
			Pb4:                     4,
			Pb5:                     5,
			Pb6:                     6,
			Pb7:                     7,
		},
		PastureFeedCd: pasture.PastureFeedCd{
			PastureFeedHeavyMetalID: nil,
			Cd1:                     1,
			Cd2:                     2,
			Cd3:                     3,
			Cd4:                     4,
			Cd5:                     5,
			Cd6:                     6,
			Cd7:                     7,
		},
		PastureFeedCr: pasture.PastureFeedCr{
			PastureFeedHeavyMetalID: nil,
			Cr1:                     1,
			Cr2:                     2,
			Cr3:                     3,
			Cr4:                     4,
		},
	}

	m, err := utils.StructToMap(heavyMetal)
	if err != nil {
		panic(err)
	}

	mm := utils.FlattenMap(m)
	fmt.Println(mm)
}
