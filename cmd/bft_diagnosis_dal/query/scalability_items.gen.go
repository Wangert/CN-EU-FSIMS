// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"CN-EU-FSIMS/cmd/bft_diagnosis_dal/model"
)

func newScalabilityItem(db *gorm.DB, opts ...gen.DOOption) scalabilityItem {
	_scalabilityItem := scalabilityItem{}

	_scalabilityItem.scalabilityItemDo.UseDB(db, opts...)
	_scalabilityItem.scalabilityItemDo.UseModel(&model.ScalabilityItem{})

	tableName := _scalabilityItem.scalabilityItemDo.TableName()
	_scalabilityItem.ALL = field.NewAsterisk(tableName)
	_scalabilityItem.ID = field.NewInt32(tableName, "id")
	_scalabilityItem.Item = field.NewString(tableName, "item")

	_scalabilityItem.fillFieldMap()

	return _scalabilityItem
}

type scalabilityItem struct {
	scalabilityItemDo scalabilityItemDo

	ALL  field.Asterisk
	ID   field.Int32
	Item field.String

	fieldMap map[string]field.Expr
}

func (s scalabilityItem) Table(newTableName string) *scalabilityItem {
	s.scalabilityItemDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s scalabilityItem) As(alias string) *scalabilityItem {
	s.scalabilityItemDo.DO = *(s.scalabilityItemDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *scalabilityItem) updateTableName(table string) *scalabilityItem {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Item = field.NewString(table, "item")

	s.fillFieldMap()

	return s
}

func (s *scalabilityItem) WithContext(ctx context.Context) IScalabilityItemDo {
	return s.scalabilityItemDo.WithContext(ctx)
}

func (s scalabilityItem) TableName() string { return s.scalabilityItemDo.TableName() }

func (s scalabilityItem) Alias() string { return s.scalabilityItemDo.Alias() }

func (s scalabilityItem) Columns(cols ...field.Expr) gen.Columns {
	return s.scalabilityItemDo.Columns(cols...)
}

func (s *scalabilityItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *scalabilityItem) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["id"] = s.ID
	s.fieldMap["item"] = s.Item
}

func (s scalabilityItem) clone(db *gorm.DB) scalabilityItem {
	s.scalabilityItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s scalabilityItem) replaceDB(db *gorm.DB) scalabilityItem {
	s.scalabilityItemDo.ReplaceDB(db)
	return s
}

type scalabilityItemDo struct{ gen.DO }

type IScalabilityItemDo interface {
	gen.SubQuery
	Debug() IScalabilityItemDo
	WithContext(ctx context.Context) IScalabilityItemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IScalabilityItemDo
	WriteDB() IScalabilityItemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IScalabilityItemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IScalabilityItemDo
	Not(conds ...gen.Condition) IScalabilityItemDo
	Or(conds ...gen.Condition) IScalabilityItemDo
	Select(conds ...field.Expr) IScalabilityItemDo
	Where(conds ...gen.Condition) IScalabilityItemDo
	Order(conds ...field.Expr) IScalabilityItemDo
	Distinct(cols ...field.Expr) IScalabilityItemDo
	Omit(cols ...field.Expr) IScalabilityItemDo
	Join(table schema.Tabler, on ...field.Expr) IScalabilityItemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IScalabilityItemDo
	RightJoin(table schema.Tabler, on ...field.Expr) IScalabilityItemDo
	Group(cols ...field.Expr) IScalabilityItemDo
	Having(conds ...gen.Condition) IScalabilityItemDo
	Limit(limit int) IScalabilityItemDo
	Offset(offset int) IScalabilityItemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IScalabilityItemDo
	Unscoped() IScalabilityItemDo
	Create(values ...*model.ScalabilityItem) error
	CreateInBatches(values []*model.ScalabilityItem, batchSize int) error
	Save(values ...*model.ScalabilityItem) error
	First() (*model.ScalabilityItem, error)
	Take() (*model.ScalabilityItem, error)
	Last() (*model.ScalabilityItem, error)
	Find() ([]*model.ScalabilityItem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ScalabilityItem, err error)
	FindInBatches(result *[]*model.ScalabilityItem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ScalabilityItem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IScalabilityItemDo
	Assign(attrs ...field.AssignExpr) IScalabilityItemDo
	Joins(fields ...field.RelationField) IScalabilityItemDo
	Preload(fields ...field.RelationField) IScalabilityItemDo
	FirstOrInit() (*model.ScalabilityItem, error)
	FirstOrCreate() (*model.ScalabilityItem, error)
	FindByPage(offset int, limit int) (result []*model.ScalabilityItem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IScalabilityItemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s scalabilityItemDo) Debug() IScalabilityItemDo {
	return s.withDO(s.DO.Debug())
}

func (s scalabilityItemDo) WithContext(ctx context.Context) IScalabilityItemDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s scalabilityItemDo) ReadDB() IScalabilityItemDo {
	return s.Clauses(dbresolver.Read)
}

func (s scalabilityItemDo) WriteDB() IScalabilityItemDo {
	return s.Clauses(dbresolver.Write)
}

func (s scalabilityItemDo) Session(config *gorm.Session) IScalabilityItemDo {
	return s.withDO(s.DO.Session(config))
}

func (s scalabilityItemDo) Clauses(conds ...clause.Expression) IScalabilityItemDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s scalabilityItemDo) Returning(value interface{}, columns ...string) IScalabilityItemDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s scalabilityItemDo) Not(conds ...gen.Condition) IScalabilityItemDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s scalabilityItemDo) Or(conds ...gen.Condition) IScalabilityItemDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s scalabilityItemDo) Select(conds ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s scalabilityItemDo) Where(conds ...gen.Condition) IScalabilityItemDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s scalabilityItemDo) Order(conds ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s scalabilityItemDo) Distinct(cols ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s scalabilityItemDo) Omit(cols ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s scalabilityItemDo) Join(table schema.Tabler, on ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s scalabilityItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s scalabilityItemDo) RightJoin(table schema.Tabler, on ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s scalabilityItemDo) Group(cols ...field.Expr) IScalabilityItemDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s scalabilityItemDo) Having(conds ...gen.Condition) IScalabilityItemDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s scalabilityItemDo) Limit(limit int) IScalabilityItemDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s scalabilityItemDo) Offset(offset int) IScalabilityItemDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s scalabilityItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IScalabilityItemDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s scalabilityItemDo) Unscoped() IScalabilityItemDo {
	return s.withDO(s.DO.Unscoped())
}

func (s scalabilityItemDo) Create(values ...*model.ScalabilityItem) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s scalabilityItemDo) CreateInBatches(values []*model.ScalabilityItem, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s scalabilityItemDo) Save(values ...*model.ScalabilityItem) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s scalabilityItemDo) First() (*model.ScalabilityItem, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScalabilityItem), nil
	}
}

func (s scalabilityItemDo) Take() (*model.ScalabilityItem, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScalabilityItem), nil
	}
}

func (s scalabilityItemDo) Last() (*model.ScalabilityItem, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScalabilityItem), nil
	}
}

func (s scalabilityItemDo) Find() ([]*model.ScalabilityItem, error) {
	result, err := s.DO.Find()
	return result.([]*model.ScalabilityItem), err
}

func (s scalabilityItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ScalabilityItem, err error) {
	buf := make([]*model.ScalabilityItem, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s scalabilityItemDo) FindInBatches(result *[]*model.ScalabilityItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s scalabilityItemDo) Attrs(attrs ...field.AssignExpr) IScalabilityItemDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s scalabilityItemDo) Assign(attrs ...field.AssignExpr) IScalabilityItemDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s scalabilityItemDo) Joins(fields ...field.RelationField) IScalabilityItemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s scalabilityItemDo) Preload(fields ...field.RelationField) IScalabilityItemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s scalabilityItemDo) FirstOrInit() (*model.ScalabilityItem, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScalabilityItem), nil
	}
}

func (s scalabilityItemDo) FirstOrCreate() (*model.ScalabilityItem, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScalabilityItem), nil
	}
}

func (s scalabilityItemDo) FindByPage(offset int, limit int) (result []*model.ScalabilityItem, count int64, err error) {
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

func (s scalabilityItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s scalabilityItemDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s scalabilityItemDo) Delete(models ...*model.ScalabilityItem) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *scalabilityItemDo) withDO(do gen.Dao) *scalabilityItemDo {
	s.DO = *do.(*gen.DO)
	return s
}
