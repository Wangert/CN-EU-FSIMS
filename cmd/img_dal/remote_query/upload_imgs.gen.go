// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package remote_query

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

func newUploadImgs(db *gorm.DB, opts ...gen.DOOption) uploadImgs {
	_uploadImgs := uploadImgs{}

	_uploadImgs.uploadImgsDo.UseDB(db, opts...)
	_uploadImgs.uploadImgsDo.UseModel(&models.UploadImgs{})

	tableName := _uploadImgs.uploadImgsDo.TableName()
	_uploadImgs.ALL = field.NewAsterisk(tableName)
	_uploadImgs.Id = field.NewUint(tableName, "id")
	_uploadImgs.Filename = field.NewString(tableName, "filename")
	_uploadImgs.Source = field.NewString(tableName, "source")
	_uploadImgs.UploadTime = field.NewTime(tableName, "upload_time")
	_uploadImgs.Result = field.NewBytes(tableName, "result")

	_uploadImgs.fillFieldMap()

	return _uploadImgs
}

type uploadImgs struct {
	uploadImgsDo uploadImgsDo

	ALL        field.Asterisk
	Id         field.Uint
	Filename   field.String
	Source     field.String
	UploadTime field.Time
	Result     field.Bytes

	fieldMap map[string]field.Expr
}

func (u uploadImgs) Table(newTableName string) *uploadImgs {
	u.uploadImgsDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uploadImgs) As(alias string) *uploadImgs {
	u.uploadImgsDo.DO = *(u.uploadImgsDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uploadImgs) updateTableName(table string) *uploadImgs {
	u.ALL = field.NewAsterisk(table)
	u.Id = field.NewUint(table, "id")
	u.Filename = field.NewString(table, "filename")
	u.Source = field.NewString(table, "source")
	u.UploadTime = field.NewTime(table, "upload_time")
	u.Result = field.NewBytes(table, "result")

	u.fillFieldMap()

	return u
}

func (u *uploadImgs) WithContext(ctx context.Context) IUploadImgsDo {
	return u.uploadImgsDo.WithContext(ctx)
}

func (u uploadImgs) TableName() string { return u.uploadImgsDo.TableName() }

func (u uploadImgs) Alias() string { return u.uploadImgsDo.Alias() }

func (u uploadImgs) Columns(cols ...field.Expr) gen.Columns { return u.uploadImgsDo.Columns(cols...) }

func (u *uploadImgs) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uploadImgs) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["id"] = u.Id
	u.fieldMap["filename"] = u.Filename
	u.fieldMap["source"] = u.Source
	u.fieldMap["upload_time"] = u.UploadTime
	u.fieldMap["result"] = u.Result
}

func (u uploadImgs) clone(db *gorm.DB) uploadImgs {
	u.uploadImgsDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uploadImgs) replaceDB(db *gorm.DB) uploadImgs {
	u.uploadImgsDo.ReplaceDB(db)
	return u
}

type uploadImgsDo struct{ gen.DO }

type IUploadImgsDo interface {
	gen.SubQuery
	Debug() IUploadImgsDo
	WithContext(ctx context.Context) IUploadImgsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUploadImgsDo
	WriteDB() IUploadImgsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUploadImgsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUploadImgsDo
	Not(conds ...gen.Condition) IUploadImgsDo
	Or(conds ...gen.Condition) IUploadImgsDo
	Select(conds ...field.Expr) IUploadImgsDo
	Where(conds ...gen.Condition) IUploadImgsDo
	Order(conds ...field.Expr) IUploadImgsDo
	Distinct(cols ...field.Expr) IUploadImgsDo
	Omit(cols ...field.Expr) IUploadImgsDo
	Join(table schema.Tabler, on ...field.Expr) IUploadImgsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUploadImgsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUploadImgsDo
	Group(cols ...field.Expr) IUploadImgsDo
	Having(conds ...gen.Condition) IUploadImgsDo
	Limit(limit int) IUploadImgsDo
	Offset(offset int) IUploadImgsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUploadImgsDo
	Unscoped() IUploadImgsDo
	Create(values ...*models.UploadImgs) error
	CreateInBatches(values []*models.UploadImgs, batchSize int) error
	Save(values ...*models.UploadImgs) error
	First() (*models.UploadImgs, error)
	Take() (*models.UploadImgs, error)
	Last() (*models.UploadImgs, error)
	Find() ([]*models.UploadImgs, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UploadImgs, err error)
	FindInBatches(result *[]*models.UploadImgs, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.UploadImgs) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUploadImgsDo
	Assign(attrs ...field.AssignExpr) IUploadImgsDo
	Joins(fields ...field.RelationField) IUploadImgsDo
	Preload(fields ...field.RelationField) IUploadImgsDo
	FirstOrInit() (*models.UploadImgs, error)
	FirstOrCreate() (*models.UploadImgs, error)
	FindByPage(offset int, limit int) (result []*models.UploadImgs, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUploadImgsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u uploadImgsDo) Debug() IUploadImgsDo {
	return u.withDO(u.DO.Debug())
}

func (u uploadImgsDo) WithContext(ctx context.Context) IUploadImgsDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uploadImgsDo) ReadDB() IUploadImgsDo {
	return u.Clauses(dbresolver.Read)
}

func (u uploadImgsDo) WriteDB() IUploadImgsDo {
	return u.Clauses(dbresolver.Write)
}

func (u uploadImgsDo) Session(config *gorm.Session) IUploadImgsDo {
	return u.withDO(u.DO.Session(config))
}

func (u uploadImgsDo) Clauses(conds ...clause.Expression) IUploadImgsDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uploadImgsDo) Returning(value interface{}, columns ...string) IUploadImgsDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uploadImgsDo) Not(conds ...gen.Condition) IUploadImgsDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uploadImgsDo) Or(conds ...gen.Condition) IUploadImgsDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uploadImgsDo) Select(conds ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uploadImgsDo) Where(conds ...gen.Condition) IUploadImgsDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uploadImgsDo) Order(conds ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uploadImgsDo) Distinct(cols ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uploadImgsDo) Omit(cols ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uploadImgsDo) Join(table schema.Tabler, on ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uploadImgsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uploadImgsDo) RightJoin(table schema.Tabler, on ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uploadImgsDo) Group(cols ...field.Expr) IUploadImgsDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uploadImgsDo) Having(conds ...gen.Condition) IUploadImgsDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uploadImgsDo) Limit(limit int) IUploadImgsDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uploadImgsDo) Offset(offset int) IUploadImgsDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uploadImgsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUploadImgsDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uploadImgsDo) Unscoped() IUploadImgsDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uploadImgsDo) Create(values ...*models.UploadImgs) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uploadImgsDo) CreateInBatches(values []*models.UploadImgs, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uploadImgsDo) Save(values ...*models.UploadImgs) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uploadImgsDo) First() (*models.UploadImgs, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.UploadImgs), nil
	}
}

func (u uploadImgsDo) Take() (*models.UploadImgs, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.UploadImgs), nil
	}
}

func (u uploadImgsDo) Last() (*models.UploadImgs, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.UploadImgs), nil
	}
}

func (u uploadImgsDo) Find() ([]*models.UploadImgs, error) {
	result, err := u.DO.Find()
	return result.([]*models.UploadImgs), err
}

func (u uploadImgsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UploadImgs, err error) {
	buf := make([]*models.UploadImgs, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uploadImgsDo) FindInBatches(result *[]*models.UploadImgs, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uploadImgsDo) Attrs(attrs ...field.AssignExpr) IUploadImgsDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uploadImgsDo) Assign(attrs ...field.AssignExpr) IUploadImgsDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uploadImgsDo) Joins(fields ...field.RelationField) IUploadImgsDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uploadImgsDo) Preload(fields ...field.RelationField) IUploadImgsDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uploadImgsDo) FirstOrInit() (*models.UploadImgs, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.UploadImgs), nil
	}
}

func (u uploadImgsDo) FirstOrCreate() (*models.UploadImgs, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.UploadImgs), nil
	}
}

func (u uploadImgsDo) FindByPage(offset int, limit int) (result []*models.UploadImgs, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u uploadImgsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uploadImgsDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uploadImgsDo) Delete(models ...*models.UploadImgs) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uploadImgsDo) withDO(do gen.Dao) *uploadImgsDo {
	u.DO = *do.(*gen.DO)
	return u
}
