// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package imgquery

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"CN-EU-FSIMS/cmd/img_dal/model"
)

func newSpectra(db *gorm.DB, opts ...gen.DOOption) spectra {
	_spectra := spectra{}

	_spectra.spectraDo.UseDB(db, opts...)
	_spectra.spectraDo.UseModel(&model.Spectra{})

	tableName := _spectra.spectraDo.TableName()
	_spectra.ALL = field.NewAsterisk(tableName)
	_spectra.ID = field.NewInt32(tableName, "id")
	_spectra.Data = field.NewString(tableName, "data")
	_spectra.UploadTime = field.NewTime(tableName, "upload_time")

	_spectra.fillFieldMap()

	return _spectra
}

type spectra struct {
	spectraDo spectraDo

	ALL        field.Asterisk
	ID         field.Int32
	Data       field.String
	UploadTime field.Time

	fieldMap map[string]field.Expr
}

func (s spectra) Table(newTableName string) *spectra {
	s.spectraDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s spectra) As(alias string) *spectra {
	s.spectraDo.DO = *(s.spectraDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *spectra) updateTableName(table string) *spectra {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Data = field.NewString(table, "data")
	s.UploadTime = field.NewTime(table, "upload_time")

	s.fillFieldMap()

	return s
}

func (s *spectra) WithContext(ctx context.Context) ISpectraDo { return s.spectraDo.WithContext(ctx) }

func (s spectra) TableName() string { return s.spectraDo.TableName() }

func (s spectra) Alias() string { return s.spectraDo.Alias() }

func (s spectra) Columns(cols ...field.Expr) gen.Columns { return s.spectraDo.Columns(cols...) }

func (s *spectra) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *spectra) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.ID
	s.fieldMap["data"] = s.Data
	s.fieldMap["upload_time"] = s.UploadTime
}

func (s spectra) clone(db *gorm.DB) spectra {
	s.spectraDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s spectra) replaceDB(db *gorm.DB) spectra {
	s.spectraDo.ReplaceDB(db)
	return s
}

type spectraDo struct{ gen.DO }

type ISpectraDo interface {
	gen.SubQuery
	Debug() ISpectraDo
	WithContext(ctx context.Context) ISpectraDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISpectraDo
	WriteDB() ISpectraDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISpectraDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISpectraDo
	Not(conds ...gen.Condition) ISpectraDo
	Or(conds ...gen.Condition) ISpectraDo
	Select(conds ...field.Expr) ISpectraDo
	Where(conds ...gen.Condition) ISpectraDo
	Order(conds ...field.Expr) ISpectraDo
	Distinct(cols ...field.Expr) ISpectraDo
	Omit(cols ...field.Expr) ISpectraDo
	Join(table schema.Tabler, on ...field.Expr) ISpectraDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISpectraDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISpectraDo
	Group(cols ...field.Expr) ISpectraDo
	Having(conds ...gen.Condition) ISpectraDo
	Limit(limit int) ISpectraDo
	Offset(offset int) ISpectraDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISpectraDo
	Unscoped() ISpectraDo
	Create(values ...*model.Spectra) error
	CreateInBatches(values []*model.Spectra, batchSize int) error
	Save(values ...*model.Spectra) error
	First() (*model.Spectra, error)
	Take() (*model.Spectra, error)
	Last() (*model.Spectra, error)
	Find() ([]*model.Spectra, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Spectra, err error)
	FindInBatches(result *[]*model.Spectra, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Spectra) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISpectraDo
	Assign(attrs ...field.AssignExpr) ISpectraDo
	Joins(fields ...field.RelationField) ISpectraDo
	Preload(fields ...field.RelationField) ISpectraDo
	FirstOrInit() (*model.Spectra, error)
	FirstOrCreate() (*model.Spectra, error)
	FindByPage(offset int, limit int) (result []*model.Spectra, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISpectraDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s spectraDo) Debug() ISpectraDo {
	return s.withDO(s.DO.Debug())
}

func (s spectraDo) WithContext(ctx context.Context) ISpectraDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s spectraDo) ReadDB() ISpectraDo {
	return s.Clauses(dbresolver.Read)
}

func (s spectraDo) WriteDB() ISpectraDo {
	return s.Clauses(dbresolver.Write)
}

func (s spectraDo) Session(config *gorm.Session) ISpectraDo {
	return s.withDO(s.DO.Session(config))
}

func (s spectraDo) Clauses(conds ...clause.Expression) ISpectraDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s spectraDo) Returning(value interface{}, columns ...string) ISpectraDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s spectraDo) Not(conds ...gen.Condition) ISpectraDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s spectraDo) Or(conds ...gen.Condition) ISpectraDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s spectraDo) Select(conds ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s spectraDo) Where(conds ...gen.Condition) ISpectraDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s spectraDo) Order(conds ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s spectraDo) Distinct(cols ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s spectraDo) Omit(cols ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s spectraDo) Join(table schema.Tabler, on ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s spectraDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s spectraDo) RightJoin(table schema.Tabler, on ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s spectraDo) Group(cols ...field.Expr) ISpectraDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s spectraDo) Having(conds ...gen.Condition) ISpectraDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s spectraDo) Limit(limit int) ISpectraDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s spectraDo) Offset(offset int) ISpectraDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s spectraDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISpectraDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s spectraDo) Unscoped() ISpectraDo {
	return s.withDO(s.DO.Unscoped())
}

func (s spectraDo) Create(values ...*model.Spectra) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s spectraDo) CreateInBatches(values []*model.Spectra, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s spectraDo) Save(values ...*model.Spectra) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s spectraDo) First() (*model.Spectra, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spectra), nil
	}
}

func (s spectraDo) Take() (*model.Spectra, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spectra), nil
	}
}

func (s spectraDo) Last() (*model.Spectra, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spectra), nil
	}
}

func (s spectraDo) Find() ([]*model.Spectra, error) {
	result, err := s.DO.Find()
	return result.([]*model.Spectra), err
}

func (s spectraDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Spectra, err error) {
	buf := make([]*model.Spectra, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s spectraDo) FindInBatches(result *[]*model.Spectra, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s spectraDo) Attrs(attrs ...field.AssignExpr) ISpectraDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s spectraDo) Assign(attrs ...field.AssignExpr) ISpectraDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s spectraDo) Joins(fields ...field.RelationField) ISpectraDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s spectraDo) Preload(fields ...field.RelationField) ISpectraDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s spectraDo) FirstOrInit() (*model.Spectra, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spectra), nil
	}
}

func (s spectraDo) FirstOrCreate() (*model.Spectra, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Spectra), nil
	}
}

func (s spectraDo) FindByPage(offset int, limit int) (result []*model.Spectra, count int64, err error) {
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

func (s spectraDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s spectraDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s spectraDo) Delete(models ...*model.Spectra) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *spectraDo) withDO(do gen.Dao) *spectraDo {
	s.DO = *do.(*gen.DO)
	return s
}
