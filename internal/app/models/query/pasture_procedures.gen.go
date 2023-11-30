// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPastureProcedure(db *gorm.DB, opts ...gen.DOOption) pastureProcedure {
	_pastureProcedure := pastureProcedure{}

	_pastureProcedure.pastureProcedureDo.UseDB(db, opts...)
	_pastureProcedure.pastureProcedureDo.UseModel(&pasture.PastureProcedure{})

	tableName := _pastureProcedure.pastureProcedureDo.TableName()
	_pastureProcedure.ALL = field.NewAsterisk(tableName)
	_pastureProcedure.ID = field.NewUint(tableName, "id")
	_pastureProcedure.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureProcedure.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureProcedure.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureProcedure.PasPID = field.NewString(tableName, "pas_p_id")
	_pastureProcedure.Water = pastureProcedureHasOneWater{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Water", "pasture.PastureWater"),
		PhysicalHazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Water.PhysicalHazard", "pasture.PastureWaterPhysicalHazard"),
		},
		ChemicalHazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Water.ChemicalHazard", "pasture.PastureWaterChemicalHazard"),
		},
		Biohazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Water.Biohazard", "pasture.PastureWaterBiohazard"),
		},
		SensoryTraits: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Water.SensoryTraits", "pasture.PastureWaterSensoryTraits"),
		},
	}

	_pastureProcedure.Fodder = pastureProcedureHasOneFodder{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Fodder", "pasture.PastureFodder"),
		PhysicalHazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Fodder.PhysicalHazard", "pasture.PastureFodderPhysicalHazard"),
		},
		Biohazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Fodder.Biohazard", "pasture.PastureFodderBiohazard"),
		},
	}

	_pastureProcedure.Soil = pastureProcedureHasOneSoil{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Soil", "pasture.PastureSoil"),
		PhysicalHazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Soil.PhysicalHazard", "pasture.PastureSoilPhysicalHazard"),
		},
		Biohazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Soil.Biohazard", "pasture.PastureSoilBiohazard"),
		},
	}

	_pastureProcedure.Air = pastureProcedureHasOneAir{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Air", "pasture.PastureAir"),
	}

	_pastureProcedure.FloorBedding = pastureProcedureHasOneFloorBedding{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("FloorBedding", "pasture.PastureFloorBedding"),
		PhysicalHazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("FloorBedding.PhysicalHazard", "pasture.PastureFloorBeddingPhysicalHazard"),
		},
		Biohazard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("FloorBedding.Biohazard", "pasture.PastureFloorBeddingBiohazard"),
		},
	}

	_pastureProcedure.WasteDischarge = pastureProcedureHasOneWasteDischarge{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("WasteDischarge", "pasture.PastureWasteDischarge"),
	}

	_pastureProcedure.fillFieldMap()

	return _pastureProcedure
}

type pastureProcedure struct {
	pastureProcedureDo pastureProcedureDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	PasPID    field.String
	Water     pastureProcedureHasOneWater

	Fodder pastureProcedureHasOneFodder

	Soil pastureProcedureHasOneSoil

	Air pastureProcedureHasOneAir

	FloorBedding pastureProcedureHasOneFloorBedding

	WasteDischarge pastureProcedureHasOneWasteDischarge

	fieldMap map[string]field.Expr
}

func (p pastureProcedure) Table(newTableName string) *pastureProcedure {
	p.pastureProcedureDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureProcedure) As(alias string) *pastureProcedure {
	p.pastureProcedureDo.DO = *(p.pastureProcedureDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureProcedure) updateTableName(table string) *pastureProcedure {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.PasPID = field.NewString(table, "pas_p_id")

	p.fillFieldMap()

	return p
}

func (p *pastureProcedure) WithContext(ctx context.Context) IPastureProcedureDo {
	return p.pastureProcedureDo.WithContext(ctx)
}

func (p pastureProcedure) TableName() string { return p.pastureProcedureDo.TableName() }

func (p pastureProcedure) Alias() string { return p.pastureProcedureDo.Alias() }

func (p pastureProcedure) Columns(cols ...field.Expr) gen.Columns {
	return p.pastureProcedureDo.Columns(cols...)
}

func (p *pastureProcedure) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureProcedure) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["pas_p_id"] = p.PasPID

}

func (p pastureProcedure) clone(db *gorm.DB) pastureProcedure {
	p.pastureProcedureDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureProcedure) replaceDB(db *gorm.DB) pastureProcedure {
	p.pastureProcedureDo.ReplaceDB(db)
	return p
}

type pastureProcedureHasOneWater struct {
	db *gorm.DB

	field.RelationField

	PhysicalHazard struct {
		field.RelationField
	}
	ChemicalHazard struct {
		field.RelationField
	}
	Biohazard struct {
		field.RelationField
	}
	SensoryTraits struct {
		field.RelationField
	}
}

func (a pastureProcedureHasOneWater) Where(conds ...field.Expr) *pastureProcedureHasOneWater {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a pastureProcedureHasOneWater) WithContext(ctx context.Context) *pastureProcedureHasOneWater {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pastureProcedureHasOneWater) Session(session *gorm.Session) *pastureProcedureHasOneWater {
	a.db = a.db.Session(session)
	return &a
}

func (a pastureProcedureHasOneWater) Model(m *pasture.PastureProcedure) *pastureProcedureHasOneWaterTx {
	return &pastureProcedureHasOneWaterTx{a.db.Model(m).Association(a.Name())}
}

type pastureProcedureHasOneWaterTx struct{ tx *gorm.Association }

func (a pastureProcedureHasOneWaterTx) Find() (result *pasture.PastureWater, err error) {
	return result, a.tx.Find(&result)
}

func (a pastureProcedureHasOneWaterTx) Append(values ...*pasture.PastureWater) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pastureProcedureHasOneWaterTx) Replace(values ...*pasture.PastureWater) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pastureProcedureHasOneWaterTx) Delete(values ...*pasture.PastureWater) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pastureProcedureHasOneWaterTx) Clear() error {
	return a.tx.Clear()
}

func (a pastureProcedureHasOneWaterTx) Count() int64 {
	return a.tx.Count()
}

type pastureProcedureHasOneFodder struct {
	db *gorm.DB

	field.RelationField

	PhysicalHazard struct {
		field.RelationField
	}
	Biohazard struct {
		field.RelationField
	}
}

func (a pastureProcedureHasOneFodder) Where(conds ...field.Expr) *pastureProcedureHasOneFodder {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a pastureProcedureHasOneFodder) WithContext(ctx context.Context) *pastureProcedureHasOneFodder {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pastureProcedureHasOneFodder) Session(session *gorm.Session) *pastureProcedureHasOneFodder {
	a.db = a.db.Session(session)
	return &a
}

func (a pastureProcedureHasOneFodder) Model(m *pasture.PastureProcedure) *pastureProcedureHasOneFodderTx {
	return &pastureProcedureHasOneFodderTx{a.db.Model(m).Association(a.Name())}
}

type pastureProcedureHasOneFodderTx struct{ tx *gorm.Association }

func (a pastureProcedureHasOneFodderTx) Find() (result *pasture.PastureFodder, err error) {
	return result, a.tx.Find(&result)
}

func (a pastureProcedureHasOneFodderTx) Append(values ...*pasture.PastureFodder) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pastureProcedureHasOneFodderTx) Replace(values ...*pasture.PastureFodder) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pastureProcedureHasOneFodderTx) Delete(values ...*pasture.PastureFodder) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pastureProcedureHasOneFodderTx) Clear() error {
	return a.tx.Clear()
}

func (a pastureProcedureHasOneFodderTx) Count() int64 {
	return a.tx.Count()
}

type pastureProcedureHasOneSoil struct {
	db *gorm.DB

	field.RelationField

	PhysicalHazard struct {
		field.RelationField
	}
	Biohazard struct {
		field.RelationField
	}
}

func (a pastureProcedureHasOneSoil) Where(conds ...field.Expr) *pastureProcedureHasOneSoil {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a pastureProcedureHasOneSoil) WithContext(ctx context.Context) *pastureProcedureHasOneSoil {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pastureProcedureHasOneSoil) Session(session *gorm.Session) *pastureProcedureHasOneSoil {
	a.db = a.db.Session(session)
	return &a
}

func (a pastureProcedureHasOneSoil) Model(m *pasture.PastureProcedure) *pastureProcedureHasOneSoilTx {
	return &pastureProcedureHasOneSoilTx{a.db.Model(m).Association(a.Name())}
}

type pastureProcedureHasOneSoilTx struct{ tx *gorm.Association }

func (a pastureProcedureHasOneSoilTx) Find() (result *pasture.PastureSoil, err error) {
	return result, a.tx.Find(&result)
}

func (a pastureProcedureHasOneSoilTx) Append(values ...*pasture.PastureSoil) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pastureProcedureHasOneSoilTx) Replace(values ...*pasture.PastureSoil) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pastureProcedureHasOneSoilTx) Delete(values ...*pasture.PastureSoil) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pastureProcedureHasOneSoilTx) Clear() error {
	return a.tx.Clear()
}

func (a pastureProcedureHasOneSoilTx) Count() int64 {
	return a.tx.Count()
}

type pastureProcedureHasOneAir struct {
	db *gorm.DB

	field.RelationField
}

func (a pastureProcedureHasOneAir) Where(conds ...field.Expr) *pastureProcedureHasOneAir {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a pastureProcedureHasOneAir) WithContext(ctx context.Context) *pastureProcedureHasOneAir {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pastureProcedureHasOneAir) Session(session *gorm.Session) *pastureProcedureHasOneAir {
	a.db = a.db.Session(session)
	return &a
}

func (a pastureProcedureHasOneAir) Model(m *pasture.PastureProcedure) *pastureProcedureHasOneAirTx {
	return &pastureProcedureHasOneAirTx{a.db.Model(m).Association(a.Name())}
}

type pastureProcedureHasOneAirTx struct{ tx *gorm.Association }

func (a pastureProcedureHasOneAirTx) Find() (result *pasture.PastureAir, err error) {
	return result, a.tx.Find(&result)
}

func (a pastureProcedureHasOneAirTx) Append(values ...*pasture.PastureAir) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pastureProcedureHasOneAirTx) Replace(values ...*pasture.PastureAir) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pastureProcedureHasOneAirTx) Delete(values ...*pasture.PastureAir) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pastureProcedureHasOneAirTx) Clear() error {
	return a.tx.Clear()
}

func (a pastureProcedureHasOneAirTx) Count() int64 {
	return a.tx.Count()
}

type pastureProcedureHasOneFloorBedding struct {
	db *gorm.DB

	field.RelationField

	PhysicalHazard struct {
		field.RelationField
	}
	Biohazard struct {
		field.RelationField
	}
}

func (a pastureProcedureHasOneFloorBedding) Where(conds ...field.Expr) *pastureProcedureHasOneFloorBedding {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a pastureProcedureHasOneFloorBedding) WithContext(ctx context.Context) *pastureProcedureHasOneFloorBedding {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pastureProcedureHasOneFloorBedding) Session(session *gorm.Session) *pastureProcedureHasOneFloorBedding {
	a.db = a.db.Session(session)
	return &a
}

func (a pastureProcedureHasOneFloorBedding) Model(m *pasture.PastureProcedure) *pastureProcedureHasOneFloorBeddingTx {
	return &pastureProcedureHasOneFloorBeddingTx{a.db.Model(m).Association(a.Name())}
}

type pastureProcedureHasOneFloorBeddingTx struct{ tx *gorm.Association }

func (a pastureProcedureHasOneFloorBeddingTx) Find() (result *pasture.PastureFloorBedding, err error) {
	return result, a.tx.Find(&result)
}

func (a pastureProcedureHasOneFloorBeddingTx) Append(values ...*pasture.PastureFloorBedding) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pastureProcedureHasOneFloorBeddingTx) Replace(values ...*pasture.PastureFloorBedding) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pastureProcedureHasOneFloorBeddingTx) Delete(values ...*pasture.PastureFloorBedding) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pastureProcedureHasOneFloorBeddingTx) Clear() error {
	return a.tx.Clear()
}

func (a pastureProcedureHasOneFloorBeddingTx) Count() int64 {
	return a.tx.Count()
}

type pastureProcedureHasOneWasteDischarge struct {
	db *gorm.DB

	field.RelationField
}

func (a pastureProcedureHasOneWasteDischarge) Where(conds ...field.Expr) *pastureProcedureHasOneWasteDischarge {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a pastureProcedureHasOneWasteDischarge) WithContext(ctx context.Context) *pastureProcedureHasOneWasteDischarge {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pastureProcedureHasOneWasteDischarge) Session(session *gorm.Session) *pastureProcedureHasOneWasteDischarge {
	a.db = a.db.Session(session)
	return &a
}

func (a pastureProcedureHasOneWasteDischarge) Model(m *pasture.PastureProcedure) *pastureProcedureHasOneWasteDischargeTx {
	return &pastureProcedureHasOneWasteDischargeTx{a.db.Model(m).Association(a.Name())}
}

type pastureProcedureHasOneWasteDischargeTx struct{ tx *gorm.Association }

func (a pastureProcedureHasOneWasteDischargeTx) Find() (result *pasture.PastureWasteDischarge, err error) {
	return result, a.tx.Find(&result)
}

func (a pastureProcedureHasOneWasteDischargeTx) Append(values ...*pasture.PastureWasteDischarge) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pastureProcedureHasOneWasteDischargeTx) Replace(values ...*pasture.PastureWasteDischarge) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pastureProcedureHasOneWasteDischargeTx) Delete(values ...*pasture.PastureWasteDischarge) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pastureProcedureHasOneWasteDischargeTx) Clear() error {
	return a.tx.Clear()
}

func (a pastureProcedureHasOneWasteDischargeTx) Count() int64 {
	return a.tx.Count()
}

type pastureProcedureDo struct{ gen.DO }

type IPastureProcedureDo interface {
	gen.SubQuery
	Debug() IPastureProcedureDo
	WithContext(ctx context.Context) IPastureProcedureDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureProcedureDo
	WriteDB() IPastureProcedureDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureProcedureDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureProcedureDo
	Not(conds ...gen.Condition) IPastureProcedureDo
	Or(conds ...gen.Condition) IPastureProcedureDo
	Select(conds ...field.Expr) IPastureProcedureDo
	Where(conds ...gen.Condition) IPastureProcedureDo
	Order(conds ...field.Expr) IPastureProcedureDo
	Distinct(cols ...field.Expr) IPastureProcedureDo
	Omit(cols ...field.Expr) IPastureProcedureDo
	Join(table schema.Tabler, on ...field.Expr) IPastureProcedureDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureProcedureDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureProcedureDo
	Group(cols ...field.Expr) IPastureProcedureDo
	Having(conds ...gen.Condition) IPastureProcedureDo
	Limit(limit int) IPastureProcedureDo
	Offset(offset int) IPastureProcedureDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureProcedureDo
	Unscoped() IPastureProcedureDo
	Create(values ...*pasture.PastureProcedure) error
	CreateInBatches(values []*pasture.PastureProcedure, batchSize int) error
	Save(values ...*pasture.PastureProcedure) error
	First() (*pasture.PastureProcedure, error)
	Take() (*pasture.PastureProcedure, error)
	Last() (*pasture.PastureProcedure, error)
	Find() ([]*pasture.PastureProcedure, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureProcedure, err error)
	FindInBatches(result *[]*pasture.PastureProcedure, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureProcedure) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureProcedureDo
	Assign(attrs ...field.AssignExpr) IPastureProcedureDo
	Joins(fields ...field.RelationField) IPastureProcedureDo
	Preload(fields ...field.RelationField) IPastureProcedureDo
	FirstOrInit() (*pasture.PastureProcedure, error)
	FirstOrCreate() (*pasture.PastureProcedure, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureProcedure, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureProcedureDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureProcedureDo) Debug() IPastureProcedureDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureProcedureDo) WithContext(ctx context.Context) IPastureProcedureDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureProcedureDo) ReadDB() IPastureProcedureDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureProcedureDo) WriteDB() IPastureProcedureDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureProcedureDo) Session(config *gorm.Session) IPastureProcedureDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureProcedureDo) Clauses(conds ...clause.Expression) IPastureProcedureDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureProcedureDo) Returning(value interface{}, columns ...string) IPastureProcedureDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureProcedureDo) Not(conds ...gen.Condition) IPastureProcedureDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureProcedureDo) Or(conds ...gen.Condition) IPastureProcedureDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureProcedureDo) Select(conds ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureProcedureDo) Where(conds ...gen.Condition) IPastureProcedureDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureProcedureDo) Order(conds ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureProcedureDo) Distinct(cols ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureProcedureDo) Omit(cols ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureProcedureDo) Join(table schema.Tabler, on ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureProcedureDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureProcedureDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureProcedureDo) Group(cols ...field.Expr) IPastureProcedureDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureProcedureDo) Having(conds ...gen.Condition) IPastureProcedureDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureProcedureDo) Limit(limit int) IPastureProcedureDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureProcedureDo) Offset(offset int) IPastureProcedureDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureProcedureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureProcedureDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureProcedureDo) Unscoped() IPastureProcedureDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureProcedureDo) Create(values ...*pasture.PastureProcedure) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureProcedureDo) CreateInBatches(values []*pasture.PastureProcedure, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureProcedureDo) Save(values ...*pasture.PastureProcedure) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureProcedureDo) First() (*pasture.PastureProcedure, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureProcedure), nil
	}
}

func (p pastureProcedureDo) Take() (*pasture.PastureProcedure, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureProcedure), nil
	}
}

func (p pastureProcedureDo) Last() (*pasture.PastureProcedure, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureProcedure), nil
	}
}

func (p pastureProcedureDo) Find() ([]*pasture.PastureProcedure, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureProcedure), err
}

func (p pastureProcedureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureProcedure, err error) {
	buf := make([]*pasture.PastureProcedure, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureProcedureDo) FindInBatches(result *[]*pasture.PastureProcedure, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureProcedureDo) Attrs(attrs ...field.AssignExpr) IPastureProcedureDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureProcedureDo) Assign(attrs ...field.AssignExpr) IPastureProcedureDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureProcedureDo) Joins(fields ...field.RelationField) IPastureProcedureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureProcedureDo) Preload(fields ...field.RelationField) IPastureProcedureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureProcedureDo) FirstOrInit() (*pasture.PastureProcedure, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureProcedure), nil
	}
}

func (p pastureProcedureDo) FirstOrCreate() (*pasture.PastureProcedure, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureProcedure), nil
	}
}

func (p pastureProcedureDo) FindByPage(offset int, limit int) (result []*pasture.PastureProcedure, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pastureProcedureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureProcedureDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureProcedureDo) Delete(models ...*pasture.PastureProcedure) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureProcedureDo) withDO(do gen.Dao) *pastureProcedureDo {
	p.DO = *do.(*gen.DO)
	return p
}