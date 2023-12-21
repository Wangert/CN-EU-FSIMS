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

func newSlaughterBatch(db *gorm.DB, opts ...gen.DOOption) slaughterBatch {
	_slaughterBatch := slaughterBatch{}

	_slaughterBatch.slaughterBatchDo.UseDB(db, opts...)
	_slaughterBatch.slaughterBatchDo.UseModel(&slaughter.SlaughterBatch{})

	tableName := _slaughterBatch.slaughterBatchDo.TableName()
	_slaughterBatch.ALL = field.NewAsterisk(tableName)
	_slaughterBatch.ID = field.NewUint(tableName, "id")
	_slaughterBatch.CreatedAt = field.NewTime(tableName, "created_at")
	_slaughterBatch.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaughterBatch.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaughterBatch.BatchNumber = field.NewString(tableName, "batch_number")
	_slaughterBatch.HouseNumber = field.NewString(tableName, "house_number")
	_slaughterBatch.State = field.NewInt(tableName, "state")
	_slaughterBatch.PID = field.NewString(tableName, "p_id")
	_slaughterBatch.Worker = field.NewString(tableName, "worker")
	_slaughterBatch.CowNumber = field.NewString(tableName, "cow_number")

	_slaughterBatch.fillFieldMap()

	return _slaughterBatch
}

type slaughterBatch struct {
	slaughterBatchDo slaughterBatchDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	BatchNumber field.String
	HouseNumber field.String
	State       field.Int
	PID         field.String
	Worker      field.String
	CowNumber   field.String

	fieldMap map[string]field.Expr
}

func (s slaughterBatch) Table(newTableName string) *slaughterBatch {
	s.slaughterBatchDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaughterBatch) As(alias string) *slaughterBatch {
	s.slaughterBatchDo.DO = *(s.slaughterBatchDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaughterBatch) updateTableName(table string) *slaughterBatch {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.BatchNumber = field.NewString(table, "batch_number")
	s.HouseNumber = field.NewString(table, "house_number")
	s.State = field.NewInt(table, "state")
	s.PID = field.NewString(table, "p_id")
	s.Worker = field.NewString(table, "worker")
	s.CowNumber = field.NewString(table, "cow_number")

	s.fillFieldMap()

	return s
}

func (s *slaughterBatch) WithContext(ctx context.Context) ISlaughterBatchDo {
	return s.slaughterBatchDo.WithContext(ctx)
}

func (s slaughterBatch) TableName() string { return s.slaughterBatchDo.TableName() }

func (s slaughterBatch) Alias() string { return s.slaughterBatchDo.Alias() }

func (s slaughterBatch) Columns(cols ...field.Expr) gen.Columns {
	return s.slaughterBatchDo.Columns(cols...)
}

func (s *slaughterBatch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaughterBatch) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["batch_number"] = s.BatchNumber
	s.fieldMap["house_number"] = s.HouseNumber
	s.fieldMap["state"] = s.State
	s.fieldMap["p_id"] = s.PID
	s.fieldMap["worker"] = s.Worker
	s.fieldMap["cow_number"] = s.CowNumber
}

func (s slaughterBatch) clone(db *gorm.DB) slaughterBatch {
	s.slaughterBatchDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaughterBatch) replaceDB(db *gorm.DB) slaughterBatch {
	s.slaughterBatchDo.ReplaceDB(db)
	return s
}

type slaughterBatchDo struct{ gen.DO }

type ISlaughterBatchDo interface {
	gen.SubQuery
	Debug() ISlaughterBatchDo
	WithContext(ctx context.Context) ISlaughterBatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaughterBatchDo
	WriteDB() ISlaughterBatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaughterBatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaughterBatchDo
	Not(conds ...gen.Condition) ISlaughterBatchDo
	Or(conds ...gen.Condition) ISlaughterBatchDo
	Select(conds ...field.Expr) ISlaughterBatchDo
	Where(conds ...gen.Condition) ISlaughterBatchDo
	Order(conds ...field.Expr) ISlaughterBatchDo
	Distinct(cols ...field.Expr) ISlaughterBatchDo
	Omit(cols ...field.Expr) ISlaughterBatchDo
	Join(table schema.Tabler, on ...field.Expr) ISlaughterBatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterBatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterBatchDo
	Group(cols ...field.Expr) ISlaughterBatchDo
	Having(conds ...gen.Condition) ISlaughterBatchDo
	Limit(limit int) ISlaughterBatchDo
	Offset(offset int) ISlaughterBatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterBatchDo
	Unscoped() ISlaughterBatchDo
	Create(values ...*slaughter.SlaughterBatch) error
	CreateInBatches(values []*slaughter.SlaughterBatch, batchSize int) error
	Save(values ...*slaughter.SlaughterBatch) error
	First() (*slaughter.SlaughterBatch, error)
	Take() (*slaughter.SlaughterBatch, error)
	Last() (*slaughter.SlaughterBatch, error)
	Find() ([]*slaughter.SlaughterBatch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterBatch, err error)
	FindInBatches(result *[]*slaughter.SlaughterBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaughterBatch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaughterBatchDo
	Assign(attrs ...field.AssignExpr) ISlaughterBatchDo
	Joins(fields ...field.RelationField) ISlaughterBatchDo
	Preload(fields ...field.RelationField) ISlaughterBatchDo
	FirstOrInit() (*slaughter.SlaughterBatch, error)
	FirstOrCreate() (*slaughter.SlaughterBatch, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaughterBatch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaughterBatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaughterBatchDo) Debug() ISlaughterBatchDo {
	return s.withDO(s.DO.Debug())
}

func (s slaughterBatchDo) WithContext(ctx context.Context) ISlaughterBatchDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaughterBatchDo) ReadDB() ISlaughterBatchDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaughterBatchDo) WriteDB() ISlaughterBatchDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaughterBatchDo) Session(config *gorm.Session) ISlaughterBatchDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaughterBatchDo) Clauses(conds ...clause.Expression) ISlaughterBatchDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaughterBatchDo) Returning(value interface{}, columns ...string) ISlaughterBatchDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaughterBatchDo) Not(conds ...gen.Condition) ISlaughterBatchDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaughterBatchDo) Or(conds ...gen.Condition) ISlaughterBatchDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaughterBatchDo) Select(conds ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaughterBatchDo) Where(conds ...gen.Condition) ISlaughterBatchDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaughterBatchDo) Order(conds ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaughterBatchDo) Distinct(cols ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaughterBatchDo) Omit(cols ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaughterBatchDo) Join(table schema.Tabler, on ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaughterBatchDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaughterBatchDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaughterBatchDo) Group(cols ...field.Expr) ISlaughterBatchDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaughterBatchDo) Having(conds ...gen.Condition) ISlaughterBatchDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaughterBatchDo) Limit(limit int) ISlaughterBatchDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaughterBatchDo) Offset(offset int) ISlaughterBatchDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaughterBatchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaughterBatchDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaughterBatchDo) Unscoped() ISlaughterBatchDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaughterBatchDo) Create(values ...*slaughter.SlaughterBatch) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaughterBatchDo) CreateInBatches(values []*slaughter.SlaughterBatch, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaughterBatchDo) Save(values ...*slaughter.SlaughterBatch) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaughterBatchDo) First() (*slaughter.SlaughterBatch, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterBatch), nil
	}
}

func (s slaughterBatchDo) Take() (*slaughter.SlaughterBatch, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterBatch), nil
	}
}

func (s slaughterBatchDo) Last() (*slaughter.SlaughterBatch, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterBatch), nil
	}
}

func (s slaughterBatchDo) Find() ([]*slaughter.SlaughterBatch, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaughterBatch), err
}

func (s slaughterBatchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaughterBatch, err error) {
	buf := make([]*slaughter.SlaughterBatch, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaughterBatchDo) FindInBatches(result *[]*slaughter.SlaughterBatch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaughterBatchDo) Attrs(attrs ...field.AssignExpr) ISlaughterBatchDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaughterBatchDo) Assign(attrs ...field.AssignExpr) ISlaughterBatchDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaughterBatchDo) Joins(fields ...field.RelationField) ISlaughterBatchDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaughterBatchDo) Preload(fields ...field.RelationField) ISlaughterBatchDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaughterBatchDo) FirstOrInit() (*slaughter.SlaughterBatch, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterBatch), nil
	}
}

func (s slaughterBatchDo) FirstOrCreate() (*slaughter.SlaughterBatch, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaughterBatch), nil
	}
}

func (s slaughterBatchDo) FindByPage(offset int, limit int) (result []*slaughter.SlaughterBatch, count int64, err error) {
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

func (s slaughterBatchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaughterBatchDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaughterBatchDo) Delete(models ...*slaughter.SlaughterBatch) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaughterBatchDo) withDO(do gen.Dao) *slaughterBatchDo {
	s.DO = *do.(*gen.DO)
	return s
}
