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

func newSlaShop(db *gorm.DB, opts ...gen.DOOption) slaShop {
	_slaShop := slaShop{}

	_slaShop.slaShopDo.UseDB(db, opts...)
	_slaShop.slaShopDo.UseModel(&slaughter.SlaShop{})

	tableName := _slaShop.slaShopDo.TableName()
	_slaShop.ALL = field.NewAsterisk(tableName)
	_slaShop.ID = field.NewUint(tableName, "id")
	_slaShop.CreatedAt = field.NewTime(tableName, "created_at")
	_slaShop.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slaShop.DeletedAt = field.NewField(tableName, "deleted_at")
	_slaShop.HouseNumber = field.NewString(tableName, "house_number")
	_slaShop.TimeRecordAt = field.NewTime(tableName, "time_record_at")
	_slaShop.SlaShop1 = field.NewFloat32(tableName, "sla_shop1")
	_slaShop.SlaShop2 = field.NewFloat32(tableName, "sla_shop2")
	_slaShop.SlaShop3 = field.NewFloat32(tableName, "sla_shop3")
	_slaShop.SlaShop4 = field.NewFloat32(tableName, "sla_shop4")
	_slaShop.SlaShop5 = field.NewFloat32(tableName, "sla_shop5")
	_slaShop.SlaShop6 = field.NewFloat32(tableName, "sla_shop6")
	_slaShop.SlaShop7 = field.NewFloat32(tableName, "sla_shop7")
	_slaShop.SlaShop8 = field.NewFloat32(tableName, "sla_shop8")
	_slaShop.SlaShop9 = field.NewFloat32(tableName, "sla_shop9")
	_slaShop.SlaShop10 = field.NewFloat32(tableName, "sla_shop10")
	_slaShop.SlaShop11 = field.NewFloat32(tableName, "sla_shop11")
	_slaShop.SlaShop12 = field.NewFloat32(tableName, "sla_shop12")
	_slaShop.SlaShop13 = field.NewFloat32(tableName, "sla_shop13")
	_slaShop.SlaShop14 = field.NewFloat32(tableName, "sla_shop14")

	_slaShop.fillFieldMap()

	return _slaShop
}

type slaShop struct {
	slaShopDo slaShopDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	HouseNumber  field.String
	TimeRecordAt field.Time
	SlaShop1     field.Float32
	SlaShop2     field.Float32
	SlaShop3     field.Float32
	SlaShop4     field.Float32
	SlaShop5     field.Float32
	SlaShop6     field.Float32
	SlaShop7     field.Float32
	SlaShop8     field.Float32
	SlaShop9     field.Float32
	SlaShop10    field.Float32
	SlaShop11    field.Float32
	SlaShop12    field.Float32
	SlaShop13    field.Float32
	SlaShop14    field.Float32

	fieldMap map[string]field.Expr
}

func (s slaShop) Table(newTableName string) *slaShop {
	s.slaShopDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slaShop) As(alias string) *slaShop {
	s.slaShopDo.DO = *(s.slaShopDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slaShop) updateTableName(table string) *slaShop {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.HouseNumber = field.NewString(table, "house_number")
	s.TimeRecordAt = field.NewTime(table, "time_record_at")
	s.SlaShop1 = field.NewFloat32(table, "sla_shop1")
	s.SlaShop2 = field.NewFloat32(table, "sla_shop2")
	s.SlaShop3 = field.NewFloat32(table, "sla_shop3")
	s.SlaShop4 = field.NewFloat32(table, "sla_shop4")
	s.SlaShop5 = field.NewFloat32(table, "sla_shop5")
	s.SlaShop6 = field.NewFloat32(table, "sla_shop6")
	s.SlaShop7 = field.NewFloat32(table, "sla_shop7")
	s.SlaShop8 = field.NewFloat32(table, "sla_shop8")
	s.SlaShop9 = field.NewFloat32(table, "sla_shop9")
	s.SlaShop10 = field.NewFloat32(table, "sla_shop10")
	s.SlaShop11 = field.NewFloat32(table, "sla_shop11")
	s.SlaShop12 = field.NewFloat32(table, "sla_shop12")
	s.SlaShop13 = field.NewFloat32(table, "sla_shop13")
	s.SlaShop14 = field.NewFloat32(table, "sla_shop14")

	s.fillFieldMap()

	return s
}

func (s *slaShop) WithContext(ctx context.Context) ISlaShopDo { return s.slaShopDo.WithContext(ctx) }

func (s slaShop) TableName() string { return s.slaShopDo.TableName() }

func (s slaShop) Alias() string { return s.slaShopDo.Alias() }

func (s slaShop) Columns(cols ...field.Expr) gen.Columns { return s.slaShopDo.Columns(cols...) }

func (s *slaShop) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slaShop) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["house_number"] = s.HouseNumber
	s.fieldMap["time_record_at"] = s.TimeRecordAt
	s.fieldMap["sla_shop1"] = s.SlaShop1
	s.fieldMap["sla_shop2"] = s.SlaShop2
	s.fieldMap["sla_shop3"] = s.SlaShop3
	s.fieldMap["sla_shop4"] = s.SlaShop4
	s.fieldMap["sla_shop5"] = s.SlaShop5
	s.fieldMap["sla_shop6"] = s.SlaShop6
	s.fieldMap["sla_shop7"] = s.SlaShop7
	s.fieldMap["sla_shop8"] = s.SlaShop8
	s.fieldMap["sla_shop9"] = s.SlaShop9
	s.fieldMap["sla_shop10"] = s.SlaShop10
	s.fieldMap["sla_shop11"] = s.SlaShop11
	s.fieldMap["sla_shop12"] = s.SlaShop12
	s.fieldMap["sla_shop13"] = s.SlaShop13
	s.fieldMap["sla_shop14"] = s.SlaShop14
}

func (s slaShop) clone(db *gorm.DB) slaShop {
	s.slaShopDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slaShop) replaceDB(db *gorm.DB) slaShop {
	s.slaShopDo.ReplaceDB(db)
	return s
}

type slaShopDo struct{ gen.DO }

type ISlaShopDo interface {
	gen.SubQuery
	Debug() ISlaShopDo
	WithContext(ctx context.Context) ISlaShopDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlaShopDo
	WriteDB() ISlaShopDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlaShopDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlaShopDo
	Not(conds ...gen.Condition) ISlaShopDo
	Or(conds ...gen.Condition) ISlaShopDo
	Select(conds ...field.Expr) ISlaShopDo
	Where(conds ...gen.Condition) ISlaShopDo
	Order(conds ...field.Expr) ISlaShopDo
	Distinct(cols ...field.Expr) ISlaShopDo
	Omit(cols ...field.Expr) ISlaShopDo
	Join(table schema.Tabler, on ...field.Expr) ISlaShopDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlaShopDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlaShopDo
	Group(cols ...field.Expr) ISlaShopDo
	Having(conds ...gen.Condition) ISlaShopDo
	Limit(limit int) ISlaShopDo
	Offset(offset int) ISlaShopDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaShopDo
	Unscoped() ISlaShopDo
	Create(values ...*slaughter.SlaShop) error
	CreateInBatches(values []*slaughter.SlaShop, batchSize int) error
	Save(values ...*slaughter.SlaShop) error
	First() (*slaughter.SlaShop, error)
	Take() (*slaughter.SlaShop, error)
	Last() (*slaughter.SlaShop, error)
	Find() ([]*slaughter.SlaShop, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaShop, err error)
	FindInBatches(result *[]*slaughter.SlaShop, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.SlaShop) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlaShopDo
	Assign(attrs ...field.AssignExpr) ISlaShopDo
	Joins(fields ...field.RelationField) ISlaShopDo
	Preload(fields ...field.RelationField) ISlaShopDo
	FirstOrInit() (*slaughter.SlaShop, error)
	FirstOrCreate() (*slaughter.SlaShop, error)
	FindByPage(offset int, limit int) (result []*slaughter.SlaShop, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlaShopDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slaShopDo) Debug() ISlaShopDo {
	return s.withDO(s.DO.Debug())
}

func (s slaShopDo) WithContext(ctx context.Context) ISlaShopDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slaShopDo) ReadDB() ISlaShopDo {
	return s.Clauses(dbresolver.Read)
}

func (s slaShopDo) WriteDB() ISlaShopDo {
	return s.Clauses(dbresolver.Write)
}

func (s slaShopDo) Session(config *gorm.Session) ISlaShopDo {
	return s.withDO(s.DO.Session(config))
}

func (s slaShopDo) Clauses(conds ...clause.Expression) ISlaShopDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slaShopDo) Returning(value interface{}, columns ...string) ISlaShopDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slaShopDo) Not(conds ...gen.Condition) ISlaShopDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slaShopDo) Or(conds ...gen.Condition) ISlaShopDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slaShopDo) Select(conds ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slaShopDo) Where(conds ...gen.Condition) ISlaShopDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slaShopDo) Order(conds ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slaShopDo) Distinct(cols ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slaShopDo) Omit(cols ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slaShopDo) Join(table schema.Tabler, on ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slaShopDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slaShopDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slaShopDo) Group(cols ...field.Expr) ISlaShopDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slaShopDo) Having(conds ...gen.Condition) ISlaShopDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slaShopDo) Limit(limit int) ISlaShopDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slaShopDo) Offset(offset int) ISlaShopDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slaShopDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlaShopDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slaShopDo) Unscoped() ISlaShopDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slaShopDo) Create(values ...*slaughter.SlaShop) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slaShopDo) CreateInBatches(values []*slaughter.SlaShop, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slaShopDo) Save(values ...*slaughter.SlaShop) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slaShopDo) First() (*slaughter.SlaShop, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaShop), nil
	}
}

func (s slaShopDo) Take() (*slaughter.SlaShop, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaShop), nil
	}
}

func (s slaShopDo) Last() (*slaughter.SlaShop, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaShop), nil
	}
}

func (s slaShopDo) Find() ([]*slaughter.SlaShop, error) {
	result, err := s.DO.Find()
	return result.([]*slaughter.SlaShop), err
}

func (s slaShopDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.SlaShop, err error) {
	buf := make([]*slaughter.SlaShop, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slaShopDo) FindInBatches(result *[]*slaughter.SlaShop, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slaShopDo) Attrs(attrs ...field.AssignExpr) ISlaShopDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slaShopDo) Assign(attrs ...field.AssignExpr) ISlaShopDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slaShopDo) Joins(fields ...field.RelationField) ISlaShopDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slaShopDo) Preload(fields ...field.RelationField) ISlaShopDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slaShopDo) FirstOrInit() (*slaughter.SlaShop, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaShop), nil
	}
}

func (s slaShopDo) FirstOrCreate() (*slaughter.SlaShop, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.SlaShop), nil
	}
}

func (s slaShopDo) FindByPage(offset int, limit int) (result []*slaughter.SlaShop, count int64, err error) {
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

func (s slaShopDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slaShopDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slaShopDo) Delete(models ...*slaughter.SlaShop) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slaShopDo) withDO(do gen.Dao) *slaShopDo {
	s.DO = *do.(*gen.DO)
	return s
}
