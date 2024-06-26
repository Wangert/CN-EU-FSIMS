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

func newSlaughterStun(db *gorm.DB, opts ...gen.DOOption) slaughterStun {
	_slaughterStun := slaughterStun{}

	_slaughterStun.slaughterStunDo.UseDB(db, opts...)
	_slaughterStun.slaughterStunDo.UseModel(&slaughter.SlaughterStun{})

	tableName := _slaughterStun.slaughterStunDo.TableName()
	_slaughterStun.ALL = field.NewAsterisk(tableName)
	_slaughterStun.ID = field.NewUint(tableName, "id")
	_slaughterStun.CreatedAt = field.NewTime(tableName, "created_at")
	_slaughterStun.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaughterStun.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaughterStun.SlaughterProcedureMonitoringDataID = field.NewUint(tableName, "slaughter_procedure_monitoring_data_id")
	_slaughterStun.Stun1 = field.NewFloat64(tableName, "stun1")
	_slaughterStun.Stun2 = field.NewFloat64(tableName, "stun2")
	_slaughterStun.Stun3 = field.NewFloat64(tableName, "stun3")

	_slaughterStun.fillFieldMap()

	return _slaughterStun
}

type slaughterStun struct {
	slaughterStunDo slaughterStunDo

	ALL                                field.Asterisk
	ID                                 field.Uint
	CreatedAt                          field.Time
	UpdatedAt                          field.Time
	DeletedAt                          field.Field
	SlaughterProcedureMonitoringDataID field.Uint
	Stun1                              field.Float64
	Stun2                              field.Float64
	Stun3                              field.Float64

	fieldMap map[string]field.Expr
}

func (s slaughterStun) Table(newTableName string) *slaughterStun {
	s.slaughterStunDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaughterStun) As(alias string) *slaughterStun {
	s.slaughterStunDo.DO = *(s.slaughterStunDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaughterStun) updateTableName(table string) *slaughterStun {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.SlaughterProcedureMonitoringDataID = field.NewUint(table, "slaughter_procedure_monitoring_data_id")
	s.Stun1 = field.NewFloat64(table, "stun1")
	s.Stun2 = field.NewFloat64(table, "stun2")
	s.Stun3 = field.NewFloat64(table, "stun3")

	s.fillFieldMap()

	return s
}

func (s *slaughterStun) WithContext(ctx context.Context) ISlaughterStunDo {
	return s.slaughterStunDo.WithContext(ctx)
}

func (s slaughterStun) TableName() string { return s.slaughterStunDo.TableName() }

func (s slaughterStun) Alias() string { return s.slaughterStunDo.Alias() }

func (s slaughterStun) Columns(cols ...field.Expr) gen.Columns {
	return s.slaughterStunDo.Columns(cols...)
}

func (s *slaughterStun) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaughterStun) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["slaughter_procedure_monitoring_data_id"] = s.SlaughterProcedureMonitoringDataID
	s.fieldMap["stun1"] = s.Stun1
	s.fieldMap["stun2"] = s.Stun2
	s.fieldMap["stun3"] = s.Stun3
}

func (s slaughterStun) clone(db *gorm.DB) slaughterStun {
	s.slaughterStunDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaughterStun) replaceDB(db *gorm.DB) slaughterStun {
	s.slaughterStunDo.ReplaceDB(db)
	return s
}

type slaughterStunDo struct{ gen.DO }

type ISlaughterStunDo interface {
	gen.SubQuery
	Debug() ISlaughterStunDo
	WithContext(ctx context.Context) ISlaughterStunDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaughterStunDo
	WriteDB() ISlaughterStunDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaughterStunDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaughterStunDo
	Not(conds ...gen.Condition) ISlaughterStunDo
	Or(conds ...gen.Condition) ISlaughterStunDo
	Select(conds ...field.Expr) ISlaughterStunDo
	Where(conds ...gen.Condition) ISlaughterStunDo
	Order(conds ...field.Expr) ISlaughterStunDo
	Distinct(cols ...field.Expr) ISlaughterStunDo
	Omit(cols ...field.Expr) ISlaughterStunDo
	Join(table schema.Tabler, on ...field.Expr) ISlaughterStunDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterStunDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterStunDo
	Group(cols ...field.Expr) ISlaughterStunDo
	Having(conds ...gen.Condition) ISlaughterStunDo
	Limit(limit int) ISlaughterStunDo
	Offset(offset int) ISlaughterStunDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterStunDo
	Unscoped() ISlaughterStunDo
	Create(values ...*slaughter.SlaughterStun) error
	CreateInBatches(values []*slaughter.SlaughterStun, batchSize int) error
	Save(values ...*slaughter.SlaughterStun) error
	First() (*slaughter.SlaughterStun, error)
	Take() (*slaughter.SlaughterStun, error)
	Last() (*slaughter.SlaughterStun, error)
	Find() ([]*slaughter.SlaughterStun, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterStun, err error)
	FindInBatches(result *[]*slaughter.SlaughterStun, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaughterStun) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaughterStunDo
	Assign(attrs ...field.AssignExpr) ISlaughterStunDo
	Joins(fields ...field.RelationField) ISlaughterStunDo
	Preload(fields ...field.RelationField) ISlaughterStunDo
	FirstOrInit() (*slaughter.SlaughterStun, error)
	FirstOrCreate() (*slaughter.SlaughterStun, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaughterStun, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaughterStunDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaughterStunDo) Debug() ISlaughterStunDo {
	return s.withDO(s.DO.Debug())
}

func (s slaughterStunDo) WithContext(ctx context.Context) ISlaughterStunDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaughterStunDo) ReadDB() ISlaughterStunDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaughterStunDo) WriteDB() ISlaughterStunDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaughterStunDo) Session(config *gorm.Session) ISlaughterStunDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaughterStunDo) Clauses(conds ...clause.Expression) ISlaughterStunDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaughterStunDo) Returning(value interface{}, columns ...string) ISlaughterStunDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaughterStunDo) Not(conds ...gen.Condition) ISlaughterStunDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaughterStunDo) Or(conds ...gen.Condition) ISlaughterStunDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaughterStunDo) Select(conds ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaughterStunDo) Where(conds ...gen.Condition) ISlaughterStunDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaughterStunDo) Order(conds ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaughterStunDo) Distinct(cols ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaughterStunDo) Omit(cols ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaughterStunDo) Join(table schema.Tabler, on ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaughterStunDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaughterStunDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaughterStunDo) Group(cols ...field.Expr) ISlaughterStunDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaughterStunDo) Having(conds ...gen.Condition) ISlaughterStunDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaughterStunDo) Limit(limit int) ISlaughterStunDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaughterStunDo) Offset(offset int) ISlaughterStunDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaughterStunDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterStunDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaughterStunDo) Unscoped() ISlaughterStunDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaughterStunDo) Create(values ...*slaughter.SlaughterStun) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaughterStunDo) CreateInBatches(values []*slaughter.SlaughterStun, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaughterStunDo) Save(values ...*slaughter.SlaughterStun) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaughterStunDo) First() (*slaughter.SlaughterStun, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterStun), nil
	}
}

func (s slaughterStunDo) Take() (*slaughter.SlaughterStun, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterStun), nil
	}
}

func (s slaughterStunDo) Last() (*slaughter.SlaughterStun, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterStun), nil
	}
}

func (s slaughterStunDo) Find() ([]*slaughter.SlaughterStun, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaughterStun), err
}

func (s slaughterStunDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterStun, err error) {
	buf := make([]*slaughter.SlaughterStun, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaughterStunDo) FindInBatches(result *[]*slaughter.SlaughterStun, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaughterStunDo) Attrs(attrs ...field.AssignExpr) ISlaughterStunDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaughterStunDo) Assign(attrs ...field.AssignExpr) ISlaughterStunDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaughterStunDo) Joins(fields ...field.RelationField) ISlaughterStunDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaughterStunDo) Preload(fields ...field.RelationField) ISlaughterStunDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaughterStunDo) FirstOrInit() (*slaughter.SlaughterStun, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterStun), nil
	}
}

func (s slaughterStunDo) FirstOrCreate() (*slaughter.SlaughterStun, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterStun), nil
	}
}

func (s slaughterStunDo) FindByPage(offset int, limit int) (result []*slaughter.SlaughterStun, count int64, err error) {
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

func (s slaughterStunDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaughterStunDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaughterStunDo) Delete(models ...*slaughter.SlaughterStun) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaughterStunDo) withDO(do gen.Dao) *slaughterStunDo {
	s.DO = *do.(*gen.DO)
	return s
}
