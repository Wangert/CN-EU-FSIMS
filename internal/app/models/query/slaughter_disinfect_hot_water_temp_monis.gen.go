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

func newSlaughterDisinfectHotWaterTempMoni(db *gorm.DB, opts ...gen.DOOption) slaughterDisinfectHotWaterTempMoni {
	_slaughterDisinfectHotWaterTempMoni := slaughterDisinfectHotWaterTempMoni{}

	_slaughterDisinfectHotWaterTempMoni.slaughterDisinfectHotWaterTempMoniDo.UseDB(db, opts...)
	_slaughterDisinfectHotWaterTempMoni.slaughterDisinfectHotWaterTempMoniDo.UseModel(&slaughter.SlaughterDisinfectHotWaterTempMoni{})

	tableName := _slaughterDisinfectHotWaterTempMoni.slaughterDisinfectHotWaterTempMoniDo.TableName()
	_slaughterDisinfectHotWaterTempMoni.ALL = field.NewAsterisk(tableName)
	_slaughterDisinfectHotWaterTempMoni.ID = field.NewUint(tableName, "id")
	_slaughterDisinfectHotWaterTempMoni.CreatedAt = field.NewTime(tableName, "created_at")
	_slaughterDisinfectHotWaterTempMoni.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaughterDisinfectHotWaterTempMoni.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaughterDisinfectHotWaterTempMoni.SlaughterProcedureMonitoringDataID = field.NewString(tableName, "slaughter_procedure_monitoring_data_id")
	_slaughterDisinfectHotWaterTempMoni.SlaughterDisinfectHotWaterTempMoni1 = field.NewFloat64(tableName, "slaughter_disinfect_hot_water_temp_moni1")
	_slaughterDisinfectHotWaterTempMoni.SlaughterDisinfectHotWaterTempMoni2 = field.NewFloat64(tableName, "slaughter_disinfect_hot_water_temp_moni2")
	_slaughterDisinfectHotWaterTempMoni.SlaughterDisinfectHotWaterTempMoni3 = field.NewFloat64(tableName, "slaughter_disinfect_hot_water_temp_moni3")
	_slaughterDisinfectHotWaterTempMoni.SlaughterDisinfectHotWaterTempMoni4 = field.NewFloat64(tableName, "slaughter_disinfect_hot_water_temp_moni4")
	_slaughterDisinfectHotWaterTempMoni.SlaughterDisinfectHotWaterTempMoni5 = field.NewFloat64(tableName, "slaughter_disinfect_hot_water_temp_moni5")
	_slaughterDisinfectHotWaterTempMoni.SlaughterDisinfectHotWaterTempMoni6 = field.NewFloat64(tableName, "slaughter_disinfect_hot_water_temp_moni6")

	_slaughterDisinfectHotWaterTempMoni.fillFieldMap()

	return _slaughterDisinfectHotWaterTempMoni
}

type slaughterDisinfectHotWaterTempMoni struct {
	slaughterDisinfectHotWaterTempMoniDo slaughterDisinfectHotWaterTempMoniDo

	ALL                                 field.Asterisk
	ID                                  field.Uint
	CreatedAt                           field.Time
	UpdatedAt                           field.Time
	DeletedAt                           field.Field
	SlaughterProcedureMonitoringDataID  field.String
	SlaughterDisinfectHotWaterTempMoni1 field.Float64
	SlaughterDisinfectHotWaterTempMoni2 field.Float64
	SlaughterDisinfectHotWaterTempMoni3 field.Float64
	SlaughterDisinfectHotWaterTempMoni4 field.Float64
	SlaughterDisinfectHotWaterTempMoni5 field.Float64
	SlaughterDisinfectHotWaterTempMoni6 field.Float64

	fieldMap map[string]field.Expr
}

func (s slaughterDisinfectHotWaterTempMoni) Table(newTableName string) *slaughterDisinfectHotWaterTempMoni {
	s.slaughterDisinfectHotWaterTempMoniDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaughterDisinfectHotWaterTempMoni) As(alias string) *slaughterDisinfectHotWaterTempMoni {
	s.slaughterDisinfectHotWaterTempMoniDo.DO = *(s.slaughterDisinfectHotWaterTempMoniDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaughterDisinfectHotWaterTempMoni) updateTableName(table string) *slaughterDisinfectHotWaterTempMoni {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.SlaughterProcedureMonitoringDataID = field.NewString(table, "slaughter_procedure_monitoring_data_id")
	s.SlaughterDisinfectHotWaterTempMoni1 = field.NewFloat64(table, "slaughter_disinfect_hot_water_temp_moni1")
	s.SlaughterDisinfectHotWaterTempMoni2 = field.NewFloat64(table, "slaughter_disinfect_hot_water_temp_moni2")
	s.SlaughterDisinfectHotWaterTempMoni3 = field.NewFloat64(table, "slaughter_disinfect_hot_water_temp_moni3")
	s.SlaughterDisinfectHotWaterTempMoni4 = field.NewFloat64(table, "slaughter_disinfect_hot_water_temp_moni4")
	s.SlaughterDisinfectHotWaterTempMoni5 = field.NewFloat64(table, "slaughter_disinfect_hot_water_temp_moni5")
	s.SlaughterDisinfectHotWaterTempMoni6 = field.NewFloat64(table, "slaughter_disinfect_hot_water_temp_moni6")

	s.fillFieldMap()

	return s
}

func (s *slaughterDisinfectHotWaterTempMoni) WithContext(ctx context.Context) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.slaughterDisinfectHotWaterTempMoniDo.WithContext(ctx)
}

func (s slaughterDisinfectHotWaterTempMoni) TableName() string {
	return s.slaughterDisinfectHotWaterTempMoniDo.TableName()
}

func (s slaughterDisinfectHotWaterTempMoni) Alias() string {
	return s.slaughterDisinfectHotWaterTempMoniDo.Alias()
}

func (s slaughterDisinfectHotWaterTempMoni) Columns(cols ...field.Expr) gen.Columns {
	return s.slaughterDisinfectHotWaterTempMoniDo.Columns(cols...)
}

func (s *slaughterDisinfectHotWaterTempMoni) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaughterDisinfectHotWaterTempMoni) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["slaughter_procedure_monitoring_data_id"] = s.SlaughterProcedureMonitoringDataID
	s.fieldMap["slaughter_disinfect_hot_water_temp_moni1"] = s.SlaughterDisinfectHotWaterTempMoni1
	s.fieldMap["slaughter_disinfect_hot_water_temp_moni2"] = s.SlaughterDisinfectHotWaterTempMoni2
	s.fieldMap["slaughter_disinfect_hot_water_temp_moni3"] = s.SlaughterDisinfectHotWaterTempMoni3
	s.fieldMap["slaughter_disinfect_hot_water_temp_moni4"] = s.SlaughterDisinfectHotWaterTempMoni4
	s.fieldMap["slaughter_disinfect_hot_water_temp_moni5"] = s.SlaughterDisinfectHotWaterTempMoni5
	s.fieldMap["slaughter_disinfect_hot_water_temp_moni6"] = s.SlaughterDisinfectHotWaterTempMoni6
}

func (s slaughterDisinfectHotWaterTempMoni) clone(db *gorm.DB) slaughterDisinfectHotWaterTempMoni {
	s.slaughterDisinfectHotWaterTempMoniDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaughterDisinfectHotWaterTempMoni) replaceDB(db *gorm.DB) slaughterDisinfectHotWaterTempMoni {
	s.slaughterDisinfectHotWaterTempMoniDo.ReplaceDB(db)
	return s
}

type slaughterDisinfectHotWaterTempMoniDo struct{ gen.DO }

type ISlaughterDisinfectHotWaterTempMoniDo interface {
	gen.SubQuery
	Debug() ISlaughterDisinfectHotWaterTempMoniDo
	WithContext(ctx context.Context) ISlaughterDisinfectHotWaterTempMoniDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaughterDisinfectHotWaterTempMoniDo
	WriteDB() ISlaughterDisinfectHotWaterTempMoniDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaughterDisinfectHotWaterTempMoniDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaughterDisinfectHotWaterTempMoniDo
	Not(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo
	Or(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo
	Select(conds ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	Where(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo
	Order(conds ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	Distinct(cols ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	Omit(cols ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	Join(table schema.Tabler, on ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	Group(cols ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo
	Having(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo
	Limit(limit int) ISlaughterDisinfectHotWaterTempMoniDo
	Offset(offset int) ISlaughterDisinfectHotWaterTempMoniDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterDisinfectHotWaterTempMoniDo
	Unscoped() ISlaughterDisinfectHotWaterTempMoniDo
	Create(values ...*slaughter.SlaughterDisinfectHotWaterTempMoni) error
	CreateInBatches(values []*slaughter.SlaughterDisinfectHotWaterTempMoni, batchSize int) error
	Save(values ...*slaughter.SlaughterDisinfectHotWaterTempMoni) error
	First() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error)
	Take() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error)
	Last() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error)
	Find() ([]*slaughter.SlaughterDisinfectHotWaterTempMoni, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterDisinfectHotWaterTempMoni, err error)
	FindInBatches(result *[]*slaughter.SlaughterDisinfectHotWaterTempMoni, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaughterDisinfectHotWaterTempMoni) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaughterDisinfectHotWaterTempMoniDo
	Assign(attrs ...field.AssignExpr) ISlaughterDisinfectHotWaterTempMoniDo
	Joins(fields ...field.RelationField) ISlaughterDisinfectHotWaterTempMoniDo
	Preload(fields ...field.RelationField) ISlaughterDisinfectHotWaterTempMoniDo
	FirstOrInit() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error)
	FirstOrCreate() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaughterDisinfectHotWaterTempMoni, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaughterDisinfectHotWaterTempMoniDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaughterDisinfectHotWaterTempMoniDo) Debug() ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Debug())
}

func (s slaughterDisinfectHotWaterTempMoniDo) WithContext(ctx context.Context) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaughterDisinfectHotWaterTempMoniDo) ReadDB() ISlaughterDisinfectHotWaterTempMoniDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaughterDisinfectHotWaterTempMoniDo) WriteDB() ISlaughterDisinfectHotWaterTempMoniDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaughterDisinfectHotWaterTempMoniDo) Session(config *gorm.Session) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Clauses(conds ...clause.Expression) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Returning(value interface{}, columns ...string) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Not(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Or(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Select(conds ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Where(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Order(conds ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Distinct(cols ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Omit(cols ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Join(table schema.Tabler, on ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Group(cols ...field.Expr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Having(conds ...gen.Condition) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Limit(limit int) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Offset(offset int) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Unscoped() ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaughterDisinfectHotWaterTempMoniDo) Create(values ...*slaughter.SlaughterDisinfectHotWaterTempMoni) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaughterDisinfectHotWaterTempMoniDo) CreateInBatches(values []*slaughter.SlaughterDisinfectHotWaterTempMoni, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaughterDisinfectHotWaterTempMoniDo) Save(values ...*slaughter.SlaughterDisinfectHotWaterTempMoni) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaughterDisinfectHotWaterTempMoniDo) First() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterDisinfectHotWaterTempMoni), nil
	}
}

func (s slaughterDisinfectHotWaterTempMoniDo) Take() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterDisinfectHotWaterTempMoni), nil
	}
}

func (s slaughterDisinfectHotWaterTempMoniDo) Last() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterDisinfectHotWaterTempMoni), nil
	}
}

func (s slaughterDisinfectHotWaterTempMoniDo) Find() ([]*slaughter.SlaughterDisinfectHotWaterTempMoni, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaughterDisinfectHotWaterTempMoni), err
}

func (s slaughterDisinfectHotWaterTempMoniDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterDisinfectHotWaterTempMoni, err error) {
	buf := make([]*slaughter.SlaughterDisinfectHotWaterTempMoni, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaughterDisinfectHotWaterTempMoniDo) FindInBatches(result *[]*slaughter.SlaughterDisinfectHotWaterTempMoni, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaughterDisinfectHotWaterTempMoniDo) Attrs(attrs ...field.AssignExpr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Assign(attrs ...field.AssignExpr) ISlaughterDisinfectHotWaterTempMoniDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaughterDisinfectHotWaterTempMoniDo) Joins(fields ...field.RelationField) ISlaughterDisinfectHotWaterTempMoniDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaughterDisinfectHotWaterTempMoniDo) Preload(fields ...field.RelationField) ISlaughterDisinfectHotWaterTempMoniDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaughterDisinfectHotWaterTempMoniDo) FirstOrInit() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterDisinfectHotWaterTempMoni), nil
	}
}

func (s slaughterDisinfectHotWaterTempMoniDo) FirstOrCreate() (*slaughter.SlaughterDisinfectHotWaterTempMoni, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterDisinfectHotWaterTempMoni), nil
	}
}

func (s slaughterDisinfectHotWaterTempMoniDo) FindByPage(offset int, limit int) (result []*slaughter.SlaughterDisinfectHotWaterTempMoni, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s slaughterDisinfectHotWaterTempMoniDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaughterDisinfectHotWaterTempMoniDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaughterDisinfectHotWaterTempMoniDo) Delete(models ...*slaughter.SlaughterDisinfectHotWaterTempMoni) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaughterDisinfectHotWaterTempMoniDo) withDO(do gen.Dao) *slaughterDisinfectHotWaterTempMoniDo {
	s.DO = *do.(*gen.DO)
	return s
}
