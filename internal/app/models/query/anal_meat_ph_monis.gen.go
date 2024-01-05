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

func newAnalMeatPhMoni(db *gorm.DB, opts ...gen.DOOption) analMeatPhMoni {
	_analMeatPhMoni := analMeatPhMoni{}

	_analMeatPhMoni.analMeatPhMoniDo.UseDB(db, opts...)
	_analMeatPhMoni.analMeatPhMoniDo.UseModel(&slaughter.AnalMeatPhMoni{})

	tableName := _analMeatPhMoni.analMeatPhMoniDo.TableName()
	_analMeatPhMoni.ALL = field.NewAsterisk(tableName)
	_analMeatPhMoni.ID = field.NewUint(tableName, "id")
	_analMeatPhMoni.CreatedAt = field.NewTime(tableName, "created_at")
	_analMeatPhMoni.UpdatedAt = field.NewTime(tableName, "updated_at")
	_analMeatPhMoni.DeletedAt = field.NewField(tableName, "deleted_at")
	_analMeatPhMoni.PID = field.NewString(tableName, "p_id")
	_analMeatPhMoni.AnalMeatPhMoni1 = field.NewFloat32(tableName, "anal_meat_ph_moni1")
	_analMeatPhMoni.AnalMeatPhMoni2 = field.NewFloat32(tableName, "anal_meat_ph_moni2")
	_analMeatPhMoni.AnalMeatPhMoni3 = field.NewFloat32(tableName, "anal_meat_ph_moni3")
	_analMeatPhMoni.AnalMeatPhMoni4 = field.NewFloat32(tableName, "anal_meat_ph_moni4")
	_analMeatPhMoni.AnalMeatPhMoni5 = field.NewFloat32(tableName, "anal_meat_ph_moni5")

	_analMeatPhMoni.fillFieldMap()

	return _analMeatPhMoni
}

type analMeatPhMoni struct {
	analMeatPhMoniDo analMeatPhMoniDo

	ALL             field.Asterisk
	ID              field.Uint
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field
	PID             field.String
	AnalMeatPhMoni1 field.Float32
	AnalMeatPhMoni2 field.Float32
	AnalMeatPhMoni3 field.Float32
	AnalMeatPhMoni4 field.Float32
	AnalMeatPhMoni5 field.Float32

	fieldMap map[string]field.Expr
}

func (a analMeatPhMoni) Table(newTableName string) *analMeatPhMoni {
	a.analMeatPhMoniDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a analMeatPhMoni) As(alias string) *analMeatPhMoni {
	a.analMeatPhMoniDo.DO = *(a.analMeatPhMoniDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *analMeatPhMoni) updateTableName(table string) *analMeatPhMoni {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.PID = field.NewString(table, "p_id")
	a.AnalMeatPhMoni1 = field.NewFloat32(table, "anal_meat_ph_moni1")
	a.AnalMeatPhMoni2 = field.NewFloat32(table, "anal_meat_ph_moni2")
	a.AnalMeatPhMoni3 = field.NewFloat32(table, "anal_meat_ph_moni3")
	a.AnalMeatPhMoni4 = field.NewFloat32(table, "anal_meat_ph_moni4")
	a.AnalMeatPhMoni5 = field.NewFloat32(table, "anal_meat_ph_moni5")

	a.fillFieldMap()

	return a
}

func (a *analMeatPhMoni) WithContext(ctx context.Context) IAnalMeatPhMoniDo {
	return a.analMeatPhMoniDo.WithContext(ctx)
}

func (a analMeatPhMoni) TableName() string { return a.analMeatPhMoniDo.TableName() }

func (a analMeatPhMoni) Alias() string { return a.analMeatPhMoniDo.Alias() }

func (a analMeatPhMoni) Columns(cols ...field.Expr) gen.Columns {
	return a.analMeatPhMoniDo.Columns(cols...)
}

func (a *analMeatPhMoni) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *analMeatPhMoni) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["p_id"] = a.PID
	a.fieldMap["anal_meat_ph_moni1"] = a.AnalMeatPhMoni1
	a.fieldMap["anal_meat_ph_moni2"] = a.AnalMeatPhMoni2
	a.fieldMap["anal_meat_ph_moni3"] = a.AnalMeatPhMoni3
	a.fieldMap["anal_meat_ph_moni4"] = a.AnalMeatPhMoni4
	a.fieldMap["anal_meat_ph_moni5"] = a.AnalMeatPhMoni5
}

func (a analMeatPhMoni) clone(db *gorm.DB) analMeatPhMoni {
	a.analMeatPhMoniDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a analMeatPhMoni) replaceDB(db *gorm.DB) analMeatPhMoni {
	a.analMeatPhMoniDo.ReplaceDB(db)
	return a
}

type analMeatPhMoniDo struct{ gen.DO }

type IAnalMeatPhMoniDo interface {
	gen.SubQuery
	Debug() IAnalMeatPhMoniDo
	WithContext(ctx context.Context) IAnalMeatPhMoniDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnalMeatPhMoniDo
	WriteDB() IAnalMeatPhMoniDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnalMeatPhMoniDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnalMeatPhMoniDo
	Not(conds ...gen.Condition) IAnalMeatPhMoniDo
	Or(conds ...gen.Condition) IAnalMeatPhMoniDo
	Select(conds ...field.Expr) IAnalMeatPhMoniDo
	Where(conds ...gen.Condition) IAnalMeatPhMoniDo
	Order(conds ...field.Expr) IAnalMeatPhMoniDo
	Distinct(cols ...field.Expr) IAnalMeatPhMoniDo
	Omit(cols ...field.Expr) IAnalMeatPhMoniDo
	Join(table schema.Tabler, on ...field.Expr) IAnalMeatPhMoniDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnalMeatPhMoniDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnalMeatPhMoniDo
	Group(cols ...field.Expr) IAnalMeatPhMoniDo
	Having(conds ...gen.Condition) IAnalMeatPhMoniDo
	Limit(limit int) IAnalMeatPhMoniDo
	Offset(offset int) IAnalMeatPhMoniDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnalMeatPhMoniDo
	Unscoped() IAnalMeatPhMoniDo
	Create(values ...*slaughter.AnalMeatPhMoni) error
	CreateInBatches(values []*slaughter.AnalMeatPhMoni, batchSize int) error
	Save(values ...*slaughter.AnalMeatPhMoni) error
	First() (*slaughter.AnalMeatPhMoni, error)
	Take() (*slaughter.AnalMeatPhMoni, error)
	Last() (*slaughter.AnalMeatPhMoni, error)
	Find() ([]*slaughter.AnalMeatPhMoni, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.AnalMeatPhMoni, err error)
	FindInBatches(result *[]*slaughter.AnalMeatPhMoni, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.AnalMeatPhMoni) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnalMeatPhMoniDo
	Assign(attrs ...field.AssignExpr) IAnalMeatPhMoniDo
	Joins(fields ...field.RelationField) IAnalMeatPhMoniDo
	Preload(fields ...field.RelationField) IAnalMeatPhMoniDo
	FirstOrInit() (*slaughter.AnalMeatPhMoni, error)
	FirstOrCreate() (*slaughter.AnalMeatPhMoni, error)
	FindByPage(offset int, limit int) (result []*slaughter.AnalMeatPhMoni, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnalMeatPhMoniDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a analMeatPhMoniDo) Debug() IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Debug())
}

func (a analMeatPhMoniDo) WithContext(ctx context.Context) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a analMeatPhMoniDo) ReadDB() IAnalMeatPhMoniDo {
	return a.Clauses(dbresolver.Read)
}

func (a analMeatPhMoniDo) WriteDB() IAnalMeatPhMoniDo {
	return a.Clauses(dbresolver.Write)
}

func (a analMeatPhMoniDo) Session(config *gorm.Session) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Session(config))
}

func (a analMeatPhMoniDo) Clauses(conds ...clause.Expression) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a analMeatPhMoniDo) Returning(value interface{}, columns ...string) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a analMeatPhMoniDo) Not(conds ...gen.Condition) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a analMeatPhMoniDo) Or(conds ...gen.Condition) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a analMeatPhMoniDo) Select(conds ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a analMeatPhMoniDo) Where(conds ...gen.Condition) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a analMeatPhMoniDo) Order(conds ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a analMeatPhMoniDo) Distinct(cols ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a analMeatPhMoniDo) Omit(cols ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a analMeatPhMoniDo) Join(table schema.Tabler, on ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a analMeatPhMoniDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a analMeatPhMoniDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a analMeatPhMoniDo) Group(cols ...field.Expr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a analMeatPhMoniDo) Having(conds ...gen.Condition) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a analMeatPhMoniDo) Limit(limit int) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a analMeatPhMoniDo) Offset(offset int) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a analMeatPhMoniDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a analMeatPhMoniDo) Unscoped() IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Unscoped())
}

func (a analMeatPhMoniDo) Create(values ...*slaughter.AnalMeatPhMoni) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a analMeatPhMoniDo) CreateInBatches(values []*slaughter.AnalMeatPhMoni, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a analMeatPhMoniDo) Save(values ...*slaughter.AnalMeatPhMoni) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a analMeatPhMoniDo) First() (*slaughter.AnalMeatPhMoni, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalMeatPhMoni), nil
	}
}

func (a analMeatPhMoniDo) Take() (*slaughter.AnalMeatPhMoni, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalMeatPhMoni), nil
	}
}

func (a analMeatPhMoniDo) Last() (*slaughter.AnalMeatPhMoni, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalMeatPhMoni), nil
	}
}

func (a analMeatPhMoniDo) Find() ([]*slaughter.AnalMeatPhMoni, error) {
	result, err := a.DO.Find()
	return result.([]*slaughter.AnalMeatPhMoni), err
}

func (a analMeatPhMoniDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.AnalMeatPhMoni, err error) {
	buf := make([]*slaughter.AnalMeatPhMoni, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a analMeatPhMoniDo) FindInBatches(result *[]*slaughter.AnalMeatPhMoni, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a analMeatPhMoniDo) Attrs(attrs ...field.AssignExpr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a analMeatPhMoniDo) Assign(attrs ...field.AssignExpr) IAnalMeatPhMoniDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a analMeatPhMoniDo) Joins(fields ...field.RelationField) IAnalMeatPhMoniDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a analMeatPhMoniDo) Preload(fields ...field.RelationField) IAnalMeatPhMoniDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a analMeatPhMoniDo) FirstOrInit() (*slaughter.AnalMeatPhMoni, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalMeatPhMoni), nil
	}
}

func (a analMeatPhMoniDo) FirstOrCreate() (*slaughter.AnalMeatPhMoni, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalMeatPhMoni), nil
	}
}

func (a analMeatPhMoniDo) FindByPage(offset int, limit int) (result []*slaughter.AnalMeatPhMoni, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a analMeatPhMoniDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a analMeatPhMoniDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a analMeatPhMoniDo) Delete(models ...*slaughter.AnalMeatPhMoni) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *analMeatPhMoniDo) withDO(do gen.Dao) *analMeatPhMoniDo {
	a.DO = *do.(*gen.DO)
	return a
}
