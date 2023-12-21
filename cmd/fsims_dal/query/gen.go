// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                 = new(Query)
	Cow                               *cow
	FSIMSUser                         *fSIMSUser
	FattenProcedure                   *fattenProcedure
	FattenSoil                        *fattenSoil
	FattenSoilBiohazard               *fattenSoilBiohazard
	FattenSoilPhysicalHazard          *fattenSoilPhysicalHazard
	FattenWater                       *fattenWater
	FattenWaterBiohazard              *fattenWaterBiohazard
	FattenWaterChemicalHazard         *fattenWaterChemicalHazard
	FattenWaterPhysicalHazard         *fattenWaterPhysicalHazard
	FattenWaterSensoryTraits          *fattenWaterSensoryTraits
	FeedingBatch                      *feedingBatch
	Log                               *log
	PackWarehouse                     *packWarehouse
	PackageBatch                      *packageBatch
	PackageHouse                      *packageHouse
	PackageReceiveRecord              *packageReceiveRecord
	PastureAir                        *pastureAir
	PastureFloorBedding               *pastureFloorBedding
	PastureFloorBeddingBiohazard      *pastureFloorBeddingBiohazard
	PastureFloorBeddingPhysicalHazard *pastureFloorBeddingPhysicalHazard
	PastureFodder                     *pastureFodder
	PastureFodderBiohazard            *pastureFodderBiohazard
	PastureFodderPhysicalHazard       *pastureFodderPhysicalHazard
	PastureHouse                      *pastureHouse
	PastureProcedure                  *pastureProcedure
	PastureSiteDisinfectionRecord     *pastureSiteDisinfectionRecord
	PastureSoil                       *pastureSoil
	PastureSoilBiohazard              *pastureSoilBiohazard
	PastureSoilPhysicalHazard         *pastureSoilPhysicalHazard
	PastureTruckDisinfectionRecord    *pastureTruckDisinfectionRecord
	PastureWarehouse                  *pastureWarehouse
	PastureWasteDischarge             *pastureWasteDischarge
	PastureWater                      *pastureWater
	PastureWaterBiohazard             *pastureWaterBiohazard
	PastureWaterChemicalHazard        *pastureWaterChemicalHazard
	PastureWaterPhysicalHazard        *pastureWaterPhysicalHazard
	PastureWaterSensoryTraits         *pastureWaterSensoryTraits
	PastureWorksuitDisinfectionRecord *pastureWorksuitDisinfectionRecord
	Procedure                         *procedure
	SlaughterBatch                    *slaughterBatch
	SlaughterHouse                    *slaughterHouse
	SlaughterProduct                  *slaughterProduct
	SlaughterReceiveRecord            *slaughterReceiveRecord
	SlaughterWarehouse                *slaughterWarehouse
	TransportProcedureData            *transportProcedureData
	TransportVehicle                  *transportVehicle
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Cow = &Q.Cow
	FSIMSUser = &Q.FSIMSUser
	FattenProcedure = &Q.FattenProcedure
	FattenSoil = &Q.FattenSoil
	FattenSoilBiohazard = &Q.FattenSoilBiohazard
	FattenSoilPhysicalHazard = &Q.FattenSoilPhysicalHazard
	FattenWater = &Q.FattenWater
	FattenWaterBiohazard = &Q.FattenWaterBiohazard
	FattenWaterChemicalHazard = &Q.FattenWaterChemicalHazard
	FattenWaterPhysicalHazard = &Q.FattenWaterPhysicalHazard
	FattenWaterSensoryTraits = &Q.FattenWaterSensoryTraits
	FeedingBatch = &Q.FeedingBatch
	Log = &Q.Log
	PackWarehouse = &Q.PackWarehouse
	PackageBatch = &Q.PackageBatch
	PackageHouse = &Q.PackageHouse
	PackageReceiveRecord = &Q.PackageReceiveRecord
	PastureAir = &Q.PastureAir
	PastureFloorBedding = &Q.PastureFloorBedding
	PastureFloorBeddingBiohazard = &Q.PastureFloorBeddingBiohazard
	PastureFloorBeddingPhysicalHazard = &Q.PastureFloorBeddingPhysicalHazard
	PastureFodder = &Q.PastureFodder
	PastureFodderBiohazard = &Q.PastureFodderBiohazard
	PastureFodderPhysicalHazard = &Q.PastureFodderPhysicalHazard
	PastureHouse = &Q.PastureHouse
	PastureProcedure = &Q.PastureProcedure
	PastureSiteDisinfectionRecord = &Q.PastureSiteDisinfectionRecord
	PastureSoil = &Q.PastureSoil
	PastureSoilBiohazard = &Q.PastureSoilBiohazard
	PastureSoilPhysicalHazard = &Q.PastureSoilPhysicalHazard
	PastureTruckDisinfectionRecord = &Q.PastureTruckDisinfectionRecord
	PastureWarehouse = &Q.PastureWarehouse
	PastureWasteDischarge = &Q.PastureWasteDischarge
	PastureWater = &Q.PastureWater
	PastureWaterBiohazard = &Q.PastureWaterBiohazard
	PastureWaterChemicalHazard = &Q.PastureWaterChemicalHazard
	PastureWaterPhysicalHazard = &Q.PastureWaterPhysicalHazard
	PastureWaterSensoryTraits = &Q.PastureWaterSensoryTraits
	PastureWorksuitDisinfectionRecord = &Q.PastureWorksuitDisinfectionRecord
	Procedure = &Q.Procedure
	SlaughterBatch = &Q.SlaughterBatch
	SlaughterHouse = &Q.SlaughterHouse
	SlaughterProduct = &Q.SlaughterProduct
	SlaughterReceiveRecord = &Q.SlaughterReceiveRecord
	SlaughterWarehouse = &Q.SlaughterWarehouse
	TransportProcedureData = &Q.TransportProcedureData
	TransportVehicle = &Q.TransportVehicle
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                                db,
		Cow:                               newCow(db, opts...),
		FSIMSUser:                         newFSIMSUser(db, opts...),
		FattenProcedure:                   newFattenProcedure(db, opts...),
		FattenSoil:                        newFattenSoil(db, opts...),
		FattenSoilBiohazard:               newFattenSoilBiohazard(db, opts...),
		FattenSoilPhysicalHazard:          newFattenSoilPhysicalHazard(db, opts...),
		FattenWater:                       newFattenWater(db, opts...),
		FattenWaterBiohazard:              newFattenWaterBiohazard(db, opts...),
		FattenWaterChemicalHazard:         newFattenWaterChemicalHazard(db, opts...),
		FattenWaterPhysicalHazard:         newFattenWaterPhysicalHazard(db, opts...),
		FattenWaterSensoryTraits:          newFattenWaterSensoryTraits(db, opts...),
		FeedingBatch:                      newFeedingBatch(db, opts...),
		Log:                               newLog(db, opts...),
		PackWarehouse:                     newPackWarehouse(db, opts...),
		PackageBatch:                      newPackageBatch(db, opts...),
		PackageHouse:                      newPackageHouse(db, opts...),
		PackageReceiveRecord:              newPackageReceiveRecord(db, opts...),
		PastureAir:                        newPastureAir(db, opts...),
		PastureFloorBedding:               newPastureFloorBedding(db, opts...),
		PastureFloorBeddingBiohazard:      newPastureFloorBeddingBiohazard(db, opts...),
		PastureFloorBeddingPhysicalHazard: newPastureFloorBeddingPhysicalHazard(db, opts...),
		PastureFodder:                     newPastureFodder(db, opts...),
		PastureFodderBiohazard:            newPastureFodderBiohazard(db, opts...),
		PastureFodderPhysicalHazard:       newPastureFodderPhysicalHazard(db, opts...),
		PastureHouse:                      newPastureHouse(db, opts...),
		PastureProcedure:                  newPastureProcedure(db, opts...),
		PastureSiteDisinfectionRecord:     newPastureSiteDisinfectionRecord(db, opts...),
		PastureSoil:                       newPastureSoil(db, opts...),
		PastureSoilBiohazard:              newPastureSoilBiohazard(db, opts...),
		PastureSoilPhysicalHazard:         newPastureSoilPhysicalHazard(db, opts...),
		PastureTruckDisinfectionRecord:    newPastureTruckDisinfectionRecord(db, opts...),
		PastureWarehouse:                  newPastureWarehouse(db, opts...),
		PastureWasteDischarge:             newPastureWasteDischarge(db, opts...),
		PastureWater:                      newPastureWater(db, opts...),
		PastureWaterBiohazard:             newPastureWaterBiohazard(db, opts...),
		PastureWaterChemicalHazard:        newPastureWaterChemicalHazard(db, opts...),
		PastureWaterPhysicalHazard:        newPastureWaterPhysicalHazard(db, opts...),
		PastureWaterSensoryTraits:         newPastureWaterSensoryTraits(db, opts...),
		PastureWorksuitDisinfectionRecord: newPastureWorksuitDisinfectionRecord(db, opts...),
		Procedure:                         newProcedure(db, opts...),
		SlaughterBatch:                    newSlaughterBatch(db, opts...),
		SlaughterHouse:                    newSlaughterHouse(db, opts...),
		SlaughterProduct:                  newSlaughterProduct(db, opts...),
		SlaughterReceiveRecord:            newSlaughterReceiveRecord(db, opts...),
		SlaughterWarehouse:                newSlaughterWarehouse(db, opts...),
		TransportProcedureData:            newTransportProcedureData(db, opts...),
		TransportVehicle:                  newTransportVehicle(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Cow                               cow
	FSIMSUser                         fSIMSUser
	FattenProcedure                   fattenProcedure
	FattenSoil                        fattenSoil
	FattenSoilBiohazard               fattenSoilBiohazard
	FattenSoilPhysicalHazard          fattenSoilPhysicalHazard
	FattenWater                       fattenWater
	FattenWaterBiohazard              fattenWaterBiohazard
	FattenWaterChemicalHazard         fattenWaterChemicalHazard
	FattenWaterPhysicalHazard         fattenWaterPhysicalHazard
	FattenWaterSensoryTraits          fattenWaterSensoryTraits
	FeedingBatch                      feedingBatch
	Log                               log
	PackWarehouse                     packWarehouse
	PackageBatch                      packageBatch
	PackageHouse                      packageHouse
	PackageReceiveRecord              packageReceiveRecord
	PastureAir                        pastureAir
	PastureFloorBedding               pastureFloorBedding
	PastureFloorBeddingBiohazard      pastureFloorBeddingBiohazard
	PastureFloorBeddingPhysicalHazard pastureFloorBeddingPhysicalHazard
	PastureFodder                     pastureFodder
	PastureFodderBiohazard            pastureFodderBiohazard
	PastureFodderPhysicalHazard       pastureFodderPhysicalHazard
	PastureHouse                      pastureHouse
	PastureProcedure                  pastureProcedure
	PastureSiteDisinfectionRecord     pastureSiteDisinfectionRecord
	PastureSoil                       pastureSoil
	PastureSoilBiohazard              pastureSoilBiohazard
	PastureSoilPhysicalHazard         pastureSoilPhysicalHazard
	PastureTruckDisinfectionRecord    pastureTruckDisinfectionRecord
	PastureWarehouse                  pastureWarehouse
	PastureWasteDischarge             pastureWasteDischarge
	PastureWater                      pastureWater
	PastureWaterBiohazard             pastureWaterBiohazard
	PastureWaterChemicalHazard        pastureWaterChemicalHazard
	PastureWaterPhysicalHazard        pastureWaterPhysicalHazard
	PastureWaterSensoryTraits         pastureWaterSensoryTraits
	PastureWorksuitDisinfectionRecord pastureWorksuitDisinfectionRecord
	Procedure                         procedure
	SlaughterBatch                    slaughterBatch
	SlaughterHouse                    slaughterHouse
	SlaughterProduct                  slaughterProduct
	SlaughterReceiveRecord            slaughterReceiveRecord
	SlaughterWarehouse                slaughterWarehouse
	TransportProcedureData            transportProcedureData
	TransportVehicle                  transportVehicle
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                                db,
		Cow:                               q.Cow.clone(db),
		FSIMSUser:                         q.FSIMSUser.clone(db),
		FattenProcedure:                   q.FattenProcedure.clone(db),
		FattenSoil:                        q.FattenSoil.clone(db),
		FattenSoilBiohazard:               q.FattenSoilBiohazard.clone(db),
		FattenSoilPhysicalHazard:          q.FattenSoilPhysicalHazard.clone(db),
		FattenWater:                       q.FattenWater.clone(db),
		FattenWaterBiohazard:              q.FattenWaterBiohazard.clone(db),
		FattenWaterChemicalHazard:         q.FattenWaterChemicalHazard.clone(db),
		FattenWaterPhysicalHazard:         q.FattenWaterPhysicalHazard.clone(db),
		FattenWaterSensoryTraits:          q.FattenWaterSensoryTraits.clone(db),
		FeedingBatch:                      q.FeedingBatch.clone(db),
		Log:                               q.Log.clone(db),
		PackWarehouse:                     q.PackWarehouse.clone(db),
		PackageBatch:                      q.PackageBatch.clone(db),
		PackageHouse:                      q.PackageHouse.clone(db),
		PackageReceiveRecord:              q.PackageReceiveRecord.clone(db),
		PastureAir:                        q.PastureAir.clone(db),
		PastureFloorBedding:               q.PastureFloorBedding.clone(db),
		PastureFloorBeddingBiohazard:      q.PastureFloorBeddingBiohazard.clone(db),
		PastureFloorBeddingPhysicalHazard: q.PastureFloorBeddingPhysicalHazard.clone(db),
		PastureFodder:                     q.PastureFodder.clone(db),
		PastureFodderBiohazard:            q.PastureFodderBiohazard.clone(db),
		PastureFodderPhysicalHazard:       q.PastureFodderPhysicalHazard.clone(db),
		PastureHouse:                      q.PastureHouse.clone(db),
		PastureProcedure:                  q.PastureProcedure.clone(db),
		PastureSiteDisinfectionRecord:     q.PastureSiteDisinfectionRecord.clone(db),
		PastureSoil:                       q.PastureSoil.clone(db),
		PastureSoilBiohazard:              q.PastureSoilBiohazard.clone(db),
		PastureSoilPhysicalHazard:         q.PastureSoilPhysicalHazard.clone(db),
		PastureTruckDisinfectionRecord:    q.PastureTruckDisinfectionRecord.clone(db),
		PastureWarehouse:                  q.PastureWarehouse.clone(db),
		PastureWasteDischarge:             q.PastureWasteDischarge.clone(db),
		PastureWater:                      q.PastureWater.clone(db),
		PastureWaterBiohazard:             q.PastureWaterBiohazard.clone(db),
		PastureWaterChemicalHazard:        q.PastureWaterChemicalHazard.clone(db),
		PastureWaterPhysicalHazard:        q.PastureWaterPhysicalHazard.clone(db),
		PastureWaterSensoryTraits:         q.PastureWaterSensoryTraits.clone(db),
		PastureWorksuitDisinfectionRecord: q.PastureWorksuitDisinfectionRecord.clone(db),
		Procedure:                         q.Procedure.clone(db),
		SlaughterBatch:                    q.SlaughterBatch.clone(db),
		SlaughterHouse:                    q.SlaughterHouse.clone(db),
		SlaughterProduct:                  q.SlaughterProduct.clone(db),
		SlaughterReceiveRecord:            q.SlaughterReceiveRecord.clone(db),
		SlaughterWarehouse:                q.SlaughterWarehouse.clone(db),
		TransportProcedureData:            q.TransportProcedureData.clone(db),
		TransportVehicle:                  q.TransportVehicle.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                                db,
		Cow:                               q.Cow.replaceDB(db),
		FSIMSUser:                         q.FSIMSUser.replaceDB(db),
		FattenProcedure:                   q.FattenProcedure.replaceDB(db),
		FattenSoil:                        q.FattenSoil.replaceDB(db),
		FattenSoilBiohazard:               q.FattenSoilBiohazard.replaceDB(db),
		FattenSoilPhysicalHazard:          q.FattenSoilPhysicalHazard.replaceDB(db),
		FattenWater:                       q.FattenWater.replaceDB(db),
		FattenWaterBiohazard:              q.FattenWaterBiohazard.replaceDB(db),
		FattenWaterChemicalHazard:         q.FattenWaterChemicalHazard.replaceDB(db),
		FattenWaterPhysicalHazard:         q.FattenWaterPhysicalHazard.replaceDB(db),
		FattenWaterSensoryTraits:          q.FattenWaterSensoryTraits.replaceDB(db),
		FeedingBatch:                      q.FeedingBatch.replaceDB(db),
		Log:                               q.Log.replaceDB(db),
		PackWarehouse:                     q.PackWarehouse.replaceDB(db),
		PackageBatch:                      q.PackageBatch.replaceDB(db),
		PackageHouse:                      q.PackageHouse.replaceDB(db),
		PackageReceiveRecord:              q.PackageReceiveRecord.replaceDB(db),
		PastureAir:                        q.PastureAir.replaceDB(db),
		PastureFloorBedding:               q.PastureFloorBedding.replaceDB(db),
		PastureFloorBeddingBiohazard:      q.PastureFloorBeddingBiohazard.replaceDB(db),
		PastureFloorBeddingPhysicalHazard: q.PastureFloorBeddingPhysicalHazard.replaceDB(db),
		PastureFodder:                     q.PastureFodder.replaceDB(db),
		PastureFodderBiohazard:            q.PastureFodderBiohazard.replaceDB(db),
		PastureFodderPhysicalHazard:       q.PastureFodderPhysicalHazard.replaceDB(db),
		PastureHouse:                      q.PastureHouse.replaceDB(db),
		PastureProcedure:                  q.PastureProcedure.replaceDB(db),
		PastureSiteDisinfectionRecord:     q.PastureSiteDisinfectionRecord.replaceDB(db),
		PastureSoil:                       q.PastureSoil.replaceDB(db),
		PastureSoilBiohazard:              q.PastureSoilBiohazard.replaceDB(db),
		PastureSoilPhysicalHazard:         q.PastureSoilPhysicalHazard.replaceDB(db),
		PastureTruckDisinfectionRecord:    q.PastureTruckDisinfectionRecord.replaceDB(db),
		PastureWarehouse:                  q.PastureWarehouse.replaceDB(db),
		PastureWasteDischarge:             q.PastureWasteDischarge.replaceDB(db),
		PastureWater:                      q.PastureWater.replaceDB(db),
		PastureWaterBiohazard:             q.PastureWaterBiohazard.replaceDB(db),
		PastureWaterChemicalHazard:        q.PastureWaterChemicalHazard.replaceDB(db),
		PastureWaterPhysicalHazard:        q.PastureWaterPhysicalHazard.replaceDB(db),
		PastureWaterSensoryTraits:         q.PastureWaterSensoryTraits.replaceDB(db),
		PastureWorksuitDisinfectionRecord: q.PastureWorksuitDisinfectionRecord.replaceDB(db),
		Procedure:                         q.Procedure.replaceDB(db),
		SlaughterBatch:                    q.SlaughterBatch.replaceDB(db),
		SlaughterHouse:                    q.SlaughterHouse.replaceDB(db),
		SlaughterProduct:                  q.SlaughterProduct.replaceDB(db),
		SlaughterReceiveRecord:            q.SlaughterReceiveRecord.replaceDB(db),
		SlaughterWarehouse:                q.SlaughterWarehouse.replaceDB(db),
		TransportProcedureData:            q.TransportProcedureData.replaceDB(db),
		TransportVehicle:                  q.TransportVehicle.replaceDB(db),
	}
}

type queryCtx struct {
	Cow                               ICowDo
	FSIMSUser                         IFSIMSUserDo
	FattenProcedure                   IFattenProcedureDo
	FattenSoil                        IFattenSoilDo
	FattenSoilBiohazard               IFattenSoilBiohazardDo
	FattenSoilPhysicalHazard          IFattenSoilPhysicalHazardDo
	FattenWater                       IFattenWaterDo
	FattenWaterBiohazard              IFattenWaterBiohazardDo
	FattenWaterChemicalHazard         IFattenWaterChemicalHazardDo
	FattenWaterPhysicalHazard         IFattenWaterPhysicalHazardDo
	FattenWaterSensoryTraits          IFattenWaterSensoryTraitsDo
	FeedingBatch                      IFeedingBatchDo
	Log                               ILogDo
	PackWarehouse                     IPackWarehouseDo
	PackageBatch                      IPackageBatchDo
	PackageHouse                      IPackageHouseDo
	PackageReceiveRecord              IPackageReceiveRecordDo
	PastureAir                        IPastureAirDo
	PastureFloorBedding               IPastureFloorBeddingDo
	PastureFloorBeddingBiohazard      IPastureFloorBeddingBiohazardDo
	PastureFloorBeddingPhysicalHazard IPastureFloorBeddingPhysicalHazardDo
	PastureFodder                     IPastureFodderDo
	PastureFodderBiohazard            IPastureFodderBiohazardDo
	PastureFodderPhysicalHazard       IPastureFodderPhysicalHazardDo
	PastureHouse                      IPastureHouseDo
	PastureProcedure                  IPastureProcedureDo
	PastureSiteDisinfectionRecord     IPastureSiteDisinfectionRecordDo
	PastureSoil                       IPastureSoilDo
	PastureSoilBiohazard              IPastureSoilBiohazardDo
	PastureSoilPhysicalHazard         IPastureSoilPhysicalHazardDo
	PastureTruckDisinfectionRecord    IPastureTruckDisinfectionRecordDo
	PastureWarehouse                  IPastureWarehouseDo
	PastureWasteDischarge             IPastureWasteDischargeDo
	PastureWater                      IPastureWaterDo
	PastureWaterBiohazard             IPastureWaterBiohazardDo
	PastureWaterChemicalHazard        IPastureWaterChemicalHazardDo
	PastureWaterPhysicalHazard        IPastureWaterPhysicalHazardDo
	PastureWaterSensoryTraits         IPastureWaterSensoryTraitsDo
	PastureWorksuitDisinfectionRecord IPastureWorksuitDisinfectionRecordDo
	Procedure                         IProcedureDo
	SlaughterBatch                    ISlaughterBatchDo
	SlaughterHouse                    ISlaughterHouseDo
	SlaughterProduct                  ISlaughterProductDo
	SlaughterReceiveRecord            ISlaughterReceiveRecordDo
	SlaughterWarehouse                ISlaughterWarehouseDo
	TransportProcedureData            ITransportProcedureDataDo
	TransportVehicle                  ITransportVehicleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Cow:                               q.Cow.WithContext(ctx),
		FSIMSUser:                         q.FSIMSUser.WithContext(ctx),
		FattenProcedure:                   q.FattenProcedure.WithContext(ctx),
		FattenSoil:                        q.FattenSoil.WithContext(ctx),
		FattenSoilBiohazard:               q.FattenSoilBiohazard.WithContext(ctx),
		FattenSoilPhysicalHazard:          q.FattenSoilPhysicalHazard.WithContext(ctx),
		FattenWater:                       q.FattenWater.WithContext(ctx),
		FattenWaterBiohazard:              q.FattenWaterBiohazard.WithContext(ctx),
		FattenWaterChemicalHazard:         q.FattenWaterChemicalHazard.WithContext(ctx),
		FattenWaterPhysicalHazard:         q.FattenWaterPhysicalHazard.WithContext(ctx),
		FattenWaterSensoryTraits:          q.FattenWaterSensoryTraits.WithContext(ctx),
		FeedingBatch:                      q.FeedingBatch.WithContext(ctx),
		Log:                               q.Log.WithContext(ctx),
		PackWarehouse:                     q.PackWarehouse.WithContext(ctx),
		PackageBatch:                      q.PackageBatch.WithContext(ctx),
		PackageHouse:                      q.PackageHouse.WithContext(ctx),
		PackageReceiveRecord:              q.PackageReceiveRecord.WithContext(ctx),
		PastureAir:                        q.PastureAir.WithContext(ctx),
		PastureFloorBedding:               q.PastureFloorBedding.WithContext(ctx),
		PastureFloorBeddingBiohazard:      q.PastureFloorBeddingBiohazard.WithContext(ctx),
		PastureFloorBeddingPhysicalHazard: q.PastureFloorBeddingPhysicalHazard.WithContext(ctx),
		PastureFodder:                     q.PastureFodder.WithContext(ctx),
		PastureFodderBiohazard:            q.PastureFodderBiohazard.WithContext(ctx),
		PastureFodderPhysicalHazard:       q.PastureFodderPhysicalHazard.WithContext(ctx),
		PastureHouse:                      q.PastureHouse.WithContext(ctx),
		PastureProcedure:                  q.PastureProcedure.WithContext(ctx),
		PastureSiteDisinfectionRecord:     q.PastureSiteDisinfectionRecord.WithContext(ctx),
		PastureSoil:                       q.PastureSoil.WithContext(ctx),
		PastureSoilBiohazard:              q.PastureSoilBiohazard.WithContext(ctx),
		PastureSoilPhysicalHazard:         q.PastureSoilPhysicalHazard.WithContext(ctx),
		PastureTruckDisinfectionRecord:    q.PastureTruckDisinfectionRecord.WithContext(ctx),
		PastureWarehouse:                  q.PastureWarehouse.WithContext(ctx),
		PastureWasteDischarge:             q.PastureWasteDischarge.WithContext(ctx),
		PastureWater:                      q.PastureWater.WithContext(ctx),
		PastureWaterBiohazard:             q.PastureWaterBiohazard.WithContext(ctx),
		PastureWaterChemicalHazard:        q.PastureWaterChemicalHazard.WithContext(ctx),
		PastureWaterPhysicalHazard:        q.PastureWaterPhysicalHazard.WithContext(ctx),
		PastureWaterSensoryTraits:         q.PastureWaterSensoryTraits.WithContext(ctx),
		PastureWorksuitDisinfectionRecord: q.PastureWorksuitDisinfectionRecord.WithContext(ctx),
		Procedure:                         q.Procedure.WithContext(ctx),
		SlaughterBatch:                    q.SlaughterBatch.WithContext(ctx),
		SlaughterHouse:                    q.SlaughterHouse.WithContext(ctx),
		SlaughterProduct:                  q.SlaughterProduct.WithContext(ctx),
		SlaughterReceiveRecord:            q.SlaughterReceiveRecord.WithContext(ctx),
		SlaughterWarehouse:                q.SlaughterWarehouse.WithContext(ctx),
		TransportProcedureData:            q.TransportProcedureData.WithContext(ctx),
		TransportVehicle:                  q.TransportVehicle.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
