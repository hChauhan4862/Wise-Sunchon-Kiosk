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

func newHCV_vw_busTest(db *gorm.DB, opts ...gen.DOOption) hCV_vw_busTest {
	_hCV_vw_busTest := hCV_vw_busTest{}

	_hCV_vw_busTest.hCV_vw_busTestDo.UseDB(db, opts...)
	_hCV_vw_busTest.hCV_vw_busTestDo.UseModel(&model.HCV_vw_busTest{})

	tableName := _hCV_vw_busTest.hCV_vw_busTestDo.TableName()
	_hCV_vw_busTest.ALL = field.NewAsterisk(tableName)
	_hCV_vw_busTest.BusStopID = field.NewInt64(tableName, "busStopID")
	_hCV_vw_busTest.StopCoordX = field.NewFloat64(tableName, "stopCoordX")
	_hCV_vw_busTest.StopCoordY = field.NewFloat64(tableName, "stopCoordY")
	_hCV_vw_busTest.StopName = field.NewString(tableName, "stopName")
	_hCV_vw_busTest.Distance = field.NewFloat64(tableName, "distance")

	_hCV_vw_busTest.fillFieldMap()

	return _hCV_vw_busTest
}

type hCV_vw_busTest struct {
	hCV_vw_busTestDo

	ALL        field.Asterisk
	BusStopID  field.Int64
	StopCoordX field.Float64
	StopCoordY field.Float64
	StopName   field.String
	Distance   field.Float64

	fieldMap map[string]field.Expr
}

func (h hCV_vw_busTest) Table(newTableName string) *hCV_vw_busTest {
	h.hCV_vw_busTestDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_vw_busTest) As(alias string) *hCV_vw_busTest {
	h.hCV_vw_busTestDo.DO = *(h.hCV_vw_busTestDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_vw_busTest) updateTableName(table string) *hCV_vw_busTest {
	h.ALL = field.NewAsterisk(table)
	h.BusStopID = field.NewInt64(table, "busStopID")
	h.StopCoordX = field.NewFloat64(table, "stopCoordX")
	h.StopCoordY = field.NewFloat64(table, "stopCoordY")
	h.StopName = field.NewString(table, "stopName")
	h.Distance = field.NewFloat64(table, "distance")

	h.fillFieldMap()

	return h
}

func (h *hCV_vw_busTest) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_vw_busTest) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 5)
	h.fieldMap["busStopID"] = h.BusStopID
	h.fieldMap["stopCoordX"] = h.StopCoordX
	h.fieldMap["stopCoordY"] = h.StopCoordY
	h.fieldMap["stopName"] = h.StopName
	h.fieldMap["distance"] = h.Distance
}

func (h hCV_vw_busTest) clone(db *gorm.DB) hCV_vw_busTest {
	h.hCV_vw_busTestDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_vw_busTest) replaceDB(db *gorm.DB) hCV_vw_busTest {
	h.hCV_vw_busTestDo.ReplaceDB(db)
	return h
}

type hCV_vw_busTestDo struct{ gen.DO }

type IHCV_vw_busTestDo interface {
	gen.SubQuery
	Debug() IHCV_vw_busTestDo
	WithContext(ctx context.Context) IHCV_vw_busTestDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_vw_busTestDo
	WriteDB() IHCV_vw_busTestDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_vw_busTestDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_vw_busTestDo
	Not(conds ...gen.Condition) IHCV_vw_busTestDo
	Or(conds ...gen.Condition) IHCV_vw_busTestDo
	Select(conds ...field.Expr) IHCV_vw_busTestDo
	Where(conds ...gen.Condition) IHCV_vw_busTestDo
	Order(conds ...field.Expr) IHCV_vw_busTestDo
	Distinct(cols ...field.Expr) IHCV_vw_busTestDo
	Omit(cols ...field.Expr) IHCV_vw_busTestDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_vw_busTestDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_busTestDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_busTestDo
	Group(cols ...field.Expr) IHCV_vw_busTestDo
	Having(conds ...gen.Condition) IHCV_vw_busTestDo
	Limit(limit int) IHCV_vw_busTestDo
	Offset(offset int) IHCV_vw_busTestDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_busTestDo
	Unscoped() IHCV_vw_busTestDo
	Create(values ...*model.HCV_vw_busTest) error
	CreateInBatches(values []*model.HCV_vw_busTest, batchSize int) error
	Save(values ...*model.HCV_vw_busTest) error
	First() (*model.HCV_vw_busTest, error)
	Take() (*model.HCV_vw_busTest, error)
	Last() (*model.HCV_vw_busTest, error)
	Find() ([]*model.HCV_vw_busTest, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_busTest, err error)
	FindInBatches(result *[]*model.HCV_vw_busTest, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_vw_busTest) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_vw_busTestDo
	Assign(attrs ...field.AssignExpr) IHCV_vw_busTestDo
	Joins(fields ...field.RelationField) IHCV_vw_busTestDo
	Preload(fields ...field.RelationField) IHCV_vw_busTestDo
	FirstOrInit() (*model.HCV_vw_busTest, error)
	FirstOrCreate() (*model.HCV_vw_busTest, error)
	FindByPage(offset int, limit int) (result []*model.HCV_vw_busTest, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_vw_busTestDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_vw_busTestDo) Debug() IHCV_vw_busTestDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_vw_busTestDo) WithContext(ctx context.Context) IHCV_vw_busTestDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_vw_busTestDo) ReadDB() IHCV_vw_busTestDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_vw_busTestDo) WriteDB() IHCV_vw_busTestDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_vw_busTestDo) Session(config *gorm.Session) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_vw_busTestDo) Clauses(conds ...clause.Expression) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_vw_busTestDo) Returning(value interface{}, columns ...string) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_vw_busTestDo) Not(conds ...gen.Condition) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_vw_busTestDo) Or(conds ...gen.Condition) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_vw_busTestDo) Select(conds ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_vw_busTestDo) Where(conds ...gen.Condition) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_vw_busTestDo) Order(conds ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_vw_busTestDo) Distinct(cols ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_vw_busTestDo) Omit(cols ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_vw_busTestDo) Join(table schema.Tabler, on ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_vw_busTestDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_vw_busTestDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_vw_busTestDo) Group(cols ...field.Expr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_vw_busTestDo) Having(conds ...gen.Condition) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_vw_busTestDo) Limit(limit int) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_vw_busTestDo) Offset(offset int) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_vw_busTestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_vw_busTestDo) Unscoped() IHCV_vw_busTestDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_vw_busTestDo) Create(values ...*model.HCV_vw_busTest) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_vw_busTestDo) CreateInBatches(values []*model.HCV_vw_busTest, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_vw_busTestDo) Save(values ...*model.HCV_vw_busTest) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_vw_busTestDo) First() (*model.HCV_vw_busTest, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_busTest), nil
	}
}

func (h hCV_vw_busTestDo) Take() (*model.HCV_vw_busTest, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_busTest), nil
	}
}

func (h hCV_vw_busTestDo) Last() (*model.HCV_vw_busTest, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_busTest), nil
	}
}

func (h hCV_vw_busTestDo) Find() ([]*model.HCV_vw_busTest, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_vw_busTest), err
}

func (h hCV_vw_busTestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_busTest, err error) {
	buf := make([]*model.HCV_vw_busTest, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_vw_busTestDo) FindInBatches(result *[]*model.HCV_vw_busTest, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_vw_busTestDo) Attrs(attrs ...field.AssignExpr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_vw_busTestDo) Assign(attrs ...field.AssignExpr) IHCV_vw_busTestDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_vw_busTestDo) Joins(fields ...field.RelationField) IHCV_vw_busTestDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_vw_busTestDo) Preload(fields ...field.RelationField) IHCV_vw_busTestDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_vw_busTestDo) FirstOrInit() (*model.HCV_vw_busTest, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_busTest), nil
	}
}

func (h hCV_vw_busTestDo) FirstOrCreate() (*model.HCV_vw_busTest, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_busTest), nil
	}
}

func (h hCV_vw_busTestDo) FindByPage(offset int, limit int) (result []*model.HCV_vw_busTest, count int64, err error) {
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

func (h hCV_vw_busTestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_vw_busTestDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_vw_busTestDo) Delete(models ...*model.HCV_vw_busTest) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_vw_busTestDo) withDO(do gen.Dao) *hCV_vw_busTestDo {
	h.DO = *do.(*gen.DO)
	return h
}