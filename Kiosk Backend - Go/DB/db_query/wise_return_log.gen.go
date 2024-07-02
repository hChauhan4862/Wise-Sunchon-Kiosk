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

func newWiseReturnLog(db *gorm.DB, opts ...gen.DOOption) wiseReturnLog {
	_wiseReturnLog := wiseReturnLog{}

	_wiseReturnLog.wiseReturnLogDo.UseDB(db, opts...)
	_wiseReturnLog.wiseReturnLogDo.UseModel(&model.WiseReturnLog{})

	tableName := _wiseReturnLog.wiseReturnLogDo.TableName()
	_wiseReturnLog.ALL = field.NewAsterisk(tableName)
	_wiseReturnLog.Rid = field.NewInt64(tableName, "rid")
	_wiseReturnLog.PcNo = field.NewInt64(tableName, "pc_no")
	_wiseReturnLog.ReturnTime = field.NewString(tableName, "return_time")
	_wiseReturnLog.ReturnGubun = field.NewString(tableName, "return_gubun")
	_wiseReturnLog.UserID = field.NewString(tableName, "user_id")
	_wiseReturnLog.UserName = field.NewString(tableName, "user_name")
	_wiseReturnLog.BookNo = field.NewString(tableName, "book_no")
	_wiseReturnLog.Title = field.NewString(tableName, "title")
	_wiseReturnLog.Author = field.NewString(tableName, "author")
	_wiseReturnLog.Publisher = field.NewString(tableName, "publisher")
	_wiseReturnLog.CallNo = field.NewString(tableName, "call_no")
	_wiseReturnLog.Libcode = field.NewString(tableName, "libcode")
	_wiseReturnLog.PositionName = field.NewString(tableName, "position_name")
	_wiseReturnLog.DeptName = field.NewString(tableName, "dept_name")

	_wiseReturnLog.fillFieldMap()

	return _wiseReturnLog
}

type wiseReturnLog struct {
	wiseReturnLogDo

	ALL          field.Asterisk
	Rid          field.Int64
	PcNo         field.Int64
	ReturnTime   field.String
	ReturnGubun  field.String
	UserID       field.String
	UserName     field.String
	BookNo       field.String
	Title        field.String
	Author       field.String
	Publisher    field.String
	CallNo       field.String
	Libcode      field.String
	PositionName field.String
	DeptName     field.String

	fieldMap map[string]field.Expr
}

func (w wiseReturnLog) Table(newTableName string) *wiseReturnLog {
	w.wiseReturnLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wiseReturnLog) As(alias string) *wiseReturnLog {
	w.wiseReturnLogDo.DO = *(w.wiseReturnLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wiseReturnLog) updateTableName(table string) *wiseReturnLog {
	w.ALL = field.NewAsterisk(table)
	w.Rid = field.NewInt64(table, "rid")
	w.PcNo = field.NewInt64(table, "pc_no")
	w.ReturnTime = field.NewString(table, "return_time")
	w.ReturnGubun = field.NewString(table, "return_gubun")
	w.UserID = field.NewString(table, "user_id")
	w.UserName = field.NewString(table, "user_name")
	w.BookNo = field.NewString(table, "book_no")
	w.Title = field.NewString(table, "title")
	w.Author = field.NewString(table, "author")
	w.Publisher = field.NewString(table, "publisher")
	w.CallNo = field.NewString(table, "call_no")
	w.Libcode = field.NewString(table, "libcode")
	w.PositionName = field.NewString(table, "position_name")
	w.DeptName = field.NewString(table, "dept_name")

	w.fillFieldMap()

	return w
}

func (w *wiseReturnLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wiseReturnLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["rid"] = w.Rid
	w.fieldMap["pc_no"] = w.PcNo
	w.fieldMap["return_time"] = w.ReturnTime
	w.fieldMap["return_gubun"] = w.ReturnGubun
	w.fieldMap["user_id"] = w.UserID
	w.fieldMap["user_name"] = w.UserName
	w.fieldMap["book_no"] = w.BookNo
	w.fieldMap["title"] = w.Title
	w.fieldMap["author"] = w.Author
	w.fieldMap["publisher"] = w.Publisher
	w.fieldMap["call_no"] = w.CallNo
	w.fieldMap["libcode"] = w.Libcode
	w.fieldMap["position_name"] = w.PositionName
	w.fieldMap["dept_name"] = w.DeptName
}

func (w wiseReturnLog) clone(db *gorm.DB) wiseReturnLog {
	w.wiseReturnLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wiseReturnLog) replaceDB(db *gorm.DB) wiseReturnLog {
	w.wiseReturnLogDo.ReplaceDB(db)
	return w
}

type wiseReturnLogDo struct{ gen.DO }

type IWiseReturnLogDo interface {
	gen.SubQuery
	Debug() IWiseReturnLogDo
	WithContext(ctx context.Context) IWiseReturnLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWiseReturnLogDo
	WriteDB() IWiseReturnLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWiseReturnLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWiseReturnLogDo
	Not(conds ...gen.Condition) IWiseReturnLogDo
	Or(conds ...gen.Condition) IWiseReturnLogDo
	Select(conds ...field.Expr) IWiseReturnLogDo
	Where(conds ...gen.Condition) IWiseReturnLogDo
	Order(conds ...field.Expr) IWiseReturnLogDo
	Distinct(cols ...field.Expr) IWiseReturnLogDo
	Omit(cols ...field.Expr) IWiseReturnLogDo
	Join(table schema.Tabler, on ...field.Expr) IWiseReturnLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWiseReturnLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWiseReturnLogDo
	Group(cols ...field.Expr) IWiseReturnLogDo
	Having(conds ...gen.Condition) IWiseReturnLogDo
	Limit(limit int) IWiseReturnLogDo
	Offset(offset int) IWiseReturnLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseReturnLogDo
	Unscoped() IWiseReturnLogDo
	Create(values ...*model.WiseReturnLog) error
	CreateInBatches(values []*model.WiseReturnLog, batchSize int) error
	Save(values ...*model.WiseReturnLog) error
	First() (*model.WiseReturnLog, error)
	Take() (*model.WiseReturnLog, error)
	Last() (*model.WiseReturnLog, error)
	Find() ([]*model.WiseReturnLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseReturnLog, err error)
	FindInBatches(result *[]*model.WiseReturnLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WiseReturnLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWiseReturnLogDo
	Assign(attrs ...field.AssignExpr) IWiseReturnLogDo
	Joins(fields ...field.RelationField) IWiseReturnLogDo
	Preload(fields ...field.RelationField) IWiseReturnLogDo
	FirstOrInit() (*model.WiseReturnLog, error)
	FirstOrCreate() (*model.WiseReturnLog, error)
	FindByPage(offset int, limit int) (result []*model.WiseReturnLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWiseReturnLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wiseReturnLogDo) Debug() IWiseReturnLogDo {
	return w.withDO(w.DO.Debug())
}

func (w wiseReturnLogDo) WithContext(ctx context.Context) IWiseReturnLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wiseReturnLogDo) ReadDB() IWiseReturnLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w wiseReturnLogDo) WriteDB() IWiseReturnLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w wiseReturnLogDo) Session(config *gorm.Session) IWiseReturnLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w wiseReturnLogDo) Clauses(conds ...clause.Expression) IWiseReturnLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wiseReturnLogDo) Returning(value interface{}, columns ...string) IWiseReturnLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wiseReturnLogDo) Not(conds ...gen.Condition) IWiseReturnLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wiseReturnLogDo) Or(conds ...gen.Condition) IWiseReturnLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wiseReturnLogDo) Select(conds ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wiseReturnLogDo) Where(conds ...gen.Condition) IWiseReturnLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wiseReturnLogDo) Order(conds ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wiseReturnLogDo) Distinct(cols ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wiseReturnLogDo) Omit(cols ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wiseReturnLogDo) Join(table schema.Tabler, on ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wiseReturnLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wiseReturnLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wiseReturnLogDo) Group(cols ...field.Expr) IWiseReturnLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wiseReturnLogDo) Having(conds ...gen.Condition) IWiseReturnLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wiseReturnLogDo) Limit(limit int) IWiseReturnLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wiseReturnLogDo) Offset(offset int) IWiseReturnLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wiseReturnLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseReturnLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wiseReturnLogDo) Unscoped() IWiseReturnLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wiseReturnLogDo) Create(values ...*model.WiseReturnLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wiseReturnLogDo) CreateInBatches(values []*model.WiseReturnLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wiseReturnLogDo) Save(values ...*model.WiseReturnLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wiseReturnLogDo) First() (*model.WiseReturnLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseReturnLog), nil
	}
}

func (w wiseReturnLogDo) Take() (*model.WiseReturnLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseReturnLog), nil
	}
}

func (w wiseReturnLogDo) Last() (*model.WiseReturnLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseReturnLog), nil
	}
}

func (w wiseReturnLogDo) Find() ([]*model.WiseReturnLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WiseReturnLog), err
}

func (w wiseReturnLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseReturnLog, err error) {
	buf := make([]*model.WiseReturnLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wiseReturnLogDo) FindInBatches(result *[]*model.WiseReturnLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wiseReturnLogDo) Attrs(attrs ...field.AssignExpr) IWiseReturnLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wiseReturnLogDo) Assign(attrs ...field.AssignExpr) IWiseReturnLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wiseReturnLogDo) Joins(fields ...field.RelationField) IWiseReturnLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wiseReturnLogDo) Preload(fields ...field.RelationField) IWiseReturnLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wiseReturnLogDo) FirstOrInit() (*model.WiseReturnLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseReturnLog), nil
	}
}

func (w wiseReturnLogDo) FirstOrCreate() (*model.WiseReturnLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseReturnLog), nil
	}
}

func (w wiseReturnLogDo) FindByPage(offset int, limit int) (result []*model.WiseReturnLog, count int64, err error) {
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

func (w wiseReturnLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wiseReturnLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wiseReturnLogDo) Delete(models ...*model.WiseReturnLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wiseReturnLogDo) withDO(do gen.Dao) *wiseReturnLogDo {
	w.DO = *do.(*gen.DO)
	return w
}
