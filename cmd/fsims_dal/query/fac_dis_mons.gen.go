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

func newFacDisMon(db *gorm.DB, opts ...gen.DOOption) facDisMon {
	_facDisMon := facDisMon{}

	_facDisMon.facDisMonDo.UseDB(db, opts...)
	_facDisMon.facDisMonDo.UseModel(&slaughter.FacDisMon{})

	tableName := _facDisMon.facDisMonDo.TableName()
	_facDisMon.ALL = field.NewAsterisk(tableName)
	_facDisMon.ID = field.NewUint(tableName, "id")
	_facDisMon.CreatedAt = field.NewTime(tableName, "created_at")
	_facDisMon.UpdatedAt = field.NewTime(tableName, "updated_at")
	_facDisMon.DeletedAt = field.NewField(tableName, "deleted_at")
	_facDisMon.SlaInfoMonID = field.NewUint(tableName, "sla_info_mon_id")
	_facDisMon.SlaShop = facDisMonHasOneSlaShop{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SlaShop", "slaughter.SlaShop"),
		SlaShop8: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("SlaShop.SlaShop8", "slaughter.DisRecord"),
		},
	}

	_facDisMon.DivShop = facDisMonHasOneDivShop{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("DivShop", "slaughter.DivShop"),
		DivShop8: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("DivShop.DivShop8", "slaughter.DisRecord"),
		},
	}

	_facDisMon.AcidShop = facDisMonHasOneAcidShop{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AcidShop", "slaughter.AcidShop"),
		AcidShop8: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("AcidShop.AcidShop8", "slaughter.DisRecord"),
		},
	}

	_facDisMon.FroShop = facDisMonHasOneFroShop{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("FroShop", "slaughter.FroShop"),
		FroShop8: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("FroShop.FroShop8", "slaughter.DisRecord"),
		},
	}

	_facDisMon.PackShop = facDisMonHasOnePackShop{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PackShop", "slaughter.PackShop"),
		PackShop8: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PackShop.PackShop8", "slaughter.DisRecord"),
		},
	}

	_facDisMon.StaUni = facDisMonHasOneStaUni{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("StaUni", "slaughter.StaUni"),
		StaUni8: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("StaUni.StaUni8", "slaughter.DisRecord"),
		},
	}

	_facDisMon.fillFieldMap()

	return _facDisMon
}

type facDisMon struct {
	facDisMonDo facDisMonDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	SlaInfoMonID field.Uint
	SlaShop      facDisMonHasOneSlaShop

	DivShop facDisMonHasOneDivShop

	AcidShop facDisMonHasOneAcidShop

	FroShop facDisMonHasOneFroShop

	PackShop facDisMonHasOnePackShop

	StaUni facDisMonHasOneStaUni

	fieldMap map[string]field.Expr
}

func (f facDisMon) Table(newTableName string) *facDisMon {
	f.facDisMonDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f facDisMon) As(alias string) *facDisMon {
	f.facDisMonDo.DO = *(f.facDisMonDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *facDisMon) updateTableName(table string) *facDisMon {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewUint(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.SlaInfoMonID = field.NewUint(table, "sla_info_mon_id")

	f.fillFieldMap()

	return f
}

func (f *facDisMon) WithContext(ctx context.Context) IFacDisMonDo {
	return f.facDisMonDo.WithContext(ctx)
}

func (f facDisMon) TableName() string { return f.facDisMonDo.TableName() }

func (f facDisMon) Alias() string { return f.facDisMonDo.Alias() }

func (f facDisMon) Columns(cols ...field.Expr) gen.Columns { return f.facDisMonDo.Columns(cols...) }

func (f *facDisMon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *facDisMon) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 11)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["sla_info_mon_id"] = f.SlaInfoMonID

}

func (f facDisMon) clone(db *gorm.DB) facDisMon {
	f.facDisMonDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f facDisMon) replaceDB(db *gorm.DB) facDisMon {
	f.facDisMonDo.ReplaceDB(db)
	return f
}

type facDisMonHasOneSlaShop struct {
	db *gorm.DB

	field.RelationField

	SlaShop8 struct {
		field.RelationField
	}
}

func (a facDisMonHasOneSlaShop) Where(conds ...field.Expr) *facDisMonHasOneSlaShop {
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

func (a facDisMonHasOneSlaShop) WithContext(ctx context.Context) *facDisMonHasOneSlaShop {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a facDisMonHasOneSlaShop) Session(session *gorm.Session) *facDisMonHasOneSlaShop {
	a.db = a.db.Session(session)
	return &a
}

func (a facDisMonHasOneSlaShop) Model(m *slaughter.FacDisMon) *facDisMonHasOneSlaShopTx {
	return &facDisMonHasOneSlaShopTx{a.db.Model(m).Association(a.Name())}
}

type facDisMonHasOneSlaShopTx struct{ tx *gorm.Association }

func (a facDisMonHasOneSlaShopTx) Find() (result *slaughter.SlaShop, err error) {
	return result, a.tx.Find(&result)
}

func (a facDisMonHasOneSlaShopTx) Append(values ...*slaughter.SlaShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a facDisMonHasOneSlaShopTx) Replace(values ...*slaughter.SlaShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a facDisMonHasOneSlaShopTx) Delete(values ...*slaughter.SlaShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a facDisMonHasOneSlaShopTx) Clear() error {
	return a.tx.Clear()
}

func (a facDisMonHasOneSlaShopTx) Count() int64 {
	return a.tx.Count()
}

type facDisMonHasOneDivShop struct {
	db *gorm.DB

	field.RelationField

	DivShop8 struct {
		field.RelationField
	}
}

func (a facDisMonHasOneDivShop) Where(conds ...field.Expr) *facDisMonHasOneDivShop {
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

func (a facDisMonHasOneDivShop) WithContext(ctx context.Context) *facDisMonHasOneDivShop {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a facDisMonHasOneDivShop) Session(session *gorm.Session) *facDisMonHasOneDivShop {
	a.db = a.db.Session(session)
	return &a
}

func (a facDisMonHasOneDivShop) Model(m *slaughter.FacDisMon) *facDisMonHasOneDivShopTx {
	return &facDisMonHasOneDivShopTx{a.db.Model(m).Association(a.Name())}
}

type facDisMonHasOneDivShopTx struct{ tx *gorm.Association }

func (a facDisMonHasOneDivShopTx) Find() (result *slaughter.DivShop, err error) {
	return result, a.tx.Find(&result)
}

func (a facDisMonHasOneDivShopTx) Append(values ...*slaughter.DivShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a facDisMonHasOneDivShopTx) Replace(values ...*slaughter.DivShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a facDisMonHasOneDivShopTx) Delete(values ...*slaughter.DivShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a facDisMonHasOneDivShopTx) Clear() error {
	return a.tx.Clear()
}

func (a facDisMonHasOneDivShopTx) Count() int64 {
	return a.tx.Count()
}

type facDisMonHasOneAcidShop struct {
	db *gorm.DB

	field.RelationField

	AcidShop8 struct {
		field.RelationField
	}
}

func (a facDisMonHasOneAcidShop) Where(conds ...field.Expr) *facDisMonHasOneAcidShop {
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

func (a facDisMonHasOneAcidShop) WithContext(ctx context.Context) *facDisMonHasOneAcidShop {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a facDisMonHasOneAcidShop) Session(session *gorm.Session) *facDisMonHasOneAcidShop {
	a.db = a.db.Session(session)
	return &a
}

func (a facDisMonHasOneAcidShop) Model(m *slaughter.FacDisMon) *facDisMonHasOneAcidShopTx {
	return &facDisMonHasOneAcidShopTx{a.db.Model(m).Association(a.Name())}
}

type facDisMonHasOneAcidShopTx struct{ tx *gorm.Association }

func (a facDisMonHasOneAcidShopTx) Find() (result *slaughter.AcidShop, err error) {
	return result, a.tx.Find(&result)
}

func (a facDisMonHasOneAcidShopTx) Append(values ...*slaughter.AcidShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a facDisMonHasOneAcidShopTx) Replace(values ...*slaughter.AcidShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a facDisMonHasOneAcidShopTx) Delete(values ...*slaughter.AcidShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a facDisMonHasOneAcidShopTx) Clear() error {
	return a.tx.Clear()
}

func (a facDisMonHasOneAcidShopTx) Count() int64 {
	return a.tx.Count()
}

type facDisMonHasOneFroShop struct {
	db *gorm.DB

	field.RelationField

	FroShop8 struct {
		field.RelationField
	}
}

func (a facDisMonHasOneFroShop) Where(conds ...field.Expr) *facDisMonHasOneFroShop {
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

func (a facDisMonHasOneFroShop) WithContext(ctx context.Context) *facDisMonHasOneFroShop {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a facDisMonHasOneFroShop) Session(session *gorm.Session) *facDisMonHasOneFroShop {
	a.db = a.db.Session(session)
	return &a
}

func (a facDisMonHasOneFroShop) Model(m *slaughter.FacDisMon) *facDisMonHasOneFroShopTx {
	return &facDisMonHasOneFroShopTx{a.db.Model(m).Association(a.Name())}
}

type facDisMonHasOneFroShopTx struct{ tx *gorm.Association }

func (a facDisMonHasOneFroShopTx) Find() (result *slaughter.FroShop, err error) {
	return result, a.tx.Find(&result)
}

func (a facDisMonHasOneFroShopTx) Append(values ...*slaughter.FroShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a facDisMonHasOneFroShopTx) Replace(values ...*slaughter.FroShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a facDisMonHasOneFroShopTx) Delete(values ...*slaughter.FroShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a facDisMonHasOneFroShopTx) Clear() error {
	return a.tx.Clear()
}

func (a facDisMonHasOneFroShopTx) Count() int64 {
	return a.tx.Count()
}

type facDisMonHasOnePackShop struct {
	db *gorm.DB

	field.RelationField

	PackShop8 struct {
		field.RelationField
	}
}

func (a facDisMonHasOnePackShop) Where(conds ...field.Expr) *facDisMonHasOnePackShop {
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

func (a facDisMonHasOnePackShop) WithContext(ctx context.Context) *facDisMonHasOnePackShop {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a facDisMonHasOnePackShop) Session(session *gorm.Session) *facDisMonHasOnePackShop {
	a.db = a.db.Session(session)
	return &a
}

func (a facDisMonHasOnePackShop) Model(m *slaughter.FacDisMon) *facDisMonHasOnePackShopTx {
	return &facDisMonHasOnePackShopTx{a.db.Model(m).Association(a.Name())}
}

type facDisMonHasOnePackShopTx struct{ tx *gorm.Association }

func (a facDisMonHasOnePackShopTx) Find() (result *slaughter.PackShop, err error) {
	return result, a.tx.Find(&result)
}

func (a facDisMonHasOnePackShopTx) Append(values ...*slaughter.PackShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a facDisMonHasOnePackShopTx) Replace(values ...*slaughter.PackShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a facDisMonHasOnePackShopTx) Delete(values ...*slaughter.PackShop) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a facDisMonHasOnePackShopTx) Clear() error {
	return a.tx.Clear()
}

func (a facDisMonHasOnePackShopTx) Count() int64 {
	return a.tx.Count()
}

type facDisMonHasOneStaUni struct {
	db *gorm.DB

	field.RelationField

	StaUni8 struct {
		field.RelationField
	}
}

func (a facDisMonHasOneStaUni) Where(conds ...field.Expr) *facDisMonHasOneStaUni {
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

func (a facDisMonHasOneStaUni) WithContext(ctx context.Context) *facDisMonHasOneStaUni {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a facDisMonHasOneStaUni) Session(session *gorm.Session) *facDisMonHasOneStaUni {
	a.db = a.db.Session(session)
	return &a
}

func (a facDisMonHasOneStaUni) Model(m *slaughter.FacDisMon) *facDisMonHasOneStaUniTx {
	return &facDisMonHasOneStaUniTx{a.db.Model(m).Association(a.Name())}
}

type facDisMonHasOneStaUniTx struct{ tx *gorm.Association }

func (a facDisMonHasOneStaUniTx) Find() (result *slaughter.StaUni, err error) {
	return result, a.tx.Find(&result)
}

func (a facDisMonHasOneStaUniTx) Append(values ...*slaughter.StaUni) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a facDisMonHasOneStaUniTx) Replace(values ...*slaughter.StaUni) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a facDisMonHasOneStaUniTx) Delete(values ...*slaughter.StaUni) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a facDisMonHasOneStaUniTx) Clear() error {
	return a.tx.Clear()
}

func (a facDisMonHasOneStaUniTx) Count() int64 {
	return a.tx.Count()
}

type facDisMonDo struct{ gen.DO }

type IFacDisMonDo interface {
	gen.SubQuery
	Debug() IFacDisMonDo
	WithContext(ctx context.Context) IFacDisMonDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFacDisMonDo
	WriteDB() IFacDisMonDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFacDisMonDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFacDisMonDo
	Not(conds ...gen.Condition) IFacDisMonDo
	Or(conds ...gen.Condition) IFacDisMonDo
	Select(conds ...field.Expr) IFacDisMonDo
	Where(conds ...gen.Condition) IFacDisMonDo
	Order(conds ...field.Expr) IFacDisMonDo
	Distinct(cols ...field.Expr) IFacDisMonDo
	Omit(cols ...field.Expr) IFacDisMonDo
	Join(table schema.Tabler, on ...field.Expr) IFacDisMonDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFacDisMonDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFacDisMonDo
	Group(cols ...field.Expr) IFacDisMonDo
	Having(conds ...gen.Condition) IFacDisMonDo
	Limit(limit int) IFacDisMonDo
	Offset(offset int) IFacDisMonDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFacDisMonDo
	Unscoped() IFacDisMonDo
	Create(values ...*slaughter.FacDisMon) error
	CreateInBatches(values []*slaughter.FacDisMon, batchSize int) error
	Save(values ...*slaughter.FacDisMon) error
	First() (*slaughter.FacDisMon, error)
	Take() (*slaughter.FacDisMon, error)
	Last() (*slaughter.FacDisMon, error)
	Find() ([]*slaughter.FacDisMon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.FacDisMon, err error)
	FindInBatches(result *[]*slaughter.FacDisMon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*slaughter.FacDisMon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFacDisMonDo
	Assign(attrs ...field.AssignExpr) IFacDisMonDo
	Joins(fields ...field.RelationField) IFacDisMonDo
	Preload(fields ...field.RelationField) IFacDisMonDo
	FirstOrInit() (*slaughter.FacDisMon, error)
	FirstOrCreate() (*slaughter.FacDisMon, error)
	FindByPage(offset int, limit int) (result []*slaughter.FacDisMon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFacDisMonDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f facDisMonDo) Debug() IFacDisMonDo {
	return f.withDO(f.DO.Debug())
}

func (f facDisMonDo) WithContext(ctx context.Context) IFacDisMonDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f facDisMonDo) ReadDB() IFacDisMonDo {
	return f.Clauses(dbresolver.Read)
}

func (f facDisMonDo) WriteDB() IFacDisMonDo {
	return f.Clauses(dbresolver.Write)
}

func (f facDisMonDo) Session(config *gorm.Session) IFacDisMonDo {
	return f.withDO(f.DO.Session(config))
}

func (f facDisMonDo) Clauses(conds ...clause.Expression) IFacDisMonDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f facDisMonDo) Returning(value interface{}, columns ...string) IFacDisMonDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f facDisMonDo) Not(conds ...gen.Condition) IFacDisMonDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f facDisMonDo) Or(conds ...gen.Condition) IFacDisMonDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f facDisMonDo) Select(conds ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f facDisMonDo) Where(conds ...gen.Condition) IFacDisMonDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f facDisMonDo) Order(conds ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f facDisMonDo) Distinct(cols ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f facDisMonDo) Omit(cols ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f facDisMonDo) Join(table schema.Tabler, on ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f facDisMonDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f facDisMonDo) RightJoin(table schema.Tabler, on ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f facDisMonDo) Group(cols ...field.Expr) IFacDisMonDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f facDisMonDo) Having(conds ...gen.Condition) IFacDisMonDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f facDisMonDo) Limit(limit int) IFacDisMonDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f facDisMonDo) Offset(offset int) IFacDisMonDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f facDisMonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFacDisMonDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f facDisMonDo) Unscoped() IFacDisMonDo {
	return f.withDO(f.DO.Unscoped())
}

func (f facDisMonDo) Create(values ...*slaughter.FacDisMon) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f facDisMonDo) CreateInBatches(values []*slaughter.FacDisMon, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f facDisMonDo) Save(values ...*slaughter.FacDisMon) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f facDisMonDo) First() (*slaughter.FacDisMon, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FacDisMon), nil
	}
}

func (f facDisMonDo) Take() (*slaughter.FacDisMon, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FacDisMon), nil
	}
}

func (f facDisMonDo) Last() (*slaughter.FacDisMon, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FacDisMon), nil
	}
}

func (f facDisMonDo) Find() ([]*slaughter.FacDisMon, error) {
	result, err := f.DO.Find()
	return result.([]*slaughter.FacDisMon), err
}

func (f facDisMonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*slaughter.FacDisMon, err error) {
	buf := make([]*slaughter.FacDisMon, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f facDisMonDo) FindInBatches(result *[]*slaughter.FacDisMon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f facDisMonDo) Attrs(attrs ...field.AssignExpr) IFacDisMonDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f facDisMonDo) Assign(attrs ...field.AssignExpr) IFacDisMonDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f facDisMonDo) Joins(fields ...field.RelationField) IFacDisMonDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f facDisMonDo) Preload(fields ...field.RelationField) IFacDisMonDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f facDisMonDo) FirstOrInit() (*slaughter.FacDisMon, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FacDisMon), nil
	}
}

func (f facDisMonDo) FirstOrCreate() (*slaughter.FacDisMon, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*slaughter.FacDisMon), nil
	}
}

func (f facDisMonDo) FindByPage(offset int, limit int) (result []*slaughter.FacDisMon, count int64, err error) {
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

func (f facDisMonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f facDisMonDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f facDisMonDo) Delete(models ...*slaughter.FacDisMon) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *facDisMonDo) withDO(do gen.Dao) *facDisMonDo {
	f.DO = *do.(*gen.DO)
	return f
}