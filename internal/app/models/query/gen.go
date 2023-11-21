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
	PastureAir                        *pastureAir
	PastureFloorBedding               *pastureFloorBedding
	PastureFloorBeddingBiohazard      *pastureFloorBeddingBiohazard
	PastureFloorBeddingPhysicalHazard *pastureFloorBeddingPhysicalHazard
	PastureFodder                     *pastureFodder
	PastureFodderBiohazard            *pastureFodderBiohazard
	PastureFodderPhysicalHazard       *pastureFodderPhysicalHazard
	PastureProcedure                  *pastureProcedure
	PastureSiteDisinfectionRecord     *pastureSiteDisinfectionRecord
	PastureSoil                       *pastureSoil
	PastureSoilBiohazard              *pastureSoilBiohazard
	PastureSoilPhysicalHazard         *pastureSoilPhysicalHazard
	PastureTruckDisinfectionRecord    *pastureTruckDisinfectionRecord
	PastureWasteDischarge             *pastureWasteDischarge
	PastureWater                      *pastureWater
	PastureWaterBiohazard             *pastureWaterBiohazard
	PastureWaterChemicalHazard        *pastureWaterChemicalHazard
	PastureWaterPhysicalHazard        *pastureWaterPhysicalHazard
	PastureWaterSensoryTraits         *pastureWaterSensoryTraits
	PastureWorksuitDisinfectionRecord *pastureWorksuitDisinfectionRecord
	Procedure                         *procedure
	SubProcedure                      *subProcedure
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
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
	PastureAir = &Q.PastureAir
	PastureFloorBedding = &Q.PastureFloorBedding
	PastureFloorBeddingBiohazard = &Q.PastureFloorBeddingBiohazard
	PastureFloorBeddingPhysicalHazard = &Q.PastureFloorBeddingPhysicalHazard
	PastureFodder = &Q.PastureFodder
	PastureFodderBiohazard = &Q.PastureFodderBiohazard
	PastureFodderPhysicalHazard = &Q.PastureFodderPhysicalHazard
	PastureProcedure = &Q.PastureProcedure
	PastureSiteDisinfectionRecord = &Q.PastureSiteDisinfectionRecord
	PastureSoil = &Q.PastureSoil
	PastureSoilBiohazard = &Q.PastureSoilBiohazard
	PastureSoilPhysicalHazard = &Q.PastureSoilPhysicalHazard
	PastureTruckDisinfectionRecord = &Q.PastureTruckDisinfectionRecord
	PastureWasteDischarge = &Q.PastureWasteDischarge
	PastureWater = &Q.PastureWater
	PastureWaterBiohazard = &Q.PastureWaterBiohazard
	PastureWaterChemicalHazard = &Q.PastureWaterChemicalHazard
	PastureWaterPhysicalHazard = &Q.PastureWaterPhysicalHazard
	PastureWaterSensoryTraits = &Q.PastureWaterSensoryTraits
	PastureWorksuitDisinfectionRecord = &Q.PastureWorksuitDisinfectionRecord
	Procedure = &Q.Procedure
	SubProcedure = &Q.SubProcedure
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                                db,
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
		PastureAir:                        newPastureAir(db, opts...),
		PastureFloorBedding:               newPastureFloorBedding(db, opts...),
		PastureFloorBeddingBiohazard:      newPastureFloorBeddingBiohazard(db, opts...),
		PastureFloorBeddingPhysicalHazard: newPastureFloorBeddingPhysicalHazard(db, opts...),
		PastureFodder:                     newPastureFodder(db, opts...),
		PastureFodderBiohazard:            newPastureFodderBiohazard(db, opts...),
		PastureFodderPhysicalHazard:       newPastureFodderPhysicalHazard(db, opts...),
		PastureProcedure:                  newPastureProcedure(db, opts...),
		PastureSiteDisinfectionRecord:     newPastureSiteDisinfectionRecord(db, opts...),
		PastureSoil:                       newPastureSoil(db, opts...),
		PastureSoilBiohazard:              newPastureSoilBiohazard(db, opts...),
		PastureSoilPhysicalHazard:         newPastureSoilPhysicalHazard(db, opts...),
		PastureTruckDisinfectionRecord:    newPastureTruckDisinfectionRecord(db, opts...),
		PastureWasteDischarge:             newPastureWasteDischarge(db, opts...),
		PastureWater:                      newPastureWater(db, opts...),
		PastureWaterBiohazard:             newPastureWaterBiohazard(db, opts...),
		PastureWaterChemicalHazard:        newPastureWaterChemicalHazard(db, opts...),
		PastureWaterPhysicalHazard:        newPastureWaterPhysicalHazard(db, opts...),
		PastureWaterSensoryTraits:         newPastureWaterSensoryTraits(db, opts...),
		PastureWorksuitDisinfectionRecord: newPastureWorksuitDisinfectionRecord(db, opts...),
		Procedure:                         newProcedure(db, opts...),
		SubProcedure:                      newSubProcedure(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

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
	PastureAir                        pastureAir
	PastureFloorBedding               pastureFloorBedding
	PastureFloorBeddingBiohazard      pastureFloorBeddingBiohazard
	PastureFloorBeddingPhysicalHazard pastureFloorBeddingPhysicalHazard
	PastureFodder                     pastureFodder
	PastureFodderBiohazard            pastureFodderBiohazard
	PastureFodderPhysicalHazard       pastureFodderPhysicalHazard
	PastureProcedure                  pastureProcedure
	PastureSiteDisinfectionRecord     pastureSiteDisinfectionRecord
	PastureSoil                       pastureSoil
	PastureSoilBiohazard              pastureSoilBiohazard
	PastureSoilPhysicalHazard         pastureSoilPhysicalHazard
	PastureTruckDisinfectionRecord    pastureTruckDisinfectionRecord
	PastureWasteDischarge             pastureWasteDischarge
	PastureWater                      pastureWater
	PastureWaterBiohazard             pastureWaterBiohazard
	PastureWaterChemicalHazard        pastureWaterChemicalHazard
	PastureWaterPhysicalHazard        pastureWaterPhysicalHazard
	PastureWaterSensoryTraits         pastureWaterSensoryTraits
	PastureWorksuitDisinfectionRecord pastureWorksuitDisinfectionRecord
	Procedure                         procedure
	SubProcedure                      subProcedure
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                                db,
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
		PastureAir:                        q.PastureAir.clone(db),
		PastureFloorBedding:               q.PastureFloorBedding.clone(db),
		PastureFloorBeddingBiohazard:      q.PastureFloorBeddingBiohazard.clone(db),
		PastureFloorBeddingPhysicalHazard: q.PastureFloorBeddingPhysicalHazard.clone(db),
		PastureFodder:                     q.PastureFodder.clone(db),
		PastureFodderBiohazard:            q.PastureFodderBiohazard.clone(db),
		PastureFodderPhysicalHazard:       q.PastureFodderPhysicalHazard.clone(db),
		PastureProcedure:                  q.PastureProcedure.clone(db),
		PastureSiteDisinfectionRecord:     q.PastureSiteDisinfectionRecord.clone(db),
		PastureSoil:                       q.PastureSoil.clone(db),
		PastureSoilBiohazard:              q.PastureSoilBiohazard.clone(db),
		PastureSoilPhysicalHazard:         q.PastureSoilPhysicalHazard.clone(db),
		PastureTruckDisinfectionRecord:    q.PastureTruckDisinfectionRecord.clone(db),
		PastureWasteDischarge:             q.PastureWasteDischarge.clone(db),
		PastureWater:                      q.PastureWater.clone(db),
		PastureWaterBiohazard:             q.PastureWaterBiohazard.clone(db),
		PastureWaterChemicalHazard:        q.PastureWaterChemicalHazard.clone(db),
		PastureWaterPhysicalHazard:        q.PastureWaterPhysicalHazard.clone(db),
		PastureWaterSensoryTraits:         q.PastureWaterSensoryTraits.clone(db),
		PastureWorksuitDisinfectionRecord: q.PastureWorksuitDisinfectionRecord.clone(db),
		Procedure:                         q.Procedure.clone(db),
		SubProcedure:                      q.SubProcedure.clone(db),
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
		PastureAir:                        q.PastureAir.replaceDB(db),
		PastureFloorBedding:               q.PastureFloorBedding.replaceDB(db),
		PastureFloorBeddingBiohazard:      q.PastureFloorBeddingBiohazard.replaceDB(db),
		PastureFloorBeddingPhysicalHazard: q.PastureFloorBeddingPhysicalHazard.replaceDB(db),
		PastureFodder:                     q.PastureFodder.replaceDB(db),
		PastureFodderBiohazard:            q.PastureFodderBiohazard.replaceDB(db),
		PastureFodderPhysicalHazard:       q.PastureFodderPhysicalHazard.replaceDB(db),
		PastureProcedure:                  q.PastureProcedure.replaceDB(db),
		PastureSiteDisinfectionRecord:     q.PastureSiteDisinfectionRecord.replaceDB(db),
		PastureSoil:                       q.PastureSoil.replaceDB(db),
		PastureSoilBiohazard:              q.PastureSoilBiohazard.replaceDB(db),
		PastureSoilPhysicalHazard:         q.PastureSoilPhysicalHazard.replaceDB(db),
		PastureTruckDisinfectionRecord:    q.PastureTruckDisinfectionRecord.replaceDB(db),
		PastureWasteDischarge:             q.PastureWasteDischarge.replaceDB(db),
		PastureWater:                      q.PastureWater.replaceDB(db),
		PastureWaterBiohazard:             q.PastureWaterBiohazard.replaceDB(db),
		PastureWaterChemicalHazard:        q.PastureWaterChemicalHazard.replaceDB(db),
		PastureWaterPhysicalHazard:        q.PastureWaterPhysicalHazard.replaceDB(db),
		PastureWaterSensoryTraits:         q.PastureWaterSensoryTraits.replaceDB(db),
		PastureWorksuitDisinfectionRecord: q.PastureWorksuitDisinfectionRecord.replaceDB(db),
		Procedure:                         q.Procedure.replaceDB(db),
		SubProcedure:                      q.SubProcedure.replaceDB(db),
	}
}

type queryCtx struct {
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
	PastureAir                        IPastureAirDo
	PastureFloorBedding               IPastureFloorBeddingDo
	PastureFloorBeddingBiohazard      IPastureFloorBeddingBiohazardDo
	PastureFloorBeddingPhysicalHazard IPastureFloorBeddingPhysicalHazardDo
	PastureFodder                     IPastureFodderDo
	PastureFodderBiohazard            IPastureFodderBiohazardDo
	PastureFodderPhysicalHazard       IPastureFodderPhysicalHazardDo
	PastureProcedure                  IPastureProcedureDo
	PastureSiteDisinfectionRecord     IPastureSiteDisinfectionRecordDo
	PastureSoil                       IPastureSoilDo
	PastureSoilBiohazard              IPastureSoilBiohazardDo
	PastureSoilPhysicalHazard         IPastureSoilPhysicalHazardDo
	PastureTruckDisinfectionRecord    IPastureTruckDisinfectionRecordDo
	PastureWasteDischarge             IPastureWasteDischargeDo
	PastureWater                      IPastureWaterDo
	PastureWaterBiohazard             IPastureWaterBiohazardDo
	PastureWaterChemicalHazard        IPastureWaterChemicalHazardDo
	PastureWaterPhysicalHazard        IPastureWaterPhysicalHazardDo
	PastureWaterSensoryTraits         IPastureWaterSensoryTraitsDo
	PastureWorksuitDisinfectionRecord IPastureWorksuitDisinfectionRecordDo
	Procedure                         IProcedureDo
	SubProcedure                      ISubProcedureDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
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
		PastureAir:                        q.PastureAir.WithContext(ctx),
		PastureFloorBedding:               q.PastureFloorBedding.WithContext(ctx),
		PastureFloorBeddingBiohazard:      q.PastureFloorBeddingBiohazard.WithContext(ctx),
		PastureFloorBeddingPhysicalHazard: q.PastureFloorBeddingPhysicalHazard.WithContext(ctx),
		PastureFodder:                     q.PastureFodder.WithContext(ctx),
		PastureFodderBiohazard:            q.PastureFodderBiohazard.WithContext(ctx),
		PastureFodderPhysicalHazard:       q.PastureFodderPhysicalHazard.WithContext(ctx),
		PastureProcedure:                  q.PastureProcedure.WithContext(ctx),
		PastureSiteDisinfectionRecord:     q.PastureSiteDisinfectionRecord.WithContext(ctx),
		PastureSoil:                       q.PastureSoil.WithContext(ctx),
		PastureSoilBiohazard:              q.PastureSoilBiohazard.WithContext(ctx),
		PastureSoilPhysicalHazard:         q.PastureSoilPhysicalHazard.WithContext(ctx),
		PastureTruckDisinfectionRecord:    q.PastureTruckDisinfectionRecord.WithContext(ctx),
		PastureWasteDischarge:             q.PastureWasteDischarge.WithContext(ctx),
		PastureWater:                      q.PastureWater.WithContext(ctx),
		PastureWaterBiohazard:             q.PastureWaterBiohazard.WithContext(ctx),
		PastureWaterChemicalHazard:        q.PastureWaterChemicalHazard.WithContext(ctx),
		PastureWaterPhysicalHazard:        q.PastureWaterPhysicalHazard.WithContext(ctx),
		PastureWaterSensoryTraits:         q.PastureWaterSensoryTraits.WithContext(ctx),
		PastureWorksuitDisinfectionRecord: q.PastureWorksuitDisinfectionRecord.WithContext(ctx),
		Procedure:                         q.Procedure.WithContext(ctx),
		SubProcedure:                      q.SubProcedure.WithContext(ctx),
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
