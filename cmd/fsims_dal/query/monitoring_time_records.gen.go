// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newMonitoringTimeRecord(db *gorm.DB, opts ...gen.DOOption) monitoringTimeRecord {
	_monitoringTimeRecord := monitoringTimeRecord{}

	_monitoringTimeRecord.monitoringTimeRecordDo.UseDB(db, opts...)
	_monitoringTimeRecord.monitoringTimeRecordDo.UseModel(&models.MonitoringTimeRecord{})

	tableName := _monitoringTimeRecord.monitoringTimeRecordDo.TableName()
	_monitoringTimeRecord.ALL = field.NewAsterisk(tableName)
	_monitoringTimeRecord.ID = field.NewUint(tableName, "id")
	_monitoringTimeRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_monitoringTimeRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_monitoringTimeRecord.DeletedAt = field.NewField(tableName, "deleted_at")
	_monitoringTimeRecord.IndexName = field.NewString(tableName, "index_name")
	_monitoringTimeRecord.LastTime = field.NewTime(tableName, "last_time")

	_monitoringTimeRecord.fillFieldMap()

	return _monitoringTimeRecord
}

type monitoringTimeRecord struct {
	monitoringTimeRecordDo monitoringTimeRecordDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	IndexName field.String
	LastTime  field.Time

	fieldMap map[string]field.Expr
}

func (m monitoringTimeRecord) Table(newTableName string) *monitoringTimeRecord {
	m.monitoringTimeRecordDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m monitoringTimeRecord) As(alias string) *monitoringTimeRecord {
	m.monitoringTimeRecordDo.DO = *(m.monitoringTimeRecordDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *monitoringTimeRecord) updateTableName(table string) *monitoringTimeRecord {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.IndexName = field.NewString(table, "index_name")
	m.LastTime = field.NewTime(table, "last_time")

	m.fillFieldMap()

	return m
}

func (m *monitoringTimeRecord) WithContext(ctx context.Context) IMonitoringTimeRecordDo {
	return m.monitoringTimeRecordDo.WithContext(ctx)
}

func (m monitoringTimeRecord) TableName() string { return m.monitoringTimeRecordDo.TableName() }

func (m monitoringTimeRecord) Alias() string { return m.monitoringTimeRecordDo.Alias() }

func (m monitoringTimeRecord) Columns(cols ...field.Expr) gen.Columns {
	return m.monitoringTimeRecordDo.Columns(cols...)
}

func (m *monitoringTimeRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *monitoringTimeRecord) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["index_name"] = m.IndexName
	m.fieldMap["last_time"] = m.LastTime
}

func (m monitoringTimeRecord) clone(db *gorm.DB) monitoringTimeRecord {
	m.monitoringTimeRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m monitoringTimeRecord) replaceDB(db *gorm.DB) monitoringTimeRecord {
	m.monitoringTimeRecordDo.ReplaceDB(db)
	return m
}

type monitoringTimeRecordDo struct{ gen.DO }

type IMonitoringTimeRecordDo interface {
	gen.SubQuery
	Debug() IMonitoringTimeRecordDo
	WithContext(ctx context.Context) IMonitoringTimeRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMonitoringTimeRecordDo
	WriteDB() IMonitoringTimeRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMonitoringTimeRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMonitoringTimeRecordDo
	Not(conds ...gen.Condition) IMonitoringTimeRecordDo
	Or(conds ...gen.Condition) IMonitoringTimeRecordDo
	Select(conds ...field.Expr) IMonitoringTimeRecordDo
	Where(conds ...gen.Condition) IMonitoringTimeRecordDo
	Order(conds ...field.Expr) IMonitoringTimeRecordDo
	Distinct(cols ...field.Expr) IMonitoringTimeRecordDo
	Omit(cols ...field.Expr) IMonitoringTimeRecordDo
	Join(table schema.Tabler, on ...field.Expr) IMonitoringTimeRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMonitoringTimeRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMonitoringTimeRecordDo
	Group(cols ...field.Expr) IMonitoringTimeRecordDo
	Having(conds ...gen.Condition) IMonitoringTimeRecordDo
	Limit(limit int) IMonitoringTimeRecordDo
	Offset(offset int) IMonitoringTimeRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMonitoringTimeRecordDo
	Unscoped() IMonitoringTimeRecordDo
	Create(values ...*models.MonitoringTimeRecord) error
	CreateInBatches(values []*models.MonitoringTimeRecord, batchSize int) error
	Save(values ...*models.MonitoringTimeRecord) error
	First() (*models.MonitoringTimeRecord, error)
	Take() (*models.MonitoringTimeRecord, error)
	Last() (*models.MonitoringTimeRecord, error)
	Find() ([]*models.MonitoringTimeRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.MonitoringTimeRecord, err error)
	FindInBatches(result *[]*models.MonitoringTimeRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.MonitoringTimeRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMonitoringTimeRecordDo
	Assign(attrs ...field.AssignExpr) IMonitoringTimeRecordDo
	Joins(fields ...field.RelationField) IMonitoringTimeRecordDo
	Preload(fields ...field.RelationField) IMonitoringTimeRecordDo
	FirstOrInit() (*models.MonitoringTimeRecord, error)
	FirstOrCreate() (*models.MonitoringTimeRecord, error)
	FindByPage(offset int, limit int) (result []*models.MonitoringTimeRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMonitoringTimeRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m monitoringTimeRecordDo) Debug() IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Debug())
}

func (m monitoringTimeRecordDo) WithContext(ctx context.Context) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m monitoringTimeRecordDo) ReadDB() IMonitoringTimeRecordDo {
	return m.Clauses(dbresolver.Read)
}

func (m monitoringTimeRecordDo) WriteDB() IMonitoringTimeRecordDo {
	return m.Clauses(dbresolver.Write)
}

func (m monitoringTimeRecordDo) Session(config *gorm.Session) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Session(config))
}

func (m monitoringTimeRecordDo) Clauses(conds ...clause.Expression) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m monitoringTimeRecordDo) Returning(value interface{}, columns ...string) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m monitoringTimeRecordDo) Not(conds ...gen.Condition) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m monitoringTimeRecordDo) Or(conds ...gen.Condition) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m monitoringTimeRecordDo) Select(conds ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m monitoringTimeRecordDo) Where(conds ...gen.Condition) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m monitoringTimeRecordDo) Order(conds ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m monitoringTimeRecordDo) Distinct(cols ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m monitoringTimeRecordDo) Omit(cols ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m monitoringTimeRecordDo) Join(table schema.Tabler, on ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m monitoringTimeRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m monitoringTimeRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m monitoringTimeRecordDo) Group(cols ...field.Expr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m monitoringTimeRecordDo) Having(conds ...gen.Condition) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m monitoringTimeRecordDo) Limit(limit int) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m monitoringTimeRecordDo) Offset(offset int) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m monitoringTimeRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m monitoringTimeRecordDo) Unscoped() IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Unscoped())
}

func (m monitoringTimeRecordDo) Create(values ...*models.MonitoringTimeRecord) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m monitoringTimeRecordDo) CreateInBatches(values []*models.MonitoringTimeRecord, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m monitoringTimeRecordDo) Save(values ...*models.MonitoringTimeRecord) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m monitoringTimeRecordDo) First() (*models.MonitoringTimeRecord, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.MonitoringTimeRecord), nil
	}
}

func (m monitoringTimeRecordDo) Take() (*models.MonitoringTimeRecord, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.MonitoringTimeRecord), nil
	}
}

func (m monitoringTimeRecordDo) Last() (*models.MonitoringTimeRecord, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.MonitoringTimeRecord), nil
	}
}

func (m monitoringTimeRecordDo) Find() ([]*models.MonitoringTimeRecord, error) {
	result, err := m.DO.Find()
	return result.([]*models.MonitoringTimeRecord), err
}

func (m monitoringTimeRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.MonitoringTimeRecord, err error) {
	buf := make([]*models.MonitoringTimeRecord, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m monitoringTimeRecordDo) FindInBatches(result *[]*models.MonitoringTimeRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m monitoringTimeRecordDo) Attrs(attrs ...field.AssignExpr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m monitoringTimeRecordDo) Assign(attrs ...field.AssignExpr) IMonitoringTimeRecordDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m monitoringTimeRecordDo) Joins(fields ...field.RelationField) IMonitoringTimeRecordDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m monitoringTimeRecordDo) Preload(fields ...field.RelationField) IMonitoringTimeRecordDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m monitoringTimeRecordDo) FirstOrInit() (*models.MonitoringTimeRecord, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.MonitoringTimeRecord), nil
	}
}

func (m monitoringTimeRecordDo) FirstOrCreate() (*models.MonitoringTimeRecord, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.MonitoringTimeRecord), nil
	}
}

func (m monitoringTimeRecordDo) FindByPage(offset int, limit int) (result []*models.MonitoringTimeRecord, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m monitoringTimeRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m monitoringTimeRecordDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m monitoringTimeRecordDo) Delete(models ...*models.MonitoringTimeRecord) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *monitoringTimeRecordDo) withDO(do gen.Dao) *monitoringTimeRecordDo {
	m.DO = *do.(*gen.DO)
	return m
}