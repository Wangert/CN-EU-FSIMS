package test

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/service"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"testing"
)

const TESTCONFIGPATH = "../conf/config.yaml"

func TestMigrate(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)
	mysql.AutoMigrate()
}

func TestCreateFsimsUser(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	newUser := models.FSIMSUser{
		UUID:         "UUID_test",
		Name:         "Name_test",
		Account:      "Account_test",
		PasswordHash: "PasswordHash_test",
		Type:         1,
		Status:       1,
		Company:      "Company_test",
		Phone:        "Phone_test",
	}

	err := query.FSIMSUser.WithContext(context.Background()).Create(&newUser)
	if err != nil {
		glog.Errorln("create new user error: %v", err)
	}

	fmt.Println("Create user Successful")

	queryUsers, err := query.FSIMSUser.WithContext(context.Background()).Find()
	if err != nil {
		glog.Errorln(err)
	} else {
		for _, r := range queryUsers {
			fmt.Println(*r)
		}
	}
}

func TestFindFsimsUsers(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	ics, err := query.FSIMSUser.WithContext(context.Background()).Find()
	if err != nil {
		panic(err)
	} else {
		for _, ic := range ics {
			fmt.Println(*ic)
		}
	}
}

func TestUpdateFsimsUsers(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)
	account := "Account_test"
	err := service.ResetFsimsPassWord(account)
	if err != nil {
		glog.Errorln("user's password updated error!")
	}
}

func createPastureProcedureWithPasPID(pasPID string) pasture.PastureProcedure {
	pastureProcedure := pasture.PastureProcedure{
		PasPID: pasPID,
		Water: pasture.PastureWater{
			PhysicalHazard: pasture.PastureWaterPhysicalHazard{
				Mercury:  0,
				Cadmium:  0,
				Lead:     0,
				Chromium: 0,
				Arsenic:  0,
				Copper:   0,
			},
			ChemicalHazard: pasture.PastureWaterChemicalHazard{
				Fluoride:        0,
				Cyanide:         0,
				Chloride:        0,
				Nitrate:         0,
				Sulfate:         0,
				Sixsixsix:       0,
				DDT:             0,
				AmmoniaNitrogen: 0,
			},
			Biohazard: pasture.PastureWaterBiohazard{
				ColiformBacteria: 0,
			},
			SensoryTraits: pasture.PastureWaterSensoryTraits{
				Color:          0,
				Turbidity:      0,
				Smell:          "",
				PH:             0,
				Hardness:       0,
				DissolvedSolid: 0,
			},
			PasPID: pasPID,
		},
		Fodder: pasture.PastureFodder{
			PhysicalHazard: pasture.PastureFodderPhysicalHazard{
				Mercury:                  0,
				Cadmium:                  0,
				Lead:                     0,
				Chromium:                 0,
				Arsenic:                  0,
				Fluorine:                 0,
				Nitrite:                  0,
				PolychlorinatedBiphenyls: 0,
				Sixsixsix:                0,
				DDT:                      0,
				Bexachlorobenzene:        0,
			},
			Biohazard: pasture.PastureFodderBiohazard{
				AflatoxinB1:       0,
				OchratoxinA:       0,
				Zearalenone:       0,
				Vomitoxin:         0,
				T2:                0,
				B1B2:              0,
				Cyanide:           0,
				FreeGossypol:      0,
				Isothiocyanate:    0,
				Oxazolidinethione: 0,
				TotalBacteria:     0,
				TotalMold:         0,
				Salmonella:        0,
			},
			PasPID: pasPID,
		},
		Soil: pasture.PastureSoil{
			PhysicalHazard: pasture.PastureSoilPhysicalHazard{
				Cadmium:  0,
				Arsenic:  0,
				Copper:   0,
				Lead:     0,
				Chromium: 0,
				Zinc:     0,
			},
			Biohazard: pasture.PastureSoilBiohazard{
				ColiformBacteria: 0,
				TotalBacteria:    0,
			},
			PasPID: pasPID,
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
			PhysicalHazard: pasture.PastureFloorBeddingPhysicalHazard{
				Mercury:  0,
				Cadmium:  0,
				Lead:     0,
				Chromium: 0,
			},
			Biohazard: pasture.PastureFloorBeddingBiohazard{
				ColiformBacteria: 0,
				TotalBacteria:    0,
				TotalMold:        0,
				AflatoxinB1:      0,
				Salmonella:       0,
			},
			PasPID: pasPID,
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

func createPastureWater() pasture.PastureWater {
	return pasture.PastureWater{
		PhysicalHazard: pasture.PastureWaterPhysicalHazard{
			Mercury:  1,
			Cadmium:  1,
			Lead:     0,
			Chromium: 0,
			Arsenic:  0,
			Copper:   0,
		},
		ChemicalHazard: pasture.PastureWaterChemicalHazard{
			Fluoride:        1,
			Cyanide:         1,
			Chloride:        0,
			Nitrate:         0,
			Sulfate:         0,
			Sixsixsix:       0,
			DDT:             0,
			AmmoniaNitrogen: 0,
		},
		Biohazard: pasture.PastureWaterBiohazard{
			ColiformBacteria: 1,
		},
		SensoryTraits: pasture.PastureWaterSensoryTraits{
			Color:          1,
			Turbidity:      0,
			Smell:          "",
			PH:             0,
			Hardness:       0,
			DissolvedSolid: 0,
		},
		PasPID: "PASPID-111",
	}
}

func createPastureWaterPhysicalHazard(pastureWaterID uint) pasture.PastureWaterPhysicalHazard {
	return pasture.PastureWaterPhysicalHazard{
		Mercury:        0,
		Cadmium:        0,
		Lead:           0,
		Chromium:       0,
		Arsenic:        0,
		Copper:         0,
		PastureWaterID: pastureWaterID,
	}
}

func createPastureWaterChemicalHazard(pastureWaterID uint) pasture.PastureWaterChemicalHazard {
	return pasture.PastureWaterChemicalHazard{
		Fluoride:        0,
		Cyanide:         0,
		Chloride:        0,
		Nitrate:         0,
		Sulfate:         0,
		Sixsixsix:       0,
		DDT:             0,
		AmmoniaNitrogen: 0,
		PastureWaterID:  pastureWaterID,
	}
}

func createPastureWaterBiohazard(pastureWaterID uint) pasture.PastureWaterBiohazard {
	return pasture.PastureWaterBiohazard{
		ColiformBacteria: 0,
		PastureWaterID:   pastureWaterID,
	}
}

func createPastureWaterSensoryTraits(pastureWaterID uint) pasture.PastureWaterSensoryTraits {
	return pasture.PastureWaterSensoryTraits{
		Color:          0,
		Turbidity:      0,
		Smell:          "",
		PH:             0,
		Hardness:       0,
		DissolvedSolid: 0,
		PastureWaterID: pastureWaterID,
	}
}

func TestCreatePastureWater(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

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

			fmt.Println("create pasture procedure successful!")
		}
	}()

	pastureProcedure1 := createPastureProcedureWithPasPID("PASPID-111")
	err = tx.PastureProcedure.WithContext(context.Background()).Create(&pastureProcedure1)
	pastureWater := createPastureWater()
	err = tx.PastureWater.WithContext(context.Background()).Create(&pastureWater)

	fmt.Println("create pasture water successful!")
}

func TestCreatePastureProcedures(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

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

			fmt.Println("create pasture procedure successful!")
		}
	}()

	pastureProcedure1 := createPastureProcedureWithPasPID("PASPID-222")
	err = tx.PastureProcedure.WithContext(context.Background()).Create(&pastureProcedure1)
	pastureWater1, err := tx.PastureWater.WithContext(context.Background()).Where(tx.PastureWater.PasPID.Eq(pastureProcedure1.PasPID)).First()

	pwph := createPastureWaterPhysicalHazard(pastureWater1.ID)
	err = tx.PastureWaterPhysicalHazard.WithContext(context.Background()).Create(&pwph)
	pwch := createPastureWaterChemicalHazard(pastureWater1.ID)
	err = tx.PastureWaterChemicalHazard.WithContext(context.Background()).Create(&pwch)
	pwbh := createPastureWaterBiohazard(pastureWater1.ID)
	err = tx.PastureWaterBiohazard.WithContext(context.Background()).Create(&pwbh)
	pwst := createPastureWaterSensoryTraits(pastureWater1.ID)
	err = tx.PastureWaterSensoryTraits.WithContext(context.Background()).Create(&pwst)
}

func TestFindPastureProcedure(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	pasP := query.PastureProcedure
	pastureProcedures, err := pasP.WithContext(context.Background()).Where(pasP.PasPID.Eq("PASPID-222")).
		Preload(pasP.Water).Preload(pasP.Water.PhysicalHazard, pasP.Water.ChemicalHazard, pasP.Water.Biohazard, pasP.Water.SensoryTraits).Find()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pasture Procedure:")
	for _, p := range pastureProcedures {
		js, _ := json.Marshal(*p)
		var out bytes.Buffer
		json.Indent(&out, js, "", "\t")
		fmt.Printf("%+v\n", out.String())
	}
}

//func TestCreateIndustrialChainWithProcedures(t *testing.T) {
//	mysql.Init(TESTCONFIGPATH)
//
//	procedure1 := models.Procedure{
//		PID:          "PID-3",
//		Name:         "Production",
//		SerialNumber: 1,
//		PrePID:       "PID-0",
//		ICID:         "ICID-11111",
//	}
//
//	procedure2 := models.Procedure{
//		PID:          "PID-4",
//		Name:         "Slaughter",
//		SerialNumber: 2,
//		PrePID:       "PID-1",
//		ICID:         "ICID-11111",
//	}
//
//	ic := models.IndustrialChain{
//		ICID:               "ICID-22222",
//		Type:               "fish",
//		State:              2,
//		StartTimestamp:     time.Now(),
//		CompletedTimestamp: time.Now(),
//		Procedures: []models.Procedure{
//			procedure1, procedure2,
//		},
//	}
//
//	err := query.IndustrialChain.WithContext(context.Background()).Create(&ic)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("create industrial chain successful!")
//}
//
//func TestFindIndustrialChains(t *testing.T) {
//	mysql.Init(TESTCONFIGPATH)
//
//	ics, err := query.IndustrialChain.WithContext(context.Background()).Find()
//	if err != nil {
//		panic(err)
//	} else {
//		for _, ic := range ics {
//			fmt.Println(*ic)
//		}
//	}
//}
//
//func TestFindAllProdecuresWithIndustrialChain(t *testing.T) {
//	mysql.Init(TESTCONFIGPATH)
//
//	ic, err := query.IndustrialChain.WithContext(context.Background()).First()
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("industrial_chain:")
//		fmt.Println(*ic)
//	}
//
//	fmt.Println("=======================================")
//	procedures, err := query.IndustrialChain.Procedures.WithContext(context.Background()).Model(ic).Find()
//	if err != nil {
//		panic(err)
//	} else {
//		for _, procedure := range procedures {
//			fmt.Println(*procedure)
//		}
//	}
//}

//func TestDeleteIndustrialChains(t *testing.T) {
//	infoResult, err := query.FSIMSUser.WithContext(context.Background()).Where(query.IndustrialChain.ID.Gt(0)).Delete()
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("Delete IndustrialChains(ID > 0) Successful")
//		fmt.Println(infoResult)
//	}
//}

func TestDeleteProcedures(t *testing.T) {
	infoResult, err := query.FSIMSUser.WithContext(context.Background()).Where(query.Procedure.ID.Gt(0)).Delete()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Procedures(ID > 0) Successful")
		fmt.Println(infoResult)
	}
}
