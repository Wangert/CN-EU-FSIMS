// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCow(db *gorm.DB, opts ...gen.DOOption) cow {
	_cow := cow{}

	_cow.cowDo.UseDB(db, opts...)
	_cow.cowDo.UseModel(&product.Cow{})

	tableName := _cow.cowDo.TableName()
	_cow.ALL = field.NewAsterisk(tableName)
	_cow.ID = field.NewUint(tableName, "id")
	_cow.CreatedAt = field.NewTime(tableName, "created_at")
	_cow.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cow.DeletedAt = field.NewField(tableName, "deleted_at")
	_cow.Number = field.NewString(tableName, "number")
	_cow.Age = field.NewInt(tableName, "age")
	_cow.Weight = field.NewFloat64(tableName, "weight")
	_cow.State = field.NewInt(tableName, "state")
	_cow.HouseNumber = field.NewString(tableName, "house_number")
	_cow.BatchNumber = field.NewString(tableName, "batch_number")

	_cow.fillFieldMap()

	return _cow
}

type cow struct {
	cowDo cowDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Number      field.String
	Age         field.Int
	Weight      field.Float64
	State       field.Int
	HouseNumber field.String
	BatchNumber field.String

	fieldMap map[string]field.Expr
}

func (c cow) Table(newTableName string) *cow {
	c.cowDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cow) As(alias string) *cow {
	c.cowDo.DO = *(c.cowDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cow) updateTableName(table string) *cow {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Number = field.NewString(table, "number")
	c.Age = field.NewInt(table, "age")
	c.Weight = field.NewFloat64(table, "weight")
	c.State = field.NewInt(table, "state")
	c.HouseNumber = field.NewString(table, "house_number")
	c.BatchNumber = field.NewString(table, "batch_number")

	c.fillFieldMap()

	return c
}

func (c *cow) WithContext(ctx context.Context) ICowDo { return c.cowDo.WithContext(ctx) }

func (c cow) TableName() string { return c.cowDo.TableName() }

func (c cow) Alias() string { return c.cowDo.Alias() }

func (c cow) Columns(cols ...field.Expr) gen.Columns { return c.cowDo.Columns(cols...) }

func (c *cow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cow) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["number"] = c.Number
	c.fieldMap["age"] = c.Age
	c.fieldMap["weight"] = c.Weight
	c.fieldMap["state"] = c.State
	c.fieldMap["house_number"] = c.HouseNumber
	c.fieldMap["batch_number"] = c.BatchNumber
}

func (c cow) clone(db *gorm.DB) cow {
	c.cowDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cow) replaceDB(db *gorm.DB) cow {
	c.cowDo.ReplaceDB(db)
	return c
}

type cowDo struct{ gen.DO }

type ICowDo interface {
	gen.SubQuery
	Debug() ICowDo
	WithContext(ctx context.Context) ICowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICowDo
	WriteDB() ICowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICowDo
	Not(conds ...gen.Condition) ICowDo
	Or(conds ...gen.Condition) ICowDo
	Select(conds ...field.Expr) ICowDo
	Where(conds ...gen.Condition) ICowDo
	Order(conds ...field.Expr) ICowDo
	Distinct(cols ...field.Expr) ICowDo
	Omit(cols ...field.Expr) ICowDo
	Join(table schema.Tabler, on ...field.Expr) ICowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICowDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICowDo
	Group(cols ...field.Expr) ICowDo
	Having(conds ...gen.Condition) ICowDo
	Limit(limit int) ICowDo
	Offset(offset int) ICowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICowDo
	Unscoped() ICowDo
	Create(values ...*product.Cow) error
	CreateInBatches(values []*product.Cow, batchSize int) error
	Save(values ...*product.Cow) error
	First() (*product.Cow, error)
	Take() (*product.Cow, error)
	Last() (*product.Cow, error)
	Find() ([]*product.Cow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*product.Cow, err error)
	FindInBatches(result *[]*product.Cow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*product.Cow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICowDo
	Assign(attrs ...field.AssignExpr) ICowDo
	Joins(fields ...field.RelationField) ICowDo
	Preload(fields ...field.RelationField) ICowDo
	FirstOrInit() (*product.Cow, error)
	FirstOrCreate() (*product.Cow, error)
	FindByPage(offset int, limit int) (result []*product.Cow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cowDo) Debug() ICowDo {
	return c.withDO(c.DO.Debug())
}

func (c cowDo) WithContext(ctx context.Context) ICowDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cowDo) ReadDB() ICowDo {
	return c.Clauses(dbresolver.Read)
}

func (c cowDo) WriteDB() ICowDo {
	return c.Clauses(dbresolver.Write)
}

func (c cowDo) Session(config *gorm.Session) ICowDo {
	return c.withDO(c.DO.Session(config))
}

func (c cowDo) Clauses(conds ...clause.Expression) ICowDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cowDo) Returning(value interface{}, columns ...string) ICowDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cowDo) Not(conds ...gen.Condition) ICowDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cowDo) Or(conds ...gen.Condition) ICowDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cowDo) Select(conds ...field.Expr) ICowDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cowDo) Where(conds ...gen.Condition) ICowDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cowDo) Order(conds ...field.Expr) ICowDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cowDo) Distinct(cols ...field.Expr) ICowDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cowDo) Omit(cols ...field.Expr) ICowDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cowDo) Join(table schema.Tabler, on ...field.Expr) ICowDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cowDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICowDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cowDo) RightJoin(table schema.Tabler, on ...field.Expr) ICowDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cowDo) Group(cols ...field.Expr) ICowDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cowDo) Having(conds ...gen.Condition) ICowDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cowDo) Limit(limit int) ICowDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cowDo) Offset(offset int) ICowDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICowDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cowDo) Unscoped() ICowDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cowDo) Create(values ...*product.Cow) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cowDo) CreateInBatches(values []*product.Cow, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cowDo) Save(values ...*product.Cow) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cowDo) First() (*product.Cow, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*product.Cow), nil
	}
}

func (c cowDo) Take() (*product.Cow, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*product.Cow), nil
	}
}

func (c cowDo) Last() (*product.Cow, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*product.Cow), nil
	}
}

func (c cowDo) Find() ([]*product.Cow, error) {
	result, err := c.DO.Find()
	return result.([]*product.Cow), err
}

func (c cowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*product.Cow, err error) {
	buf := make([]*product.Cow, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cowDo) FindInBatches(result *[]*product.Cow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cowDo) Attrs(attrs ...field.AssignExpr) ICowDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cowDo) Assign(attrs ...field.AssignExpr) ICowDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cowDo) Joins(fields ...field.RelationField) ICowDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cowDo) Preload(fields ...field.RelationField) ICowDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cowDo) FirstOrInit() (*product.Cow, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*product.Cow), nil
	}
}

func (c cowDo) FirstOrCreate() (*product.Cow, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*product.Cow), nil
	}
}

func (c cowDo) FindByPage(offset int, limit int) (result []*product.Cow, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cowDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cowDo) Delete(models ...*product.Cow) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cowDo) withDO(do gen.Dao) *cowDo {
	c.DO = *do.(*gen.DO)
	return c
}