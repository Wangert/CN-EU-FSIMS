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

func newPastureBuffer(db *gorm.DB, opts ...gen.DOOption) pastureBuffer {
	_pastureBuffer := pastureBuffer{}

	_pastureBuffer.pastureBufferDo.UseDB(db, opts...)
	_pastureBuffer.pastureBufferDo.UseModel(&pasture.PastureBuffer{})

	tableName := _pastureBuffer.pastureBufferDo.TableName()
	_pastureBuffer.ALL = field.NewAsterisk(tableName)
	_pastureBuffer.ID = field.NewUint(tableName, "id")
	_pastureBuffer.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureBuffer.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureBuffer.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureBuffer.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_pastureBuffer.HouseNumber = field.NewString(tableName, "house_number")
	_pastureBuffer.Buffer1 = field.NewFloat32(tableName, "buffer1")
	_pastureBuffer.Buffer2 = field.NewFloat32(tableName, "buffer2")
	_pastureBuffer.Buffer3 = field.NewFloat32(tableName, "buffer3")
	_pastureBuffer.Buffer4 = field.NewFloat32(tableName, "buffer4")
	_pastureBuffer.Buffer5 = field.NewFloat32(tableName, "buffer5")
	_pastureBuffer.Buffer6 = field.NewFloat32(tableName, "buffer6")
	_pastureBuffer.Buffer7 = field.NewFloat32(tableName, "buffer7")
	_pastureBuffer.Buffer8 = field.NewFloat32(tableName, "buffer8")
	_pastureBuffer.Buffer9 = field.NewFloat32(tableName, "buffer9")
	_pastureBuffer.Buffer10 = field.NewFloat32(tableName, "buffer10")
	_pastureBuffer.Buffer11 = field.NewFloat32(tableName, "buffer11")
	_pastureBuffer.Buffer12 = field.NewFloat32(tableName, "buffer12")

	_pastureBuffer.fillFieldMap()

	return _pastureBuffer
}

type pastureBuffer struct {
	pastureBufferDo pastureBufferDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	TimeRecordAt field.Time
	HouseNumber  field.String
	Buffer1      field.Float32
	Buffer2      field.Float32
	Buffer3      field.Float32
	Buffer4      field.Float32
	Buffer5      field.Float32
	Buffer6      field.Float32
	Buffer7      field.Float32
	Buffer8      field.Float32
	Buffer9      field.Float32
	Buffer10     field.Float32
	Buffer11     field.Float32
	Buffer12     field.Float32

	fieldMap map[string]field.Expr
}

func (p pastureBuffer) Table(newTableName string) *pastureBuffer {
	p.pastureBufferDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureBuffer) As(alias string) *pastureBuffer {
	p.pastureBufferDo.DO = *(p.pastureBufferDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureBuffer) updateTableName(table string) *pastureBuffer {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.TimeRecordAt = field.NewTime(table, "time_record_at")
	p.HouseNumber = field.NewString(table, "house_number")
	p.Buffer1 = field.NewFloat32(table, "buffer1")
	p.Buffer2 = field.NewFloat32(table, "buffer2")
	p.Buffer3 = field.NewFloat32(table, "buffer3")
	p.Buffer4 = field.NewFloat32(table, "buffer4")
	p.Buffer5 = field.NewFloat32(table, "buffer5")
	p.Buffer6 = field.NewFloat32(table, "buffer6")
	p.Buffer7 = field.NewFloat32(table, "buffer7")
	p.Buffer8 = field.NewFloat32(table, "buffer8")
	p.Buffer9 = field.NewFloat32(table, "buffer9")
	p.Buffer10 = field.NewFloat32(table, "buffer10")
	p.Buffer11 = field.NewFloat32(table, "buffer11")
	p.Buffer12 = field.NewFloat32(table, "buffer12")

	p.fillFieldMap()

	return p
}

func (p *pastureBuffer) WithContext(ctx context.Context) IPastureBufferDo {
	return p.pastureBufferDo.WithContext(ctx)
}

func (p pastureBuffer) TableName() string { return p.pastureBufferDo.TableName() }

func (p pastureBuffer) Alias() string { return p.pastureBufferDo.Alias() }

func (p pastureBuffer) Columns(cols ...field.Expr) gen.Columns {
	return p.pastureBufferDo.Columns(cols...)
}

func (p *pastureBuffer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureBuffer) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 18)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["time_record_at"] = p.TimeRecordAt
	p.fieldMap["house_number"] = p.HouseNumber
	p.fieldMap["buffer1"] = p.Buffer1
	p.fieldMap["buffer2"] = p.Buffer2
	p.fieldMap["buffer3"] = p.Buffer3
	p.fieldMap["buffer4"] = p.Buffer4
	p.fieldMap["buffer5"] = p.Buffer5
	p.fieldMap["buffer6"] = p.Buffer6
	p.fieldMap["buffer7"] = p.Buffer7
	p.fieldMap["buffer8"] = p.Buffer8
	p.fieldMap["buffer9"] = p.Buffer9
	p.fieldMap["buffer10"] = p.Buffer10
	p.fieldMap["buffer11"] = p.Buffer11
	p.fieldMap["buffer12"] = p.Buffer12
}

func (p pastureBuffer) clone(db *gorm.DB) pastureBuffer {
	p.pastureBufferDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureBuffer) replaceDB(db *gorm.DB) pastureBuffer {
	p.pastureBufferDo.ReplaceDB(db)
	return p
}

type pastureBufferDo struct{ gen.DO }

type IPastureBufferDo interface {
	gen.SubQuery
	Debug() IPastureBufferDo
	WithContext(ctx context.Context) IPastureBufferDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureBufferDo
	WriteDB() IPastureBufferDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureBufferDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureBufferDo
	Not(conds ...gen.Condition) IPastureBufferDo
	Or(conds ...gen.Condition) IPastureBufferDo
	Select(conds ...field.Expr) IPastureBufferDo
	Where(conds ...gen.Condition) IPastureBufferDo
	Order(conds ...field.Expr) IPastureBufferDo
	Distinct(cols ...field.Expr) IPastureBufferDo
	Omit(cols ...field.Expr) IPastureBufferDo
	Join(table schema.Tabler, on ...field.Expr) IPastureBufferDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureBufferDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureBufferDo
	Group(cols ...field.Expr) IPastureBufferDo
	Having(conds ...gen.Condition) IPastureBufferDo
	Limit(limit int) IPastureBufferDo
	Offset(offset int) IPastureBufferDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureBufferDo
	Unscoped() IPastureBufferDo
	Create(values ...*pasture.PastureBuffer) error
	CreateInBatches(values []*pasture.PastureBuffer, batchSize int) error
	Save(values ...*pasture.PastureBuffer) error
	First() (*pasture.PastureBuffer, error)
	Take() (*pasture.PastureBuffer, error)
	Last() (*pasture.PastureBuffer, error)
	Find() ([]*pasture.PastureBuffer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureBuffer, err error)
	FindInBatches(result *[]*pasture.PastureBuffer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureBuffer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureBufferDo
	Assign(attrs ...field.AssignExpr) IPastureBufferDo
	Joins(fields ...field.RelationField) IPastureBufferDo
	Preload(fields ...field.RelationField) IPastureBufferDo
	FirstOrInit() (*pasture.PastureBuffer, error)
	FirstOrCreate() (*pasture.PastureBuffer, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureBuffer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureBufferDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureBufferDo) Debug() IPastureBufferDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureBufferDo) WithContext(ctx context.Context) IPastureBufferDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureBufferDo) ReadDB() IPastureBufferDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureBufferDo) WriteDB() IPastureBufferDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureBufferDo) Session(config *gorm.Session) IPastureBufferDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureBufferDo) Clauses(conds ...clause.Expression) IPastureBufferDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureBufferDo) Returning(value interface{}, columns ...string) IPastureBufferDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureBufferDo) Not(conds ...gen.Condition) IPastureBufferDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureBufferDo) Or(conds ...gen.Condition) IPastureBufferDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureBufferDo) Select(conds ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureBufferDo) Where(conds ...gen.Condition) IPastureBufferDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureBufferDo) Order(conds ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureBufferDo) Distinct(cols ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureBufferDo) Omit(cols ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureBufferDo) Join(table schema.Tabler, on ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureBufferDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureBufferDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureBufferDo) Group(cols ...field.Expr) IPastureBufferDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureBufferDo) Having(conds ...gen.Condition) IPastureBufferDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureBufferDo) Limit(limit int) IPastureBufferDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureBufferDo) Offset(offset int) IPastureBufferDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureBufferDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureBufferDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureBufferDo) Unscoped() IPastureBufferDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureBufferDo) Create(values ...*pasture.PastureBuffer) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureBufferDo) CreateInBatches(values []*pasture.PastureBuffer, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureBufferDo) Save(values ...*pasture.PastureBuffer) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureBufferDo) First() (*pasture.PastureBuffer, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureBuffer), nil
	}
}

func (p pastureBufferDo) Take() (*pasture.PastureBuffer, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureBuffer), nil
	}
}

func (p pastureBufferDo) Last() (*pasture.PastureBuffer, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureBuffer), nil
	}
}

func (p pastureBufferDo) Find() ([]*pasture.PastureBuffer, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureBuffer), err
}

func (p pastureBufferDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureBuffer, err error) {
	buf := make([]*pasture.PastureBuffer, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureBufferDo) FindInBatches(result *[]*pasture.PastureBuffer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureBufferDo) Attrs(attrs ...field.AssignExpr) IPastureBufferDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureBufferDo) Assign(attrs ...field.AssignExpr) IPastureBufferDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureBufferDo) Joins(fields ...field.RelationField) IPastureBufferDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureBufferDo) Preload(fields ...field.RelationField) IPastureBufferDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureBufferDo) FirstOrInit() (*pasture.PastureBuffer, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureBuffer), nil
	}
}

func (p pastureBufferDo) FirstOrCreate() (*pasture.PastureBuffer, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureBuffer), nil
	}
}

func (p pastureBufferDo) FindByPage(offset int, limit int) (result []*pasture.PastureBuffer, count int64, err error) {
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

func (p pastureBufferDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureBufferDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureBufferDo) Delete(models ...*pasture.PastureBuffer) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureBufferDo) withDO(do gen.Dao) *pastureBufferDo {
	p.DO = *do.(*gen.DO)
	return p
}