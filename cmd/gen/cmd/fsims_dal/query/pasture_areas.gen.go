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

func newPastureArea(db *gorm.DB, opts ...gen.DOOption) pastureArea {
	_pastureArea := pastureArea{}

	_pastureArea.pastureAreaDo.UseDB(db, opts...)
	_pastureArea.pastureAreaDo.UseModel(&pasture.PastureArea{})

	tableName := _pastureArea.pastureAreaDo.TableName()
	_pastureArea.ALL = field.NewAsterisk(tableName)
	_pastureArea.ID = field.NewUint(tableName, "id")
	_pastureArea.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureArea.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureArea.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureArea.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_pastureArea.HouseNumber = field.NewString(tableName, "house_number")
	_pastureArea.CattleFarm1 = field.NewFloat32(tableName, "cattle_farm1")
	_pastureArea.CattleFarm2 = field.NewFloat32(tableName, "cattle_farm2")
	_pastureArea.CattleFarm3 = field.NewFloat32(tableName, "cattle_farm3")
	_pastureArea.CattleFarm4 = field.NewFloat32(tableName, "cattle_farm4")
	_pastureArea.CattleFarm5 = field.NewFloat32(tableName, "cattle_farm5")
	_pastureArea.CattleFarm6 = field.NewFloat32(tableName, "cattle_farm6")
	_pastureArea.CattleFarm7 = field.NewFloat32(tableName, "cattle_farm7")
	_pastureArea.CattleFarm8 = field.NewFloat32(tableName, "cattle_farm8")
	_pastureArea.CattleFarm9 = field.NewFloat32(tableName, "cattle_farm9")
	_pastureArea.CattleFarm10 = field.NewFloat32(tableName, "cattle_farm10")
	_pastureArea.CattleFarm11 = field.NewFloat32(tableName, "cattle_farm11")
	_pastureArea.CattleFarm12 = field.NewFloat32(tableName, "cattle_farm12")

	_pastureArea.fillFieldMap()

	return _pastureArea
}

type pastureArea struct {
	pastureAreaDo pastureAreaDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	TimeRecordAt field.Time
	HouseNumber  field.String
	CattleFarm1  field.Float32
	CattleFarm2  field.Float32
	CattleFarm3  field.Float32
	CattleFarm4  field.Float32
	CattleFarm5  field.Float32
	CattleFarm6  field.Float32
	CattleFarm7  field.Float32
	CattleFarm8  field.Float32
	CattleFarm9  field.Float32
	CattleFarm10 field.Float32
	CattleFarm11 field.Float32
	CattleFarm12 field.Float32

	fieldMap map[string]field.Expr
}

func (p pastureArea) Table(newTableName string) *pastureArea {
	p.pastureAreaDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureArea) As(alias string) *pastureArea {
	p.pastureAreaDo.DO = *(p.pastureAreaDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureArea) updateTableName(table string) *pastureArea {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.TimeRecordAt = field.NewTime(table, "time_record_at")
	p.HouseNumber = field.NewString(table, "house_number")
	p.CattleFarm1 = field.NewFloat32(table, "cattle_farm1")
	p.CattleFarm2 = field.NewFloat32(table, "cattle_farm2")
	p.CattleFarm3 = field.NewFloat32(table, "cattle_farm3")
	p.CattleFarm4 = field.NewFloat32(table, "cattle_farm4")
	p.CattleFarm5 = field.NewFloat32(table, "cattle_farm5")
	p.CattleFarm6 = field.NewFloat32(table, "cattle_farm6")
	p.CattleFarm7 = field.NewFloat32(table, "cattle_farm7")
	p.CattleFarm8 = field.NewFloat32(table, "cattle_farm8")
	p.CattleFarm9 = field.NewFloat32(table, "cattle_farm9")
	p.CattleFarm10 = field.NewFloat32(table, "cattle_farm10")
	p.CattleFarm11 = field.NewFloat32(table, "cattle_farm11")
	p.CattleFarm12 = field.NewFloat32(table, "cattle_farm12")

	p.fillFieldMap()

	return p
}

func (p *pastureArea) WithContext(ctx context.Context) IPastureAreaDo {
	return p.pastureAreaDo.WithContext(ctx)
}

func (p pastureArea) TableName() string { return p.pastureAreaDo.TableName() }

func (p pastureArea) Alias() string { return p.pastureAreaDo.Alias() }

func (p pastureArea) Columns(cols ...field.Expr) gen.Columns { return p.pastureAreaDo.Columns(cols...) }

func (p *pastureArea) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureArea) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 18)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["time_record_at"] = p.TimeRecordAt
	p.fieldMap["house_number"] = p.HouseNumber
	p.fieldMap["cattle_farm1"] = p.CattleFarm1
	p.fieldMap["cattle_farm2"] = p.CattleFarm2
	p.fieldMap["cattle_farm3"] = p.CattleFarm3
	p.fieldMap["cattle_farm4"] = p.CattleFarm4
	p.fieldMap["cattle_farm5"] = p.CattleFarm5
	p.fieldMap["cattle_farm6"] = p.CattleFarm6
	p.fieldMap["cattle_farm7"] = p.CattleFarm7
	p.fieldMap["cattle_farm8"] = p.CattleFarm8
	p.fieldMap["cattle_farm9"] = p.CattleFarm9
	p.fieldMap["cattle_farm10"] = p.CattleFarm10
	p.fieldMap["cattle_farm11"] = p.CattleFarm11
	p.fieldMap["cattle_farm12"] = p.CattleFarm12
}

func (p pastureArea) clone(db *gorm.DB) pastureArea {
	p.pastureAreaDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureArea) replaceDB(db *gorm.DB) pastureArea {
	p.pastureAreaDo.ReplaceDB(db)
	return p
}

type pastureAreaDo struct{ gen.DO }

type IPastureAreaDo interface {
	gen.SubQuery
	Debug() IPastureAreaDo
	WithContext(ctx context.Context) IPastureAreaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureAreaDo
	WriteDB() IPastureAreaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureAreaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureAreaDo
	Not(conds ...gen.Condition) IPastureAreaDo
	Or(conds ...gen.Condition) IPastureAreaDo
	Select(conds ...field.Expr) IPastureAreaDo
	Where(conds ...gen.Condition) IPastureAreaDo
	Order(conds ...field.Expr) IPastureAreaDo
	Distinct(cols ...field.Expr) IPastureAreaDo
	Omit(cols ...field.Expr) IPastureAreaDo
	Join(table schema.Tabler, on ...field.Expr) IPastureAreaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureAreaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureAreaDo
	Group(cols ...field.Expr) IPastureAreaDo
	Having(conds ...gen.Condition) IPastureAreaDo
	Limit(limit int) IPastureAreaDo
	Offset(offset int) IPastureAreaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureAreaDo
	Unscoped() IPastureAreaDo
	Create(values ...*pasture.PastureArea) error
	CreateInBatches(values []*pasture.PastureArea, batchSize int) error
	Save(values ...*pasture.PastureArea) error
	First() (*pasture.PastureArea, error)
	Take() (*pasture.PastureArea, error)
	Last() (*pasture.PastureArea, error)
	Find() ([]*pasture.PastureArea, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureArea, err error)
	FindInBatches(result *[]*pasture.PastureArea, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureArea) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureAreaDo
	Assign(attrs ...field.AssignExpr) IPastureAreaDo
	Joins(fields ...field.RelationField) IPastureAreaDo
	Preload(fields ...field.RelationField) IPastureAreaDo
	FirstOrInit() (*pasture.PastureArea, error)
	FirstOrCreate() (*pasture.PastureArea, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureArea, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureAreaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureAreaDo) Debug() IPastureAreaDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureAreaDo) WithContext(ctx context.Context) IPastureAreaDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureAreaDo) ReadDB() IPastureAreaDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureAreaDo) WriteDB() IPastureAreaDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureAreaDo) Session(config *gorm.Session) IPastureAreaDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureAreaDo) Clauses(conds ...clause.Expression) IPastureAreaDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureAreaDo) Returning(value interface{}, columns ...string) IPastureAreaDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureAreaDo) Not(conds ...gen.Condition) IPastureAreaDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureAreaDo) Or(conds ...gen.Condition) IPastureAreaDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureAreaDo) Select(conds ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureAreaDo) Where(conds ...gen.Condition) IPastureAreaDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureAreaDo) Order(conds ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureAreaDo) Distinct(cols ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureAreaDo) Omit(cols ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureAreaDo) Join(table schema.Tabler, on ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureAreaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureAreaDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureAreaDo) Group(cols ...field.Expr) IPastureAreaDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureAreaDo) Having(conds ...gen.Condition) IPastureAreaDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureAreaDo) Limit(limit int) IPastureAreaDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureAreaDo) Offset(offset int) IPastureAreaDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureAreaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureAreaDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureAreaDo) Unscoped() IPastureAreaDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureAreaDo) Create(values ...*pasture.PastureArea) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureAreaDo) CreateInBatches(values []*pasture.PastureArea, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureAreaDo) Save(values ...*pasture.PastureArea) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureAreaDo) First() (*pasture.PastureArea, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureArea), nil
	}
}

func (p pastureAreaDo) Take() (*pasture.PastureArea, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureArea), nil
	}
}

func (p pastureAreaDo) Last() (*pasture.PastureArea, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureArea), nil
	}
}

func (p pastureAreaDo) Find() ([]*pasture.PastureArea, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureArea), err
}

func (p pastureAreaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureArea, err error) {
	buf := make([]*pasture.PastureArea, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureAreaDo) FindInBatches(result *[]*pasture.PastureArea, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureAreaDo) Attrs(attrs ...field.AssignExpr) IPastureAreaDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureAreaDo) Assign(attrs ...field.AssignExpr) IPastureAreaDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureAreaDo) Joins(fields ...field.RelationField) IPastureAreaDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureAreaDo) Preload(fields ...field.RelationField) IPastureAreaDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureAreaDo) FirstOrInit() (*pasture.PastureArea, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureArea), nil
	}
}

func (p pastureAreaDo) FirstOrCreate() (*pasture.PastureArea, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureArea), nil
	}
}

func (p pastureAreaDo) FindByPage(offset int, limit int) (result []*pasture.PastureArea, count int64, err error) {
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

func (p pastureAreaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureAreaDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureAreaDo) Delete(models ...*pasture.PastureArea) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureAreaDo) withDO(do gen.Dao) *pastureAreaDo {
	p.DO = *do.(*gen.DO)
	return p
}