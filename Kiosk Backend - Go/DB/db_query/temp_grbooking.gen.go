// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db_query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"WISE_SOFTWARE/DB/model"
)

func newTempGrbooking(db *gorm.DB, opts ...gen.DOOption) tempGrbooking {
	_tempGrbooking := tempGrbooking{}

	_tempGrbooking.tempGrbookingDo.UseDB(db, opts...)
	_tempGrbooking.tempGrbookingDo.UseModel(&model.TempGrbooking{})

	tableName := _tempGrbooking.tempGrbookingDo.TableName()
	_tempGrbooking.ALL = field.NewAsterisk(tableName)
	_tempGrbooking.BSEQNO = field.NewInt64(tableName, "BSEQNO")
	_tempGrbooking.MEMBERS = field.NewString(tableName, "MEMBERS")

	_tempGrbooking.fillFieldMap()

	return _tempGrbooking
}

type tempGrbooking struct {
	tempGrbookingDo

	ALL     field.Asterisk
	BSEQNO  field.Int64
	MEMBERS field.String

	fieldMap map[string]field.Expr
}

func (t tempGrbooking) Table(newTableName string) *tempGrbooking {
	t.tempGrbookingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tempGrbooking) As(alias string) *tempGrbooking {
	t.tempGrbookingDo.DO = *(t.tempGrbookingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tempGrbooking) updateTableName(table string) *tempGrbooking {
	t.ALL = field.NewAsterisk(table)
	t.BSEQNO = field.NewInt64(table, "BSEQNO")
	t.MEMBERS = field.NewString(table, "MEMBERS")

	t.fillFieldMap()

	return t
}

func (t *tempGrbooking) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tempGrbooking) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["BSEQNO"] = t.BSEQNO
	t.fieldMap["MEMBERS"] = t.MEMBERS
}

func (t tempGrbooking) clone(db *gorm.DB) tempGrbooking {
	t.tempGrbookingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tempGrbooking) replaceDB(db *gorm.DB) tempGrbooking {
	t.tempGrbookingDo.ReplaceDB(db)
	return t
}

type tempGrbookingDo struct{ gen.DO }

type ITempGrbookingDo interface {
	gen.SubQuery
	Debug() ITempGrbookingDo
	WithContext(ctx context.Context) ITempGrbookingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITempGrbookingDo
	WriteDB() ITempGrbookingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITempGrbookingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITempGrbookingDo
	Not(conds ...gen.Condition) ITempGrbookingDo
	Or(conds ...gen.Condition) ITempGrbookingDo
	Select(conds ...field.Expr) ITempGrbookingDo
	Where(conds ...gen.Condition) ITempGrbookingDo
	Order(conds ...field.Expr) ITempGrbookingDo
	Distinct(cols ...field.Expr) ITempGrbookingDo
	Omit(cols ...field.Expr) ITempGrbookingDo
	Join(table schema.Tabler, on ...field.Expr) ITempGrbookingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITempGrbookingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITempGrbookingDo
	Group(cols ...field.Expr) ITempGrbookingDo
	Having(conds ...gen.Condition) ITempGrbookingDo
	Limit(limit int) ITempGrbookingDo
	Offset(offset int) ITempGrbookingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITempGrbookingDo
	Unscoped() ITempGrbookingDo
	Create(values ...*model.TempGrbooking) error
	CreateInBatches(values []*model.TempGrbooking, batchSize int) error
	Save(values ...*model.TempGrbooking) error
	First() (*model.TempGrbooking, error)
	Take() (*model.TempGrbooking, error)
	Last() (*model.TempGrbooking, error)
	Find() ([]*model.TempGrbooking, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TempGrbooking, err error)
	FindInBatches(result *[]*model.TempGrbooking, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TempGrbooking) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITempGrbookingDo
	Assign(attrs ...field.AssignExpr) ITempGrbookingDo
	Joins(fields ...field.RelationField) ITempGrbookingDo
	Preload(fields ...field.RelationField) ITempGrbookingDo
	FirstOrInit() (*model.TempGrbooking, error)
	FirstOrCreate() (*model.TempGrbooking, error)
	FindByPage(offset int, limit int) (result []*model.TempGrbooking, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITempGrbookingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tempGrbookingDo) Debug() ITempGrbookingDo {
	return t.withDO(t.DO.Debug())
}

func (t tempGrbookingDo) WithContext(ctx context.Context) ITempGrbookingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tempGrbookingDo) ReadDB() ITempGrbookingDo {
	return t.Clauses(dbresolver.Read)
}

func (t tempGrbookingDo) WriteDB() ITempGrbookingDo {
	return t.Clauses(dbresolver.Write)
}

func (t tempGrbookingDo) Session(config *gorm.Session) ITempGrbookingDo {
	return t.withDO(t.DO.Session(config))
}

func (t tempGrbookingDo) Clauses(conds ...clause.Expression) ITempGrbookingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tempGrbookingDo) Returning(value interface{}, columns ...string) ITempGrbookingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tempGrbookingDo) Not(conds ...gen.Condition) ITempGrbookingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tempGrbookingDo) Or(conds ...gen.Condition) ITempGrbookingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tempGrbookingDo) Select(conds ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tempGrbookingDo) Where(conds ...gen.Condition) ITempGrbookingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tempGrbookingDo) Order(conds ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tempGrbookingDo) Distinct(cols ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tempGrbookingDo) Omit(cols ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tempGrbookingDo) Join(table schema.Tabler, on ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tempGrbookingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tempGrbookingDo) RightJoin(table schema.Tabler, on ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tempGrbookingDo) Group(cols ...field.Expr) ITempGrbookingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tempGrbookingDo) Having(conds ...gen.Condition) ITempGrbookingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tempGrbookingDo) Limit(limit int) ITempGrbookingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tempGrbookingDo) Offset(offset int) ITempGrbookingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tempGrbookingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITempGrbookingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tempGrbookingDo) Unscoped() ITempGrbookingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tempGrbookingDo) Create(values ...*model.TempGrbooking) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tempGrbookingDo) CreateInBatches(values []*model.TempGrbooking, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tempGrbookingDo) Save(values ...*model.TempGrbooking) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tempGrbookingDo) First() (*model.TempGrbooking, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TempGrbooking), nil
	}
}

func (t tempGrbookingDo) Take() (*model.TempGrbooking, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TempGrbooking), nil
	}
}

func (t tempGrbookingDo) Last() (*model.TempGrbooking, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TempGrbooking), nil
	}
}

func (t tempGrbookingDo) Find() ([]*model.TempGrbooking, error) {
	result, err := t.DO.Find()
	return result.([]*model.TempGrbooking), err
}

func (t tempGrbookingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TempGrbooking, err error) {
	buf := make([]*model.TempGrbooking, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tempGrbookingDo) FindInBatches(result *[]*model.TempGrbooking, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tempGrbookingDo) Attrs(attrs ...field.AssignExpr) ITempGrbookingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tempGrbookingDo) Assign(attrs ...field.AssignExpr) ITempGrbookingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tempGrbookingDo) Joins(fields ...field.RelationField) ITempGrbookingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tempGrbookingDo) Preload(fields ...field.RelationField) ITempGrbookingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tempGrbookingDo) FirstOrInit() (*model.TempGrbooking, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TempGrbooking), nil
	}
}

func (t tempGrbookingDo) FirstOrCreate() (*model.TempGrbooking, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TempGrbooking), nil
	}
}

func (t tempGrbookingDo) FindByPage(offset int, limit int) (result []*model.TempGrbooking, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tempGrbookingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tempGrbookingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tempGrbookingDo) Delete(models ...*model.TempGrbooking) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tempGrbookingDo) withDO(do gen.Dao) *tempGrbookingDo {
	t.DO = *do.(*gen.DO)
	return t
}
