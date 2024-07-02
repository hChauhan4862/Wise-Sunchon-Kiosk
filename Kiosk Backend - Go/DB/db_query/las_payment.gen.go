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

func newLASPAYMENT(db *gorm.DB, opts ...gen.DOOption) lASPAYMENT {
	_lASPAYMENT := lASPAYMENT{}

	_lASPAYMENT.lASPAYMENTDo.UseDB(db, opts...)
	_lASPAYMENT.lASPAYMENTDo.UseModel(&model.LASPAYMENT{})

	tableName := _lASPAYMENT.lASPAYMENTDo.TableName()
	_lASPAYMENT.ALL = field.NewAsterisk(tableName)
	_lASPAYMENT.LASORDERNO = field.NewString(tableName, "LAS_ORDER_NO")
	_lASPAYMENT.USERID = field.NewString(tableName, "USER_ID")
	_lASPAYMENT.BOOK = field.NewString(tableName, "BOOK")
	_lASPAYMENT.DATECNT = field.NewString(tableName, "DATE_CNT")
	_lASPAYMENT.PAY = field.NewInt64(tableName, "PAY")
	_lASPAYMENT.LASNO = field.NewString(tableName, "LAS_NO")

	_lASPAYMENT.fillFieldMap()

	return _lASPAYMENT
}

type lASPAYMENT struct {
	lASPAYMENTDo

	ALL        field.Asterisk
	LASORDERNO field.String
	USERID     field.String
	BOOK       field.String
	DATECNT    field.String
	PAY        field.Int64
	LASNO      field.String

	fieldMap map[string]field.Expr
}

func (l lASPAYMENT) Table(newTableName string) *lASPAYMENT {
	l.lASPAYMENTDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lASPAYMENT) As(alias string) *lASPAYMENT {
	l.lASPAYMENTDo.DO = *(l.lASPAYMENTDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lASPAYMENT) updateTableName(table string) *lASPAYMENT {
	l.ALL = field.NewAsterisk(table)
	l.LASORDERNO = field.NewString(table, "LAS_ORDER_NO")
	l.USERID = field.NewString(table, "USER_ID")
	l.BOOK = field.NewString(table, "BOOK")
	l.DATECNT = field.NewString(table, "DATE_CNT")
	l.PAY = field.NewInt64(table, "PAY")
	l.LASNO = field.NewString(table, "LAS_NO")

	l.fillFieldMap()

	return l
}

func (l *lASPAYMENT) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lASPAYMENT) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["LAS_ORDER_NO"] = l.LASORDERNO
	l.fieldMap["USER_ID"] = l.USERID
	l.fieldMap["BOOK"] = l.BOOK
	l.fieldMap["DATE_CNT"] = l.DATECNT
	l.fieldMap["PAY"] = l.PAY
	l.fieldMap["LAS_NO"] = l.LASNO
}

func (l lASPAYMENT) clone(db *gorm.DB) lASPAYMENT {
	l.lASPAYMENTDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l lASPAYMENT) replaceDB(db *gorm.DB) lASPAYMENT {
	l.lASPAYMENTDo.ReplaceDB(db)
	return l
}

type lASPAYMENTDo struct{ gen.DO }

type ILASPAYMENTDo interface {
	gen.SubQuery
	Debug() ILASPAYMENTDo
	WithContext(ctx context.Context) ILASPAYMENTDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILASPAYMENTDo
	WriteDB() ILASPAYMENTDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILASPAYMENTDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILASPAYMENTDo
	Not(conds ...gen.Condition) ILASPAYMENTDo
	Or(conds ...gen.Condition) ILASPAYMENTDo
	Select(conds ...field.Expr) ILASPAYMENTDo
	Where(conds ...gen.Condition) ILASPAYMENTDo
	Order(conds ...field.Expr) ILASPAYMENTDo
	Distinct(cols ...field.Expr) ILASPAYMENTDo
	Omit(cols ...field.Expr) ILASPAYMENTDo
	Join(table schema.Tabler, on ...field.Expr) ILASPAYMENTDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILASPAYMENTDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILASPAYMENTDo
	Group(cols ...field.Expr) ILASPAYMENTDo
	Having(conds ...gen.Condition) ILASPAYMENTDo
	Limit(limit int) ILASPAYMENTDo
	Offset(offset int) ILASPAYMENTDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILASPAYMENTDo
	Unscoped() ILASPAYMENTDo
	Create(values ...*model.LASPAYMENT) error
	CreateInBatches(values []*model.LASPAYMENT, batchSize int) error
	Save(values ...*model.LASPAYMENT) error
	First() (*model.LASPAYMENT, error)
	Take() (*model.LASPAYMENT, error)
	Last() (*model.LASPAYMENT, error)
	Find() ([]*model.LASPAYMENT, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LASPAYMENT, err error)
	FindInBatches(result *[]*model.LASPAYMENT, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LASPAYMENT) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILASPAYMENTDo
	Assign(attrs ...field.AssignExpr) ILASPAYMENTDo
	Joins(fields ...field.RelationField) ILASPAYMENTDo
	Preload(fields ...field.RelationField) ILASPAYMENTDo
	FirstOrInit() (*model.LASPAYMENT, error)
	FirstOrCreate() (*model.LASPAYMENT, error)
	FindByPage(offset int, limit int) (result []*model.LASPAYMENT, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILASPAYMENTDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l lASPAYMENTDo) Debug() ILASPAYMENTDo {
	return l.withDO(l.DO.Debug())
}

func (l lASPAYMENTDo) WithContext(ctx context.Context) ILASPAYMENTDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lASPAYMENTDo) ReadDB() ILASPAYMENTDo {
	return l.Clauses(dbresolver.Read)
}

func (l lASPAYMENTDo) WriteDB() ILASPAYMENTDo {
	return l.Clauses(dbresolver.Write)
}

func (l lASPAYMENTDo) Session(config *gorm.Session) ILASPAYMENTDo {
	return l.withDO(l.DO.Session(config))
}

func (l lASPAYMENTDo) Clauses(conds ...clause.Expression) ILASPAYMENTDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lASPAYMENTDo) Returning(value interface{}, columns ...string) ILASPAYMENTDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lASPAYMENTDo) Not(conds ...gen.Condition) ILASPAYMENTDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lASPAYMENTDo) Or(conds ...gen.Condition) ILASPAYMENTDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lASPAYMENTDo) Select(conds ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lASPAYMENTDo) Where(conds ...gen.Condition) ILASPAYMENTDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lASPAYMENTDo) Order(conds ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lASPAYMENTDo) Distinct(cols ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lASPAYMENTDo) Omit(cols ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lASPAYMENTDo) Join(table schema.Tabler, on ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lASPAYMENTDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lASPAYMENTDo) RightJoin(table schema.Tabler, on ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lASPAYMENTDo) Group(cols ...field.Expr) ILASPAYMENTDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lASPAYMENTDo) Having(conds ...gen.Condition) ILASPAYMENTDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lASPAYMENTDo) Limit(limit int) ILASPAYMENTDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lASPAYMENTDo) Offset(offset int) ILASPAYMENTDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lASPAYMENTDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILASPAYMENTDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lASPAYMENTDo) Unscoped() ILASPAYMENTDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lASPAYMENTDo) Create(values ...*model.LASPAYMENT) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lASPAYMENTDo) CreateInBatches(values []*model.LASPAYMENT, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lASPAYMENTDo) Save(values ...*model.LASPAYMENT) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lASPAYMENTDo) First() (*model.LASPAYMENT, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LASPAYMENT), nil
	}
}

func (l lASPAYMENTDo) Take() (*model.LASPAYMENT, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LASPAYMENT), nil
	}
}

func (l lASPAYMENTDo) Last() (*model.LASPAYMENT, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LASPAYMENT), nil
	}
}

func (l lASPAYMENTDo) Find() ([]*model.LASPAYMENT, error) {
	result, err := l.DO.Find()
	return result.([]*model.LASPAYMENT), err
}

func (l lASPAYMENTDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LASPAYMENT, err error) {
	buf := make([]*model.LASPAYMENT, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lASPAYMENTDo) FindInBatches(result *[]*model.LASPAYMENT, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lASPAYMENTDo) Attrs(attrs ...field.AssignExpr) ILASPAYMENTDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lASPAYMENTDo) Assign(attrs ...field.AssignExpr) ILASPAYMENTDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lASPAYMENTDo) Joins(fields ...field.RelationField) ILASPAYMENTDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lASPAYMENTDo) Preload(fields ...field.RelationField) ILASPAYMENTDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lASPAYMENTDo) FirstOrInit() (*model.LASPAYMENT, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LASPAYMENT), nil
	}
}

func (l lASPAYMENTDo) FirstOrCreate() (*model.LASPAYMENT, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LASPAYMENT), nil
	}
}

func (l lASPAYMENTDo) FindByPage(offset int, limit int) (result []*model.LASPAYMENT, count int64, err error) {
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

func (l lASPAYMENTDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lASPAYMENTDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lASPAYMENTDo) Delete(models ...*model.LASPAYMENT) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lASPAYMENTDo) withDO(do gen.Dao) *lASPAYMENTDo {
	l.DO = *do.(*gen.DO)
	return l
}