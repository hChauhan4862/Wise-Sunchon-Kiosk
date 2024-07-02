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

func newHCV_TV_YESTERDAY_USE_TIME2(db *gorm.DB, opts ...gen.DOOption) hCV_TV_YESTERDAY_USE_TIME2 {
	_hCV_TV_YESTERDAY_USE_TIME2 := hCV_TV_YESTERDAY_USE_TIME2{}

	_hCV_TV_YESTERDAY_USE_TIME2.hCV_TV_YESTERDAY_USE_TIME2Do.UseDB(db, opts...)
	_hCV_TV_YESTERDAY_USE_TIME2.hCV_TV_YESTERDAY_USE_TIME2Do.UseModel(&model.HCV_TV_YESTERDAY_USE_TIME2{})

	tableName := _hCV_TV_YESTERDAY_USE_TIME2.hCV_TV_YESTERDAY_USE_TIME2Do.TableName()
	_hCV_TV_YESTERDAY_USE_TIME2.ALL = field.NewAsterisk(tableName)
	_hCV_TV_YESTERDAY_USE_TIME2.Libno = field.NewInt64(tableName, "libno")
	_hCV_TV_YESTERDAY_USE_TIME2.UseDate = field.NewString(tableName, "use_date")
	_hCV_TV_YESTERDAY_USE_TIME2.Status = field.NewInt64(tableName, "status")
	_hCV_TV_YESTERDAY_USE_TIME2.Gubun = field.NewInt64(tableName, "gubun")
	_hCV_TV_YESTERDAY_USE_TIME2.T00 = field.NewInt64(tableName, "t00")
	_hCV_TV_YESTERDAY_USE_TIME2.T01 = field.NewInt64(tableName, "t01")
	_hCV_TV_YESTERDAY_USE_TIME2.T02 = field.NewInt64(tableName, "t02")
	_hCV_TV_YESTERDAY_USE_TIME2.T03 = field.NewInt64(tableName, "t03")
	_hCV_TV_YESTERDAY_USE_TIME2.T04 = field.NewInt64(tableName, "t04")
	_hCV_TV_YESTERDAY_USE_TIME2.T05 = field.NewInt64(tableName, "t05")
	_hCV_TV_YESTERDAY_USE_TIME2.T06 = field.NewInt64(tableName, "t06")
	_hCV_TV_YESTERDAY_USE_TIME2.T07 = field.NewInt64(tableName, "t07")
	_hCV_TV_YESTERDAY_USE_TIME2.T08 = field.NewInt64(tableName, "t08")
	_hCV_TV_YESTERDAY_USE_TIME2.T09 = field.NewInt64(tableName, "t09")
	_hCV_TV_YESTERDAY_USE_TIME2.T10 = field.NewInt64(tableName, "t10")
	_hCV_TV_YESTERDAY_USE_TIME2.T11 = field.NewInt64(tableName, "t11")
	_hCV_TV_YESTERDAY_USE_TIME2.T12 = field.NewInt64(tableName, "t12")
	_hCV_TV_YESTERDAY_USE_TIME2.T13 = field.NewInt64(tableName, "t13")
	_hCV_TV_YESTERDAY_USE_TIME2.T14 = field.NewInt64(tableName, "t14")
	_hCV_TV_YESTERDAY_USE_TIME2.T15 = field.NewInt64(tableName, "t15")
	_hCV_TV_YESTERDAY_USE_TIME2.T16 = field.NewInt64(tableName, "t16")
	_hCV_TV_YESTERDAY_USE_TIME2.T17 = field.NewInt64(tableName, "t17")
	_hCV_TV_YESTERDAY_USE_TIME2.T18 = field.NewInt64(tableName, "t18")
	_hCV_TV_YESTERDAY_USE_TIME2.T19 = field.NewInt64(tableName, "t19")
	_hCV_TV_YESTERDAY_USE_TIME2.T20 = field.NewInt64(tableName, "t20")
	_hCV_TV_YESTERDAY_USE_TIME2.T21 = field.NewInt64(tableName, "t21")
	_hCV_TV_YESTERDAY_USE_TIME2.T22 = field.NewInt64(tableName, "t22")
	_hCV_TV_YESTERDAY_USE_TIME2.T23 = field.NewInt64(tableName, "t23")

	_hCV_TV_YESTERDAY_USE_TIME2.fillFieldMap()

	return _hCV_TV_YESTERDAY_USE_TIME2
}

type hCV_TV_YESTERDAY_USE_TIME2 struct {
	hCV_TV_YESTERDAY_USE_TIME2Do

	ALL     field.Asterisk
	Libno   field.Int64
	UseDate field.String
	Status  field.Int64
	Gubun   field.Int64
	T00     field.Int64
	T01     field.Int64
	T02     field.Int64
	T03     field.Int64
	T04     field.Int64
	T05     field.Int64
	T06     field.Int64
	T07     field.Int64
	T08     field.Int64
	T09     field.Int64
	T10     field.Int64
	T11     field.Int64
	T12     field.Int64
	T13     field.Int64
	T14     field.Int64
	T15     field.Int64
	T16     field.Int64
	T17     field.Int64
	T18     field.Int64
	T19     field.Int64
	T20     field.Int64
	T21     field.Int64
	T22     field.Int64
	T23     field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_TV_YESTERDAY_USE_TIME2) Table(newTableName string) *hCV_TV_YESTERDAY_USE_TIME2 {
	h.hCV_TV_YESTERDAY_USE_TIME2Do.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_TV_YESTERDAY_USE_TIME2) As(alias string) *hCV_TV_YESTERDAY_USE_TIME2 {
	h.hCV_TV_YESTERDAY_USE_TIME2Do.DO = *(h.hCV_TV_YESTERDAY_USE_TIME2Do.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_TV_YESTERDAY_USE_TIME2) updateTableName(table string) *hCV_TV_YESTERDAY_USE_TIME2 {
	h.ALL = field.NewAsterisk(table)
	h.Libno = field.NewInt64(table, "libno")
	h.UseDate = field.NewString(table, "use_date")
	h.Status = field.NewInt64(table, "status")
	h.Gubun = field.NewInt64(table, "gubun")
	h.T00 = field.NewInt64(table, "t00")
	h.T01 = field.NewInt64(table, "t01")
	h.T02 = field.NewInt64(table, "t02")
	h.T03 = field.NewInt64(table, "t03")
	h.T04 = field.NewInt64(table, "t04")
	h.T05 = field.NewInt64(table, "t05")
	h.T06 = field.NewInt64(table, "t06")
	h.T07 = field.NewInt64(table, "t07")
	h.T08 = field.NewInt64(table, "t08")
	h.T09 = field.NewInt64(table, "t09")
	h.T10 = field.NewInt64(table, "t10")
	h.T11 = field.NewInt64(table, "t11")
	h.T12 = field.NewInt64(table, "t12")
	h.T13 = field.NewInt64(table, "t13")
	h.T14 = field.NewInt64(table, "t14")
	h.T15 = field.NewInt64(table, "t15")
	h.T16 = field.NewInt64(table, "t16")
	h.T17 = field.NewInt64(table, "t17")
	h.T18 = field.NewInt64(table, "t18")
	h.T19 = field.NewInt64(table, "t19")
	h.T20 = field.NewInt64(table, "t20")
	h.T21 = field.NewInt64(table, "t21")
	h.T22 = field.NewInt64(table, "t22")
	h.T23 = field.NewInt64(table, "t23")

	h.fillFieldMap()

	return h
}

func (h *hCV_TV_YESTERDAY_USE_TIME2) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_TV_YESTERDAY_USE_TIME2) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 28)
	h.fieldMap["libno"] = h.Libno
	h.fieldMap["use_date"] = h.UseDate
	h.fieldMap["status"] = h.Status
	h.fieldMap["gubun"] = h.Gubun
	h.fieldMap["t00"] = h.T00
	h.fieldMap["t01"] = h.T01
	h.fieldMap["t02"] = h.T02
	h.fieldMap["t03"] = h.T03
	h.fieldMap["t04"] = h.T04
	h.fieldMap["t05"] = h.T05
	h.fieldMap["t06"] = h.T06
	h.fieldMap["t07"] = h.T07
	h.fieldMap["t08"] = h.T08
	h.fieldMap["t09"] = h.T09
	h.fieldMap["t10"] = h.T10
	h.fieldMap["t11"] = h.T11
	h.fieldMap["t12"] = h.T12
	h.fieldMap["t13"] = h.T13
	h.fieldMap["t14"] = h.T14
	h.fieldMap["t15"] = h.T15
	h.fieldMap["t16"] = h.T16
	h.fieldMap["t17"] = h.T17
	h.fieldMap["t18"] = h.T18
	h.fieldMap["t19"] = h.T19
	h.fieldMap["t20"] = h.T20
	h.fieldMap["t21"] = h.T21
	h.fieldMap["t22"] = h.T22
	h.fieldMap["t23"] = h.T23
}

func (h hCV_TV_YESTERDAY_USE_TIME2) clone(db *gorm.DB) hCV_TV_YESTERDAY_USE_TIME2 {
	h.hCV_TV_YESTERDAY_USE_TIME2Do.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_TV_YESTERDAY_USE_TIME2) replaceDB(db *gorm.DB) hCV_TV_YESTERDAY_USE_TIME2 {
	h.hCV_TV_YESTERDAY_USE_TIME2Do.ReplaceDB(db)
	return h
}

type hCV_TV_YESTERDAY_USE_TIME2Do struct{ gen.DO }

type IHCV_TV_YESTERDAY_USE_TIME2Do interface {
	gen.SubQuery
	Debug() IHCV_TV_YESTERDAY_USE_TIME2Do
	WithContext(ctx context.Context) IHCV_TV_YESTERDAY_USE_TIME2Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_TV_YESTERDAY_USE_TIME2Do
	WriteDB() IHCV_TV_YESTERDAY_USE_TIME2Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_TV_YESTERDAY_USE_TIME2Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_TV_YESTERDAY_USE_TIME2Do
	Not(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do
	Or(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do
	Select(conds ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Where(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do
	Order(conds ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Distinct(cols ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Omit(cols ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Join(table schema.Tabler, on ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Group(cols ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Having(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do
	Limit(limit int) IHCV_TV_YESTERDAY_USE_TIME2Do
	Offset(offset int) IHCV_TV_YESTERDAY_USE_TIME2Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_TV_YESTERDAY_USE_TIME2Do
	Unscoped() IHCV_TV_YESTERDAY_USE_TIME2Do
	Create(values ...*model.HCV_TV_YESTERDAY_USE_TIME2) error
	CreateInBatches(values []*model.HCV_TV_YESTERDAY_USE_TIME2, batchSize int) error
	Save(values ...*model.HCV_TV_YESTERDAY_USE_TIME2) error
	First() (*model.HCV_TV_YESTERDAY_USE_TIME2, error)
	Take() (*model.HCV_TV_YESTERDAY_USE_TIME2, error)
	Last() (*model.HCV_TV_YESTERDAY_USE_TIME2, error)
	Find() ([]*model.HCV_TV_YESTERDAY_USE_TIME2, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_TV_YESTERDAY_USE_TIME2, err error)
	FindInBatches(result *[]*model.HCV_TV_YESTERDAY_USE_TIME2, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_TV_YESTERDAY_USE_TIME2) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Assign(attrs ...field.AssignExpr) IHCV_TV_YESTERDAY_USE_TIME2Do
	Joins(fields ...field.RelationField) IHCV_TV_YESTERDAY_USE_TIME2Do
	Preload(fields ...field.RelationField) IHCV_TV_YESTERDAY_USE_TIME2Do
	FirstOrInit() (*model.HCV_TV_YESTERDAY_USE_TIME2, error)
	FirstOrCreate() (*model.HCV_TV_YESTERDAY_USE_TIME2, error)
	FindByPage(offset int, limit int) (result []*model.HCV_TV_YESTERDAY_USE_TIME2, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_TV_YESTERDAY_USE_TIME2Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Debug() IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Debug())
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) WithContext(ctx context.Context) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) ReadDB() IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) WriteDB() IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Session(config *gorm.Session) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Clauses(conds ...clause.Expression) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Returning(value interface{}, columns ...string) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Not(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Or(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Select(conds ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Where(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Order(conds ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Distinct(cols ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Omit(cols ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Join(table schema.Tabler, on ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Group(cols ...field.Expr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Having(conds ...gen.Condition) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Limit(limit int) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Offset(offset int) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Unscoped() IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Create(values ...*model.HCV_TV_YESTERDAY_USE_TIME2) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) CreateInBatches(values []*model.HCV_TV_YESTERDAY_USE_TIME2, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_TV_YESTERDAY_USE_TIME2Do) Save(values ...*model.HCV_TV_YESTERDAY_USE_TIME2) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) First() (*model.HCV_TV_YESTERDAY_USE_TIME2, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_TV_YESTERDAY_USE_TIME2), nil
	}
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Take() (*model.HCV_TV_YESTERDAY_USE_TIME2, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_TV_YESTERDAY_USE_TIME2), nil
	}
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Last() (*model.HCV_TV_YESTERDAY_USE_TIME2, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_TV_YESTERDAY_USE_TIME2), nil
	}
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Find() ([]*model.HCV_TV_YESTERDAY_USE_TIME2, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_TV_YESTERDAY_USE_TIME2), err
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_TV_YESTERDAY_USE_TIME2, err error) {
	buf := make([]*model.HCV_TV_YESTERDAY_USE_TIME2, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) FindInBatches(result *[]*model.HCV_TV_YESTERDAY_USE_TIME2, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Attrs(attrs ...field.AssignExpr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Assign(attrs ...field.AssignExpr) IHCV_TV_YESTERDAY_USE_TIME2Do {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Joins(fields ...field.RelationField) IHCV_TV_YESTERDAY_USE_TIME2Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Preload(fields ...field.RelationField) IHCV_TV_YESTERDAY_USE_TIME2Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) FirstOrInit() (*model.HCV_TV_YESTERDAY_USE_TIME2, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_TV_YESTERDAY_USE_TIME2), nil
	}
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) FirstOrCreate() (*model.HCV_TV_YESTERDAY_USE_TIME2, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_TV_YESTERDAY_USE_TIME2), nil
	}
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) FindByPage(offset int, limit int) (result []*model.HCV_TV_YESTERDAY_USE_TIME2, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_TV_YESTERDAY_USE_TIME2Do) Delete(models ...*model.HCV_TV_YESTERDAY_USE_TIME2) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_TV_YESTERDAY_USE_TIME2Do) withDO(do gen.Dao) *hCV_TV_YESTERDAY_USE_TIME2Do {
	h.DO = *do.(*gen.DO)
	return h
}