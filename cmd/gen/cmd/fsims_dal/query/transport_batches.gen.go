// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/product"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newTransportBatch(db *gorm.DB, opts ...gen.DOOption) transportBatch {
	_transportBatch := transportBatch{}

	_transportBatch.transportBatchDo.UseDB(db, opts...)
	_transportBatch.transportBatchDo.UseModel(&coldchain.TransportBatch{})

	tableName := _transportBatch.transportBatchDo.TableName()
	_transportBatch.ALL = field.NewAsterisk(tableName)
	_transportBatch.ID = field.NewUint(tableName, "id")
	_transportBatch.CreatedAt = field.NewTime(tableName, "created_at")
	_transportBatch.UpdatedAt = field.NewTime(tableName, "updated_at")
	_transportBatch.DeletedAt = field.NewField(tableName, "deleted_at")
	_transportBatch.BatchNumber = field.NewString(tableName, "batch_number")
	_transportBatch.TVNumber = field.NewString(tableName, "tv_number")
	_transportBatch.State = field.NewInt(tableName, "state")
	_transportBatch.Worker = field.NewString(tableName, "worker")
	_transportBatch.MallNumber = field.NewString(tableName, "mall_number")
	_transportBatch.StartTime = field.NewTime(tableName, "start_time")
	_transportBatch.EndTime = field.NewTime(tableName, "end_time")
	_transportBatch.Products = transportBatchHasManyProducts{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Products", "product.PackageProduct"),
	}

	_transportBatch.fillFieldMap()

	return _transportBatch
}

type transportBatch struct {
	transportBatchDo transportBatchDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	BatchNumber field.String
	TVNumber    field.String
	State       field.Int
	Worker      field.String
	MallNumber  field.String
	StartTime   field.Time
	EndTime     field.Time
	Products    transportBatchHasManyProducts

	fieldMap map[string]field.Expr
}

func (t transportBatch) Table(newTableName string) *transportBatch {
	t.transportBatchDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t transportBatch) As(alias string) *transportBatch {
	t.transportBatchDo.DO = *(t.transportBatchDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *transportBatch) updateTableName(table string) *transportBatch {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.BatchNumber = field.NewString(table, "batch_number")
	t.TVNumber = field.NewString(table, "tv_number")
	t.State = field.NewInt(table, "state")
	t.Worker = field.NewString(table, "worker")
	t.MallNumber = field.NewString(table, "mall_number")
	t.StartTime = field.NewTime(table, "start_time")
	t.EndTime = field.NewTime(table, "end_time")

	t.fillFieldMap()

	return t
}

func (t *transportBatch) WithContext(ctx context.Context) ITransportBatchDo {
	return t.transportBatchDo.WithContext(ctx)
}

func (t transportBatch) TableName() string { return t.transportBatchDo.TableName() }

func (t transportBatch) Alias() string { return t.transportBatchDo.Alias() }

func (t transportBatch) Columns(cols ...field.Expr) gen.Columns {
	return t.transportBatchDo.Columns(cols...)
}

func (t *transportBatch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *transportBatch) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["batch_number"] = t.BatchNumber
	t.fieldMap["tv_number"] = t.TVNumber
	t.fieldMap["state"] = t.State
	t.fieldMap["worker"] = t.Worker
	t.fieldMap["mall_number"] = t.MallNumber
	t.fieldMap["start_time"] = t.StartTime
	t.fieldMap["end_time"] = t.EndTime

}

func (t transportBatch) clone(db *gorm.DB) transportBatch {
	t.transportBatchDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t transportBatch) replaceDB(db *gorm.DB) transportBatch {
	t.transportBatchDo.ReplaceDB(db)
	return t
}

type transportBatchHasManyProducts struct {
	db *gorm.DB

	field.RelationField
}

func (a transportBatchHasManyProducts) Where(conds ...field.Expr) *transportBatchHasManyProducts {
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

func (a transportBatchHasManyProducts) WithContext(ctx context.Context) *transportBatchHasManyProducts {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a transportBatchHasManyProducts) Session(session *gorm.Session) *transportBatchHasManyProducts {
	a.db = a.db.Session(session)
	return &a
}

func (a transportBatchHasManyProducts) Model(m *coldchain.TransportBatch) *transportBatchHasManyProductsTx {
	return &transportBatchHasManyProductsTx{a.db.Model(m).Association(a.Name())}
}

type transportBatchHasManyProductsTx struct{ tx *gorm.Association }

func (a transportBatchHasManyProductsTx) Find() (result []*product.PackageProduct, err error) {
	return result, a.tx.Find(&result)
}

func (a transportBatchHasManyProductsTx) Append(values ...*product.PackageProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a transportBatchHasManyProductsTx) Replace(values ...*product.PackageProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a transportBatchHasManyProductsTx) Delete(values ...*product.PackageProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a transportBatchHasManyProductsTx) Clear() error {
	return a.tx.Clear()
}

func (a transportBatchHasManyProductsTx) Count() int64 {
	return a.tx.Count()
}

type transportBatchDo struct{ gen.DO }

type ITransportBatchDo interface {
	gen.SubQuery
	Debug() ITransportBatchDo
	WithContext(ctx context.Context) ITransportBatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITransportBatchDo
	WriteDB() ITransportBatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITransportBatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITransportBatchDo
	Not(conds ...gen.Condition) ITransportBatchDo
	Or(conds ...gen.Condition) ITransportBatchDo
	Select(conds ...field.Expr) ITransportBatchDo
	Where(conds ...gen.Condition) ITransportBatchDo
	Order(conds ...field.Expr) ITransportBatchDo
	Distinct(cols ...field.Expr) ITransportBatchDo
	Omit(cols ...field.Expr) ITransportBatchDo
	Join(table schema.Tabler, on ...field.Expr) ITransportBatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITransportBatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITransportBatchDo
	Group(cols ...field.Expr) ITransportBatchDo
	Having(conds ...gen.Condition) ITransportBatchDo
	Limit(limit int) ITransportBatchDo
	Offset(offset int) ITransportBatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITransportBatchDo
	Unscoped() ITransportBatchDo
	Create(values ...*coldchain.TransportBatch) error
	CreateInBatches(values []*coldchain.TransportBatch, batchSize int) error
	Save(values ...*coldchain.TransportBatch) error
	First() (*coldchain.TransportBatch, error)
	Take() (*coldchain.TransportBatch, error)
	Last() (*coldchain.TransportBatch, error)
	Find() ([]*coldchain.TransportBatch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*coldchain.TransportBatch, err error)
	FindInBatches(result *[]*coldchain.TransportBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*coldchain.TransportBatch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITransportBatchDo
	Assign(attrs ...field.AssignExpr) ITransportBatchDo
	Joins(fields ...field.RelationField) ITransportBatchDo
	Preload(fields ...field.RelationField) ITransportBatchDo
	FirstOrInit() (*coldchain.TransportBatch, error)
	FirstOrCreate() (*coldchain.TransportBatch, error)
	FindByPage(offset int, limit int) (result []*coldchain.TransportBatch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITransportBatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t transportBatchDo) Debug() ITransportBatchDo {
	return t.withDO(t.DO.Debug())
}

func (t transportBatchDo) WithContext(ctx context.Context) ITransportBatchDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t transportBatchDo) ReadDB() ITransportBatchDo {
	return t.Clauses(dbresolver.Read)
}

func (t transportBatchDo) WriteDB() ITransportBatchDo {
	return t.Clauses(dbresolver.Write)
}

func (t transportBatchDo) Session(config *gorm.Session) ITransportBatchDo {
	return t.withDO(t.DO.Session(config))
}

func (t transportBatchDo) Clauses(conds ...clause.Expression) ITransportBatchDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t transportBatchDo) Returning(value interface{}, columns ...string) ITransportBatchDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t transportBatchDo) Not(conds ...gen.Condition) ITransportBatchDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t transportBatchDo) Or(conds ...gen.Condition) ITransportBatchDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t transportBatchDo) Select(conds ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t transportBatchDo) Where(conds ...gen.Condition) ITransportBatchDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t transportBatchDo) Order(conds ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t transportBatchDo) Distinct(cols ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t transportBatchDo) Omit(cols ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t transportBatchDo) Join(table schema.Tabler, on ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t transportBatchDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t transportBatchDo) RightJoin(table schema.Tabler, on ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t transportBatchDo) Group(cols ...field.Expr) ITransportBatchDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t transportBatchDo) Having(conds ...gen.Condition) ITransportBatchDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t transportBatchDo) Limit(limit int) ITransportBatchDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t transportBatchDo) Offset(offset int) ITransportBatchDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t transportBatchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITransportBatchDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t transportBatchDo) Unscoped() ITransportBatchDo {
	return t.withDO(t.DO.Unscoped())
}

func (t transportBatchDo) Create(values ...*coldchain.TransportBatch) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t transportBatchDo) CreateInBatches(values []*coldchain.TransportBatch, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t transportBatchDo) Save(values ...*coldchain.TransportBatch) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t transportBatchDo) First() (*coldchain.TransportBatch, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*coldchain.TransportBatch), nil
	}
}

func (t transportBatchDo) Take() (*coldchain.TransportBatch, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*coldchain.TransportBatch), nil
	}
}

func (t transportBatchDo) Last() (*coldchain.TransportBatch, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*coldchain.TransportBatch), nil
	}
}

func (t transportBatchDo) Find() ([]*coldchain.TransportBatch, error) {
	result, err := t.DO.Find()
	return result.([]*coldchain.TransportBatch), err
}

func (t transportBatchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*coldchain.TransportBatch, err error) {
	buf := make([]*coldchain.TransportBatch, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t transportBatchDo) FindInBatches(result *[]*coldchain.TransportBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t transportBatchDo) Attrs(attrs ...field.AssignExpr) ITransportBatchDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t transportBatchDo) Assign(attrs ...field.AssignExpr) ITransportBatchDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t transportBatchDo) Joins(fields ...field.RelationField) ITransportBatchDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t transportBatchDo) Preload(fields ...field.RelationField) ITransportBatchDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t transportBatchDo) FirstOrInit() (*coldchain.TransportBatch, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*coldchain.TransportBatch), nil
	}
}

func (t transportBatchDo) FirstOrCreate() (*coldchain.TransportBatch, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*coldchain.TransportBatch), nil
	}
}

func (t transportBatchDo) FindByPage(offset int, limit int) (result []*coldchain.TransportBatch, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t transportBatchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t transportBatchDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t transportBatchDo) Delete(models ...*coldchain.TransportBatch) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *transportBatchDo) withDO(do gen.Dao) *transportBatchDo {
	t.DO = *do.(*gen.DO)
	return t
}
