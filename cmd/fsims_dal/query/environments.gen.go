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

func newEnvironment(db *gorm.DB, opts ...gen.DOOption) environment {
	_environment := environment{}

	_environment.environmentDo.UseDB(db, opts...)
	_environment.environmentDo.UseModel(&pasture.Environment{})

	tableName := _environment.environmentDo.TableName()
	_environment.ALL = field.NewAsterisk(tableName)
	_environment.ID = field.NewUint(tableName, "id")
	_environment.CreatedAt = field.NewTime(tableName, "created_at")
	_environment.UpdatedAt = field.NewTime(tableName, "updated_at")
	_environment.DeletedAt = field.NewField(tableName, "deleted_at")
	_environment.FarmEnvironmentID = field.NewUint(tableName, "farm_environment_id")
	_environment.Environment1 = field.NewFloat32(tableName, "environment1")
	_environment.Environment2 = field.NewFloat32(tableName, "environment2")
	_environment.Environment3 = field.NewFloat32(tableName, "environment3")
	_environment.Environment4 = field.NewFloat32(tableName, "environment4")
	_environment.Environment5 = field.NewFloat32(tableName, "environment5")
	_environment.Environment6 = field.NewFloat32(tableName, "environment6")

	_environment.fillFieldMap()

	return _environment
}

type environment struct {
	environmentDo environmentDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	FarmEnvironmentID field.Uint
	Environment1      field.Float32
	Environment2      field.Float32
	Environment3      field.Float32
	Environment4      field.Float32
	Environment5      field.Float32
	Environment6      field.Float32

	fieldMap map[string]field.Expr
}

func (e environment) Table(newTableName string) *environment {
	e.environmentDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e environment) As(alias string) *environment {
	e.environmentDo.DO = *(e.environmentDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *environment) updateTableName(table string) *environment {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewUint(table, "id")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")
	e.FarmEnvironmentID = field.NewUint(table, "farm_environment_id")
	e.Environment1 = field.NewFloat32(table, "environment1")
	e.Environment2 = field.NewFloat32(table, "environment2")
	e.Environment3 = field.NewFloat32(table, "environment3")
	e.Environment4 = field.NewFloat32(table, "environment4")
	e.Environment5 = field.NewFloat32(table, "environment5")
	e.Environment6 = field.NewFloat32(table, "environment6")

	e.fillFieldMap()

	return e
}

func (e *environment) WithContext(ctx context.Context) IEnvironmentDo {
	return e.environmentDo.WithContext(ctx)
}

func (e environment) TableName() string { return e.environmentDo.TableName() }

func (e environment) Alias() string { return e.environmentDo.Alias() }

func (e environment) Columns(cols ...field.Expr) gen.Columns { return e.environmentDo.Columns(cols...) }

func (e *environment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *environment) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 11)
	e.fieldMap["id"] = e.ID
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
	e.fieldMap["farm_environment_id"] = e.FarmEnvironmentID
	e.fieldMap["environment1"] = e.Environment1
	e.fieldMap["environment2"] = e.Environment2
	e.fieldMap["environment3"] = e.Environment3
	e.fieldMap["environment4"] = e.Environment4
	e.fieldMap["environment5"] = e.Environment5
	e.fieldMap["environment6"] = e.Environment6
}

func (e environment) clone(db *gorm.DB) environment {
	e.environmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e environment) replaceDB(db *gorm.DB) environment {
	e.environmentDo.ReplaceDB(db)
	return e
}

type environmentDo struct{ gen.DO }

type IEnvironmentDo interface {
	gen.SubQuery
	Debug() IEnvironmentDo
	WithContext(ctx context.Context) IEnvironmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEnvironmentDo
	WriteDB() IEnvironmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEnvironmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEnvironmentDo
	Not(conds ...gen.Condition) IEnvironmentDo
	Or(conds ...gen.Condition) IEnvironmentDo
	Select(conds ...field.Expr) IEnvironmentDo
	Where(conds ...gen.Condition) IEnvironmentDo
	Order(conds ...field.Expr) IEnvironmentDo
	Distinct(cols ...field.Expr) IEnvironmentDo
	Omit(cols ...field.Expr) IEnvironmentDo
	Join(table schema.Tabler, on ...field.Expr) IEnvironmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEnvironmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEnvironmentDo
	Group(cols ...field.Expr) IEnvironmentDo
	Having(conds ...gen.Condition) IEnvironmentDo
	Limit(limit int) IEnvironmentDo
	Offset(offset int) IEnvironmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEnvironmentDo
	Unscoped() IEnvironmentDo
	Create(values ...*pasture.Environment) error
	CreateInBatches(values []*pasture.Environment, batchSize int) error
	Save(values ...*pasture.Environment) error
	First() (*pasture.Environment, error)
	Take() (*pasture.Environment, error)
	Last() (*pasture.Environment, error)
	Find() ([]*pasture.Environment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.Environment, err error)
	FindInBatches(result *[]*pasture.Environment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.Environment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEnvironmentDo
	Assign(attrs ...field.AssignExpr) IEnvironmentDo
	Joins(fields ...field.RelationField) IEnvironmentDo
	Preload(fields ...field.RelationField) IEnvironmentDo
	FirstOrInit() (*pasture.Environment, error)
	FirstOrCreate() (*pasture.Environment, error)
	FindByPage(offset int, limit int) (result []*pasture.Environment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEnvironmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e environmentDo) Debug() IEnvironmentDo {
	return e.withDO(e.DO.Debug())
}

func (e environmentDo) WithContext(ctx context.Context) IEnvironmentDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e environmentDo) ReadDB() IEnvironmentDo {
	return e.Clauses(dbresolver.Read)
}

func (e environmentDo) WriteDB() IEnvironmentDo {
	return e.Clauses(dbresolver.Write)
}

func (e environmentDo) Session(config *gorm.Session) IEnvironmentDo {
	return e.withDO(e.DO.Session(config))
}

func (e environmentDo) Clauses(conds ...clause.Expression) IEnvironmentDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e environmentDo) Returning(value interface{}, columns ...string) IEnvironmentDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e environmentDo) Not(conds ...gen.Condition) IEnvironmentDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e environmentDo) Or(conds ...gen.Condition) IEnvironmentDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e environmentDo) Select(conds ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e environmentDo) Where(conds ...gen.Condition) IEnvironmentDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e environmentDo) Order(conds ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e environmentDo) Distinct(cols ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e environmentDo) Omit(cols ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e environmentDo) Join(table schema.Tabler, on ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e environmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e environmentDo) RightJoin(table schema.Tabler, on ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e environmentDo) Group(cols ...field.Expr) IEnvironmentDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e environmentDo) Having(conds ...gen.Condition) IEnvironmentDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e environmentDo) Limit(limit int) IEnvironmentDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e environmentDo) Offset(offset int) IEnvironmentDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e environmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEnvironmentDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e environmentDo) Unscoped() IEnvironmentDo {
	return e.withDO(e.DO.Unscoped())
}

func (e environmentDo) Create(values ...*pasture.Environment) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e environmentDo) CreateInBatches(values []*pasture.Environment, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e environmentDo) Save(values ...*pasture.Environment) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e environmentDo) First() (*pasture.Environment, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Environment), nil
	}
}

func (e environmentDo) Take() (*pasture.Environment, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Environment), nil
	}
}

func (e environmentDo) Last() (*pasture.Environment, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Environment), nil
	}
}

func (e environmentDo) Find() ([]*pasture.Environment, error) {
	result, err := e.DO.Find()
	return result.([]*pasture.Environment), err
}

func (e environmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.Environment, err error) {
	buf := make([]*pasture.Environment, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e environmentDo) FindInBatches(result *[]*pasture.Environment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e environmentDo) Attrs(attrs ...field.AssignExpr) IEnvironmentDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e environmentDo) Assign(attrs ...field.AssignExpr) IEnvironmentDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e environmentDo) Joins(fields ...field.RelationField) IEnvironmentDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e environmentDo) Preload(fields ...field.RelationField) IEnvironmentDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e environmentDo) FirstOrInit() (*pasture.Environment, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Environment), nil
	}
}

func (e environmentDo) FirstOrCreate() (*pasture.Environment, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Environment), nil
	}
}

func (e environmentDo) FindByPage(offset int, limit int) (result []*pasture.Environment, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e environmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e environmentDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e environmentDo) Delete(models ...*pasture.Environment) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *environmentDo) withDO(do gen.Dao) *environmentDo {
	e.DO = *do.(*gen.DO)
	return e
}
