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

func newHCV_VIEW_BLACKLIST3(db *gorm.DB, opts ...gen.DOOption) hCV_VIEW_BLACKLIST3 {
	_hCV_VIEW_BLACKLIST3 := hCV_VIEW_BLACKLIST3{}

	_hCV_VIEW_BLACKLIST3.hCV_VIEW_BLACKLIST3Do.UseDB(db, opts...)
	_hCV_VIEW_BLACKLIST3.hCV_VIEW_BLACKLIST3Do.UseModel(&model.HCV_VIEW_BLACKLIST3{})

	tableName := _hCV_VIEW_BLACKLIST3.hCV_VIEW_BLACKLIST3Do.TableName()
	_hCV_VIEW_BLACKLIST3.ALL = field.NewAsterisk(tableName)
	_hCV_VIEW_BLACKLIST3.SCHOOLNO = field.NewString(tableName, "SCHOOLNO")
	_hCV_VIEW_BLACKLIST3.USERNAME = field.NewString(tableName, "USER_NAME")
	_hCV_VIEW_BLACKLIST3.REASON = field.NewString(tableName, "REASON")
	_hCV_VIEW_BLACKLIST3.REGTIME = field.NewString(tableName, "REGTIME")
	_hCV_VIEW_BLACKLIST3.BLOCKSTART = field.NewTime(tableName, "BLOCKSTART")
	_hCV_VIEW_BLACKLIST3.BLOCKEXPIRE = field.NewTime(tableName, "BLOCKEXPIRE")
	_hCV_VIEW_BLACKLIST3.STATUS = field.NewInt64(tableName, "STATUS")
	_hCV_VIEW_BLACKLIST3.SETFROM = field.NewString(tableName, "SETFROM")
	_hCV_VIEW_BLACKLIST3.BLDAYBEFORE = field.NewInt64(tableName, "BL_DAYBEFORE")
	_hCV_VIEW_BLACKLIST3.BLLIMITCNT = field.NewInt64(tableName, "BL_LIMITCNT")
	_hCV_VIEW_BLACKLIST3.BLBOOKINGDAY = field.NewInt64(tableName, "BL_BOOKINGDAY")

	_hCV_VIEW_BLACKLIST3.fillFieldMap()

	return _hCV_VIEW_BLACKLIST3
}

type hCV_VIEW_BLACKLIST3 struct {
	hCV_VIEW_BLACKLIST3Do

	ALL          field.Asterisk
	SCHOOLNO     field.String
	USERNAME     field.String
	REASON       field.String
	REGTIME      field.String
	BLOCKSTART   field.Time
	BLOCKEXPIRE  field.Time
	STATUS       field.Int64
	SETFROM      field.String
	BLDAYBEFORE  field.Int64
	BLLIMITCNT   field.Int64
	BLBOOKINGDAY field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_VIEW_BLACKLIST3) Table(newTableName string) *hCV_VIEW_BLACKLIST3 {
	h.hCV_VIEW_BLACKLIST3Do.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_VIEW_BLACKLIST3) As(alias string) *hCV_VIEW_BLACKLIST3 {
	h.hCV_VIEW_BLACKLIST3Do.DO = *(h.hCV_VIEW_BLACKLIST3Do.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_VIEW_BLACKLIST3) updateTableName(table string) *hCV_VIEW_BLACKLIST3 {
	h.ALL = field.NewAsterisk(table)
	h.SCHOOLNO = field.NewString(table, "SCHOOLNO")
	h.USERNAME = field.NewString(table, "USER_NAME")
	h.REASON = field.NewString(table, "REASON")
	h.REGTIME = field.NewString(table, "REGTIME")
	h.BLOCKSTART = field.NewTime(table, "BLOCKSTART")
	h.BLOCKEXPIRE = field.NewTime(table, "BLOCKEXPIRE")
	h.STATUS = field.NewInt64(table, "STATUS")
	h.SETFROM = field.NewString(table, "SETFROM")
	h.BLDAYBEFORE = field.NewInt64(table, "BL_DAYBEFORE")
	h.BLLIMITCNT = field.NewInt64(table, "BL_LIMITCNT")
	h.BLBOOKINGDAY = field.NewInt64(table, "BL_BOOKINGDAY")

	h.fillFieldMap()

	return h
}

func (h *hCV_VIEW_BLACKLIST3) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_VIEW_BLACKLIST3) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 11)
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

func (h hCV_VIEW_BLACKLIST3) clone(db *gorm.DB) hCV_VIEW_BLACKLIST3 {
	h.hCV_VIEW_BLACKLIST3Do.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_VIEW_BLACKLIST3) replaceDB(db *gorm.DB) hCV_VIEW_BLACKLIST3 {
	h.hCV_VIEW_BLACKLIST3Do.ReplaceDB(db)
	return h
}

type hCV_VIEW_BLACKLIST3Do struct{ gen.DO }

type IHCV_VIEW_BLACKLIST3Do interface {
	gen.SubQuery
	Debug() IHCV_VIEW_BLACKLIST3Do
	WithContext(ctx context.Context) IHCV_VIEW_BLACKLIST3Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_VIEW_BLACKLIST3Do
	WriteDB() IHCV_VIEW_BLACKLIST3Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_VIEW_BLACKLIST3Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_VIEW_BLACKLIST3Do
	Not(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do
	Or(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do
	Select(conds ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	Where(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do
	Order(conds ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	Distinct(cols ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	Omit(cols ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	Join(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	Group(cols ...field.Expr) IHCV_VIEW_BLACKLIST3Do
	Having(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do
	Limit(limit int) IHCV_VIEW_BLACKLIST3Do
	Offset(offset int) IHCV_VIEW_BLACKLIST3Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_VIEW_BLACKLIST3Do
	Unscoped() IHCV_VIEW_BLACKLIST3Do
	Create(values ...*model.HCV_VIEW_BLACKLIST3) error
	CreateInBatches(values []*model.HCV_VIEW_BLACKLIST3, batchSize int) error
	Save(values ...*model.HCV_VIEW_BLACKLIST3) error
	First() (*model.HCV_VIEW_BLACKLIST3, error)
	Take() (*model.HCV_VIEW_BLACKLIST3, error)
	Last() (*model.HCV_VIEW_BLACKLIST3, error)
	Find() ([]*model.HCV_VIEW_BLACKLIST3, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_VIEW_BLACKLIST3, err error)
	FindInBatches(result *[]*model.HCV_VIEW_BLACKLIST3, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_VIEW_BLACKLIST3) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST3Do
	Assign(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST3Do
	Joins(fields ...field.RelationField) IHCV_VIEW_BLACKLIST3Do
	Preload(fields ...field.RelationField) IHCV_VIEW_BLACKLIST3Do
	FirstOrInit() (*model.HCV_VIEW_BLACKLIST3, error)
	FirstOrCreate() (*model.HCV_VIEW_BLACKLIST3, error)
	FindByPage(offset int, limit int) (result []*model.HCV_VIEW_BLACKLIST3, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_VIEW_BLACKLIST3Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_VIEW_BLACKLIST3Do) Debug() IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Debug())
}

func (h hCV_VIEW_BLACKLIST3Do) WithContext(ctx context.Context) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_VIEW_BLACKLIST3Do) ReadDB() IHCV_VIEW_BLACKLIST3Do {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_VIEW_BLACKLIST3Do) WriteDB() IHCV_VIEW_BLACKLIST3Do {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_VIEW_BLACKLIST3Do) Session(config *gorm.Session) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_VIEW_BLACKLIST3Do) Clauses(conds ...clause.Expression) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Returning(value interface{}, columns ...string) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_VIEW_BLACKLIST3Do) Not(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Or(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Select(conds ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Where(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Order(conds ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Distinct(cols ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_VIEW_BLACKLIST3Do) Omit(cols ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_VIEW_BLACKLIST3Do) Join(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_VIEW_BLACKLIST3Do) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_VIEW_BLACKLIST3Do) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_VIEW_BLACKLIST3Do) Group(cols ...field.Expr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_VIEW_BLACKLIST3Do) Having(conds ...gen.Condition) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_VIEW_BLACKLIST3Do) Limit(limit int) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_VIEW_BLACKLIST3Do) Offset(offset int) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_VIEW_BLACKLIST3Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_VIEW_BLACKLIST3Do) Unscoped() IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_VIEW_BLACKLIST3Do) Create(values ...*model.HCV_VIEW_BLACKLIST3) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_VIEW_BLACKLIST3Do) CreateInBatches(values []*model.HCV_VIEW_BLACKLIST3, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_VIEW_BLACKLIST3Do) Save(values ...*model.HCV_VIEW_BLACKLIST3) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_VIEW_BLACKLIST3Do) First() (*model.HCV_VIEW_BLACKLIST3, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST3), nil
	}
}

func (h hCV_VIEW_BLACKLIST3Do) Take() (*model.HCV_VIEW_BLACKLIST3, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST3), nil
	}
}

func (h hCV_VIEW_BLACKLIST3Do) Last() (*model.HCV_VIEW_BLACKLIST3, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST3), nil
	}
}

func (h hCV_VIEW_BLACKLIST3Do) Find() ([]*model.HCV_VIEW_BLACKLIST3, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_VIEW_BLACKLIST3), err
}

func (h hCV_VIEW_BLACKLIST3Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_VIEW_BLACKLIST3, err error) {
	buf := make([]*model.HCV_VIEW_BLACKLIST3, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_VIEW_BLACKLIST3Do) FindInBatches(result *[]*model.HCV_VIEW_BLACKLIST3, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_VIEW_BLACKLIST3Do) Attrs(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_VIEW_BLACKLIST3Do) Assign(attrs ...field.AssignExpr) IHCV_VIEW_BLACKLIST3Do {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_VIEW_BLACKLIST3Do) Joins(fields ...field.RelationField) IHCV_VIEW_BLACKLIST3Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_VIEW_BLACKLIST3Do) Preload(fields ...field.RelationField) IHCV_VIEW_BLACKLIST3Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_VIEW_BLACKLIST3Do) FirstOrInit() (*model.HCV_VIEW_BLACKLIST3, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST3), nil
	}
}

func (h hCV_VIEW_BLACKLIST3Do) FirstOrCreate() (*model.HCV_VIEW_BLACKLIST3, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKLIST3), nil
	}
}

func (h hCV_VIEW_BLACKLIST3Do) FindByPage(offset int, limit int) (result []*model.HCV_VIEW_BLACKLIST3, count int64, err error) {
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

func (h hCV_VIEW_BLACKLIST3Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_VIEW_BLACKLIST3Do) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_VIEW_BLACKLIST3Do) Delete(models ...*model.HCV_VIEW_BLACKLIST3) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_VIEW_BLACKLIST3Do) withDO(do gen.Dao) *hCV_VIEW_BLACKLIST3Do {
	h.DO = *do.(*gen.DO)
	return h
}
