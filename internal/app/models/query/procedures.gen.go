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

func newProcedure(db *gorm.DB, opts ...gen.DOOption) procedure {
	_procedure := procedure{}

	_procedure.procedureDo.UseDB(db, opts...)
	_procedure.procedureDo.UseModel(&models.Procedure{})

	tableName := _procedure.procedureDo.TableName()
	_procedure.ALL = field.NewAsterisk(tableName)
	_procedure.ID = field.NewUint(tableName, "id")
	_procedure.CreatedAt = field.NewTime(tableName, "created_at")
	_procedure.UpdatedAt = field.NewTime(tableName, "updated_at")
	_procedure.DeletedAt = field.NewField(tableName, "deleted_at")
	_procedure.PID = field.NewString(tableName, "p_id")
	_procedure.Type = field.NewUint(tableName, "type")
	_procedure.Name = field.NewString(tableName, "name")
	_procedure.State = field.NewUint(tableName, "state")
	_procedure.PHash = field.NewString(tableName, "p_hash")
	_procedure.CheckCode = field.NewString(tableName, "check_code")
	_procedure.SerialNumber = field.NewInt64(tableName, "serial_number")
	_procedure.Operator = field.NewString(tableName, "operator")
	_procedure.StartTimestamp = field.NewTime(tableName, "start_timestamp")
	_procedure.CompletedTimestamp = field.NewTime(tableName, "completed_timestamp")
	_procedure.PrePID = field.NewString(tableName, "pre_p_id")
	_procedure.LastFlag = field.NewString(tableName, "last_flag")
	_procedure.ICID = field.NewString(tableName, "ic_id")

	_procedure.fillFieldMap()

	return _procedure
}

type procedure struct {
	procedureDo procedureDo

	ALL                field.Asterisk
	ID                 field.Uint
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field
	PID                field.String
	Type               field.Uint
	Name               field.String
	State              field.Uint
	PHash              field.String
	CheckCode          field.String
	SerialNumber       field.Int64
	Operator           field.String
	StartTimestamp     field.Time
	CompletedTimestamp field.Time
	PrePID             field.String
	LastFlag           field.String
	ICID               field.String

	fieldMap map[string]field.Expr
}

func (p procedure) Table(newTableName string) *procedure {
	p.procedureDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p procedure) As(alias string) *procedure {
	p.procedureDo.DO = *(p.procedureDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *procedure) updateTableName(table string) *procedure {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.PID = field.NewString(table, "p_id")
	p.Type = field.NewUint(table, "type")
	p.Name = field.NewString(table, "name")
	p.State = field.NewUint(table, "state")
	p.PHash = field.NewString(table, "p_hash")
	p.CheckCode = field.NewString(table, "check_code")
	p.SerialNumber = field.NewInt64(table, "serial_number")
	p.Operator = field.NewString(table, "operator")
	p.StartTimestamp = field.NewTime(table, "start_timestamp")
	p.CompletedTimestamp = field.NewTime(table, "completed_timestamp")
	p.PrePID = field.NewString(table, "pre_p_id")
	p.LastFlag = field.NewString(table, "last_flag")
	p.ICID = field.NewString(table, "ic_id")

	p.fillFieldMap()

	return p
}

func (p *procedure) WithContext(ctx context.Context) IProcedureDo {
	return p.procedureDo.WithContext(ctx)
}

func (p procedure) TableName() string { return p.procedureDo.TableName() }

func (p procedure) Alias() string { return p.procedureDo.Alias() }

func (p procedure) Columns(cols ...field.Expr) gen.Columns { return p.procedureDo.Columns(cols...) }

func (p *procedure) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *procedure) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 17)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["p_id"] = p.PID
	p.fieldMap["type"] = p.Type
	p.fieldMap["name"] = p.Name
	p.fieldMap["state"] = p.State
	p.fieldMap["p_hash"] = p.PHash
	p.fieldMap["check_code"] = p.CheckCode
	p.fieldMap["serial_number"] = p.SerialNumber
	p.fieldMap["operator"] = p.Operator
	p.fieldMap["start_timestamp"] = p.StartTimestamp
	p.fieldMap["completed_timestamp"] = p.CompletedTimestamp
	p.fieldMap["pre_p_id"] = p.PrePID
	p.fieldMap["last_flag"] = p.LastFlag
	p.fieldMap["ic_id"] = p.ICID
}

func (p procedure) clone(db *gorm.DB) procedure {
	p.procedureDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p procedure) replaceDB(db *gorm.DB) procedure {
	p.procedureDo.ReplaceDB(db)
	return p
}

type procedureDo struct{ gen.DO }

type IProcedureDo interface {
	gen.SubQuery
	Debug() IProcedureDo
	WithContext(ctx context.Context) IProcedureDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProcedureDo
	WriteDB() IProcedureDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProcedureDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProcedureDo
	Not(conds ...gen.Condition) IProcedureDo
	Or(conds ...gen.Condition) IProcedureDo
	Select(conds ...field.Expr) IProcedureDo
	Where(conds ...gen.Condition) IProcedureDo
	Order(conds ...field.Expr) IProcedureDo
	Distinct(cols ...field.Expr) IProcedureDo
	Omit(cols ...field.Expr) IProcedureDo
	Join(table schema.Tabler, on ...field.Expr) IProcedureDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProcedureDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProcedureDo
	Group(cols ...field.Expr) IProcedureDo
	Having(conds ...gen.Condition) IProcedureDo
	Limit(limit int) IProcedureDo
	Offset(offset int) IProcedureDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProcedureDo
	Unscoped() IProcedureDo
	Create(values ...*models.Procedure) error
	CreateInBatches(values []*models.Procedure, batchSize int) error
	Save(values ...*models.Procedure) error
	First() (*models.Procedure, error)
	Take() (*models.Procedure, error)
	Last() (*models.Procedure, error)
	Find() ([]*models.Procedure, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Procedure, err error)
	FindInBatches(result *[]*models.Procedure, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Procedure) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProcedureDo
	Assign(attrs ...field.AssignExpr) IProcedureDo
	Joins(fields ...field.RelationField) IProcedureDo
	Preload(fields ...field.RelationField) IProcedureDo
	FirstOrInit() (*models.Procedure, error)
	FirstOrCreate() (*models.Procedure, error)
	FindByPage(offset int, limit int) (result []*models.Procedure, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProcedureDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p procedureDo) Debug() IProcedureDo {
	return p.withDO(p.DO.Debug())
}

func (p procedureDo) WithContext(ctx context.Context) IProcedureDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p procedureDo) ReadDB() IProcedureDo {
	return p.Clauses(dbresolver.Read)
}

func (p procedureDo) WriteDB() IProcedureDo {
	return p.Clauses(dbresolver.Write)
}

func (p procedureDo) Session(config *gorm.Session) IProcedureDo {
	return p.withDO(p.DO.Session(config))
}

func (p procedureDo) Clauses(conds ...clause.Expression) IProcedureDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p procedureDo) Returning(value interface{}, columns ...string) IProcedureDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p procedureDo) Not(conds ...gen.Condition) IProcedureDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p procedureDo) Or(conds ...gen.Condition) IProcedureDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p procedureDo) Select(conds ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p procedureDo) Where(conds ...gen.Condition) IProcedureDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p procedureDo) Order(conds ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p procedureDo) Distinct(cols ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p procedureDo) Omit(cols ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p procedureDo) Join(table schema.Tabler, on ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p procedureDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p procedureDo) RightJoin(table schema.Tabler, on ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p procedureDo) Group(cols ...field.Expr) IProcedureDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p procedureDo) Having(conds ...gen.Condition) IProcedureDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p procedureDo) Limit(limit int) IProcedureDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p procedureDo) Offset(offset int) IProcedureDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p procedureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProcedureDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p procedureDo) Unscoped() IProcedureDo {
	return p.withDO(p.DO.Unscoped())
}

func (p procedureDo) Create(values ...*models.Procedure) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p procedureDo) CreateInBatches(values []*models.Procedure, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p procedureDo) Save(values ...*models.Procedure) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p procedureDo) First() (*models.Procedure, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Procedure), nil
	}
}

func (p procedureDo) Take() (*models.Procedure, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Procedure), nil
	}
}

func (p procedureDo) Last() (*models.Procedure, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Procedure), nil
	}
}

func (p procedureDo) Find() ([]*models.Procedure, error) {
	result, err := p.DO.Find()
	return result.([]*models.Procedure), err
}

func (p procedureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Procedure, err error) {
	buf := make([]*models.Procedure, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p procedureDo) FindInBatches(result *[]*models.Procedure, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p procedureDo) Attrs(attrs ...field.AssignExpr) IProcedureDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p procedureDo) Assign(attrs ...field.AssignExpr) IProcedureDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p procedureDo) Joins(fields ...field.RelationField) IProcedureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p procedureDo) Preload(fields ...field.RelationField) IProcedureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p procedureDo) FirstOrInit() (*models.Procedure, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Procedure), nil
	}
}

func (p procedureDo) FirstOrCreate() (*models.Procedure, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Procedure), nil
	}
}

func (p procedureDo) FindByPage(offset int, limit int) (result []*models.Procedure, count int64, err error) {
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

func (p procedureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p procedureDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p procedureDo) Delete(models ...*models.Procedure) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *procedureDo) withDO(do gen.Dao) *procedureDo {
	p.DO = *do.(*gen.DO)
	return p
}
