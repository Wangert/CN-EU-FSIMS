package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/utils"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"github.com/golang/glog"
	"time"
)

const (
	COW_NUMBER_PREFIX   = "COW-"
	BATCH_NUMBER_PREFIX = "BATCH-"

	// cow state
	INIT_STATE      = 1
	FEEDING_STATE   = 2
	WAREHOUSE_STATE = 3
	SENDING_STATE   = 4
	END_STATE       = 5
)

func NewFeedingBatch(r *request.ReqNewFeedingBatch) (string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				panic(err)
			}

			glog.Infoln("new feeding batch successful!")
		}
	}()

	pp := NewProcedureParams{
		Type:     PASTURE_TYPE,
		Operator: r.Worker,
		PrePID:   r.PrePID,
	}
	procedure, err := NewProcedure(&pp)

	glog.Infoln("CowNumbers:")
	glog.Infoln(r.CowNumbers)
	cows, err := GetCowsByNumbers(r.CowNumbers)

	glog.Infoln("Cows:")
	glog.Infoln(cows)

	bNum := BATCH_NUMBER_PREFIX + generateNumber(r)
	fb := pasture.FeedingBatch{
		BatchNumber: bNum,
		HouseNumber: r.HouseNumber,
		State:       1,
		PID:         procedure.PID,
		Worker:      r.Worker,
		Cows:        cows,
	}

	err = tx.Procedure.WithContext(context.Background()).Create(&procedure)
	err = tx.FeedingBatch.WithContext(context.Background()).Create(&fb)

	return bNum, nil
}

func GetCowsByNumbers(numbers []string) ([]product.Cow, error) {
	cows := make([]product.Cow, len(numbers))
	q := query.Cow
	for i, number := range numbers {
		c, err := q.WithContext(context.Background()).Where(q.Number.Eq(number)).First()
		if err != nil {
			return nil, err
		}

		cows[i] = *c
	}

	return cows, nil
}

func PastureIsExist(number string) (bool, error) {
	q := query.PastureHouse
	count, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(number)).Count()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	return true, nil
}

func AddCow(r *request.ReqAddCow) (product.CowInfo, error) {
	cowNum := generateNumber(r)
	cow := product.Cow{
		Number:      COW_NUMBER_PREFIX + cowNum,
		Age:         r.Age,
		Weight:      r.Weight,
		State:       INIT_STATE,
		HouseNumber: r.HouseNumber,
		BatchNumber: nil,
	}

	q := query.Cow
	err := q.WithContext(context.Background()).Create(&cow)
	if err != nil {
		return product.CowInfo{}, err
	}

	cowInfo := product.ToCowInfo(&cow)
	return cowInfo, nil
}

func generateNumber(i interface{}) string {
	t := time.Now().String()
	s := utils.SerializeStructToString(i)
	h := crypto.CalculateSHA256(s+t, "number")

	return h
}

func createPastureProcedure(pasPID string) pasture.PastureProcedure {
	pastureProcedure := pasture.PastureProcedure{
		PasPID: pasPID,
		Water: pasture.PastureWater{
			PhysicalHazard: pasture.PastureWaterPhysicalHazard{},
			ChemicalHazard: pasture.PastureWaterChemicalHazard{},
			Biohazard:      pasture.PastureWaterBiohazard{},
			SensoryTraits:  pasture.PastureWaterSensoryTraits{},
			PasPID:         pasPID,
		},
		Fodder: pasture.PastureFodder{
			PhysicalHazard: pasture.PastureFodderPhysicalHazard{},
			Biohazard:      pasture.PastureFodderBiohazard{},
			PasPID:         pasPID,
		},
		Soil: pasture.PastureSoil{
			PhysicalHazard: pasture.PastureSoilPhysicalHazard{},
			Biohazard:      pasture.PastureSoilBiohazard{},
			PasPID:         pasPID,
		},
		Air: pasture.PastureAir{
			TotalBacteria:   0,
			AmmoniaGas:      0,
			HydrogenSulfide: 0,
			CarbonDioxide:   0,
			PM10:            0,
			TSP:             0,
			Stench:          0,
			PasPID:          pasPID,
		},
		FloorBedding: pasture.PastureFloorBedding{
			PhysicalHazard: pasture.PastureFloorBeddingPhysicalHazard{},
			Biohazard:      pasture.PastureFloorBeddingBiohazard{},
			PasPID:         pasPID,
		},
		WasteDischarge: pasture.PastureWasteDischarge{
			WaterFivedayBOD:          0,
			WaterChemicalOxygen:      0,
			WaterAmmoniaNitrogen:     0,
			WaterPhosphorus:          0,
			WaterSuspendedMatter:     0,
			WaterFecalColiform:       0,
			WaterAO:                  0,
			WasteSlagFecalColiform:   0,
			WasteSlagAOMortalityRate: 0,
			O3Concentration:          0,
			PasPID:                   pasPID,
		},
	}

	return pastureProcedure
}
