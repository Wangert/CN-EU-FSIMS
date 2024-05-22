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

func newAllPasturesTrashDisposal(db *gorm.DB, opts ...gen.DOOption) allPasturesTrashDisposal {
	_allPasturesTrashDisposal := allPasturesTrashDisposal{}

	_allPasturesTrashDisposal.allPasturesTrashDisposalDo.UseDB(db, opts...)
	_allPasturesTrashDisposal.allPasturesTrashDisposalDo.UseModel(&pasture.AllPasturesTrashDisposal{})

	tableName := _allPasturesTrashDisposal.allPasturesTrashDisposalDo.TableName()
	_allPasturesTrashDisposal.ALL = field.NewAsterisk(tableName)
	_allPasturesTrashDisposal.ID = field.NewUint(tableName, "id")
	_allPasturesTrashDisposal.CreatedAt = field.NewTime(tableName, "created_at")
	_allPasturesTrashDisposal.UpdatedAt = field.NewTime(tableName, "updated_at")
	_allPasturesTrashDisposal.DeletedAt = field.NewField(tableName, "deleted_at")
	_allPasturesTrashDisposal.TimeStamp = field.NewTime(tableName, "time_stamp")
	_allPasturesTrashDisposal.OdorPasturesTrashDisposal1 = field.NewFloat64(tableName, "odor_pastures_trash_disposal1")
	_allPasturesTrashDisposal.OdorPasturesTrashDisposal2 = field.NewFloat64(tableName, "odor_pastures_trash_disposal2")
	_allPasturesTrashDisposal.OdorPasturesTrashDisposal3 = field.NewFloat64(tableName, "odor_pastures_trash_disposal3")
	_allPasturesTrashDisposal.OdorPasturesTrashDisposal4 = field.NewFloat64(tableName, "odor_pastures_trash_disposal4")
	_allPasturesTrashDisposal.ResiduePasturesTrashDisposal1 = field.NewFloat64(tableName, "residue_pastures_trash_disposal1")
	_allPasturesTrashDisposal.ResiduePasturesTrashDisposal2 = field.NewFloat64(tableName, "residue_pastures_trash_disposal2")
	_allPasturesTrashDisposal.ResiduePasturesTrashDisposal3 = field.NewFloat64(tableName, "residue_pastures_trash_disposal3")
	_allPasturesTrashDisposal.ResiduePasturesTrashDisposal4 = field.NewFloat64(tableName, "residue_pastures_trash_disposal4")
	_allPasturesTrashDisposal.WaterPasturesTrashDisposal1 = field.NewFloat64(tableName, "water_pastures_trash_disposal1")
	_allPasturesTrashDisposal.WaterPasturesTrashDisposal2 = field.NewFloat64(tableName, "water_pastures_trash_disposal2")
	_allPasturesTrashDisposal.WaterPasturesTrashDisposal3 = field.NewFloat64(tableName, "water_pastures_trash_disposal3")
	_allPasturesTrashDisposal.WaterPasturesTrashDisposal4 = field.NewFloat64(tableName, "water_pastures_trash_disposal4")

	_allPasturesTrashDisposal.fillFieldMap()

	return _allPasturesTrashDisposal
}

type allPasturesTrashDisposal struct {
	allPasturesTrashDisposalDo allPasturesTrashDisposalDo

	ALL                           field.Asterisk
	ID                            field.Uint
	CreatedAt                     field.Time
	UpdatedAt                     field.Time
	DeletedAt                     field.Field
	TimeStamp                     field.Time
	OdorPasturesTrashDisposal1    field.Float64
	OdorPasturesTrashDisposal2    field.Float64
	OdorPasturesTrashDisposal3    field.Float64
	OdorPasturesTrashDisposal4    field.Float64
	ResiduePasturesTrashDisposal1 field.Float64
	ResiduePasturesTrashDisposal2 field.Float64
	ResiduePasturesTrashDisposal3 field.Float64
	ResiduePasturesTrashDisposal4 field.Float64
	WaterPasturesTrashDisposal1   field.Float64
	WaterPasturesTrashDisposal2   field.Float64
	WaterPasturesTrashDisposal3   field.Float64
	WaterPasturesTrashDisposal4   field.Float64

	fieldMap map[string]field.Expr
}

func (a allPasturesTrashDisposal) Table(newTableName string) *allPasturesTrashDisposal {
	a.allPasturesTrashDisposalDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a allPasturesTrashDisposal) As(alias string) *allPasturesTrashDisposal {
	a.allPasturesTrashDisposalDo.DO = *(a.allPasturesTrashDisposalDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *allPasturesTrashDisposal) updateTableName(table string) *allPasturesTrashDisposal {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.TimeStamp = field.NewTime(table, "time_stamp")
	a.OdorPasturesTrashDisposal1 = field.NewFloat64(table, "odor_pastures_trash_disposal1")
	a.OdorPasturesTrashDisposal2 = field.NewFloat64(table, "odor_pastures_trash_disposal2")
	a.OdorPasturesTrashDisposal3 = field.NewFloat64(table, "odor_pastures_trash_disposal3")
	a.OdorPasturesTrashDisposal4 = field.NewFloat64(table, "odor_pastures_trash_disposal4")
	a.ResiduePasturesTrashDisposal1 = field.NewFloat64(table, "residue_pastures_trash_disposal1")
	a.ResiduePasturesTrashDisposal2 = field.NewFloat64(table, "residue_pastures_trash_disposal2")
	a.ResiduePasturesTrashDisposal3 = field.NewFloat64(table, "residue_pastures_trash_disposal3")
	a.ResiduePasturesTrashDisposal4 = field.NewFloat64(table, "residue_pastures_trash_disposal4")
	a.WaterPasturesTrashDisposal1 = field.NewFloat64(table, "water_pastures_trash_disposal1")
	a.WaterPasturesTrashDisposal2 = field.NewFloat64(table, "water_pastures_trash_disposal2")
	a.WaterPasturesTrashDisposal3 = field.NewFloat64(table, "water_pastures_trash_disposal3")
	a.WaterPasturesTrashDisposal4 = field.NewFloat64(table, "water_pastures_trash_disposal4")

	a.fillFieldMap()

	return a
}

func (a *allPasturesTrashDisposal) WithContext(ctx context.Context) IAllPasturesTrashDisposalDo {
	return a.allPasturesTrashDisposalDo.WithContext(ctx)
}

func (a allPasturesTrashDisposal) TableName() string { return a.allPasturesTrashDisposalDo.TableName() }

func (a allPasturesTrashDisposal) Alias() string { return a.allPasturesTrashDisposalDo.Alias() }

func (a allPasturesTrashDisposal) Columns(cols ...field.Expr) gen.Columns {
	return a.allPasturesTrashDisposalDo.Columns(cols...)
}

func (a *allPasturesTrashDisposal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *allPasturesTrashDisposal) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 17)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["time_stamp"] = a.TimeStamp
	a.fieldMap["odor_pastures_trash_disposal1"] = a.OdorPasturesTrashDisposal1
	a.fieldMap["odor_pastures_trash_disposal2"] = a.OdorPasturesTrashDisposal2
	a.fieldMap["odor_pastures_trash_disposal3"] = a.OdorPasturesTrashDisposal3
	a.fieldMap["odor_pastures_trash_disposal4"] = a.OdorPasturesTrashDisposal4
	a.fieldMap["residue_pastures_trash_disposal1"] = a.ResiduePasturesTrashDisposal1
	a.fieldMap["residue_pastures_trash_disposal2"] = a.ResiduePasturesTrashDisposal2
	a.fieldMap["residue_pastures_trash_disposal3"] = a.ResiduePasturesTrashDisposal3
	a.fieldMap["residue_pastures_trash_disposal4"] = a.ResiduePasturesTrashDisposal4
	a.fieldMap["water_pastures_trash_disposal1"] = a.WaterPasturesTrashDisposal1
	a.fieldMap["water_pastures_trash_disposal2"] = a.WaterPasturesTrashDisposal2
	a.fieldMap["water_pastures_trash_disposal3"] = a.WaterPasturesTrashDisposal3
	a.fieldMap["water_pastures_trash_disposal4"] = a.WaterPasturesTrashDisposal4
}

func (a allPasturesTrashDisposal) clone(db *gorm.DB) allPasturesTrashDisposal {
	a.allPasturesTrashDisposalDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a allPasturesTrashDisposal) replaceDB(db *gorm.DB) allPasturesTrashDisposal {
	a.allPasturesTrashDisposalDo.ReplaceDB(db)
	return a
}

type allPasturesTrashDisposalDo struct{ gen.DO }

type IAllPasturesTrashDisposalDo interface {
	gen.SubQuery
	Debug() IAllPasturesTrashDisposalDo
	WithContext(ctx context.Context) IAllPasturesTrashDisposalDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAllPasturesTrashDisposalDo
	WriteDB() IAllPasturesTrashDisposalDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAllPasturesTrashDisposalDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAllPasturesTrashDisposalDo
	Not(conds ...gen.Condition) IAllPasturesTrashDisposalDo
	Or(conds ...gen.Condition) IAllPasturesTrashDisposalDo
	Select(conds ...field.Expr) IAllPasturesTrashDisposalDo
	Where(conds ...gen.Condition) IAllPasturesTrashDisposalDo
	Order(conds ...field.Expr) IAllPasturesTrashDisposalDo
	Distinct(cols ...field.Expr) IAllPasturesTrashDisposalDo
	Omit(cols ...field.Expr) IAllPasturesTrashDisposalDo
	Join(table schema.Tabler, on ...field.Expr) IAllPasturesTrashDisposalDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAllPasturesTrashDisposalDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAllPasturesTrashDisposalDo
	Group(cols ...field.Expr) IAllPasturesTrashDisposalDo
	Having(conds ...gen.Condition) IAllPasturesTrashDisposalDo
	Limit(limit int) IAllPasturesTrashDisposalDo
	Offset(offset int) IAllPasturesTrashDisposalDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAllPasturesTrashDisposalDo
	Unscoped() IAllPasturesTrashDisposalDo
	Create(values ...*pasture.AllPasturesTrashDisposal) error
	CreateInBatches(values []*pasture.AllPasturesTrashDisposal, batchSize int) error
	Save(values ...*pasture.AllPasturesTrashDisposal) error
	First() (*pasture.AllPasturesTrashDisposal, error)
	Take() (*pasture.AllPasturesTrashDisposal, error)
	Last() (*pasture.AllPasturesTrashDisposal, error)
	Find() ([]*pasture.AllPasturesTrashDisposal, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.AllPasturesTrashDisposal, err error)
	FindInBatches(result *[]*pasture.AllPasturesTrashDisposal, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.AllPasturesTrashDisposal) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAllPasturesTrashDisposalDo
	Assign(attrs ...field.AssignExpr) IAllPasturesTrashDisposalDo
	Joins(fields ...field.RelationField) IAllPasturesTrashDisposalDo
	Preload(fields ...field.RelationField) IAllPasturesTrashDisposalDo
	FirstOrInit() (*pasture.AllPasturesTrashDisposal, error)
	FirstOrCreate() (*pasture.AllPasturesTrashDisposal, error)
	FindByPage(offset int, limit int) (result []*pasture.AllPasturesTrashDisposal, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAllPasturesTrashDisposalDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a allPasturesTrashDisposalDo) Debug() IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Debug())
}

func (a allPasturesTrashDisposalDo) WithContext(ctx context.Context) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a allPasturesTrashDisposalDo) ReadDB() IAllPasturesTrashDisposalDo {
	return a.Clauses(dbresolver.Read)
}

func (a allPasturesTrashDisposalDo) WriteDB() IAllPasturesTrashDisposalDo {
	return a.Clauses(dbresolver.Write)
}

func (a allPasturesTrashDisposalDo) Session(config *gorm.Session) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Session(config))
}

func (a allPasturesTrashDisposalDo) Clauses(conds ...clause.Expression) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a allPasturesTrashDisposalDo) Returning(value interface{}, columns ...string) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a allPasturesTrashDisposalDo) Not(conds ...gen.Condition) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a allPasturesTrashDisposalDo) Or(conds ...gen.Condition) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a allPasturesTrashDisposalDo) Select(conds ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a allPasturesTrashDisposalDo) Where(conds ...gen.Condition) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a allPasturesTrashDisposalDo) Order(conds ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a allPasturesTrashDisposalDo) Distinct(cols ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a allPasturesTrashDisposalDo) Omit(cols ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a allPasturesTrashDisposalDo) Join(table schema.Tabler, on ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a allPasturesTrashDisposalDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a allPasturesTrashDisposalDo) RightJoin(table schema.Tabler, on ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a allPasturesTrashDisposalDo) Group(cols ...field.Expr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a allPasturesTrashDisposalDo) Having(conds ...gen.Condition) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a allPasturesTrashDisposalDo) Limit(limit int) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a allPasturesTrashDisposalDo) Offset(offset int) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a allPasturesTrashDisposalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a allPasturesTrashDisposalDo) Unscoped() IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Unscoped())
}

func (a allPasturesTrashDisposalDo) Create(values ...*pasture.AllPasturesTrashDisposal) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a allPasturesTrashDisposalDo) CreateInBatches(values []*pasture.AllPasturesTrashDisposal, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a allPasturesTrashDisposalDo) Save(values ...*pasture.AllPasturesTrashDisposal) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a allPasturesTrashDisposalDo) First() (*pasture.AllPasturesTrashDisposal, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.AllPasturesTrashDisposal), nil
	}
}

func (a allPasturesTrashDisposalDo) Take() (*pasture.AllPasturesTrashDisposal, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.AllPasturesTrashDisposal), nil
	}
}

func (a allPasturesTrashDisposalDo) Last() (*pasture.AllPasturesTrashDisposal, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.AllPasturesTrashDisposal), nil
	}
}

func (a allPasturesTrashDisposalDo) Find() ([]*pasture.AllPasturesTrashDisposal, error) {
	result, err := a.DO.Find()
	return result.([]*pasture.AllPasturesTrashDisposal), err
}

func (a allPasturesTrashDisposalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.AllPasturesTrashDisposal, err error) {
	buf := make([]*pasture.AllPasturesTrashDisposal, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a allPasturesTrashDisposalDo) FindInBatches(result *[]*pasture.AllPasturesTrashDisposal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a allPasturesTrashDisposalDo) Attrs(attrs ...field.AssignExpr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a allPasturesTrashDisposalDo) Assign(attrs ...field.AssignExpr) IAllPasturesTrashDisposalDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a allPasturesTrashDisposalDo) Joins(fields ...field.RelationField) IAllPasturesTrashDisposalDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a allPasturesTrashDisposalDo) Preload(fields ...field.RelationField) IAllPasturesTrashDisposalDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a allPasturesTrashDisposalDo) FirstOrInit() (*pasture.AllPasturesTrashDisposal, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.AllPasturesTrashDisposal), nil
	}
}

func (a allPasturesTrashDisposalDo) FirstOrCreate() (*pasture.AllPasturesTrashDisposal, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.AllPasturesTrashDisposal), nil
	}
}

func (a allPasturesTrashDisposalDo) FindByPage(offset int, limit int) (result []*pasture.AllPasturesTrashDisposal, count int64, err error) {
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

func (a allPasturesTrashDisposalDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a allPasturesTrashDisposalDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a allPasturesTrashDisposalDo) Delete(models ...*pasture.AllPasturesTrashDisposal) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *allPasturesTrashDisposalDo) withDO(do gen.Dao) *allPasturesTrashDisposalDo {
	a.DO = *do.(*gen.DO)
	return a
}