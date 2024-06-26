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

func newLymphNode(db *gorm.DB, opts ...gen.DOOption) lymphNode {
	_lymphNode := lymphNode{}

	_lymphNode.lymphNodeDo.UseDB(db, opts...)
	_lymphNode.lymphNodeDo.UseModel(&model.LymphNode{})

	tableName := _lymphNode.lymphNodeDo.TableName()
	_lymphNode.ALL = field.NewAsterisk(tableName)
	_lymphNode.ID = field.NewInt32(tableName, "id")
	_lymphNode.Grade = field.NewString(tableName, "grade")
	_lymphNode.ImgPath = field.NewString(tableName, "img_path")
	_lymphNode.UploadTime = field.NewTime(tableName, "upload_time")

	_lymphNode.fillFieldMap()

	return _lymphNode
}

type lymphNode struct {
	lymphNodeDo lymphNodeDo

	ALL        field.Asterisk
	ID         field.Int32  // 自增主键
	Grade      field.String // 淋巴结等级
	ImgPath    field.String // 图片路径
	UploadTime field.Time   // 上传时间

	fieldMap map[string]field.Expr
}

func (l lymphNode) Table(newTableName string) *lymphNode {
	l.lymphNodeDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lymphNode) As(alias string) *lymphNode {
	l.lymphNodeDo.DO = *(l.lymphNodeDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lymphNode) updateTableName(table string) *lymphNode {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.Grade = field.NewString(table, "grade")
	l.ImgPath = field.NewString(table, "img_path")
	l.UploadTime = field.NewTime(table, "upload_time")

	l.fillFieldMap()

	return l
}

func (l *lymphNode) WithContext(ctx context.Context) ILymphNodeDo {
	return l.lymphNodeDo.WithContext(ctx)
}

func (l lymphNode) TableName() string { return l.lymphNodeDo.TableName() }

func (l lymphNode) Alias() string { return l.lymphNodeDo.Alias() }

func (l lymphNode) Columns(cols ...field.Expr) gen.Columns { return l.lymphNodeDo.Columns(cols...) }

func (l *lymphNode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lymphNode) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 4)
	l.fieldMap["id"] = l.ID
	l.fieldMap["grade"] = l.Grade
	l.fieldMap["img_path"] = l.ImgPath
	l.fieldMap["upload_time"] = l.UploadTime
}

func (l lymphNode) clone(db *gorm.DB) lymphNode {
	l.lymphNodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l lymphNode) replaceDB(db *gorm.DB) lymphNode {
	l.lymphNodeDo.ReplaceDB(db)
	return l
}

type lymphNodeDo struct{ gen.DO }

type ILymphNodeDo interface {
	gen.SubQuery
	Debug() ILymphNodeDo
	WithContext(ctx context.Context) ILymphNodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILymphNodeDo
	WriteDB() ILymphNodeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILymphNodeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILymphNodeDo
	Not(conds ...gen.Condition) ILymphNodeDo
	Or(conds ...gen.Condition) ILymphNodeDo
	Select(conds ...field.Expr) ILymphNodeDo
	Where(conds ...gen.Condition) ILymphNodeDo
	Order(conds ...field.Expr) ILymphNodeDo
	Distinct(cols ...field.Expr) ILymphNodeDo
	Omit(cols ...field.Expr) ILymphNodeDo
	Join(table schema.Tabler, on ...field.Expr) ILymphNodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILymphNodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILymphNodeDo
	Group(cols ...field.Expr) ILymphNodeDo
	Having(conds ...gen.Condition) ILymphNodeDo
	Limit(limit int) ILymphNodeDo
	Offset(offset int) ILymphNodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILymphNodeDo
	Unscoped() ILymphNodeDo
	Create(values ...*model.LymphNode) error
	CreateInBatches(values []*model.LymphNode, batchSize int) error
	Save(values ...*model.LymphNode) error
	First() (*model.LymphNode, error)
	Take() (*model.LymphNode, error)
	Last() (*model.LymphNode, error)
	Find() ([]*model.LymphNode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LymphNode, err error)
	FindInBatches(result *[]*model.LymphNode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LymphNode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILymphNodeDo
	Assign(attrs ...field.AssignExpr) ILymphNodeDo
	Joins(fields ...field.RelationField) ILymphNodeDo
	Preload(fields ...field.RelationField) ILymphNodeDo
	FirstOrInit() (*model.LymphNode, error)
	FirstOrCreate() (*model.LymphNode, error)
	FindByPage(offset int, limit int) (result []*model.LymphNode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILymphNodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l lymphNodeDo) Debug() ILymphNodeDo {
	return l.withDO(l.DO.Debug())
}

func (l lymphNodeDo) WithContext(ctx context.Context) ILymphNodeDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lymphNodeDo) ReadDB() ILymphNodeDo {
	return l.Clauses(dbresolver.Read)
}

func (l lymphNodeDo) WriteDB() ILymphNodeDo {
	return l.Clauses(dbresolver.Write)
}

func (l lymphNodeDo) Session(config *gorm.Session) ILymphNodeDo {
	return l.withDO(l.DO.Session(config))
}

func (l lymphNodeDo) Clauses(conds ...clause.Expression) ILymphNodeDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lymphNodeDo) Returning(value interface{}, columns ...string) ILymphNodeDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lymphNodeDo) Not(conds ...gen.Condition) ILymphNodeDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lymphNodeDo) Or(conds ...gen.Condition) ILymphNodeDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lymphNodeDo) Select(conds ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lymphNodeDo) Where(conds ...gen.Condition) ILymphNodeDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lymphNodeDo) Order(conds ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lymphNodeDo) Distinct(cols ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lymphNodeDo) Omit(cols ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lymphNodeDo) Join(table schema.Tabler, on ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lymphNodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lymphNodeDo) RightJoin(table schema.Tabler, on ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lymphNodeDo) Group(cols ...field.Expr) ILymphNodeDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lymphNodeDo) Having(conds ...gen.Condition) ILymphNodeDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lymphNodeDo) Limit(limit int) ILymphNodeDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lymphNodeDo) Offset(offset int) ILymphNodeDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lymphNodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILymphNodeDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lymphNodeDo) Unscoped() ILymphNodeDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lymphNodeDo) Create(values ...*model.LymphNode) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lymphNodeDo) CreateInBatches(values []*model.LymphNode, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lymphNodeDo) Save(values ...*model.LymphNode) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lymphNodeDo) First() (*model.LymphNode, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LymphNode), nil
	}
}

func (l lymphNodeDo) Take() (*model.LymphNode, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LymphNode), nil
	}
}

func (l lymphNodeDo) Last() (*model.LymphNode, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LymphNode), nil
	}
}

func (l lymphNodeDo) Find() ([]*model.LymphNode, error) {
	result, err := l.DO.Find()
	return result.([]*model.LymphNode), err
}

func (l lymphNodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LymphNode, err error) {
	buf := make([]*model.LymphNode, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lymphNodeDo) FindInBatches(result *[]*model.LymphNode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lymphNodeDo) Attrs(attrs ...field.AssignExpr) ILymphNodeDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lymphNodeDo) Assign(attrs ...field.AssignExpr) ILymphNodeDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lymphNodeDo) Joins(fields ...field.RelationField) ILymphNodeDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lymphNodeDo) Preload(fields ...field.RelationField) ILymphNodeDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lymphNodeDo) FirstOrInit() (*model.LymphNode, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LymphNode), nil
	}
}

func (l lymphNodeDo) FirstOrCreate() (*model.LymphNode, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LymphNode), nil
	}
}

func (l lymphNodeDo) FindByPage(offset int, limit int) (result []*model.LymphNode, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l lymphNodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lymphNodeDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lymphNodeDo) Delete(models ...*model.LymphNode) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lymphNodeDo) withDO(do gen.Dao) *lymphNodeDo {
	l.DO = *do.(*gen.DO)
	return l
}
