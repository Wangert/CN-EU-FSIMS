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

func newBleedElectronic(db *gorm.DB, opts ...gen.DOOption) bleedElectronic {
	_bleedElectronic := bleedElectronic{}

	_bleedElectronic.bleedElectronicDo.UseDB(db, opts...)
	_bleedElectronic.bleedElectronicDo.UseModel(&slaughter.BleedElectronic{})

	tableName := _bleedElectronic.bleedElectronicDo.TableName()
	_bleedElectronic.ALL = field.NewAsterisk(tableName)
	_bleedElectronic.ID = field.NewUint(tableName, "id")
	_bleedElectronic.CreatedAt = field.NewTime(tableName, "created_at")
	_bleedElectronic.UpdatedAt = field.NewTime(tableName, "updated_at")
	_bleedElectronic.DeletedAt = field.NewField(tableName, "deleted_at")
	_bleedElectronic.SlaughterProcedureMonitoringDataID = field.NewUint(tableName, "slaughter_procedure_monitoring_data_id")
	_bleedElectronic.BleedElectronic1 = field.NewFloat64(tableName, "bleed_electronic1")
	_bleedElectronic.BleedElectronic2 = field.NewFloat64(tableName, "bleed_electronic2")
	_bleedElectronic.BleedElectronic3 = field.NewFloat64(tableName, "bleed_electronic3")
	_bleedElectronic.BleedElectronic4 = field.NewFloat64(tableName, "bleed_electronic4")
	_bleedElectronic.BleedElectronic5 = field.NewFloat64(tableName, "bleed_electronic5")

	_bleedElectronic.fillFieldMap()

	return _bleedElectronic
}

type bleedElectronic struct {
	bleedElectronicDo bleedElectronicDo

	ALL                                field.Asterisk
	ID                                 field.Uint
	CreatedAt                          field.Time
	UpdatedAt                          field.Time
	DeletedAt                          field.Field
	SlaughterProcedureMonitoringDataID field.Uint
	BleedElectronic1                   field.Float64
	BleedElectronic2                   field.Float64
	BleedElectronic3                   field.Float64
	BleedElectronic4                   field.Float64
	BleedElectronic5                   field.Float64

	fieldMap map[string]field.Expr
}

func (b bleedElectronic) Table(newTableName string) *bleedElectronic {
	b.bleedElectronicDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bleedElectronic) As(alias string) *bleedElectronic {
	b.bleedElectronicDo.DO = *(b.bleedElectronicDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bleedElectronic) updateTableName(table string) *bleedElectronic {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewUint(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.SlaughterProcedureMonitoringDataID = field.NewUint(table, "slaughter_procedure_monitoring_data_id")
	b.BleedElectronic1 = field.NewFloat64(table, "bleed_electronic1")
	b.BleedElectronic2 = field.NewFloat64(table, "bleed_electronic2")
	b.BleedElectronic3 = field.NewFloat64(table, "bleed_electronic3")
	b.BleedElectronic4 = field.NewFloat64(table, "bleed_electronic4")
	b.BleedElectronic5 = field.NewFloat64(table, "bleed_electronic5")

	b.fillFieldMap()

	return b
}

func (b *bleedElectronic) WithContext(ctx context.Context) IBleedElectronicDo {
	return b.bleedElectronicDo.WithContext(ctx)
}

func (b bleedElectronic) TableName() string { return b.bleedElectronicDo.TableName() }

func (b bleedElectronic) Alias() string { return b.bleedElectronicDo.Alias() }

func (b bleedElectronic) Columns(cols ...field.Expr) gen.Columns {
	return b.bleedElectronicDo.Columns(cols...)
}

func (b *bleedElectronic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bleedElectronic) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 10)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["slaughter_procedure_monitoring_data_id"] = b.SlaughterProcedureMonitoringDataID
	b.fieldMap["bleed_electronic1"] = b.BleedElectronic1
	b.fieldMap["bleed_electronic2"] = b.BleedElectronic2
	b.fieldMap["bleed_electronic3"] = b.BleedElectronic3
	b.fieldMap["bleed_electronic4"] = b.BleedElectronic4
	b.fieldMap["bleed_electronic5"] = b.BleedElectronic5
}

func (b bleedElectronic) clone(db *gorm.DB) bleedElectronic {
	b.bleedElectronicDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bleedElectronic) replaceDB(db *gorm.DB) bleedElectronic {
	b.bleedElectronicDo.ReplaceDB(db)
	return b
}

type bleedElectronicDo struct{ gen.DO }

type IBleedElectronicDo interface {
	gen.SubQuery
	Debug() IBleedElectronicDo
	WithContext(ctx context.Context) IBleedElectronicDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBleedElectronicDo
	WriteDB() IBleedElectronicDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBleedElectronicDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBleedElectronicDo
	Not(conds ...gen.Condition) IBleedElectronicDo
	Or(conds ...gen.Condition) IBleedElectronicDo
	Select(conds ...field.Expr) IBleedElectronicDo
	Where(conds ...gen.Condition) IBleedElectronicDo
	Order(conds ...field.Expr) IBleedElectronicDo
	Distinct(cols ...field.Expr) IBleedElectronicDo
	Omit(cols ...field.Expr) IBleedElectronicDo
	Join(table schema.Tabler, on ...field.Expr) IBleedElectronicDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBleedElectronicDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBleedElectronicDo
	Group(cols ...field.Expr) IBleedElectronicDo
	Having(conds ...gen.Condition) IBleedElectronicDo
	Limit(limit int) IBleedElectronicDo
	Offset(offset int) IBleedElectronicDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBleedElectronicDo
	Unscoped() IBleedElectronicDo
	Create(values ...*slaughter.BleedElectronic) error
	CreateInBatches(values []*slaughter.BleedElectronic, batchSize int) error
	Save(values ...*slaughter.BleedElectronic) error
	First() (*slaughter.BleedElectronic, error)
	Take() (*slaughter.BleedElectronic, error)
	Last() (*slaughter.BleedElectronic, error)
	Find() ([]*slaughter.BleedElectronic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.BleedElectronic, err error)
	FindInBatches(result *[]*slaughter.BleedElectronic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.BleedElectronic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBleedElectronicDo
	Assign(attrs ...field.AssignExpr) IBleedElectronicDo
	Joins(fields ...field.RelationField) IBleedElectronicDo
	Preload(fields ...field.RelationField) IBleedElectronicDo
	FirstOrInit() (*slaughter.BleedElectronic, error)
	FirstOrCreate() (*slaughter.BleedElectronic, error)
	FindByPage(offset int, limit int) (result []*slaughter.BleedElectronic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBleedElectronicDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bleedElectronicDo) Debug() IBleedElectronicDo {
	return b.withDO(b.DO.Debug())
}

func (b bleedElectronicDo) WithContext(ctx context.Context) IBleedElectronicDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bleedElectronicDo) ReadDB() IBleedElectronicDo {
	return b.Clauses(dbresolver.Read)
}

func (b bleedElectronicDo) WriteDB() IBleedElectronicDo {
	return b.Clauses(dbresolver.Write)
}

func (b bleedElectronicDo) Session(config *gorm.Session) IBleedElectronicDo {
	return b.withDO(b.DO.Session(config))
}

func (b bleedElectronicDo) Clauses(conds ...clause.Expression) IBleedElectronicDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bleedElectronicDo) Returning(value interface{}, columns ...string) IBleedElectronicDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bleedElectronicDo) Not(conds ...gen.Condition) IBleedElectronicDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bleedElectronicDo) Or(conds ...gen.Condition) IBleedElectronicDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bleedElectronicDo) Select(conds ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bleedElectronicDo) Where(conds ...gen.Condition) IBleedElectronicDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bleedElectronicDo) Order(conds ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bleedElectronicDo) Distinct(cols ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bleedElectronicDo) Omit(cols ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bleedElectronicDo) Join(table schema.Tabler, on ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bleedElectronicDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bleedElectronicDo) RightJoin(table schema.Tabler, on ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bleedElectronicDo) Group(cols ...field.Expr) IBleedElectronicDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bleedElectronicDo) Having(conds ...gen.Condition) IBleedElectronicDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bleedElectronicDo) Limit(limit int) IBleedElectronicDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bleedElectronicDo) Offset(offset int) IBleedElectronicDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bleedElectronicDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBleedElectronicDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bleedElectronicDo) Unscoped() IBleedElectronicDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bleedElectronicDo) Create(values ...*slaughter.BleedElectronic) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bleedElectronicDo) CreateInBatches(values []*slaughter.BleedElectronic, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bleedElectronicDo) Save(values ...*slaughter.BleedElectronic) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bleedElectronicDo) First() (*slaughter.BleedElectronic, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.BleedElectronic), nil
	}
}

func (b bleedElectronicDo) Take() (*slaughter.BleedElectronic, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.BleedElectronic), nil
	}
}

func (b bleedElectronicDo) Last() (*slaughter.BleedElectronic, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.BleedElectronic), nil
	}
}

func (b bleedElectronicDo) Find() ([]*slaughter.BleedElectronic, error) {
	result, err := b.DO.Find()
	return result.([]*slaughter.BleedElectronic), err
}

func (b bleedElectronicDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.BleedElectronic, err error) {
	buf := make([]*slaughter.BleedElectronic, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bleedElectronicDo) FindInBatches(result *[]*slaughter.BleedElectronic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bleedElectronicDo) Attrs(attrs ...field.AssignExpr) IBleedElectronicDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bleedElectronicDo) Assign(attrs ...field.AssignExpr) IBleedElectronicDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bleedElectronicDo) Joins(fields ...field.RelationField) IBleedElectronicDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bleedElectronicDo) Preload(fields ...field.RelationField) IBleedElectronicDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bleedElectronicDo) FirstOrInit() (*slaughter.BleedElectronic, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.BleedElectronic), nil
	}
}

func (b bleedElectronicDo) FirstOrCreate() (*slaughter.BleedElectronic, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.BleedElectronic), nil
	}
}

func (b bleedElectronicDo) FindByPage(offset int, limit int) (result []*slaughter.BleedElectronic, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bleedElectronicDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bleedElectronicDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bleedElectronicDo) Delete(models ...*slaughter.BleedElectronic) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bleedElectronicDo) withDO(do gen.Dao) *bleedElectronicDo {
	b.DO = *do.(*gen.DO)
	return b
}
