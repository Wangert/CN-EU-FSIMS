package service

import "CN-EU-FSIMS/internal/app/models/pasture"

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
