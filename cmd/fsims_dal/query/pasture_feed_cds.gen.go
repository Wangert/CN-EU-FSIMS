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

func newPastureFeedCd(db *gorm.DB, opts ...gen.DOOption) pastureFeedCd {
	_pastureFeedCd := pastureFeedCd{}

	_pastureFeedCd.pastureFeedCdDo.UseDB(db, opts...)
	_pastureFeedCd.pastureFeedCdDo.UseModel(&pasture.PastureFeedCd{})

	tableName := _pastureFeedCd.pastureFeedCdDo.TableName()
	_pastureFeedCd.ALL = field.NewAsterisk(tableName)
	_pastureFeedCd.ID = field.NewUint(tableName, "id")
	_pastureFeedCd.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureFeedCd.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureFeedCd.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureFeedCd.PastureFeedHeavyMetalID = field.NewUint(tableName, "pasture_feed_heavy_metal_id")
	_pastureFeedCd.Cd1 = field.NewFloat64(tableName, "cd1")
	_pastureFeedCd.Cd2 = field.NewFloat64(tableName, "cd2")
	_pastureFeedCd.Cd3 = field.NewFloat64(tableName, "cd3")
	_pastureFeedCd.Cd4 = field.NewFloat64(tableName, "cd4")
	_pastureFeedCd.Cd5 = field.NewFloat64(tableName, "cd5")
	_pastureFeedCd.Cd6 = field.NewFloat64(tableName, "cd6")
	_pastureFeedCd.Cd7 = field.NewFloat64(tableName, "cd7")

	_pastureFeedCd.fillFieldMap()

	return _pastureFeedCd
}

type pastureFeedCd struct {
	pastureFeedCdDo pastureFeedCdDo

	ALL                     field.Asterisk
	ID                      field.Uint
	CreatedAt               field.Time
	UpdatedAt               field.Time
	DeletedAt               field.Field
	PastureFeedHeavyMetalID field.Uint
	Cd1                     field.Float64
	Cd2                     field.Float64
	Cd3                     field.Float64
	Cd4                     field.Float64
	Cd5                     field.Float64
	Cd6                     field.Float64
	Cd7                     field.Float64

	fieldMap map[string]field.Expr
}

func (p pastureFeedCd) Table(newTableName string) *pastureFeedCd {
	p.pastureFeedCdDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureFeedCd) As(alias string) *pastureFeedCd {
	p.pastureFeedCdDo.DO = *(p.pastureFeedCdDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureFeedCd) updateTableName(table string) *pastureFeedCd {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.PastureFeedHeavyMetalID = field.NewUint(table, "pasture_feed_heavy_metal_id")
	p.Cd1 = field.NewFloat64(table, "cd1")
	p.Cd2 = field.NewFloat64(table, "cd2")
	p.Cd3 = field.NewFloat64(table, "cd3")
	p.Cd4 = field.NewFloat64(table, "cd4")
	p.Cd5 = field.NewFloat64(table, "cd5")
	p.Cd6 = field.NewFloat64(table, "cd6")
	p.Cd7 = field.NewFloat64(table, "cd7")

	p.fillFieldMap()

	return p
}

func (p *pastureFeedCd) WithContext(ctx context.Context) IPastureFeedCdDo {
	return p.pastureFeedCdDo.WithContext(ctx)
}

func (p pastureFeedCd) TableName() string { return p.pastureFeedCdDo.TableName() }

func (p pastureFeedCd) Alias() string { return p.pastureFeedCdDo.Alias() }

func (p pastureFeedCd) Columns(cols ...field.Expr) gen.Columns {
	return p.pastureFeedCdDo.Columns(cols...)
}

func (p *pastureFeedCd) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureFeedCd) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["pasture_feed_heavy_metal_id"] = p.PastureFeedHeavyMetalID
	p.fieldMap["cd1"] = p.Cd1
	p.fieldMap["cd2"] = p.Cd2
	p.fieldMap["cd3"] = p.Cd3
	p.fieldMap["cd4"] = p.Cd4
	p.fieldMap["cd5"] = p.Cd5
	p.fieldMap["cd6"] = p.Cd6
	p.fieldMap["cd7"] = p.Cd7
}

func (p pastureFeedCd) clone(db *gorm.DB) pastureFeedCd {
	p.pastureFeedCdDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureFeedCd) replaceDB(db *gorm.DB) pastureFeedCd {
	p.pastureFeedCdDo.ReplaceDB(db)
	return p
}

type pastureFeedCdDo struct{ gen.DO }

type IPastureFeedCdDo interface {
	gen.SubQuery
	Debug() IPastureFeedCdDo
	WithContext(ctx context.Context) IPastureFeedCdDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureFeedCdDo
	WriteDB() IPastureFeedCdDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureFeedCdDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureFeedCdDo
	Not(conds ...gen.Condition) IPastureFeedCdDo
	Or(conds ...gen.Condition) IPastureFeedCdDo
	Select(conds ...field.Expr) IPastureFeedCdDo
	Where(conds ...gen.Condition) IPastureFeedCdDo
	Order(conds ...field.Expr) IPastureFeedCdDo
	Distinct(cols ...field.Expr) IPastureFeedCdDo
	Omit(cols ...field.Expr) IPastureFeedCdDo
	Join(table schema.Tabler, on ...field.Expr) IPastureFeedCdDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureFeedCdDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureFeedCdDo
	Group(cols ...field.Expr) IPastureFeedCdDo
	Having(conds ...gen.Condition) IPastureFeedCdDo
	Limit(limit int) IPastureFeedCdDo
	Offset(offset int) IPastureFeedCdDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureFeedCdDo
	Unscoped() IPastureFeedCdDo
	Create(values ...*pasture.PastureFeedCd) error
	CreateInBatches(values []*pasture.PastureFeedCd, batchSize int) error
	Save(values ...*pasture.PastureFeedCd) error
	First() (*pasture.PastureFeedCd, error)
	Take() (*pasture.PastureFeedCd, error)
	Last() (*pasture.PastureFeedCd, error)
	Find() ([]*pasture.PastureFeedCd, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureFeedCd, err error)
	FindInBatches(result *[]*pasture.PastureFeedCd, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureFeedCd) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureFeedCdDo
	Assign(attrs ...field.AssignExpr) IPastureFeedCdDo
	Joins(fields ...field.RelationField) IPastureFeedCdDo
	Preload(fields ...field.RelationField) IPastureFeedCdDo
	FirstOrInit() (*pasture.PastureFeedCd, error)
	FirstOrCreate() (*pasture.PastureFeedCd, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureFeedCd, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureFeedCdDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureFeedCdDo) Debug() IPastureFeedCdDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureFeedCdDo) WithContext(ctx context.Context) IPastureFeedCdDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureFeedCdDo) ReadDB() IPastureFeedCdDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureFeedCdDo) WriteDB() IPastureFeedCdDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureFeedCdDo) Session(config *gorm.Session) IPastureFeedCdDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureFeedCdDo) Clauses(conds ...clause.Expression) IPastureFeedCdDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureFeedCdDo) Returning(value interface{}, columns ...string) IPastureFeedCdDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureFeedCdDo) Not(conds ...gen.Condition) IPastureFeedCdDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureFeedCdDo) Or(conds ...gen.Condition) IPastureFeedCdDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureFeedCdDo) Select(conds ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureFeedCdDo) Where(conds ...gen.Condition) IPastureFeedCdDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureFeedCdDo) Order(conds ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureFeedCdDo) Distinct(cols ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureFeedCdDo) Omit(cols ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureFeedCdDo) Join(table schema.Tabler, on ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureFeedCdDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureFeedCdDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureFeedCdDo) Group(cols ...field.Expr) IPastureFeedCdDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureFeedCdDo) Having(conds ...gen.Condition) IPastureFeedCdDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureFeedCdDo) Limit(limit int) IPastureFeedCdDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureFeedCdDo) Offset(offset int) IPastureFeedCdDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureFeedCdDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureFeedCdDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureFeedCdDo) Unscoped() IPastureFeedCdDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureFeedCdDo) Create(values ...*pasture.PastureFeedCd) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureFeedCdDo) CreateInBatches(values []*pasture.PastureFeedCd, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureFeedCdDo) Save(values ...*pasture.PastureFeedCd) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureFeedCdDo) First() (*pasture.PastureFeedCd, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedCd), nil
	}
}

func (p pastureFeedCdDo) Take() (*pasture.PastureFeedCd, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedCd), nil
	}
}

func (p pastureFeedCdDo) Last() (*pasture.PastureFeedCd, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedCd), nil
	}
}

func (p pastureFeedCdDo) Find() ([]*pasture.PastureFeedCd, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureFeedCd), err
}

func (p pastureFeedCdDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureFeedCd, err error) {
	buf := make([]*pasture.PastureFeedCd, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureFeedCdDo) FindInBatches(result *[]*pasture.PastureFeedCd, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureFeedCdDo) Attrs(attrs ...field.AssignExpr) IPastureFeedCdDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureFeedCdDo) Assign(attrs ...field.AssignExpr) IPastureFeedCdDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureFeedCdDo) Joins(fields ...field.RelationField) IPastureFeedCdDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureFeedCdDo) Preload(fields ...field.RelationField) IPastureFeedCdDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureFeedCdDo) FirstOrInit() (*pasture.PastureFeedCd, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedCd), nil
	}
}

func (p pastureFeedCdDo) FirstOrCreate() (*pasture.PastureFeedCd, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureFeedCd), nil
	}
}

func (p pastureFeedCdDo) FindByPage(offset int, limit int) (result []*pasture.PastureFeedCd, count int64, err error) {
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

func (p pastureFeedCdDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureFeedCdDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureFeedCdDo) Delete(models ...*pasture.PastureFeedCd) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureFeedCdDo) withDO(do gen.Dao) *pastureFeedCdDo {
	p.DO = *do.(*gen.DO)
	return p
}