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

func newWastedWaterIndex(db *gorm.DB, opts ...gen.DOOption) wastedWaterIndex {
	_wastedWaterIndex := wastedWaterIndex{}

	_wastedWaterIndex.wastedWaterIndexDo.UseDB(db, opts...)
	_wastedWaterIndex.wastedWaterIndexDo.UseModel(&pasture.WastedWaterIndex{})

	tableName := _wastedWaterIndex.wastedWaterIndexDo.TableName()
	_wastedWaterIndex.ALL = field.NewAsterisk(tableName)
	_wastedWaterIndex.ID = field.NewUint(tableName, "id")
	_wastedWaterIndex.CreatedAt = field.NewTime(tableName, "created_at")
	_wastedWaterIndex.UpdatedAt = field.NewTime(tableName, "updated_at")
	_wastedWaterIndex.DeletedAt = field.NewField(tableName, "deleted_at")
	_wastedWaterIndex.FarmEnvironmentID = field.NewUint(tableName, "farm_environment_id")
	_wastedWaterIndex.WastedWaterIndex1 = field.NewFloat32(tableName, "wasted_water_index1")
	_wastedWaterIndex.WastedWaterIndex2 = field.NewFloat32(tableName, "wasted_water_index2")
	_wastedWaterIndex.WastedWaterIndex3 = field.NewFloat32(tableName, "wasted_water_index3")
	_wastedWaterIndex.WastedWaterIndex4 = field.NewFloat32(tableName, "wasted_water_index4")
	_wastedWaterIndex.WastedWaterIndex5 = field.NewFloat32(tableName, "wasted_water_index5")
	_wastedWaterIndex.WastedWaterIndex6 = field.NewFloat32(tableName, "wasted_water_index6")
	_wastedWaterIndex.WastedWaterIndex7 = field.NewFloat32(tableName, "wasted_water_index7")
	_wastedWaterIndex.WastedWaterIndex8 = field.NewFloat32(tableName, "wasted_water_index8")
	_wastedWaterIndex.WastedWaterIndex9 = field.NewFloat32(tableName, "wasted_water_index9")

	_wastedWaterIndex.fillFieldMap()

	return _wastedWaterIndex
}

type wastedWaterIndex struct {
	wastedWaterIndexDo wastedWaterIndexDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	FarmEnvironmentID field.Uint
	WastedWaterIndex1 field.Float32
	WastedWaterIndex2 field.Float32
	WastedWaterIndex3 field.Float32
	WastedWaterIndex4 field.Float32
	WastedWaterIndex5 field.Float32
	WastedWaterIndex6 field.Float32
	WastedWaterIndex7 field.Float32
	WastedWaterIndex8 field.Float32
	WastedWaterIndex9 field.Float32

	fieldMap map[string]field.Expr
}

func (w wastedWaterIndex) Table(newTableName string) *wastedWaterIndex {
	w.wastedWaterIndexDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wastedWaterIndex) As(alias string) *wastedWaterIndex {
	w.wastedWaterIndexDo.DO = *(w.wastedWaterIndexDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wastedWaterIndex) updateTableName(table string) *wastedWaterIndex {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.FarmEnvironmentID = field.NewUint(table, "farm_environment_id")
	w.WastedWaterIndex1 = field.NewFloat32(table, "wasted_water_index1")
	w.WastedWaterIndex2 = field.NewFloat32(table, "wasted_water_index2")
	w.WastedWaterIndex3 = field.NewFloat32(table, "wasted_water_index3")
	w.WastedWaterIndex4 = field.NewFloat32(table, "wasted_water_index4")
	w.WastedWaterIndex5 = field.NewFloat32(table, "wasted_water_index5")
	w.WastedWaterIndex6 = field.NewFloat32(table, "wasted_water_index6")
	w.WastedWaterIndex7 = field.NewFloat32(table, "wasted_water_index7")
	w.WastedWaterIndex8 = field.NewFloat32(table, "wasted_water_index8")
	w.WastedWaterIndex9 = field.NewFloat32(table, "wasted_water_index9")

	w.fillFieldMap()

	return w
}

func (w *wastedWaterIndex) WithContext(ctx context.Context) IWastedWaterIndexDo {
	return w.wastedWaterIndexDo.WithContext(ctx)
}

func (w wastedWaterIndex) TableName() string { return w.wastedWaterIndexDo.TableName() }

func (w wastedWaterIndex) Alias() string { return w.wastedWaterIndexDo.Alias() }

func (w wastedWaterIndex) Columns(cols ...field.Expr) gen.Columns {
	return w.wastedWaterIndexDo.Columns(cols...)
}

func (w *wastedWaterIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wastedWaterIndex) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["farm_environment_id"] = w.FarmEnvironmentID
	w.fieldMap["wasted_water_index1"] = w.WastedWaterIndex1
	w.fieldMap["wasted_water_index2"] = w.WastedWaterIndex2
	w.fieldMap["wasted_water_index3"] = w.WastedWaterIndex3
	w.fieldMap["wasted_water_index4"] = w.WastedWaterIndex4
	w.fieldMap["wasted_water_index5"] = w.WastedWaterIndex5
	w.fieldMap["wasted_water_index6"] = w.WastedWaterIndex6
	w.fieldMap["wasted_water_index7"] = w.WastedWaterIndex7
	w.fieldMap["wasted_water_index8"] = w.WastedWaterIndex8
	w.fieldMap["wasted_water_index9"] = w.WastedWaterIndex9
}

func (w wastedWaterIndex) clone(db *gorm.DB) wastedWaterIndex {
	w.wastedWaterIndexDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wastedWaterIndex) replaceDB(db *gorm.DB) wastedWaterIndex {
	w.wastedWaterIndexDo.ReplaceDB(db)
	return w
}

type wastedWaterIndexDo struct{ gen.DO }

type IWastedWaterIndexDo interface {
	gen.SubQuery
	Debug() IWastedWaterIndexDo
	WithContext(ctx context.Context) IWastedWaterIndexDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWastedWaterIndexDo
	WriteDB() IWastedWaterIndexDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWastedWaterIndexDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWastedWaterIndexDo
	Not(conds ...gen.Condition) IWastedWaterIndexDo
	Or(conds ...gen.Condition) IWastedWaterIndexDo
	Select(conds ...field.Expr) IWastedWaterIndexDo
	Where(conds ...gen.Condition) IWastedWaterIndexDo
	Order(conds ...field.Expr) IWastedWaterIndexDo
	Distinct(cols ...field.Expr) IWastedWaterIndexDo
	Omit(cols ...field.Expr) IWastedWaterIndexDo
	Join(table schema.Tabler, on ...field.Expr) IWastedWaterIndexDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWastedWaterIndexDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWastedWaterIndexDo
	Group(cols ...field.Expr) IWastedWaterIndexDo
	Having(conds ...gen.Condition) IWastedWaterIndexDo
	Limit(limit int) IWastedWaterIndexDo
	Offset(offset int) IWastedWaterIndexDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWastedWaterIndexDo
	Unscoped() IWastedWaterIndexDo
	Create(values ...*pasture.WastedWaterIndex) error
	CreateInBatches(values []*pasture.WastedWaterIndex, batchSize int) error
	Save(values ...*pasture.WastedWaterIndex) error
	First() (*pasture.WastedWaterIndex, error)
	Take() (*pasture.WastedWaterIndex, error)
	Last() (*pasture.WastedWaterIndex, error)
	Find() ([]*pasture.WastedWaterIndex, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.WastedWaterIndex, err error)
	FindInBatches(result *[]*pasture.WastedWaterIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.WastedWaterIndex) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWastedWaterIndexDo
	Assign(attrs ...field.AssignExpr) IWastedWaterIndexDo
	Joins(fields ...field.RelationField) IWastedWaterIndexDo
	Preload(fields ...field.RelationField) IWastedWaterIndexDo
	FirstOrInit() (*pasture.WastedWaterIndex, error)
	FirstOrCreate() (*pasture.WastedWaterIndex, error)
	FindByPage(offset int, limit int) (result []*pasture.WastedWaterIndex, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWastedWaterIndexDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wastedWaterIndexDo) Debug() IWastedWaterIndexDo {
	return w.withDO(w.DO.Debug())
}

func (w wastedWaterIndexDo) WithContext(ctx context.Context) IWastedWaterIndexDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wastedWaterIndexDo) ReadDB() IWastedWaterIndexDo {
	return w.Clauses(dbresolver.Read)
}

func (w wastedWaterIndexDo) WriteDB() IWastedWaterIndexDo {
	return w.Clauses(dbresolver.Write)
}

func (w wastedWaterIndexDo) Session(config *gorm.Session) IWastedWaterIndexDo {
	return w.withDO(w.DO.Session(config))
}

func (w wastedWaterIndexDo) Clauses(conds ...clause.Expression) IWastedWaterIndexDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wastedWaterIndexDo) Returning(value interface{}, columns ...string) IWastedWaterIndexDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wastedWaterIndexDo) Not(conds ...gen.Condition) IWastedWaterIndexDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wastedWaterIndexDo) Or(conds ...gen.Condition) IWastedWaterIndexDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wastedWaterIndexDo) Select(conds ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wastedWaterIndexDo) Where(conds ...gen.Condition) IWastedWaterIndexDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wastedWaterIndexDo) Order(conds ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wastedWaterIndexDo) Distinct(cols ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wastedWaterIndexDo) Omit(cols ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wastedWaterIndexDo) Join(table schema.Tabler, on ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wastedWaterIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wastedWaterIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wastedWaterIndexDo) Group(cols ...field.Expr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wastedWaterIndexDo) Having(conds ...gen.Condition) IWastedWaterIndexDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wastedWaterIndexDo) Limit(limit int) IWastedWaterIndexDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wastedWaterIndexDo) Offset(offset int) IWastedWaterIndexDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wastedWaterIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWastedWaterIndexDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wastedWaterIndexDo) Unscoped() IWastedWaterIndexDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wastedWaterIndexDo) Create(values ...*pasture.WastedWaterIndex) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wastedWaterIndexDo) CreateInBatches(values []*pasture.WastedWaterIndex, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wastedWaterIndexDo) Save(values ...*pasture.WastedWaterIndex) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wastedWaterIndexDo) First() (*pasture.WastedWaterIndex, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WastedWaterIndex), nil
	}
}

func (w wastedWaterIndexDo) Take() (*pasture.WastedWaterIndex, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WastedWaterIndex), nil
	}
}

func (w wastedWaterIndexDo) Last() (*pasture.WastedWaterIndex, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WastedWaterIndex), nil
	}
}

func (w wastedWaterIndexDo) Find() ([]*pasture.WastedWaterIndex, error) {
	result, err := w.DO.Find()
	return result.([]*pasture.WastedWaterIndex), err
}

func (w wastedWaterIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.WastedWaterIndex, err error) {
	buf := make([]*pasture.WastedWaterIndex, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wastedWaterIndexDo) FindInBatches(result *[]*pasture.WastedWaterIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wastedWaterIndexDo) Attrs(attrs ...field.AssignExpr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wastedWaterIndexDo) Assign(attrs ...field.AssignExpr) IWastedWaterIndexDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wastedWaterIndexDo) Joins(fields ...field.RelationField) IWastedWaterIndexDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wastedWaterIndexDo) Preload(fields ...field.RelationField) IWastedWaterIndexDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wastedWaterIndexDo) FirstOrInit() (*pasture.WastedWaterIndex, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WastedWaterIndex), nil
	}
}

func (w wastedWaterIndexDo) FirstOrCreate() (*pasture.WastedWaterIndex, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WastedWaterIndex), nil
	}
}

func (w wastedWaterIndexDo) FindByPage(offset int, limit int) (result []*pasture.WastedWaterIndex, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wastedWaterIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wastedWaterIndexDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wastedWaterIndexDo) Delete(models ...*pasture.WastedWaterIndex) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wastedWaterIndexDo) withDO(do gen.Dao) *wastedWaterIndexDo {
	w.DO = *do.(*gen.DO)
	return w
}
