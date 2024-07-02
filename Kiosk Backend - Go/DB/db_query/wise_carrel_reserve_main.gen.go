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

func newWiseCarrelReserveMain(db *gorm.DB, opts ...gen.DOOption) wiseCarrelReserveMain {
	_wiseCarrelReserveMain := wiseCarrelReserveMain{}

	_wiseCarrelReserveMain.wiseCarrelReserveMainDo.UseDB(db, opts...)
	_wiseCarrelReserveMain.wiseCarrelReserveMainDo.UseModel(&model.WiseCarrelReserveMain{})

	tableName := _wiseCarrelReserveMain.wiseCarrelReserveMainDo.TableName()
	_wiseCarrelReserveMain.ALL = field.NewAsterisk(tableName)
	_wiseCarrelReserveMain.ReserveNo = field.NewString(tableName, "reserve_no")
	_wiseCarrelReserveMain.UserID = field.NewString(tableName, "user_id")
	_wiseCarrelReserveMain.UserName = field.NewString(tableName, "user_name")
	_wiseCarrelReserveMain.StartDate = field.NewString(tableName, "start_date")
	_wiseCarrelReserveMain.EndDate = field.NewString(tableName, "end_date")
	_wiseCarrelReserveMain.ReserveStat = field.NewString(tableName, "reserve_stat")
	_wiseCarrelReserveMain.RegDate = field.NewString(tableName, "reg_date")
	_wiseCarrelReserveMain.UptDate = field.NewString(tableName, "upt_date")
	_wiseCarrelReserveMain.CarrelNo = field.NewString(tableName, "carrel_no")
	_wiseCarrelReserveMain.UseMonth = field.NewString(tableName, "use_month")

	_wiseCarrelReserveMain.fillFieldMap()

	return _wiseCarrelReserveMain
}

type wiseCarrelReserveMain struct {
	wiseCarrelReserveMainDo

	ALL         field.Asterisk
	ReserveNo   field.String
	UserID      field.String
	UserName    field.String
	StartDate   field.String
	EndDate     field.String
	ReserveStat field.String
	RegDate     field.String
	UptDate     field.String
	CarrelNo    field.String
	UseMonth    field.String

	fieldMap map[string]field.Expr
}

func (w wiseCarrelReserveMain) Table(newTableName string) *wiseCarrelReserveMain {
	w.wiseCarrelReserveMainDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wiseCarrelReserveMain) As(alias string) *wiseCarrelReserveMain {
	w.wiseCarrelReserveMainDo.DO = *(w.wiseCarrelReserveMainDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wiseCarrelReserveMain) updateTableName(table string) *wiseCarrelReserveMain {
	w.ALL = field.NewAsterisk(table)
	w.ReserveNo = field.NewString(table, "reserve_no")
	w.UserID = field.NewString(table, "user_id")
	w.UserName = field.NewString(table, "user_name")
	w.StartDate = field.NewString(table, "start_date")
	w.EndDate = field.NewString(table, "end_date")
	w.ReserveStat = field.NewString(table, "reserve_stat")
	w.RegDate = field.NewString(table, "reg_date")
	w.UptDate = field.NewString(table, "upt_date")
	w.CarrelNo = field.NewString(table, "carrel_no")
	w.UseMonth = field.NewString(table, "use_month")

	w.fillFieldMap()

	return w
}

func (w *wiseCarrelReserveMain) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wiseCarrelReserveMain) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["reserve_no"] = w.ReserveNo
	w.fieldMap["user_id"] = w.UserID
	w.fieldMap["user_name"] = w.UserName
	w.fieldMap["start_date"] = w.StartDate
	w.fieldMap["end_date"] = w.EndDate
	w.fieldMap["reserve_stat"] = w.ReserveStat
	w.fieldMap["reg_date"] = w.RegDate
	w.fieldMap["upt_date"] = w.UptDate
	w.fieldMap["carrel_no"] = w.CarrelNo
	w.fieldMap["use_month"] = w.UseMonth
}

func (w wiseCarrelReserveMain) clone(db *gorm.DB) wiseCarrelReserveMain {
	w.wiseCarrelReserveMainDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wiseCarrelReserveMain) replaceDB(db *gorm.DB) wiseCarrelReserveMain {
	w.wiseCarrelReserveMainDo.ReplaceDB(db)
	return w
}

type wiseCarrelReserveMainDo struct{ gen.DO }

type IWiseCarrelReserveMainDo interface {
	gen.SubQuery
	Debug() IWiseCarrelReserveMainDo
	WithContext(ctx context.Context) IWiseCarrelReserveMainDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWiseCarrelReserveMainDo
	WriteDB() IWiseCarrelReserveMainDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWiseCarrelReserveMainDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWiseCarrelReserveMainDo
	Not(conds ...gen.Condition) IWiseCarrelReserveMainDo
	Or(conds ...gen.Condition) IWiseCarrelReserveMainDo
	Select(conds ...field.Expr) IWiseCarrelReserveMainDo
	Where(conds ...gen.Condition) IWiseCarrelReserveMainDo
	Order(conds ...field.Expr) IWiseCarrelReserveMainDo
	Distinct(cols ...field.Expr) IWiseCarrelReserveMainDo
	Omit(cols ...field.Expr) IWiseCarrelReserveMainDo
	Join(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveMainDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveMainDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveMainDo
	Group(cols ...field.Expr) IWiseCarrelReserveMainDo
	Having(conds ...gen.Condition) IWiseCarrelReserveMainDo
	Limit(limit int) IWiseCarrelReserveMainDo
	Offset(offset int) IWiseCarrelReserveMainDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseCarrelReserveMainDo
	Unscoped() IWiseCarrelReserveMainDo
	Create(values ...*model.WiseCarrelReserveMain) error
	CreateInBatches(values []*model.WiseCarrelReserveMain, batchSize int) error
	Save(values ...*model.WiseCarrelReserveMain) error
	First() (*model.WiseCarrelReserveMain, error)
	Take() (*model.WiseCarrelReserveMain, error)
	Last() (*model.WiseCarrelReserveMain, error)
	Find() ([]*model.WiseCarrelReserveMain, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseCarrelReserveMain, err error)
	FindInBatches(result *[]*model.WiseCarrelReserveMain, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WiseCarrelReserveMain) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWiseCarrelReserveMainDo
	Assign(attrs ...field.AssignExpr) IWiseCarrelReserveMainDo
	Joins(fields ...field.RelationField) IWiseCarrelReserveMainDo
	Preload(fields ...field.RelationField) IWiseCarrelReserveMainDo
	FirstOrInit() (*model.WiseCarrelReserveMain, error)
	FirstOrCreate() (*model.WiseCarrelReserveMain, error)
	FindByPage(offset int, limit int) (result []*model.WiseCarrelReserveMain, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWiseCarrelReserveMainDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wiseCarrelReserveMainDo) Debug() IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Debug())
}

func (w wiseCarrelReserveMainDo) WithContext(ctx context.Context) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wiseCarrelReserveMainDo) ReadDB() IWiseCarrelReserveMainDo {
	return w.Clauses(dbresolver.Read)
}

func (w wiseCarrelReserveMainDo) WriteDB() IWiseCarrelReserveMainDo {
	return w.Clauses(dbresolver.Write)
}

func (w wiseCarrelReserveMainDo) Session(config *gorm.Session) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Session(config))
}

func (w wiseCarrelReserveMainDo) Clauses(conds ...clause.Expression) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wiseCarrelReserveMainDo) Returning(value interface{}, columns ...string) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wiseCarrelReserveMainDo) Not(conds ...gen.Condition) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wiseCarrelReserveMainDo) Or(conds ...gen.Condition) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wiseCarrelReserveMainDo) Select(conds ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wiseCarrelReserveMainDo) Where(conds ...gen.Condition) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wiseCarrelReserveMainDo) Order(conds ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wiseCarrelReserveMainDo) Distinct(cols ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wiseCarrelReserveMainDo) Omit(cols ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wiseCarrelReserveMainDo) Join(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wiseCarrelReserveMainDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wiseCarrelReserveMainDo) RightJoin(table schema.Tabler, on ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wiseCarrelReserveMainDo) Group(cols ...field.Expr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wiseCarrelReserveMainDo) Having(conds ...gen.Condition) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wiseCarrelReserveMainDo) Limit(limit int) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wiseCarrelReserveMainDo) Offset(offset int) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wiseCarrelReserveMainDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wiseCarrelReserveMainDo) Unscoped() IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wiseCarrelReserveMainDo) Create(values ...*model.WiseCarrelReserveMain) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wiseCarrelReserveMainDo) CreateInBatches(values []*model.WiseCarrelReserveMain, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wiseCarrelReserveMainDo) Save(values ...*model.WiseCarrelReserveMain) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wiseCarrelReserveMainDo) First() (*model.WiseCarrelReserveMain, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveMain), nil
	}
}

func (w wiseCarrelReserveMainDo) Take() (*model.WiseCarrelReserveMain, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveMain), nil
	}
}

func (w wiseCarrelReserveMainDo) Last() (*model.WiseCarrelReserveMain, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveMain), nil
	}
}

func (w wiseCarrelReserveMainDo) Find() ([]*model.WiseCarrelReserveMain, error) {
	result, err := w.DO.Find()
	return result.([]*model.WiseCarrelReserveMain), err
}

func (w wiseCarrelReserveMainDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseCarrelReserveMain, err error) {
	buf := make([]*model.WiseCarrelReserveMain, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wiseCarrelReserveMainDo) FindInBatches(result *[]*model.WiseCarrelReserveMain, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wiseCarrelReserveMainDo) Attrs(attrs ...field.AssignExpr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wiseCarrelReserveMainDo) Assign(attrs ...field.AssignExpr) IWiseCarrelReserveMainDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wiseCarrelReserveMainDo) Joins(fields ...field.RelationField) IWiseCarrelReserveMainDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wiseCarrelReserveMainDo) Preload(fields ...field.RelationField) IWiseCarrelReserveMainDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wiseCarrelReserveMainDo) FirstOrInit() (*model.WiseCarrelReserveMain, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveMain), nil
	}
}

func (w wiseCarrelReserveMainDo) FirstOrCreate() (*model.WiseCarrelReserveMain, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseCarrelReserveMain), nil
	}
}

func (w wiseCarrelReserveMainDo) FindByPage(offset int, limit int) (result []*model.WiseCarrelReserveMain, count int64, err error) {
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

func (w wiseCarrelReserveMainDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wiseCarrelReserveMainDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wiseCarrelReserveMainDo) Delete(models ...*model.WiseCarrelReserveMain) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wiseCarrelReserveMainDo) withDO(do gen.Dao) *wiseCarrelReserveMainDo {
	w.DO = *do.(*gen.DO)
	return w
}
