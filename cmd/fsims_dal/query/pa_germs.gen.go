// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPaGerm(db *gorm.DB, opts ...gen.DOOption) paGerm {
	_paGerm := paGerm{}

	_paGerm.paGermDo.UseDB(db, opts...)
	_paGerm.paGermDo.UseModel(&slaughter.PaGerm{})

	tableName := _paGerm.paGermDo.TableName()
	_paGerm.ALL = field.NewAsterisk(tableName)
	_paGerm.ID = field.NewUint(tableName, "id")
	_paGerm.CreatedAt = field.NewTime(tableName, "created_at")
	_paGerm.UpdatedAt = field.NewTime(tableName, "updated_at")
	_paGerm.DeletedAt = field.NewField(tableName, "deleted_at")
	_paGerm.ChiledFreshIndexID = field.NewUint(tableName, "chiled_fresh_index_id")
	_paGerm.PaGerm1 = field.NewFloat32(tableName, "pa_germ1")
	_paGerm.PaGerm2 = field.NewFloat32(tableName, "pa_germ2")
	_paGerm.PaGerm3 = field.NewFloat32(tableName, "pa_germ3")
	_paGerm.PaGerm4 = field.NewFloat32(tableName, "pa_germ4")
	_paGerm.PaGerm5 = field.NewFloat32(tableName, "pa_germ5")

	_paGerm.fillFieldMap()

	return _paGerm
}

type paGerm struct {
	paGermDo paGermDo

	ALL                field.Asterisk
	ID                 field.Uint
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field
	ChiledFreshIndexID field.Uint
	PaGerm1            field.Float32
	PaGerm2            field.Float32
	PaGerm3            field.Float32
	PaGerm4            field.Float32
	PaGerm5            field.Float32

	fieldMap map[string]field.Expr
}

func (p paGerm) Table(newTableName string) *paGerm {
	p.paGermDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p paGerm) As(alias string) *paGerm {
	p.paGermDo.DO = *(p.paGermDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *paGerm) updateTableName(table string) *paGerm {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.ChiledFreshIndexID = field.NewUint(table, "chiled_fresh_index_id")
	p.PaGerm1 = field.NewFloat32(table, "pa_germ1")
	p.PaGerm2 = field.NewFloat32(table, "pa_germ2")
	p.PaGerm3 = field.NewFloat32(table, "pa_germ3")
	p.PaGerm4 = field.NewFloat32(table, "pa_germ4")
	p.PaGerm5 = field.NewFloat32(table, "pa_germ5")

	p.fillFieldMap()

	return p
}

func (p *paGerm) WithContext(ctx context.Context) IPaGermDo { return p.paGermDo.WithContext(ctx) }

func (p paGerm) TableName() string { return p.paGermDo.TableName() }

func (p paGerm) Alias() string { return p.paGermDo.Alias() }

func (p paGerm) Columns(cols ...field.Expr) gen.Columns { return p.paGermDo.Columns(cols...) }

func (p *paGerm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *paGerm) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["chiled_fresh_index_id"] = p.ChiledFreshIndexID
	p.fieldMap["pa_germ1"] = p.PaGerm1
	p.fieldMap["pa_germ2"] = p.PaGerm2
	p.fieldMap["pa_germ3"] = p.PaGerm3
	p.fieldMap["pa_germ4"] = p.PaGerm4
	p.fieldMap["pa_germ5"] = p.PaGerm5
}

func (p paGerm) clone(db *gorm.DB) paGerm {
	p.paGermDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p paGerm) replaceDB(db *gorm.DB) paGerm {
	p.paGermDo.ReplaceDB(db)
	return p
}

type paGermDo struct{ gen.DO }

type IPaGermDo interface {
	gen.SubQuery
	Debug() IPaGermDo
	WithContext(ctx context.Context) IPaGermDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPaGermDo
	WriteDB() IPaGermDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPaGermDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPaGermDo
	Not(conds ...gen.Condition) IPaGermDo
	Or(conds ...gen.Condition) IPaGermDo
	Select(conds ...field.Expr) IPaGermDo
	Where(conds ...gen.Condition) IPaGermDo
	Order(conds ...field.Expr) IPaGermDo
	Distinct(cols ...field.Expr) IPaGermDo
	Omit(cols ...field.Expr) IPaGermDo
	Join(table schema.Tabler, on ...field.Expr) IPaGermDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPaGermDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPaGermDo
	Group(cols ...field.Expr) IPaGermDo
	Having(conds ...gen.Condition) IPaGermDo
	Limit(limit int) IPaGermDo
	Offset(offset int) IPaGermDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPaGermDo
	Unscoped() IPaGermDo
	Create(values ...*slaughter.PaGerm) error
	CreateInBatches(values []*slaughter.PaGerm, batchSize int) error
	Save(values ...*slaughter.PaGerm) error
	First() (*slaughter.PaGerm, error)
	Take() (*slaughter.PaGerm, error)
	Last() (*slaughter.PaGerm, error)
	Find() ([]*slaughter.PaGerm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.PaGerm, err error)
	FindInBatches(result *[]*slaughter.PaGerm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.PaGerm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPaGermDo
	Assign(attrs ...field.AssignExpr) IPaGermDo
	Joins(fields ...field.RelationField) IPaGermDo
	Preload(fields ...field.RelationField) IPaGermDo
	FirstOrInit() (*slaughter.PaGerm, error)
	FirstOrCreate() (*slaughter.PaGerm, error)
	FindByPage(offset int, limit int) (result []*slaughter.PaGerm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPaGermDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p paGermDo) Debug() IPaGermDo {
	return p.withDO(p.DO.Debug())
}

func (p paGermDo) WithContext(ctx context.Context) IPaGermDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paGermDo) ReadDB() IPaGermDo {
	return p.Clauses(dbresolver.Read)
}

func (p paGermDo) WriteDB() IPaGermDo {
	return p.Clauses(dbresolver.Write)
}

func (p paGermDo) Session(config *gorm.Session) IPaGermDo {
	return p.withDO(p.DO.Session(config))
}

func (p paGermDo) Clauses(conds ...clause.Expression) IPaGermDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paGermDo) Returning(value interface{}, columns ...string) IPaGermDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paGermDo) Not(conds ...gen.Condition) IPaGermDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paGermDo) Or(conds ...gen.Condition) IPaGermDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paGermDo) Select(conds ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paGermDo) Where(conds ...gen.Condition) IPaGermDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paGermDo) Order(conds ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paGermDo) Distinct(cols ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paGermDo) Omit(cols ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paGermDo) Join(table schema.Tabler, on ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paGermDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paGermDo) RightJoin(table schema.Tabler, on ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paGermDo) Group(cols ...field.Expr) IPaGermDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paGermDo) Having(conds ...gen.Condition) IPaGermDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paGermDo) Limit(limit int) IPaGermDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paGermDo) Offset(offset int) IPaGermDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paGermDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPaGermDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paGermDo) Unscoped() IPaGermDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paGermDo) Create(values ...*slaughter.PaGerm) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paGermDo) CreateInBatches(values []*slaughter.PaGerm, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paGermDo) Save(values ...*slaughter.PaGerm) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paGermDo) First() (*slaughter.PaGerm, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PaGerm), nil
	}
}

func (p paGermDo) Take() (*slaughter.PaGerm, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PaGerm), nil
	}
}

func (p paGermDo) Last() (*slaughter.PaGerm, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PaGerm), nil
	}
}

func (p paGermDo) Find() ([]*slaughter.PaGerm, error) {
	result, err := p.DO.Find()
	return result.([]*slaughter.PaGerm), err
}

func (p paGermDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.PaGerm, err error) {
	buf := make([]*slaughter.PaGerm, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paGermDo) FindInBatches(result *[]*slaughter.PaGerm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paGermDo) Attrs(attrs ...field.AssignExpr) IPaGermDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paGermDo) Assign(attrs ...field.AssignExpr) IPaGermDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paGermDo) Joins(fields ...field.RelationField) IPaGermDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paGermDo) Preload(fields ...field.RelationField) IPaGermDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paGermDo) FirstOrInit() (*slaughter.PaGerm, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PaGerm), nil
	}
}

func (p paGermDo) FirstOrCreate() (*slaughter.PaGerm, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PaGerm), nil
	}
}

func (p paGermDo) FindByPage(offset int, limit int) (result []*slaughter.PaGerm, count int64, err error) {
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

func (p paGermDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paGermDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paGermDo) Delete(models ...*slaughter.PaGerm) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paGermDo) withDO(do gen.Dao) *paGermDo {
	p.DO = *do.(*gen.DO)
	return p
}