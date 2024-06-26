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

func newPastureWasteDischarge(db *gorm.DB, opts ...gen.DOOption) pastureWasteDischarge {
	_pastureWasteDischarge := pastureWasteDischarge{}

	_pastureWasteDischarge.pastureWasteDischargeDo.UseDB(db, opts...)
	_pastureWasteDischarge.pastureWasteDischargeDo.UseModel(&pasture.PastureWasteDischarge{})

	tableName := _pastureWasteDischarge.pastureWasteDischargeDo.TableName()
	_pastureWasteDischarge.ALL = field.NewAsterisk(tableName)
	_pastureWasteDischarge.ID = field.NewUint(tableName, "id")
	_pastureWasteDischarge.CreatedAt = field.NewTime(tableName, "created_at")
	_pastureWasteDischarge.UpdatedAt = field.NewTime(tableName, "updated_at")
	_pastureWasteDischarge.DeletedAt = field.NewField(tableName, "deleted_at")
	_pastureWasteDischarge.WaterFivedayBOD = field.NewFloat64(tableName, "water_fiveday_bod")
	_pastureWasteDischarge.WaterChemicalOxygen = field.NewFloat64(tableName, "water_chemical_oxygen")
	_pastureWasteDischarge.WaterAmmoniaNitrogen = field.NewFloat64(tableName, "water_ammonia_nitrogen")
	_pastureWasteDischarge.WaterPhosphorus = field.NewFloat64(tableName, "water_phosphorus")
	_pastureWasteDischarge.WaterSuspendedMatter = field.NewFloat64(tableName, "water_suspended_matter")
	_pastureWasteDischarge.WaterFecalColiform = field.NewUint(tableName, "water_fecal_coliform")
	_pastureWasteDischarge.WaterAO = field.NewFloat64(tableName, "water_ao")
	_pastureWasteDischarge.WasteSlagFecalColiform = field.NewUint(tableName, "waste_slag_fecal_coliform")
	_pastureWasteDischarge.WasteSlagAOMortalityRate = field.NewFloat64(tableName, "waste_slag_ao_mortality_rate")
	_pastureWasteDischarge.O3Concentration = field.NewFloat64(tableName, "o3_concentration")
	_pastureWasteDischarge.PasPID = field.NewString(tableName, "pas_p_id")

	_pastureWasteDischarge.fillFieldMap()

	return _pastureWasteDischarge
}

type pastureWasteDischarge struct {
	pastureWasteDischargeDo pastureWasteDischargeDo

	ALL                      field.Asterisk
	ID                       field.Uint
	CreatedAt                field.Time
	UpdatedAt                field.Time
	DeletedAt                field.Field
	WaterFivedayBOD          field.Float64
	WaterChemicalOxygen      field.Float64
	WaterAmmoniaNitrogen     field.Float64
	WaterPhosphorus          field.Float64
	WaterSuspendedMatter     field.Float64
	WaterFecalColiform       field.Uint
	WaterAO                  field.Float64
	WasteSlagFecalColiform   field.Uint
	WasteSlagAOMortalityRate field.Float64
	O3Concentration          field.Float64
	PasPID                   field.String

	fieldMap map[string]field.Expr
}

func (p pastureWasteDischarge) Table(newTableName string) *pastureWasteDischarge {
	p.pastureWasteDischargeDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pastureWasteDischarge) As(alias string) *pastureWasteDischarge {
	p.pastureWasteDischargeDo.DO = *(p.pastureWasteDischargeDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pastureWasteDischarge) updateTableName(table string) *pastureWasteDischarge {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.WaterFivedayBOD = field.NewFloat64(table, "water_fiveday_bod")
	p.WaterChemicalOxygen = field.NewFloat64(table, "water_chemical_oxygen")
	p.WaterAmmoniaNitrogen = field.NewFloat64(table, "water_ammonia_nitrogen")
	p.WaterPhosphorus = field.NewFloat64(table, "water_phosphorus")
	p.WaterSuspendedMatter = field.NewFloat64(table, "water_suspended_matter")
	p.WaterFecalColiform = field.NewUint(table, "water_fecal_coliform")
	p.WaterAO = field.NewFloat64(table, "water_ao")
	p.WasteSlagFecalColiform = field.NewUint(table, "waste_slag_fecal_coliform")
	p.WasteSlagAOMortalityRate = field.NewFloat64(table, "waste_slag_ao_mortality_rate")
	p.O3Concentration = field.NewFloat64(table, "o3_concentration")
	p.PasPID = field.NewString(table, "pas_p_id")

	p.fillFieldMap()

	return p
}

func (p *pastureWasteDischarge) WithContext(ctx context.Context) IPastureWasteDischargeDo {
	return p.pastureWasteDischargeDo.WithContext(ctx)
}

func (p pastureWasteDischarge) TableName() string { return p.pastureWasteDischargeDo.TableName() }

func (p pastureWasteDischarge) Alias() string { return p.pastureWasteDischargeDo.Alias() }

func (p pastureWasteDischarge) Columns(cols ...field.Expr) gen.Columns {
	return p.pastureWasteDischargeDo.Columns(cols...)
}

func (p *pastureWasteDischarge) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pastureWasteDischarge) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 15)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["water_fiveday_bod"] = p.WaterFivedayBOD
	p.fieldMap["water_chemical_oxygen"] = p.WaterChemicalOxygen
	p.fieldMap["water_ammonia_nitrogen"] = p.WaterAmmoniaNitrogen
	p.fieldMap["water_phosphorus"] = p.WaterPhosphorus
	p.fieldMap["water_suspended_matter"] = p.WaterSuspendedMatter
	p.fieldMap["water_fecal_coliform"] = p.WaterFecalColiform
	p.fieldMap["water_ao"] = p.WaterAO
	p.fieldMap["waste_slag_fecal_coliform"] = p.WasteSlagFecalColiform
	p.fieldMap["waste_slag_ao_mortality_rate"] = p.WasteSlagAOMortalityRate
	p.fieldMap["o3_concentration"] = p.O3Concentration
	p.fieldMap["pas_p_id"] = p.PasPID
}

func (p pastureWasteDischarge) clone(db *gorm.DB) pastureWasteDischarge {
	p.pastureWasteDischargeDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pastureWasteDischarge) replaceDB(db *gorm.DB) pastureWasteDischarge {
	p.pastureWasteDischargeDo.ReplaceDB(db)
	return p
}

type pastureWasteDischargeDo struct{ gen.DO }

type IPastureWasteDischargeDo interface {
	gen.SubQuery
	Debug() IPastureWasteDischargeDo
	WithContext(ctx context.Context) IPastureWasteDischargeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPastureWasteDischargeDo
	WriteDB() IPastureWasteDischargeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPastureWasteDischargeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPastureWasteDischargeDo
	Not(conds ...gen.Condition) IPastureWasteDischargeDo
	Or(conds ...gen.Condition) IPastureWasteDischargeDo
	Select(conds ...field.Expr) IPastureWasteDischargeDo
	Where(conds ...gen.Condition) IPastureWasteDischargeDo
	Order(conds ...field.Expr) IPastureWasteDischargeDo
	Distinct(cols ...field.Expr) IPastureWasteDischargeDo
	Omit(cols ...field.Expr) IPastureWasteDischargeDo
	Join(table schema.Tabler, on ...field.Expr) IPastureWasteDischargeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPastureWasteDischargeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPastureWasteDischargeDo
	Group(cols ...field.Expr) IPastureWasteDischargeDo
	Having(conds ...gen.Condition) IPastureWasteDischargeDo
	Limit(limit int) IPastureWasteDischargeDo
	Offset(offset int) IPastureWasteDischargeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureWasteDischargeDo
	Unscoped() IPastureWasteDischargeDo
	Create(values ...*pasture.PastureWasteDischarge) error
	CreateInBatches(values []*pasture.PastureWasteDischarge, batchSize int) error
	Save(values ...*pasture.PastureWasteDischarge) error
	First() (*pasture.PastureWasteDischarge, error)
	Take() (*pasture.PastureWasteDischarge, error)
	Last() (*pasture.PastureWasteDischarge, error)
	Find() ([]*pasture.PastureWasteDischarge, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureWasteDischarge, err error)
	FindInBatches(result *[]*pasture.PastureWasteDischarge, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.PastureWasteDischarge) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPastureWasteDischargeDo
	Assign(attrs ...field.AssignExpr) IPastureWasteDischargeDo
	Joins(fields ...field.RelationField) IPastureWasteDischargeDo
	Preload(fields ...field.RelationField) IPastureWasteDischargeDo
	FirstOrInit() (*pasture.PastureWasteDischarge, error)
	FirstOrCreate() (*pasture.PastureWasteDischarge, error)
	FindByPage(offset int, limit int) (result []*pasture.PastureWasteDischarge, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPastureWasteDischargeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pastureWasteDischargeDo) Debug() IPastureWasteDischargeDo {
	return p.withDO(p.DO.Debug())
}

func (p pastureWasteDischargeDo) WithContext(ctx context.Context) IPastureWasteDischargeDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pastureWasteDischargeDo) ReadDB() IPastureWasteDischargeDo {
	return p.Clauses(dbresolver.Read)
}

func (p pastureWasteDischargeDo) WriteDB() IPastureWasteDischargeDo {
	return p.Clauses(dbresolver.Write)
}

func (p pastureWasteDischargeDo) Session(config *gorm.Session) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Session(config))
}

func (p pastureWasteDischargeDo) Clauses(conds ...clause.Expression) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pastureWasteDischargeDo) Returning(value interface{}, columns ...string) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pastureWasteDischargeDo) Not(conds ...gen.Condition) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pastureWasteDischargeDo) Or(conds ...gen.Condition) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pastureWasteDischargeDo) Select(conds ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pastureWasteDischargeDo) Where(conds ...gen.Condition) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pastureWasteDischargeDo) Order(conds ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pastureWasteDischargeDo) Distinct(cols ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pastureWasteDischargeDo) Omit(cols ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pastureWasteDischargeDo) Join(table schema.Tabler, on ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pastureWasteDischargeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pastureWasteDischargeDo) RightJoin(table schema.Tabler, on ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pastureWasteDischargeDo) Group(cols ...field.Expr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pastureWasteDischargeDo) Having(conds ...gen.Condition) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pastureWasteDischargeDo) Limit(limit int) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pastureWasteDischargeDo) Offset(offset int) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pastureWasteDischargeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pastureWasteDischargeDo) Unscoped() IPastureWasteDischargeDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pastureWasteDischargeDo) Create(values ...*pasture.PastureWasteDischarge) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pastureWasteDischargeDo) CreateInBatches(values []*pasture.PastureWasteDischarge, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pastureWasteDischargeDo) Save(values ...*pasture.PastureWasteDischarge) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pastureWasteDischargeDo) First() (*pasture.PastureWasteDischarge, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureWasteDischarge), nil
	}
}

func (p pastureWasteDischargeDo) Take() (*pasture.PastureWasteDischarge, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureWasteDischarge), nil
	}
}

func (p pastureWasteDischargeDo) Last() (*pasture.PastureWasteDischarge, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureWasteDischarge), nil
	}
}

func (p pastureWasteDischargeDo) Find() ([]*pasture.PastureWasteDischarge, error) {
	result, err := p.DO.Find()
	return result.([]*pasture.PastureWasteDischarge), err
}

func (p pastureWasteDischargeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.PastureWasteDischarge, err error) {
	buf := make([]*pasture.PastureWasteDischarge, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pastureWasteDischargeDo) FindInBatches(result *[]*pasture.PastureWasteDischarge, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pastureWasteDischargeDo) Attrs(attrs ...field.AssignExpr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pastureWasteDischargeDo) Assign(attrs ...field.AssignExpr) IPastureWasteDischargeDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pastureWasteDischargeDo) Joins(fields ...field.RelationField) IPastureWasteDischargeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pastureWasteDischargeDo) Preload(fields ...field.RelationField) IPastureWasteDischargeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pastureWasteDischargeDo) FirstOrInit() (*pasture.PastureWasteDischarge, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureWasteDischarge), nil
	}
}

func (p pastureWasteDischargeDo) FirstOrCreate() (*pasture.PastureWasteDischarge, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.PastureWasteDischarge), nil
	}
}

func (p pastureWasteDischargeDo) FindByPage(offset int, limit int) (result []*pasture.PastureWasteDischarge, count int64, err error) {
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

func (p pastureWasteDischargeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pastureWasteDischargeDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pastureWasteDischargeDo) Delete(models ...*pasture.PastureWasteDischarge) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pastureWasteDischargeDo) withDO(do gen.Dao) *pastureWasteDischargeDo {
	p.DO = *do.(*gen.DO)
	return p
}
