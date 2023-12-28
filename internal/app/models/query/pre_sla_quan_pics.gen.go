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

func newPreSlaQuanPic(db *gorm.DB, opts ...gen.DOOption) preSlaQuanPic {
	_preSlaQuanPic := preSlaQuanPic{}

	_preSlaQuanPic.preSlaQuanPicDo.UseDB(db, opts...)
	_preSlaQuanPic.preSlaQuanPicDo.UseModel(&slaughter.PreSlaQuanPic{})

	tableName := _preSlaQuanPic.preSlaQuanPicDo.TableName()
	_preSlaQuanPic.ALL = field.NewAsterisk(tableName)
	_preSlaQuanPic.ID = field.NewUint(tableName, "id")
	_preSlaQuanPic.CreatedAt = field.NewTime(tableName, "created_at")
	_preSlaQuanPic.UpdatedAt = field.NewTime(tableName, "updated_at")
	_preSlaQuanPic.DeletedAt = field.NewField(tableName, "deleted_at")
	_preSlaQuanPic.SlaInfoMonID = field.NewUint(tableName, "sla_info_mon_id")
	_preSlaQuanPic.PreSlaQuanPic1 = field.NewString(tableName, "pre_sla_quan_pic1")
	_preSlaQuanPic.PreSlaQuanPic2 = field.NewString(tableName, "pre_sla_quan_pic2")
	_preSlaQuanPic.PreSlaQuanPic3 = field.NewString(tableName, "pre_sla_quan_pic3")
	_preSlaQuanPic.PreSlaQuanPic4 = field.NewString(tableName, "pre_sla_quan_pic4")
	_preSlaQuanPic.PreSlaQuanPic5 = field.NewString(tableName, "pre_sla_quan_pic5")
	_preSlaQuanPic.PreSlaQuanPic6 = field.NewString(tableName, "pre_sla_quan_pic6")
	_preSlaQuanPic.PreSlaQuanPic7 = field.NewString(tableName, "pre_sla_quan_pic7")
	_preSlaQuanPic.PreSlaQuanPic8 = field.NewString(tableName, "pre_sla_quan_pic8")
	_preSlaQuanPic.PreSlaQuanPic9 = field.NewString(tableName, "pre_sla_quan_pic9")

	_preSlaQuanPic.fillFieldMap()

	return _preSlaQuanPic
}

type preSlaQuanPic struct {
	preSlaQuanPicDo preSlaQuanPicDo

	ALL            field.Asterisk
	ID             field.Uint
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	SlaInfoMonID   field.Uint
	PreSlaQuanPic1 field.String
	PreSlaQuanPic2 field.String
	PreSlaQuanPic3 field.String
	PreSlaQuanPic4 field.String
	PreSlaQuanPic5 field.String
	PreSlaQuanPic6 field.String
	PreSlaQuanPic7 field.String
	PreSlaQuanPic8 field.String
	PreSlaQuanPic9 field.String

	fieldMap map[string]field.Expr
}

func (p preSlaQuanPic) Table(newTableName string) *preSlaQuanPic {
	p.preSlaQuanPicDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p preSlaQuanPic) As(alias string) *preSlaQuanPic {
	p.preSlaQuanPicDo.DO = *(p.preSlaQuanPicDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *preSlaQuanPic) updateTableName(table string) *preSlaQuanPic {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.SlaInfoMonID = field.NewUint(table, "sla_info_mon_id")
	p.PreSlaQuanPic1 = field.NewString(table, "pre_sla_quan_pic1")
	p.PreSlaQuanPic2 = field.NewString(table, "pre_sla_quan_pic2")
	p.PreSlaQuanPic3 = field.NewString(table, "pre_sla_quan_pic3")
	p.PreSlaQuanPic4 = field.NewString(table, "pre_sla_quan_pic4")
	p.PreSlaQuanPic5 = field.NewString(table, "pre_sla_quan_pic5")
	p.PreSlaQuanPic6 = field.NewString(table, "pre_sla_quan_pic6")
	p.PreSlaQuanPic7 = field.NewString(table, "pre_sla_quan_pic7")
	p.PreSlaQuanPic8 = field.NewString(table, "pre_sla_quan_pic8")
	p.PreSlaQuanPic9 = field.NewString(table, "pre_sla_quan_pic9")

	p.fillFieldMap()

	return p
}

func (p *preSlaQuanPic) WithContext(ctx context.Context) IPreSlaQuanPicDo {
	return p.preSlaQuanPicDo.WithContext(ctx)
}

func (p preSlaQuanPic) TableName() string { return p.preSlaQuanPicDo.TableName() }

func (p preSlaQuanPic) Alias() string { return p.preSlaQuanPicDo.Alias() }

func (p preSlaQuanPic) Columns(cols ...field.Expr) gen.Columns {
	return p.preSlaQuanPicDo.Columns(cols...)
}

func (p *preSlaQuanPic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *preSlaQuanPic) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 14)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["sla_info_mon_id"] = p.SlaInfoMonID
	p.fieldMap["pre_sla_quan_pic1"] = p.PreSlaQuanPic1
	p.fieldMap["pre_sla_quan_pic2"] = p.PreSlaQuanPic2
	p.fieldMap["pre_sla_quan_pic3"] = p.PreSlaQuanPic3
	p.fieldMap["pre_sla_quan_pic4"] = p.PreSlaQuanPic4
	p.fieldMap["pre_sla_quan_pic5"] = p.PreSlaQuanPic5
	p.fieldMap["pre_sla_quan_pic6"] = p.PreSlaQuanPic6
	p.fieldMap["pre_sla_quan_pic7"] = p.PreSlaQuanPic7
	p.fieldMap["pre_sla_quan_pic8"] = p.PreSlaQuanPic8
	p.fieldMap["pre_sla_quan_pic9"] = p.PreSlaQuanPic9
}

func (p preSlaQuanPic) clone(db *gorm.DB) preSlaQuanPic {
	p.preSlaQuanPicDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p preSlaQuanPic) replaceDB(db *gorm.DB) preSlaQuanPic {
	p.preSlaQuanPicDo.ReplaceDB(db)
	return p
}

type preSlaQuanPicDo struct{ gen.DO }

type IPreSlaQuanPicDo interface {
	gen.SubQuery
	Debug() IPreSlaQuanPicDo
	WithContext(ctx context.Context) IPreSlaQuanPicDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPreSlaQuanPicDo
	WriteDB() IPreSlaQuanPicDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPreSlaQuanPicDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPreSlaQuanPicDo
	Not(conds ...gen.Condition) IPreSlaQuanPicDo
	Or(conds ...gen.Condition) IPreSlaQuanPicDo
	Select(conds ...field.Expr) IPreSlaQuanPicDo
	Where(conds ...gen.Condition) IPreSlaQuanPicDo
	Order(conds ...field.Expr) IPreSlaQuanPicDo
	Distinct(cols ...field.Expr) IPreSlaQuanPicDo
	Omit(cols ...field.Expr) IPreSlaQuanPicDo
	Join(table schema.Tabler, on ...field.Expr) IPreSlaQuanPicDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPreSlaQuanPicDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPreSlaQuanPicDo
	Group(cols ...field.Expr) IPreSlaQuanPicDo
	Having(conds ...gen.Condition) IPreSlaQuanPicDo
	Limit(limit int) IPreSlaQuanPicDo
	Offset(offset int) IPreSlaQuanPicDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPreSlaQuanPicDo
	Unscoped() IPreSlaQuanPicDo
	Create(values ...*slaughter.PreSlaQuanPic) error
	CreateInBatches(values []*slaughter.PreSlaQuanPic, batchSize int) error
	Save(values ...*slaughter.PreSlaQuanPic) error
	First() (*slaughter.PreSlaQuanPic, error)
	Take() (*slaughter.PreSlaQuanPic, error)
	Last() (*slaughter.PreSlaQuanPic, error)
	Find() ([]*slaughter.PreSlaQuanPic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.PreSlaQuanPic, err error)
	FindInBatches(result *[]*slaughter.PreSlaQuanPic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.PreSlaQuanPic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPreSlaQuanPicDo
	Assign(attrs ...field.AssignExpr) IPreSlaQuanPicDo
	Joins(fields ...field.RelationField) IPreSlaQuanPicDo
	Preload(fields ...field.RelationField) IPreSlaQuanPicDo
	FirstOrInit() (*slaughter.PreSlaQuanPic, error)
	FirstOrCreate() (*slaughter.PreSlaQuanPic, error)
	FindByPage(offset int, limit int) (result []*slaughter.PreSlaQuanPic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPreSlaQuanPicDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p preSlaQuanPicDo) Debug() IPreSlaQuanPicDo {
	return p.withDO(p.DO.Debug())
}

func (p preSlaQuanPicDo) WithContext(ctx context.Context) IPreSlaQuanPicDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p preSlaQuanPicDo) ReadDB() IPreSlaQuanPicDo {
	return p.Clauses(dbresolver.Read)
}

func (p preSlaQuanPicDo) WriteDB() IPreSlaQuanPicDo {
	return p.Clauses(dbresolver.Write)
}

func (p preSlaQuanPicDo) Session(config *gorm.Session) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Session(config))
}

func (p preSlaQuanPicDo) Clauses(conds ...clause.Expression) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p preSlaQuanPicDo) Returning(value interface{}, columns ...string) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p preSlaQuanPicDo) Not(conds ...gen.Condition) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p preSlaQuanPicDo) Or(conds ...gen.Condition) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p preSlaQuanPicDo) Select(conds ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p preSlaQuanPicDo) Where(conds ...gen.Condition) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p preSlaQuanPicDo) Order(conds ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p preSlaQuanPicDo) Distinct(cols ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p preSlaQuanPicDo) Omit(cols ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p preSlaQuanPicDo) Join(table schema.Tabler, on ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p preSlaQuanPicDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p preSlaQuanPicDo) RightJoin(table schema.Tabler, on ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p preSlaQuanPicDo) Group(cols ...field.Expr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p preSlaQuanPicDo) Having(conds ...gen.Condition) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p preSlaQuanPicDo) Limit(limit int) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p preSlaQuanPicDo) Offset(offset int) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p preSlaQuanPicDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p preSlaQuanPicDo) Unscoped() IPreSlaQuanPicDo {
	return p.withDO(p.DO.Unscoped())
}

func (p preSlaQuanPicDo) Create(values ...*slaughter.PreSlaQuanPic) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p preSlaQuanPicDo) CreateInBatches(values []*slaughter.PreSlaQuanPic, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p preSlaQuanPicDo) Save(values ...*slaughter.PreSlaQuanPic) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p preSlaQuanPicDo) First() (*slaughter.PreSlaQuanPic, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PreSlaQuanPic), nil
	}
}

func (p preSlaQuanPicDo) Take() (*slaughter.PreSlaQuanPic, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PreSlaQuanPic), nil
	}
}

func (p preSlaQuanPicDo) Last() (*slaughter.PreSlaQuanPic, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PreSlaQuanPic), nil
	}
}

func (p preSlaQuanPicDo) Find() ([]*slaughter.PreSlaQuanPic, error) {
	result, err := p.DO.Find()
	return result.([]*slaughter.PreSlaQuanPic), err
}

func (p preSlaQuanPicDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.PreSlaQuanPic, err error) {
	buf := make([]*slaughter.PreSlaQuanPic, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p preSlaQuanPicDo) FindInBatches(result *[]*slaughter.PreSlaQuanPic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p preSlaQuanPicDo) Attrs(attrs ...field.AssignExpr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p preSlaQuanPicDo) Assign(attrs ...field.AssignExpr) IPreSlaQuanPicDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p preSlaQuanPicDo) Joins(fields ...field.RelationField) IPreSlaQuanPicDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p preSlaQuanPicDo) Preload(fields ...field.RelationField) IPreSlaQuanPicDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p preSlaQuanPicDo) FirstOrInit() (*slaughter.PreSlaQuanPic, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PreSlaQuanPic), nil
	}
}

func (p preSlaQuanPicDo) FirstOrCreate() (*slaughter.PreSlaQuanPic, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.PreSlaQuanPic), nil
	}
}

func (p preSlaQuanPicDo) FindByPage(offset int, limit int) (result []*slaughter.PreSlaQuanPic, count int64, err error) {
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

func (p preSlaQuanPicDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p preSlaQuanPicDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p preSlaQuanPicDo) Delete(models ...*slaughter.PreSlaQuanPic) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *preSlaQuanPicDo) withDO(do gen.Dao) *preSlaQuanPicDo {
	p.DO = *do.(*gen.DO)
	return p
}
