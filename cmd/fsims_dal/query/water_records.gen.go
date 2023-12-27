// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newWaterRecord(db *gorm.DB, opts ...gen.DOOption) waterRecord {
	_waterRecord := waterRecord{}

	_waterRecord.waterRecordDo.UseDB(db, opts...)
	_waterRecord.waterRecordDo.UseModel(&pasture.WaterRecord{})

	tableName := _waterRecord.waterRecordDo.TableName()
	_waterRecord.ALL = field.NewAsterisk(tableName)
	_waterRecord.ID = field.NewUint(tableName, "id")
	_waterRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_waterRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_waterRecord.DeletedAt = field.NewField(tableName, "deleted_at")
	_waterRecord.TimeRecordAt = field.NewString(tableName, "time_record_at")
	_waterRecord.OapGci = waterRecordHasOneOapGci{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("OapGci", "pasture.OapGci"),
	}

	_waterRecord.ToxIndex = waterRecordHasOneToxIndex{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ToxIndex", "pasture.ToxIndex"),
	}

	_waterRecord.MicroIndex = waterRecordHasOneMicroIndex{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("MicroIndex", "pasture.MicroIndex"),
	}

	_waterRecord.fillFieldMap()

	return _waterRecord
}

type waterRecord struct {
	waterRecordDo waterRecordDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	TimeRecordAt field.String
	OapGci       waterRecordHasOneOapGci

	ToxIndex waterRecordHasOneToxIndex

	MicroIndex waterRecordHasOneMicroIndex

	fieldMap map[string]field.Expr
}

func (w waterRecord) Table(newTableName string) *waterRecord {
	w.waterRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w waterRecord) As(alias string) *waterRecord {
	w.waterRecordDo.DO = *(w.waterRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *waterRecord) updateTableName(table string) *waterRecord {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.TimeRecordAt = field.NewString(table, "time_record_at")

	w.fillFieldMap()

	return w
}

func (w *waterRecord) WithContext(ctx context.Context) IWaterRecordDo {
	return w.waterRecordDo.WithContext(ctx)
}

func (w waterRecord) TableName() string { return w.waterRecordDo.TableName() }

func (w waterRecord) Alias() string { return w.waterRecordDo.Alias() }

func (w waterRecord) Columns(cols ...field.Expr) gen.Columns { return w.waterRecordDo.Columns(cols...) }

func (w *waterRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *waterRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["time_record_at"] = w.TimeRecordAt

}

func (w waterRecord) clone(db *gorm.DB) waterRecord {
	w.waterRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w waterRecord) replaceDB(db *gorm.DB) waterRecord {
	w.waterRecordDo.ReplaceDB(db)
	return w
}

type waterRecordHasOneOapGci struct {
	db *gorm.DB

	field.RelationField
}

func (a waterRecordHasOneOapGci) Where(conds ...field.Expr) *waterRecordHasOneOapGci {
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

func (a waterRecordHasOneOapGci) WithContext(ctx context.Context) *waterRecordHasOneOapGci {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a waterRecordHasOneOapGci) Session(session *gorm.Session) *waterRecordHasOneOapGci {
	a.db = a.db.Session(session)
	return &a
}

func (a waterRecordHasOneOapGci) Model(m *pasture.WaterRecord) *waterRecordHasOneOapGciTx {
	return &waterRecordHasOneOapGciTx{a.db.Model(m).Association(a.Name())}
}

type waterRecordHasOneOapGciTx struct{ tx *gorm.Association }

func (a waterRecordHasOneOapGciTx) Find() (result *pasture.OapGci, err error) {
	return result, a.tx.Find(&result)
}

func (a waterRecordHasOneOapGciTx) Append(values ...*pasture.OapGci) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a waterRecordHasOneOapGciTx) Replace(values ...*pasture.OapGci) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a waterRecordHasOneOapGciTx) Delete(values ...*pasture.OapGci) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a waterRecordHasOneOapGciTx) Clear() error {
	return a.tx.Clear()
}

func (a waterRecordHasOneOapGciTx) Count() int64 {
	return a.tx.Count()
}

type waterRecordHasOneToxIndex struct {
	db *gorm.DB

	field.RelationField
}

func (a waterRecordHasOneToxIndex) Where(conds ...field.Expr) *waterRecordHasOneToxIndex {
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

func (a waterRecordHasOneToxIndex) WithContext(ctx context.Context) *waterRecordHasOneToxIndex {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a waterRecordHasOneToxIndex) Session(session *gorm.Session) *waterRecordHasOneToxIndex {
	a.db = a.db.Session(session)
	return &a
}

func (a waterRecordHasOneToxIndex) Model(m *pasture.WaterRecord) *waterRecordHasOneToxIndexTx {
	return &waterRecordHasOneToxIndexTx{a.db.Model(m).Association(a.Name())}
}

type waterRecordHasOneToxIndexTx struct{ tx *gorm.Association }

func (a waterRecordHasOneToxIndexTx) Find() (result *pasture.ToxIndex, err error) {
	return result, a.tx.Find(&result)
}

func (a waterRecordHasOneToxIndexTx) Append(values ...*pasture.ToxIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a waterRecordHasOneToxIndexTx) Replace(values ...*pasture.ToxIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a waterRecordHasOneToxIndexTx) Delete(values ...*pasture.ToxIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a waterRecordHasOneToxIndexTx) Clear() error {
	return a.tx.Clear()
}

func (a waterRecordHasOneToxIndexTx) Count() int64 {
	return a.tx.Count()
}

type waterRecordHasOneMicroIndex struct {
	db *gorm.DB

	field.RelationField
}

func (a waterRecordHasOneMicroIndex) Where(conds ...field.Expr) *waterRecordHasOneMicroIndex {
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

func (a waterRecordHasOneMicroIndex) WithContext(ctx context.Context) *waterRecordHasOneMicroIndex {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a waterRecordHasOneMicroIndex) Session(session *gorm.Session) *waterRecordHasOneMicroIndex {
	a.db = a.db.Session(session)
	return &a
}

func (a waterRecordHasOneMicroIndex) Model(m *pasture.WaterRecord) *waterRecordHasOneMicroIndexTx {
	return &waterRecordHasOneMicroIndexTx{a.db.Model(m).Association(a.Name())}
}

type waterRecordHasOneMicroIndexTx struct{ tx *gorm.Association }

func (a waterRecordHasOneMicroIndexTx) Find() (result *pasture.MicroIndex, err error) {
	return result, a.tx.Find(&result)
}

func (a waterRecordHasOneMicroIndexTx) Append(values ...*pasture.MicroIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a waterRecordHasOneMicroIndexTx) Replace(values ...*pasture.MicroIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a waterRecordHasOneMicroIndexTx) Delete(values ...*pasture.MicroIndex) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a waterRecordHasOneMicroIndexTx) Clear() error {
	return a.tx.Clear()
}

func (a waterRecordHasOneMicroIndexTx) Count() int64 {
	return a.tx.Count()
}

type waterRecordDo struct{ gen.DO }

type IWaterRecordDo interface {
	gen.SubQuery
	Debug() IWaterRecordDo
	WithContext(ctx context.Context) IWaterRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWaterRecordDo
	WriteDB() IWaterRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWaterRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWaterRecordDo
	Not(conds ...gen.Condition) IWaterRecordDo
	Or(conds ...gen.Condition) IWaterRecordDo
	Select(conds ...field.Expr) IWaterRecordDo
	Where(conds ...gen.Condition) IWaterRecordDo
	Order(conds ...field.Expr) IWaterRecordDo
	Distinct(cols ...field.Expr) IWaterRecordDo
	Omit(cols ...field.Expr) IWaterRecordDo
	Join(table schema.Tabler, on ...field.Expr) IWaterRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWaterRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWaterRecordDo
	Group(cols ...field.Expr) IWaterRecordDo
	Having(conds ...gen.Condition) IWaterRecordDo
	Limit(limit int) IWaterRecordDo
	Offset(offset int) IWaterRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWaterRecordDo
	Unscoped() IWaterRecordDo
	Create(values ...*pasture.WaterRecord) error
	CreateInBatches(values []*pasture.WaterRecord, batchSize int) error
	Save(values ...*pasture.WaterRecord) error
	First() (*pasture.WaterRecord, error)
	Take() (*pasture.WaterRecord, error)
	Last() (*pasture.WaterRecord, error)
	Find() ([]*pasture.WaterRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.WaterRecord, err error)
	FindInBatches(result *[]*pasture.WaterRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pasture.WaterRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWaterRecordDo
	Assign(attrs ...field.AssignExpr) IWaterRecordDo
	Joins(fields ...field.RelationField) IWaterRecordDo
	Preload(fields ...field.RelationField) IWaterRecordDo
	FirstOrInit() (*pasture.WaterRecord, error)
	FirstOrCreate() (*pasture.WaterRecord, error)
	FindByPage(offset int, limit int) (result []*pasture.WaterRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWaterRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w waterRecordDo) Debug() IWaterRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w waterRecordDo) WithContext(ctx context.Context) IWaterRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w waterRecordDo) ReadDB() IWaterRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w waterRecordDo) WriteDB() IWaterRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w waterRecordDo) Session(config *gorm.Session) IWaterRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w waterRecordDo) Clauses(conds ...clause.Expression) IWaterRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w waterRecordDo) Returning(value interface{}, columns ...string) IWaterRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w waterRecordDo) Not(conds ...gen.Condition) IWaterRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w waterRecordDo) Or(conds ...gen.Condition) IWaterRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w waterRecordDo) Select(conds ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w waterRecordDo) Where(conds ...gen.Condition) IWaterRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w waterRecordDo) Order(conds ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w waterRecordDo) Distinct(cols ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w waterRecordDo) Omit(cols ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w waterRecordDo) Join(table schema.Tabler, on ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w waterRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w waterRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w waterRecordDo) Group(cols ...field.Expr) IWaterRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w waterRecordDo) Having(conds ...gen.Condition) IWaterRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w waterRecordDo) Limit(limit int) IWaterRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w waterRecordDo) Offset(offset int) IWaterRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w waterRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWaterRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w waterRecordDo) Unscoped() IWaterRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w waterRecordDo) Create(values ...*pasture.WaterRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w waterRecordDo) CreateInBatches(values []*pasture.WaterRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w waterRecordDo) Save(values ...*pasture.WaterRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w waterRecordDo) First() (*pasture.WaterRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WaterRecord), nil
	}
}

func (w waterRecordDo) Take() (*pasture.WaterRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WaterRecord), nil
	}
}

func (w waterRecordDo) Last() (*pasture.WaterRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WaterRecord), nil
	}
}

func (w waterRecordDo) Find() ([]*pasture.WaterRecord, error) {
	result, err := w.DO.Find()
	return result.([]*pasture.WaterRecord), err
}

func (w waterRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pasture.WaterRecord, err error) {
	buf := make([]*pasture.WaterRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w waterRecordDo) FindInBatches(result *[]*pasture.WaterRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w waterRecordDo) Attrs(attrs ...field.AssignExpr) IWaterRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w waterRecordDo) Assign(attrs ...field.AssignExpr) IWaterRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w waterRecordDo) Joins(fields ...field.RelationField) IWaterRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w waterRecordDo) Preload(fields ...field.RelationField) IWaterRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w waterRecordDo) FirstOrInit() (*pasture.WaterRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WaterRecord), nil
	}
}

func (w waterRecordDo) FirstOrCreate() (*pasture.WaterRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pasture.WaterRecord), nil
	}
}

func (w waterRecordDo) FindByPage(offset int, limit int) (result []*pasture.WaterRecord, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w waterRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w waterRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w waterRecordDo) Delete(models ...*pasture.WaterRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *waterRecordDo) withDO(do gen.Dao) *waterRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}
