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

func newStaUni(db *gorm.DB, opts ...gen.DOOption) staUni {
	_staUni := staUni{}

	_staUni.staUniDo.UseDB(db, opts...)
	_staUni.staUniDo.UseModel(&slaughter.StaUni{})

	tableName := _staUni.staUniDo.TableName()
	_staUni.ALL = field.NewAsterisk(tableName)
	_staUni.ID = field.NewUint(tableName, "id")
	_staUni.CreatedAt = field.NewTime(tableName, "created_at")
	_staUni.UpdatedAt = field.NewTime(tableName, "updated_at")
	_staUni.DeletedAt = field.NewField(tableName, "deleted_at")
	_staUni.HouseNumber = field.NewString(tableName, "house_number")
	_staUni.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_staUni.StaUni1 = field.NewFloat64(tableName, "sta_uni1")
	_staUni.StaUni2 = field.NewFloat64(tableName, "sta_uni2")
	_staUni.StaUni3 = field.NewFloat64(tableName, "sta_uni3")
	_staUni.StaUni4 = field.NewFloat64(tableName, "sta_uni4")
	_staUni.StaUni5 = field.NewFloat64(tableName, "sta_uni5")
	_staUni.StaUni6 = field.NewFloat64(tableName, "sta_uni6")
	_staUni.StaUni7 = field.NewFloat64(tableName, "sta_uni7")
	_staUni.StaUni8 = field.NewFloat64(tableName, "sta_uni8")
	_staUni.StaUni9 = field.NewFloat64(tableName, "sta_uni9")
	_staUni.StaUni10 = field.NewFloat64(tableName, "sta_uni10")
	_staUni.StaUni11 = field.NewFloat64(tableName, "sta_uni11")
	_staUni.StaUni12 = field.NewFloat64(tableName, "sta_uni12")
	_staUni.StaUni13 = field.NewFloat64(tableName, "sta_uni13")
	_staUni.StaUni14 = field.NewFloat64(tableName, "sta_uni14")

	_staUni.fillFieldMap()

	return _staUni
}

type staUni struct {
	staUniDo staUniDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HouseNumber  field.String
	TimeRecordAt field.Time
	StaUni1      field.Float64
	StaUni2      field.Float64
	StaUni3      field.Float64
	StaUni4      field.Float64
	StaUni5      field.Float64
	StaUni6      field.Float64
	StaUni7      field.Float64
	StaUni8      field.Float64
	StaUni9      field.Float64
	StaUni10     field.Float64
	StaUni11     field.Float64
	StaUni12     field.Float64
	StaUni13     field.Float64
	StaUni14     field.Float64

	fieldMap map[string]field.Expr
}

func (s staUni) Table(newTableName string) *staUni {
	s.staUniDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s staUni) As(alias string) *staUni {
	s.staUniDo.DO = *(s.staUniDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *staUni) updateTableName(table string) *staUni {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.HouseNumber = field.NewString(table, "house_number")
	s.TimeRecordAt = field.NewTime(table, "time_record_at")
	s.StaUni1 = field.NewFloat64(table, "sta_uni1")
	s.StaUni2 = field.NewFloat64(table, "sta_uni2")
	s.StaUni3 = field.NewFloat64(table, "sta_uni3")
	s.StaUni4 = field.NewFloat64(table, "sta_uni4")
	s.StaUni5 = field.NewFloat64(table, "sta_uni5")
	s.StaUni6 = field.NewFloat64(table, "sta_uni6")
	s.StaUni7 = field.NewFloat64(table, "sta_uni7")
	s.StaUni8 = field.NewFloat64(table, "sta_uni8")
	s.StaUni9 = field.NewFloat64(table, "sta_uni9")
	s.StaUni10 = field.NewFloat64(table, "sta_uni10")
	s.StaUni11 = field.NewFloat64(table, "sta_uni11")
	s.StaUni12 = field.NewFloat64(table, "sta_uni12")
	s.StaUni13 = field.NewFloat64(table, "sta_uni13")
	s.StaUni14 = field.NewFloat64(table, "sta_uni14")

	s.fillFieldMap()

	return s
}

func (s *staUni) WithContext(ctx context.Context) IStaUniDo { return s.staUniDo.WithContext(ctx) }

func (s staUni) TableName() string { return s.staUniDo.TableName() }

func (s staUni) Alias() string { return s.staUniDo.Alias() }

func (s staUni) Columns(cols ...field.Expr) gen.Columns { return s.staUniDo.Columns(cols...) }

func (s *staUni) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *staUni) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["house_number"] = s.HouseNumber
	s.fieldMap["time_record_at"] = s.TimeRecordAt
	s.fieldMap["sta_uni1"] = s.StaUni1
	s.fieldMap["sta_uni2"] = s.StaUni2
	s.fieldMap["sta_uni3"] = s.StaUni3
	s.fieldMap["sta_uni4"] = s.StaUni4
	s.fieldMap["sta_uni5"] = s.StaUni5
	s.fieldMap["sta_uni6"] = s.StaUni6
	s.fieldMap["sta_uni7"] = s.StaUni7
	s.fieldMap["sta_uni8"] = s.StaUni8
	s.fieldMap["sta_uni9"] = s.StaUni9
	s.fieldMap["sta_uni10"] = s.StaUni10
	s.fieldMap["sta_uni11"] = s.StaUni11
	s.fieldMap["sta_uni12"] = s.StaUni12
	s.fieldMap["sta_uni13"] = s.StaUni13
	s.fieldMap["sta_uni14"] = s.StaUni14
}

func (s staUni) clone(db *gorm.DB) staUni {
	s.staUniDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s staUni) replaceDB(db *gorm.DB) staUni {
	s.staUniDo.ReplaceDB(db)
	return s
}

type staUniDo struct{ gen.DO }

type IStaUniDo interface {
	gen.SubQuery
	Debug() IStaUniDo
	WithContext(ctx context.Context) IStaUniDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStaUniDo
	WriteDB() IStaUniDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStaUniDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStaUniDo
	Not(conds ...gen.Condition) IStaUniDo
	Or(conds ...gen.Condition) IStaUniDo
	Select(conds ...field.Expr) IStaUniDo
	Where(conds ...gen.Condition) IStaUniDo
	Order(conds ...field.Expr) IStaUniDo
	Distinct(cols ...field.Expr) IStaUniDo
	Omit(cols ...field.Expr) IStaUniDo
	Join(table schema.Tabler, on ...field.Expr) IStaUniDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStaUniDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStaUniDo
	Group(cols ...field.Expr) IStaUniDo
	Having(conds ...gen.Condition) IStaUniDo
	Limit(limit int) IStaUniDo
	Offset(offset int) IStaUniDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStaUniDo
	Unscoped() IStaUniDo
	Create(values ...*slaughter.StaUni) error
	CreateInBatches(values []*slaughter.StaUni, batchSize int) error
	Save(values ...*slaughter.StaUni) error
	First() (*slaughter.StaUni, error)
	Take() (*slaughter.StaUni, error)
	Last() (*slaughter.StaUni, error)
	Find() ([]*slaughter.StaUni, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.StaUni, err error)
	FindInBatches(result *[]*slaughter.StaUni, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.StaUni) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStaUniDo
	Assign(attrs ...field.AssignExpr) IStaUniDo
	Joins(fields ...field.RelationField) IStaUniDo
	Preload(fields ...field.RelationField) IStaUniDo
	FirstOrInit() (*slaughter.StaUni, error)
	FirstOrCreate() (*slaughter.StaUni, error)
	FindByPage(offset int, limit int) (result []*slaughter.StaUni, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStaUniDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s staUniDo) Debug() IStaUniDo {
	return s.withDO(s.DO.Debug())
}

func (s staUniDo) WithContext(ctx context.Context) IStaUniDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s staUniDo) ReadDB() IStaUniDo {
	return s.Clauses(dbresolver.Read)
}

func (s staUniDo) WriteDB() IStaUniDo {
	return s.Clauses(dbresolver.Write)
}

func (s staUniDo) Session(config *gorm.Session) IStaUniDo {
	return s.withDO(s.DO.Session(config))
}

func (s staUniDo) Clauses(conds ...clause.Expression) IStaUniDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s staUniDo) Returning(value interface{}, columns ...string) IStaUniDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s staUniDo) Not(conds ...gen.Condition) IStaUniDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s staUniDo) Or(conds ...gen.Condition) IStaUniDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s staUniDo) Select(conds ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s staUniDo) Where(conds ...gen.Condition) IStaUniDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s staUniDo) Order(conds ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s staUniDo) Distinct(cols ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s staUniDo) Omit(cols ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s staUniDo) Join(table schema.Tabler, on ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s staUniDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s staUniDo) RightJoin(table schema.Tabler, on ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s staUniDo) Group(cols ...field.Expr) IStaUniDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s staUniDo) Having(conds ...gen.Condition) IStaUniDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s staUniDo) Limit(limit int) IStaUniDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s staUniDo) Offset(offset int) IStaUniDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s staUniDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStaUniDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s staUniDo) Unscoped() IStaUniDo {
	return s.withDO(s.DO.Unscoped())
}

func (s staUniDo) Create(values ...*slaughter.StaUni) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s staUniDo) CreateInBatches(values []*slaughter.StaUni, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s staUniDo) Save(values ...*slaughter.StaUni) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s staUniDo) First() (*slaughter.StaUni, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.StaUni), nil
	}
}

func (s staUniDo) Take() (*slaughter.StaUni, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.StaUni), nil
	}
}

func (s staUniDo) Last() (*slaughter.StaUni, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.StaUni), nil
	}
}

func (s staUniDo) Find() ([]*slaughter.StaUni, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.StaUni), err
}

func (s staUniDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.StaUni, err error) {
	buf := make([]*slaughter.StaUni, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s staUniDo) FindInBatches(result *[]*slaughter.StaUni, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s staUniDo) Attrs(attrs ...field.AssignExpr) IStaUniDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s staUniDo) Assign(attrs ...field.AssignExpr) IStaUniDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s staUniDo) Joins(fields ...field.RelationField) IStaUniDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s staUniDo) Preload(fields ...field.RelationField) IStaUniDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s staUniDo) FirstOrInit() (*slaughter.StaUni, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.StaUni), nil
	}
}

func (s staUniDo) FirstOrCreate() (*slaughter.StaUni, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.StaUni), nil
	}
}

func (s staUniDo) FindByPage(offset int, limit int) (result []*slaughter.StaUni, count int64, err error) {
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

func (s staUniDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s staUniDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s staUniDo) Delete(models ...*slaughter.StaUni) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *staUniDo) withDO(do gen.Dao) *staUniDo {
	s.DO = *do.(*gen.DO)
	return s
}
