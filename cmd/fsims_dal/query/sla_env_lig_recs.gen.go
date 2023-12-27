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

func newSlaEnvLigRec(db *gorm.DB, opts ...gen.DOOption) slaEnvLigRec {
	_slaEnvLigRec := slaEnvLigRec{}

	_slaEnvLigRec.slaEnvLigRecDo.UseDB(db, opts...)
	_slaEnvLigRec.slaEnvLigRecDo.UseModel(&slaughter.SlaEnvLigRec{})

	tableName := _slaEnvLigRec.slaEnvLigRecDo.TableName()
	_slaEnvLigRec.ALL = field.NewAsterisk(tableName)
	_slaEnvLigRec.ID = field.NewUint(tableName, "id")
	_slaEnvLigRec.CreatedAt = field.NewTime(tableName, "created_at")
	_slaEnvLigRec.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaEnvLigRec.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaEnvLigRec.SlaInfoMonID = field.NewUint(tableName, "sla_info_mon_id")
	_slaEnvLigRec.SlaEnvLigRec1 = field.NewFloat32(tableName, "sla_env_lig_rec1")
	_slaEnvLigRec.SlaEnvLigRec2 = field.NewFloat32(tableName, "sla_env_lig_rec2")
	_slaEnvLigRec.SlaEnvLigRec3 = field.NewFloat32(tableName, "sla_env_lig_rec3")
	_slaEnvLigRec.SlaEnvLigRec4 = field.NewFloat32(tableName, "sla_env_lig_rec4")

	_slaEnvLigRec.fillFieldMap()

	return _slaEnvLigRec
}

type slaEnvLigRec struct {
	slaEnvLigRecDo slaEnvLigRecDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	SlaInfoMonID  field.Uint
	SlaEnvLigRec1 field.Float32
	SlaEnvLigRec2 field.Float32
	SlaEnvLigRec3 field.Float32
	SlaEnvLigRec4 field.Float32

	fieldMap map[string]field.Expr
}

func (s slaEnvLigRec) Table(newTableName string) *slaEnvLigRec {
	s.slaEnvLigRecDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaEnvLigRec) As(alias string) *slaEnvLigRec {
	s.slaEnvLigRecDo.DO = *(s.slaEnvLigRecDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaEnvLigRec) updateTableName(table string) *slaEnvLigRec {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.SlaInfoMonID = field.NewUint(table, "sla_info_mon_id")
	s.SlaEnvLigRec1 = field.NewFloat32(table, "sla_env_lig_rec1")
	s.SlaEnvLigRec2 = field.NewFloat32(table, "sla_env_lig_rec2")
	s.SlaEnvLigRec3 = field.NewFloat32(table, "sla_env_lig_rec3")
	s.SlaEnvLigRec4 = field.NewFloat32(table, "sla_env_lig_rec4")

	s.fillFieldMap()

	return s
}

func (s *slaEnvLigRec) WithContext(ctx context.Context) ISlaEnvLigRecDo {
	return s.slaEnvLigRecDo.WithContext(ctx)
}

func (s slaEnvLigRec) TableName() string { return s.slaEnvLigRecDo.TableName() }

func (s slaEnvLigRec) Alias() string { return s.slaEnvLigRecDo.Alias() }

func (s slaEnvLigRec) Columns(cols ...field.Expr) gen.Columns {
	return s.slaEnvLigRecDo.Columns(cols...)
}

func (s *slaEnvLigRec) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaEnvLigRec) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["sla_info_mon_id"] = s.SlaInfoMonID
	s.fieldMap["sla_env_lig_rec1"] = s.SlaEnvLigRec1
	s.fieldMap["sla_env_lig_rec2"] = s.SlaEnvLigRec2
	s.fieldMap["sla_env_lig_rec3"] = s.SlaEnvLigRec3
	s.fieldMap["sla_env_lig_rec4"] = s.SlaEnvLigRec4
}

func (s slaEnvLigRec) clone(db *gorm.DB) slaEnvLigRec {
	s.slaEnvLigRecDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaEnvLigRec) replaceDB(db *gorm.DB) slaEnvLigRec {
	s.slaEnvLigRecDo.ReplaceDB(db)
	return s
}

type slaEnvLigRecDo struct{ gen.DO }

type ISlaEnvLigRecDo interface {
	gen.SubQuery
	Debug() ISlaEnvLigRecDo
	WithContext(ctx context.Context) ISlaEnvLigRecDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaEnvLigRecDo
	WriteDB() ISlaEnvLigRecDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaEnvLigRecDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaEnvLigRecDo
	Not(conds ...gen.Condition) ISlaEnvLigRecDo
	Or(conds ...gen.Condition) ISlaEnvLigRecDo
	Select(conds ...field.Expr) ISlaEnvLigRecDo
	Where(conds ...gen.Condition) ISlaEnvLigRecDo
	Order(conds ...field.Expr) ISlaEnvLigRecDo
	Distinct(cols ...field.Expr) ISlaEnvLigRecDo
	Omit(cols ...field.Expr) ISlaEnvLigRecDo
	Join(table schema.Tabler, on ...field.Expr) ISlaEnvLigRecDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaEnvLigRecDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaEnvLigRecDo
	Group(cols ...field.Expr) ISlaEnvLigRecDo
	Having(conds ...gen.Condition) ISlaEnvLigRecDo
	Limit(limit int) ISlaEnvLigRecDo
	Offset(offset int) ISlaEnvLigRecDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaEnvLigRecDo
	Unscoped() ISlaEnvLigRecDo
	Create(values ...*slaughter.SlaEnvLigRec) error
	CreateInBatches(values []*slaughter.SlaEnvLigRec, batchSize int) error
	Save(values ...*slaughter.SlaEnvLigRec) error
	First() (*slaughter.SlaEnvLigRec, error)
	Take() (*slaughter.SlaEnvLigRec, error)
	Last() (*slaughter.SlaEnvLigRec, error)
	Find() ([]*slaughter.SlaEnvLigRec, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaEnvLigRec, err error)
	FindInBatches(result *[]*slaughter.SlaEnvLigRec, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaEnvLigRec) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaEnvLigRecDo
	Assign(attrs ...field.AssignExpr) ISlaEnvLigRecDo
	Joins(fields ...field.RelationField) ISlaEnvLigRecDo
	Preload(fields ...field.RelationField) ISlaEnvLigRecDo
	FirstOrInit() (*slaughter.SlaEnvLigRec, error)
	FirstOrCreate() (*slaughter.SlaEnvLigRec, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaEnvLigRec, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaEnvLigRecDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaEnvLigRecDo) Debug() ISlaEnvLigRecDo {
	return s.withDO(s.DO.Debug())
}

func (s slaEnvLigRecDo) WithContext(ctx context.Context) ISlaEnvLigRecDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaEnvLigRecDo) ReadDB() ISlaEnvLigRecDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaEnvLigRecDo) WriteDB() ISlaEnvLigRecDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaEnvLigRecDo) Session(config *gorm.Session) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaEnvLigRecDo) Clauses(conds ...clause.Expression) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaEnvLigRecDo) Returning(value interface{}, columns ...string) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaEnvLigRecDo) Not(conds ...gen.Condition) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaEnvLigRecDo) Or(conds ...gen.Condition) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaEnvLigRecDo) Select(conds ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaEnvLigRecDo) Where(conds ...gen.Condition) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaEnvLigRecDo) Order(conds ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaEnvLigRecDo) Distinct(cols ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaEnvLigRecDo) Omit(cols ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaEnvLigRecDo) Join(table schema.Tabler, on ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaEnvLigRecDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaEnvLigRecDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaEnvLigRecDo) Group(cols ...field.Expr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaEnvLigRecDo) Having(conds ...gen.Condition) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaEnvLigRecDo) Limit(limit int) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaEnvLigRecDo) Offset(offset int) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaEnvLigRecDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaEnvLigRecDo) Unscoped() ISlaEnvLigRecDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaEnvLigRecDo) Create(values ...*slaughter.SlaEnvLigRec) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaEnvLigRecDo) CreateInBatches(values []*slaughter.SlaEnvLigRec, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaEnvLigRecDo) Save(values ...*slaughter.SlaEnvLigRec) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaEnvLigRecDo) First() (*slaughter.SlaEnvLigRec, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaEnvLigRec), nil
	}
}

func (s slaEnvLigRecDo) Take() (*slaughter.SlaEnvLigRec, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaEnvLigRec), nil
	}
}

func (s slaEnvLigRecDo) Last() (*slaughter.SlaEnvLigRec, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaEnvLigRec), nil
	}
}

func (s slaEnvLigRecDo) Find() ([]*slaughter.SlaEnvLigRec, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaEnvLigRec), err
}

func (s slaEnvLigRecDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaEnvLigRec, err error) {
	buf := make([]*slaughter.SlaEnvLigRec, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaEnvLigRecDo) FindInBatches(result *[]*slaughter.SlaEnvLigRec, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaEnvLigRecDo) Attrs(attrs ...field.AssignExpr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaEnvLigRecDo) Assign(attrs ...field.AssignExpr) ISlaEnvLigRecDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaEnvLigRecDo) Joins(fields ...field.RelationField) ISlaEnvLigRecDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaEnvLigRecDo) Preload(fields ...field.RelationField) ISlaEnvLigRecDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaEnvLigRecDo) FirstOrInit() (*slaughter.SlaEnvLigRec, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaEnvLigRec), nil
	}
}

func (s slaEnvLigRecDo) FirstOrCreate() (*slaughter.SlaEnvLigRec, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaEnvLigRec), nil
	}
}

func (s slaEnvLigRecDo) FindByPage(offset int, limit int) (result []*slaughter.SlaEnvLigRec, count int64, err error) {
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

func (s slaEnvLigRecDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaEnvLigRecDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaEnvLigRecDo) Delete(models ...*slaughter.SlaEnvLigRec) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaEnvLigRecDo) withDO(do gen.Dao) *slaEnvLigRecDo {
	s.DO = *do.(*gen.DO)
	return s
}
