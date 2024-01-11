// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPackageReceiveRecord(db *gorm.DB, opts ...gen.DOOption) packageReceiveRecord {
	_packageReceiveRecord := packageReceiveRecord{}

	_packageReceiveRecord.packageReceiveRecordDo.UseDB(db, opts...)
	_packageReceiveRecord.packageReceiveRecordDo.UseModel(&warehouse.PackageReceiveRecord{})

	tableName := _packageReceiveRecord.packageReceiveRecordDo.TableName()
	_packageReceiveRecord.ALL = field.NewAsterisk(tableName)
	_packageReceiveRecord.ID = field.NewUint(tableName, "id")
	_packageReceiveRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_packageReceiveRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_packageReceiveRecord.DeletedAt = field.NewField(tableName, "deleted_at")
	_packageReceiveRecord.ProductNumber = field.NewString(tableName, "product_number")
	_packageReceiveRecord.PID = field.NewString(tableName, "p_id")
	_packageReceiveRecord.SourceNumber = field.NewString(tableName, "source_number")
	_packageReceiveRecord.SourceName = field.NewString(tableName, "source_name")
	_packageReceiveRecord.State = field.NewInt(tableName, "state")
	_packageReceiveRecord.Operator = field.NewString(tableName, "operator")
	_packageReceiveRecord.ReceiveTime = field.NewTime(tableName, "receive_time")
	_packageReceiveRecord.ConfirmTime = field.NewTime(tableName, "confirm_time")
	_packageReceiveRecord.PacNumber = field.NewString(tableName, "pac_number")

	_packageReceiveRecord.fillFieldMap()

	return _packageReceiveRecord
}

type packageReceiveRecord struct {
	packageReceiveRecordDo packageReceiveRecordDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	ProductNumber field.String
	PID           field.String
	SourceNumber  field.String
	SourceName    field.String
	State         field.Int
	Operator      field.String
	ReceiveTime   field.Time
	ConfirmTime   field.Time
	PacNumber     field.String

	fieldMap map[string]field.Expr
}

func (p packageReceiveRecord) Table(newTableName string) *packageReceiveRecord {
	p.packageReceiveRecordDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p packageReceiveRecord) As(alias string) *packageReceiveRecord {
	p.packageReceiveRecordDo.DO = *(p.packageReceiveRecordDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *packageReceiveRecord) updateTableName(table string) *packageReceiveRecord {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.ProductNumber = field.NewString(table, "product_number")
	p.PID = field.NewString(table, "p_id")
	p.SourceNumber = field.NewString(table, "source_number")
	p.SourceName = field.NewString(table, "source_name")
	p.State = field.NewInt(table, "state")
	p.Operator = field.NewString(table, "operator")
	p.ReceiveTime = field.NewTime(table, "receive_time")
	p.ConfirmTime = field.NewTime(table, "confirm_time")
	p.PacNumber = field.NewString(table, "pac_number")

	p.fillFieldMap()

	return p
}

func (p *packageReceiveRecord) WithContext(ctx context.Context) IPackageReceiveRecordDo {
	return p.packageReceiveRecordDo.WithContext(ctx)
}

func (p packageReceiveRecord) TableName() string { return p.packageReceiveRecordDo.TableName() }

func (p packageReceiveRecord) Alias() string { return p.packageReceiveRecordDo.Alias() }

func (p packageReceiveRecord) Columns(cols ...field.Expr) gen.Columns {
	return p.packageReceiveRecordDo.Columns(cols...)
}

func (p *packageReceiveRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *packageReceiveRecord) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["product_number"] = p.ProductNumber
	p.fieldMap["p_id"] = p.PID
	p.fieldMap["source_number"] = p.SourceNumber
	p.fieldMap["source_name"] = p.SourceName
	p.fieldMap["state"] = p.State
	p.fieldMap["operator"] = p.Operator
	p.fieldMap["receive_time"] = p.ReceiveTime
	p.fieldMap["confirm_time"] = p.ConfirmTime
	p.fieldMap["pac_number"] = p.PacNumber
}

func (p packageReceiveRecord) clone(db *gorm.DB) packageReceiveRecord {
	p.packageReceiveRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p packageReceiveRecord) replaceDB(db *gorm.DB) packageReceiveRecord {
	p.packageReceiveRecordDo.ReplaceDB(db)
	return p
}

type packageReceiveRecordDo struct{ gen.DO }

type IPackageReceiveRecordDo interface {
	gen.SubQuery
	Debug() IPackageReceiveRecordDo
	WithContext(ctx context.Context) IPackageReceiveRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPackageReceiveRecordDo
	WriteDB() IPackageReceiveRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPackageReceiveRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPackageReceiveRecordDo
	Not(conds ...gen.Condition) IPackageReceiveRecordDo
	Or(conds ...gen.Condition) IPackageReceiveRecordDo
	Select(conds ...field.Expr) IPackageReceiveRecordDo
	Where(conds ...gen.Condition) IPackageReceiveRecordDo
	Order(conds ...field.Expr) IPackageReceiveRecordDo
	Distinct(cols ...field.Expr) IPackageReceiveRecordDo
	Omit(cols ...field.Expr) IPackageReceiveRecordDo
	Join(table schema.Tabler, on ...field.Expr) IPackageReceiveRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPackageReceiveRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPackageReceiveRecordDo
	Group(cols ...field.Expr) IPackageReceiveRecordDo
	Having(conds ...gen.Condition) IPackageReceiveRecordDo
	Limit(limit int) IPackageReceiveRecordDo
	Offset(offset int) IPackageReceiveRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPackageReceiveRecordDo
	Unscoped() IPackageReceiveRecordDo
	Create(values ...*warehouse.PackageReceiveRecord) error
	CreateInBatches(values []*warehouse.PackageReceiveRecord, batchSize int) error
	Save(values ...*warehouse.PackageReceiveRecord) error
	First() (*warehouse.PackageReceiveRecord, error)
	Take() (*warehouse.PackageReceiveRecord, error)
	Last() (*warehouse.PackageReceiveRecord, error)
	Find() ([]*warehouse.PackageReceiveRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*warehouse.PackageReceiveRecord, err error)
	FindInBatches(result *[]*warehouse.PackageReceiveRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*warehouse.PackageReceiveRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPackageReceiveRecordDo
	Assign(attrs ...field.AssignExpr) IPackageReceiveRecordDo
	Joins(fields ...field.RelationField) IPackageReceiveRecordDo
	Preload(fields ...field.RelationField) IPackageReceiveRecordDo
	FirstOrInit() (*warehouse.PackageReceiveRecord, error)
	FirstOrCreate() (*warehouse.PackageReceiveRecord, error)
	FindByPage(offset int, limit int) (result []*warehouse.PackageReceiveRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPackageReceiveRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p packageReceiveRecordDo) Debug() IPackageReceiveRecordDo {
	return p.withDO(p.DO.Debug())
}

func (p packageReceiveRecordDo) WithContext(ctx context.Context) IPackageReceiveRecordDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p packageReceiveRecordDo) ReadDB() IPackageReceiveRecordDo {
	return p.Clauses(dbresolver.Read)
}

func (p packageReceiveRecordDo) WriteDB() IPackageReceiveRecordDo {
	return p.Clauses(dbresolver.Write)
}

func (p packageReceiveRecordDo) Session(config *gorm.Session) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Session(config))
}

func (p packageReceiveRecordDo) Clauses(conds ...clause.Expression) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p packageReceiveRecordDo) Returning(value interface{}, columns ...string) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p packageReceiveRecordDo) Not(conds ...gen.Condition) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p packageReceiveRecordDo) Or(conds ...gen.Condition) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p packageReceiveRecordDo) Select(conds ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p packageReceiveRecordDo) Where(conds ...gen.Condition) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p packageReceiveRecordDo) Order(conds ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p packageReceiveRecordDo) Distinct(cols ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p packageReceiveRecordDo) Omit(cols ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p packageReceiveRecordDo) Join(table schema.Tabler, on ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p packageReceiveRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p packageReceiveRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p packageReceiveRecordDo) Group(cols ...field.Expr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p packageReceiveRecordDo) Having(conds ...gen.Condition) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p packageReceiveRecordDo) Limit(limit int) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p packageReceiveRecordDo) Offset(offset int) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p packageReceiveRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p packageReceiveRecordDo) Unscoped() IPackageReceiveRecordDo {
	return p.withDO(p.DO.Unscoped())
}

func (p packageReceiveRecordDo) Create(values ...*warehouse.PackageReceiveRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p packageReceiveRecordDo) CreateInBatches(values []*warehouse.PackageReceiveRecord, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p packageReceiveRecordDo) Save(values ...*warehouse.PackageReceiveRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p packageReceiveRecordDo) First() (*warehouse.PackageReceiveRecord, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*warehouse.PackageReceiveRecord), nil
	}
}

func (p packageReceiveRecordDo) Take() (*warehouse.PackageReceiveRecord, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*warehouse.PackageReceiveRecord), nil
	}
}

func (p packageReceiveRecordDo) Last() (*warehouse.PackageReceiveRecord, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*warehouse.PackageReceiveRecord), nil
	}
}

func (p packageReceiveRecordDo) Find() ([]*warehouse.PackageReceiveRecord, error) {
	result, err := p.DO.Find()
	return result.([]*warehouse.PackageReceiveRecord), err
}

func (p packageReceiveRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*warehouse.PackageReceiveRecord, err error) {
	buf := make([]*warehouse.PackageReceiveRecord, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p packageReceiveRecordDo) FindInBatches(result *[]*warehouse.PackageReceiveRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p packageReceiveRecordDo) Attrs(attrs ...field.AssignExpr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p packageReceiveRecordDo) Assign(attrs ...field.AssignExpr) IPackageReceiveRecordDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p packageReceiveRecordDo) Joins(fields ...field.RelationField) IPackageReceiveRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p packageReceiveRecordDo) Preload(fields ...field.RelationField) IPackageReceiveRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p packageReceiveRecordDo) FirstOrInit() (*warehouse.PackageReceiveRecord, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*warehouse.PackageReceiveRecord), nil
	}
}

func (p packageReceiveRecordDo) FirstOrCreate() (*warehouse.PackageReceiveRecord, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*warehouse.PackageReceiveRecord), nil
	}
}

func (p packageReceiveRecordDo) FindByPage(offset int, limit int) (result []*warehouse.PackageReceiveRecord, count int64, err error) {
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

func (p packageReceiveRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p packageReceiveRecordDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p packageReceiveRecordDo) Delete(models ...*warehouse.PackageReceiveRecord) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *packageReceiveRecordDo) withDO(do gen.Dao) *packageReceiveRecordDo {
	p.DO = *do.(*gen.DO)
	return p
}
