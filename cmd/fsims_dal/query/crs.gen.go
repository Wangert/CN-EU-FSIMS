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

func newCr(db *gorm.DB, opts ...gen.DOOption) cr {
	_cr := cr{}

	_cr.crDo.UseDB(db, opts...)
	_cr.crDo.UseModel(&pasture.Cr{})

	tableName := _cr.crDo.TableName()
	_cr.ALL = field.NewAsterisk(tableName)
	_cr.ID = field.NewUint(tableName, "id")
	_cr.CreatedAt = field.NewTime(tableName, "created_at")
	_cr.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cr.DeletedAt = field.NewField(tableName, "deleted_at")
	_cr.HeavyMetalID = field.NewUint(tableName, "heavy_metal_id")
	_cr.Cr1 = field.NewFloat64(tableName, "cr1")
	_cr.Cr2 = field.NewFloat64(tableName, "cr2")
	_cr.Cr3 = field.NewFloat64(tableName, "cr3")
	_cr.Cr4 = field.NewFloat64(tableName, "cr4")

	_cr.fillFieldMap()

	return _cr
}

type cr struct {
	crDo crDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HeavyMetalID field.Uint
	Cr1          field.Float64
	Cr2          field.Float64
	Cr3          field.Float64
	Cr4          field.Float64

	fieldMap map[string]field.Expr
}

func (c cr) Table(newTableName string) *cr {
	c.crDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cr) As(alias string) *cr {
	c.crDo.DO = *(c.crDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cr) updateTableName(table string) *cr {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.HeavyMetalID = field.NewUint(table, "heavy_metal_id")
	c.Cr1 = field.NewFloat64(table, "cr1")
	c.Cr2 = field.NewFloat64(table, "cr2")
	c.Cr3 = field.NewFloat64(table, "cr3")
	c.Cr4 = field.NewFloat64(table, "cr4")

	c.fillFieldMap()

	return c
}

func (c *cr) WithContext(ctx context.Context) ICrDo { return c.crDo.WithContext(ctx) }

func (c cr) TableName() string { return c.crDo.TableName() }

func (c cr) Alias() string { return c.crDo.Alias() }

func (c cr) Columns(cols ...field.Expr) gen.Columns { return c.crDo.Columns(cols...) }

func (c *cr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cr) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["heavy_metal_id"] = c.HeavyMetalID
	c.fieldMap["cr1"] = c.Cr1
	c.fieldMap["cr2"] = c.Cr2
	c.fieldMap["cr3"] = c.Cr3
	c.fieldMap["cr4"] = c.Cr4
}

func (c cr) clone(db *gorm.DB) cr {
	c.crDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cr) replaceDB(db *gorm.DB) cr {
	c.crDo.ReplaceDB(db)
	return c
}

type crDo struct{ gen.DO }

type ICrDo interface {
	gen.SubQuery
	Debug() ICrDo
	WithContext(ctx context.Context) ICrDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICrDo
	WriteDB() ICrDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICrDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICrDo
	Not(conds ...gen.Condition) ICrDo
	Or(conds ...gen.Condition) ICrDo
	Select(conds ...field.Expr) ICrDo
	Where(conds ...gen.Condition) ICrDo
	Order(conds ...field.Expr) ICrDo
	Distinct(cols ...field.Expr) ICrDo
	Omit(cols ...field.Expr) ICrDo
	Join(table schema.Tabler, on ...field.Expr) ICrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICrDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICrDo
	Group(cols ...field.Expr) ICrDo
	Having(conds ...gen.Condition) ICrDo
	Limit(limit int) ICrDo
	Offset(offset int) ICrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICrDo
	Unscoped() ICrDo
	Create(values ...*pasture.Cr) error
	CreateInBatches(values []*pasture.Cr, batchSize int) error
	Save(values ...*pasture.Cr) error
	First() (*pasture.Cr, error)
	Take() (*pasture.Cr, error)
	Last() (*pasture.Cr, error)
	Find() ([]*pasture.Cr, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.Cr, err error)
	FindInBatches(result *[]*pasture.Cr, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.Cr) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICrDo
	Assign(attrs ...field.AssignExpr) ICrDo
	Joins(fields ...field.RelationField) ICrDo
	Preload(fields ...field.RelationField) ICrDo
	FirstOrInit() (*pasture.Cr, error)
	FirstOrCreate() (*pasture.Cr, error)
	FindByPage(offset int, limit int) (result []*pasture.Cr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c crDo) Debug() ICrDo {
	return c.withDO(c.DO.Debug())
}

func (c crDo) WithContext(ctx context.Context) ICrDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c crDo) ReadDB() ICrDo {
	return c.Clauses(dbresolver.Read)
}

func (c crDo) WriteDB() ICrDo {
	return c.Clauses(dbresolver.Write)
}

func (c crDo) Session(config *gorm.Session) ICrDo {
	return c.withDO(c.DO.Session(config))
}

func (c crDo) Clauses(conds ...clause.Expression) ICrDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c crDo) Returning(value interface{}, columns ...string) ICrDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c crDo) Not(conds ...gen.Condition) ICrDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c crDo) Or(conds ...gen.Condition) ICrDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c crDo) Select(conds ...field.Expr) ICrDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c crDo) Where(conds ...gen.Condition) ICrDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c crDo) Order(conds ...field.Expr) ICrDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c crDo) Distinct(cols ...field.Expr) ICrDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c crDo) Omit(cols ...field.Expr) ICrDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c crDo) Join(table schema.Tabler, on ...field.Expr) ICrDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c crDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICrDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c crDo) RightJoin(table schema.Tabler, on ...field.Expr) ICrDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c crDo) Group(cols ...field.Expr) ICrDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c crDo) Having(conds ...gen.Condition) ICrDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c crDo) Limit(limit int) ICrDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c crDo) Offset(offset int) ICrDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c crDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICrDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c crDo) Unscoped() ICrDo {
	return c.withDO(c.DO.Unscoped())
}

func (c crDo) Create(values ...*pasture.Cr) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c crDo) CreateInBatches(values []*pasture.Cr, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c crDo) Save(values ...*pasture.Cr) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c crDo) First() (*pasture.Cr, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Cr), nil
	}
}

func (c crDo) Take() (*pasture.Cr, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Cr), nil
	}
}

func (c crDo) Last() (*pasture.Cr, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Cr), nil
	}
}

func (c crDo) Find() ([]*pasture.Cr, error) {
	result, err := c.DO.Find()
	return result.([]*pasture.Cr), err
}

func (c crDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.Cr, err error) {
	buf := make([]*pasture.Cr, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c crDo) FindInBatches(result *[]*pasture.Cr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c crDo) Attrs(attrs ...field.AssignExpr) ICrDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c crDo) Assign(attrs ...field.AssignExpr) ICrDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c crDo) Joins(fields ...field.RelationField) ICrDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c crDo) Preload(fields ...field.RelationField) ICrDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c crDo) FirstOrInit() (*pasture.Cr, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Cr), nil
	}
}

func (c crDo) FirstOrCreate() (*pasture.Cr, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Cr), nil
	}
}

func (c crDo) FindByPage(offset int, limit int) (result []*pasture.Cr, count int64, err error) {
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

func (c crDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c crDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c crDo) Delete(models ...*pasture.Cr) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *crDo) withDO(do gen.Dao) *crDo {
	c.DO = *do.(*gen.DO)
	return c
}
