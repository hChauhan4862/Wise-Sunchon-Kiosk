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

func newLimitUser(db *gorm.DB, opts ...gen.DOOption) limitUser {
	_limitUser := limitUser{}

	_limitUser.limitUserDo.UseDB(db, opts...)
	_limitUser.limitUserDo.UseModel(&model.LimitUser{})

	tableName := _limitUser.limitUserDo.TableName()
	_limitUser.ALL = field.NewAsterisk(tableName)
	_limitUser.LimitSeq = field.NewInt64(tableName, "limit_seq")
	_limitUser.LimitID = field.NewInt64(tableName, "limit_id")
	_limitUser.UserID = field.NewString(tableName, "user_id")
	_limitUser.UserName = field.NewString(tableName, "user_name")
	_limitUser.LimitCode = field.NewString(tableName, "limit_code")
	_limitUser.LimitDate = field.NewString(tableName, "limit_date")

	_limitUser.fillFieldMap()

	return _limitUser
}

type limitUser struct {
	limitUserDo

	ALL       field.Asterisk
	LimitSeq  field.Int64
	LimitID   field.Int64
	UserID    field.String
	UserName  field.String
	LimitCode field.String
	LimitDate field.String

	fieldMap map[string]field.Expr
}

func (l limitUser) Table(newTableName string) *limitUser {
	l.limitUserDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l limitUser) As(alias string) *limitUser {
	l.limitUserDo.DO = *(l.limitUserDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *limitUser) updateTableName(table string) *limitUser {
	l.ALL = field.NewAsterisk(table)
	l.LimitSeq = field.NewInt64(table, "limit_seq")
	l.LimitID = field.NewInt64(table, "limit_id")
	l.UserID = field.NewString(table, "user_id")
	l.UserName = field.NewString(table, "user_name")
	l.LimitCode = field.NewString(table, "limit_code")
	l.LimitDate = field.NewString(table, "limit_date")

	l.fillFieldMap()

	return l
}

func (l *limitUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *limitUser) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["limit_seq"] = l.LimitSeq
	l.fieldMap["limit_id"] = l.LimitID
	l.fieldMap["user_id"] = l.UserID
	l.fieldMap["user_name"] = l.UserName
	l.fieldMap["limit_code"] = l.LimitCode
	l.fieldMap["limit_date"] = l.LimitDate
}

func (l limitUser) clone(db *gorm.DB) limitUser {
	l.limitUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l limitUser) replaceDB(db *gorm.DB) limitUser {
	l.limitUserDo.ReplaceDB(db)
	return l
}

type limitUserDo struct{ gen.DO }

type ILimitUserDo interface {
	gen.SubQuery
	Debug() ILimitUserDo
	WithContext(ctx context.Context) ILimitUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILimitUserDo
	WriteDB() ILimitUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILimitUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILimitUserDo
	Not(conds ...gen.Condition) ILimitUserDo
	Or(conds ...gen.Condition) ILimitUserDo
	Select(conds ...field.Expr) ILimitUserDo
	Where(conds ...gen.Condition) ILimitUserDo
	Order(conds ...field.Expr) ILimitUserDo
	Distinct(cols ...field.Expr) ILimitUserDo
	Omit(cols ...field.Expr) ILimitUserDo
	Join(table schema.Tabler, on ...field.Expr) ILimitUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILimitUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILimitUserDo
	Group(cols ...field.Expr) ILimitUserDo
	Having(conds ...gen.Condition) ILimitUserDo
	Limit(limit int) ILimitUserDo
	Offset(offset int) ILimitUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILimitUserDo
	Unscoped() ILimitUserDo
	Create(values ...*model.LimitUser) error
	CreateInBatches(values []*model.LimitUser, batchSize int) error
	Save(values ...*model.LimitUser) error
	First() (*model.LimitUser, error)
	Take() (*model.LimitUser, error)
	Last() (*model.LimitUser, error)
	Find() ([]*model.LimitUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LimitUser, err error)
	FindInBatches(result *[]*model.LimitUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LimitUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILimitUserDo
	Assign(attrs ...field.AssignExpr) ILimitUserDo
	Joins(fields ...field.RelationField) ILimitUserDo
	Preload(fields ...field.RelationField) ILimitUserDo
	FirstOrInit() (*model.LimitUser, error)
	FirstOrCreate() (*model.LimitUser, error)
	FindByPage(offset int, limit int) (result []*model.LimitUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILimitUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l limitUserDo) Debug() ILimitUserDo {
	return l.withDO(l.DO.Debug())
}

func (l limitUserDo) WithContext(ctx context.Context) ILimitUserDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l limitUserDo) ReadDB() ILimitUserDo {
	return l.Clauses(dbresolver.Read)
}

func (l limitUserDo) WriteDB() ILimitUserDo {
	return l.Clauses(dbresolver.Write)
}

func (l limitUserDo) Session(config *gorm.Session) ILimitUserDo {
	return l.withDO(l.DO.Session(config))
}

func (l limitUserDo) Clauses(conds ...clause.Expression) ILimitUserDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l limitUserDo) Returning(value interface{}, columns ...string) ILimitUserDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l limitUserDo) Not(conds ...gen.Condition) ILimitUserDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l limitUserDo) Or(conds ...gen.Condition) ILimitUserDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l limitUserDo) Select(conds ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l limitUserDo) Where(conds ...gen.Condition) ILimitUserDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l limitUserDo) Order(conds ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l limitUserDo) Distinct(cols ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l limitUserDo) Omit(cols ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l limitUserDo) Join(table schema.Tabler, on ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l limitUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l limitUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l limitUserDo) Group(cols ...field.Expr) ILimitUserDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l limitUserDo) Having(conds ...gen.Condition) ILimitUserDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l limitUserDo) Limit(limit int) ILimitUserDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l limitUserDo) Offset(offset int) ILimitUserDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l limitUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILimitUserDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l limitUserDo) Unscoped() ILimitUserDo {
	return l.withDO(l.DO.Unscoped())
}

func (l limitUserDo) Create(values ...*model.LimitUser) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l limitUserDo) CreateInBatches(values []*model.LimitUser, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l limitUserDo) Save(values ...*model.LimitUser) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l limitUserDo) First() (*model.LimitUser, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LimitUser), nil
	}
}

func (l limitUserDo) Take() (*model.LimitUser, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LimitUser), nil
	}
}

func (l limitUserDo) Last() (*model.LimitUser, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LimitUser), nil
	}
}

func (l limitUserDo) Find() ([]*model.LimitUser, error) {
	result, err := l.DO.Find()
	return result.([]*model.LimitUser), err
}

func (l limitUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LimitUser, err error) {
	buf := make([]*model.LimitUser, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l limitUserDo) FindInBatches(result *[]*model.LimitUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l limitUserDo) Attrs(attrs ...field.AssignExpr) ILimitUserDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l limitUserDo) Assign(attrs ...field.AssignExpr) ILimitUserDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l limitUserDo) Joins(fields ...field.RelationField) ILimitUserDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l limitUserDo) Preload(fields ...field.RelationField) ILimitUserDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l limitUserDo) FirstOrInit() (*model.LimitUser, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LimitUser), nil
	}
}

func (l limitUserDo) FirstOrCreate() (*model.LimitUser, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LimitUser), nil
	}
}

func (l limitUserDo) FindByPage(offset int, limit int) (result []*model.LimitUser, count int64, err error) {
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

func (l limitUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l limitUserDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l limitUserDo) Delete(models ...*model.LimitUser) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *limitUserDo) withDO(do gen.Dao) *limitUserDo {
	l.DO = *do.(*gen.DO)
	return l
}