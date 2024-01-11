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

func newWaitingSlaughterCircle(db *gorm.DB, opts ...gen.DOOption) waitingSlaughterCircle {
	_waitingSlaughterCircle := waitingSlaughterCircle{}

	_waitingSlaughterCircle.waitingSlaughterCircleDo.UseDB(db, opts...)
	_waitingSlaughterCircle.waitingSlaughterCircleDo.UseModel(&premortem.WaitingSlaughterCircle{})

	tableName := _waitingSlaughterCircle.waitingSlaughterCircleDo.TableName()
	_waitingSlaughterCircle.ALL = field.NewAsterisk(tableName)
	_waitingSlaughterCircle.ID = field.NewUint(tableName, "id")
	_waitingSlaughterCircle.CreatedAt = field.NewTime(tableName, "created_at")
	_waitingSlaughterCircle.UpdatedAt = field.NewTime(tableName, "updated_at")
	_waitingSlaughterCircle.DeletedAt = field.NewField(tableName, "deleted_at")
	_waitingSlaughterCircle.WaitingSlaughterCircle2 = field.NewFloat32(tableName, "waiting_slaughter_circle2")
	_waitingSlaughterCircle.WaitingSlaughterCircle3 = field.NewFloat32(tableName, "waiting_slaughter_circle3")
	_waitingSlaughterCircle.WaitingSlaughterCircle4 = field.NewFloat32(tableName, "waiting_slaughter_circle4")
	_waitingSlaughterCircle.WaitingSlaughterCircle1 = waitingSlaughterCircleHasOneWaitingSlaughterCircle1{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("WaitingSlaughterCircle1", "premortem.WaitingSlaughterCircleGerms"),
	}

	_waitingSlaughterCircle.fillFieldMap()

	return _waitingSlaughterCircle
}

type waitingSlaughterCircle struct {
	waitingSlaughterCircleDo waitingSlaughterCircleDo

	ALL                     field.Asterisk
	ID                      field.Uint
	CreatedAt               field.Time
	UpdatedAt               field.Time
	DeletedAt               field.Field
	WaitingSlaughterCircle2 field.Float32
	WaitingSlaughterCircle3 field.Float32
	WaitingSlaughterCircle4 field.Float32
	WaitingSlaughterCircle1 waitingSlaughterCircleHasOneWaitingSlaughterCircle1

	fieldMap map[string]field.Expr
}

func (w waitingSlaughterCircle) Table(newTableName string) *waitingSlaughterCircle {
	w.waitingSlaughterCircleDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w waitingSlaughterCircle) As(alias string) *waitingSlaughterCircle {
	w.waitingSlaughterCircleDo.DO = *(w.waitingSlaughterCircleDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *waitingSlaughterCircle) updateTableName(table string) *waitingSlaughterCircle {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.WaitingSlaughterCircle2 = field.NewFloat32(table, "waiting_slaughter_circle2")
	w.WaitingSlaughterCircle3 = field.NewFloat32(table, "waiting_slaughter_circle3")
	w.WaitingSlaughterCircle4 = field.NewFloat32(table, "waiting_slaughter_circle4")

	w.fillFieldMap()

	return w
}

func (w *waitingSlaughterCircle) WithContext(ctx context.Context) IWaitingSlaughterCircleDo {
	return w.waitingSlaughterCircleDo.WithContext(ctx)
}

func (w waitingSlaughterCircle) TableName() string { return w.waitingSlaughterCircleDo.TableName() }

func (w waitingSlaughterCircle) Alias() string { return w.waitingSlaughterCircleDo.Alias() }

func (w waitingSlaughterCircle) Columns(cols ...field.Expr) gen.Columns {
	return w.waitingSlaughterCircleDo.Columns(cols...)
}

func (w *waitingSlaughterCircle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *waitingSlaughterCircle) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["waiting_slaughter_circle2"] = w.WaitingSlaughterCircle2
	w.fieldMap["waiting_slaughter_circle3"] = w.WaitingSlaughterCircle3
	w.fieldMap["waiting_slaughter_circle4"] = w.WaitingSlaughterCircle4

}

func (w waitingSlaughterCircle) clone(db *gorm.DB) waitingSlaughterCircle {
	w.waitingSlaughterCircleDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w waitingSlaughterCircle) replaceDB(db *gorm.DB) waitingSlaughterCircle {
	w.waitingSlaughterCircleDo.ReplaceDB(db)
	return w
}

type waitingSlaughterCircleHasOneWaitingSlaughterCircle1 struct {
	db *gorm.DB

	field.RelationField
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1) Where(conds ...field.Expr) *waitingSlaughterCircleHasOneWaitingSlaughterCircle1 {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1) WithContext(ctx context.Context) *waitingSlaughterCircleHasOneWaitingSlaughterCircle1 {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1) Session(session *gorm.Session) *waitingSlaughterCircleHasOneWaitingSlaughterCircle1 {
	a.db = a.db.Session(session)
	return &a
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1) Model(m *premortem.WaitingSlaughterCircle) *waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx {
	return &waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx{a.db.Model(m).Association(a.Name())}
}

type waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx struct{ tx *gorm.Association }

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx) Find() (result *premortem.WaitingSlaughterCircleGerms, err error) {
	return result, a.tx.Find(&result)
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx) Append(values ...*premortem.WaitingSlaughterCircleGerms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx) Replace(values ...*premortem.WaitingSlaughterCircleGerms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx) Delete(values ...*premortem.WaitingSlaughterCircleGerms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx) Clear() error {
	return a.tx.Clear()
}

func (a waitingSlaughterCircleHasOneWaitingSlaughterCircle1Tx) Count() int64 {
	return a.tx.Count()
}

type waitingSlaughterCircleDo struct{ gen.DO }

type IWaitingSlaughterCircleDo interface {
	gen.SubQuery
	Debug() IWaitingSlaughterCircleDo
	WithContext(ctx context.Context) IWaitingSlaughterCircleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWaitingSlaughterCircleDo
	WriteDB() IWaitingSlaughterCircleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWaitingSlaughterCircleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWaitingSlaughterCircleDo
	Not(conds ...gen.Condition) IWaitingSlaughterCircleDo
	Or(conds ...gen.Condition) IWaitingSlaughterCircleDo
	Select(conds ...field.Expr) IWaitingSlaughterCircleDo
	Where(conds ...gen.Condition) IWaitingSlaughterCircleDo
	Order(conds ...field.Expr) IWaitingSlaughterCircleDo
	Distinct(cols ...field.Expr) IWaitingSlaughterCircleDo
	Omit(cols ...field.Expr) IWaitingSlaughterCircleDo
	Join(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleDo
	Group(cols ...field.Expr) IWaitingSlaughterCircleDo
	Having(conds ...gen.Condition) IWaitingSlaughterCircleDo
	Limit(limit int) IWaitingSlaughterCircleDo
	Offset(offset int) IWaitingSlaughterCircleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWaitingSlaughterCircleDo
	Unscoped() IWaitingSlaughterCircleDo
	Create(values ...*premortem.WaitingSlaughterCircle) error
	CreateInBatches(values []*premortem.WaitingSlaughterCircle, batchSize int) error
	Save(values ...*premortem.WaitingSlaughterCircle) error
	First() (*premortem.WaitingSlaughterCircle, error)
	Take() (*premortem.WaitingSlaughterCircle, error)
	Last() (*premortem.WaitingSlaughterCircle, error)
	Find() ([]*premortem.WaitingSlaughterCircle, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*premortem.WaitingSlaughterCircle, err error)
	FindInBatches(result *[]*premortem.WaitingSlaughterCircle, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*premortem.WaitingSlaughterCircle) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWaitingSlaughterCircleDo
	Assign(attrs ...field.AssignExpr) IWaitingSlaughterCircleDo
	Joins(fields ...field.RelationField) IWaitingSlaughterCircleDo
	Preload(fields ...field.RelationField) IWaitingSlaughterCircleDo
	FirstOrInit() (*premortem.WaitingSlaughterCircle, error)
	FirstOrCreate() (*premortem.WaitingSlaughterCircle, error)
	FindByPage(offset int, limit int) (result []*premortem.WaitingSlaughterCircle, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWaitingSlaughterCircleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w waitingSlaughterCircleDo) Debug() IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Debug())
}

func (w waitingSlaughterCircleDo) WithContext(ctx context.Context) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w waitingSlaughterCircleDo) ReadDB() IWaitingSlaughterCircleDo {
	return w.Clauses(dbresolver.Read)
}

func (w waitingSlaughterCircleDo) WriteDB() IWaitingSlaughterCircleDo {
	return w.Clauses(dbresolver.Write)
}

func (w waitingSlaughterCircleDo) Session(config *gorm.Session) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Session(config))
}

func (w waitingSlaughterCircleDo) Clauses(conds ...clause.Expression) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w waitingSlaughterCircleDo) Returning(value interface{}, columns ...string) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w waitingSlaughterCircleDo) Not(conds ...gen.Condition) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w waitingSlaughterCircleDo) Or(conds ...gen.Condition) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w waitingSlaughterCircleDo) Select(conds ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w waitingSlaughterCircleDo) Where(conds ...gen.Condition) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w waitingSlaughterCircleDo) Order(conds ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w waitingSlaughterCircleDo) Distinct(cols ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w waitingSlaughterCircleDo) Omit(cols ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w waitingSlaughterCircleDo) Join(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w waitingSlaughterCircleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w waitingSlaughterCircleDo) RightJoin(table schema.Tabler, on ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w waitingSlaughterCircleDo) Group(cols ...field.Expr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w waitingSlaughterCircleDo) Having(conds ...gen.Condition) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w waitingSlaughterCircleDo) Limit(limit int) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w waitingSlaughterCircleDo) Offset(offset int) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w waitingSlaughterCircleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w waitingSlaughterCircleDo) Unscoped() IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Unscoped())
}

func (w waitingSlaughterCircleDo) Create(values ...*premortem.WaitingSlaughterCircle) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w waitingSlaughterCircleDo) CreateInBatches(values []*premortem.WaitingSlaughterCircle, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w waitingSlaughterCircleDo) Save(values ...*premortem.WaitingSlaughterCircle) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w waitingSlaughterCircleDo) First() (*premortem.WaitingSlaughterCircle, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircle), nil
	}
}

func (w waitingSlaughterCircleDo) Take() (*premortem.WaitingSlaughterCircle, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircle), nil
	}
}

func (w waitingSlaughterCircleDo) Last() (*premortem.WaitingSlaughterCircle, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircle), nil
	}
}

func (w waitingSlaughterCircleDo) Find() ([]*premortem.WaitingSlaughterCircle, error) {
	result, err := w.DO.Find()
	return result.([]*premortem.WaitingSlaughterCircle), err
}

func (w waitingSlaughterCircleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*premortem.WaitingSlaughterCircle, err error) {
	buf := make([]*premortem.WaitingSlaughterCircle, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w waitingSlaughterCircleDo) FindInBatches(result *[]*premortem.WaitingSlaughterCircle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w waitingSlaughterCircleDo) Attrs(attrs ...field.AssignExpr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w waitingSlaughterCircleDo) Assign(attrs ...field.AssignExpr) IWaitingSlaughterCircleDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w waitingSlaughterCircleDo) Joins(fields ...field.RelationField) IWaitingSlaughterCircleDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w waitingSlaughterCircleDo) Preload(fields ...field.RelationField) IWaitingSlaughterCircleDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w waitingSlaughterCircleDo) FirstOrInit() (*premortem.WaitingSlaughterCircle, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircle), nil
	}
}

func (w waitingSlaughterCircleDo) FirstOrCreate() (*premortem.WaitingSlaughterCircle, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.WaitingSlaughterCircle), nil
	}
}

func (w waitingSlaughterCircleDo) FindByPage(offset int, limit int) (result []*premortem.WaitingSlaughterCircle, count int64, err error) {
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

func (w waitingSlaughterCircleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w waitingSlaughterCircleDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w waitingSlaughterCircleDo) Delete(models ...*premortem.WaitingSlaughterCircle) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *waitingSlaughterCircleDo) withDO(do gen.Dao) *waitingSlaughterCircleDo {
	w.DO = *do.(*gen.DO)
	return w
}
