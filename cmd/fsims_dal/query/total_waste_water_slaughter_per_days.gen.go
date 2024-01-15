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

func newTotalWasteWaterSlaughterPerDay(db *gorm.DB, opts ...gen.DOOption) totalWasteWaterSlaughterPerDay {
	_totalWasteWaterSlaughterPerDay := totalWasteWaterSlaughterPerDay{}

	_totalWasteWaterSlaughterPerDay.totalWasteWaterSlaughterPerDayDo.UseDB(db, opts...)
	_totalWasteWaterSlaughterPerDay.totalWasteWaterSlaughterPerDayDo.UseModel(&slaughter.TotalWasteWaterSlaughterPerDay{})

	tableName := _totalWasteWaterSlaughterPerDay.totalWasteWaterSlaughterPerDayDo.TableName()
	_totalWasteWaterSlaughterPerDay.ALL = field.NewAsterisk(tableName)
	_totalWasteWaterSlaughterPerDay.ID = field.NewUint(tableName, "id")
	_totalWasteWaterSlaughterPerDay.CreatedAt = field.NewTime(tableName, "created_at")
	_totalWasteWaterSlaughterPerDay.UpdatedAt = field.NewTime(tableName, "updated_at")
	_totalWasteWaterSlaughterPerDay.DeletedAt = field.NewField(tableName, "deleted_at")
	_totalWasteWaterSlaughterPerDay.TimeStamp = field.NewTime(tableName, "time_stamp")
	_totalWasteWaterSlaughterPerDay.HouseNumber = field.NewString(tableName, "house_number")
	_totalWasteWaterSlaughterPerDay.TotalWasteWaterSlaughterPerDay1 = field.NewFloat32(tableName, "total_waste_water_slaughter_per_day1")
	_totalWasteWaterSlaughterPerDay.TotalWasteWaterSlaughterPerDay2 = field.NewFloat32(tableName, "total_waste_water_slaughter_per_day2")
	_totalWasteWaterSlaughterPerDay.TotalWasteWaterSlaughterPerDay3 = field.NewFloat32(tableName, "total_waste_water_slaughter_per_day3")
	_totalWasteWaterSlaughterPerDay.TotalWasteWaterSlaughterPerDay4 = field.NewFloat32(tableName, "total_waste_water_slaughter_per_day4")

	_totalWasteWaterSlaughterPerDay.fillFieldMap()

	return _totalWasteWaterSlaughterPerDay
}

type totalWasteWaterSlaughterPerDay struct {
	totalWasteWaterSlaughterPerDayDo totalWasteWaterSlaughterPerDayDo

	ALL                             field.Asterisk
	ID                              field.Uint
	CreatedAt                       field.Time
	UpdatedAt                       field.Time
	DeletedAt                       field.Field
	TimeStamp                       field.Time
	HouseNumber                     field.String
	TotalWasteWaterSlaughterPerDay1 field.Float32
	TotalWasteWaterSlaughterPerDay2 field.Float32
	TotalWasteWaterSlaughterPerDay3 field.Float32
	TotalWasteWaterSlaughterPerDay4 field.Float32

	fieldMap map[string]field.Expr
}

func (t totalWasteWaterSlaughterPerDay) Table(newTableName string) *totalWasteWaterSlaughterPerDay {
	t.totalWasteWaterSlaughterPerDayDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t totalWasteWaterSlaughterPerDay) As(alias string) *totalWasteWaterSlaughterPerDay {
	t.totalWasteWaterSlaughterPerDayDo.DO = *(t.totalWasteWaterSlaughterPerDayDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *totalWasteWaterSlaughterPerDay) updateTableName(table string) *totalWasteWaterSlaughterPerDay {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TimeStamp = field.NewTime(table, "time_stamp")
	t.HouseNumber = field.NewString(table, "house_number")
	t.TotalWasteWaterSlaughterPerDay1 = field.NewFloat32(table, "total_waste_water_slaughter_per_day1")
	t.TotalWasteWaterSlaughterPerDay2 = field.NewFloat32(table, "total_waste_water_slaughter_per_day2")
	t.TotalWasteWaterSlaughterPerDay3 = field.NewFloat32(table, "total_waste_water_slaughter_per_day3")
	t.TotalWasteWaterSlaughterPerDay4 = field.NewFloat32(table, "total_waste_water_slaughter_per_day4")

	t.fillFieldMap()

	return t
}

func (t *totalWasteWaterSlaughterPerDay) WithContext(ctx context.Context) ITotalWasteWaterSlaughterPerDayDo {
	return t.totalWasteWaterSlaughterPerDayDo.WithContext(ctx)
}

func (t totalWasteWaterSlaughterPerDay) TableName() string {
	return t.totalWasteWaterSlaughterPerDayDo.TableName()
}

func (t totalWasteWaterSlaughterPerDay) Alias() string {
	return t.totalWasteWaterSlaughterPerDayDo.Alias()
}

func (t totalWasteWaterSlaughterPerDay) Columns(cols ...field.Expr) gen.Columns {
	return t.totalWasteWaterSlaughterPerDayDo.Columns(cols...)
}

func (t *totalWasteWaterSlaughterPerDay) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *totalWasteWaterSlaughterPerDay) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["time_stamp"] = t.TimeStamp
	t.fieldMap["house_number"] = t.HouseNumber
	t.fieldMap["total_waste_water_slaughter_per_day1"] = t.TotalWasteWaterSlaughterPerDay1
	t.fieldMap["total_waste_water_slaughter_per_day2"] = t.TotalWasteWaterSlaughterPerDay2
	t.fieldMap["total_waste_water_slaughter_per_day3"] = t.TotalWasteWaterSlaughterPerDay3
	t.fieldMap["total_waste_water_slaughter_per_day4"] = t.TotalWasteWaterSlaughterPerDay4
}

func (t totalWasteWaterSlaughterPerDay) clone(db *gorm.DB) totalWasteWaterSlaughterPerDay {
	t.totalWasteWaterSlaughterPerDayDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t totalWasteWaterSlaughterPerDay) replaceDB(db *gorm.DB) totalWasteWaterSlaughterPerDay {
	t.totalWasteWaterSlaughterPerDayDo.ReplaceDB(db)
	return t
}

type totalWasteWaterSlaughterPerDayDo struct{ gen.DO }

type ITotalWasteWaterSlaughterPerDayDo interface {
	gen.SubQuery
	Debug() ITotalWasteWaterSlaughterPerDayDo
	WithContext(ctx context.Context) ITotalWasteWaterSlaughterPerDayDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITotalWasteWaterSlaughterPerDayDo
	WriteDB() ITotalWasteWaterSlaughterPerDayDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITotalWasteWaterSlaughterPerDayDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITotalWasteWaterSlaughterPerDayDo
	Not(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo
	Or(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo
	Select(conds ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	Where(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo
	Order(conds ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	Distinct(cols ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	Omit(cols ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	Join(table schema.Tabler, on ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	Group(cols ...field.Expr) ITotalWasteWaterSlaughterPerDayDo
	Having(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo
	Limit(limit int) ITotalWasteWaterSlaughterPerDayDo
	Offset(offset int) ITotalWasteWaterSlaughterPerDayDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITotalWasteWaterSlaughterPerDayDo
	Unscoped() ITotalWasteWaterSlaughterPerDayDo
	Create(values ...*slaughter.TotalWasteWaterSlaughterPerDay) error
	CreateInBatches(values []*slaughter.TotalWasteWaterSlaughterPerDay, batchSize int) error
	Save(values ...*slaughter.TotalWasteWaterSlaughterPerDay) error
	First() (*slaughter.TotalWasteWaterSlaughterPerDay, error)
	Take() (*slaughter.TotalWasteWaterSlaughterPerDay, error)
	Last() (*slaughter.TotalWasteWaterSlaughterPerDay, error)
	Find() ([]*slaughter.TotalWasteWaterSlaughterPerDay, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.TotalWasteWaterSlaughterPerDay, err error)
	FindInBatches(result *[]*slaughter.TotalWasteWaterSlaughterPerDay, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.TotalWasteWaterSlaughterPerDay) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITotalWasteWaterSlaughterPerDayDo
	Assign(attrs ...field.AssignExpr) ITotalWasteWaterSlaughterPerDayDo
	Joins(fields ...field.RelationField) ITotalWasteWaterSlaughterPerDayDo
	Preload(fields ...field.RelationField) ITotalWasteWaterSlaughterPerDayDo
	FirstOrInit() (*slaughter.TotalWasteWaterSlaughterPerDay, error)
	FirstOrCreate() (*slaughter.TotalWasteWaterSlaughterPerDay, error)
	FindByPage(offset int, limit int) (result []*slaughter.TotalWasteWaterSlaughterPerDay, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITotalWasteWaterSlaughterPerDayDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t totalWasteWaterSlaughterPerDayDo) Debug() ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Debug())
}

func (t totalWasteWaterSlaughterPerDayDo) WithContext(ctx context.Context) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t totalWasteWaterSlaughterPerDayDo) ReadDB() ITotalWasteWaterSlaughterPerDayDo {
	return t.Clauses(dbresolver.Read)
}

func (t totalWasteWaterSlaughterPerDayDo) WriteDB() ITotalWasteWaterSlaughterPerDayDo {
	return t.Clauses(dbresolver.Write)
}

func (t totalWasteWaterSlaughterPerDayDo) Session(config *gorm.Session) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Session(config))
}

func (t totalWasteWaterSlaughterPerDayDo) Clauses(conds ...clause.Expression) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Returning(value interface{}, columns ...string) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t totalWasteWaterSlaughterPerDayDo) Not(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Or(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Select(conds ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Where(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Order(conds ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Distinct(cols ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t totalWasteWaterSlaughterPerDayDo) Omit(cols ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t totalWasteWaterSlaughterPerDayDo) Join(table schema.Tabler, on ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t totalWasteWaterSlaughterPerDayDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t totalWasteWaterSlaughterPerDayDo) RightJoin(table schema.Tabler, on ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t totalWasteWaterSlaughterPerDayDo) Group(cols ...field.Expr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t totalWasteWaterSlaughterPerDayDo) Having(conds ...gen.Condition) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t totalWasteWaterSlaughterPerDayDo) Limit(limit int) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t totalWasteWaterSlaughterPerDayDo) Offset(offset int) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t totalWasteWaterSlaughterPerDayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t totalWasteWaterSlaughterPerDayDo) Unscoped() ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Unscoped())
}

func (t totalWasteWaterSlaughterPerDayDo) Create(values ...*slaughter.TotalWasteWaterSlaughterPerDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t totalWasteWaterSlaughterPerDayDo) CreateInBatches(values []*slaughter.TotalWasteWaterSlaughterPerDay, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t totalWasteWaterSlaughterPerDayDo) Save(values ...*slaughter.TotalWasteWaterSlaughterPerDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t totalWasteWaterSlaughterPerDayDo) First() (*slaughter.TotalWasteWaterSlaughterPerDay, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalWasteWaterSlaughterPerDay), nil
	}
}

func (t totalWasteWaterSlaughterPerDayDo) Take() (*slaughter.TotalWasteWaterSlaughterPerDay, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalWasteWaterSlaughterPerDay), nil
	}
}

func (t totalWasteWaterSlaughterPerDayDo) Last() (*slaughter.TotalWasteWaterSlaughterPerDay, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalWasteWaterSlaughterPerDay), nil
	}
}

func (t totalWasteWaterSlaughterPerDayDo) Find() ([]*slaughter.TotalWasteWaterSlaughterPerDay, error) {
	result, err := t.DO.Find()
	return result.([]*slaughter.TotalWasteWaterSlaughterPerDay), err
}

func (t totalWasteWaterSlaughterPerDayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.TotalWasteWaterSlaughterPerDay, err error) {
	buf := make([]*slaughter.TotalWasteWaterSlaughterPerDay, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t totalWasteWaterSlaughterPerDayDo) FindInBatches(result *[]*slaughter.TotalWasteWaterSlaughterPerDay, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t totalWasteWaterSlaughterPerDayDo) Attrs(attrs ...field.AssignExpr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t totalWasteWaterSlaughterPerDayDo) Assign(attrs ...field.AssignExpr) ITotalWasteWaterSlaughterPerDayDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t totalWasteWaterSlaughterPerDayDo) Joins(fields ...field.RelationField) ITotalWasteWaterSlaughterPerDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t totalWasteWaterSlaughterPerDayDo) Preload(fields ...field.RelationField) ITotalWasteWaterSlaughterPerDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t totalWasteWaterSlaughterPerDayDo) FirstOrInit() (*slaughter.TotalWasteWaterSlaughterPerDay, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalWasteWaterSlaughterPerDay), nil
	}
}

func (t totalWasteWaterSlaughterPerDayDo) FirstOrCreate() (*slaughter.TotalWasteWaterSlaughterPerDay, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalWasteWaterSlaughterPerDay), nil
	}
}

func (t totalWasteWaterSlaughterPerDayDo) FindByPage(offset int, limit int) (result []*slaughter.TotalWasteWaterSlaughterPerDay, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t totalWasteWaterSlaughterPerDayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t totalWasteWaterSlaughterPerDayDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t totalWasteWaterSlaughterPerDayDo) Delete(models ...*slaughter.TotalWasteWaterSlaughterPerDay) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *totalWasteWaterSlaughterPerDayDo) withDO(do gen.Dao) *totalWasteWaterSlaughterPerDayDo {
	t.DO = *do.(*gen.DO)
	return t
}