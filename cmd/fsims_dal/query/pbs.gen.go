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

func newPb(db *gorm.DB, opts ...gen.DOOption) pb {
	_pb := pb{}

	_pb.pbDo.UseDB(db, opts...)
	_pb.pbDo.UseModel(&pasture.Pb{})

	tableName := _pb.pbDo.TableName()
	_pb.ALL = field.NewAsterisk(tableName)
	_pb.ID = field.NewUint(tableName, "id")
	_pb.CreatedAt = field.NewTime(tableName, "created_at")
	_pb.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pb.DeletedAt = field.NewField(tableName, "deleted_at")
	_pb.HeavyMetalID = field.NewUint(tableName, "heavy_metal_id")
	_pb.Pb1 = field.NewFloat64(tableName, "pb1")
	_pb.Pb2 = field.NewFloat64(tableName, "pb2")
	_pb.Pb3 = field.NewFloat64(tableName, "pb3")
	_pb.Pb4 = field.NewFloat64(tableName, "pb4")
	_pb.Pb5 = field.NewFloat64(tableName, "pb5")
	_pb.Pb6 = field.NewFloat64(tableName, "pb6")
	_pb.Pb7 = field.NewFloat64(tableName, "pb7")

	_pb.fillFieldMap()

	return _pb
}

type pb struct {
	pbDo pbDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HeavyMetalID field.Uint
	Pb1          field.Float64
	Pb2          field.Float64
	Pb3          field.Float64
	Pb4          field.Float64
	Pb5          field.Float64
	Pb6          field.Float64
	Pb7          field.Float64

	fieldMap map[string]field.Expr
}

func (p pb) Table(newTableName string) *pb {
	p.pbDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pb) As(alias string) *pb {
	p.pbDo.DO = *(p.pbDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pb) updateTableName(table string) *pb {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.HeavyMetalID = field.NewUint(table, "heavy_metal_id")
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

func (p *pb) WithContext(ctx context.Context) IPbDo { return p.pbDo.WithContext(ctx) }

func (p pb) TableName() string { return p.pbDo.TableName() }

func (p pb) Alias() string { return p.pbDo.Alias() }

func (p pb) Columns(cols ...field.Expr) gen.Columns { return p.pbDo.Columns(cols...) }

func (p *pb) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pb) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["heavy_metal_id"] = p.HeavyMetalID
	p.fieldMap["pb1"] = p.Pb1
	p.fieldMap["pb2"] = p.Pb2
	p.fieldMap["pb3"] = p.Pb3
	p.fieldMap["pb4"] = p.Pb4
	p.fieldMap["pb5"] = p.Pb5
	p.fieldMap["pb6"] = p.Pb6
	p.fieldMap["pb7"] = p.Pb7
}

func (p pb) clone(db *gorm.DB) pb {
	p.pbDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pb) replaceDB(db *gorm.DB) pb {
	p.pbDo.ReplaceDB(db)
	return p
}

type pbDo struct{ gen.DO }

type IPbDo interface {
	gen.SubQuery
	Debug() IPbDo
	WithContext(ctx context.Context) IPbDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPbDo
	WriteDB() IPbDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPbDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPbDo
	Not(conds ...gen.Condition) IPbDo
	Or(conds ...gen.Condition) IPbDo
	Select(conds ...field.Expr) IPbDo
	Where(conds ...gen.Condition) IPbDo
	Order(conds ...field.Expr) IPbDo
	Distinct(cols ...field.Expr) IPbDo
	Omit(cols ...field.Expr) IPbDo
	Join(table schema.Tabler, on ...field.Expr) IPbDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPbDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPbDo
	Group(cols ...field.Expr) IPbDo
	Having(conds ...gen.Condition) IPbDo
	Limit(limit int) IPbDo
	Offset(offset int) IPbDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPbDo
	Unscoped() IPbDo
	Create(values ...*pasture.Pb) error
	CreateInBatches(values []*pasture.Pb, batchSize int) error
	Save(values ...*pasture.Pb) error
	First() (*pasture.Pb, error)
	Take() (*pasture.Pb, error)
	Last() (*pasture.Pb, error)
	Find() ([]*pasture.Pb, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.Pb, err error)
	FindInBatches(result *[]*pasture.Pb, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.Pb) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPbDo
	Assign(attrs ...field.AssignExpr) IPbDo
	Joins(fields ...field.RelationField) IPbDo
	Preload(fields ...field.RelationField) IPbDo
	FirstOrInit() (*pasture.Pb, error)
	FirstOrCreate() (*pasture.Pb, error)
	FindByPage(offset int, limit int) (result []*pasture.Pb, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPbDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pbDo) Debug() IPbDo {
	return p.withDO(p.DO.Debug())
}

func (p pbDo) WithContext(ctx context.Context) IPbDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pbDo) ReadDB() IPbDo {
	return p.Clauses(dbresolver.Read)
}

func (p pbDo) WriteDB() IPbDo {
	return p.Clauses(dbresolver.Write)
}

func (p pbDo) Session(config *gorm.Session) IPbDo {
	return p.withDO(p.DO.Session(config))
}

func (p pbDo) Clauses(conds ...clause.Expression) IPbDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pbDo) Returning(value interface{}, columns ...string) IPbDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pbDo) Not(conds ...gen.Condition) IPbDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pbDo) Or(conds ...gen.Condition) IPbDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pbDo) Select(conds ...field.Expr) IPbDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pbDo) Where(conds ...gen.Condition) IPbDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pbDo) Order(conds ...field.Expr) IPbDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pbDo) Distinct(cols ...field.Expr) IPbDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pbDo) Omit(cols ...field.Expr) IPbDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pbDo) Join(table schema.Tabler, on ...field.Expr) IPbDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pbDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPbDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pbDo) RightJoin(table schema.Tabler, on ...field.Expr) IPbDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pbDo) Group(cols ...field.Expr) IPbDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pbDo) Having(conds ...gen.Condition) IPbDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pbDo) Limit(limit int) IPbDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pbDo) Offset(offset int) IPbDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pbDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPbDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pbDo) Unscoped() IPbDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pbDo) Create(values ...*pasture.Pb) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pbDo) CreateInBatches(values []*pasture.Pb, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pbDo) Save(values ...*pasture.Pb) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pbDo) First() (*pasture.Pb, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Pb), nil
	}
}

func (p pbDo) Take() (*pasture.Pb, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Pb), nil
	}
}

func (p pbDo) Last() (*pasture.Pb, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Pb), nil
	}
}

func (p pbDo) Find() ([]*pasture.Pb, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.Pb), err
}

func (p pbDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.Pb, err error) {
	buf := make([]*pasture.Pb, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pbDo) FindInBatches(result *[]*pasture.Pb, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pbDo) Attrs(attrs ...field.AssignExpr) IPbDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pbDo) Assign(attrs ...field.AssignExpr) IPbDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pbDo) Joins(fields ...field.RelationField) IPbDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pbDo) Preload(fields ...field.RelationField) IPbDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pbDo) FirstOrInit() (*pasture.Pb, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Pb), nil
	}
}

func (p pbDo) FirstOrCreate() (*pasture.Pb, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.Pb), nil
	}
}

func (p pbDo) FindByPage(offset int, limit int) (result []*pasture.Pb, count int64, err error) {
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

func (p pbDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pbDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pbDo) Delete(models ...*pasture.Pb) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pbDo) withDO(do gen.Dao) *pbDo {
	p.DO = *do.(*gen.DO)
	return p
}