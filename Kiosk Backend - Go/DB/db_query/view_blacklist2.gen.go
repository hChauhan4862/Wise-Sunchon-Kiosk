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

func newHCV_VIEW_BLACKLIST2(db *gorm.DB, opts ...gen.DOOption) hCV_VIEW_BLACKLIST2 {
	_hCV_VIEW_BLACKLIST2 := hCV_VIEW_BLACKLIST2{}

	_hCV_VIEW_BLACKLIST2.hCV_VIEW_BLACKLIST2Do.UseDB(db, opts...)
	_hCV_VIEW_BLACKLIST2.hCV_VIEW_BLACKLIST2Do.UseModel(&model.HCV_VIEW_BLACKLIST2{})

	tableName := _hCV_VIEW_BLACKLIST2.hCV_VIEW_BLACKLIST2Do.TableName()
	_hCV_VIEW_BLACKLIST2.ALL = field.NewAsterisk(tableName)
	_hCV_VIEW_BLACKLIST2.BLSEQNO = field.NewInt64(tableName, "BLSEQNO")
	_hCV_VIEW_BLACKLIST2.SCHOOLNO = field.NewString(tableName, "SCHOOLNO")
	_hCV_VIEW_BLACKLIST2.USERNAME = field.NewString(tableName, "USER_NAME")
	_hCV_VIEW_BLACKLIST2.REASON = field.NewString(tableName, "REASON")
	_hCV_VIEW_BLACKLIST2.REGTIME = field.NewString(tableName, "REGTIME")
	_hCV_VIEW_BLACKLIST2.BLOCKSTART = field.NewString(tableName, "BLOCKSTART")
	_hCV_VIEW_BLACKLIST2.BLOCKEXPIRE = field.NewString(tableName, "BLOCKEXPIRE")
	_hCV_VIEW_BLACKLIST2.STATUS = field.NewInt64(tableName, "STATUS")
	_hCV_VIEW_BLACKLIST2.SETFROM = field.NewString(tableName, "SETFROM")
	_hCV_VIEW_BLACKLIST2.BLDAYBEFORE = field.NewInt64(tableName, "BL_DAYBEFORE")
	_hCV_VIEW_BLACKLIST2.BLLIMITCNT = field.NewInt64(tableName, "BL_LIMITCNT")
	_hCV_VIEW_BLACKLIST2.BLBOOKINGDAY = field.NewInt64(tableName, "BL_BOOKINGDAY")

	_hCV_VIEW_BLACKLIST2.fillFieldMap()

	return _hCV_VIEW_BLACKLIST2
}

type hCV_VIEW_BLACKLIST2 struct {
	hCV_VIEW_BLACKLIST2Do

	ALL          field.Asterisk
	BLSEQNO      field.Int64
	SCHOOLNO     field.String
	USERNAME     field.String
	REASON       field.String
	REGTIME      field.String
	BLOCKSTART   field.String
	BLOCKEXPIRE  field.String
	STATUS       field.Int64
	SETFROM      field.String
	BLDAYBEFORE  field.Int64
	BLLIMITCNT   field.Int64
	BLBOOKINGDAY field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_VIEW_BLACKLIST2) Table(newTableName string) *hCV_VIEW_BLACKLIST2 {
	h.hCV_VIEW_BLACKLIST2Do.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_VIEW_BLACKLIST2) As(alias string) *hCV_VIEW_BLACKLIST2 {
	h.hCV_VIEW_BLACKLIST2Do.DO = *(h.hCV_VIEW_BLACKLIST2Do.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_VIEW_BLACKLIST2) updateTableName(table string) *hCV_VIEW_BLACKLIST2 {
	h.ALL = field.NewAsterisk(table)
	h.BLSEQNO = field.NewInt64(table, "BLSEQNO")
	h.SCHOOLNO = field.NewString(table, "SCHOOLNO")
	h.USERNAME = field.NewString(table, "USER_NAME")
	h.REASON = field.NewString(table, "REASON")
	h.REGTIME = field.NewString(table, "REGTIME")
	h.BLOCKSTART = field.NewString(table, "BLOCKSTART")
	h.BLOCKEXPIRE = field.NewString(table, "BLOCKEXPIRE")
	h.STATUS = field.NewInt64(table, "STATUS")
	h.SETFROM = field.NewString(table, "SETFROM")
	h.BLDAYBEFORE = field.NewInt64(table, "BL_DAYBEFORE")
	h.BLLIMITCNT = field.NewInt64(table, "BL_LIMITCNT")
	h.BLBOOKINGDAY = field.NewInt64(table, "BL_BOOKINGDAY")

	h.fillFieldMap()

	return h
}

func (h *hCV_VIEW_BLACKLIST2) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_VIEW_BLACKLIST2) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 12)
	h.fieldMap["BLSEQNO"] = h.BLSEQNO
	h.fieldMap["SCHOOLNO"] = h.SCHOOLNO
	h.fieldMap["USER_NAME"] = h.USERNAME
	h.fieldMap["REASON"] = h.REASON
	h.fieldMap["REGTIME"] = h.REGTIME
	h.fieldMap["BLOCKSTART"] = h.BLOCKSTART
	h.fieldMap["BLOCKEXPIRE"] = h.BLOCKEXPIRE
	h.fieldMap["STATUS"] = h.STATUS
	h.fieldMap["SETFROM"] = h.SETFROM
	h.fieldMap["BL_DAYBEFORE"] = h.BLDAYBEFORE
	h.fieldMap["BL_LIMITCNT"] = h.BLLIMITCNT
	h.fieldMap["BL_BOOKINGDAY"] = h.BLBOOKINGDAY
}

func (h hCV_VIEW_BLACKLIST2) clone(db *gorm.DB) hCV_VIEW_BLACKLIST2 {
	h.hCV_VIEW_BLACKLIST2Do.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_VIEW_BLACKLIST2) replaceDB(db *gorm.DB) hCV_VIEW_BLACKLIST2 {
	h.hCV_VIEW_BLACKLIST2Do.ReplaceDB(db)
	return h
}

type hCV_VIEW_BLACKLIST2Do struct{ gen.DO }

type IHCV_VIEW_BLACKLIST2Do interface {
	gen.SubQuery
	Debug() IHCV_VIEW_BLACKLIST2Do
	WithContext(ctx context.Context) IHCV_VIEW_BLACKLIST2Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_VIEW_BLACKLIST2Do
	WriteDB() IHCV_VIEW_BLACKLIST2Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_VIEW_BLACKLIST2Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_VIEW_BLACKLIST2Do
	Not(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do
	Or(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do
	Select(conds ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	Where(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do
	Order(conds ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	Distinct(cols ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	Omit(cols ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	Join(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	Group(cols ...field.Expr) IHCV_VIEW_BLACKLIST2Do
	Having(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do
	Limit(limit int) IHCV_VIEW_BLACKLIST2Do
	Offset(offset int) IHCV_VIEW_BLACKLIST2Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_VIEW_BLACKLIST2Do
	Unscoped() IHCV_VIEW_BLACKLIST2Do
	Create(values ...*model.HCV_VIEW_BLACKLIST2) error
	CreateInBatches(values []*model.HCV_VIEW_BLACKLIST2, batchSize int) error
	Save(values ...*model.HCV_VIEW_BLACKLIST2) error
	First() (*model.HCV_VIEW_BLACKLIST2, error)
	Take() (*model.HCV_VIEW_BLACKLIST2, error)
	Last() (*model.HCV_VIEW_BLACKLIST2, error)
	Find() ([]*model.HCV_VIEW_BLACKLIST2, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_VIEW_BLACKLIST2, err error)
	FindInBatches(result *[]*model.HCV_VIEW_BLACKLIST2, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_VIEW_BLACKLIST2) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST2Do
	Assign(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST2Do
	Joins(fields ...field.RelationField) IHCV_VIEW_BLACKLIST2Do
	Preload(fields ...field.RelationField) IHCV_VIEW_BLACKLIST2Do
	FirstOrInit() (*model.HCV_VIEW_BLACKLIST2, error)
	FirstOrCreate() (*model.HCV_VIEW_BLACKLIST2, error)
	FindByPage(offset int, limit int) (result []*model.HCV_VIEW_BLACKLIST2, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_VIEW_BLACKLIST2Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_VIEW_BLACKLIST2Do) Debug() IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Debug())
}

func (h hCV_VIEW_BLACKLIST2Do) WithContext(ctx context.Context) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_VIEW_BLACKLIST2Do) ReadDB() IHCV_VIEW_BLACKLIST2Do {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_VIEW_BLACKLIST2Do) WriteDB() IHCV_VIEW_BLACKLIST2Do {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_VIEW_BLACKLIST2Do) Session(config *gorm.Session) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_VIEW_BLACKLIST2Do) Clauses(conds ...clause.Expression) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Returning(value interface{}, columns ...string) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_VIEW_BLACKLIST2Do) Not(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Or(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Select(conds ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Where(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Order(conds ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Distinct(cols ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_VIEW_BLACKLIST2Do) Omit(cols ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_VIEW_BLACKLIST2Do) Join(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_VIEW_BLACKLIST2Do) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_VIEW_BLACKLIST2Do) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_VIEW_BLACKLIST2Do) Group(cols ...field.Expr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_VIEW_BLACKLIST2Do) Having(conds ...gen.Condition) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_VIEW_BLACKLIST2Do) Limit(limit int) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_VIEW_BLACKLIST2Do) Offset(offset int) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_VIEW_BLACKLIST2Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_VIEW_BLACKLIST2Do) Unscoped() IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_VIEW_BLACKLIST2Do) Create(values ...*model.HCV_VIEW_BLACKLIST2) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_VIEW_BLACKLIST2Do) CreateInBatches(values []*model.HCV_VIEW_BLACKLIST2, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_VIEW_BLACKLIST2Do) Save(values ...*model.HCV_VIEW_BLACKLIST2) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_VIEW_BLACKLIST2Do) First() (*model.HCV_VIEW_BLACKLIST2, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST2), nil
	}
}

func (h hCV_VIEW_BLACKLIST2Do) Take() (*model.HCV_VIEW_BLACKLIST2, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST2), nil
	}
}

func (h hCV_VIEW_BLACKLIST2Do) Last() (*model.HCV_VIEW_BLACKLIST2, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST2), nil
	}
}

func (h hCV_VIEW_BLACKLIST2Do) Find() ([]*model.HCV_VIEW_BLACKLIST2, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_VIEW_BLACKLIST2), err
}

func (h hCV_VIEW_BLACKLIST2Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_VIEW_BLACKLIST2, err error) {
	buf := make([]*model.HCV_VIEW_BLACKLIST2, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_VIEW_BLACKLIST2Do) FindInBatches(result *[]*model.HCV_VIEW_BLACKLIST2, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_VIEW_BLACKLIST2Do) Attrs(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_VIEW_BLACKLIST2Do) Assign(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST2Do {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_VIEW_BLACKLIST2Do) Joins(fields ...field.RelationField) IHCV_VIEW_BLACKLIST2Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_VIEW_BLACKLIST2Do) Preload(fields ...field.RelationField) IHCV_VIEW_BLACKLIST2Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_VIEW_BLACKLIST2Do) FirstOrInit() (*model.HCV_VIEW_BLACKLIST2, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST2), nil
	}
}

func (h hCV_VIEW_BLACKLIST2Do) FirstOrCreate() (*model.HCV_VIEW_BLACKLIST2, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST2), nil
	}
}

func (h hCV_VIEW_BLACKLIST2Do) FindByPage(offset int, limit int) (result []*model.HCV_VIEW_BLACKLIST2, count int64, err error) {
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

func (h hCV_VIEW_BLACKLIST2Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_VIEW_BLACKLIST2Do) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_VIEW_BLACKLIST2Do) Delete(models ...*model.HCV_VIEW_BLACKLIST2) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_VIEW_BLACKLIST2Do) withDO(do gen.Dao) *hCV_VIEW_BLACKLIST2Do {
	h.DO = *do.(*gen.DO)
	return h
}
