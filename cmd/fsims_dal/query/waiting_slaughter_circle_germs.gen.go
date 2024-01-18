// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/premortem"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newWaitingSlaughterCircleGerms(db *gorm.DB, opts ...gen.DOOption) waitingSlaughterCircleGerms {
	_waitingSlaughterCircleGerms := waitingSlaughterCircleGerms{}

	_waitingSlaughterCircleGerms.waitingSlaughterCircleGermsDo.UseDB(db, opts...)
	_waitingSlaughterCircleGerms.waitingSlaughterCircleGermsDo.UseModel(&premortem.WaitingSlaughterCircleGerms{})

	tableName := _waitingSlaughterCircleGerms.waitingSlaughterCircleGermsDo.TableName()
	_waitingSlaughterCircleGerms.ALL = field.NewAsterisk(tableName)
	_waitingSlaughterCircleGerms.ID = field.NewUint(tableName, "id")
	_waitingSlaughterCircleGerms.CreatedAt = field.NewTime(tableName, "created_at")
	_waitingSlaughterCircleGerms.UpdatedAt = field.NewTime(tableName, "updated_at")
	_waitingSlaughterCircleGerms.DeletedAt = field.NewField(tableName, "deleted_at")
	_waitingSlaughterCircleGerms.WaitingSlaughterCircleID = field.NewUint(tableName, "waiting_slaughter_circle_id")
	_waitingSlaughterCircleGerms.WaitingSlaughterCircleGerms1 = field.NewFloat64(tableName, "waiting_slaughter_circle_germs1")
	_waitingSlaughterCircleGerms.WaitingSlaughterCircleGerms2 = field.NewFloat64(tableName, "waiting_slaughter_circle_germs2")
	_waitingSlaughterCircleGerms.WaitingSlaughterCircleGerms3 = field.NewFloat64(tableName, "waiting_slaughter_circle_germs3")

	_waitingSlaughterCircleGerms.fillFieldMap()

	return _waitingSlaughterCircleGerms
}

type waitingSlaughterCircleGerms struct {
	waitingSlaughterCircleGermsDo waitingSlaughterCircleGermsDo

	ALL                          field.Asterisk
	ID                           field.Uint
	CreatedAt                    field.Time
	UpdatedAt                    field.Time
	DeletedAt                    field.Field
	WaitingSlaughterCircleID     field.Uint
	WaitingSlaughterCircleGerms1 field.Float64
	WaitingSlaughterCircleGerms2 field.Float64
	WaitingSlaughterCircleGerms3 field.Float64

	fieldMap map[string]field.Expr
}

func (w waitingSlaughterCircleGerms) Table(newTableName string) *waitingSlaughterCircleGerms {
	w.waitingSlaughterCircleGermsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w waitingSlaughterCircleGerms) As(alias string) *waitingSlaughterCircleGerms {
	w.waitingSlaughterCircleGermsDo.DO = *(w.waitingSlaughterCircleGermsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *waitingSlaughterCircleGerms) updateTableName(table string) *waitingSlaughterCircleGerms {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.WaitingSlaughterCircleID = field.NewUint(table, "waiting_slaughter_circle_id")
	w.WaitingSlaughterCircleGerms1 = field.NewFloat64(table, "waiting_slaughter_circle_germs1")
	w.WaitingSlaughterCircleGerms2 = field.NewFloat64(table, "waiting_slaughter_circle_germs2")
	w.WaitingSlaughterCircleGerms3 = field.NewFloat64(table, "waiting_slaughter_circle_germs3")

	w.fillFieldMap()

	return w
}

func (w *waitingSlaughterCircleGerms) WithContext(ctx context.Context) IWaitingSlaughterCircleGermsDo {
	return w.waitingSlaughterCircleGermsDo.WithContext(ctx)
}

func (w waitingSlaughterCircleGerms) TableName() string {
	return w.waitingSlaughterCircleGermsDo.TableName()
}

func (w waitingSlaughterCircleGerms) Alias() string { return w.waitingSlaughterCircleGermsDo.Alias() }

func (w waitingSlaughterCircleGerms) Columns(cols ...field.Expr) gen.Columns {
	return w.waitingSlaughterCircleGermsDo.Columns(cols...)
}

func (w *waitingSlaughterCircleGerms) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *waitingSlaughterCircleGerms) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["waiting_slaughter_circle_id"] = w.WaitingSlaughterCircleID
	w.fieldMap["waiting_slaughter_circle_germs1"] = w.WaitingSlaughterCircleGerms1
	w.fieldMap["waiting_slaughter_circle_germs2"] = w.WaitingSlaughterCircleGerms2
	w.fieldMap["waiting_slaughter_circle_germs3"] = w.WaitingSlaughterCircleGerms3
}

func (w waitingSlaughterCircleGerms) clone(db *gorm.DB) waitingSlaughterCircleGerms {
	w.waitingSlaughterCircleGermsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w waitingSlaughterCircleGerms) replaceDB(db *gorm.DB) waitingSlaughterCircleGerms {
	w.waitingSlaughterCircleGermsDo.ReplaceDB(db)
	return w
}

type waitingSlaughterCircleGermsDo struct{ gen.DO }

type IWaitingSlaughterCircleGermsDo interface {
	gen.SubQuery
	Debug() IWaitingSlaughterCircleGermsDo
	WithContext(ctx context.Context) IWaitingSlaughterCircleGermsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWaitingSlaughterCircleGermsDo
	WriteDB() IWaitingSlaughterCircleGermsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWaitingSlaughterCircleGermsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWaitingSlaughterCircleGermsDo
	Not(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo
	Or(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo
	Select(conds ...field.Expr) IWaitingSlaughterCircleGermsDo
	Where(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo
	Order(conds ...field.Expr) IWaitingSlaughterCircleGermsDo
	Distinct(cols ...field.Expr) IWaitingSlaughterCircleGermsDo
	Omit(cols ...field.Expr) IWaitingSlaughterCircleGermsDo
	Join(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleGermsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleGermsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleGermsDo
	Group(cols ...field.Expr) IWaitingSlaughterCircleGermsDo
	Having(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo
	Limit(limit int) IWaitingSlaughterCircleGermsDo
	Offset(offset int) IWaitingSlaughterCircleGermsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWaitingSlaughterCircleGermsDo
	Unscoped() IWaitingSlaughterCircleGermsDo
	Create(values ...*premortem.WaitingSlaughterCircleGerms) error
	CreateInBatches(values []*premortem.WaitingSlaughterCircleGerms, batchSize int) error
	Save(values ...*premortem.WaitingSlaughterCircleGerms) error
	First() (*premortem.WaitingSlaughterCircleGerms, error)
	Take() (*premortem.WaitingSlaughterCircleGerms, error)
	Last() (*premortem.WaitingSlaughterCircleGerms, error)
	Find() ([]*premortem.WaitingSlaughterCircleGerms, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*premortem.WaitingSlaughterCircleGerms, err error)
	FindInBatches(result *[]*premortem.WaitingSlaughterCircleGerms, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*premortem.WaitingSlaughterCircleGerms) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWaitingSlaughterCircleGermsDo
	Assign(attrs ...field.AssignExpr) IWaitingSlaughterCircleGermsDo
	Joins(fields ...field.RelationField) IWaitingSlaughterCircleGermsDo
	Preload(fields ...field.RelationField) IWaitingSlaughterCircleGermsDo
	FirstOrInit() (*premortem.WaitingSlaughterCircleGerms, error)
	FirstOrCreate() (*premortem.WaitingSlaughterCircleGerms, error)
	FindByPage(offset int, limit int) (result []*premortem.WaitingSlaughterCircleGerms, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWaitingSlaughterCircleGermsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w waitingSlaughterCircleGermsDo) Debug() IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Debug())
}

func (w waitingSlaughterCircleGermsDo) WithContext(ctx context.Context) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w waitingSlaughterCircleGermsDo) ReadDB() IWaitingSlaughterCircleGermsDo {
	return w.Clauses(dbresolver.Read)
}

func (w waitingSlaughterCircleGermsDo) WriteDB() IWaitingSlaughterCircleGermsDo {
	return w.Clauses(dbresolver.Write)
}

func (w waitingSlaughterCircleGermsDo) Session(config *gorm.Session) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Session(config))
}

func (w waitingSlaughterCircleGermsDo) Clauses(conds ...clause.Expression) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w waitingSlaughterCircleGermsDo) Returning(value interface{}, columns ...string) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w waitingSlaughterCircleGermsDo) Not(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w waitingSlaughterCircleGermsDo) Or(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w waitingSlaughterCircleGermsDo) Select(conds ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w waitingSlaughterCircleGermsDo) Where(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w waitingSlaughterCircleGermsDo) Order(conds ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w waitingSlaughterCircleGermsDo) Distinct(cols ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w waitingSlaughterCircleGermsDo) Omit(cols ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w waitingSlaughterCircleGermsDo) Join(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w waitingSlaughterCircleGermsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w waitingSlaughterCircleGermsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w waitingSlaughterCircleGermsDo) Group(cols ...field.Expr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w waitingSlaughterCircleGermsDo) Having(conds ...gen.Condition) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w waitingSlaughterCircleGermsDo) Limit(limit int) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w waitingSlaughterCircleGermsDo) Offset(offset int) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w waitingSlaughterCircleGermsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w waitingSlaughterCircleGermsDo) Unscoped() IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w waitingSlaughterCircleGermsDo) Create(values ...*premortem.WaitingSlaughterCircleGerms) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w waitingSlaughterCircleGermsDo) CreateInBatches(values []*premortem.WaitingSlaughterCircleGerms, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w waitingSlaughterCircleGermsDo) Save(values ...*premortem.WaitingSlaughterCircleGerms) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w waitingSlaughterCircleGermsDo) First() (*premortem.WaitingSlaughterCircleGerms, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircleGerms), nil
	}
}

func (w waitingSlaughterCircleGermsDo) Take() (*premortem.WaitingSlaughterCircleGerms, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircleGerms), nil
	}
}

func (w waitingSlaughterCircleGermsDo) Last() (*premortem.WaitingSlaughterCircleGerms, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircleGerms), nil
	}
}

func (w waitingSlaughterCircleGermsDo) Find() ([]*premortem.WaitingSlaughterCircleGerms, error) {
	result, err := w.DO.Find()
	return result.([]*premortem.WaitingSlaughterCircleGerms), err
}

func (w waitingSlaughterCircleGermsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*premortem.WaitingSlaughterCircleGerms, err error) {
	buf := make([]*premortem.WaitingSlaughterCircleGerms, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w waitingSlaughterCircleGermsDo) FindInBatches(result *[]*premortem.WaitingSlaughterCircleGerms, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w waitingSlaughterCircleGermsDo) Attrs(attrs ...field.AssignExpr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w waitingSlaughterCircleGermsDo) Assign(attrs ...field.AssignExpr) IWaitingSlaughterCircleGermsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w waitingSlaughterCircleGermsDo) Joins(fields ...field.RelationField) IWaitingSlaughterCircleGermsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w waitingSlaughterCircleGermsDo) Preload(fields ...field.RelationField) IWaitingSlaughterCircleGermsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w waitingSlaughterCircleGermsDo) FirstOrInit() (*premortem.WaitingSlaughterCircleGerms, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircleGerms), nil
	}
}

func (w waitingSlaughterCircleGermsDo) FirstOrCreate() (*premortem.WaitingSlaughterCircleGerms, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircleGerms), nil
	}
}

func (w waitingSlaughterCircleGermsDo) FindByPage(offset int, limit int) (result []*premortem.WaitingSlaughterCircleGerms, count int64, err error) {
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

func (w waitingSlaughterCircleGermsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w waitingSlaughterCircleGermsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w waitingSlaughterCircleGermsDo) Delete(models ...*premortem.WaitingSlaughterCircleGerms) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *waitingSlaughterCircleGermsDo) withDO(do gen.Dao) *waitingSlaughterCircleGermsDo {
	w.DO = *do.(*gen.DO)
	return w
}
