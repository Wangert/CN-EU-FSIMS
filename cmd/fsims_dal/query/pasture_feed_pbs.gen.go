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

func newPastureFeedPb(db *gorm.DB, opts ...gen.DOOption) pastureFeedPb {
	_pastureFeedPb := pastureFeedPb{}

	_pastureFeedPb.pastureFeedPbDo.UseDB(db, opts...)
	_pastureFeedPb.pastureFeedPbDo.UseModel(&pasture.PastureFeedPb{})

	tableName := _pastureFeedPb.pastureFeedPbDo.TableName()
	_pastureFeedPb.ALL = field.NewAsterisk(tableName)
	_pastureFeedPb.ID = field.NewUint(tableName, "id")
	_pastureFeedPb.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureFeedPb.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureFeedPb.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureFeedPb.PastureFeedHeavyMetalID = field.NewUint(tableName, "pasture_feed_heavy_metal_id")
	_pastureFeedPb.Pb1 = field.NewFloat64(tableName, "pb1")
	_pastureFeedPb.Pb2 = field.NewFloat64(tableName, "pb2")
	_pastureFeedPb.Pb3 = field.NewFloat64(tableName, "pb3")
	_pastureFeedPb.Pb4 = field.NewFloat64(tableName, "pb4")
	_pastureFeedPb.Pb5 = field.NewFloat64(tableName, "pb5")
	_pastureFeedPb.Pb6 = field.NewFloat64(tableName, "pb6")
	_pastureFeedPb.Pb7 = field.NewFloat64(tableName, "pb7")

	_pastureFeedPb.fillFieldMap()

	return _pastureFeedPb
}

type pastureFeedPb struct {
	pastureFeedPbDo pastureFeedPbDo

	ALL                     field.Asterisk
	ID                      field.Uint
	CreatedAt               field.Time
	UpdatedAt               field.Time
	DeletedAt               field.Field
	PastureFeedHeavyMetalID field.Uint
	Pb1                     field.Float64
	Pb2                     field.Float64
	Pb3                     field.Float64
	Pb4                     field.Float64
	Pb5                     field.Float64
	Pb6                     field.Float64
	Pb7                     field.Float64

	fieldMap map[string]field.Expr
}

func (p pastureFeedPb) Table(newTableName string) *pastureFeedPb {
	p.pastureFeedPbDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureFeedPb) As(alias string) *pastureFeedPb {
	p.pastureFeedPbDo.DO = *(p.pastureFeedPbDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureFeedPb) updateTableName(table string) *pastureFeedPb {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.PastureFeedHeavyMetalID = field.NewUint(table, "pasture_feed_heavy_metal_id")
	p.Pb1 = field.NewFloat64(table, "pb1")
	p.Pb2 = field.NewFloat64(table, "pb2")
	p.Pb3 = field.NewFloat64(table, "pb3")
	p.Pb4 = field.NewFloat64(table, "pb4")
	p.Pb5 = field.NewFloat64(table, "pb5")
	p.Pb6 = field.NewFloat64(table, "pb6")
	p.Pb7 = field.NewFloat64(table, "pb7")

	p.fillFieldMap()

	return p
}

func (p *pastureFeedPb) WithContext(ctx context.Context) IPastureFeedPbDo {
	return p.pastureFeedPbDo.WithContext(ctx)
}

func (p pastureFeedPb) TableName() string { return p.pastureFeedPbDo.TableName() }

func (p pastureFeedPb) Alias() string { return p.pastureFeedPbDo.Alias() }

func (p pastureFeedPb) Columns(cols ...field.Expr) gen.Columns {
	return p.pastureFeedPbDo.Columns(cols...)
}

func (p *pastureFeedPb) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureFeedPb) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["pasture_feed_heavy_metal_id"] = p.PastureFeedHeavyMetalID
	p.fieldMap["pb1"] = p.Pb1
	p.fieldMap["pb2"] = p.Pb2
	p.fieldMap["pb3"] = p.Pb3
	p.fieldMap["pb4"] = p.Pb4
	p.fieldMap["pb5"] = p.Pb5
	p.fieldMap["pb6"] = p.Pb6
	p.fieldMap["pb7"] = p.Pb7
}

func (p pastureFeedPb) clone(db *gorm.DB) pastureFeedPb {
	p.pastureFeedPbDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureFeedPb) replaceDB(db *gorm.DB) pastureFeedPb {
	p.pastureFeedPbDo.ReplaceDB(db)
	return p
}

type pastureFeedPbDo struct{ gen.DO }

type IPastureFeedPbDo interface {
	gen.SubQuery
	Debug() IPastureFeedPbDo
	WithContext(ctx context.Context) IPastureFeedPbDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureFeedPbDo
	WriteDB() IPastureFeedPbDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureFeedPbDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureFeedPbDo
	Not(conds ...gen.Condition) IPastureFeedPbDo
	Or(conds ...gen.Condition) IPastureFeedPbDo
	Select(conds ...field.Expr) IPastureFeedPbDo
	Where(conds ...gen.Condition) IPastureFeedPbDo
	Order(conds ...field.Expr) IPastureFeedPbDo
	Distinct(cols ...field.Expr) IPastureFeedPbDo
	Omit(cols ...field.Expr) IPastureFeedPbDo
	Join(table schema.Tabler, on ...field.Expr) IPastureFeedPbDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureFeedPbDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureFeedPbDo
	Group(cols ...field.Expr) IPastureFeedPbDo
	Having(conds ...gen.Condition) IPastureFeedPbDo
	Limit(limit int) IPastureFeedPbDo
	Offset(offset int) IPastureFeedPbDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureFeedPbDo
	Unscoped() IPastureFeedPbDo
	Create(values ...*pasture.PastureFeedPb) error
	CreateInBatches(values []*pasture.PastureFeedPb, batchSize int) error
	Save(values ...*pasture.PastureFeedPb) error
	First() (*pasture.PastureFeedPb, error)
	Take() (*pasture.PastureFeedPb, error)
	Last() (*pasture.PastureFeedPb, error)
	Find() ([]*pasture.PastureFeedPb, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureFeedPb, err error)
	FindInBatches(result *[]*pasture.PastureFeedPb, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureFeedPb) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureFeedPbDo
	Assign(attrs ...field.AssignExpr) IPastureFeedPbDo
	Joins(fields ...field.RelationField) IPastureFeedPbDo
	Preload(fields ...field.RelationField) IPastureFeedPbDo
	FirstOrInit() (*pasture.PastureFeedPb, error)
	FirstOrCreate() (*pasture.PastureFeedPb, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureFeedPb, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureFeedPbDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureFeedPbDo) Debug() IPastureFeedPbDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureFeedPbDo) WithContext(ctx context.Context) IPastureFeedPbDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureFeedPbDo) ReadDB() IPastureFeedPbDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureFeedPbDo) WriteDB() IPastureFeedPbDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureFeedPbDo) Session(config *gorm.Session) IPastureFeedPbDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureFeedPbDo) Clauses(conds ...clause.Expression) IPastureFeedPbDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureFeedPbDo) Returning(value interface{}, columns ...string) IPastureFeedPbDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureFeedPbDo) Not(conds ...gen.Condition) IPastureFeedPbDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureFeedPbDo) Or(conds ...gen.Condition) IPastureFeedPbDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureFeedPbDo) Select(conds ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureFeedPbDo) Where(conds ...gen.Condition) IPastureFeedPbDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureFeedPbDo) Order(conds ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureFeedPbDo) Distinct(cols ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureFeedPbDo) Omit(cols ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureFeedPbDo) Join(table schema.Tabler, on ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureFeedPbDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureFeedPbDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureFeedPbDo) Group(cols ...field.Expr) IPastureFeedPbDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureFeedPbDo) Having(conds ...gen.Condition) IPastureFeedPbDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureFeedPbDo) Limit(limit int) IPastureFeedPbDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureFeedPbDo) Offset(offset int) IPastureFeedPbDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureFeedPbDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureFeedPbDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureFeedPbDo) Unscoped() IPastureFeedPbDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureFeedPbDo) Create(values ...*pasture.PastureFeedPb) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureFeedPbDo) CreateInBatches(values []*pasture.PastureFeedPb, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureFeedPbDo) Save(values ...*pasture.PastureFeedPb) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureFeedPbDo) First() (*pasture.PastureFeedPb, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedPb), nil
	}
}

func (p pastureFeedPbDo) Take() (*pasture.PastureFeedPb, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedPb), nil
	}
}

func (p pastureFeedPbDo) Last() (*pasture.PastureFeedPb, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedPb), nil
	}
}

func (p pastureFeedPbDo) Find() ([]*pasture.PastureFeedPb, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureFeedPb), err
}

func (p pastureFeedPbDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureFeedPb, err error) {
	buf := make([]*pasture.PastureFeedPb, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureFeedPbDo) FindInBatches(result *[]*pasture.PastureFeedPb, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureFeedPbDo) Attrs(attrs ...field.AssignExpr) IPastureFeedPbDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureFeedPbDo) Assign(attrs ...field.AssignExpr) IPastureFeedPbDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureFeedPbDo) Joins(fields ...field.RelationField) IPastureFeedPbDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureFeedPbDo) Preload(fields ...field.RelationField) IPastureFeedPbDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureFeedPbDo) FirstOrInit() (*pasture.PastureFeedPb, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedPb), nil
	}
}

func (p pastureFeedPbDo) FirstOrCreate() (*pasture.PastureFeedPb, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedPb), nil
	}
}

func (p pastureFeedPbDo) FindByPage(offset int, limit int) (result []*pasture.PastureFeedPb, count int64, err error) {
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

func (p pastureFeedPbDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureFeedPbDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureFeedPbDo) Delete(models ...*pasture.PastureFeedPb) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureFeedPbDo) withDO(do gen.Dao) *pastureFeedPbDo {
	p.DO = *do.(*gen.DO)
	return p
}