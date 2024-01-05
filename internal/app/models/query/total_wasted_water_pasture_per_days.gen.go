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

func newTotalWastedWaterPasturePerDay(db *gorm.DB, opts ...gen.DOOption) totalWastedWaterPasturePerDay {
	_totalWastedWaterPasturePerDay := totalWastedWaterPasturePerDay{}

	_totalWastedWaterPasturePerDay.totalWastedWaterPasturePerDayDo.UseDB(db, opts...)
	_totalWastedWaterPasturePerDay.totalWastedWaterPasturePerDayDo.UseModel(&pasture.TotalWastedWaterPasturePerDay{})

	tableName := _totalWastedWaterPasturePerDay.totalWastedWaterPasturePerDayDo.TableName()
	_totalWastedWaterPasturePerDay.ALL = field.NewAsterisk(tableName)
	_totalWastedWaterPasturePerDay.ID = field.NewUint(tableName, "id")
	_totalWastedWaterPasturePerDay.CreatedAt = field.NewTime(tableName, "created_at")
	_totalWastedWaterPasturePerDay.UpdatedAt = field.NewTime(tableName, "updated_at")
	_totalWastedWaterPasturePerDay.DeletedAt = field.NewField(tableName, "deleted_at")
	_totalWastedWaterPasturePerDay.TimeStamp = field.NewTime(tableName, "time_stamp")
	_totalWastedWaterPasturePerDay.HouseNumber = field.NewString(tableName, "house_number")
	_totalWastedWaterPasturePerDay.TotalWastedWaterPerDay1 = field.NewFloat32(tableName, "total_wasted_water_per_day1")
	_totalWastedWaterPasturePerDay.TotalWastedWaterPerDay2 = field.NewFloat32(tableName, "total_wasted_water_per_day2")
	_totalWastedWaterPasturePerDay.TotalWastedWaterPerDay3 = field.NewFloat32(tableName, "total_wasted_water_per_day3")
	_totalWastedWaterPasturePerDay.TotalWastedWaterPerDay4 = field.NewFloat32(tableName, "total_wasted_water_per_day4")

	_totalWastedWaterPasturePerDay.fillFieldMap()

	return _totalWastedWaterPasturePerDay
}

type totalWastedWaterPasturePerDay struct {
	totalWastedWaterPasturePerDayDo totalWastedWaterPasturePerDayDo

	ALL                     field.Asterisk
	ID                      field.Uint
	CreatedAt               field.Time
	UpdatedAt               field.Time
	DeletedAt               field.Field
	TimeStamp               field.Time
	HouseNumber             field.String
	TotalWastedWaterPerDay1 field.Float32
	TotalWastedWaterPerDay2 field.Float32
	TotalWastedWaterPerDay3 field.Float32
	TotalWastedWaterPerDay4 field.Float32

	fieldMap map[string]field.Expr
}

func (t totalWastedWaterPasturePerDay) Table(newTableName string) *totalWastedWaterPasturePerDay {
	t.totalWastedWaterPasturePerDayDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t totalWastedWaterPasturePerDay) As(alias string) *totalWastedWaterPasturePerDay {
	t.totalWastedWaterPasturePerDayDo.DO = *(t.totalWastedWaterPasturePerDayDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *totalWastedWaterPasturePerDay) updateTableName(table string) *totalWastedWaterPasturePerDay {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TimeStamp = field.NewTime(table, "time_stamp")
	t.HouseNumber = field.NewString(table, "house_number")
	t.TotalWastedWaterPerDay1 = field.NewFloat32(table, "total_wasted_water_per_day1")
	t.TotalWastedWaterPerDay2 = field.NewFloat32(table, "total_wasted_water_per_day2")
	t.TotalWastedWaterPerDay3 = field.NewFloat32(table, "total_wasted_water_per_day3")
	t.TotalWastedWaterPerDay4 = field.NewFloat32(table, "total_wasted_water_per_day4")

	t.fillFieldMap()

	return t
}

func (t *totalWastedWaterPasturePerDay) WithContext(ctx context.Context) ITotalWastedWaterPasturePerDayDo {
	return t.totalWastedWaterPasturePerDayDo.WithContext(ctx)
}

func (t totalWastedWaterPasturePerDay) TableName() string {
	return t.totalWastedWaterPasturePerDayDo.TableName()
}

func (t totalWastedWaterPasturePerDay) Alias() string {
	return t.totalWastedWaterPasturePerDayDo.Alias()
}

func (t totalWastedWaterPasturePerDay) Columns(cols ...field.Expr) gen.Columns {
	return t.totalWastedWaterPasturePerDayDo.Columns(cols...)
}

func (t *totalWastedWaterPasturePerDay) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *totalWastedWaterPasturePerDay) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["time_stamp"] = t.TimeStamp
	t.fieldMap["house_number"] = t.HouseNumber
	t.fieldMap["total_wasted_water_per_day1"] = t.TotalWastedWaterPerDay1
	t.fieldMap["total_wasted_water_per_day2"] = t.TotalWastedWaterPerDay2
	t.fieldMap["total_wasted_water_per_day3"] = t.TotalWastedWaterPerDay3
	t.fieldMap["total_wasted_water_per_day4"] = t.TotalWastedWaterPerDay4
}

func (t totalWastedWaterPasturePerDay) clone(db *gorm.DB) totalWastedWaterPasturePerDay {
	t.totalWastedWaterPasturePerDayDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t totalWastedWaterPasturePerDay) replaceDB(db *gorm.DB) totalWastedWaterPasturePerDay {
	t.totalWastedWaterPasturePerDayDo.ReplaceDB(db)
	return t
}

type totalWastedWaterPasturePerDayDo struct{ gen.DO }

type ITotalWastedWaterPasturePerDayDo interface {
	gen.SubQuery
	Debug() ITotalWastedWaterPasturePerDayDo
	WithContext(ctx context.Context) ITotalWastedWaterPasturePerDayDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITotalWastedWaterPasturePerDayDo
	WriteDB() ITotalWastedWaterPasturePerDayDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITotalWastedWaterPasturePerDayDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITotalWastedWaterPasturePerDayDo
	Not(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo
	Or(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo
	Select(conds ...field.Expr) ITotalWastedWaterPasturePerDayDo
	Where(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo
	Order(conds ...field.Expr) ITotalWastedWaterPasturePerDayDo
	Distinct(cols ...field.Expr) ITotalWastedWaterPasturePerDayDo
	Omit(cols ...field.Expr) ITotalWastedWaterPasturePerDayDo
	Join(table schema.Tabler, on ...field.Expr) ITotalWastedWaterPasturePerDayDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITotalWastedWaterPasturePerDayDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITotalWastedWaterPasturePerDayDo
	Group(cols ...field.Expr) ITotalWastedWaterPasturePerDayDo
	Having(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo
	Limit(limit int) ITotalWastedWaterPasturePerDayDo
	Offset(offset int) ITotalWastedWaterPasturePerDayDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITotalWastedWaterPasturePerDayDo
	Unscoped() ITotalWastedWaterPasturePerDayDo
	Create(values ...*pasture.TotalWastedWaterPasturePerDay) error
	CreateInBatches(values []*pasture.TotalWastedWaterPasturePerDay, batchSize int) error
	Save(values ...*pasture.TotalWastedWaterPasturePerDay) error
	First() (*pasture.TotalWastedWaterPasturePerDay, error)
	Take() (*pasture.TotalWastedWaterPasturePerDay, error)
	Last() (*pasture.TotalWastedWaterPasturePerDay, error)
	Find() ([]*pasture.TotalWastedWaterPasturePerDay, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.TotalWastedWaterPasturePerDay, err error)
	FindInBatches(result *[]*pasture.TotalWastedWaterPasturePerDay, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.TotalWastedWaterPasturePerDay) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITotalWastedWaterPasturePerDayDo
	Assign(attrs ...field.AssignExpr) ITotalWastedWaterPasturePerDayDo
	Joins(fields ...field.RelationField) ITotalWastedWaterPasturePerDayDo
	Preload(fields ...field.RelationField) ITotalWastedWaterPasturePerDayDo
	FirstOrInit() (*pasture.TotalWastedWaterPasturePerDay, error)
	FirstOrCreate() (*pasture.TotalWastedWaterPasturePerDay, error)
	FindByPage(offset int, limit int) (result []*pasture.TotalWastedWaterPasturePerDay, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITotalWastedWaterPasturePerDayDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t totalWastedWaterPasturePerDayDo) Debug() ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Debug())
}

func (t totalWastedWaterPasturePerDayDo) WithContext(ctx context.Context) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t totalWastedWaterPasturePerDayDo) ReadDB() ITotalWastedWaterPasturePerDayDo {
	return t.Clauses(dbresolver.Read)
}

func (t totalWastedWaterPasturePerDayDo) WriteDB() ITotalWastedWaterPasturePerDayDo {
	return t.Clauses(dbresolver.Write)
}

func (t totalWastedWaterPasturePerDayDo) Session(config *gorm.Session) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Session(config))
}

func (t totalWastedWaterPasturePerDayDo) Clauses(conds ...clause.Expression) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Returning(value interface{}, columns ...string) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t totalWastedWaterPasturePerDayDo) Not(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Or(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Select(conds ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Where(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Order(conds ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Distinct(cols ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t totalWastedWaterPasturePerDayDo) Omit(cols ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t totalWastedWaterPasturePerDayDo) Join(table schema.Tabler, on ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t totalWastedWaterPasturePerDayDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t totalWastedWaterPasturePerDayDo) RightJoin(table schema.Tabler, on ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t totalWastedWaterPasturePerDayDo) Group(cols ...field.Expr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t totalWastedWaterPasturePerDayDo) Having(conds ...gen.Condition) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t totalWastedWaterPasturePerDayDo) Limit(limit int) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t totalWastedWaterPasturePerDayDo) Offset(offset int) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t totalWastedWaterPasturePerDayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t totalWastedWaterPasturePerDayDo) Unscoped() ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Unscoped())
}

func (t totalWastedWaterPasturePerDayDo) Create(values ...*pasture.TotalWastedWaterPasturePerDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t totalWastedWaterPasturePerDayDo) CreateInBatches(values []*pasture.TotalWastedWaterPasturePerDay, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t totalWastedWaterPasturePerDayDo) Save(values ...*pasture.TotalWastedWaterPasturePerDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t totalWastedWaterPasturePerDayDo) First() (*pasture.TotalWastedWaterPasturePerDay, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.TotalWastedWaterPasturePerDay), nil
	}
}

func (t totalWastedWaterPasturePerDayDo) Take() (*pasture.TotalWastedWaterPasturePerDay, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.TotalWastedWaterPasturePerDay), nil
	}
}

func (t totalWastedWaterPasturePerDayDo) Last() (*pasture.TotalWastedWaterPasturePerDay, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.TotalWastedWaterPasturePerDay), nil
	}
}

func (t totalWastedWaterPasturePerDayDo) Find() ([]*pasture.TotalWastedWaterPasturePerDay, error) {
	result, err := t.DO.Find()
	return result.([]*pasture.TotalWastedWaterPasturePerDay), err
}

func (t totalWastedWaterPasturePerDayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.TotalWastedWaterPasturePerDay, err error) {
	buf := make([]*pasture.TotalWastedWaterPasturePerDay, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t totalWastedWaterPasturePerDayDo) FindInBatches(result *[]*pasture.TotalWastedWaterPasturePerDay, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t totalWastedWaterPasturePerDayDo) Attrs(attrs ...field.AssignExpr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t totalWastedWaterPasturePerDayDo) Assign(attrs ...field.AssignExpr) ITotalWastedWaterPasturePerDayDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t totalWastedWaterPasturePerDayDo) Joins(fields ...field.RelationField) ITotalWastedWaterPasturePerDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t totalWastedWaterPasturePerDayDo) Preload(fields ...field.RelationField) ITotalWastedWaterPasturePerDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t totalWastedWaterPasturePerDayDo) FirstOrInit() (*pasture.TotalWastedWaterPasturePerDay, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.TotalWastedWaterPasturePerDay), nil
	}
}

func (t totalWastedWaterPasturePerDayDo) FirstOrCreate() (*pasture.TotalWastedWaterPasturePerDay, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.TotalWastedWaterPasturePerDay), nil
	}
}

func (t totalWastedWaterPasturePerDayDo) FindByPage(offset int, limit int) (result []*pasture.TotalWastedWaterPasturePerDay, count int64, err error) {
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

func (t totalWastedWaterPasturePerDayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t totalWastedWaterPasturePerDayDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t totalWastedWaterPasturePerDayDo) Delete(models ...*pasture.TotalWastedWaterPasturePerDay) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *totalWastedWaterPasturePerDayDo) withDO(do gen.Dao) *totalWastedWaterPasturePerDayDo {
	t.DO = *do.(*gen.DO)
	return t
}
