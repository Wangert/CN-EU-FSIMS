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

func newAcidShop(db *gorm.DB, opts ...gen.DOOption) acidShop {
	_acidShop := acidShop{}

	_acidShop.acidShopDo.UseDB(db, opts...)
	_acidShop.acidShopDo.UseModel(&slaughter.AcidShop{})

	tableName := _acidShop.acidShopDo.TableName()
	_acidShop.ALL = field.NewAsterisk(tableName)
	_acidShop.ID = field.NewUint(tableName, "id")
	_acidShop.CreatedAt = field.NewTime(tableName, "created_at")
	_acidShop.UpdatedAt = field.NewTime(tableName, "updated_at")
	_acidShop.DeletedAt = field.NewField(tableName, "deleted_at")
	_acidShop.HouseNumber = field.NewString(tableName, "house_number")
	_acidShop.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_acidShop.AcidShop1 = field.NewFloat64(tableName, "acid_shop1")
	_acidShop.AcidShop2 = field.NewFloat64(tableName, "acid_shop2")
	_acidShop.AcidShop3 = field.NewFloat64(tableName, "acid_shop3")
	_acidShop.AcidShop4 = field.NewFloat64(tableName, "acid_shop4")
	_acidShop.AcidShop5 = field.NewFloat64(tableName, "acid_shop5")
	_acidShop.AcidShop6 = field.NewFloat64(tableName, "acid_shop6")
	_acidShop.AcidShop7 = field.NewFloat64(tableName, "acid_shop7")
	_acidShop.AcidShop8 = field.NewFloat64(tableName, "acid_shop8")
	_acidShop.AcidShop9 = field.NewFloat64(tableName, "acid_shop9")
	_acidShop.AcidShop10 = field.NewString(tableName, "acid_shop10")
	_acidShop.AcidShop11 = field.NewString(tableName, "acid_shop11")
	_acidShop.AcidShop12 = field.NewString(tableName, "acid_shop12")
	_acidShop.AcidShop13 = field.NewString(tableName, "acid_shop13")
	_acidShop.AcidShop14 = field.NewString(tableName, "acid_shop14")

	_acidShop.fillFieldMap()

	return _acidShop
}

type acidShop struct {
	acidShopDo acidShopDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HouseNumber  field.String
	TimeRecordAt field.Time
	AcidShop1    field.Float32
	AcidShop2    field.Float32
	AcidShop3    field.Float32
	AcidShop4    field.Float32
	AcidShop5    field.Float32
	AcidShop6    field.Float32
	AcidShop7    field.Float32
	AcidShop8    field.Float32
	AcidShop9    field.Float32
	AcidShop10   field.Float32
	AcidShop11   field.Float32
	AcidShop12   field.Float32
	AcidShop13   field.Float32
	AcidShop14   field.Float32

	fieldMap map[string]field.Expr
}

func (a acidShop) Table(newTableName string) *acidShop {
	a.acidShopDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a acidShop) As(alias string) *acidShop {
	a.acidShopDo.DO = *(a.acidShopDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *acidShop) updateTableName(table string) *acidShop {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.HouseNumber = field.NewString(table, "house_number")
	a.TimeRecordAt = field.NewTime(table, "time_record_at")
	a.AcidShop1 = field.NewFloat32(table, "acid_shop1")
	a.AcidShop2 = field.NewFloat32(table, "acid_shop2")
	a.AcidShop3 = field.NewFloat32(table, "acid_shop3")
	a.AcidShop4 = field.NewFloat32(table, "acid_shop4")
	a.AcidShop5 = field.NewFloat32(table, "acid_shop5")
	a.AcidShop6 = field.NewFloat32(table, "acid_shop6")
	a.AcidShop7 = field.NewFloat32(table, "acid_shop7")
	a.AcidShop8 = field.NewFloat32(table, "acid_shop8")
	a.AcidShop9 = field.NewFloat32(table, "acid_shop9")
	a.AcidShop10 = field.NewFloat32(table, "acid_shop10")
	a.AcidShop11 = field.NewFloat32(table, "acid_shop11")
	a.AcidShop12 = field.NewFloat32(table, "acid_shop12")
	a.AcidShop13 = field.NewFloat32(table, "acid_shop13")
	a.AcidShop14 = field.NewFloat32(table, "acid_shop14")

	a.fillFieldMap()

	return a
}

func (a *acidShop) WithContext(ctx context.Context) IAcidShopDo { return a.acidShopDo.WithContext(ctx) }

func (a acidShop) TableName() string { return a.acidShopDo.TableName() }

func (a acidShop) Alias() string { return a.acidShopDo.Alias() }

func (a acidShop) Columns(cols ...field.Expr) gen.Columns { return a.acidShopDo.Columns(cols...) }

func (a *acidShop) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *acidShop) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 20)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["house_number"] = a.HouseNumber
	a.fieldMap["time_record_at"] = a.TimeRecordAt
	a.fieldMap["acid_shop1"] = a.AcidShop1
	a.fieldMap["acid_shop2"] = a.AcidShop2
	a.fieldMap["acid_shop3"] = a.AcidShop3
	a.fieldMap["acid_shop4"] = a.AcidShop4
	a.fieldMap["acid_shop5"] = a.AcidShop5
	a.fieldMap["acid_shop6"] = a.AcidShop6
	a.fieldMap["acid_shop7"] = a.AcidShop7
	a.fieldMap["acid_shop8"] = a.AcidShop8
	a.fieldMap["acid_shop9"] = a.AcidShop9
	a.fieldMap["acid_shop10"] = a.AcidShop10
	a.fieldMap["acid_shop11"] = a.AcidShop11
	a.fieldMap["acid_shop12"] = a.AcidShop12
	a.fieldMap["acid_shop13"] = a.AcidShop13
	a.fieldMap["acid_shop14"] = a.AcidShop14
}

func (a acidShop) clone(db *gorm.DB) acidShop {
	a.acidShopDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a acidShop) replaceDB(db *gorm.DB) acidShop {
	a.acidShopDo.ReplaceDB(db)
	return a
}

type acidShopDo struct{ gen.DO }

type IAcidShopDo interface {
	gen.SubQuery
	Debug() IAcidShopDo
	WithContext(ctx context.Context) IAcidShopDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAcidShopDo
	WriteDB() IAcidShopDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAcidShopDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAcidShopDo
	Not(conds ...gen.Condition) IAcidShopDo
	Or(conds ...gen.Condition) IAcidShopDo
	Select(conds ...field.Expr) IAcidShopDo
	Where(conds ...gen.Condition) IAcidShopDo
	Order(conds ...field.Expr) IAcidShopDo
	Distinct(cols ...field.Expr) IAcidShopDo
	Omit(cols ...field.Expr) IAcidShopDo
	Join(table schema.Tabler, on ...field.Expr) IAcidShopDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAcidShopDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAcidShopDo
	Group(cols ...field.Expr) IAcidShopDo
	Having(conds ...gen.Condition) IAcidShopDo
	Limit(limit int) IAcidShopDo
	Offset(offset int) IAcidShopDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAcidShopDo
	Unscoped() IAcidShopDo
	Create(values ...*slaughter.AcidShop) error
	CreateInBatches(values []*slaughter.AcidShop, batchSize int) error
	Save(values ...*slaughter.AcidShop) error
	First() (*slaughter.AcidShop, error)
	Take() (*slaughter.AcidShop, error)
	Last() (*slaughter.AcidShop, error)
	Find() ([]*slaughter.AcidShop, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.AcidShop, err error)
	FindInBatches(result *[]*slaughter.AcidShop, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.AcidShop) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAcidShopDo
	Assign(attrs ...field.AssignExpr) IAcidShopDo
	Joins(fields ...field.RelationField) IAcidShopDo
	Preload(fields ...field.RelationField) IAcidShopDo
	FirstOrInit() (*slaughter.AcidShop, error)
	FirstOrCreate() (*slaughter.AcidShop, error)
	FindByPage(offset int, limit int) (result []*slaughter.AcidShop, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAcidShopDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a acidShopDo) Debug() IAcidShopDo {
	return a.withDO(a.DO.Debug())
}

func (a acidShopDo) WithContext(ctx context.Context) IAcidShopDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a acidShopDo) ReadDB() IAcidShopDo {
	return a.Clauses(dbresolver.Read)
}

func (a acidShopDo) WriteDB() IAcidShopDo {
	return a.Clauses(dbresolver.Write)
}

func (a acidShopDo) Session(config *gorm.Session) IAcidShopDo {
	return a.withDO(a.DO.Session(config))
}

func (a acidShopDo) Clauses(conds ...clause.Expression) IAcidShopDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a acidShopDo) Returning(value interface{}, columns ...string) IAcidShopDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a acidShopDo) Not(conds ...gen.Condition) IAcidShopDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a acidShopDo) Or(conds ...gen.Condition) IAcidShopDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a acidShopDo) Select(conds ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a acidShopDo) Where(conds ...gen.Condition) IAcidShopDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a acidShopDo) Order(conds ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a acidShopDo) Distinct(cols ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a acidShopDo) Omit(cols ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a acidShopDo) Join(table schema.Tabler, on ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a acidShopDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a acidShopDo) RightJoin(table schema.Tabler, on ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a acidShopDo) Group(cols ...field.Expr) IAcidShopDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a acidShopDo) Having(conds ...gen.Condition) IAcidShopDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a acidShopDo) Limit(limit int) IAcidShopDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a acidShopDo) Offset(offset int) IAcidShopDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a acidShopDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAcidShopDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a acidShopDo) Unscoped() IAcidShopDo {
	return a.withDO(a.DO.Unscoped())
}

func (a acidShopDo) Create(values ...*slaughter.AcidShop) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a acidShopDo) CreateInBatches(values []*slaughter.AcidShop, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a acidShopDo) Save(values ...*slaughter.AcidShop) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a acidShopDo) First() (*slaughter.AcidShop, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AcidShop), nil
	}
}

func (a acidShopDo) Take() (*slaughter.AcidShop, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AcidShop), nil
	}
}

func (a acidShopDo) Last() (*slaughter.AcidShop, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AcidShop), nil
	}
}

func (a acidShopDo) Find() ([]*slaughter.AcidShop, error) {
	result, err := a.DO.Find()
	return result.([]*slaughter.AcidShop), err
}

func (a acidShopDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.AcidShop, err error) {
	buf := make([]*slaughter.AcidShop, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a acidShopDo) FindInBatches(result *[]*slaughter.AcidShop, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a acidShopDo) Attrs(attrs ...field.AssignExpr) IAcidShopDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a acidShopDo) Assign(attrs ...field.AssignExpr) IAcidShopDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a acidShopDo) Joins(fields ...field.RelationField) IAcidShopDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a acidShopDo) Preload(fields ...field.RelationField) IAcidShopDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a acidShopDo) FirstOrInit() (*slaughter.AcidShop, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AcidShop), nil
	}
}

func (a acidShopDo) FirstOrCreate() (*slaughter.AcidShop, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.AcidShop), nil
	}
}

func (a acidShopDo) FindByPage(offset int, limit int) (result []*slaughter.AcidShop, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a acidShopDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a acidShopDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a acidShopDo) Delete(models ...*slaughter.AcidShop) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *acidShopDo) withDO(do gen.Dao) *acidShopDo {
	a.DO = *do.(*gen.DO)
	return a
}
