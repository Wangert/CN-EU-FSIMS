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

func newSlaughterWaterToxinIndex(db *gorm.DB, opts ...gen.DOOption) slaughterWaterToxinIndex {
	_slaughterWaterToxinIndex := slaughterWaterToxinIndex{}

	_slaughterWaterToxinIndex.slaughterWaterToxinIndexDo.UseDB(db, opts...)
	_slaughterWaterToxinIndex.slaughterWaterToxinIndexDo.UseModel(&slaughter.SlaughterWaterToxinIndex{})

	tableName := _slaughterWaterToxinIndex.slaughterWaterToxinIndexDo.TableName()
	_slaughterWaterToxinIndex.ALL = field.NewAsterisk(tableName)
	_slaughterWaterToxinIndex.ID = field.NewUint(tableName, "id")
	_slaughterWaterToxinIndex.CreatedAt = field.NewTime(tableName, "created_at")
	_slaughterWaterToxinIndex.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaughterWaterToxinIndex.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaughterWaterToxinIndex.SlaughterToxinIndexID = field.NewUint(tableName, "slaughter_toxin_index_id")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex1 = field.NewFloat32(tableName, "slaughter_water_toxin_index1")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex2 = field.NewFloat32(tableName, "slaughter_water_toxin_index2")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex3 = field.NewFloat32(tableName, "slaughter_water_toxin_index3")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex4 = field.NewFloat32(tableName, "slaughter_water_toxin_index4")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex5 = field.NewFloat32(tableName, "slaughter_water_toxin_index5")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex6 = field.NewFloat32(tableName, "slaughter_water_toxin_index6")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex7 = field.NewFloat32(tableName, "slaughter_water_toxin_index7")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex8 = field.NewFloat32(tableName, "slaughter_water_toxin_index8")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex9 = field.NewFloat32(tableName, "slaughter_water_toxin_index9")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex10 = field.NewFloat32(tableName, "slaughter_water_toxin_index10")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex11 = field.NewFloat32(tableName, "slaughter_water_toxin_index11")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex12 = field.NewFloat32(tableName, "slaughter_water_toxin_index12")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex13 = field.NewFloat32(tableName, "slaughter_water_toxin_index13")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex14 = field.NewFloat32(tableName, "slaughter_water_toxin_index14")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex15 = field.NewFloat32(tableName, "slaughter_water_toxin_index15")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex16 = field.NewFloat32(tableName, "slaughter_water_toxin_index16")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex17 = field.NewFloat32(tableName, "slaughter_water_toxin_index17")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex18 = field.NewFloat32(tableName, "slaughter_water_toxin_index18")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex19 = field.NewFloat32(tableName, "slaughter_water_toxin_index19")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex20 = field.NewFloat32(tableName, "slaughter_water_toxin_index20")
	_slaughterWaterToxinIndex.SlaughterWaterToxinIndex21 = field.NewFloat32(tableName, "slaughter_water_toxin_index21")

	_slaughterWaterToxinIndex.fillFieldMap()

	return _slaughterWaterToxinIndex
}

type slaughterWaterToxinIndex struct {
	slaughterWaterToxinIndexDo slaughterWaterToxinIndexDo

	ALL                        field.Asterisk
	ID                         field.Uint
	CreatedAt                  field.Time
	UpdatedAt                  field.Time
	DeletedAt                  field.Field
	SlaughterToxinIndexID      field.Uint
	SlaughterWaterToxinIndex1  field.Float32
	SlaughterWaterToxinIndex2  field.Float32
	SlaughterWaterToxinIndex3  field.Float32
	SlaughterWaterToxinIndex4  field.Float32
	SlaughterWaterToxinIndex5  field.Float32
	SlaughterWaterToxinIndex6  field.Float32
	SlaughterWaterToxinIndex7  field.Float32
	SlaughterWaterToxinIndex8  field.Float32
	SlaughterWaterToxinIndex9  field.Float32
	SlaughterWaterToxinIndex10 field.Float32
	SlaughterWaterToxinIndex11 field.Float32
	SlaughterWaterToxinIndex12 field.Float32
	SlaughterWaterToxinIndex13 field.Float32
	SlaughterWaterToxinIndex14 field.Float32
	SlaughterWaterToxinIndex15 field.Float32
	SlaughterWaterToxinIndex16 field.Float32
	SlaughterWaterToxinIndex17 field.Float32
	SlaughterWaterToxinIndex18 field.Float32
	SlaughterWaterToxinIndex19 field.Float32
	SlaughterWaterToxinIndex20 field.Float32
	SlaughterWaterToxinIndex21 field.Float32

	fieldMap map[string]field.Expr
}

func (s slaughterWaterToxinIndex) Table(newTableName string) *slaughterWaterToxinIndex {
	s.slaughterWaterToxinIndexDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaughterWaterToxinIndex) As(alias string) *slaughterWaterToxinIndex {
	s.slaughterWaterToxinIndexDo.DO = *(s.slaughterWaterToxinIndexDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaughterWaterToxinIndex) updateTableName(table string) *slaughterWaterToxinIndex {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.SlaughterToxinIndexID = field.NewUint(table, "slaughter_toxin_index_id")
	s.SlaughterWaterToxinIndex1 = field.NewFloat32(table, "slaughter_water_toxin_index1")
	s.SlaughterWaterToxinIndex2 = field.NewFloat32(table, "slaughter_water_toxin_index2")
	s.SlaughterWaterToxinIndex3 = field.NewFloat32(table, "slaughter_water_toxin_index3")
	s.SlaughterWaterToxinIndex4 = field.NewFloat32(table, "slaughter_water_toxin_index4")
	s.SlaughterWaterToxinIndex5 = field.NewFloat32(table, "slaughter_water_toxin_index5")
	s.SlaughterWaterToxinIndex6 = field.NewFloat32(table, "slaughter_water_toxin_index6")
	s.SlaughterWaterToxinIndex7 = field.NewFloat32(table, "slaughter_water_toxin_index7")
	s.SlaughterWaterToxinIndex8 = field.NewFloat32(table, "slaughter_water_toxin_index8")
	s.SlaughterWaterToxinIndex9 = field.NewFloat32(table, "slaughter_water_toxin_index9")
	s.SlaughterWaterToxinIndex10 = field.NewFloat32(table, "slaughter_water_toxin_index10")
	s.SlaughterWaterToxinIndex11 = field.NewFloat32(table, "slaughter_water_toxin_index11")
	s.SlaughterWaterToxinIndex12 = field.NewFloat32(table, "slaughter_water_toxin_index12")
	s.SlaughterWaterToxinIndex13 = field.NewFloat32(table, "slaughter_water_toxin_index13")
	s.SlaughterWaterToxinIndex14 = field.NewFloat32(table, "slaughter_water_toxin_index14")
	s.SlaughterWaterToxinIndex15 = field.NewFloat32(table, "slaughter_water_toxin_index15")
	s.SlaughterWaterToxinIndex16 = field.NewFloat32(table, "slaughter_water_toxin_index16")
	s.SlaughterWaterToxinIndex17 = field.NewFloat32(table, "slaughter_water_toxin_index17")
	s.SlaughterWaterToxinIndex18 = field.NewFloat32(table, "slaughter_water_toxin_index18")
	s.SlaughterWaterToxinIndex19 = field.NewFloat32(table, "slaughter_water_toxin_index19")
	s.SlaughterWaterToxinIndex20 = field.NewFloat32(table, "slaughter_water_toxin_index20")
	s.SlaughterWaterToxinIndex21 = field.NewFloat32(table, "slaughter_water_toxin_index21")

	s.fillFieldMap()

	return s
}

func (s *slaughterWaterToxinIndex) WithContext(ctx context.Context) ISlaughterWaterToxinIndexDo {
	return s.slaughterWaterToxinIndexDo.WithContext(ctx)
}

func (s slaughterWaterToxinIndex) TableName() string { return s.slaughterWaterToxinIndexDo.TableName() }

func (s slaughterWaterToxinIndex) Alias() string { return s.slaughterWaterToxinIndexDo.Alias() }

func (s slaughterWaterToxinIndex) Columns(cols ...field.Expr) gen.Columns {
	return s.slaughterWaterToxinIndexDo.Columns(cols...)
}

func (s *slaughterWaterToxinIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaughterWaterToxinIndex) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 26)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["slaughter_toxin_index_id"] = s.SlaughterToxinIndexID
	s.fieldMap["slaughter_water_toxin_index1"] = s.SlaughterWaterToxinIndex1
	s.fieldMap["slaughter_water_toxin_index2"] = s.SlaughterWaterToxinIndex2
	s.fieldMap["slaughter_water_toxin_index3"] = s.SlaughterWaterToxinIndex3
	s.fieldMap["slaughter_water_toxin_index4"] = s.SlaughterWaterToxinIndex4
	s.fieldMap["slaughter_water_toxin_index5"] = s.SlaughterWaterToxinIndex5
	s.fieldMap["slaughter_water_toxin_index6"] = s.SlaughterWaterToxinIndex6
	s.fieldMap["slaughter_water_toxin_index7"] = s.SlaughterWaterToxinIndex7
	s.fieldMap["slaughter_water_toxin_index8"] = s.SlaughterWaterToxinIndex8
	s.fieldMap["slaughter_water_toxin_index9"] = s.SlaughterWaterToxinIndex9
	s.fieldMap["slaughter_water_toxin_index10"] = s.SlaughterWaterToxinIndex10
	s.fieldMap["slaughter_water_toxin_index11"] = s.SlaughterWaterToxinIndex11
	s.fieldMap["slaughter_water_toxin_index12"] = s.SlaughterWaterToxinIndex12
	s.fieldMap["slaughter_water_toxin_index13"] = s.SlaughterWaterToxinIndex13
	s.fieldMap["slaughter_water_toxin_index14"] = s.SlaughterWaterToxinIndex14
	s.fieldMap["slaughter_water_toxin_index15"] = s.SlaughterWaterToxinIndex15
	s.fieldMap["slaughter_water_toxin_index16"] = s.SlaughterWaterToxinIndex16
	s.fieldMap["slaughter_water_toxin_index17"] = s.SlaughterWaterToxinIndex17
	s.fieldMap["slaughter_water_toxin_index18"] = s.SlaughterWaterToxinIndex18
	s.fieldMap["slaughter_water_toxin_index19"] = s.SlaughterWaterToxinIndex19
	s.fieldMap["slaughter_water_toxin_index20"] = s.SlaughterWaterToxinIndex20
	s.fieldMap["slaughter_water_toxin_index21"] = s.SlaughterWaterToxinIndex21
}

func (s slaughterWaterToxinIndex) clone(db *gorm.DB) slaughterWaterToxinIndex {
	s.slaughterWaterToxinIndexDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaughterWaterToxinIndex) replaceDB(db *gorm.DB) slaughterWaterToxinIndex {
	s.slaughterWaterToxinIndexDo.ReplaceDB(db)
	return s
}

type slaughterWaterToxinIndexDo struct{ gen.DO }

type ISlaughterWaterToxinIndexDo interface {
	gen.SubQuery
	Debug() ISlaughterWaterToxinIndexDo
	WithContext(ctx context.Context) ISlaughterWaterToxinIndexDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaughterWaterToxinIndexDo
	WriteDB() ISlaughterWaterToxinIndexDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaughterWaterToxinIndexDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaughterWaterToxinIndexDo
	Not(conds ...gen.Condition) ISlaughterWaterToxinIndexDo
	Or(conds ...gen.Condition) ISlaughterWaterToxinIndexDo
	Select(conds ...field.Expr) ISlaughterWaterToxinIndexDo
	Where(conds ...gen.Condition) ISlaughterWaterToxinIndexDo
	Order(conds ...field.Expr) ISlaughterWaterToxinIndexDo
	Distinct(cols ...field.Expr) ISlaughterWaterToxinIndexDo
	Omit(cols ...field.Expr) ISlaughterWaterToxinIndexDo
	Join(table schema.Tabler, on ...field.Expr) ISlaughterWaterToxinIndexDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterToxinIndexDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterToxinIndexDo
	Group(cols ...field.Expr) ISlaughterWaterToxinIndexDo
	Having(conds ...gen.Condition) ISlaughterWaterToxinIndexDo
	Limit(limit int) ISlaughterWaterToxinIndexDo
	Offset(offset int) ISlaughterWaterToxinIndexDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterWaterToxinIndexDo
	Unscoped() ISlaughterWaterToxinIndexDo
	Create(values ...*slaughter.SlaughterWaterToxinIndex) error
	CreateInBatches(values []*slaughter.SlaughterWaterToxinIndex, batchSize int) error
	Save(values ...*slaughter.SlaughterWaterToxinIndex) error
	First() (*slaughter.SlaughterWaterToxinIndex, error)
	Take() (*slaughter.SlaughterWaterToxinIndex, error)
	Last() (*slaughter.SlaughterWaterToxinIndex, error)
	Find() ([]*slaughter.SlaughterWaterToxinIndex, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterWaterToxinIndex, err error)
	FindInBatches(result *[]*slaughter.SlaughterWaterToxinIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaughterWaterToxinIndex) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaughterWaterToxinIndexDo
	Assign(attrs ...field.AssignExpr) ISlaughterWaterToxinIndexDo
	Joins(fields ...field.RelationField) ISlaughterWaterToxinIndexDo
	Preload(fields ...field.RelationField) ISlaughterWaterToxinIndexDo
	FirstOrInit() (*slaughter.SlaughterWaterToxinIndex, error)
	FirstOrCreate() (*slaughter.SlaughterWaterToxinIndex, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaughterWaterToxinIndex, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaughterWaterToxinIndexDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaughterWaterToxinIndexDo) Debug() ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Debug())
}

func (s slaughterWaterToxinIndexDo) WithContext(ctx context.Context) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaughterWaterToxinIndexDo) ReadDB() ISlaughterWaterToxinIndexDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaughterWaterToxinIndexDo) WriteDB() ISlaughterWaterToxinIndexDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaughterWaterToxinIndexDo) Session(config *gorm.Session) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaughterWaterToxinIndexDo) Clauses(conds ...clause.Expression) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaughterWaterToxinIndexDo) Returning(value interface{}, columns ...string) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaughterWaterToxinIndexDo) Not(conds ...gen.Condition) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaughterWaterToxinIndexDo) Or(conds ...gen.Condition) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaughterWaterToxinIndexDo) Select(conds ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaughterWaterToxinIndexDo) Where(conds ...gen.Condition) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaughterWaterToxinIndexDo) Order(conds ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaughterWaterToxinIndexDo) Distinct(cols ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaughterWaterToxinIndexDo) Omit(cols ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaughterWaterToxinIndexDo) Join(table schema.Tabler, on ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaughterWaterToxinIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaughterWaterToxinIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaughterWaterToxinIndexDo) Group(cols ...field.Expr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaughterWaterToxinIndexDo) Having(conds ...gen.Condition) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaughterWaterToxinIndexDo) Limit(limit int) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaughterWaterToxinIndexDo) Offset(offset int) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaughterWaterToxinIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaughterWaterToxinIndexDo) Unscoped() ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaughterWaterToxinIndexDo) Create(values ...*slaughter.SlaughterWaterToxinIndex) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaughterWaterToxinIndexDo) CreateInBatches(values []*slaughter.SlaughterWaterToxinIndex, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaughterWaterToxinIndexDo) Save(values ...*slaughter.SlaughterWaterToxinIndex) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaughterWaterToxinIndexDo) First() (*slaughter.SlaughterWaterToxinIndex, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterToxinIndex), nil
	}
}

func (s slaughterWaterToxinIndexDo) Take() (*slaughter.SlaughterWaterToxinIndex, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterToxinIndex), nil
	}
}

func (s slaughterWaterToxinIndexDo) Last() (*slaughter.SlaughterWaterToxinIndex, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterToxinIndex), nil
	}
}

func (s slaughterWaterToxinIndexDo) Find() ([]*slaughter.SlaughterWaterToxinIndex, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaughterWaterToxinIndex), err
}

func (s slaughterWaterToxinIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterWaterToxinIndex, err error) {
	buf := make([]*slaughter.SlaughterWaterToxinIndex, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaughterWaterToxinIndexDo) FindInBatches(result *[]*slaughter.SlaughterWaterToxinIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaughterWaterToxinIndexDo) Attrs(attrs ...field.AssignExpr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaughterWaterToxinIndexDo) Assign(attrs ...field.AssignExpr) ISlaughterWaterToxinIndexDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaughterWaterToxinIndexDo) Joins(fields ...field.RelationField) ISlaughterWaterToxinIndexDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaughterWaterToxinIndexDo) Preload(fields ...field.RelationField) ISlaughterWaterToxinIndexDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaughterWaterToxinIndexDo) FirstOrInit() (*slaughter.SlaughterWaterToxinIndex, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterToxinIndex), nil
	}
}

func (s slaughterWaterToxinIndexDo) FirstOrCreate() (*slaughter.SlaughterWaterToxinIndex, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterToxinIndex), nil
	}
}

func (s slaughterWaterToxinIndexDo) FindByPage(offset int, limit int) (result []*slaughter.SlaughterWaterToxinIndex, count int64, err error) {
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

func (s slaughterWaterToxinIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaughterWaterToxinIndexDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaughterWaterToxinIndexDo) Delete(models ...*slaughter.SlaughterWaterToxinIndex) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaughterWaterToxinIndexDo) withDO(do gen.Dao) *slaughterWaterToxinIndexDo {
	s.DO = *do.(*gen.DO)
	return s
}
