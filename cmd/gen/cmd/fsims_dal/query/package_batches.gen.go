// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/pack"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPackageBatch(db *gorm.DB, opts ...gen.DOOption) packageBatch {
	_packageBatch := packageBatch{}

	_packageBatch.packageBatchDo.UseDB(db, opts...)
	_packageBatch.packageBatchDo.UseModel(&pack.PackageBatch{})

	tableName := _packageBatch.packageBatchDo.TableName()
	_packageBatch.ALL = field.NewAsterisk(tableName)
	_packageBatch.ID = field.NewUint(tableName, "id")
	_packageBatch.CreatedAt = field.NewTime(tableName, "created_at")
	_packageBatch.UpdatedAt = field.NewTime(tableName, "updated_at")
	_packageBatch.DeletedAt = field.NewField(tableName, "deleted_at")
	_packageBatch.BatchNumber = field.NewString(tableName, "batch_number")
	_packageBatch.HouseNumber = field.NewString(tableName, "house_number")
	_packageBatch.State = field.NewInt(tableName, "state")
	_packageBatch.PID = field.NewString(tableName, "p_id")
	_packageBatch.Worker = field.NewString(tableName, "worker")
	_packageBatch.StartTime = field.NewTime(tableName, "start_time")
	_packageBatch.EndTime = field.NewTime(tableName, "end_time")
	_packageBatch.ProductNumber = field.NewString(tableName, "product_number")

	_packageBatch.fillFieldMap()

	return _packageBatch
}

type packageBatch struct {
	packageBatchDo packageBatchDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	BatchNumber   field.String
	HouseNumber   field.String
	State         field.Int
	PID           field.String
	Worker        field.String
	StartTime     field.Time
	EndTime       field.Time
	ProductNumber field.String

	fieldMap map[string]field.Expr
}

func (p packageBatch) Table(newTableName string) *packageBatch {
	p.packageBatchDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p packageBatch) As(alias string) *packageBatch {
	p.packageBatchDo.DO = *(p.packageBatchDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *packageBatch) updateTableName(table string) *packageBatch {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.BatchNumber = field.NewString(table, "batch_number")
	p.HouseNumber = field.NewString(table, "house_number")
	p.State = field.NewInt(table, "state")
	p.PID = field.NewString(table, "p_id")
	p.Worker = field.NewString(table, "worker")
	p.StartTime = field.NewTime(table, "start_time")
	p.EndTime = field.NewTime(table, "end_time")
	p.ProductNumber = field.NewString(table, "product_number")

	p.fillFieldMap()

	return p
}

func (p *packageBatch) WithContext(ctx context.Context) IPackageBatchDo {
	return p.packageBatchDo.WithContext(ctx)
}

func (p packageBatch) TableName() string { return p.packageBatchDo.TableName() }

func (p packageBatch) Alias() string { return p.packageBatchDo.Alias() }

func (p packageBatch) Columns(cols ...field.Expr) gen.Columns {
	return p.packageBatchDo.Columns(cols...)
}

func (p *packageBatch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *packageBatch) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["batch_number"] = p.BatchNumber
	p.fieldMap["house_number"] = p.HouseNumber
	p.fieldMap["state"] = p.State
	p.fieldMap["p_id"] = p.PID
	p.fieldMap["worker"] = p.Worker
	p.fieldMap["start_time"] = p.StartTime
	p.fieldMap["end_time"] = p.EndTime
	p.fieldMap["product_number"] = p.ProductNumber
}

func (p packageBatch) clone(db *gorm.DB) packageBatch {
	p.packageBatchDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p packageBatch) replaceDB(db *gorm.DB) packageBatch {
	p.packageBatchDo.ReplaceDB(db)
	return p
}

type packageBatchDo struct{ gen.DO }

type IPackageBatchDo interface {
	gen.SubQuery
	Debug() IPackageBatchDo
	WithContext(ctx context.Context) IPackageBatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPackageBatchDo
	WriteDB() IPackageBatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPackageBatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPackageBatchDo
	Not(conds ...gen.Condition) IPackageBatchDo
	Or(conds ...gen.Condition) IPackageBatchDo
	Select(conds ...field.Expr) IPackageBatchDo
	Where(conds ...gen.Condition) IPackageBatchDo
	Order(conds ...field.Expr) IPackageBatchDo
	Distinct(cols ...field.Expr) IPackageBatchDo
	Omit(cols ...field.Expr) IPackageBatchDo
	Join(table schema.Tabler, on ...field.Expr) IPackageBatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPackageBatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPackageBatchDo
	Group(cols ...field.Expr) IPackageBatchDo
	Having(conds ...gen.Condition) IPackageBatchDo
	Limit(limit int) IPackageBatchDo
	Offset(offset int) IPackageBatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPackageBatchDo
	Unscoped() IPackageBatchDo
	Create(values ...*pack.PackageBatch) error
	CreateInBatches(values []*pack.PackageBatch, batchSize int) error
	Save(values ...*pack.PackageBatch) error
	First() (*pack.PackageBatch, error)
	Take() (*pack.PackageBatch, error)
	Last() (*pack.PackageBatch, error)
	Find() ([]*pack.PackageBatch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pack.PackageBatch, err error)
	FindInBatches(result *[]*pack.PackageBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pack.PackageBatch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPackageBatchDo
	Assign(attrs ...field.AssignExpr) IPackageBatchDo
	Joins(fields ...field.RelationField) IPackageBatchDo
	Preload(fields ...field.RelationField) IPackageBatchDo
	FirstOrInit() (*pack.PackageBatch, error)
	FirstOrCreate() (*pack.PackageBatch, error)
	FindByPage(offset int, limit int) (result []*pack.PackageBatch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPackageBatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p packageBatchDo) Debug() IPackageBatchDo {
	return p.withDO(p.DO.Debug())
}

func (p packageBatchDo) WithContext(ctx context.Context) IPackageBatchDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p packageBatchDo) ReadDB() IPackageBatchDo {
	return p.Clauses(dbresolver.Read)
}

func (p packageBatchDo) WriteDB() IPackageBatchDo {
	return p.Clauses(dbresolver.Write)
}

func (p packageBatchDo) Session(config *gorm.Session) IPackageBatchDo {
	return p.withDO(p.DO.Session(config))
}

func (p packageBatchDo) Clauses(conds ...clause.Expression) IPackageBatchDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p packageBatchDo) Returning(value interface{}, columns ...string) IPackageBatchDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p packageBatchDo) Not(conds ...gen.Condition) IPackageBatchDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p packageBatchDo) Or(conds ...gen.Condition) IPackageBatchDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p packageBatchDo) Select(conds ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p packageBatchDo) Where(conds ...gen.Condition) IPackageBatchDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p packageBatchDo) Order(conds ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p packageBatchDo) Distinct(cols ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p packageBatchDo) Omit(cols ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p packageBatchDo) Join(table schema.Tabler, on ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p packageBatchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p packageBatchDo) RightJoin(table schema.Tabler, on ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p packageBatchDo) Group(cols ...field.Expr) IPackageBatchDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p packageBatchDo) Having(conds ...gen.Condition) IPackageBatchDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p packageBatchDo) Limit(limit int) IPackageBatchDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p packageBatchDo) Offset(offset int) IPackageBatchDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p packageBatchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPackageBatchDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p packageBatchDo) Unscoped() IPackageBatchDo {
	return p.withDO(p.DO.Unscoped())
}

func (p packageBatchDo) Create(values ...*pack.PackageBatch) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p packageBatchDo) CreateInBatches(values []*pack.PackageBatch, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p packageBatchDo) Save(values ...*pack.PackageBatch) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p packageBatchDo) First() (*pack.PackageBatch, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pack.PackageBatch), nil
	}
}

func (p packageBatchDo) Take() (*pack.PackageBatch, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pack.PackageBatch), nil
	}
}

func (p packageBatchDo) Last() (*pack.PackageBatch, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pack.PackageBatch), nil
	}
}

func (p packageBatchDo) Find() ([]*pack.PackageBatch, error) {
	result, err := p.DO.Find()
	return result.([]*pack.PackageBatch), err
}

func (p packageBatchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pack.PackageBatch, err error) {
	buf := make([]*pack.PackageBatch, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p packageBatchDo) FindInBatches(result *[]*pack.PackageBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p packageBatchDo) Attrs(attrs ...field.AssignExpr) IPackageBatchDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p packageBatchDo) Assign(attrs ...field.AssignExpr) IPackageBatchDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p packageBatchDo) Joins(fields ...field.RelationField) IPackageBatchDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p packageBatchDo) Preload(fields ...field.RelationField) IPackageBatchDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p packageBatchDo) FirstOrInit() (*pack.PackageBatch, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pack.PackageBatch), nil
	}
}

func (p packageBatchDo) FirstOrCreate() (*pack.PackageBatch, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pack.PackageBatch), nil
	}
}

func (p packageBatchDo) FindByPage(offset int, limit int) (result []*pack.PackageBatch, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p packageBatchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p packageBatchDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p packageBatchDo) Delete(models ...*pack.PackageBatch) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *packageBatchDo) withDO(do gen.Dao) *packageBatchDo {
	p.DO = *do.(*gen.DO)
	return p
}
