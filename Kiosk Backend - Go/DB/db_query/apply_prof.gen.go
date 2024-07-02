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

func newAPPLYPROF(db *gorm.DB, opts ...gen.DOOption) aPPLYPROF {
	_aPPLYPROF := aPPLYPROF{}

	_aPPLYPROF.aPPLYPROFDo.UseDB(db, opts...)
	_aPPLYPROF.aPPLYPROFDo.UseModel(&model.APPLYPROF{})

	tableName := _aPPLYPROF.aPPLYPROFDo.TableName()
	_aPPLYPROF.ALL = field.NewAsterisk(tableName)
	_aPPLYPROF.SEQID = field.NewInt64(tableName, "SEQ_ID")
	_aPPLYPROF.PROFID = field.NewString(tableName, "PROF_ID")
	_aPPLYPROF.PROFNAME = field.NewString(tableName, "PROF_NAME")
	_aPPLYPROF.PHOTOURL = field.NewString(tableName, "PHOTO_URL")
	_aPPLYPROF.POSITIONCODE = field.NewString(tableName, "POSITION_CODE")
	_aPPLYPROF.POSITIONNAME = field.NewString(tableName, "POSITION_NAME")
	_aPPLYPROF.STATUSCODE = field.NewString(tableName, "STATUS_CODE")
	_aPPLYPROF.STATUSNAME = field.NewString(tableName, "STATUS_NAME")
	_aPPLYPROF.INSERTDATE = field.NewString(tableName, "INSERT_DATE")

	_aPPLYPROF.fillFieldMap()

	return _aPPLYPROF
}

type aPPLYPROF struct {
	aPPLYPROFDo

	ALL          field.Asterisk
	SEQID        field.Int64
	PROFID       field.String
	PROFNAME     field.String
	PHOTOURL     field.String
	POSITIONCODE field.String
	POSITIONNAME field.String
	STATUSCODE   field.String
	STATUSNAME   field.String
	INSERTDATE   field.String

	fieldMap map[string]field.Expr
}

func (a aPPLYPROF) Table(newTableName string) *aPPLYPROF {
	a.aPPLYPROFDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aPPLYPROF) As(alias string) *aPPLYPROF {
	a.aPPLYPROFDo.DO = *(a.aPPLYPROFDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aPPLYPROF) updateTableName(table string) *aPPLYPROF {
	a.ALL = field.NewAsterisk(table)
	a.SEQID = field.NewInt64(table, "SEQ_ID")
	a.PROFID = field.NewString(table, "PROF_ID")
	a.PROFNAME = field.NewString(table, "PROF_NAME")
	a.PHOTOURL = field.NewString(table, "PHOTO_URL")
	a.POSITIONCODE = field.NewString(table, "POSITION_CODE")
	a.POSITIONNAME = field.NewString(table, "POSITION_NAME")
	a.STATUSCODE = field.NewString(table, "STATUS_CODE")
	a.STATUSNAME = field.NewString(table, "STATUS_NAME")
	a.INSERTDATE = field.NewString(table, "INSERT_DATE")

	a.fillFieldMap()

	return a
}

func (a *aPPLYPROF) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aPPLYPROF) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["SEQ_ID"] = a.SEQID
	a.fieldMap["PROF_ID"] = a.PROFID
	a.fieldMap["PROF_NAME"] = a.PROFNAME
	a.fieldMap["PHOTO_URL"] = a.PHOTOURL
	a.fieldMap["POSITION_CODE"] = a.POSITIONCODE
	a.fieldMap["POSITION_NAME"] = a.POSITIONNAME
	a.fieldMap["STATUS_CODE"] = a.STATUSCODE
	a.fieldMap["STATUS_NAME"] = a.STATUSNAME
	a.fieldMap["INSERT_DATE"] = a.INSERTDATE
}

func (a aPPLYPROF) clone(db *gorm.DB) aPPLYPROF {
	a.aPPLYPROFDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aPPLYPROF) replaceDB(db *gorm.DB) aPPLYPROF {
	a.aPPLYPROFDo.ReplaceDB(db)
	return a
}

type aPPLYPROFDo struct{ gen.DO }

type IAPPLYPROFDo interface {
	gen.SubQuery
	Debug() IAPPLYPROFDo
	WithContext(ctx context.Context) IAPPLYPROFDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAPPLYPROFDo
	WriteDB() IAPPLYPROFDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAPPLYPROFDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAPPLYPROFDo
	Not(conds ...gen.Condition) IAPPLYPROFDo
	Or(conds ...gen.Condition) IAPPLYPROFDo
	Select(conds ...field.Expr) IAPPLYPROFDo
	Where(conds ...gen.Condition) IAPPLYPROFDo
	Order(conds ...field.Expr) IAPPLYPROFDo
	Distinct(cols ...field.Expr) IAPPLYPROFDo
	Omit(cols ...field.Expr) IAPPLYPROFDo
	Join(table schema.Tabler, on ...field.Expr) IAPPLYPROFDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAPPLYPROFDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAPPLYPROFDo
	Group(cols ...field.Expr) IAPPLYPROFDo
	Having(conds ...gen.Condition) IAPPLYPROFDo
	Limit(limit int) IAPPLYPROFDo
	Offset(offset int) IAPPLYPROFDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAPPLYPROFDo
	Unscoped() IAPPLYPROFDo
	Create(values ...*model.APPLYPROF) error
	CreateInBatches(values []*model.APPLYPROF, batchSize int) error
	Save(values ...*model.APPLYPROF) error
	First() (*model.APPLYPROF, error)
	Take() (*model.APPLYPROF, error)
	Last() (*model.APPLYPROF, error)
	Find() ([]*model.APPLYPROF, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.APPLYPROF, err error)
	FindInBatches(result *[]*model.APPLYPROF, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.APPLYPROF) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAPPLYPROFDo
	Assign(attrs ...field.AssignExpr) IAPPLYPROFDo
	Joins(fields ...field.RelationField) IAPPLYPROFDo
	Preload(fields ...field.RelationField) IAPPLYPROFDo
	FirstOrInit() (*model.APPLYPROF, error)
	FirstOrCreate() (*model.APPLYPROF, error)
	FindByPage(offset int, limit int) (result []*model.APPLYPROF, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAPPLYPROFDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aPPLYPROFDo) Debug() IAPPLYPROFDo {
	return a.withDO(a.DO.Debug())
}

func (a aPPLYPROFDo) WithContext(ctx context.Context) IAPPLYPROFDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aPPLYPROFDo) ReadDB() IAPPLYPROFDo {
	return a.Clauses(dbresolver.Read)
}

func (a aPPLYPROFDo) WriteDB() IAPPLYPROFDo {
	return a.Clauses(dbresolver.Write)
}

func (a aPPLYPROFDo) Session(config *gorm.Session) IAPPLYPROFDo {
	return a.withDO(a.DO.Session(config))
}

func (a aPPLYPROFDo) Clauses(conds ...clause.Expression) IAPPLYPROFDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aPPLYPROFDo) Returning(value interface{}, columns ...string) IAPPLYPROFDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aPPLYPROFDo) Not(conds ...gen.Condition) IAPPLYPROFDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aPPLYPROFDo) Or(conds ...gen.Condition) IAPPLYPROFDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aPPLYPROFDo) Select(conds ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aPPLYPROFDo) Where(conds ...gen.Condition) IAPPLYPROFDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aPPLYPROFDo) Order(conds ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aPPLYPROFDo) Distinct(cols ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aPPLYPROFDo) Omit(cols ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aPPLYPROFDo) Join(table schema.Tabler, on ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aPPLYPROFDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aPPLYPROFDo) RightJoin(table schema.Tabler, on ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aPPLYPROFDo) Group(cols ...field.Expr) IAPPLYPROFDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aPPLYPROFDo) Having(conds ...gen.Condition) IAPPLYPROFDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aPPLYPROFDo) Limit(limit int) IAPPLYPROFDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aPPLYPROFDo) Offset(offset int) IAPPLYPROFDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aPPLYPROFDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAPPLYPROFDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aPPLYPROFDo) Unscoped() IAPPLYPROFDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aPPLYPROFDo) Create(values ...*model.APPLYPROF) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aPPLYPROFDo) CreateInBatches(values []*model.APPLYPROF, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aPPLYPROFDo) Save(values ...*model.APPLYPROF) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aPPLYPROFDo) First() (*model.APPLYPROF, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYPROF), nil
	}
}

func (a aPPLYPROFDo) Take() (*model.APPLYPROF, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYPROF), nil
	}
}

func (a aPPLYPROFDo) Last() (*model.APPLYPROF, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYPROF), nil
	}
}

func (a aPPLYPROFDo) Find() ([]*model.APPLYPROF, error) {
	result, err := a.DO.Find()
	return result.([]*model.APPLYPROF), err
}

func (a aPPLYPROFDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.APPLYPROF, err error) {
	buf := make([]*model.APPLYPROF, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aPPLYPROFDo) FindInBatches(result *[]*model.APPLYPROF, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aPPLYPROFDo) Attrs(attrs ...field.AssignExpr) IAPPLYPROFDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aPPLYPROFDo) Assign(attrs ...field.AssignExpr) IAPPLYPROFDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aPPLYPROFDo) Joins(fields ...field.RelationField) IAPPLYPROFDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aPPLYPROFDo) Preload(fields ...field.RelationField) IAPPLYPROFDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aPPLYPROFDo) FirstOrInit() (*model.APPLYPROF, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYPROF), nil
	}
}

func (a aPPLYPROFDo) FirstOrCreate() (*model.APPLYPROF, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYPROF), nil
	}
}

func (a aPPLYPROFDo) FindByPage(offset int, limit int) (result []*model.APPLYPROF, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aPPLYPROFDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aPPLYPROFDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aPPLYPROFDo) Delete(models ...*model.APPLYPROF) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aPPLYPROFDo) withDO(do gen.Dao) *aPPLYPROFDo {
	a.DO = *do.(*gen.DO)
	return a
}
