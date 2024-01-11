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

func newFroShop(db *gorm.DB, opts ...gen.DOOption) froShop {
	_froShop := froShop{}

	_froShop.froShopDo.UseDB(db, opts...)
	_froShop.froShopDo.UseModel(&slaughter.FroShop{})

	tableName := _froShop.froShopDo.TableName()
	_froShop.ALL = field.NewAsterisk(tableName)
	_froShop.ID = field.NewUint(tableName, "id")
	_froShop.CreatedAt = field.NewTime(tableName, "created_at")
	_froShop.UpdatedAt = field.NewTime(tableName, "updated_at")
	_froShop.DeletedAt = field.NewField(tableName, "deleted_at")
	_froShop.HouseNumber = field.NewString(tableName, "house_number")
	_froShop.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_froShop.FroShop1 = field.NewFloat32(tableName, "fro_shop1")
	_froShop.FroShop2 = field.NewFloat32(tableName, "fro_shop2")
	_froShop.FroShop3 = field.NewFloat32(tableName, "fro_shop3")
	_froShop.FroShop4 = field.NewFloat32(tableName, "fro_shop4")
	_froShop.FroShop5 = field.NewFloat32(tableName, "fro_shop5")
	_froShop.FroShop6 = field.NewFloat32(tableName, "fro_shop6")
	_froShop.FroShop7 = field.NewFloat32(tableName, "fro_shop7")
	_froShop.FroShop8 = field.NewFloat32(tableName, "fro_shop8")
	_froShop.FroShop9 = field.NewFloat32(tableName, "fro_shop9")
	_froShop.FroShop10 = field.NewFloat32(tableName, "fro_shop10")
	_froShop.FroShop11 = field.NewFloat32(tableName, "fro_shop11")
	_froShop.FroShop12 = field.NewFloat32(tableName, "fro_shop12")
	_froShop.FroShop13 = field.NewFloat32(tableName, "fro_shop13")
	_froShop.FroShop14 = field.NewFloat32(tableName, "fro_shop14")

	_froShop.fillFieldMap()

	return _froShop
}

type froShop struct {
	froShopDo froShopDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HouseNumber  field.String
	TimeRecordAt field.Time
	FroShop1     field.Float32
	FroShop2     field.Float32
	FroShop3     field.Float32
	FroShop4     field.Float32
	FroShop5     field.Float32
	FroShop6     field.Float32
	FroShop7     field.Float32
	FroShop8     field.Float32
	FroShop9     field.Float32
	FroShop10    field.Float32
	FroShop11    field.Float32
	FroShop12    field.Float32
	FroShop13    field.Float32
	FroShop14    field.Float32

	fieldMap map[string]field.Expr
}

func (f froShop) Table(newTableName string) *froShop {
	f.froShopDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f froShop) As(alias string) *froShop {
	f.froShopDo.DO = *(f.froShopDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *froShop) updateTableName(table string) *froShop {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewUint(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.HouseNumber = field.NewString(table, "house_number")
	f.TimeRecordAt = field.NewTime(table, "time_record_at")
	f.FroShop1 = field.NewFloat32(table, "fro_shop1")
	f.FroShop2 = field.NewFloat32(table, "fro_shop2")
	f.FroShop3 = field.NewFloat32(table, "fro_shop3")
	f.FroShop4 = field.NewFloat32(table, "fro_shop4")
	f.FroShop5 = field.NewFloat32(table, "fro_shop5")
	f.FroShop6 = field.NewFloat32(table, "fro_shop6")
	f.FroShop7 = field.NewFloat32(table, "fro_shop7")
	f.FroShop8 = field.NewFloat32(table, "fro_shop8")
	f.FroShop9 = field.NewFloat32(table, "fro_shop9")
	f.FroShop10 = field.NewFloat32(table, "fro_shop10")
	f.FroShop11 = field.NewFloat32(table, "fro_shop11")
	f.FroShop12 = field.NewFloat32(table, "fro_shop12")
	f.FroShop13 = field.NewFloat32(table, "fro_shop13")
	f.FroShop14 = field.NewFloat32(table, "fro_shop14")

	f.fillFieldMap()

	return f
}

func (f *froShop) WithContext(ctx context.Context) IFroShopDo { return f.froShopDo.WithContext(ctx) }

func (f froShop) TableName() string { return f.froShopDo.TableName() }

func (f froShop) Alias() string { return f.froShopDo.Alias() }

func (f froShop) Columns(cols ...field.Expr) gen.Columns { return f.froShopDo.Columns(cols...) }

func (f *froShop) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *froShop) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 20)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["house_number"] = f.HouseNumber
	f.fieldMap["time_record_at"] = f.TimeRecordAt
	f.fieldMap["fro_shop1"] = f.FroShop1
	f.fieldMap["fro_shop2"] = f.FroShop2
	f.fieldMap["fro_shop3"] = f.FroShop3
	f.fieldMap["fro_shop4"] = f.FroShop4
	f.fieldMap["fro_shop5"] = f.FroShop5
	f.fieldMap["fro_shop6"] = f.FroShop6
	f.fieldMap["fro_shop7"] = f.FroShop7
	f.fieldMap["fro_shop8"] = f.FroShop8
	f.fieldMap["fro_shop9"] = f.FroShop9
	f.fieldMap["fro_shop10"] = f.FroShop10
	f.fieldMap["fro_shop11"] = f.FroShop11
	f.fieldMap["fro_shop12"] = f.FroShop12
	f.fieldMap["fro_shop13"] = f.FroShop13
	f.fieldMap["fro_shop14"] = f.FroShop14
}

func (f froShop) clone(db *gorm.DB) froShop {
	f.froShopDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f froShop) replaceDB(db *gorm.DB) froShop {
	f.froShopDo.ReplaceDB(db)
	return f
}

type froShopDo struct{ gen.DO }

type IFroShopDo interface {
	gen.SubQuery
	Debug() IFroShopDo
	WithContext(ctx context.Context) IFroShopDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFroShopDo
	WriteDB() IFroShopDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFroShopDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFroShopDo
	Not(conds ...gen.Condition) IFroShopDo
	Or(conds ...gen.Condition) IFroShopDo
	Select(conds ...field.Expr) IFroShopDo
	Where(conds ...gen.Condition) IFroShopDo
	Order(conds ...field.Expr) IFroShopDo
	Distinct(cols ...field.Expr) IFroShopDo
	Omit(cols ...field.Expr) IFroShopDo
	Join(table schema.Tabler, on ...field.Expr) IFroShopDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFroShopDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFroShopDo
	Group(cols ...field.Expr) IFroShopDo
	Having(conds ...gen.Condition) IFroShopDo
	Limit(limit int) IFroShopDo
	Offset(offset int) IFroShopDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFroShopDo
	Unscoped() IFroShopDo
	Create(values ...*slaughter.FroShop) error
	CreateInBatches(values []*slaughter.FroShop, batchSize int) error
	Save(values ...*slaughter.FroShop) error
	First() (*slaughter.FroShop, error)
	Take() (*slaughter.FroShop, error)
	Last() (*slaughter.FroShop, error)
	Find() ([]*slaughter.FroShop, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.FroShop, err error)
	FindInBatches(result *[]*slaughter.FroShop, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.FroShop) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFroShopDo
	Assign(attrs ...field.AssignExpr) IFroShopDo
	Joins(fields ...field.RelationField) IFroShopDo
	Preload(fields ...field.RelationField) IFroShopDo
	FirstOrInit() (*slaughter.FroShop, error)
	FirstOrCreate() (*slaughter.FroShop, error)
	FindByPage(offset int, limit int) (result []*slaughter.FroShop, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFroShopDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f froShopDo) Debug() IFroShopDo {
	return f.withDO(f.DO.Debug())
}

func (f froShopDo) WithContext(ctx context.Context) IFroShopDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f froShopDo) ReadDB() IFroShopDo {
	return f.Clauses(dbresolver.Read)
}

func (f froShopDo) WriteDB() IFroShopDo {
	return f.Clauses(dbresolver.Write)
}

func (f froShopDo) Session(config *gorm.Session) IFroShopDo {
	return f.withDO(f.DO.Session(config))
}

func (f froShopDo) Clauses(conds ...clause.Expression) IFroShopDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f froShopDo) Returning(value interface{}, columns ...string) IFroShopDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f froShopDo) Not(conds ...gen.Condition) IFroShopDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f froShopDo) Or(conds ...gen.Condition) IFroShopDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f froShopDo) Select(conds ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f froShopDo) Where(conds ...gen.Condition) IFroShopDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f froShopDo) Order(conds ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f froShopDo) Distinct(cols ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f froShopDo) Omit(cols ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f froShopDo) Join(table schema.Tabler, on ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f froShopDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f froShopDo) RightJoin(table schema.Tabler, on ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f froShopDo) Group(cols ...field.Expr) IFroShopDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f froShopDo) Having(conds ...gen.Condition) IFroShopDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f froShopDo) Limit(limit int) IFroShopDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f froShopDo) Offset(offset int) IFroShopDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f froShopDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFroShopDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f froShopDo) Unscoped() IFroShopDo {
	return f.withDO(f.DO.Unscoped())
}

func (f froShopDo) Create(values ...*slaughter.FroShop) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f froShopDo) CreateInBatches(values []*slaughter.FroShop, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f froShopDo) Save(values ...*slaughter.FroShop) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f froShopDo) First() (*slaughter.FroShop, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FroShop), nil
	}
}

func (f froShopDo) Take() (*slaughter.FroShop, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FroShop), nil
	}
}

func (f froShopDo) Last() (*slaughter.FroShop, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FroShop), nil
	}
}

func (f froShopDo) Find() ([]*slaughter.FroShop, error) {
	result, err := f.DO.Find()
	return result.([]*slaughter.FroShop), err
}

func (f froShopDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.FroShop, err error) {
	buf := make([]*slaughter.FroShop, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f froShopDo) FindInBatches(result *[]*slaughter.FroShop, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f froShopDo) Attrs(attrs ...field.AssignExpr) IFroShopDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f froShopDo) Assign(attrs ...field.AssignExpr) IFroShopDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f froShopDo) Joins(fields ...field.RelationField) IFroShopDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f froShopDo) Preload(fields ...field.RelationField) IFroShopDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f froShopDo) FirstOrInit() (*slaughter.FroShop, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FroShop), nil
	}
}

func (f froShopDo) FirstOrCreate() (*slaughter.FroShop, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FroShop), nil
	}
}

func (f froShopDo) FindByPage(offset int, limit int) (result []*slaughter.FroShop, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f froShopDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f froShopDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f froShopDo) Delete(models ...*slaughter.FroShop) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *froShopDo) withDO(do gen.Dao) *froShopDo {
	f.DO = *do.(*gen.DO)
	return f
}
