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

func newPastureMicroIndex(db *gorm.DB, opts ...gen.DOOption) pastureMicroIndex {
	_pastureMicroIndex := pastureMicroIndex{}

	_pastureMicroIndex.pastureMicroIndexDo.UseDB(db, opts...)
	_pastureMicroIndex.pastureMicroIndexDo.UseModel(&pasture.PastureMicroIndex{})

	tableName := _pastureMicroIndex.pastureMicroIndexDo.TableName()
	_pastureMicroIndex.ALL = field.NewAsterisk(tableName)
	_pastureMicroIndex.ID = field.NewUint(tableName, "id")
	_pastureMicroIndex.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureMicroIndex.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureMicroIndex.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureMicroIndex.PastureWaterRecordID = field.NewUint(tableName, "pasture_water_record_id")
	_pastureMicroIndex.MicroIndex1 = field.NewFloat64(tableName, "micro_index1")
	_pastureMicroIndex.MicroIndex2 = field.NewFloat64(tableName, "micro_index2")
	_pastureMicroIndex.MicroIndex3 = field.NewFloat64(tableName, "micro_index3")

	_pastureMicroIndex.fillFieldMap()

	return _pastureMicroIndex
}

type pastureMicroIndex struct {
	pastureMicroIndexDo pastureMicroIndexDo

	ALL                  field.Asterisk
	ID                   field.Uint
	CreatedAt            field.Time
	UpdatedAt            field.Time
	DeletedAt            field.Field
	PastureWaterRecordID field.Uint
	MicroIndex1          field.Float64
	MicroIndex2          field.Float64
	MicroIndex3          field.Float64

	fieldMap map[string]field.Expr
}

func (p pastureMicroIndex) Table(newTableName string) *pastureMicroIndex {
	p.pastureMicroIndexDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureMicroIndex) As(alias string) *pastureMicroIndex {
	p.pastureMicroIndexDo.DO = *(p.pastureMicroIndexDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureMicroIndex) updateTableName(table string) *pastureMicroIndex {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.PastureWaterRecordID = field.NewUint(table, "pasture_water_record_id")
	p.MicroIndex1 = field.NewFloat64(table, "micro_index1")
	p.MicroIndex2 = field.NewFloat64(table, "micro_index2")
	p.MicroIndex3 = field.NewFloat64(table, "micro_index3")

	p.fillFieldMap()

	return p
}

func (p *pastureMicroIndex) WithContext(ctx context.Context) IPastureMicroIndexDo {
	return p.pastureMicroIndexDo.WithContext(ctx)
}

func (p pastureMicroIndex) TableName() string { return p.pastureMicroIndexDo.TableName() }

func (p pastureMicroIndex) Alias() string { return p.pastureMicroIndexDo.Alias() }

func (p pastureMicroIndex) Columns(cols ...field.Expr) gen.Columns {
	return p.pastureMicroIndexDo.Columns(cols...)
}

func (p *pastureMicroIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureMicroIndex) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["pasture_water_record_id"] = p.PastureWaterRecordID
	p.fieldMap["micro_index1"] = p.MicroIndex1
	p.fieldMap["micro_index2"] = p.MicroIndex2
	p.fieldMap["micro_index3"] = p.MicroIndex3
}

func (p pastureMicroIndex) clone(db *gorm.DB) pastureMicroIndex {
	p.pastureMicroIndexDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureMicroIndex) replaceDB(db *gorm.DB) pastureMicroIndex {
	p.pastureMicroIndexDo.ReplaceDB(db)
	return p
}

type pastureMicroIndexDo struct{ gen.DO }

type IPastureMicroIndexDo interface {
	gen.SubQuery
	Debug() IPastureMicroIndexDo
	WithContext(ctx context.Context) IPastureMicroIndexDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureMicroIndexDo
	WriteDB() IPastureMicroIndexDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureMicroIndexDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureMicroIndexDo
	Not(conds ...gen.Condition) IPastureMicroIndexDo
	Or(conds ...gen.Condition) IPastureMicroIndexDo
	Select(conds ...field.Expr) IPastureMicroIndexDo
	Where(conds ...gen.Condition) IPastureMicroIndexDo
	Order(conds ...field.Expr) IPastureMicroIndexDo
	Distinct(cols ...field.Expr) IPastureMicroIndexDo
	Omit(cols ...field.Expr) IPastureMicroIndexDo
	Join(table schema.Tabler, on ...field.Expr) IPastureMicroIndexDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureMicroIndexDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureMicroIndexDo
	Group(cols ...field.Expr) IPastureMicroIndexDo
	Having(conds ...gen.Condition) IPastureMicroIndexDo
	Limit(limit int) IPastureMicroIndexDo
	Offset(offset int) IPastureMicroIndexDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureMicroIndexDo
	Unscoped() IPastureMicroIndexDo
	Create(values ...*pasture.PastureMicroIndex) error
	CreateInBatches(values []*pasture.PastureMicroIndex, batchSize int) error
	Save(values ...*pasture.PastureMicroIndex) error
	First() (*pasture.PastureMicroIndex, error)
	Take() (*pasture.PastureMicroIndex, error)
	Last() (*pasture.PastureMicroIndex, error)
	Find() ([]*pasture.PastureMicroIndex, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureMicroIndex, err error)
	FindInBatches(result *[]*pasture.PastureMicroIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureMicroIndex) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureMicroIndexDo
	Assign(attrs ...field.AssignExpr) IPastureMicroIndexDo
	Joins(fields ...field.RelationField) IPastureMicroIndexDo
	Preload(fields ...field.RelationField) IPastureMicroIndexDo
	FirstOrInit() (*pasture.PastureMicroIndex, error)
	FirstOrCreate() (*pasture.PastureMicroIndex, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureMicroIndex, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureMicroIndexDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureMicroIndexDo) Debug() IPastureMicroIndexDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureMicroIndexDo) WithContext(ctx context.Context) IPastureMicroIndexDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureMicroIndexDo) ReadDB() IPastureMicroIndexDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureMicroIndexDo) WriteDB() IPastureMicroIndexDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureMicroIndexDo) Session(config *gorm.Session) IPastureMicroIndexDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureMicroIndexDo) Clauses(conds ...clause.Expression) IPastureMicroIndexDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureMicroIndexDo) Returning(value interface{}, columns ...string) IPastureMicroIndexDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureMicroIndexDo) Not(conds ...gen.Condition) IPastureMicroIndexDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureMicroIndexDo) Or(conds ...gen.Condition) IPastureMicroIndexDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureMicroIndexDo) Select(conds ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureMicroIndexDo) Where(conds ...gen.Condition) IPastureMicroIndexDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureMicroIndexDo) Order(conds ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureMicroIndexDo) Distinct(cols ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureMicroIndexDo) Omit(cols ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureMicroIndexDo) Join(table schema.Tabler, on ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureMicroIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureMicroIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureMicroIndexDo) Group(cols ...field.Expr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureMicroIndexDo) Having(conds ...gen.Condition) IPastureMicroIndexDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureMicroIndexDo) Limit(limit int) IPastureMicroIndexDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureMicroIndexDo) Offset(offset int) IPastureMicroIndexDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureMicroIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureMicroIndexDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureMicroIndexDo) Unscoped() IPastureMicroIndexDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureMicroIndexDo) Create(values ...*pasture.PastureMicroIndex) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureMicroIndexDo) CreateInBatches(values []*pasture.PastureMicroIndex, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureMicroIndexDo) Save(values ...*pasture.PastureMicroIndex) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureMicroIndexDo) First() (*pasture.PastureMicroIndex, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureMicroIndex), nil
	}
}

func (p pastureMicroIndexDo) Take() (*pasture.PastureMicroIndex, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureMicroIndex), nil
	}
}

func (p pastureMicroIndexDo) Last() (*pasture.PastureMicroIndex, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureMicroIndex), nil
	}
}

func (p pastureMicroIndexDo) Find() ([]*pasture.PastureMicroIndex, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureMicroIndex), err
}

func (p pastureMicroIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureMicroIndex, err error) {
	buf := make([]*pasture.PastureMicroIndex, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureMicroIndexDo) FindInBatches(result *[]*pasture.PastureMicroIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureMicroIndexDo) Attrs(attrs ...field.AssignExpr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureMicroIndexDo) Assign(attrs ...field.AssignExpr) IPastureMicroIndexDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureMicroIndexDo) Joins(fields ...field.RelationField) IPastureMicroIndexDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureMicroIndexDo) Preload(fields ...field.RelationField) IPastureMicroIndexDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureMicroIndexDo) FirstOrInit() (*pasture.PastureMicroIndex, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureMicroIndex), nil
	}
}

func (p pastureMicroIndexDo) FirstOrCreate() (*pasture.PastureMicroIndex, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureMicroIndex), nil
	}
}

func (p pastureMicroIndexDo) FindByPage(offset int, limit int) (result []*pasture.PastureMicroIndex, count int64, err error) {
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

func (p pastureMicroIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureMicroIndexDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureMicroIndexDo) Delete(models ...*pasture.PastureMicroIndex) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureMicroIndexDo) withDO(do gen.Dao) *pastureMicroIndexDo {
	p.DO = *do.(*gen.DO)
	return p
}
