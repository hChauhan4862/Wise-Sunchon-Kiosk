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

func newHCV_vw_today_in(db *gorm.DB, opts ...gen.DOOption) hCV_vw_today_in {
	_hCV_vw_today_in := hCV_vw_today_in{}

	_hCV_vw_today_in.hCV_vw_today_inDo.UseDB(db, opts...)
	_hCV_vw_today_in.hCV_vw_today_inDo.UseModel(&model.HCV_vw_today_in{})

	tableName := _hCV_vw_today_in.hCV_vw_today_inDo.TableName()
	_hCV_vw_today_in.ALL = field.NewAsterisk(tableName)
	_hCV_vw_today_in.InCnt = field.NewInt64(tableName, "in_cnt")

	_hCV_vw_today_in.fillFieldMap()

	return _hCV_vw_today_in
}

type hCV_vw_today_in struct {
	hCV_vw_today_inDo

	ALL   field.Asterisk
	InCnt field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_vw_today_in) Table(newTableName string) *hCV_vw_today_in {
	h.hCV_vw_today_inDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_vw_today_in) As(alias string) *hCV_vw_today_in {
	h.hCV_vw_today_inDo.DO = *(h.hCV_vw_today_inDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_vw_today_in) updateTableName(table string) *hCV_vw_today_in {
	h.ALL = field.NewAsterisk(table)
	h.InCnt = field.NewInt64(table, "in_cnt")

	h.fillFieldMap()

	return h
}

func (h *hCV_vw_today_in) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_vw_today_in) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 1)
	h.fieldMap["in_cnt"] = h.InCnt
}

func (h hCV_vw_today_in) clone(db *gorm.DB) hCV_vw_today_in {
	h.hCV_vw_today_inDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_vw_today_in) replaceDB(db *gorm.DB) hCV_vw_today_in {
	h.hCV_vw_today_inDo.ReplaceDB(db)
	return h
}

type hCV_vw_today_inDo struct{ gen.DO }

type IHCV_vw_today_inDo interface {
	gen.SubQuery
	Debug() IHCV_vw_today_inDo
	WithContext(ctx context.Context) IHCV_vw_today_inDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_vw_today_inDo
	WriteDB() IHCV_vw_today_inDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_vw_today_inDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_vw_today_inDo
	Not(conds ...gen.Condition) IHCV_vw_today_inDo
	Or(conds ...gen.Condition) IHCV_vw_today_inDo
	Select(conds ...field.Expr) IHCV_vw_today_inDo
	Where(conds ...gen.Condition) IHCV_vw_today_inDo
	Order(conds ...field.Expr) IHCV_vw_today_inDo
	Distinct(cols ...field.Expr) IHCV_vw_today_inDo
	Omit(cols ...field.Expr) IHCV_vw_today_inDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_vw_today_inDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_today_inDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_today_inDo
	Group(cols ...field.Expr) IHCV_vw_today_inDo
	Having(conds ...gen.Condition) IHCV_vw_today_inDo
	Limit(limit int) IHCV_vw_today_inDo
	Offset(offset int) IHCV_vw_today_inDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_today_inDo
	Unscoped() IHCV_vw_today_inDo
	Create(values ...*model.HCV_vw_today_in) error
	CreateInBatches(values []*model.HCV_vw_today_in, batchSize int) error
	Save(values ...*model.HCV_vw_today_in) error
	First() (*model.HCV_vw_today_in, error)
	Take() (*model.HCV_vw_today_in, error)
	Last() (*model.HCV_vw_today_in, error)
	Find() ([]*model.HCV_vw_today_in, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_today_in, err error)
	FindInBatches(result *[]*model.HCV_vw_today_in, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_vw_today_in) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_vw_today_inDo
	Assign(attrs ...field.AssignExpr) IHCV_vw_today_inDo
	Joins(fields ...field.RelationField) IHCV_vw_today_inDo
	Preload(fields ...field.RelationField) IHCV_vw_today_inDo
	FirstOrInit() (*model.HCV_vw_today_in, error)
	FirstOrCreate() (*model.HCV_vw_today_in, error)
	FindByPage(offset int, limit int) (result []*model.HCV_vw_today_in, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_vw_today_inDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_vw_today_inDo) Debug() IHCV_vw_today_inDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_vw_today_inDo) WithContext(ctx context.Context) IHCV_vw_today_inDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_vw_today_inDo) ReadDB() IHCV_vw_today_inDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_vw_today_inDo) WriteDB() IHCV_vw_today_inDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_vw_today_inDo) Session(config *gorm.Session) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_vw_today_inDo) Clauses(conds ...clause.Expression) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_vw_today_inDo) Returning(value interface{}, columns ...string) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_vw_today_inDo) Not(conds ...gen.Condition) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_vw_today_inDo) Or(conds ...gen.Condition) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_vw_today_inDo) Select(conds ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_vw_today_inDo) Where(conds ...gen.Condition) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_vw_today_inDo) Order(conds ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_vw_today_inDo) Distinct(cols ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_vw_today_inDo) Omit(cols ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_vw_today_inDo) Join(table schema.Tabler, on ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_vw_today_inDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_vw_today_inDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_vw_today_inDo) Group(cols ...field.Expr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_vw_today_inDo) Having(conds ...gen.Condition) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_vw_today_inDo) Limit(limit int) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_vw_today_inDo) Offset(offset int) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_vw_today_inDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_vw_today_inDo) Unscoped() IHCV_vw_today_inDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_vw_today_inDo) Create(values ...*model.HCV_vw_today_in) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_vw_today_inDo) CreateInBatches(values []*model.HCV_vw_today_in, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_vw_today_inDo) Save(values ...*model.HCV_vw_today_in) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_vw_today_inDo) First() (*model.HCV_vw_today_in, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_today_in), nil
	}
}

func (h hCV_vw_today_inDo) Take() (*model.HCV_vw_today_in, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_today_in), nil
	}
}

func (h hCV_vw_today_inDo) Last() (*model.HCV_vw_today_in, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_today_in), nil
	}
}

func (h hCV_vw_today_inDo) Find() ([]*model.HCV_vw_today_in, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_vw_today_in), err
}

func (h hCV_vw_today_inDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_today_in, err error) {
	buf := make([]*model.HCV_vw_today_in, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_vw_today_inDo) FindInBatches(result *[]*model.HCV_vw_today_in, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_vw_today_inDo) Attrs(attrs ...field.AssignExpr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_vw_today_inDo) Assign(attrs ...field.AssignExpr) IHCV_vw_today_inDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_vw_today_inDo) Joins(fields ...field.RelationField) IHCV_vw_today_inDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_vw_today_inDo) Preload(fields ...field.RelationField) IHCV_vw_today_inDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_vw_today_inDo) FirstOrInit() (*model.HCV_vw_today_in, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_today_in), nil
	}
}

func (h hCV_vw_today_inDo) FirstOrCreate() (*model.HCV_vw_today_in, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_today_in), nil
	}
}

func (h hCV_vw_today_inDo) FindByPage(offset int, limit int) (result []*model.HCV_vw_today_in, count int64, err error) {
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

func (h hCV_vw_today_inDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_vw_today_inDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_vw_today_inDo) Delete(models ...*model.HCV_vw_today_in) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_vw_today_inDo) withDO(do gen.Dao) *hCV_vw_today_inDo {
	h.DO = *do.(*gen.DO)
	return h
}
