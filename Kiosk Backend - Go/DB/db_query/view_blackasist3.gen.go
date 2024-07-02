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

func newHCV_VIEW_BLACKASIST3(db *gorm.DB, opts ...gen.DOOption) hCV_VIEW_BLACKASIST3 {
	_hCV_VIEW_BLACKASIST3 := hCV_VIEW_BLACKASIST3{}

	_hCV_VIEW_BLACKASIST3.hCV_VIEW_BLACKASIST3Do.UseDB(db, opts...)
	_hCV_VIEW_BLACKASIST3.hCV_VIEW_BLACKASIST3Do.UseModel(&model.HCV_VIEW_BLACKASIST3{})

	tableName := _hCV_VIEW_BLACKASIST3.hCV_VIEW_BLACKASIST3Do.TableName()
	_hCV_VIEW_BLACKASIST3.ALL = field.NewAsterisk(tableName)
	_hCV_VIEW_BLACKASIST3.ALTPID = field.NewString(tableName, "ALT_PID")
	_hCV_VIEW_BLACKASIST3.LASTREGTIME = field.NewTime(tableName, "LAST_REGTIME")
	_hCV_VIEW_BLACKASIST3.MISSCNT = field.NewInt64(tableName, "MISS_CNT")

	_hCV_VIEW_BLACKASIST3.fillFieldMap()

	return _hCV_VIEW_BLACKASIST3
}

type hCV_VIEW_BLACKASIST3 struct {
	hCV_VIEW_BLACKASIST3Do

	ALL         field.Asterisk
	ALTPID      field.String
	LASTREGTIME field.Time
	MISSCNT     field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_VIEW_BLACKASIST3) Table(newTableName string) *hCV_VIEW_BLACKASIST3 {
	h.hCV_VIEW_BLACKASIST3Do.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_VIEW_BLACKASIST3) As(alias string) *hCV_VIEW_BLACKASIST3 {
	h.hCV_VIEW_BLACKASIST3Do.DO = *(h.hCV_VIEW_BLACKASIST3Do.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_VIEW_BLACKASIST3) updateTableName(table string) *hCV_VIEW_BLACKASIST3 {
	h.ALL = field.NewAsterisk(table)
	h.ALTPID = field.NewString(table, "ALT_PID")
	h.LASTREGTIME = field.NewTime(table, "LAST_REGTIME")
	h.MISSCNT = field.NewInt64(table, "MISS_CNT")

	h.fillFieldMap()

	return h
}

func (h *hCV_VIEW_BLACKASIST3) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_VIEW_BLACKASIST3) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 3)
	h.fieldMap["ALT_PID"] = h.ALTPID
	h.fieldMap["LAST_REGTIME"] = h.LASTREGTIME
	h.fieldMap["MISS_CNT"] = h.MISSCNT
}

func (h hCV_VIEW_BLACKASIST3) clone(db *gorm.DB) hCV_VIEW_BLACKASIST3 {
	h.hCV_VIEW_BLACKASIST3Do.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_VIEW_BLACKASIST3) replaceDB(db *gorm.DB) hCV_VIEW_BLACKASIST3 {
	h.hCV_VIEW_BLACKASIST3Do.ReplaceDB(db)
	return h
}

type hCV_VIEW_BLACKASIST3Do struct{ gen.DO }

type IHCV_VIEW_BLACKASIST3Do interface {
	gen.SubQuery
	Debug() IHCV_VIEW_BLACKASIST3Do
	WithContext(ctx context.Context) IHCV_VIEW_BLACKASIST3Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_VIEW_BLACKASIST3Do
	WriteDB() IHCV_VIEW_BLACKASIST3Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_VIEW_BLACKASIST3Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_VIEW_BLACKASIST3Do
	Not(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do
	Or(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do
	Select(conds ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	Where(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do
	Order(conds ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	Distinct(cols ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	Omit(cols ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	Join(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	Group(cols ...field.Expr) IHCV_VIEW_BLACKASIST3Do
	Having(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do
	Limit(limit int) IHCV_VIEW_BLACKASIST3Do
	Offset(offset int) IHCV_VIEW_BLACKASIST3Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_VIEW_BLACKASIST3Do
	Unscoped() IHCV_VIEW_BLACKASIST3Do
	Create(values ...*model.HCV_VIEW_BLACKASIST3) error
	CreateInBatches(values []*model.HCV_VIEW_BLACKASIST3, batchSize int) error
	Save(values ...*model.HCV_VIEW_BLACKASIST3) error
	First() (*model.HCV_VIEW_BLACKASIST3, error)
	Take() (*model.HCV_VIEW_BLACKASIST3, error)
	Last() (*model.HCV_VIEW_BLACKASIST3, error)
	Find() ([]*model.HCV_VIEW_BLACKASIST3, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_VIEW_BLACKASIST3, err error)
	FindInBatches(result *[]*model.HCV_VIEW_BLACKASIST3, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_VIEW_BLACKASIST3) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_VIEW_BLACKASIST3Do
	Assign(attrs ...field.AssignExpr) IHCV_VIEW_BLACKASIST3Do
	Joins(fields ...field.RelationField) IHCV_VIEW_BLACKASIST3Do
	Preload(fields ...field.RelationField) IHCV_VIEW_BLACKASIST3Do
	FirstOrInit() (*model.HCV_VIEW_BLACKASIST3, error)
	FirstOrCreate() (*model.HCV_VIEW_BLACKASIST3, error)
	FindByPage(offset int, limit int) (result []*model.HCV_VIEW_BLACKASIST3, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_VIEW_BLACKASIST3Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_VIEW_BLACKASIST3Do) Debug() IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Debug())
}

func (h hCV_VIEW_BLACKASIST3Do) WithContext(ctx context.Context) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_VIEW_BLACKASIST3Do) ReadDB() IHCV_VIEW_BLACKASIST3Do {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_VIEW_BLACKASIST3Do) WriteDB() IHCV_VIEW_BLACKASIST3Do {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_VIEW_BLACKASIST3Do) Session(config *gorm.Session) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_VIEW_BLACKASIST3Do) Clauses(conds ...clause.Expression) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Returning(value interface{}, columns ...string) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_VIEW_BLACKASIST3Do) Not(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Or(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Select(conds ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Where(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Order(conds ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Distinct(cols ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_VIEW_BLACKASIST3Do) Omit(cols ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_VIEW_BLACKASIST3Do) Join(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_VIEW_BLACKASIST3Do) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_VIEW_BLACKASIST3Do) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_VIEW_BLACKASIST3Do) Group(cols ...field.Expr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_VIEW_BLACKASIST3Do) Having(conds ...gen.Condition) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_VIEW_BLACKASIST3Do) Limit(limit int) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_VIEW_BLACKASIST3Do) Offset(offset int) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_VIEW_BLACKASIST3Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_VIEW_BLACKASIST3Do) Unscoped() IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_VIEW_BLACKASIST3Do) Create(values ...*model.HCV_VIEW_BLACKASIST3) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_VIEW_BLACKASIST3Do) CreateInBatches(values []*model.HCV_VIEW_BLACKASIST3, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_VIEW_BLACKASIST3Do) Save(values ...*model.HCV_VIEW_BLACKASIST3) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_VIEW_BLACKASIST3Do) First() (*model.HCV_VIEW_BLACKASIST3, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKASIST3), nil
	}
}

func (h hCV_VIEW_BLACKASIST3Do) Take() (*model.HCV_VIEW_BLACKASIST3, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKASIST3), nil
	}
}

func (h hCV_VIEW_BLACKASIST3Do) Last() (*model.HCV_VIEW_BLACKASIST3, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKASIST3), nil
	}
}

func (h hCV_VIEW_BLACKASIST3Do) Find() ([]*model.HCV_VIEW_BLACKASIST3, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_VIEW_BLACKASIST3), err
}

func (h hCV_VIEW_BLACKASIST3Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_VIEW_BLACKASIST3, err error) {
	buf := make([]*model.HCV_VIEW_BLACKASIST3, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_VIEW_BLACKASIST3Do) FindInBatches(result *[]*model.HCV_VIEW_BLACKASIST3, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_VIEW_BLACKASIST3Do) Attrs(attrs ...field.AssignExpr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_VIEW_BLACKASIST3Do) Assign(attrs ...field.AssignExpr) IHCV_VIEW_BLACKASIST3Do {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_VIEW_BLACKASIST3Do) Joins(fields ...field.RelationField) IHCV_VIEW_BLACKASIST3Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_VIEW_BLACKASIST3Do) Preload(fields ...field.RelationField) IHCV_VIEW_BLACKASIST3Do {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_VIEW_BLACKASIST3Do) FirstOrInit() (*model.HCV_VIEW_BLACKASIST3, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKASIST3), nil
	}
}

func (h hCV_VIEW_BLACKASIST3Do) FirstOrCreate() (*model.HCV_VIEW_BLACKASIST3, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_VIEW_BLACKASIST3), nil
	}
}

func (h hCV_VIEW_BLACKASIST3Do) FindByPage(offset int, limit int) (result []*model.HCV_VIEW_BLACKASIST3, count int64, err error) {
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

func (h hCV_VIEW_BLACKASIST3Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_VIEW_BLACKASIST3Do) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_VIEW_BLACKASIST3Do) Delete(models ...*model.HCV_VIEW_BLACKASIST3) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_VIEW_BLACKASIST3Do) withDO(do gen.Dao) *hCV_VIEW_BLACKASIST3Do {
	h.DO = *do.(*gen.DO)
	return h
}