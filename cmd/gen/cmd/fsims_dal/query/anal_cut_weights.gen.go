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

func newAnalCutWeight(db *gorm.DB, opts ...gen.DOOption) analCutWeight {
	_analCutWeight := analCutWeight{}

	_analCutWeight.analCutWeightDo.UseDB(db, opts...)
	_analCutWeight.analCutWeightDo.UseModel(&slaughter.AnalCutWeight{})

	tableName := _analCutWeight.analCutWeightDo.TableName()
	_analCutWeight.ALL = field.NewAsterisk(tableName)
	_analCutWeight.ID = field.NewUint(tableName, "id")
	_analCutWeight.CreatedAt = field.NewTime(tableName, "created_at")
	_analCutWeight.UpdatedAt = field.NewTime(tableName, "updated_at")
	_analCutWeight.DeletedAt = field.NewField(tableName, "deleted_at")
	_analCutWeight.PID = field.NewString(tableName, "p_id")
	_analCutWeight.AnalCutWeight1 = field.NewFloat32(tableName, "anal_cut_weight1")
	_analCutWeight.AnalCutWeight2 = field.NewFloat32(tableName, "anal_cut_weight2")
	_analCutWeight.AnalCutWeight3 = field.NewFloat32(tableName, "anal_cut_weight3")
	_analCutWeight.AnalCutWeight4 = field.NewFloat32(tableName, "anal_cut_weight4")
	_analCutWeight.AnalCutWeight5 = field.NewFloat32(tableName, "anal_cut_weight5")
	_analCutWeight.AnalCutWeight6 = field.NewFloat32(tableName, "anal_cut_weight6")
	_analCutWeight.AnalCutWeight7 = field.NewFloat32(tableName, "anal_cut_weight7")
	_analCutWeight.AnalCutWeight8 = field.NewFloat32(tableName, "anal_cut_weight8")
	_analCutWeight.AnalCutWeight9 = field.NewFloat32(tableName, "anal_cut_weight9")
	_analCutWeight.AnalCutWeight10 = field.NewFloat32(tableName, "anal_cut_weight10")
	_analCutWeight.AnalCutWeight11 = field.NewFloat32(tableName, "anal_cut_weight11")
	_analCutWeight.AnalCutWeight12 = field.NewFloat32(tableName, "anal_cut_weight12")

	_analCutWeight.fillFieldMap()

	return _analCutWeight
}

type analCutWeight struct {
	analCutWeightDo analCutWeightDo

	ALL             field.Asterisk
	ID              field.Uint
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field
	PID             field.String
	AnalCutWeight1  field.Float32
	AnalCutWeight2  field.Float32
	AnalCutWeight3  field.Float32
	AnalCutWeight4  field.Float32
	AnalCutWeight5  field.Float32
	AnalCutWeight6  field.Float32
	AnalCutWeight7  field.Float32
	AnalCutWeight8  field.Float32
	AnalCutWeight9  field.Float32
	AnalCutWeight10 field.Float32
	AnalCutWeight11 field.Float32
	AnalCutWeight12 field.Float32

	fieldMap map[string]field.Expr
}

func (a analCutWeight) Table(newTableName string) *analCutWeight {
	a.analCutWeightDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a analCutWeight) As(alias string) *analCutWeight {
	a.analCutWeightDo.DO = *(a.analCutWeightDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *analCutWeight) updateTableName(table string) *analCutWeight {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.PID = field.NewString(table, "p_id")
	a.AnalCutWeight1 = field.NewFloat32(table, "anal_cut_weight1")
	a.AnalCutWeight2 = field.NewFloat32(table, "anal_cut_weight2")
	a.AnalCutWeight3 = field.NewFloat32(table, "anal_cut_weight3")
	a.AnalCutWeight4 = field.NewFloat32(table, "anal_cut_weight4")
	a.AnalCutWeight5 = field.NewFloat32(table, "anal_cut_weight5")
	a.AnalCutWeight6 = field.NewFloat32(table, "anal_cut_weight6")
	a.AnalCutWeight7 = field.NewFloat32(table, "anal_cut_weight7")
	a.AnalCutWeight8 = field.NewFloat32(table, "anal_cut_weight8")
	a.AnalCutWeight9 = field.NewFloat32(table, "anal_cut_weight9")
	a.AnalCutWeight10 = field.NewFloat32(table, "anal_cut_weight10")
	a.AnalCutWeight11 = field.NewFloat32(table, "anal_cut_weight11")
	a.AnalCutWeight12 = field.NewFloat32(table, "anal_cut_weight12")

	a.fillFieldMap()

	return a
}

func (a *analCutWeight) WithContext(ctx context.Context) IAnalCutWeightDo {
	return a.analCutWeightDo.WithContext(ctx)
}

func (a analCutWeight) TableName() string { return a.analCutWeightDo.TableName() }

func (a analCutWeight) Alias() string { return a.analCutWeightDo.Alias() }

func (a analCutWeight) Columns(cols ...field.Expr) gen.Columns {
	return a.analCutWeightDo.Columns(cols...)
}

func (a *analCutWeight) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *analCutWeight) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 17)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["p_id"] = a.PID
	a.fieldMap["anal_cut_weight1"] = a.AnalCutWeight1
	a.fieldMap["anal_cut_weight2"] = a.AnalCutWeight2
	a.fieldMap["anal_cut_weight3"] = a.AnalCutWeight3
	a.fieldMap["anal_cut_weight4"] = a.AnalCutWeight4
	a.fieldMap["anal_cut_weight5"] = a.AnalCutWeight5
	a.fieldMap["anal_cut_weight6"] = a.AnalCutWeight6
	a.fieldMap["anal_cut_weight7"] = a.AnalCutWeight7
	a.fieldMap["anal_cut_weight8"] = a.AnalCutWeight8
	a.fieldMap["anal_cut_weight9"] = a.AnalCutWeight9
	a.fieldMap["anal_cut_weight10"] = a.AnalCutWeight10
	a.fieldMap["anal_cut_weight11"] = a.AnalCutWeight11
	a.fieldMap["anal_cut_weight12"] = a.AnalCutWeight12
}

func (a analCutWeight) clone(db *gorm.DB) analCutWeight {
	a.analCutWeightDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a analCutWeight) replaceDB(db *gorm.DB) analCutWeight {
	a.analCutWeightDo.ReplaceDB(db)
	return a
}

type analCutWeightDo struct{ gen.DO }

type IAnalCutWeightDo interface {
	gen.SubQuery
	Debug() IAnalCutWeightDo
	WithContext(ctx context.Context) IAnalCutWeightDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnalCutWeightDo
	WriteDB() IAnalCutWeightDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnalCutWeightDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnalCutWeightDo
	Not(conds ...gen.Condition) IAnalCutWeightDo
	Or(conds ...gen.Condition) IAnalCutWeightDo
	Select(conds ...field.Expr) IAnalCutWeightDo
	Where(conds ...gen.Condition) IAnalCutWeightDo
	Order(conds ...field.Expr) IAnalCutWeightDo
	Distinct(cols ...field.Expr) IAnalCutWeightDo
	Omit(cols ...field.Expr) IAnalCutWeightDo
	Join(table schema.Tabler, on ...field.Expr) IAnalCutWeightDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnalCutWeightDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnalCutWeightDo
	Group(cols ...field.Expr) IAnalCutWeightDo
	Having(conds ...gen.Condition) IAnalCutWeightDo
	Limit(limit int) IAnalCutWeightDo
	Offset(offset int) IAnalCutWeightDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnalCutWeightDo
	Unscoped() IAnalCutWeightDo
	Create(values ...*slaughter.AnalCutWeight) error
	CreateInBatches(values []*slaughter.AnalCutWeight, batchSize int) error
	Save(values ...*slaughter.AnalCutWeight) error
	First() (*slaughter.AnalCutWeight, error)
	Take() (*slaughter.AnalCutWeight, error)
	Last() (*slaughter.AnalCutWeight, error)
	Find() ([]*slaughter.AnalCutWeight, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.AnalCutWeight, err error)
	FindInBatches(result *[]*slaughter.AnalCutWeight, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.AnalCutWeight) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnalCutWeightDo
	Assign(attrs ...field.AssignExpr) IAnalCutWeightDo
	Joins(fields ...field.RelationField) IAnalCutWeightDo
	Preload(fields ...field.RelationField) IAnalCutWeightDo
	FirstOrInit() (*slaughter.AnalCutWeight, error)
	FirstOrCreate() (*slaughter.AnalCutWeight, error)
	FindByPage(offset int, limit int) (result []*slaughter.AnalCutWeight, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnalCutWeightDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a analCutWeightDo) Debug() IAnalCutWeightDo {
	return a.withDO(a.DO.Debug())
}

func (a analCutWeightDo) WithContext(ctx context.Context) IAnalCutWeightDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a analCutWeightDo) ReadDB() IAnalCutWeightDo {
	return a.Clauses(dbresolver.Read)
}

func (a analCutWeightDo) WriteDB() IAnalCutWeightDo {
	return a.Clauses(dbresolver.Write)
}

func (a analCutWeightDo) Session(config *gorm.Session) IAnalCutWeightDo {
	return a.withDO(a.DO.Session(config))
}

func (a analCutWeightDo) Clauses(conds ...clause.Expression) IAnalCutWeightDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a analCutWeightDo) Returning(value interface{}, columns ...string) IAnalCutWeightDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a analCutWeightDo) Not(conds ...gen.Condition) IAnalCutWeightDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a analCutWeightDo) Or(conds ...gen.Condition) IAnalCutWeightDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a analCutWeightDo) Select(conds ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a analCutWeightDo) Where(conds ...gen.Condition) IAnalCutWeightDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a analCutWeightDo) Order(conds ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a analCutWeightDo) Distinct(cols ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a analCutWeightDo) Omit(cols ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a analCutWeightDo) Join(table schema.Tabler, on ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a analCutWeightDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a analCutWeightDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a analCutWeightDo) Group(cols ...field.Expr) IAnalCutWeightDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a analCutWeightDo) Having(conds ...gen.Condition) IAnalCutWeightDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a analCutWeightDo) Limit(limit int) IAnalCutWeightDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a analCutWeightDo) Offset(offset int) IAnalCutWeightDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a analCutWeightDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnalCutWeightDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a analCutWeightDo) Unscoped() IAnalCutWeightDo {
	return a.withDO(a.DO.Unscoped())
}

func (a analCutWeightDo) Create(values ...*slaughter.AnalCutWeight) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a analCutWeightDo) CreateInBatches(values []*slaughter.AnalCutWeight, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a analCutWeightDo) Save(values ...*slaughter.AnalCutWeight) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a analCutWeightDo) First() (*slaughter.AnalCutWeight, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalCutWeight), nil
	}
}

func (a analCutWeightDo) Take() (*slaughter.AnalCutWeight, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalCutWeight), nil
	}
}

func (a analCutWeightDo) Last() (*slaughter.AnalCutWeight, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalCutWeight), nil
	}
}

func (a analCutWeightDo) Find() ([]*slaughter.AnalCutWeight, error) {
	result, err := a.DO.Find()
	return result.([]*slaughter.AnalCutWeight), err
}

func (a analCutWeightDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.AnalCutWeight, err error) {
	buf := make([]*slaughter.AnalCutWeight, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a analCutWeightDo) FindInBatches(result *[]*slaughter.AnalCutWeight, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a analCutWeightDo) Attrs(attrs ...field.AssignExpr) IAnalCutWeightDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a analCutWeightDo) Assign(attrs ...field.AssignExpr) IAnalCutWeightDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a analCutWeightDo) Joins(fields ...field.RelationField) IAnalCutWeightDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a analCutWeightDo) Preload(fields ...field.RelationField) IAnalCutWeightDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a analCutWeightDo) FirstOrInit() (*slaughter.AnalCutWeight, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalCutWeight), nil
	}
}

func (a analCutWeightDo) FirstOrCreate() (*slaughter.AnalCutWeight, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AnalCutWeight), nil
	}
}

func (a analCutWeightDo) FindByPage(offset int, limit int) (result []*slaughter.AnalCutWeight, count int64, err error) {
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

func (a analCutWeightDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a analCutWeightDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a analCutWeightDo) Delete(models ...*slaughter.AnalCutWeight) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *analCutWeightDo) withDO(do gen.Dao) *analCutWeightDo {
	a.DO = *do.(*gen.DO)
	return a
}
