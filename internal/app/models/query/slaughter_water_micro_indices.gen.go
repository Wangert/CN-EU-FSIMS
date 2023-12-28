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

func newSlaughterWaterMicroIndex(db *gorm.DB, opts ...gen.DOOption) slaughterWaterMicroIndex {
	_slaughterWaterMicroIndex := slaughterWaterMicroIndex{}

	_slaughterWaterMicroIndex.slaughterWaterMicroIndexDo.UseDB(db, opts...)
	_slaughterWaterMicroIndex.slaughterWaterMicroIndexDo.UseModel(&slaughter.SlaughterWaterMicroIndex{})

	tableName := _slaughterWaterMicroIndex.slaughterWaterMicroIndexDo.TableName()
	_slaughterWaterMicroIndex.ALL = field.NewAsterisk(tableName)
	_slaughterWaterMicroIndex.ID = field.NewUint(tableName, "id")
	_slaughterWaterMicroIndex.CreatedAt = field.NewTime(tableName, "created_at")
	_slaughterWaterMicroIndex.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaughterWaterMicroIndex.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaughterWaterMicroIndex.WaterQualityMonID = field.NewUint(tableName, "water_quality_mon_id")
	_slaughterWaterMicroIndex.MicroIndex1 = field.NewFloat64(tableName, "micro_index1")
	_slaughterWaterMicroIndex.MicroIndex2 = field.NewFloat64(tableName, "micro_index2")
	_slaughterWaterMicroIndex.MicroIndex3 = field.NewFloat64(tableName, "micro_index3")

	_slaughterWaterMicroIndex.fillFieldMap()

	return _slaughterWaterMicroIndex
}

type slaughterWaterMicroIndex struct {
	slaughterWaterMicroIndexDo slaughterWaterMicroIndexDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	WaterQualityMonID field.Uint
	MicroIndex1       field.Float64
	MicroIndex2       field.Float64
	MicroIndex3       field.Float64

	fieldMap map[string]field.Expr
}

func (s slaughterWaterMicroIndex) Table(newTableName string) *slaughterWaterMicroIndex {
	s.slaughterWaterMicroIndexDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaughterWaterMicroIndex) As(alias string) *slaughterWaterMicroIndex {
	s.slaughterWaterMicroIndexDo.DO = *(s.slaughterWaterMicroIndexDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaughterWaterMicroIndex) updateTableName(table string) *slaughterWaterMicroIndex {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.WaterQualityMonID = field.NewUint(table, "water_quality_mon_id")
	s.MicroIndex1 = field.NewFloat64(table, "micro_index1")
	s.MicroIndex2 = field.NewFloat64(table, "micro_index2")
	s.MicroIndex3 = field.NewFloat64(table, "micro_index3")

	s.fillFieldMap()

	return s
}

func (s *slaughterWaterMicroIndex) WithContext(ctx context.Context) ISlaughterWaterMicroIndexDo {
	return s.slaughterWaterMicroIndexDo.WithContext(ctx)
}

func (s slaughterWaterMicroIndex) TableName() string { return s.slaughterWaterMicroIndexDo.TableName() }

func (s slaughterWaterMicroIndex) Alias() string { return s.slaughterWaterMicroIndexDo.Alias() }

func (s slaughterWaterMicroIndex) Columns(cols ...field.Expr) gen.Columns {
	return s.slaughterWaterMicroIndexDo.Columns(cols...)
}

func (s *slaughterWaterMicroIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaughterWaterMicroIndex) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["water_quality_mon_id"] = s.WaterQualityMonID
	s.fieldMap["micro_index1"] = s.MicroIndex1
	s.fieldMap["micro_index2"] = s.MicroIndex2
	s.fieldMap["micro_index3"] = s.MicroIndex3
}

func (s slaughterWaterMicroIndex) clone(db *gorm.DB) slaughterWaterMicroIndex {
	s.slaughterWaterMicroIndexDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaughterWaterMicroIndex) replaceDB(db *gorm.DB) slaughterWaterMicroIndex {
	s.slaughterWaterMicroIndexDo.ReplaceDB(db)
	return s
}

type slaughterWaterMicroIndexDo struct{ gen.DO }

type ISlaughterWaterMicroIndexDo interface {
	gen.SubQuery
	Debug() ISlaughterWaterMicroIndexDo
	WithContext(ctx context.Context) ISlaughterWaterMicroIndexDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaughterWaterMicroIndexDo
	WriteDB() ISlaughterWaterMicroIndexDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaughterWaterMicroIndexDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaughterWaterMicroIndexDo
	Not(conds ...gen.Condition) ISlaughterWaterMicroIndexDo
	Or(conds ...gen.Condition) ISlaughterWaterMicroIndexDo
	Select(conds ...field.Expr) ISlaughterWaterMicroIndexDo
	Where(conds ...gen.Condition) ISlaughterWaterMicroIndexDo
	Order(conds ...field.Expr) ISlaughterWaterMicroIndexDo
	Distinct(cols ...field.Expr) ISlaughterWaterMicroIndexDo
	Omit(cols ...field.Expr) ISlaughterWaterMicroIndexDo
	Join(table schema.Tabler, on ...field.Expr) ISlaughterWaterMicroIndexDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterMicroIndexDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterMicroIndexDo
	Group(cols ...field.Expr) ISlaughterWaterMicroIndexDo
	Having(conds ...gen.Condition) ISlaughterWaterMicroIndexDo
	Limit(limit int) ISlaughterWaterMicroIndexDo
	Offset(offset int) ISlaughterWaterMicroIndexDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterWaterMicroIndexDo
	Unscoped() ISlaughterWaterMicroIndexDo
	Create(values ...*slaughter.SlaughterWaterMicroIndex) error
	CreateInBatches(values []*slaughter.SlaughterWaterMicroIndex, batchSize int) error
	Save(values ...*slaughter.SlaughterWaterMicroIndex) error
	First() (*slaughter.SlaughterWaterMicroIndex, error)
	Take() (*slaughter.SlaughterWaterMicroIndex, error)
	Last() (*slaughter.SlaughterWaterMicroIndex, error)
	Find() ([]*slaughter.SlaughterWaterMicroIndex, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterWaterMicroIndex, err error)
	FindInBatches(result *[]*slaughter.SlaughterWaterMicroIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaughterWaterMicroIndex) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaughterWaterMicroIndexDo
	Assign(attrs ...field.AssignExpr) ISlaughterWaterMicroIndexDo
	Joins(fields ...field.RelationField) ISlaughterWaterMicroIndexDo
	Preload(fields ...field.RelationField) ISlaughterWaterMicroIndexDo
	FirstOrInit() (*slaughter.SlaughterWaterMicroIndex, error)
	FirstOrCreate() (*slaughter.SlaughterWaterMicroIndex, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaughterWaterMicroIndex, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaughterWaterMicroIndexDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaughterWaterMicroIndexDo) Debug() ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Debug())
}

func (s slaughterWaterMicroIndexDo) WithContext(ctx context.Context) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaughterWaterMicroIndexDo) ReadDB() ISlaughterWaterMicroIndexDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaughterWaterMicroIndexDo) WriteDB() ISlaughterWaterMicroIndexDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaughterWaterMicroIndexDo) Session(config *gorm.Session) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaughterWaterMicroIndexDo) Clauses(conds ...clause.Expression) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaughterWaterMicroIndexDo) Returning(value interface{}, columns ...string) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaughterWaterMicroIndexDo) Not(conds ...gen.Condition) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaughterWaterMicroIndexDo) Or(conds ...gen.Condition) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaughterWaterMicroIndexDo) Select(conds ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaughterWaterMicroIndexDo) Where(conds ...gen.Condition) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaughterWaterMicroIndexDo) Order(conds ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaughterWaterMicroIndexDo) Distinct(cols ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaughterWaterMicroIndexDo) Omit(cols ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaughterWaterMicroIndexDo) Join(table schema.Tabler, on ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaughterWaterMicroIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaughterWaterMicroIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaughterWaterMicroIndexDo) Group(cols ...field.Expr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaughterWaterMicroIndexDo) Having(conds ...gen.Condition) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaughterWaterMicroIndexDo) Limit(limit int) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaughterWaterMicroIndexDo) Offset(offset int) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaughterWaterMicroIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaughterWaterMicroIndexDo) Unscoped() ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaughterWaterMicroIndexDo) Create(values ...*slaughter.SlaughterWaterMicroIndex) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaughterWaterMicroIndexDo) CreateInBatches(values []*slaughter.SlaughterWaterMicroIndex, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaughterWaterMicroIndexDo) Save(values ...*slaughter.SlaughterWaterMicroIndex) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaughterWaterMicroIndexDo) First() (*slaughter.SlaughterWaterMicroIndex, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterMicroIndex), nil
	}
}

func (s slaughterWaterMicroIndexDo) Take() (*slaughter.SlaughterWaterMicroIndex, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterMicroIndex), nil
	}
}

func (s slaughterWaterMicroIndexDo) Last() (*slaughter.SlaughterWaterMicroIndex, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterMicroIndex), nil
	}
}

func (s slaughterWaterMicroIndexDo) Find() ([]*slaughter.SlaughterWaterMicroIndex, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaughterWaterMicroIndex), err
}

func (s slaughterWaterMicroIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterWaterMicroIndex, err error) {
	buf := make([]*slaughter.SlaughterWaterMicroIndex, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaughterWaterMicroIndexDo) FindInBatches(result *[]*slaughter.SlaughterWaterMicroIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaughterWaterMicroIndexDo) Attrs(attrs ...field.AssignExpr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaughterWaterMicroIndexDo) Assign(attrs ...field.AssignExpr) ISlaughterWaterMicroIndexDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaughterWaterMicroIndexDo) Joins(fields ...field.RelationField) ISlaughterWaterMicroIndexDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaughterWaterMicroIndexDo) Preload(fields ...field.RelationField) ISlaughterWaterMicroIndexDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaughterWaterMicroIndexDo) FirstOrInit() (*slaughter.SlaughterWaterMicroIndex, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterMicroIndex), nil
	}
}

func (s slaughterWaterMicroIndexDo) FirstOrCreate() (*slaughter.SlaughterWaterMicroIndex, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterWaterMicroIndex), nil
	}
}

func (s slaughterWaterMicroIndexDo) FindByPage(offset int, limit int) (result []*slaughter.SlaughterWaterMicroIndex, count int64, err error) {
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

func (s slaughterWaterMicroIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaughterWaterMicroIndexDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaughterWaterMicroIndexDo) Delete(models ...*slaughter.SlaughterWaterMicroIndex) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaughterWaterMicroIndexDo) withDO(do gen.Dao) *slaughterWaterMicroIndexDo {
	s.DO = *do.(*gen.DO)
	return s
}