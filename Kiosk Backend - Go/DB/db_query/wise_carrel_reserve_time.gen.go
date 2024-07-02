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

func newWiseCarrelReserveTime(db *gorm.DB, opts ...gen.DOOption) wiseCarrelReserveTime {
	_wiseCarrelReserveTime := wiseCarrelReserveTime{}

	_wiseCarrelReserveTime.wiseCarrelReserveTimeDo.UseDB(db, opts...)
	_wiseCarrelReserveTime.wiseCarrelReserveTimeDo.UseModel(&model.WiseCarrelReserveTime{})

	tableName := _wiseCarrelReserveTime.wiseCarrelReserveTimeDo.TableName()
	_wiseCarrelReserveTime.ALL = field.NewAsterisk(tableName)
	_wiseCarrelReserveTime.ReserveTimeNo = field.NewString(tableName, "reserve_time_no")
	_wiseCarrelReserveTime.StartTime = field.NewString(tableName, "start_time")
	_wiseCarrelReserveTime.EndTime = field.NewString(tableName, "end_time")
	_wiseCarrelReserveTime.UseYn = field.NewString(tableName, "use_yn")

	_wiseCarrelReserveTime.fillFieldMap()

	return _wiseCarrelReserveTime
}

type wiseCarrelReserveTime struct {
	wiseCarrelReserveTimeDo

	ALL           field.Asterisk
	ReserveTimeNo field.String
	StartTime     field.String
	EndTime       field.String
	UseYn         field.String

	fieldMap map[string]field.Expr
}

func (w wiseCarrelReserveTime) Table(newTableName string) *wiseCarrelReserveTime {
	w.wiseCarrelReserveTimeDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wiseCarrelReserveTime) As(alias string) *wiseCarrelReserveTime {
	w.wiseCarrelReserveTimeDo.DO = *(w.wiseCarrelReserveTimeDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wiseCarrelReserveTime) updateTableName(table string) *wiseCarrelReserveTime {
	w.ALL = field.NewAsterisk(table)
	w.ReserveTimeNo = field.NewString(table, "reserve_time_no")
	w.StartTime = field.NewString(table, "start_time")
	w.EndTime = field.NewString(table, "end_time")
	w.UseYn = field.NewString(table, "use_yn")

	w.fillFieldMap()

	return w
}

func (w *wiseCarrelReserveTime) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wiseCarrelReserveTime) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 4)
	w.fieldMap["reserve_time_no"] = w.ReserveTimeNo
	w.fieldMap["start_time"] = w.StartTime
	w.fieldMap["end_time"] = w.EndTime
	w.fieldMap["use_yn"] = w.UseYn
}

func (w wiseCarrelReserveTime) clone(db *gorm.DB) wiseCarrelReserveTime {
	w.wiseCarrelReserveTimeDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wiseCarrelReserveTime) replaceDB(db *gorm.DB) wiseCarrelReserveTime {
	w.wiseCarrelReserveTimeDo.ReplaceDB(db)
	return w
}

type wiseCarrelReserveTimeDo struct{ gen.DO }

type IWiseCarrelReserveTimeDo interface {
	gen.SubQuery
	Debug() IWiseCarrelReserveTimeDo
	WithContext(ctx context.Context) IWiseCarrelReserveTimeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWiseCarrelReserveTimeDo
	WriteDB() IWiseCarrelReserveTimeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWiseCarrelReserveTimeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWiseCarrelReserveTimeDo
	Not(conds ...gen.Condition) IWiseCarrelReserveTimeDo
	Or(conds ...gen.Condition) IWiseCarrelReserveTimeDo
	Select(conds ...field.Expr) IWiseCarrelReserveTimeDo
	Where(conds ...gen.Condition) IWiseCarrelReserveTimeDo
	Order(conds ...field.Expr) IWiseCarrelReserveTimeDo
	Distinct(cols ...field.Expr) IWiseCarrelReserveTimeDo
	Omit(cols ...field.Expr) IWiseCarrelReserveTimeDo
	Join(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveTimeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveTimeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveTimeDo
	Group(cols ...field.Expr) IWiseCarrelReserveTimeDo
	Having(conds ...gen.Condition) IWiseCarrelReserveTimeDo
	Limit(limit int) IWiseCarrelReserveTimeDo
	Offset(offset int) IWiseCarrelReserveTimeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseCarrelReserveTimeDo
	Unscoped() IWiseCarrelReserveTimeDo
	Create(values ...*model.WiseCarrelReserveTime) error
	CreateInBatches(values []*model.WiseCarrelReserveTime, batchSize int) error
	Save(values ...*model.WiseCarrelReserveTime) error
	First() (*model.WiseCarrelReserveTime, error)
	Take() (*model.WiseCarrelReserveTime, error)
	Last() (*model.WiseCarrelReserveTime, error)
	Find() ([]*model.WiseCarrelReserveTime, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseCarrelReserveTime, err error)
	FindInBatches(result *[]*model.WiseCarrelReserveTime, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WiseCarrelReserveTime) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWiseCarrelReserveTimeDo
	Assign(attrs ...field.AssignExpr) IWiseCarrelReserveTimeDo
	Joins(fields ...field.RelationField) IWiseCarrelReserveTimeDo
	Preload(fields ...field.RelationField) IWiseCarrelReserveTimeDo
	FirstOrInit() (*model.WiseCarrelReserveTime, error)
	FirstOrCreate() (*model.WiseCarrelReserveTime, error)
	FindByPage(offset int, limit int) (result []*model.WiseCarrelReserveTime, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWiseCarrelReserveTimeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wiseCarrelReserveTimeDo) Debug() IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Debug())
}

func (w wiseCarrelReserveTimeDo) WithContext(ctx context.Context) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wiseCarrelReserveTimeDo) ReadDB() IWiseCarrelReserveTimeDo {
	return w.Clauses(dbresolver.Read)
}

func (w wiseCarrelReserveTimeDo) WriteDB() IWiseCarrelReserveTimeDo {
	return w.Clauses(dbresolver.Write)
}

func (w wiseCarrelReserveTimeDo) Session(config *gorm.Session) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Session(config))
}

func (w wiseCarrelReserveTimeDo) Clauses(conds ...clause.Expression) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wiseCarrelReserveTimeDo) Returning(value interface{}, columns ...string) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wiseCarrelReserveTimeDo) Not(conds ...gen.Condition) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wiseCarrelReserveTimeDo) Or(conds ...gen.Condition) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wiseCarrelReserveTimeDo) Select(conds ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wiseCarrelReserveTimeDo) Where(conds ...gen.Condition) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wiseCarrelReserveTimeDo) Order(conds ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wiseCarrelReserveTimeDo) Distinct(cols ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wiseCarrelReserveTimeDo) Omit(cols ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wiseCarrelReserveTimeDo) Join(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wiseCarrelReserveTimeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wiseCarrelReserveTimeDo) RightJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wiseCarrelReserveTimeDo) Group(cols ...field.Expr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wiseCarrelReserveTimeDo) Having(conds ...gen.Condition) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wiseCarrelReserveTimeDo) Limit(limit int) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wiseCarrelReserveTimeDo) Offset(offset int) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wiseCarrelReserveTimeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wiseCarrelReserveTimeDo) Unscoped() IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wiseCarrelReserveTimeDo) Create(values ...*model.WiseCarrelReserveTime) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wiseCarrelReserveTimeDo) CreateInBatches(values []*model.WiseCarrelReserveTime, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wiseCarrelReserveTimeDo) Save(values ...*model.WiseCarrelReserveTime) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wiseCarrelReserveTimeDo) First() (*model.WiseCarrelReserveTime, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveTime), nil
	}
}

func (w wiseCarrelReserveTimeDo) Take() (*model.WiseCarrelReserveTime, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveTime), nil
	}
}

func (w wiseCarrelReserveTimeDo) Last() (*model.WiseCarrelReserveTime, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveTime), nil
	}
}

func (w wiseCarrelReserveTimeDo) Find() ([]*model.WiseCarrelReserveTime, error) {
	result, err := w.DO.Find()
	return result.([]*model.WiseCarrelReserveTime), err
}

func (w wiseCarrelReserveTimeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseCarrelReserveTime, err error) {
	buf := make([]*model.WiseCarrelReserveTime, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wiseCarrelReserveTimeDo) FindInBatches(result *[]*model.WiseCarrelReserveTime, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wiseCarrelReserveTimeDo) Attrs(attrs ...field.AssignExpr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wiseCarrelReserveTimeDo) Assign(attrs ...field.AssignExpr) IWiseCarrelReserveTimeDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wiseCarrelReserveTimeDo) Joins(fields ...field.RelationField) IWiseCarrelReserveTimeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wiseCarrelReserveTimeDo) Preload(fields ...field.RelationField) IWiseCarrelReserveTimeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wiseCarrelReserveTimeDo) FirstOrInit() (*model.WiseCarrelReserveTime, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveTime), nil
	}
}

func (w wiseCarrelReserveTimeDo) FirstOrCreate() (*model.WiseCarrelReserveTime, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveTime), nil
	}
}

func (w wiseCarrelReserveTimeDo) FindByPage(offset int, limit int) (result []*model.WiseCarrelReserveTime, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wiseCarrelReserveTimeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wiseCarrelReserveTimeDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wiseCarrelReserveTimeDo) Delete(models ...*model.WiseCarrelReserveTime) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wiseCarrelReserveTimeDo) withDO(do gen.Dao) *wiseCarrelReserveTimeDo {
	w.DO = *do.(*gen.DO)
	return w
}
