// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newSlaughterToxinIndex(db *gorm.DB, opts ...gen.DOOption) slaughterToxinIndex {
	_slaughterToxinIndex := slaughterToxinIndex{}

	_slaughterToxinIndex.slaughterToxinIndexDo.UseDB(db, opts...)
	_slaughterToxinIndex.slaughterToxinIndexDo.UseModel(&slaughter.SlaughterToxinIndex{})

	tableName := _slaughterToxinIndex.slaughterToxinIndexDo.TableName()
	_slaughterToxinIndex.ALL = field.NewAsterisk(tableName)
	_slaughterToxinIndex.ID = field.NewUint(tableName, "id")
	_slaughterToxinIndex.CreatedAt = field.NewTime(tableName, "created_at")
	_slaughterToxinIndex.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaughterToxinIndex.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaughterToxinIndex.SlaughterWaterQualityMonID = field.NewUint(tableName, "slaughter_water_quality_mon_id")
	_slaughterToxinIndex.ToxinIndexSla1 = field.NewFloat64(tableName, "toxin_index_sla1")
	_slaughterToxinIndex.ToxinIndexSla2 = field.NewFloat64(tableName, "toxin_index_sla2")
	_slaughterToxinIndex.ToxinIndexSla3 = field.NewFloat64(tableName, "toxin_index_sla3")
	_slaughterToxinIndex.ToxinIndexSla4 = field.NewFloat64(tableName, "toxin_index_sla4")
	_slaughterToxinIndex.ToxinIndexSla5 = field.NewFloat64(tableName, "toxin_index_sla5")
	_slaughterToxinIndex.ToxinIndexSla6 = field.NewFloat64(tableName, "toxin_index_sla6")
	_slaughterToxinIndex.ToxinIndexSla7 = field.NewFloat64(tableName, "toxin_index_sla7")
	_slaughterToxinIndex.ToxinIndexSla8 = field.NewFloat64(tableName, "toxin_index_sla8")
	_slaughterToxinIndex.ToxinIndexSla9 = field.NewFloat64(tableName, "toxin_index_sla9")
	_slaughterToxinIndex.ToxinIndexSla10 = field.NewFloat64(tableName, "toxin_index_sla10")
	_slaughterToxinIndex.ToxinIndexSla11 = field.NewFloat64(tableName, "toxin_index_sla11")
	_slaughterToxinIndex.ToxinIndexSla12 = field.NewFloat64(tableName, "toxin_index_sla12")
	_slaughterToxinIndex.ToxinIndexSla13 = field.NewFloat64(tableName, "toxin_index_sla13")
	_slaughterToxinIndex.ToxinIndexSla14 = field.NewFloat64(tableName, "toxin_index_sla14")
	_slaughterToxinIndex.ToxinIndexSla15 = field.NewFloat64(tableName, "toxin_index_sla15")
	_slaughterToxinIndex.ToxinIndexSla16 = field.NewFloat64(tableName, "toxin_index_sla16")
	_slaughterToxinIndex.ToxinIndexSla17 = field.NewFloat64(tableName, "toxin_index_sla17")
	_slaughterToxinIndex.ToxinIndexSla18 = field.NewFloat64(tableName, "toxin_index_sla18")
	_slaughterToxinIndex.ToxinIndexSla19 = field.NewFloat64(tableName, "toxin_index_sla19")
	_slaughterToxinIndex.ToxinIndexSla20 = field.NewFloat64(tableName, "toxin_index_sla20")
	_slaughterToxinIndex.ToxinIndexSla21 = field.NewFloat64(tableName, "toxin_index_sla21")
	_slaughterToxinIndex.ToxinIndexSla22 = field.NewFloat64(tableName, "toxin_index_sla22")
	_slaughterToxinIndex.ToxinIndexSla23 = field.NewFloat64(tableName, "toxin_index_sla23")
	_slaughterToxinIndex.ToxinIndexSla24 = field.NewFloat64(tableName, "toxin_index_sla24")
	_slaughterToxinIndex.ToxinIndexSla25 = field.NewFloat64(tableName, "toxin_index_sla25")
	_slaughterToxinIndex.ToxinIndexSla26 = field.NewFloat64(tableName, "toxin_index_sla26")
	_slaughterToxinIndex.ToxinIndexSla27 = field.NewFloat64(tableName, "toxin_index_sla27")
	_slaughterToxinIndex.ToxinIndexSla28 = field.NewFloat64(tableName, "toxin_index_sla28")
	_slaughterToxinIndex.ToxinIndexSla29 = field.NewFloat64(tableName, "toxin_index_sla29")
	_slaughterToxinIndex.ToxinIndexSla30 = field.NewFloat64(tableName, "toxin_index_sla30")
	_slaughterToxinIndex.ToxinIndexSla31 = field.NewFloat64(tableName, "toxin_index_sla31")
	_slaughterToxinIndex.ToxinIndexSla32 = field.NewFloat64(tableName, "toxin_index_sla32")
	_slaughterToxinIndex.ToxinIndexSla33 = field.NewFloat64(tableName, "toxin_index_sla33")
	_slaughterToxinIndex.ToxinIndexSla34 = field.NewFloat64(tableName, "toxin_index_sla34")
	_slaughterToxinIndex.ToxinIndexSla35 = field.NewFloat64(tableName, "toxin_index_sla35")
	_slaughterToxinIndex.ToxinIndexSla36 = field.NewFloat64(tableName, "toxin_index_sla36")
	_slaughterToxinIndex.ToxinIndexSla37 = field.NewFloat64(tableName, "toxin_index_sla37")
	_slaughterToxinIndex.ToxinIndexSla38 = field.NewFloat64(tableName, "toxin_index_sla38")
	_slaughterToxinIndex.ToxinIndexSla39 = field.NewFloat64(tableName, "toxin_index_sla39")
	_slaughterToxinIndex.ToxinIndexSla40 = field.NewFloat64(tableName, "toxin_index_sla40")
	_slaughterToxinIndex.ToxinIndexSla41 = field.NewFloat64(tableName, "toxin_index_sla41")
	_slaughterToxinIndex.ToxinIndexSla42 = field.NewFloat64(tableName, "toxin_index_sla42")
	_slaughterToxinIndex.ToxinIndexSla43 = field.NewFloat64(tableName, "toxin_index_sla43")
	_slaughterToxinIndex.ToxinIndexSla44 = field.NewFloat64(tableName, "toxin_index_sla44")
	_slaughterToxinIndex.ToxinIndexSla45 = field.NewFloat64(tableName, "toxin_index_sla45")
	_slaughterToxinIndex.ToxinIndexSla46 = field.NewFloat64(tableName, "toxin_index_sla46")
	_slaughterToxinIndex.ToxinIndexSla47 = field.NewFloat64(tableName, "toxin_index_sla47")
	_slaughterToxinIndex.SlaughterWaterToxinIndex = slaughterToxinIndexHasOneSlaughterWaterToxinIndex{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SlaughterWaterToxinIndex", "slaughter.SlaughterWaterToxinIndex"),
	}

	_slaughterToxinIndex.fillFieldMap()

	return _slaughterToxinIndex
}

type slaughterToxinIndex struct {
	slaughterToxinIndexDo slaughterToxinIndexDo

	ALL                        field.Asterisk
	ID                         field.Uint
	CreatedAt                  field.Time
	UpdatedAt                  field.Time
	DeletedAt                  field.Field
	SlaughterWaterQualityMonID field.Uint
	ToxinIndexSla1             field.Float64
	ToxinIndexSla2             field.Float64
	ToxinIndexSla3             field.Float64
	ToxinIndexSla4             field.Float64
	ToxinIndexSla5             field.Float64
	ToxinIndexSla6             field.Float64
	ToxinIndexSla7             field.Float64
	ToxinIndexSla8             field.Float64
	ToxinIndexSla9             field.Float64
	ToxinIndexSla10            field.Float64
	ToxinIndexSla11            field.Float64
	ToxinIndexSla12            field.Float64
	ToxinIndexSla13            field.Float64
	ToxinIndexSla14            field.Float64
	ToxinIndexSla15            field.Float64
	ToxinIndexSla16            field.Float64
	ToxinIndexSla17            field.Float64
	ToxinIndexSla18            field.Float64
	ToxinIndexSla19            field.Float64
	ToxinIndexSla20            field.Float64
	ToxinIndexSla21            field.Float64
	ToxinIndexSla22            field.Float64
	ToxinIndexSla23            field.Float64
	ToxinIndexSla24            field.Float64
	ToxinIndexSla25            field.Float64
	ToxinIndexSla26            field.Float64
	ToxinIndexSla27            field.Float64
	ToxinIndexSla28            field.Float64
	ToxinIndexSla29            field.Float64
	ToxinIndexSla30            field.Float64
	ToxinIndexSla31            field.Float64
	ToxinIndexSla32            field.Float64
	ToxinIndexSla33            field.Float64
	ToxinIndexSla34            field.Float64
	ToxinIndexSla35            field.Float64
	ToxinIndexSla36            field.Float64
	ToxinIndexSla37            field.Float64
	ToxinIndexSla38            field.Float64
	ToxinIndexSla39            field.Float64
	ToxinIndexSla40            field.Float64
	ToxinIndexSla41            field.Float64
	ToxinIndexSla42            field.Float64
	ToxinIndexSla43            field.Float64
	ToxinIndexSla44            field.Float64
	ToxinIndexSla45            field.Float64
	ToxinIndexSla46            field.Float64
	ToxinIndexSla47            field.Float64
	SlaughterWaterToxinIndex   slaughterToxinIndexHasOneSlaughterWaterToxinIndex

	fieldMap map[string]field.Expr
}

func (s slaughterToxinIndex) Table(newTableName string) *slaughterToxinIndex {
	s.slaughterToxinIndexDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaughterToxinIndex) As(alias string) *slaughterToxinIndex {
	s.slaughterToxinIndexDo.DO = *(s.slaughterToxinIndexDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaughterToxinIndex) updateTableName(table string) *slaughterToxinIndex {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.SlaughterWaterQualityMonID = field.NewUint(table, "slaughter_water_quality_mon_id")
	s.ToxinIndexSla1 = field.NewFloat64(table, "toxin_index_sla1")
	s.ToxinIndexSla2 = field.NewFloat64(table, "toxin_index_sla2")
	s.ToxinIndexSla3 = field.NewFloat64(table, "toxin_index_sla3")
	s.ToxinIndexSla4 = field.NewFloat64(table, "toxin_index_sla4")
	s.ToxinIndexSla5 = field.NewFloat64(table, "toxin_index_sla5")
	s.ToxinIndexSla6 = field.NewFloat64(table, "toxin_index_sla6")
	s.ToxinIndexSla7 = field.NewFloat64(table, "toxin_index_sla7")
	s.ToxinIndexSla8 = field.NewFloat64(table, "toxin_index_sla8")
	s.ToxinIndexSla9 = field.NewFloat64(table, "toxin_index_sla9")
	s.ToxinIndexSla10 = field.NewFloat64(table, "toxin_index_sla10")
	s.ToxinIndexSla11 = field.NewFloat64(table, "toxin_index_sla11")
	s.ToxinIndexSla12 = field.NewFloat64(table, "toxin_index_sla12")
	s.ToxinIndexSla13 = field.NewFloat64(table, "toxin_index_sla13")
	s.ToxinIndexSla14 = field.NewFloat64(table, "toxin_index_sla14")
	s.ToxinIndexSla15 = field.NewFloat64(table, "toxin_index_sla15")
	s.ToxinIndexSla16 = field.NewFloat64(table, "toxin_index_sla16")
	s.ToxinIndexSla17 = field.NewFloat64(table, "toxin_index_sla17")
	s.ToxinIndexSla18 = field.NewFloat64(table, "toxin_index_sla18")
	s.ToxinIndexSla19 = field.NewFloat64(table, "toxin_index_sla19")
	s.ToxinIndexSla20 = field.NewFloat64(table, "toxin_index_sla20")
	s.ToxinIndexSla21 = field.NewFloat64(table, "toxin_index_sla21")
	s.ToxinIndexSla22 = field.NewFloat64(table, "toxin_index_sla22")
	s.ToxinIndexSla23 = field.NewFloat64(table, "toxin_index_sla23")
	s.ToxinIndexSla24 = field.NewFloat64(table, "toxin_index_sla24")
	s.ToxinIndexSla25 = field.NewFloat64(table, "toxin_index_sla25")
	s.ToxinIndexSla26 = field.NewFloat64(table, "toxin_index_sla26")
	s.ToxinIndexSla27 = field.NewFloat64(table, "toxin_index_sla27")
	s.ToxinIndexSla28 = field.NewFloat64(table, "toxin_index_sla28")
	s.ToxinIndexSla29 = field.NewFloat64(table, "toxin_index_sla29")
	s.ToxinIndexSla30 = field.NewFloat64(table, "toxin_index_sla30")
	s.ToxinIndexSla31 = field.NewFloat64(table, "toxin_index_sla31")
	s.ToxinIndexSla32 = field.NewFloat64(table, "toxin_index_sla32")
	s.ToxinIndexSla33 = field.NewFloat64(table, "toxin_index_sla33")
	s.ToxinIndexSla34 = field.NewFloat64(table, "toxin_index_sla34")
	s.ToxinIndexSla35 = field.NewFloat64(table, "toxin_index_sla35")
	s.ToxinIndexSla36 = field.NewFloat64(table, "toxin_index_sla36")
	s.ToxinIndexSla37 = field.NewFloat64(table, "toxin_index_sla37")
	s.ToxinIndexSla38 = field.NewFloat64(table, "toxin_index_sla38")
	s.ToxinIndexSla39 = field.NewFloat64(table, "toxin_index_sla39")
	s.ToxinIndexSla40 = field.NewFloat64(table, "toxin_index_sla40")
	s.ToxinIndexSla41 = field.NewFloat64(table, "toxin_index_sla41")
	s.ToxinIndexSla42 = field.NewFloat64(table, "toxin_index_sla42")
	s.ToxinIndexSla43 = field.NewFloat64(table, "toxin_index_sla43")
	s.ToxinIndexSla44 = field.NewFloat64(table, "toxin_index_sla44")
	s.ToxinIndexSla45 = field.NewFloat64(table, "toxin_index_sla45")
	s.ToxinIndexSla46 = field.NewFloat64(table, "toxin_index_sla46")
	s.ToxinIndexSla47 = field.NewFloat64(table, "toxin_index_sla47")

	s.fillFieldMap()

	return s
}

func (s *slaughterToxinIndex) WithContext(ctx context.Context) ISlaughterToxinIndexDo {
	return s.slaughterToxinIndexDo.WithContext(ctx)
}

func (s slaughterToxinIndex) TableName() string { return s.slaughterToxinIndexDo.TableName() }

func (s slaughterToxinIndex) Alias() string { return s.slaughterToxinIndexDo.Alias() }

func (s slaughterToxinIndex) Columns(cols ...field.Expr) gen.Columns {
	return s.slaughterToxinIndexDo.Columns(cols...)
}

func (s *slaughterToxinIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaughterToxinIndex) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 53)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["slaughter_water_quality_mon_id"] = s.SlaughterWaterQualityMonID
	s.fieldMap["toxin_index_sla1"] = s.ToxinIndexSla1
	s.fieldMap["toxin_index_sla2"] = s.ToxinIndexSla2
	s.fieldMap["toxin_index_sla3"] = s.ToxinIndexSla3
	s.fieldMap["toxin_index_sla4"] = s.ToxinIndexSla4
	s.fieldMap["toxin_index_sla5"] = s.ToxinIndexSla5
	s.fieldMap["toxin_index_sla6"] = s.ToxinIndexSla6
	s.fieldMap["toxin_index_sla7"] = s.ToxinIndexSla7
	s.fieldMap["toxin_index_sla8"] = s.ToxinIndexSla8
	s.fieldMap["toxin_index_sla9"] = s.ToxinIndexSla9
	s.fieldMap["toxin_index_sla10"] = s.ToxinIndexSla10
	s.fieldMap["toxin_index_sla11"] = s.ToxinIndexSla11
	s.fieldMap["toxin_index_sla12"] = s.ToxinIndexSla12
	s.fieldMap["toxin_index_sla13"] = s.ToxinIndexSla13
	s.fieldMap["toxin_index_sla14"] = s.ToxinIndexSla14
	s.fieldMap["toxin_index_sla15"] = s.ToxinIndexSla15
	s.fieldMap["toxin_index_sla16"] = s.ToxinIndexSla16
	s.fieldMap["toxin_index_sla17"] = s.ToxinIndexSla17
	s.fieldMap["toxin_index_sla18"] = s.ToxinIndexSla18
	s.fieldMap["toxin_index_sla19"] = s.ToxinIndexSla19
	s.fieldMap["toxin_index_sla20"] = s.ToxinIndexSla20
	s.fieldMap["toxin_index_sla21"] = s.ToxinIndexSla21
	s.fieldMap["toxin_index_sla22"] = s.ToxinIndexSla22
	s.fieldMap["toxin_index_sla23"] = s.ToxinIndexSla23
	s.fieldMap["toxin_index_sla24"] = s.ToxinIndexSla24
	s.fieldMap["toxin_index_sla25"] = s.ToxinIndexSla25
	s.fieldMap["toxin_index_sla26"] = s.ToxinIndexSla26
	s.fieldMap["toxin_index_sla27"] = s.ToxinIndexSla27
	s.fieldMap["toxin_index_sla28"] = s.ToxinIndexSla28
	s.fieldMap["toxin_index_sla29"] = s.ToxinIndexSla29
	s.fieldMap["toxin_index_sla30"] = s.ToxinIndexSla30
	s.fieldMap["toxin_index_sla31"] = s.ToxinIndexSla31
	s.fieldMap["toxin_index_sla32"] = s.ToxinIndexSla32
	s.fieldMap["toxin_index_sla33"] = s.ToxinIndexSla33
	s.fieldMap["toxin_index_sla34"] = s.ToxinIndexSla34
	s.fieldMap["toxin_index_sla35"] = s.ToxinIndexSla35
	s.fieldMap["toxin_index_sla36"] = s.ToxinIndexSla36
	s.fieldMap["toxin_index_sla37"] = s.ToxinIndexSla37
	s.fieldMap["toxin_index_sla38"] = s.ToxinIndexSla38
	s.fieldMap["toxin_index_sla39"] = s.ToxinIndexSla39
	s.fieldMap["toxin_index_sla40"] = s.ToxinIndexSla40
	s.fieldMap["toxin_index_sla41"] = s.ToxinIndexSla41
	s.fieldMap["toxin_index_sla42"] = s.ToxinIndexSla42
	s.fieldMap["toxin_index_sla43"] = s.ToxinIndexSla43
	s.fieldMap["toxin_index_sla44"] = s.ToxinIndexSla44
	s.fieldMap["toxin_index_sla45"] = s.ToxinIndexSla45
	s.fieldMap["toxin_index_sla46"] = s.ToxinIndexSla46
	s.fieldMap["toxin_index_sla47"] = s.ToxinIndexSla47

}

func (s slaughterToxinIndex) clone(db *gorm.DB) slaughterToxinIndex {
	s.slaughterToxinIndexDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaughterToxinIndex) replaceDB(db *gorm.DB) slaughterToxinIndex {
	s.slaughterToxinIndexDo.ReplaceDB(db)
	return s
}

type slaughterToxinIndexHasOneSlaughterWaterToxinIndex struct {
	db *gorm.DB

	field.RelationField
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndex) Where(conds ...field.Expr) *slaughterToxinIndexHasOneSlaughterWaterToxinIndex {
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

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndex) WithContext(ctx context.Context) *slaughterToxinIndexHasOneSlaughterWaterToxinIndex {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndex) Session(session *gorm.Session) *slaughterToxinIndexHasOneSlaughterWaterToxinIndex {
	a.db = a.db.Session(session)
	return &a
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndex) Model(m *slaughter.SlaughterToxinIndex) *slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx {
	return &slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx{a.db.Model(m).Association(a.Name())}
}

type slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx struct{ tx *gorm.Association }

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx) Find() (result *slaughter.SlaughterWaterToxinIndex, err error) {
	return result, a.tx.Find(&result)
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx) Append(values ...*slaughter.SlaughterWaterToxinIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx) Replace(values ...*slaughter.SlaughterWaterToxinIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx) Delete(values ...*slaughter.SlaughterWaterToxinIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx) Clear() error {
	return a.tx.Clear()
}

func (a slaughterToxinIndexHasOneSlaughterWaterToxinIndexTx) Count() int64 {
	return a.tx.Count()
}

type slaughterToxinIndexDo struct{ gen.DO }

type ISlaughterToxinIndexDo interface {
	gen.SubQuery
	Debug() ISlaughterToxinIndexDo
	WithContext(ctx context.Context) ISlaughterToxinIndexDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaughterToxinIndexDo
	WriteDB() ISlaughterToxinIndexDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaughterToxinIndexDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaughterToxinIndexDo
	Not(conds ...gen.Condition) ISlaughterToxinIndexDo
	Or(conds ...gen.Condition) ISlaughterToxinIndexDo
	Select(conds ...field.Expr) ISlaughterToxinIndexDo
	Where(conds ...gen.Condition) ISlaughterToxinIndexDo
	Order(conds ...field.Expr) ISlaughterToxinIndexDo
	Distinct(cols ...field.Expr) ISlaughterToxinIndexDo
	Omit(cols ...field.Expr) ISlaughterToxinIndexDo
	Join(table schema.Tabler, on ...field.Expr) ISlaughterToxinIndexDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterToxinIndexDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterToxinIndexDo
	Group(cols ...field.Expr) ISlaughterToxinIndexDo
	Having(conds ...gen.Condition) ISlaughterToxinIndexDo
	Limit(limit int) ISlaughterToxinIndexDo
	Offset(offset int) ISlaughterToxinIndexDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterToxinIndexDo
	Unscoped() ISlaughterToxinIndexDo
	Create(values ...*slaughter.SlaughterToxinIndex) error
	CreateInBatches(values []*slaughter.SlaughterToxinIndex, batchSize int) error
	Save(values ...*slaughter.SlaughterToxinIndex) error
	First() (*slaughter.SlaughterToxinIndex, error)
	Take() (*slaughter.SlaughterToxinIndex, error)
	Last() (*slaughter.SlaughterToxinIndex, error)
	Find() ([]*slaughter.SlaughterToxinIndex, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterToxinIndex, err error)
	FindInBatches(result *[]*slaughter.SlaughterToxinIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaughterToxinIndex) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaughterToxinIndexDo
	Assign(attrs ...field.AssignExpr) ISlaughterToxinIndexDo
	Joins(fields ...field.RelationField) ISlaughterToxinIndexDo
	Preload(fields ...field.RelationField) ISlaughterToxinIndexDo
	FirstOrInit() (*slaughter.SlaughterToxinIndex, error)
	FirstOrCreate() (*slaughter.SlaughterToxinIndex, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaughterToxinIndex, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaughterToxinIndexDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaughterToxinIndexDo) Debug() ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Debug())
}

func (s slaughterToxinIndexDo) WithContext(ctx context.Context) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaughterToxinIndexDo) ReadDB() ISlaughterToxinIndexDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaughterToxinIndexDo) WriteDB() ISlaughterToxinIndexDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaughterToxinIndexDo) Session(config *gorm.Session) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaughterToxinIndexDo) Clauses(conds ...clause.Expression) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaughterToxinIndexDo) Returning(value interface{}, columns ...string) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaughterToxinIndexDo) Not(conds ...gen.Condition) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaughterToxinIndexDo) Or(conds ...gen.Condition) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaughterToxinIndexDo) Select(conds ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaughterToxinIndexDo) Where(conds ...gen.Condition) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaughterToxinIndexDo) Order(conds ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaughterToxinIndexDo) Distinct(cols ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaughterToxinIndexDo) Omit(cols ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaughterToxinIndexDo) Join(table schema.Tabler, on ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaughterToxinIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaughterToxinIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaughterToxinIndexDo) Group(cols ...field.Expr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaughterToxinIndexDo) Having(conds ...gen.Condition) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaughterToxinIndexDo) Limit(limit int) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaughterToxinIndexDo) Offset(offset int) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaughterToxinIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaughterToxinIndexDo) Unscoped() ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaughterToxinIndexDo) Create(values ...*slaughter.SlaughterToxinIndex) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaughterToxinIndexDo) CreateInBatches(values []*slaughter.SlaughterToxinIndex, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaughterToxinIndexDo) Save(values ...*slaughter.SlaughterToxinIndex) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaughterToxinIndexDo) First() (*slaughter.SlaughterToxinIndex, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterToxinIndex), nil
	}
}

func (s slaughterToxinIndexDo) Take() (*slaughter.SlaughterToxinIndex, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterToxinIndex), nil
	}
}

func (s slaughterToxinIndexDo) Last() (*slaughter.SlaughterToxinIndex, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterToxinIndex), nil
	}
}

func (s slaughterToxinIndexDo) Find() ([]*slaughter.SlaughterToxinIndex, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaughterToxinIndex), err
}

func (s slaughterToxinIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterToxinIndex, err error) {
	buf := make([]*slaughter.SlaughterToxinIndex, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaughterToxinIndexDo) FindInBatches(result *[]*slaughter.SlaughterToxinIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaughterToxinIndexDo) Attrs(attrs ...field.AssignExpr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaughterToxinIndexDo) Assign(attrs ...field.AssignExpr) ISlaughterToxinIndexDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaughterToxinIndexDo) Joins(fields ...field.RelationField) ISlaughterToxinIndexDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaughterToxinIndexDo) Preload(fields ...field.RelationField) ISlaughterToxinIndexDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaughterToxinIndexDo) FirstOrInit() (*slaughter.SlaughterToxinIndex, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterToxinIndex), nil
	}
}

func (s slaughterToxinIndexDo) FirstOrCreate() (*slaughter.SlaughterToxinIndex, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterToxinIndex), nil
	}
}

func (s slaughterToxinIndexDo) FindByPage(offset int, limit int) (result []*slaughter.SlaughterToxinIndex, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s slaughterToxinIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaughterToxinIndexDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaughterToxinIndexDo) Delete(models ...*slaughter.SlaughterToxinIndex) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaughterToxinIndexDo) withDO(do gen.Dao) *slaughterToxinIndexDo {
	s.DO = *do.(*gen.DO)
	return s
}
