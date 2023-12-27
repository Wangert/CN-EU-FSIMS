// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/premortem"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPreSlaInfoRec(db *gorm.DB, opts ...gen.DOOption) preSlaInfoRec {
	_preSlaInfoRec := preSlaInfoRec{}

	_preSlaInfoRec.preSlaInfoRecDo.UseDB(db, opts...)
	_preSlaInfoRec.preSlaInfoRecDo.UseModel(&premortem.PreSlaInfoRec{})

	tableName := _preSlaInfoRec.preSlaInfoRecDo.TableName()
	_preSlaInfoRec.ALL = field.NewAsterisk(tableName)
	_preSlaInfoRec.ID = field.NewUint(tableName, "id")
	_preSlaInfoRec.CreatedAt = field.NewTime(tableName, "created_at")
	_preSlaInfoRec.UpdatedAt = field.NewTime(tableName, "updated_at")
	_preSlaInfoRec.DeletedAt = field.NewField(tableName, "deleted_at")
	_preSlaInfoRec.TimeRecordAt = field.NewString(tableName, "time_record_at")
	_preSlaInfoRec.PreSlaInfoRec3 = field.NewBool(tableName, "pre_sla_info_rec3")
	_preSlaInfoRec.PreSlaInfoRec4 = field.NewString(tableName, "pre_sla_info_rec4")
	_preSlaInfoRec.PreSlaInfoRec5 = field.NewFloat32(tableName, "pre_sla_info_rec5")
	_preSlaInfoRec.PreSlaInfoRec7 = field.NewString(tableName, "pre_sla_info_rec7")
	_preSlaInfoRec.PreSlaInfoRec1 = preSlaInfoRecHasOnePreSlaInfoRec1{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PreSlaInfoRec1", "premortem.Gps"),
	}

	_preSlaInfoRec.PreSlaInfoRec2 = preSlaInfoRecHasOnePreSlaInfoRec2{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PreSlaInfoRec2", "premortem.PreSlaDietManage"),
	}

	_preSlaInfoRec.PreSlaInfoRec6 = preSlaInfoRecHasOnePreSlaInfoRec6{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PreSlaInfoRec6", "premortem.PreSlaGerms"),
	}

	_preSlaInfoRec.PreSlaInfoRec8 = preSlaInfoRecHasOnePreSlaInfoRec8{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PreSlaInfoRec8", "premortem.PreSlaPicAndEn"),
	}

	_preSlaInfoRec.fillFieldMap()

	return _preSlaInfoRec
}

type preSlaInfoRec struct {
	preSlaInfoRecDo preSlaInfoRecDo

	ALL            field.Asterisk
	ID             field.Uint
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	TimeRecordAt   field.String
	PreSlaInfoRec3 field.Bool
	PreSlaInfoRec4 field.String
	PreSlaInfoRec5 field.Float32
	PreSlaInfoRec7 field.String
	PreSlaInfoRec1 preSlaInfoRecHasOnePreSlaInfoRec1

	PreSlaInfoRec2 preSlaInfoRecHasOnePreSlaInfoRec2

	PreSlaInfoRec6 preSlaInfoRecHasOnePreSlaInfoRec6

	PreSlaInfoRec8 preSlaInfoRecHasOnePreSlaInfoRec8

	fieldMap map[string]field.Expr
}

func (p preSlaInfoRec) Table(newTableName string) *preSlaInfoRec {
	p.preSlaInfoRecDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p preSlaInfoRec) As(alias string) *preSlaInfoRec {
	p.preSlaInfoRecDo.DO = *(p.preSlaInfoRecDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *preSlaInfoRec) updateTableName(table string) *preSlaInfoRec {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.TimeRecordAt = field.NewString(table, "time_record_at")
	p.PreSlaInfoRec3 = field.NewBool(table, "pre_sla_info_rec3")
	p.PreSlaInfoRec4 = field.NewString(table, "pre_sla_info_rec4")
	p.PreSlaInfoRec5 = field.NewFloat32(table, "pre_sla_info_rec5")
	p.PreSlaInfoRec7 = field.NewString(table, "pre_sla_info_rec7")

	p.fillFieldMap()

	return p
}

func (p *preSlaInfoRec) WithContext(ctx context.Context) IPreSlaInfoRecDo {
	return p.preSlaInfoRecDo.WithContext(ctx)
}

func (p preSlaInfoRec) TableName() string { return p.preSlaInfoRecDo.TableName() }

func (p preSlaInfoRec) Alias() string { return p.preSlaInfoRecDo.Alias() }

func (p preSlaInfoRec) Columns(cols ...field.Expr) gen.Columns {
	return p.preSlaInfoRecDo.Columns(cols...)
}

func (p *preSlaInfoRec) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *preSlaInfoRec) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["time_record_at"] = p.TimeRecordAt
	p.fieldMap["pre_sla_info_rec3"] = p.PreSlaInfoRec3
	p.fieldMap["pre_sla_info_rec4"] = p.PreSlaInfoRec4
	p.fieldMap["pre_sla_info_rec5"] = p.PreSlaInfoRec5
	p.fieldMap["pre_sla_info_rec7"] = p.PreSlaInfoRec7

}

func (p preSlaInfoRec) clone(db *gorm.DB) preSlaInfoRec {
	p.preSlaInfoRecDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p preSlaInfoRec) replaceDB(db *gorm.DB) preSlaInfoRec {
	p.preSlaInfoRecDo.ReplaceDB(db)
	return p
}

type preSlaInfoRecHasOnePreSlaInfoRec1 struct {
	db *gorm.DB

	field.RelationField
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1) Where(conds ...field.Expr) *preSlaInfoRecHasOnePreSlaInfoRec1 {
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

func (a preSlaInfoRecHasOnePreSlaInfoRec1) WithContext(ctx context.Context) *preSlaInfoRecHasOnePreSlaInfoRec1 {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1) Session(session *gorm.Session) *preSlaInfoRecHasOnePreSlaInfoRec1 {
	a.db = a.db.Session(session)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1) Model(m *premortem.PreSlaInfoRec) *preSlaInfoRecHasOnePreSlaInfoRec1Tx {
	return &preSlaInfoRecHasOnePreSlaInfoRec1Tx{a.db.Model(m).Association(a.Name())}
}

type preSlaInfoRecHasOnePreSlaInfoRec1Tx struct{ tx *gorm.Association }

func (a preSlaInfoRecHasOnePreSlaInfoRec1Tx) Find() (result *premortem.Gps, err error) {
	return result, a.tx.Find(&result)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1Tx) Append(values ...*premortem.Gps) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1Tx) Replace(values ...*premortem.Gps) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1Tx) Delete(values ...*premortem.Gps) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1Tx) Clear() error {
	return a.tx.Clear()
}

func (a preSlaInfoRecHasOnePreSlaInfoRec1Tx) Count() int64 {
	return a.tx.Count()
}

type preSlaInfoRecHasOnePreSlaInfoRec2 struct {
	db *gorm.DB

	field.RelationField
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2) Where(conds ...field.Expr) *preSlaInfoRecHasOnePreSlaInfoRec2 {
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

func (a preSlaInfoRecHasOnePreSlaInfoRec2) WithContext(ctx context.Context) *preSlaInfoRecHasOnePreSlaInfoRec2 {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2) Session(session *gorm.Session) *preSlaInfoRecHasOnePreSlaInfoRec2 {
	a.db = a.db.Session(session)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2) Model(m *premortem.PreSlaInfoRec) *preSlaInfoRecHasOnePreSlaInfoRec2Tx {
	return &preSlaInfoRecHasOnePreSlaInfoRec2Tx{a.db.Model(m).Association(a.Name())}
}

type preSlaInfoRecHasOnePreSlaInfoRec2Tx struct{ tx *gorm.Association }

func (a preSlaInfoRecHasOnePreSlaInfoRec2Tx) Find() (result *premortem.PreSlaDietManage, err error) {
	return result, a.tx.Find(&result)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2Tx) Append(values ...*premortem.PreSlaDietManage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2Tx) Replace(values ...*premortem.PreSlaDietManage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2Tx) Delete(values ...*premortem.PreSlaDietManage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2Tx) Clear() error {
	return a.tx.Clear()
}

func (a preSlaInfoRecHasOnePreSlaInfoRec2Tx) Count() int64 {
	return a.tx.Count()
}

type preSlaInfoRecHasOnePreSlaInfoRec6 struct {
	db *gorm.DB

	field.RelationField
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6) Where(conds ...field.Expr) *preSlaInfoRecHasOnePreSlaInfoRec6 {
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

func (a preSlaInfoRecHasOnePreSlaInfoRec6) WithContext(ctx context.Context) *preSlaInfoRecHasOnePreSlaInfoRec6 {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6) Session(session *gorm.Session) *preSlaInfoRecHasOnePreSlaInfoRec6 {
	a.db = a.db.Session(session)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6) Model(m *premortem.PreSlaInfoRec) *preSlaInfoRecHasOnePreSlaInfoRec6Tx {
	return &preSlaInfoRecHasOnePreSlaInfoRec6Tx{a.db.Model(m).Association(a.Name())}
}

type preSlaInfoRecHasOnePreSlaInfoRec6Tx struct{ tx *gorm.Association }

func (a preSlaInfoRecHasOnePreSlaInfoRec6Tx) Find() (result *premortem.PreSlaGerms, err error) {
	return result, a.tx.Find(&result)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6Tx) Append(values ...*premortem.PreSlaGerms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6Tx) Replace(values ...*premortem.PreSlaGerms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6Tx) Delete(values ...*premortem.PreSlaGerms) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6Tx) Clear() error {
	return a.tx.Clear()
}

func (a preSlaInfoRecHasOnePreSlaInfoRec6Tx) Count() int64 {
	return a.tx.Count()
}

type preSlaInfoRecHasOnePreSlaInfoRec8 struct {
	db *gorm.DB

	field.RelationField
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8) Where(conds ...field.Expr) *preSlaInfoRecHasOnePreSlaInfoRec8 {
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

func (a preSlaInfoRecHasOnePreSlaInfoRec8) WithContext(ctx context.Context) *preSlaInfoRecHasOnePreSlaInfoRec8 {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8) Session(session *gorm.Session) *preSlaInfoRecHasOnePreSlaInfoRec8 {
	a.db = a.db.Session(session)
	return &a
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8) Model(m *premortem.PreSlaInfoRec) *preSlaInfoRecHasOnePreSlaInfoRec8Tx {
	return &preSlaInfoRecHasOnePreSlaInfoRec8Tx{a.db.Model(m).Association(a.Name())}
}

type preSlaInfoRecHasOnePreSlaInfoRec8Tx struct{ tx *gorm.Association }

func (a preSlaInfoRecHasOnePreSlaInfoRec8Tx) Find() (result *premortem.PreSlaPicAndEn, err error) {
	return result, a.tx.Find(&result)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8Tx) Append(values ...*premortem.PreSlaPicAndEn) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8Tx) Replace(values ...*premortem.PreSlaPicAndEn) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8Tx) Delete(values ...*premortem.PreSlaPicAndEn) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8Tx) Clear() error {
	return a.tx.Clear()
}

func (a preSlaInfoRecHasOnePreSlaInfoRec8Tx) Count() int64 {
	return a.tx.Count()
}

type preSlaInfoRecDo struct{ gen.DO }

type IPreSlaInfoRecDo interface {
	gen.SubQuery
	Debug() IPreSlaInfoRecDo
	WithContext(ctx context.Context) IPreSlaInfoRecDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPreSlaInfoRecDo
	WriteDB() IPreSlaInfoRecDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPreSlaInfoRecDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPreSlaInfoRecDo
	Not(conds ...gen.Condition) IPreSlaInfoRecDo
	Or(conds ...gen.Condition) IPreSlaInfoRecDo
	Select(conds ...field.Expr) IPreSlaInfoRecDo
	Where(conds ...gen.Condition) IPreSlaInfoRecDo
	Order(conds ...field.Expr) IPreSlaInfoRecDo
	Distinct(cols ...field.Expr) IPreSlaInfoRecDo
	Omit(cols ...field.Expr) IPreSlaInfoRecDo
	Join(table schema.Tabler, on ...field.Expr) IPreSlaInfoRecDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPreSlaInfoRecDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPreSlaInfoRecDo
	Group(cols ...field.Expr) IPreSlaInfoRecDo
	Having(conds ...gen.Condition) IPreSlaInfoRecDo
	Limit(limit int) IPreSlaInfoRecDo
	Offset(offset int) IPreSlaInfoRecDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPreSlaInfoRecDo
	Unscoped() IPreSlaInfoRecDo
	Create(values ...*premortem.PreSlaInfoRec) error
	CreateInBatches(values []*premortem.PreSlaInfoRec, batchSize int) error
	Save(values ...*premortem.PreSlaInfoRec) error
	First() (*premortem.PreSlaInfoRec, error)
	Take() (*premortem.PreSlaInfoRec, error)
	Last() (*premortem.PreSlaInfoRec, error)
	Find() ([]*premortem.PreSlaInfoRec, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*premortem.PreSlaInfoRec, err error)
	FindInBatches(result *[]*premortem.PreSlaInfoRec, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*premortem.PreSlaInfoRec) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPreSlaInfoRecDo
	Assign(attrs ...field.AssignExpr) IPreSlaInfoRecDo
	Joins(fields ...field.RelationField) IPreSlaInfoRecDo
	Preload(fields ...field.RelationField) IPreSlaInfoRecDo
	FirstOrInit() (*premortem.PreSlaInfoRec, error)
	FirstOrCreate() (*premortem.PreSlaInfoRec, error)
	FindByPage(offset int, limit int) (result []*premortem.PreSlaInfoRec, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPreSlaInfoRecDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p preSlaInfoRecDo) Debug() IPreSlaInfoRecDo {
	return p.withDO(p.DO.Debug())
}

func (p preSlaInfoRecDo) WithContext(ctx context.Context) IPreSlaInfoRecDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p preSlaInfoRecDo) ReadDB() IPreSlaInfoRecDo {
	return p.Clauses(dbresolver.Read)
}

func (p preSlaInfoRecDo) WriteDB() IPreSlaInfoRecDo {
	return p.Clauses(dbresolver.Write)
}

func (p preSlaInfoRecDo) Session(config *gorm.Session) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Session(config))
}

func (p preSlaInfoRecDo) Clauses(conds ...clause.Expression) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p preSlaInfoRecDo) Returning(value interface{}, columns ...string) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p preSlaInfoRecDo) Not(conds ...gen.Condition) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p preSlaInfoRecDo) Or(conds ...gen.Condition) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p preSlaInfoRecDo) Select(conds ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p preSlaInfoRecDo) Where(conds ...gen.Condition) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p preSlaInfoRecDo) Order(conds ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p preSlaInfoRecDo) Distinct(cols ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p preSlaInfoRecDo) Omit(cols ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p preSlaInfoRecDo) Join(table schema.Tabler, on ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p preSlaInfoRecDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p preSlaInfoRecDo) RightJoin(table schema.Tabler, on ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p preSlaInfoRecDo) Group(cols ...field.Expr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p preSlaInfoRecDo) Having(conds ...gen.Condition) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p preSlaInfoRecDo) Limit(limit int) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p preSlaInfoRecDo) Offset(offset int) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p preSlaInfoRecDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p preSlaInfoRecDo) Unscoped() IPreSlaInfoRecDo {
	return p.withDO(p.DO.Unscoped())
}

func (p preSlaInfoRecDo) Create(values ...*premortem.PreSlaInfoRec) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p preSlaInfoRecDo) CreateInBatches(values []*premortem.PreSlaInfoRec, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p preSlaInfoRecDo) Save(values ...*premortem.PreSlaInfoRec) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p preSlaInfoRecDo) First() (*premortem.PreSlaInfoRec, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.PreSlaInfoRec), nil
	}
}

func (p preSlaInfoRecDo) Take() (*premortem.PreSlaInfoRec, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.PreSlaInfoRec), nil
	}
}

func (p preSlaInfoRecDo) Last() (*premortem.PreSlaInfoRec, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.PreSlaInfoRec), nil
	}
}

func (p preSlaInfoRecDo) Find() ([]*premortem.PreSlaInfoRec, error) {
	result, err := p.DO.Find()
	return result.([]*premortem.PreSlaInfoRec), err
}

func (p preSlaInfoRecDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*premortem.PreSlaInfoRec, err error) {
	buf := make([]*premortem.PreSlaInfoRec, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p preSlaInfoRecDo) FindInBatches(result *[]*premortem.PreSlaInfoRec, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p preSlaInfoRecDo) Attrs(attrs ...field.AssignExpr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p preSlaInfoRecDo) Assign(attrs ...field.AssignExpr) IPreSlaInfoRecDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p preSlaInfoRecDo) Joins(fields ...field.RelationField) IPreSlaInfoRecDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p preSlaInfoRecDo) Preload(fields ...field.RelationField) IPreSlaInfoRecDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p preSlaInfoRecDo) FirstOrInit() (*premortem.PreSlaInfoRec, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.PreSlaInfoRec), nil
	}
}

func (p preSlaInfoRecDo) FirstOrCreate() (*premortem.PreSlaInfoRec, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*premortem.PreSlaInfoRec), nil
	}
}

func (p preSlaInfoRecDo) FindByPage(offset int, limit int) (result []*premortem.PreSlaInfoRec, count int64, err error) {
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

func (p preSlaInfoRecDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p preSlaInfoRecDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p preSlaInfoRecDo) Delete(models ...*premortem.PreSlaInfoRec) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *preSlaInfoRecDo) withDO(do gen.Dao) *preSlaInfoRecDo {
	p.DO = *do.(*gen.DO)
	return p
}