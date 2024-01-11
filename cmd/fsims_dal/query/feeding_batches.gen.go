// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/product"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newFeedingBatch(db *gorm.DB, opts ...gen.DOOption) feedingBatch {
	_feedingBatch := feedingBatch{}

	_feedingBatch.feedingBatchDo.UseDB(db, opts...)
	_feedingBatch.feedingBatchDo.UseModel(&pasture.FeedingBatch{})

	tableName := _feedingBatch.feedingBatchDo.TableName()
	_feedingBatch.ALL = field.NewAsterisk(tableName)
	_feedingBatch.ID = field.NewUint(tableName, "id")
	_feedingBatch.CreatedAt = field.NewTime(tableName, "created_at")
	_feedingBatch.UpdatedAt = field.NewTime(tableName, "updated_at")
	_feedingBatch.DeletedAt = field.NewField(tableName, "deleted_at")
	_feedingBatch.BatchNumber = field.NewString(tableName, "batch_number")
	_feedingBatch.HouseNumber = field.NewString(tableName, "house_number")
	_feedingBatch.State = field.NewInt(tableName, "state")
	_feedingBatch.PID = field.NewString(tableName, "p_id")
	_feedingBatch.Worker = field.NewString(tableName, "worker")
	_feedingBatch.Procedure = feedingBatchHasOneProcedure{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Procedure", "models.Procedure"),
	}

	_feedingBatch.Cows = feedingBatchHasManyCows{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Cows", "product.Cow"),
	}

	_feedingBatch.fillFieldMap()

	return _feedingBatch
}

type feedingBatch struct {
	feedingBatchDo feedingBatchDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	BatchNumber field.String
	HouseNumber field.String
	State       field.Int
	PID         field.String
	Worker      field.String
	Procedure   feedingBatchHasOneProcedure

	Cows feedingBatchHasManyCows

	fieldMap map[string]field.Expr
}

func (f feedingBatch) Table(newTableName string) *feedingBatch {
	f.feedingBatchDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f feedingBatch) As(alias string) *feedingBatch {
	f.feedingBatchDo.DO = *(f.feedingBatchDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *feedingBatch) updateTableName(table string) *feedingBatch {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewUint(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.BatchNumber = field.NewString(table, "batch_number")
	f.HouseNumber = field.NewString(table, "house_number")
	f.State = field.NewInt(table, "state")
	f.PID = field.NewString(table, "p_id")
	f.Worker = field.NewString(table, "worker")

	f.fillFieldMap()

	return f
}

func (f *feedingBatch) WithContext(ctx context.Context) IFeedingBatchDo {
	return f.feedingBatchDo.WithContext(ctx)
}

func (f feedingBatch) TableName() string { return f.feedingBatchDo.TableName() }

func (f feedingBatch) Alias() string { return f.feedingBatchDo.Alias() }

func (f feedingBatch) Columns(cols ...field.Expr) gen.Columns {
	return f.feedingBatchDo.Columns(cols...)
}

func (f *feedingBatch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *feedingBatch) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 11)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["batch_number"] = f.BatchNumber
	f.fieldMap["house_number"] = f.HouseNumber
	f.fieldMap["state"] = f.State
	f.fieldMap["p_id"] = f.PID
	f.fieldMap["worker"] = f.Worker

}

func (f feedingBatch) clone(db *gorm.DB) feedingBatch {
	f.feedingBatchDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f feedingBatch) replaceDB(db *gorm.DB) feedingBatch {
	f.feedingBatchDo.ReplaceDB(db)
	return f
}

type feedingBatchHasOneProcedure struct {
	db *gorm.DB

	field.RelationField
}

func (a feedingBatchHasOneProcedure) Where(conds ...field.Expr) *feedingBatchHasOneProcedure {
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

func (a feedingBatchHasOneProcedure) WithContext(ctx context.Context) *feedingBatchHasOneProcedure {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a feedingBatchHasOneProcedure) Session(session *gorm.Session) *feedingBatchHasOneProcedure {
	a.db = a.db.Session(session)
	return &a
}

func (a feedingBatchHasOneProcedure) Model(m *pasture.FeedingBatch) *feedingBatchHasOneProcedureTx {
	return &feedingBatchHasOneProcedureTx{a.db.Model(m).Association(a.Name())}
}

type feedingBatchHasOneProcedureTx struct{ tx *gorm.Association }

func (a feedingBatchHasOneProcedureTx) Find() (result *models.Procedure, err error) {
	return result, a.tx.Find(&result)
}

func (a feedingBatchHasOneProcedureTx) Append(values ...*models.Procedure) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a feedingBatchHasOneProcedureTx) Replace(values ...*models.Procedure) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a feedingBatchHasOneProcedureTx) Delete(values ...*models.Procedure) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a feedingBatchHasOneProcedureTx) Clear() error {
	return a.tx.Clear()
}

func (a feedingBatchHasOneProcedureTx) Count() int64 {
	return a.tx.Count()
}

type feedingBatchHasManyCows struct {
	db *gorm.DB

	field.RelationField
}

func (a feedingBatchHasManyCows) Where(conds ...field.Expr) *feedingBatchHasManyCows {
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

func (a feedingBatchHasManyCows) WithContext(ctx context.Context) *feedingBatchHasManyCows {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a feedingBatchHasManyCows) Session(session *gorm.Session) *feedingBatchHasManyCows {
	a.db = a.db.Session(session)
	return &a
}

func (a feedingBatchHasManyCows) Model(m *pasture.FeedingBatch) *feedingBatchHasManyCowsTx {
	return &feedingBatchHasManyCowsTx{a.db.Model(m).Association(a.Name())}
}

type feedingBatchHasManyCowsTx struct{ tx *gorm.Association }

func (a feedingBatchHasManyCowsTx) Find() (result []*product.Cow, err error) {
	return result, a.tx.Find(&result)
}

func (a feedingBatchHasManyCowsTx) Append(values ...*product.Cow) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a feedingBatchHasManyCowsTx) Replace(values ...*product.Cow) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a feedingBatchHasManyCowsTx) Delete(values ...*product.Cow) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a feedingBatchHasManyCowsTx) Clear() error {
	return a.tx.Clear()
}

func (a feedingBatchHasManyCowsTx) Count() int64 {
	return a.tx.Count()
}

type feedingBatchDo struct{ gen.DO }

type IFeedingBatchDo interface {
	gen.SubQuery
	Debug() IFeedingBatchDo
	WithContext(ctx context.Context) IFeedingBatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFeedingBatchDo
	WriteDB() IFeedingBatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFeedingBatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFeedingBatchDo
	Not(conds ...gen.Condition) IFeedingBatchDo
	Or(conds ...gen.Condition) IFeedingBatchDo
	Select(conds ...field.Expr) IFeedingBatchDo
	Where(conds ...gen.Condition) IFeedingBatchDo
	Order(conds ...field.Expr) IFeedingBatchDo
	Distinct(cols ...field.Expr) IFeedingBatchDo
	Omit(cols ...field.Expr) IFeedingBatchDo
	Join(table schema.Tabler, on ...field.Expr) IFeedingBatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFeedingBatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFeedingBatchDo
	Group(cols ...field.Expr) IFeedingBatchDo
	Having(conds ...gen.Condition) IFeedingBatchDo
	Limit(limit int) IFeedingBatchDo
	Offset(offset int) IFeedingBatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFeedingBatchDo
	Unscoped() IFeedingBatchDo
	Create(values ...*pasture.FeedingBatch) error
	CreateInBatches(values []*pasture.FeedingBatch, batchSize int) error
	Save(values ...*pasture.FeedingBatch) error
	First() (*pasture.FeedingBatch, error)
	Take() (*pasture.FeedingBatch, error)
	Last() (*pasture.FeedingBatch, error)
	Find() ([]*pasture.FeedingBatch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.FeedingBatch, err error)
	FindInBatches(result *[]*pasture.FeedingBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.FeedingBatch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFeedingBatchDo
	Assign(attrs ...field.AssignExpr) IFeedingBatchDo
	Joins(fields ...field.RelationField) IFeedingBatchDo
	Preload(fields ...field.RelationField) IFeedingBatchDo
	FirstOrInit() (*pasture.FeedingBatch, error)
	FirstOrCreate() (*pasture.FeedingBatch, error)
	FindByPage(offset int, limit int) (result []*pasture.FeedingBatch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFeedingBatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f feedingBatchDo) Debug() IFeedingBatchDo {
	return f.withDO(f.DO.Debug())
}

func (f feedingBatchDo) WithContext(ctx context.Context) IFeedingBatchDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f feedingBatchDo) ReadDB() IFeedingBatchDo {
	return f.Clauses(dbresolver.Read)
}

func (f feedingBatchDo) WriteDB() IFeedingBatchDo {
	return f.Clauses(dbresolver.Write)
}

func (f feedingBatchDo) Session(config *gorm.Session) IFeedingBatchDo {
	return f.withDO(f.DO.Session(config))
}

func (f feedingBatchDo) Clauses(conds ...clause.Expression) IFeedingBatchDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f feedingBatchDo) Returning(value interface{}, columns ...string) IFeedingBatchDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f feedingBatchDo) Not(conds ...gen.Condition) IFeedingBatchDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f feedingBatchDo) Or(conds ...gen.Condition) IFeedingBatchDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f feedingBatchDo) Select(conds ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f feedingBatchDo) Where(conds ...gen.Condition) IFeedingBatchDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f feedingBatchDo) Order(conds ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f feedingBatchDo) Distinct(cols ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f feedingBatchDo) Omit(cols ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f feedingBatchDo) Join(table schema.Tabler, on ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f feedingBatchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f feedingBatchDo) RightJoin(table schema.Tabler, on ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f feedingBatchDo) Group(cols ...field.Expr) IFeedingBatchDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f feedingBatchDo) Having(conds ...gen.Condition) IFeedingBatchDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f feedingBatchDo) Limit(limit int) IFeedingBatchDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f feedingBatchDo) Offset(offset int) IFeedingBatchDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f feedingBatchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFeedingBatchDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f feedingBatchDo) Unscoped() IFeedingBatchDo {
	return f.withDO(f.DO.Unscoped())
}

func (f feedingBatchDo) Create(values ...*pasture.FeedingBatch) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f feedingBatchDo) CreateInBatches(values []*pasture.FeedingBatch, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f feedingBatchDo) Save(values ...*pasture.FeedingBatch) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f feedingBatchDo) First() (*pasture.FeedingBatch, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.FeedingBatch), nil
	}
}

func (f feedingBatchDo) Take() (*pasture.FeedingBatch, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.FeedingBatch), nil
	}
}

func (f feedingBatchDo) Last() (*pasture.FeedingBatch, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.FeedingBatch), nil
	}
}

func (f feedingBatchDo) Find() ([]*pasture.FeedingBatch, error) {
	result, err := f.DO.Find()
	return result.([]*pasture.FeedingBatch), err
}

func (f feedingBatchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.FeedingBatch, err error) {
	buf := make([]*pasture.FeedingBatch, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f feedingBatchDo) FindInBatches(result *[]*pasture.FeedingBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f feedingBatchDo) Attrs(attrs ...field.AssignExpr) IFeedingBatchDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f feedingBatchDo) Assign(attrs ...field.AssignExpr) IFeedingBatchDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f feedingBatchDo) Joins(fields ...field.RelationField) IFeedingBatchDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f feedingBatchDo) Preload(fields ...field.RelationField) IFeedingBatchDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f feedingBatchDo) FirstOrInit() (*pasture.FeedingBatch, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.FeedingBatch), nil
	}
}

func (f feedingBatchDo) FirstOrCreate() (*pasture.FeedingBatch, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.FeedingBatch), nil
	}
}

func (f feedingBatchDo) FindByPage(offset int, limit int) (result []*pasture.FeedingBatch, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f feedingBatchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f feedingBatchDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f feedingBatchDo) Delete(models ...*pasture.FeedingBatch) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *feedingBatchDo) withDO(do gen.Dao) *feedingBatchDo {
	f.DO = *do.(*gen.DO)
	return f
}
