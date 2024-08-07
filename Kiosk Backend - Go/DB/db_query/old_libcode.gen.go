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

func newHCV_old_libcode(db *gorm.DB, opts ...gen.DOOption) hCV_old_libcode {
	_hCV_old_libcode := hCV_old_libcode{}

	_hCV_old_libcode.hCV_old_libcodeDo.UseDB(db, opts...)
	_hCV_old_libcode.hCV_old_libcodeDo.UseModel(&model.HCV_old_libcode{})

	tableName := _hCV_old_libcode.hCV_old_libcodeDo.TableName()
	_hCV_old_libcode.ALL = field.NewAsterisk(tableName)
	_hCV_old_libcode.Libno = field.NewInt64(tableName, "libno")
	_hCV_old_libcode.Name = field.NewString(tableName, "name")
	_hCV_old_libcode.EnName = field.NewString(tableName, "en_name")

	_hCV_old_libcode.fillFieldMap()

	return _hCV_old_libcode
}

type hCV_old_libcode struct {
	hCV_old_libcodeDo

	ALL    field.Asterisk
	Libno  field.Int64
	Name   field.String
	EnName field.String

	fieldMap map[string]field.Expr
}

func (h hCV_old_libcode) Table(newTableName string) *hCV_old_libcode {
	h.hCV_old_libcodeDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_old_libcode) As(alias string) *hCV_old_libcode {
	h.hCV_old_libcodeDo.DO = *(h.hCV_old_libcodeDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_old_libcode) updateTableName(table string) *hCV_old_libcode {
	h.ALL = field.NewAsterisk(table)
	h.Libno = field.NewInt64(table, "libno")
	h.Name = field.NewString(table, "name")
	h.EnName = field.NewString(table, "en_name")

	h.fillFieldMap()

	return h
}

func (h *hCV_old_libcode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_old_libcode) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 3)
	h.fieldMap["libno"] = h.Libno
	h.fieldMap["name"] = h.Name
	h.fieldMap["en_name"] = h.EnName
}

func (h hCV_old_libcode) clone(db *gorm.DB) hCV_old_libcode {
	h.hCV_old_libcodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_old_libcode) replaceDB(db *gorm.DB) hCV_old_libcode {
	h.hCV_old_libcodeDo.ReplaceDB(db)
	return h
}

type hCV_old_libcodeDo struct{ gen.DO }

type IHCV_old_libcodeDo interface {
	gen.SubQuery
	Debug() IHCV_old_libcodeDo
	WithContext(ctx context.Context) IHCV_old_libcodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_old_libcodeDo
	WriteDB() IHCV_old_libcodeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_old_libcodeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_old_libcodeDo
	Not(conds ...gen.Condition) IHCV_old_libcodeDo
	Or(conds ...gen.Condition) IHCV_old_libcodeDo
	Select(conds ...field.Expr) IHCV_old_libcodeDo
	Where(conds ...gen.Condition) IHCV_old_libcodeDo
	Order(conds ...field.Expr) IHCV_old_libcodeDo
	Distinct(cols ...field.Expr) IHCV_old_libcodeDo
	Omit(cols ...field.Expr) IHCV_old_libcodeDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_old_libcodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_old_libcodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_old_libcodeDo
	Group(cols ...field.Expr) IHCV_old_libcodeDo
	Having(conds ...gen.Condition) IHCV_old_libcodeDo
	Limit(limit int) IHCV_old_libcodeDo
	Offset(offset int) IHCV_old_libcodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_old_libcodeDo
	Unscoped() IHCV_old_libcodeDo
	Create(values ...*model.HCV_old_libcode) error
	CreateInBatches(values []*model.HCV_old_libcode, batchSize int) error
	Save(values ...*model.HCV_old_libcode) error
	First() (*model.HCV_old_libcode, error)
	Take() (*model.HCV_old_libcode, error)
	Last() (*model.HCV_old_libcode, error)
	Find() ([]*model.HCV_old_libcode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_old_libcode, err error)
	FindInBatches(result *[]*model.HCV_old_libcode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_old_libcode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_old_libcodeDo
	Assign(attrs ...field.AssignExpr) IHCV_old_libcodeDo
	Joins(fields ...field.RelationField) IHCV_old_libcodeDo
	Preload(fields ...field.RelationField) IHCV_old_libcodeDo
	FirstOrInit() (*model.HCV_old_libcode, error)
	FirstOrCreate() (*model.HCV_old_libcode, error)
	FindByPage(offset int, limit int) (result []*model.HCV_old_libcode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_old_libcodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_old_libcodeDo) Debug() IHCV_old_libcodeDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_old_libcodeDo) WithContext(ctx context.Context) IHCV_old_libcodeDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_old_libcodeDo) ReadDB() IHCV_old_libcodeDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_old_libcodeDo) WriteDB() IHCV_old_libcodeDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_old_libcodeDo) Session(config *gorm.Session) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_old_libcodeDo) Clauses(conds ...clause.Expression) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_old_libcodeDo) Returning(value interface{}, columns ...string) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_old_libcodeDo) Not(conds ...gen.Condition) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_old_libcodeDo) Or(conds ...gen.Condition) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_old_libcodeDo) Select(conds ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_old_libcodeDo) Where(conds ...gen.Condition) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_old_libcodeDo) Order(conds ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_old_libcodeDo) Distinct(cols ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_old_libcodeDo) Omit(cols ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_old_libcodeDo) Join(table schema.Tabler, on ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_old_libcodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_old_libcodeDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_old_libcodeDo) Group(cols ...field.Expr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_old_libcodeDo) Having(conds ...gen.Condition) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_old_libcodeDo) Limit(limit int) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_old_libcodeDo) Offset(offset int) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_old_libcodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_old_libcodeDo) Unscoped() IHCV_old_libcodeDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_old_libcodeDo) Create(values ...*model.HCV_old_libcode) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_old_libcodeDo) CreateInBatches(values []*model.HCV_old_libcode, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_old_libcodeDo) Save(values ...*model.HCV_old_libcode) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_old_libcodeDo) First() (*model.HCV_old_libcode, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_old_libcode), nil
	}
}

func (h hCV_old_libcodeDo) Take() (*model.HCV_old_libcode, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_old_libcode), nil
	}
}

func (h hCV_old_libcodeDo) Last() (*model.HCV_old_libcode, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_old_libcode), nil
	}
}

func (h hCV_old_libcodeDo) Find() ([]*model.HCV_old_libcode, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_old_libcode), err
}

func (h hCV_old_libcodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_old_libcode, err error) {
	buf := make([]*model.HCV_old_libcode, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_old_libcodeDo) FindInBatches(result *[]*model.HCV_old_libcode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_old_libcodeDo) Attrs(attrs ...field.AssignExpr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_old_libcodeDo) Assign(attrs ...field.AssignExpr) IHCV_old_libcodeDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_old_libcodeDo) Joins(fields ...field.RelationField) IHCV_old_libcodeDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_old_libcodeDo) Preload(fields ...field.RelationField) IHCV_old_libcodeDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_old_libcodeDo) FirstOrInit() (*model.HCV_old_libcode, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_old_libcode), nil
	}
}

func (h hCV_old_libcodeDo) FirstOrCreate() (*model.HCV_old_libcode, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_old_libcode), nil
	}
}

func (h hCV_old_libcodeDo) FindByPage(offset int, limit int) (result []*model.HCV_old_libcode, count int64, err error) {
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

func (h hCV_old_libcodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_old_libcodeDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_old_libcodeDo) Delete(models ...*model.HCV_old_libcode) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_old_libcodeDo) withDO(do gen.Dao) *hCV_old_libcodeDo {
	h.DO = *do.(*gen.DO)
	return h
}
