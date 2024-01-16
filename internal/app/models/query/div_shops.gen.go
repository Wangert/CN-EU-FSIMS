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

func newDivShop(db *gorm.DB, opts ...gen.DOOption) divShop {
	_divShop := divShop{}

	_divShop.divShopDo.UseDB(db, opts...)
	_divShop.divShopDo.UseModel(&slaughter.DivShop{})

	tableName := _divShop.divShopDo.TableName()
	_divShop.ALL = field.NewAsterisk(tableName)
	_divShop.ID = field.NewUint(tableName, "id")
	_divShop.CreatedAt = field.NewTime(tableName, "created_at")
	_divShop.UpdatedAt = field.NewTime(tableName, "updated_at")
	_divShop.DeletedAt = field.NewField(tableName, "deleted_at")
	_divShop.HouseNumber = field.NewString(tableName, "house_number")
	_divShop.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_divShop.DivShop1 = field.NewFloat64(tableName, "div_shop1")
	_divShop.DivShop2 = field.NewFloat64(tableName, "div_shop2")
	_divShop.DivShop3 = field.NewFloat64(tableName, "div_shop3")
	_divShop.DivShop4 = field.NewFloat64(tableName, "div_shop4")
	_divShop.DivShop5 = field.NewFloat64(tableName, "div_shop5")
	_divShop.DivShop6 = field.NewFloat64(tableName, "div_shop6")
	_divShop.DivShop7 = field.NewFloat64(tableName, "div_shop7")
	_divShop.DivShop8 = field.NewFloat64(tableName, "div_shop8")
	_divShop.DivShop9 = field.NewFloat64(tableName, "div_shop9")
	_divShop.DivShop10 = field.NewString(tableName, "div_shop10")
	_divShop.DivShop11 = field.NewString(tableName, "div_shop11")
	_divShop.DivShop12 = field.NewString(tableName, "div_shop12")
	_divShop.DivShop13 = field.NewString(tableName, "div_shop13")
	_divShop.DivShop14 = field.NewString(tableName, "div_shop14")

	_divShop.fillFieldMap()

	return _divShop
}

type divShop struct {
	divShopDo divShopDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HouseNumber  field.String
	TimeRecordAt field.Time
	DivShop1     field.Float64
	DivShop2     field.Float64
	DivShop3     field.Float64
	DivShop4     field.Float64
	DivShop5     field.Float64
	DivShop6     field.Float64
	DivShop7     field.Float64
	DivShop8     field.Float64
	DivShop9     field.Float64
	DivShop10    field.String
	DivShop11    field.String
	DivShop12    field.String
	DivShop13    field.String
	DivShop14    field.String

	fieldMap map[string]field.Expr
}

func (d divShop) Table(newTableName string) *divShop {
	d.divShopDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d divShop) As(alias string) *divShop {
	d.divShopDo.DO = *(d.divShopDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *divShop) updateTableName(table string) *divShop {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewUint(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.HouseNumber = field.NewString(table, "house_number")
	d.TimeRecordAt = field.NewTime(table, "time_record_at")
	d.DivShop1 = field.NewFloat64(table, "div_shop1")
	d.DivShop2 = field.NewFloat64(table, "div_shop2")
	d.DivShop3 = field.NewFloat64(table, "div_shop3")
	d.DivShop4 = field.NewFloat64(table, "div_shop4")
	d.DivShop5 = field.NewFloat64(table, "div_shop5")
	d.DivShop6 = field.NewFloat64(table, "div_shop6")
	d.DivShop7 = field.NewFloat64(table, "div_shop7")
	d.DivShop8 = field.NewFloat64(table, "div_shop8")
	d.DivShop9 = field.NewFloat64(table, "div_shop9")
	d.DivShop10 = field.NewString(table, "div_shop10")
	d.DivShop11 = field.NewString(table, "div_shop11")
	d.DivShop12 = field.NewString(table, "div_shop12")
	d.DivShop13 = field.NewString(table, "div_shop13")
	d.DivShop14 = field.NewString(table, "div_shop14")

	d.fillFieldMap()

	return d
}

func (d *divShop) WithContext(ctx context.Context) IDivShopDo { return d.divShopDo.WithContext(ctx) }

func (d divShop) TableName() string { return d.divShopDo.TableName() }

func (d divShop) Alias() string { return d.divShopDo.Alias() }

func (d divShop) Columns(cols ...field.Expr) gen.Columns { return d.divShopDo.Columns(cols...) }

func (d *divShop) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *divShop) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 20)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["house_number"] = d.HouseNumber
	d.fieldMap["time_record_at"] = d.TimeRecordAt
	d.fieldMap["div_shop1"] = d.DivShop1
	d.fieldMap["div_shop2"] = d.DivShop2
	d.fieldMap["div_shop3"] = d.DivShop3
	d.fieldMap["div_shop4"] = d.DivShop4
	d.fieldMap["div_shop5"] = d.DivShop5
	d.fieldMap["div_shop6"] = d.DivShop6
	d.fieldMap["div_shop7"] = d.DivShop7
	d.fieldMap["div_shop8"] = d.DivShop8
	d.fieldMap["div_shop9"] = d.DivShop9
	d.fieldMap["div_shop10"] = d.DivShop10
	d.fieldMap["div_shop11"] = d.DivShop11
	d.fieldMap["div_shop12"] = d.DivShop12
	d.fieldMap["div_shop13"] = d.DivShop13
	d.fieldMap["div_shop14"] = d.DivShop14
}

func (d divShop) clone(db *gorm.DB) divShop {
	d.divShopDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d divShop) replaceDB(db *gorm.DB) divShop {
	d.divShopDo.ReplaceDB(db)
	return d
}

type divShopDo struct{ gen.DO }

type IDivShopDo interface {
	gen.SubQuery
	Debug() IDivShopDo
	WithContext(ctx context.Context) IDivShopDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDivShopDo
	WriteDB() IDivShopDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDivShopDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDivShopDo
	Not(conds ...gen.Condition) IDivShopDo
	Or(conds ...gen.Condition) IDivShopDo
	Select(conds ...field.Expr) IDivShopDo
	Where(conds ...gen.Condition) IDivShopDo
	Order(conds ...field.Expr) IDivShopDo
	Distinct(cols ...field.Expr) IDivShopDo
	Omit(cols ...field.Expr) IDivShopDo
	Join(table schema.Tabler, on ...field.Expr) IDivShopDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDivShopDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDivShopDo
	Group(cols ...field.Expr) IDivShopDo
	Having(conds ...gen.Condition) IDivShopDo
	Limit(limit int) IDivShopDo
	Offset(offset int) IDivShopDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDivShopDo
	Unscoped() IDivShopDo
	Create(values ...*slaughter.DivShop) error
	CreateInBatches(values []*slaughter.DivShop, batchSize int) error
	Save(values ...*slaughter.DivShop) error
	First() (*slaughter.DivShop, error)
	Take() (*slaughter.DivShop, error)
	Last() (*slaughter.DivShop, error)
	Find() ([]*slaughter.DivShop, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.DivShop, err error)
	FindInBatches(result *[]*slaughter.DivShop, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.DivShop) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDivShopDo
	Assign(attrs ...field.AssignExpr) IDivShopDo
	Joins(fields ...field.RelationField) IDivShopDo
	Preload(fields ...field.RelationField) IDivShopDo
	FirstOrInit() (*slaughter.DivShop, error)
	FirstOrCreate() (*slaughter.DivShop, error)
	FindByPage(offset int, limit int) (result []*slaughter.DivShop, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDivShopDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d divShopDo) Debug() IDivShopDo {
	return d.withDO(d.DO.Debug())
}

func (d divShopDo) WithContext(ctx context.Context) IDivShopDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d divShopDo) ReadDB() IDivShopDo {
	return d.Clauses(dbresolver.Read)
}

func (d divShopDo) WriteDB() IDivShopDo {
	return d.Clauses(dbresolver.Write)
}

func (d divShopDo) Session(config *gorm.Session) IDivShopDo {
	return d.withDO(d.DO.Session(config))
}

func (d divShopDo) Clauses(conds ...clause.Expression) IDivShopDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d divShopDo) Returning(value interface{}, columns ...string) IDivShopDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d divShopDo) Not(conds ...gen.Condition) IDivShopDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d divShopDo) Or(conds ...gen.Condition) IDivShopDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d divShopDo) Select(conds ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d divShopDo) Where(conds ...gen.Condition) IDivShopDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d divShopDo) Order(conds ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d divShopDo) Distinct(cols ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d divShopDo) Omit(cols ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d divShopDo) Join(table schema.Tabler, on ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d divShopDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d divShopDo) RightJoin(table schema.Tabler, on ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d divShopDo) Group(cols ...field.Expr) IDivShopDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d divShopDo) Having(conds ...gen.Condition) IDivShopDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d divShopDo) Limit(limit int) IDivShopDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d divShopDo) Offset(offset int) IDivShopDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d divShopDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDivShopDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d divShopDo) Unscoped() IDivShopDo {
	return d.withDO(d.DO.Unscoped())
}

func (d divShopDo) Create(values ...*slaughter.DivShop) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d divShopDo) CreateInBatches(values []*slaughter.DivShop, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d divShopDo) Save(values ...*slaughter.DivShop) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d divShopDo) First() (*slaughter.DivShop, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.DivShop), nil
	}
}

func (d divShopDo) Take() (*slaughter.DivShop, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.DivShop), nil
	}
}

func (d divShopDo) Last() (*slaughter.DivShop, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.DivShop), nil
	}
}

func (d divShopDo) Find() ([]*slaughter.DivShop, error) {
	result, err := d.DO.Find()
	return result.([]*slaughter.DivShop), err
}

func (d divShopDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.DivShop, err error) {
	buf := make([]*slaughter.DivShop, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d divShopDo) FindInBatches(result *[]*slaughter.DivShop, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d divShopDo) Attrs(attrs ...field.AssignExpr) IDivShopDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d divShopDo) Assign(attrs ...field.AssignExpr) IDivShopDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d divShopDo) Joins(fields ...field.RelationField) IDivShopDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d divShopDo) Preload(fields ...field.RelationField) IDivShopDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d divShopDo) FirstOrInit() (*slaughter.DivShop, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.DivShop), nil
	}
}

func (d divShopDo) FirstOrCreate() (*slaughter.DivShop, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.DivShop), nil
	}
}

func (d divShopDo) FindByPage(offset int, limit int) (result []*slaughter.DivShop, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d divShopDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d divShopDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d divShopDo) Delete(models ...*slaughter.DivShop) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *divShopDo) withDO(do gen.Dao) *divShopDo {
	d.DO = *do.(*gen.DO)
	return d
}
