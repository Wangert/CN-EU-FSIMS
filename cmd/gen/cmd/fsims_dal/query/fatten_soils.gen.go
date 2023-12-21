// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/fatten"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newFattenSoil(db *gorm.DB, opts ...gen.DOOption) fattenSoil {
	_fattenSoil := fattenSoil{}

	_fattenSoil.fattenSoilDo.UseDB(db, opts...)
	_fattenSoil.fattenSoilDo.UseModel(&fatten.FattenSoil{})

	tableName := _fattenSoil.fattenSoilDo.TableName()
	_fattenSoil.ALL = field.NewAsterisk(tableName)
	_fattenSoil.ID = field.NewUint(tableName, "id")
	_fattenSoil.CreatedAt = field.NewTime(tableName, "created_at")
	_fattenSoil.UpdatedAt = field.NewTime(tableName, "updated_at")
	_fattenSoil.DeletedAt = field.NewField(tableName, "deleted_at")
	_fattenSoil.FatPID = field.NewString(tableName, "fat_p_id")
	_fattenSoil.PhysicalHazard = fattenSoilHasOnePhysicalHazard{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PhysicalHazard", "fatten.FattenSoilPhysicalHazard"),
	}

	_fattenSoil.Biohazard = fattenSoilHasOneBiohazard{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Biohazard", "fatten.FattenSoilBiohazard"),
	}

	_fattenSoil.fillFieldMap()

	return _fattenSoil
}

type fattenSoil struct {
	fattenSoilDo fattenSoilDo

	ALL            field.Asterisk
	ID             field.Uint
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	FatPID         field.String
	PhysicalHazard fattenSoilHasOnePhysicalHazard

	Biohazard fattenSoilHasOneBiohazard

	fieldMap map[string]field.Expr
}

func (f fattenSoil) Table(newTableName string) *fattenSoil {
	f.fattenSoilDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fattenSoil) As(alias string) *fattenSoil {
	f.fattenSoilDo.DO = *(f.fattenSoilDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fattenSoil) updateTableName(table string) *fattenSoil {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewUint(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.FatPID = field.NewString(table, "fat_p_id")

	f.fillFieldMap()

	return f
}

func (f *fattenSoil) WithContext(ctx context.Context) IFattenSoilDo {
	return f.fattenSoilDo.WithContext(ctx)
}

func (f fattenSoil) TableName() string { return f.fattenSoilDo.TableName() }

func (f fattenSoil) Alias() string { return f.fattenSoilDo.Alias() }

func (f fattenSoil) Columns(cols ...field.Expr) gen.Columns { return f.fattenSoilDo.Columns(cols...) }

func (f *fattenSoil) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fattenSoil) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["fat_p_id"] = f.FatPID

}

func (f fattenSoil) clone(db *gorm.DB) fattenSoil {
	f.fattenSoilDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fattenSoil) replaceDB(db *gorm.DB) fattenSoil {
	f.fattenSoilDo.ReplaceDB(db)
	return f
}

type fattenSoilHasOnePhysicalHazard struct {
	db *gorm.DB

	field.RelationField
}

func (a fattenSoilHasOnePhysicalHazard) Where(conds ...field.Expr) *fattenSoilHasOnePhysicalHazard {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a fattenSoilHasOnePhysicalHazard) WithContext(ctx context.Context) *fattenSoilHasOnePhysicalHazard {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a fattenSoilHasOnePhysicalHazard) Session(session *gorm.Session) *fattenSoilHasOnePhysicalHazard {
	a.db = a.db.Session(session)
	return &a
}

func (a fattenSoilHasOnePhysicalHazard) Model(m *fatten.FattenSoil) *fattenSoilHasOnePhysicalHazardTx {
	return &fattenSoilHasOnePhysicalHazardTx{a.db.Model(m).Association(a.Name())}
}

type fattenSoilHasOnePhysicalHazardTx struct{ tx *gorm.Association }

func (a fattenSoilHasOnePhysicalHazardTx) Find() (result *fatten.FattenSoilPhysicalHazard, err error) {
	return result, a.tx.Find(&result)
}

func (a fattenSoilHasOnePhysicalHazardTx) Append(values ...*fatten.FattenSoilPhysicalHazard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a fattenSoilHasOnePhysicalHazardTx) Replace(values ...*fatten.FattenSoilPhysicalHazard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a fattenSoilHasOnePhysicalHazardTx) Delete(values ...*fatten.FattenSoilPhysicalHazard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a fattenSoilHasOnePhysicalHazardTx) Clear() error {
	return a.tx.Clear()
}

func (a fattenSoilHasOnePhysicalHazardTx) Count() int64 {
	return a.tx.Count()
}

type fattenSoilHasOneBiohazard struct {
	db *gorm.DB

	field.RelationField
}

func (a fattenSoilHasOneBiohazard) Where(conds ...field.Expr) *fattenSoilHasOneBiohazard {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a fattenSoilHasOneBiohazard) WithContext(ctx context.Context) *fattenSoilHasOneBiohazard {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a fattenSoilHasOneBiohazard) Session(session *gorm.Session) *fattenSoilHasOneBiohazard {
	a.db = a.db.Session(session)
	return &a
}

func (a fattenSoilHasOneBiohazard) Model(m *fatten.FattenSoil) *fattenSoilHasOneBiohazardTx {
	return &fattenSoilHasOneBiohazardTx{a.db.Model(m).Association(a.Name())}
}

type fattenSoilHasOneBiohazardTx struct{ tx *gorm.Association }

func (a fattenSoilHasOneBiohazardTx) Find() (result *fatten.FattenSoilBiohazard, err error) {
	return result, a.tx.Find(&result)
}

func (a fattenSoilHasOneBiohazardTx) Append(values ...*fatten.FattenSoilBiohazard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a fattenSoilHasOneBiohazardTx) Replace(values ...*fatten.FattenSoilBiohazard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a fattenSoilHasOneBiohazardTx) Delete(values ...*fatten.FattenSoilBiohazard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a fattenSoilHasOneBiohazardTx) Clear() error {
	return a.tx.Clear()
}

func (a fattenSoilHasOneBiohazardTx) Count() int64 {
	return a.tx.Count()
}

type fattenSoilDo struct{ gen.DO }

type IFattenSoilDo interface {
	gen.SubQuery
	Debug() IFattenSoilDo
	WithContext(ctx context.Context) IFattenSoilDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFattenSoilDo
	WriteDB() IFattenSoilDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFattenSoilDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFattenSoilDo
	Not(conds ...gen.Condition) IFattenSoilDo
	Or(conds ...gen.Condition) IFattenSoilDo
	Select(conds ...field.Expr) IFattenSoilDo
	Where(conds ...gen.Condition) IFattenSoilDo
	Order(conds ...field.Expr) IFattenSoilDo
	Distinct(cols ...field.Expr) IFattenSoilDo
	Omit(cols ...field.Expr) IFattenSoilDo
	Join(table schema.Tabler, on ...field.Expr) IFattenSoilDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFattenSoilDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFattenSoilDo
	Group(cols ...field.Expr) IFattenSoilDo
	Having(conds ...gen.Condition) IFattenSoilDo
	Limit(limit int) IFattenSoilDo
	Offset(offset int) IFattenSoilDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFattenSoilDo
	Unscoped() IFattenSoilDo
	Create(values ...*fatten.FattenSoil) error
	CreateInBatches(values []*fatten.FattenSoil, batchSize int) error
	Save(values ...*fatten.FattenSoil) error
	First() (*fatten.FattenSoil, error)
	Take() (*fatten.FattenSoil, error)
	Last() (*fatten.FattenSoil, error)
	Find() ([]*fatten.FattenSoil, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*fatten.FattenSoil, err error)
	FindInBatches(result *[]*fatten.FattenSoil, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*fatten.FattenSoil) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFattenSoilDo
	Assign(attrs ...field.AssignExpr) IFattenSoilDo
	Joins(fields ...field.RelationField) IFattenSoilDo
	Preload(fields ...field.RelationField) IFattenSoilDo
	FirstOrInit() (*fatten.FattenSoil, error)
	FirstOrCreate() (*fatten.FattenSoil, error)
	FindByPage(offset int, limit int) (result []*fatten.FattenSoil, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFattenSoilDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fattenSoilDo) Debug() IFattenSoilDo {
	return f.withDO(f.DO.Debug())
}

func (f fattenSoilDo) WithContext(ctx context.Context) IFattenSoilDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fattenSoilDo) ReadDB() IFattenSoilDo {
	return f.Clauses(dbresolver.Read)
}

func (f fattenSoilDo) WriteDB() IFattenSoilDo {
	return f.Clauses(dbresolver.Write)
}

func (f fattenSoilDo) Session(config *gorm.Session) IFattenSoilDo {
	return f.withDO(f.DO.Session(config))
}

func (f fattenSoilDo) Clauses(conds ...clause.Expression) IFattenSoilDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fattenSoilDo) Returning(value interface{}, columns ...string) IFattenSoilDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fattenSoilDo) Not(conds ...gen.Condition) IFattenSoilDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fattenSoilDo) Or(conds ...gen.Condition) IFattenSoilDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fattenSoilDo) Select(conds ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fattenSoilDo) Where(conds ...gen.Condition) IFattenSoilDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fattenSoilDo) Order(conds ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fattenSoilDo) Distinct(cols ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fattenSoilDo) Omit(cols ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fattenSoilDo) Join(table schema.Tabler, on ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fattenSoilDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fattenSoilDo) RightJoin(table schema.Tabler, on ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fattenSoilDo) Group(cols ...field.Expr) IFattenSoilDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fattenSoilDo) Having(conds ...gen.Condition) IFattenSoilDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fattenSoilDo) Limit(limit int) IFattenSoilDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fattenSoilDo) Offset(offset int) IFattenSoilDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fattenSoilDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFattenSoilDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fattenSoilDo) Unscoped() IFattenSoilDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fattenSoilDo) Create(values ...*fatten.FattenSoil) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fattenSoilDo) CreateInBatches(values []*fatten.FattenSoil, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fattenSoilDo) Save(values ...*fatten.FattenSoil) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fattenSoilDo) First() (*fatten.FattenSoil, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*fatten.FattenSoil), nil
	}
}

func (f fattenSoilDo) Take() (*fatten.FattenSoil, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*fatten.FattenSoil), nil
	}
}

func (f fattenSoilDo) Last() (*fatten.FattenSoil, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*fatten.FattenSoil), nil
	}
}

func (f fattenSoilDo) Find() ([]*fatten.FattenSoil, error) {
	result, err := f.DO.Find()
	return result.([]*fatten.FattenSoil), err
}

func (f fattenSoilDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*fatten.FattenSoil, err error) {
	buf := make([]*fatten.FattenSoil, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fattenSoilDo) FindInBatches(result *[]*fatten.FattenSoil, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fattenSoilDo) Attrs(attrs ...field.AssignExpr) IFattenSoilDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fattenSoilDo) Assign(attrs ...field.AssignExpr) IFattenSoilDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fattenSoilDo) Joins(fields ...field.RelationField) IFattenSoilDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fattenSoilDo) Preload(fields ...field.RelationField) IFattenSoilDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fattenSoilDo) FirstOrInit() (*fatten.FattenSoil, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*fatten.FattenSoil), nil
	}
}

func (f fattenSoilDo) FirstOrCreate() (*fatten.FattenSoil, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*fatten.FattenSoil), nil
	}
}

func (f fattenSoilDo) FindByPage(offset int, limit int) (result []*fatten.FattenSoil, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fattenSoilDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fattenSoilDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fattenSoilDo) Delete(models ...*fatten.FattenSoil) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fattenSoilDo) withDO(do gen.Dao) *fattenSoilDo {
	f.DO = *do.(*gen.DO)
	return f
}