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

func newToNumGermMon(db *gorm.DB, opts ...gen.DOOption) toNumGermMon {
	_toNumGermMon := toNumGermMon{}

	_toNumGermMon.toNumGermMonDo.UseDB(db, opts...)
	_toNumGermMon.toNumGermMonDo.UseModel(&slaughter.ToNumGermMon{})

	tableName := _toNumGermMon.toNumGermMonDo.TableName()
	_toNumGermMon.ALL = field.NewAsterisk(tableName)
	_toNumGermMon.ID = field.NewUint(tableName, "id")
	_toNumGermMon.CreatedAt = field.NewTime(tableName, "created_at")
	_toNumGermMon.UpdatedAt = field.NewTime(tableName, "updated_at")
	_toNumGermMon.DeletedAt = field.NewField(tableName, "deleted_at")
	_toNumGermMon.SlaughterProcedureMonitoringDataID = field.NewString(tableName, "slaughter_procedure_monitoring_data_id")
	_toNumGermMon.ToNumGermMon1 = field.NewFloat64(tableName, "to_num_germ_mon1")
	_toNumGermMon.ToNumGermMon2 = field.NewFloat64(tableName, "to_num_germ_mon2")
	_toNumGermMon.ToNumGermMon3 = field.NewFloat64(tableName, "to_num_germ_mon3")
	_toNumGermMon.ToNumGermMon4 = field.NewFloat64(tableName, "to_num_germ_mon4")
	_toNumGermMon.ToNumGermMon5 = field.NewFloat64(tableName, "to_num_germ_mon5")
	_toNumGermMon.ToNumGermMon6 = field.NewFloat64(tableName, "to_num_germ_mon6")
	_toNumGermMon.ToNumGermMon7 = field.NewFloat64(tableName, "to_num_germ_mon7")
	_toNumGermMon.ToNumGermMon8 = field.NewFloat64(tableName, "to_num_germ_mon8")

	_toNumGermMon.fillFieldMap()

	return _toNumGermMon
}

type toNumGermMon struct {
	toNumGermMonDo toNumGermMonDo

	ALL                                field.Asterisk
	ID                                 field.Uint
	CreatedAt                          field.Time
	UpdatedAt                          field.Time
	DeletedAt                          field.Field
	SlaughterProcedureMonitoringDataID field.String
	ToNumGermMon1                      field.Float64
	ToNumGermMon2                      field.Float64
	ToNumGermMon3                      field.Float64
	ToNumGermMon4                      field.Float64
	ToNumGermMon5                      field.Float64
	ToNumGermMon6                      field.Float64
	ToNumGermMon7                      field.Float64
	ToNumGermMon8                      field.Float64

	fieldMap map[string]field.Expr
}

func (t toNumGermMon) Table(newTableName string) *toNumGermMon {
	t.toNumGermMonDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t toNumGermMon) As(alias string) *toNumGermMon {
	t.toNumGermMonDo.DO = *(t.toNumGermMonDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *toNumGermMon) updateTableName(table string) *toNumGermMon {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.SlaughterProcedureMonitoringDataID = field.NewString(table, "slaughter_procedure_monitoring_data_id")
	t.ToNumGermMon1 = field.NewFloat64(table, "to_num_germ_mon1")
	t.ToNumGermMon2 = field.NewFloat64(table, "to_num_germ_mon2")
	t.ToNumGermMon3 = field.NewFloat64(table, "to_num_germ_mon3")
	t.ToNumGermMon4 = field.NewFloat64(table, "to_num_germ_mon4")
	t.ToNumGermMon5 = field.NewFloat64(table, "to_num_germ_mon5")
	t.ToNumGermMon6 = field.NewFloat64(table, "to_num_germ_mon6")
	t.ToNumGermMon7 = field.NewFloat64(table, "to_num_germ_mon7")
	t.ToNumGermMon8 = field.NewFloat64(table, "to_num_germ_mon8")

	t.fillFieldMap()

	return t
}

func (t *toNumGermMon) WithContext(ctx context.Context) IToNumGermMonDo {
	return t.toNumGermMonDo.WithContext(ctx)
}

func (t toNumGermMon) TableName() string { return t.toNumGermMonDo.TableName() }

func (t toNumGermMon) Alias() string { return t.toNumGermMonDo.Alias() }

func (t toNumGermMon) Columns(cols ...field.Expr) gen.Columns {
	return t.toNumGermMonDo.Columns(cols...)
}

func (t *toNumGermMon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *toNumGermMon) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 13)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["slaughter_procedure_monitoring_data_id"] = t.SlaughterProcedureMonitoringDataID
	t.fieldMap["to_num_germ_mon1"] = t.ToNumGermMon1
	t.fieldMap["to_num_germ_mon2"] = t.ToNumGermMon2
	t.fieldMap["to_num_germ_mon3"] = t.ToNumGermMon3
	t.fieldMap["to_num_germ_mon4"] = t.ToNumGermMon4
	t.fieldMap["to_num_germ_mon5"] = t.ToNumGermMon5
	t.fieldMap["to_num_germ_mon6"] = t.ToNumGermMon6
	t.fieldMap["to_num_germ_mon7"] = t.ToNumGermMon7
	t.fieldMap["to_num_germ_mon8"] = t.ToNumGermMon8
}

func (t toNumGermMon) clone(db *gorm.DB) toNumGermMon {
	t.toNumGermMonDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t toNumGermMon) replaceDB(db *gorm.DB) toNumGermMon {
	t.toNumGermMonDo.ReplaceDB(db)
	return t
}

type toNumGermMonDo struct{ gen.DO }

type IToNumGermMonDo interface {
	gen.SubQuery
	Debug() IToNumGermMonDo
	WithContext(ctx context.Context) IToNumGermMonDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IToNumGermMonDo
	WriteDB() IToNumGermMonDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IToNumGermMonDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IToNumGermMonDo
	Not(conds ...gen.Condition) IToNumGermMonDo
	Or(conds ...gen.Condition) IToNumGermMonDo
	Select(conds ...field.Expr) IToNumGermMonDo
	Where(conds ...gen.Condition) IToNumGermMonDo
	Order(conds ...field.Expr) IToNumGermMonDo
	Distinct(cols ...field.Expr) IToNumGermMonDo
	Omit(cols ...field.Expr) IToNumGermMonDo
	Join(table schema.Tabler, on ...field.Expr) IToNumGermMonDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IToNumGermMonDo
	RightJoin(table schema.Tabler, on ...field.Expr) IToNumGermMonDo
	Group(cols ...field.Expr) IToNumGermMonDo
	Having(conds ...gen.Condition) IToNumGermMonDo
	Limit(limit int) IToNumGermMonDo
	Offset(offset int) IToNumGermMonDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IToNumGermMonDo
	Unscoped() IToNumGermMonDo
	Create(values ...*slaughter.ToNumGermMon) error
	CreateInBatches(values []*slaughter.ToNumGermMon, batchSize int) error
	Save(values ...*slaughter.ToNumGermMon) error
	First() (*slaughter.ToNumGermMon, error)
	Take() (*slaughter.ToNumGermMon, error)
	Last() (*slaughter.ToNumGermMon, error)
	Find() ([]*slaughter.ToNumGermMon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.ToNumGermMon, err error)
	FindInBatches(result *[]*slaughter.ToNumGermMon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.ToNumGermMon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IToNumGermMonDo
	Assign(attrs ...field.AssignExpr) IToNumGermMonDo
	Joins(fields ...field.RelationField) IToNumGermMonDo
	Preload(fields ...field.RelationField) IToNumGermMonDo
	FirstOrInit() (*slaughter.ToNumGermMon, error)
	FirstOrCreate() (*slaughter.ToNumGermMon, error)
	FindByPage(offset int, limit int) (result []*slaughter.ToNumGermMon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IToNumGermMonDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t toNumGermMonDo) Debug() IToNumGermMonDo {
	return t.withDO(t.DO.Debug())
}

func (t toNumGermMonDo) WithContext(ctx context.Context) IToNumGermMonDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t toNumGermMonDo) ReadDB() IToNumGermMonDo {
	return t.Clauses(dbresolver.Read)
}

func (t toNumGermMonDo) WriteDB() IToNumGermMonDo {
	return t.Clauses(dbresolver.Write)
}

func (t toNumGermMonDo) Session(config *gorm.Session) IToNumGermMonDo {
	return t.withDO(t.DO.Session(config))
}

func (t toNumGermMonDo) Clauses(conds ...clause.Expression) IToNumGermMonDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t toNumGermMonDo) Returning(value interface{}, columns ...string) IToNumGermMonDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t toNumGermMonDo) Not(conds ...gen.Condition) IToNumGermMonDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t toNumGermMonDo) Or(conds ...gen.Condition) IToNumGermMonDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t toNumGermMonDo) Select(conds ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t toNumGermMonDo) Where(conds ...gen.Condition) IToNumGermMonDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t toNumGermMonDo) Order(conds ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t toNumGermMonDo) Distinct(cols ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t toNumGermMonDo) Omit(cols ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t toNumGermMonDo) Join(table schema.Tabler, on ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t toNumGermMonDo) LeftJoin(table schema.Tabler, on ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t toNumGermMonDo) RightJoin(table schema.Tabler, on ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t toNumGermMonDo) Group(cols ...field.Expr) IToNumGermMonDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t toNumGermMonDo) Having(conds ...gen.Condition) IToNumGermMonDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t toNumGermMonDo) Limit(limit int) IToNumGermMonDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t toNumGermMonDo) Offset(offset int) IToNumGermMonDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t toNumGermMonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IToNumGermMonDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t toNumGermMonDo) Unscoped() IToNumGermMonDo {
	return t.withDO(t.DO.Unscoped())
}

func (t toNumGermMonDo) Create(values ...*slaughter.ToNumGermMon) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t toNumGermMonDo) CreateInBatches(values []*slaughter.ToNumGermMon, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t toNumGermMonDo) Save(values ...*slaughter.ToNumGermMon) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t toNumGermMonDo) First() (*slaughter.ToNumGermMon, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.ToNumGermMon), nil
	}
}

func (t toNumGermMonDo) Take() (*slaughter.ToNumGermMon, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.ToNumGermMon), nil
	}
}

func (t toNumGermMonDo) Last() (*slaughter.ToNumGermMon, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.ToNumGermMon), nil
	}
}

func (t toNumGermMonDo) Find() ([]*slaughter.ToNumGermMon, error) {
	result, err := t.DO.Find()
	return result.([]*slaughter.ToNumGermMon), err
}

func (t toNumGermMonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.ToNumGermMon, err error) {
	buf := make([]*slaughter.ToNumGermMon, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t toNumGermMonDo) FindInBatches(result *[]*slaughter.ToNumGermMon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t toNumGermMonDo) Attrs(attrs ...field.AssignExpr) IToNumGermMonDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t toNumGermMonDo) Assign(attrs ...field.AssignExpr) IToNumGermMonDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t toNumGermMonDo) Joins(fields ...field.RelationField) IToNumGermMonDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t toNumGermMonDo) Preload(fields ...field.RelationField) IToNumGermMonDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t toNumGermMonDo) FirstOrInit() (*slaughter.ToNumGermMon, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.ToNumGermMon), nil
	}
}

func (t toNumGermMonDo) FirstOrCreate() (*slaughter.ToNumGermMon, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.ToNumGermMon), nil
	}
}

func (t toNumGermMonDo) FindByPage(offset int, limit int) (result []*slaughter.ToNumGermMon, count int64, err error) {
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

func (t toNumGermMonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t toNumGermMonDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t toNumGermMonDo) Delete(models ...*slaughter.ToNumGermMon) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *toNumGermMonDo) withDO(do gen.Dao) *toNumGermMonDo {
	t.DO = *do.(*gen.DO)
	return t
}
