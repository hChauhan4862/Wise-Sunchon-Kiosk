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

func newHCV_view_opentype_rooms(db *gorm.DB, opts ...gen.DOOption) hCV_view_opentype_rooms {
	_hCV_view_opentype_rooms := hCV_view_opentype_rooms{}

	_hCV_view_opentype_rooms.hCV_view_opentype_roomsDo.UseDB(db, opts...)
	_hCV_view_opentype_rooms.hCV_view_opentype_roomsDo.UseModel(&model.HCV_view_opentype_rooms{})

	tableName := _hCV_view_opentype_rooms.hCV_view_opentype_roomsDo.TableName()
	_hCV_view_opentype_rooms.ALL = field.NewAsterisk(tableName)
	_hCV_view_opentype_rooms.OpenType = field.NewInt64(tableName, "open_type")
	_hCV_view_opentype_rooms.Roomnames = field.NewString(tableName, "roomnames")

	_hCV_view_opentype_rooms.fillFieldMap()

	return _hCV_view_opentype_rooms
}

type hCV_view_opentype_rooms struct {
	hCV_view_opentype_roomsDo

	ALL       field.Asterisk
	OpenType  field.Int64
	Roomnames field.String

	fieldMap map[string]field.Expr
}

func (h hCV_view_opentype_rooms) Table(newTableName string) *hCV_view_opentype_rooms {
	h.hCV_view_opentype_roomsDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_view_opentype_rooms) As(alias string) *hCV_view_opentype_rooms {
	h.hCV_view_opentype_roomsDo.DO = *(h.hCV_view_opentype_roomsDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_view_opentype_rooms) updateTableName(table string) *hCV_view_opentype_rooms {
	h.ALL = field.NewAsterisk(table)
	h.OpenType = field.NewInt64(table, "open_type")
	h.Roomnames = field.NewString(table, "roomnames")

	h.fillFieldMap()

	return h
}

func (h *hCV_view_opentype_rooms) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_view_opentype_rooms) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 2)
	h.fieldMap["open_type"] = h.OpenType
	h.fieldMap["roomnames"] = h.Roomnames
}

func (h hCV_view_opentype_rooms) clone(db *gorm.DB) hCV_view_opentype_rooms {
	h.hCV_view_opentype_roomsDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_view_opentype_rooms) replaceDB(db *gorm.DB) hCV_view_opentype_rooms {
	h.hCV_view_opentype_roomsDo.ReplaceDB(db)
	return h
}

type hCV_view_opentype_roomsDo struct{ gen.DO }

type IHCV_view_opentype_roomsDo interface {
	gen.SubQuery
	Debug() IHCV_view_opentype_roomsDo
	WithContext(ctx context.Context) IHCV_view_opentype_roomsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_view_opentype_roomsDo
	WriteDB() IHCV_view_opentype_roomsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_view_opentype_roomsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_view_opentype_roomsDo
	Not(conds ...gen.Condition) IHCV_view_opentype_roomsDo
	Or(conds ...gen.Condition) IHCV_view_opentype_roomsDo
	Select(conds ...field.Expr) IHCV_view_opentype_roomsDo
	Where(conds ...gen.Condition) IHCV_view_opentype_roomsDo
	Order(conds ...field.Expr) IHCV_view_opentype_roomsDo
	Distinct(cols ...field.Expr) IHCV_view_opentype_roomsDo
	Omit(cols ...field.Expr) IHCV_view_opentype_roomsDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_view_opentype_roomsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_opentype_roomsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_opentype_roomsDo
	Group(cols ...field.Expr) IHCV_view_opentype_roomsDo
	Having(conds ...gen.Condition) IHCV_view_opentype_roomsDo
	Limit(limit int) IHCV_view_opentype_roomsDo
	Offset(offset int) IHCV_view_opentype_roomsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_opentype_roomsDo
	Unscoped() IHCV_view_opentype_roomsDo
	Create(values ...*model.HCV_view_opentype_rooms) error
	CreateInBatches(values []*model.HCV_view_opentype_rooms, batchSize int) error
	Save(values ...*model.HCV_view_opentype_rooms) error
	First() (*model.HCV_view_opentype_rooms, error)
	Take() (*model.HCV_view_opentype_rooms, error)
	Last() (*model.HCV_view_opentype_rooms, error)
	Find() ([]*model.HCV_view_opentype_rooms, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_opentype_rooms, err error)
	FindInBatches(result *[]*model.HCV_view_opentype_rooms, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_view_opentype_rooms) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_view_opentype_roomsDo
	Assign(attrs ...field.AssignExpr) IHCV_view_opentype_roomsDo
	Joins(fields ...field.RelationField) IHCV_view_opentype_roomsDo
	Preload(fields ...field.RelationField) IHCV_view_opentype_roomsDo
	FirstOrInit() (*model.HCV_view_opentype_rooms, error)
	FirstOrCreate() (*model.HCV_view_opentype_rooms, error)
	FindByPage(offset int, limit int) (result []*model.HCV_view_opentype_rooms, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_view_opentype_roomsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_view_opentype_roomsDo) Debug() IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_view_opentype_roomsDo) WithContext(ctx context.Context) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_view_opentype_roomsDo) ReadDB() IHCV_view_opentype_roomsDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_view_opentype_roomsDo) WriteDB() IHCV_view_opentype_roomsDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_view_opentype_roomsDo) Session(config *gorm.Session) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_view_opentype_roomsDo) Clauses(conds ...clause.Expression) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_view_opentype_roomsDo) Returning(value interface{}, columns ...string) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_view_opentype_roomsDo) Not(conds ...gen.Condition) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_view_opentype_roomsDo) Or(conds ...gen.Condition) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_view_opentype_roomsDo) Select(conds ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_view_opentype_roomsDo) Where(conds ...gen.Condition) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_view_opentype_roomsDo) Order(conds ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_view_opentype_roomsDo) Distinct(cols ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_view_opentype_roomsDo) Omit(cols ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_view_opentype_roomsDo) Join(table schema.Tabler, on ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_view_opentype_roomsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_view_opentype_roomsDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_view_opentype_roomsDo) Group(cols ...field.Expr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_view_opentype_roomsDo) Having(conds ...gen.Condition) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_view_opentype_roomsDo) Limit(limit int) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_view_opentype_roomsDo) Offset(offset int) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_view_opentype_roomsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_view_opentype_roomsDo) Unscoped() IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_view_opentype_roomsDo) Create(values ...*model.HCV_view_opentype_rooms) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_view_opentype_roomsDo) CreateInBatches(values []*model.HCV_view_opentype_rooms, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_view_opentype_roomsDo) Save(values ...*model.HCV_view_opentype_rooms) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_view_opentype_roomsDo) First() (*model.HCV_view_opentype_rooms, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_opentype_rooms), nil
	}
}

func (h hCV_view_opentype_roomsDo) Take() (*model.HCV_view_opentype_rooms, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_opentype_rooms), nil
	}
}

func (h hCV_view_opentype_roomsDo) Last() (*model.HCV_view_opentype_rooms, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_opentype_rooms), nil
	}
}

func (h hCV_view_opentype_roomsDo) Find() ([]*model.HCV_view_opentype_rooms, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_view_opentype_rooms), err
}

func (h hCV_view_opentype_roomsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_opentype_rooms, err error) {
	buf := make([]*model.HCV_view_opentype_rooms, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_view_opentype_roomsDo) FindInBatches(result *[]*model.HCV_view_opentype_rooms, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_view_opentype_roomsDo) Attrs(attrs ...field.AssignExpr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_view_opentype_roomsDo) Assign(attrs ...field.AssignExpr) IHCV_view_opentype_roomsDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_view_opentype_roomsDo) Joins(fields ...field.RelationField) IHCV_view_opentype_roomsDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_view_opentype_roomsDo) Preload(fields ...field.RelationField) IHCV_view_opentype_roomsDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_view_opentype_roomsDo) FirstOrInit() (*model.HCV_view_opentype_rooms, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_opentype_rooms), nil
	}
}

func (h hCV_view_opentype_roomsDo) FirstOrCreate() (*model.HCV_view_opentype_rooms, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_opentype_rooms), nil
	}
}

func (h hCV_view_opentype_roomsDo) FindByPage(offset int, limit int) (result []*model.HCV_view_opentype_rooms, count int64, err error) {
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

func (h hCV_view_opentype_roomsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_view_opentype_roomsDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_view_opentype_roomsDo) Delete(models ...*model.HCV_view_opentype_rooms) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_view_opentype_roomsDo) withDO(do gen.Dao) *hCV_view_opentype_roomsDo {
	h.DO = *do.(*gen.DO)
	return h
}
