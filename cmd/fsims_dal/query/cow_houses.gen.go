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

func newCowHouse(db *gorm.DB, opts ...gen.DOOption) cowHouse {
	_cowHouse := cowHouse{}

	_cowHouse.cowHouseDo.UseDB(db, opts...)
	_cowHouse.cowHouseDo.UseModel(&pasture.CowHouse{})

	tableName := _cowHouse.cowHouseDo.TableName()
	_cowHouse.ALL = field.NewAsterisk(tableName)
	_cowHouse.ID = field.NewUint(tableName, "id")
	_cowHouse.CreatedAt = field.NewTime(tableName, "created_at")
	_cowHouse.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cowHouse.DeletedAt = field.NewField(tableName, "deleted_at")
	_cowHouse.FarmEnvironmentID = field.NewUint(tableName, "farm_environment_id")
	_cowHouse.CowHouse1 = field.NewFloat32(tableName, "cow_house1")
	_cowHouse.CowHouse2 = field.NewFloat32(tableName, "cow_house2")
	_cowHouse.CowHouse3 = field.NewFloat32(tableName, "cow_house3")
	_cowHouse.CowHouse4 = field.NewFloat32(tableName, "cow_house4")
	_cowHouse.CowHouse5 = field.NewFloat32(tableName, "cow_house5")
	_cowHouse.CowHouse6 = field.NewFloat32(tableName, "cow_house6")
	_cowHouse.CowHouse7 = field.NewFloat32(tableName, "cow_house7")
	_cowHouse.CowHouse8 = field.NewFloat32(tableName, "cow_house8")
	_cowHouse.CowHouse9 = field.NewFloat32(tableName, "cow_house9")
	_cowHouse.CowHouse10 = field.NewFloat32(tableName, "cow_house10")
	_cowHouse.CowHouse11 = field.NewFloat32(tableName, "cow_house11")
	_cowHouse.CowHouse12 = field.NewFloat32(tableName, "cow_house12")

	_cowHouse.fillFieldMap()

	return _cowHouse
}

type cowHouse struct {
	cowHouseDo cowHouseDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	FarmEnvironmentID field.Uint
	CowHouse1         field.Float32
	CowHouse2         field.Float32
	CowHouse3         field.Float32
	CowHouse4         field.Float32
	CowHouse5         field.Float32
	CowHouse6         field.Float32
	CowHouse7         field.Float32
	CowHouse8         field.Float32
	CowHouse9         field.Float32
	CowHouse10        field.Float32
	CowHouse11        field.Float32
	CowHouse12        field.Float32

	fieldMap map[string]field.Expr
}

func (c cowHouse) Table(newTableName string) *cowHouse {
	c.cowHouseDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cowHouse) As(alias string) *cowHouse {
	c.cowHouseDo.DO = *(c.cowHouseDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cowHouse) updateTableName(table string) *cowHouse {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.FarmEnvironmentID = field.NewUint(table, "farm_environment_id")
	c.CowHouse1 = field.NewFloat32(table, "cow_house1")
	c.CowHouse2 = field.NewFloat32(table, "cow_house2")
	c.CowHouse3 = field.NewFloat32(table, "cow_house3")
	c.CowHouse4 = field.NewFloat32(table, "cow_house4")
	c.CowHouse5 = field.NewFloat32(table, "cow_house5")
	c.CowHouse6 = field.NewFloat32(table, "cow_house6")
	c.CowHouse7 = field.NewFloat32(table, "cow_house7")
	c.CowHouse8 = field.NewFloat32(table, "cow_house8")
	c.CowHouse9 = field.NewFloat32(table, "cow_house9")
	c.CowHouse10 = field.NewFloat32(table, "cow_house10")
	c.CowHouse11 = field.NewFloat32(table, "cow_house11")
	c.CowHouse12 = field.NewFloat32(table, "cow_house12")

	c.fillFieldMap()

	return c
}

func (c *cowHouse) WithContext(ctx context.Context) ICowHouseDo { return c.cowHouseDo.WithContext(ctx) }

func (c cowHouse) TableName() string { return c.cowHouseDo.TableName() }

func (c cowHouse) Alias() string { return c.cowHouseDo.Alias() }

func (c cowHouse) Columns(cols ...field.Expr) gen.Columns { return c.cowHouseDo.Columns(cols...) }

func (c *cowHouse) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cowHouse) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 17)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["farm_environment_id"] = c.FarmEnvironmentID
	c.fieldMap["cow_house1"] = c.CowHouse1
	c.fieldMap["cow_house2"] = c.CowHouse2
	c.fieldMap["cow_house3"] = c.CowHouse3
	c.fieldMap["cow_house4"] = c.CowHouse4
	c.fieldMap["cow_house5"] = c.CowHouse5
	c.fieldMap["cow_house6"] = c.CowHouse6
	c.fieldMap["cow_house7"] = c.CowHouse7
	c.fieldMap["cow_house8"] = c.CowHouse8
	c.fieldMap["cow_house9"] = c.CowHouse9
	c.fieldMap["cow_house10"] = c.CowHouse10
	c.fieldMap["cow_house11"] = c.CowHouse11
	c.fieldMap["cow_house12"] = c.CowHouse12
}

func (c cowHouse) clone(db *gorm.DB) cowHouse {
	c.cowHouseDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cowHouse) replaceDB(db *gorm.DB) cowHouse {
	c.cowHouseDo.ReplaceDB(db)
	return c
}

type cowHouseDo struct{ gen.DO }

type ICowHouseDo interface {
	gen.SubQuery
	Debug() ICowHouseDo
	WithContext(ctx context.Context) ICowHouseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICowHouseDo
	WriteDB() ICowHouseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICowHouseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICowHouseDo
	Not(conds ...gen.Condition) ICowHouseDo
	Or(conds ...gen.Condition) ICowHouseDo
	Select(conds ...field.Expr) ICowHouseDo
	Where(conds ...gen.Condition) ICowHouseDo
	Order(conds ...field.Expr) ICowHouseDo
	Distinct(cols ...field.Expr) ICowHouseDo
	Omit(cols ...field.Expr) ICowHouseDo
	Join(table schema.Tabler, on ...field.Expr) ICowHouseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICowHouseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICowHouseDo
	Group(cols ...field.Expr) ICowHouseDo
	Having(conds ...gen.Condition) ICowHouseDo
	Limit(limit int) ICowHouseDo
	Offset(offset int) ICowHouseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICowHouseDo
	Unscoped() ICowHouseDo
	Create(values ...*pasture.CowHouse) error
	CreateInBatches(values []*pasture.CowHouse, batchSize int) error
	Save(values ...*pasture.CowHouse) error
	First() (*pasture.CowHouse, error)
	Take() (*pasture.CowHouse, error)
	Last() (*pasture.CowHouse, error)
	Find() ([]*pasture.CowHouse, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.CowHouse, err error)
	FindInBatches(result *[]*pasture.CowHouse, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.CowHouse) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICowHouseDo
	Assign(attrs ...field.AssignExpr) ICowHouseDo
	Joins(fields ...field.RelationField) ICowHouseDo
	Preload(fields ...field.RelationField) ICowHouseDo
	FirstOrInit() (*pasture.CowHouse, error)
	FirstOrCreate() (*pasture.CowHouse, error)
	FindByPage(offset int, limit int) (result []*pasture.CowHouse, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICowHouseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cowHouseDo) Debug() ICowHouseDo {
	return c.withDO(c.DO.Debug())
}

func (c cowHouseDo) WithContext(ctx context.Context) ICowHouseDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cowHouseDo) ReadDB() ICowHouseDo {
	return c.Clauses(dbresolver.Read)
}

func (c cowHouseDo) WriteDB() ICowHouseDo {
	return c.Clauses(dbresolver.Write)
}

func (c cowHouseDo) Session(config *gorm.Session) ICowHouseDo {
	return c.withDO(c.DO.Session(config))
}

func (c cowHouseDo) Clauses(conds ...clause.Expression) ICowHouseDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cowHouseDo) Returning(value interface{}, columns ...string) ICowHouseDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cowHouseDo) Not(conds ...gen.Condition) ICowHouseDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cowHouseDo) Or(conds ...gen.Condition) ICowHouseDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cowHouseDo) Select(conds ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cowHouseDo) Where(conds ...gen.Condition) ICowHouseDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cowHouseDo) Order(conds ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cowHouseDo) Distinct(cols ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cowHouseDo) Omit(cols ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cowHouseDo) Join(table schema.Tabler, on ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cowHouseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cowHouseDo) RightJoin(table schema.Tabler, on ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cowHouseDo) Group(cols ...field.Expr) ICowHouseDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cowHouseDo) Having(conds ...gen.Condition) ICowHouseDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cowHouseDo) Limit(limit int) ICowHouseDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cowHouseDo) Offset(offset int) ICowHouseDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cowHouseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICowHouseDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cowHouseDo) Unscoped() ICowHouseDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cowHouseDo) Create(values ...*pasture.CowHouse) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cowHouseDo) CreateInBatches(values []*pasture.CowHouse, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cowHouseDo) Save(values ...*pasture.CowHouse) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cowHouseDo) First() (*pasture.CowHouse, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.CowHouse), nil
	}
}

func (c cowHouseDo) Take() (*pasture.CowHouse, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.CowHouse), nil
	}
}

func (c cowHouseDo) Last() (*pasture.CowHouse, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.CowHouse), nil
	}
}

func (c cowHouseDo) Find() ([]*pasture.CowHouse, error) {
	result, err := c.DO.Find()
	return result.([]*pasture.CowHouse), err
}

func (c cowHouseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.CowHouse, err error) {
	buf := make([]*pasture.CowHouse, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cowHouseDo) FindInBatches(result *[]*pasture.CowHouse, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cowHouseDo) Attrs(attrs ...field.AssignExpr) ICowHouseDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cowHouseDo) Assign(attrs ...field.AssignExpr) ICowHouseDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cowHouseDo) Joins(fields ...field.RelationField) ICowHouseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cowHouseDo) Preload(fields ...field.RelationField) ICowHouseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cowHouseDo) FirstOrInit() (*pasture.CowHouse, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.CowHouse), nil
	}
}

func (c cowHouseDo) FirstOrCreate() (*pasture.CowHouse, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.CowHouse), nil
	}
}

func (c cowHouseDo) FindByPage(offset int, limit int) (result []*pasture.CowHouse, count int64, err error) {
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

func (c cowHouseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cowHouseDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cowHouseDo) Delete(models ...*pasture.CowHouse) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cowHouseDo) withDO(do gen.Dao) *cowHouseDo {
	c.DO = *do.(*gen.DO)
	return c
}
