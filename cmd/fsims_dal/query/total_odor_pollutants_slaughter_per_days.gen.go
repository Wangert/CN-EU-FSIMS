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

func newTotalOdorPollutantsSlaughterPerDay(db *gorm.DB, opts ...gen.DOOption) totalOdorPollutantsSlaughterPerDay {
	_totalOdorPollutantsSlaughterPerDay := totalOdorPollutantsSlaughterPerDay{}

	_totalOdorPollutantsSlaughterPerDay.totalOdorPollutantsSlaughterPerDayDo.UseDB(db, opts...)
	_totalOdorPollutantsSlaughterPerDay.totalOdorPollutantsSlaughterPerDayDo.UseModel(&slaughter.TotalOdorPollutantsSlaughterPerDay{})

	tableName := _totalOdorPollutantsSlaughterPerDay.totalOdorPollutantsSlaughterPerDayDo.TableName()
	_totalOdorPollutantsSlaughterPerDay.ALL = field.NewAsterisk(tableName)
	_totalOdorPollutantsSlaughterPerDay.ID = field.NewUint(tableName, "id")
	_totalOdorPollutantsSlaughterPerDay.CreatedAt = field.NewTime(tableName, "created_at")
	_totalOdorPollutantsSlaughterPerDay.UpdatedAt = field.NewTime(tableName, "updated_at")
	_totalOdorPollutantsSlaughterPerDay.DeletedAt = field.NewField(tableName, "deleted_at")
	_totalOdorPollutantsSlaughterPerDay.TimeStamp = field.NewTime(tableName, "time_stamp")
	_totalOdorPollutantsSlaughterPerDay.HouseNumber = field.NewString(tableName, "house_number")
	_totalOdorPollutantsSlaughterPerDay.TotalOdorPollutantsSlaughterPerDay1 = field.NewFloat64(tableName, "total_odor_pollutants_slaughter_per_day1")
	_totalOdorPollutantsSlaughterPerDay.TotalOdorPollutantsSlaughterPerDay2 = field.NewFloat64(tableName, "total_odor_pollutants_slaughter_per_day2")
	_totalOdorPollutantsSlaughterPerDay.TotalOdorPollutantsSlaughterPerDay3 = field.NewFloat64(tableName, "total_odor_pollutants_slaughter_per_day3")
	_totalOdorPollutantsSlaughterPerDay.TotalOdorPollutantsSlaughterPerDay4 = field.NewFloat64(tableName, "total_odor_pollutants_slaughter_per_day4")

	_totalOdorPollutantsSlaughterPerDay.fillFieldMap()

	return _totalOdorPollutantsSlaughterPerDay
}

type totalOdorPollutantsSlaughterPerDay struct {
	totalOdorPollutantsSlaughterPerDayDo totalOdorPollutantsSlaughterPerDayDo

	ALL                                 field.Asterisk
	ID                                  field.Uint
	CreatedAt                           field.Time
	UpdatedAt                           field.Time
	DeletedAt                           field.Field
	TimeStamp                           field.Time
	HouseNumber                         field.String
	TotalOdorPollutantsSlaughterPerDay1 field.Float64
	TotalOdorPollutantsSlaughterPerDay2 field.Float64
	TotalOdorPollutantsSlaughterPerDay3 field.Float64
	TotalOdorPollutantsSlaughterPerDay4 field.Float64

	fieldMap map[string]field.Expr
}

func (t totalOdorPollutantsSlaughterPerDay) Table(newTableName string) *totalOdorPollutantsSlaughterPerDay {
	t.totalOdorPollutantsSlaughterPerDayDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t totalOdorPollutantsSlaughterPerDay) As(alias string) *totalOdorPollutantsSlaughterPerDay {
	t.totalOdorPollutantsSlaughterPerDayDo.DO = *(t.totalOdorPollutantsSlaughterPerDayDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *totalOdorPollutantsSlaughterPerDay) updateTableName(table string) *totalOdorPollutantsSlaughterPerDay {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TimeStamp = field.NewTime(table, "time_stamp")
	t.HouseNumber = field.NewString(table, "house_number")
	t.TotalOdorPollutantsSlaughterPerDay1 = field.NewFloat64(table, "total_odor_pollutants_slaughter_per_day1")
	t.TotalOdorPollutantsSlaughterPerDay2 = field.NewFloat64(table, "total_odor_pollutants_slaughter_per_day2")
	t.TotalOdorPollutantsSlaughterPerDay3 = field.NewFloat64(table, "total_odor_pollutants_slaughter_per_day3")
	t.TotalOdorPollutantsSlaughterPerDay4 = field.NewFloat64(table, "total_odor_pollutants_slaughter_per_day4")

	t.fillFieldMap()

	return t
}

func (t *totalOdorPollutantsSlaughterPerDay) WithContext(ctx context.Context) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.totalOdorPollutantsSlaughterPerDayDo.WithContext(ctx)
}

func (t totalOdorPollutantsSlaughterPerDay) TableName() string {
	return t.totalOdorPollutantsSlaughterPerDayDo.TableName()
}

func (t totalOdorPollutantsSlaughterPerDay) Alias() string {
	return t.totalOdorPollutantsSlaughterPerDayDo.Alias()
}

func (t totalOdorPollutantsSlaughterPerDay) Columns(cols ...field.Expr) gen.Columns {
	return t.totalOdorPollutantsSlaughterPerDayDo.Columns(cols...)
}

func (t *totalOdorPollutantsSlaughterPerDay) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *totalOdorPollutantsSlaughterPerDay) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["time_stamp"] = t.TimeStamp
	t.fieldMap["house_number"] = t.HouseNumber
	t.fieldMap["total_odor_pollutants_slaughter_per_day1"] = t.TotalOdorPollutantsSlaughterPerDay1
	t.fieldMap["total_odor_pollutants_slaughter_per_day2"] = t.TotalOdorPollutantsSlaughterPerDay2
	t.fieldMap["total_odor_pollutants_slaughter_per_day3"] = t.TotalOdorPollutantsSlaughterPerDay3
	t.fieldMap["total_odor_pollutants_slaughter_per_day4"] = t.TotalOdorPollutantsSlaughterPerDay4
}

func (t totalOdorPollutantsSlaughterPerDay) clone(db *gorm.DB) totalOdorPollutantsSlaughterPerDay {
	t.totalOdorPollutantsSlaughterPerDayDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t totalOdorPollutantsSlaughterPerDay) replaceDB(db *gorm.DB) totalOdorPollutantsSlaughterPerDay {
	t.totalOdorPollutantsSlaughterPerDayDo.ReplaceDB(db)
	return t
}

type totalOdorPollutantsSlaughterPerDayDo struct{ gen.DO }

type ITotalOdorPollutantsSlaughterPerDayDo interface {
	gen.SubQuery
	Debug() ITotalOdorPollutantsSlaughterPerDayDo
	WithContext(ctx context.Context) ITotalOdorPollutantsSlaughterPerDayDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITotalOdorPollutantsSlaughterPerDayDo
	WriteDB() ITotalOdorPollutantsSlaughterPerDayDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITotalOdorPollutantsSlaughterPerDayDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITotalOdorPollutantsSlaughterPerDayDo
	Not(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo
	Or(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo
	Select(conds ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	Where(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo
	Order(conds ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	Distinct(cols ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	Omit(cols ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	Join(table schema.Tabler, on ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	Group(cols ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo
	Having(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo
	Limit(limit int) ITotalOdorPollutantsSlaughterPerDayDo
	Offset(offset int) ITotalOdorPollutantsSlaughterPerDayDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITotalOdorPollutantsSlaughterPerDayDo
	Unscoped() ITotalOdorPollutantsSlaughterPerDayDo
	Create(values ...*slaughter.TotalOdorPollutantsSlaughterPerDay) error
	CreateInBatches(values []*slaughter.TotalOdorPollutantsSlaughterPerDay, batchSize int) error
	Save(values ...*slaughter.TotalOdorPollutantsSlaughterPerDay) error
	First() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error)
	Take() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error)
	Last() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error)
	Find() ([]*slaughter.TotalOdorPollutantsSlaughterPerDay, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.TotalOdorPollutantsSlaughterPerDay, err error)
	FindInBatches(result *[]*slaughter.TotalOdorPollutantsSlaughterPerDay, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.TotalOdorPollutantsSlaughterPerDay) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITotalOdorPollutantsSlaughterPerDayDo
	Assign(attrs ...field.AssignExpr) ITotalOdorPollutantsSlaughterPerDayDo
	Joins(fields ...field.RelationField) ITotalOdorPollutantsSlaughterPerDayDo
	Preload(fields ...field.RelationField) ITotalOdorPollutantsSlaughterPerDayDo
	FirstOrInit() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error)
	FirstOrCreate() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error)
	FindByPage(offset int, limit int) (result []*slaughter.TotalOdorPollutantsSlaughterPerDay, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITotalOdorPollutantsSlaughterPerDayDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t totalOdorPollutantsSlaughterPerDayDo) Debug() ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Debug())
}

func (t totalOdorPollutantsSlaughterPerDayDo) WithContext(ctx context.Context) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t totalOdorPollutantsSlaughterPerDayDo) ReadDB() ITotalOdorPollutantsSlaughterPerDayDo {
	return t.Clauses(dbresolver.Read)
}

func (t totalOdorPollutantsSlaughterPerDayDo) WriteDB() ITotalOdorPollutantsSlaughterPerDayDo {
	return t.Clauses(dbresolver.Write)
}

func (t totalOdorPollutantsSlaughterPerDayDo) Session(config *gorm.Session) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Session(config))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Clauses(conds ...clause.Expression) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Returning(value interface{}, columns ...string) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Not(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Or(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Select(conds ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Where(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Order(conds ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Distinct(cols ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Omit(cols ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Join(table schema.Tabler, on ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) RightJoin(table schema.Tabler, on ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Group(cols ...field.Expr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Having(conds ...gen.Condition) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Limit(limit int) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Offset(offset int) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Unscoped() ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Unscoped())
}

func (t totalOdorPollutantsSlaughterPerDayDo) Create(values ...*slaughter.TotalOdorPollutantsSlaughterPerDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t totalOdorPollutantsSlaughterPerDayDo) CreateInBatches(values []*slaughter.TotalOdorPollutantsSlaughterPerDay, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t totalOdorPollutantsSlaughterPerDayDo) Save(values ...*slaughter.TotalOdorPollutantsSlaughterPerDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t totalOdorPollutantsSlaughterPerDayDo) First() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalOdorPollutantsSlaughterPerDay), nil
	}
}

func (t totalOdorPollutantsSlaughterPerDayDo) Take() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalOdorPollutantsSlaughterPerDay), nil
	}
}

func (t totalOdorPollutantsSlaughterPerDayDo) Last() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalOdorPollutantsSlaughterPerDay), nil
	}
}

func (t totalOdorPollutantsSlaughterPerDayDo) Find() ([]*slaughter.TotalOdorPollutantsSlaughterPerDay, error) {
	result, err := t.DO.Find()
	return result.([]*slaughter.TotalOdorPollutantsSlaughterPerDay), err
}

func (t totalOdorPollutantsSlaughterPerDayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.TotalOdorPollutantsSlaughterPerDay, err error) {
	buf := make([]*slaughter.TotalOdorPollutantsSlaughterPerDay, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t totalOdorPollutantsSlaughterPerDayDo) FindInBatches(result *[]*slaughter.TotalOdorPollutantsSlaughterPerDay, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t totalOdorPollutantsSlaughterPerDayDo) Attrs(attrs ...field.AssignExpr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Assign(attrs ...field.AssignExpr) ITotalOdorPollutantsSlaughterPerDayDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t totalOdorPollutantsSlaughterPerDayDo) Joins(fields ...field.RelationField) ITotalOdorPollutantsSlaughterPerDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t totalOdorPollutantsSlaughterPerDayDo) Preload(fields ...field.RelationField) ITotalOdorPollutantsSlaughterPerDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t totalOdorPollutantsSlaughterPerDayDo) FirstOrInit() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalOdorPollutantsSlaughterPerDay), nil
	}
}

func (t totalOdorPollutantsSlaughterPerDayDo) FirstOrCreate() (*slaughter.TotalOdorPollutantsSlaughterPerDay, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.TotalOdorPollutantsSlaughterPerDay), nil
	}
}

func (t totalOdorPollutantsSlaughterPerDayDo) FindByPage(offset int, limit int) (result []*slaughter.TotalOdorPollutantsSlaughterPerDay, count int64, err error) {
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

func (t totalOdorPollutantsSlaughterPerDayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t totalOdorPollutantsSlaughterPerDayDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t totalOdorPollutantsSlaughterPerDayDo) Delete(models ...*slaughter.TotalOdorPollutantsSlaughterPerDay) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *totalOdorPollutantsSlaughterPerDayDo) withDO(do gen.Dao) *totalOdorPollutantsSlaughterPerDayDo {
	t.DO = *do.(*gen.DO)
	return t
}
