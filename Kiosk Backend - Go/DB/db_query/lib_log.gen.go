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

func newLibLog(db *gorm.DB, opts ...gen.DOOption) libLog {
	_libLog := libLog{}

	_libLog.libLogDo.UseDB(db, opts...)
	_libLog.libLogDo.UseModel(&model.LibLog{})

	tableName := _libLog.libLogDo.TableName()
	_libLog.ALL = field.NewAsterisk(tableName)
	_libLog.Lib = field.NewString(tableName, "lib")
	_libLog.AdminID = field.NewString(tableName, "admin_id")
	_libLog.MenuID = field.NewString(tableName, "menu_id")
	_libLog.Crud = field.NewString(tableName, "crud")
	_libLog.RegDate = field.NewString(tableName, "reg_date")
	_libLog.Bigo = field.NewString(tableName, "bigo")

	_libLog.fillFieldMap()

	return _libLog
}

type libLog struct {
	libLogDo

	ALL     field.Asterisk
	Lib     field.String
	AdminID field.String
	MenuID  field.String
	Crud    field.String
	RegDate field.String
	Bigo    field.String

	fieldMap map[string]field.Expr
}

func (l libLog) Table(newTableName string) *libLog {
	l.libLogDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l libLog) As(alias string) *libLog {
	l.libLogDo.DO = *(l.libLogDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *libLog) updateTableName(table string) *libLog {
	l.ALL = field.NewAsterisk(table)
	l.Lib = field.NewString(table, "lib")
	l.AdminID = field.NewString(table, "admin_id")
	l.MenuID = field.NewString(table, "menu_id")
	l.Crud = field.NewString(table, "crud")
	l.RegDate = field.NewString(table, "reg_date")
	l.Bigo = field.NewString(table, "bigo")

	l.fillFieldMap()

	return l
}

func (l *libLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *libLog) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["lib"] = l.Lib
	l.fieldMap["admin_id"] = l.AdminID
	l.fieldMap["menu_id"] = l.MenuID
	l.fieldMap["crud"] = l.Crud
	l.fieldMap["reg_date"] = l.RegDate
	l.fieldMap["bigo"] = l.Bigo
}

func (l libLog) clone(db *gorm.DB) libLog {
	l.libLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l libLog) replaceDB(db *gorm.DB) libLog {
	l.libLogDo.ReplaceDB(db)
	return l
}

type libLogDo struct{ gen.DO }

type ILibLogDo interface {
	gen.SubQuery
	Debug() ILibLogDo
	WithContext(ctx context.Context) ILibLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILibLogDo
	WriteDB() ILibLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILibLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILibLogDo
	Not(conds ...gen.Condition) ILibLogDo
	Or(conds ...gen.Condition) ILibLogDo
	Select(conds ...field.Expr) ILibLogDo
	Where(conds ...gen.Condition) ILibLogDo
	Order(conds ...field.Expr) ILibLogDo
	Distinct(cols ...field.Expr) ILibLogDo
	Omit(cols ...field.Expr) ILibLogDo
	Join(table schema.Tabler, on ...field.Expr) ILibLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILibLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILibLogDo
	Group(cols ...field.Expr) ILibLogDo
	Having(conds ...gen.Condition) ILibLogDo
	Limit(limit int) ILibLogDo
	Offset(offset int) ILibLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILibLogDo
	Unscoped() ILibLogDo
	Create(values ...*model.LibLog) error
	CreateInBatches(values []*model.LibLog, batchSize int) error
	Save(values ...*model.LibLog) error
	First() (*model.LibLog, error)
	Take() (*model.LibLog, error)
	Last() (*model.LibLog, error)
	Find() ([]*model.LibLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LibLog, err error)
	FindInBatches(result *[]*model.LibLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LibLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILibLogDo
	Assign(attrs ...field.AssignExpr) ILibLogDo
	Joins(fields ...field.RelationField) ILibLogDo
	Preload(fields ...field.RelationField) ILibLogDo
	FirstOrInit() (*model.LibLog, error)
	FirstOrCreate() (*model.LibLog, error)
	FindByPage(offset int, limit int) (result []*model.LibLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILibLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l libLogDo) Debug() ILibLogDo {
	return l.withDO(l.DO.Debug())
}

func (l libLogDo) WithContext(ctx context.Context) ILibLogDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l libLogDo) ReadDB() ILibLogDo {
	return l.Clauses(dbresolver.Read)
}

func (l libLogDo) WriteDB() ILibLogDo {
	return l.Clauses(dbresolver.Write)
}

func (l libLogDo) Session(config *gorm.Session) ILibLogDo {
	return l.withDO(l.DO.Session(config))
}

func (l libLogDo) Clauses(conds ...clause.Expression) ILibLogDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l libLogDo) Returning(value interface{}, columns ...string) ILibLogDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l libLogDo) Not(conds ...gen.Condition) ILibLogDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l libLogDo) Or(conds ...gen.Condition) ILibLogDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l libLogDo) Select(conds ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l libLogDo) Where(conds ...gen.Condition) ILibLogDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l libLogDo) Order(conds ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l libLogDo) Distinct(cols ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l libLogDo) Omit(cols ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l libLogDo) Join(table schema.Tabler, on ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l libLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l libLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l libLogDo) Group(cols ...field.Expr) ILibLogDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l libLogDo) Having(conds ...gen.Condition) ILibLogDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l libLogDo) Limit(limit int) ILibLogDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l libLogDo) Offset(offset int) ILibLogDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l libLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILibLogDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l libLogDo) Unscoped() ILibLogDo {
	return l.withDO(l.DO.Unscoped())
}

func (l libLogDo) Create(values ...*model.LibLog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l libLogDo) CreateInBatches(values []*model.LibLog, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l libLogDo) Save(values ...*model.LibLog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l libLogDo) First() (*model.LibLog, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibLog), nil
	}
}

func (l libLogDo) Take() (*model.LibLog, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibLog), nil
	}
}

func (l libLogDo) Last() (*model.LibLog, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibLog), nil
	}
}

func (l libLogDo) Find() ([]*model.LibLog, error) {
	result, err := l.DO.Find()
	return result.([]*model.LibLog), err
}

func (l libLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LibLog, err error) {
	buf := make([]*model.LibLog, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l libLogDo) FindInBatches(result *[]*model.LibLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l libLogDo) Attrs(attrs ...field.AssignExpr) ILibLogDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l libLogDo) Assign(attrs ...field.AssignExpr) ILibLogDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l libLogDo) Joins(fields ...field.RelationField) ILibLogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l libLogDo) Preload(fields ...field.RelationField) ILibLogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l libLogDo) FirstOrInit() (*model.LibLog, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibLog), nil
	}
}

func (l libLogDo) FirstOrCreate() (*model.LibLog, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibLog), nil
	}
}

func (l libLogDo) FindByPage(offset int, limit int) (result []*model.LibLog, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l libLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l libLogDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l libLogDo) Delete(models ...*model.LibLog) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *libLogDo) withDO(do gen.Dao) *libLogDo {
	l.DO = *do.(*gen.DO)
	return l
}