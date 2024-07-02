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

func newEQUIPTYPE(db *gorm.DB, opts ...gen.DOOption) eQUIPTYPE {
	_eQUIPTYPE := eQUIPTYPE{}

	_eQUIPTYPE.eQUIPTYPEDo.UseDB(db, opts...)
	_eQUIPTYPE.eQUIPTYPEDo.UseModel(&model.EQUIPTYPE{})

	tableName := _eQUIPTYPE.eQUIPTYPEDo.TableName()
	_eQUIPTYPE.ALL = field.NewAsterisk(tableName)
	_eQUIPTYPE.EQTYPE = field.NewInt64(tableName, "EQTYPE")
	_eQUIPTYPE.NAME = field.NewString(tableName, "NAME")
	_eQUIPTYPE.ENNAME = field.NewString(tableName, "EN_NAME")
	_eQUIPTYPE.DETAIL = field.NewString(tableName, "DETAIL")
	_eQUIPTYPE.ISDELETE = field.NewInt64(tableName, "ISDELETE")

	_eQUIPTYPE.fillFieldMap()

	return _eQUIPTYPE
}

type eQUIPTYPE struct {
	eQUIPTYPEDo

	ALL      field.Asterisk
	EQTYPE   field.Int64
	NAME     field.String
	ENNAME   field.String
	DETAIL   field.String
	ISDELETE field.Int64

	fieldMap map[string]field.Expr
}

func (e eQUIPTYPE) Table(newTableName string) *eQUIPTYPE {
	e.eQUIPTYPEDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e eQUIPTYPE) As(alias string) *eQUIPTYPE {
	e.eQUIPTYPEDo.DO = *(e.eQUIPTYPEDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *eQUIPTYPE) updateTableName(table string) *eQUIPTYPE {
	e.ALL = field.NewAsterisk(table)
	e.EQTYPE = field.NewInt64(table, "EQTYPE")
	e.NAME = field.NewString(table, "NAME")
	e.ENNAME = field.NewString(table, "EN_NAME")
	e.DETAIL = field.NewString(table, "DETAIL")
	e.ISDELETE = field.NewInt64(table, "ISDELETE")

	e.fillFieldMap()

	return e
}

func (e *eQUIPTYPE) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *eQUIPTYPE) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 5)
	e.fieldMap["EQTYPE"] = e.EQTYPE
	e.fieldMap["NAME"] = e.NAME
	e.fieldMap["EN_NAME"] = e.ENNAME
	e.fieldMap["DETAIL"] = e.DETAIL
	e.fieldMap["ISDELETE"] = e.ISDELETE
}

func (e eQUIPTYPE) clone(db *gorm.DB) eQUIPTYPE {
	e.eQUIPTYPEDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e eQUIPTYPE) replaceDB(db *gorm.DB) eQUIPTYPE {
	e.eQUIPTYPEDo.ReplaceDB(db)
	return e
}

type eQUIPTYPEDo struct{ gen.DO }

type IEQUIPTYPEDo interface {
	gen.SubQuery
	Debug() IEQUIPTYPEDo
	WithContext(ctx context.Context) IEQUIPTYPEDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEQUIPTYPEDo
	WriteDB() IEQUIPTYPEDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEQUIPTYPEDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEQUIPTYPEDo
	Not(conds ...gen.Condition) IEQUIPTYPEDo
	Or(conds ...gen.Condition) IEQUIPTYPEDo
	Select(conds ...field.Expr) IEQUIPTYPEDo
	Where(conds ...gen.Condition) IEQUIPTYPEDo
	Order(conds ...field.Expr) IEQUIPTYPEDo
	Distinct(cols ...field.Expr) IEQUIPTYPEDo
	Omit(cols ...field.Expr) IEQUIPTYPEDo
	Join(table schema.Tabler, on ...field.Expr) IEQUIPTYPEDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEQUIPTYPEDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEQUIPTYPEDo
	Group(cols ...field.Expr) IEQUIPTYPEDo
	Having(conds ...gen.Condition) IEQUIPTYPEDo
	Limit(limit int) IEQUIPTYPEDo
	Offset(offset int) IEQUIPTYPEDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEQUIPTYPEDo
	Unscoped() IEQUIPTYPEDo
	Create(values ...*model.EQUIPTYPE) error
	CreateInBatches(values []*model.EQUIPTYPE, batchSize int) error
	Save(values ...*model.EQUIPTYPE) error
	First() (*model.EQUIPTYPE, error)
	Take() (*model.EQUIPTYPE, error)
	Last() (*model.EQUIPTYPE, error)
	Find() ([]*model.EQUIPTYPE, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EQUIPTYPE, err error)
	FindInBatches(result *[]*model.EQUIPTYPE, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EQUIPTYPE) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEQUIPTYPEDo
	Assign(attrs ...field.AssignExpr) IEQUIPTYPEDo
	Joins(fields ...field.RelationField) IEQUIPTYPEDo
	Preload(fields ...field.RelationField) IEQUIPTYPEDo
	FirstOrInit() (*model.EQUIPTYPE, error)
	FirstOrCreate() (*model.EQUIPTYPE, error)
	FindByPage(offset int, limit int) (result []*model.EQUIPTYPE, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEQUIPTYPEDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e eQUIPTYPEDo) Debug() IEQUIPTYPEDo {
	return e.withDO(e.DO.Debug())
}

func (e eQUIPTYPEDo) WithContext(ctx context.Context) IEQUIPTYPEDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e eQUIPTYPEDo) ReadDB() IEQUIPTYPEDo {
	return e.Clauses(dbresolver.Read)
}

func (e eQUIPTYPEDo) WriteDB() IEQUIPTYPEDo {
	return e.Clauses(dbresolver.Write)
}

func (e eQUIPTYPEDo) Session(config *gorm.Session) IEQUIPTYPEDo {
	return e.withDO(e.DO.Session(config))
}

func (e eQUIPTYPEDo) Clauses(conds ...clause.Expression) IEQUIPTYPEDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e eQUIPTYPEDo) Returning(value interface{}, columns ...string) IEQUIPTYPEDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e eQUIPTYPEDo) Not(conds ...gen.Condition) IEQUIPTYPEDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e eQUIPTYPEDo) Or(conds ...gen.Condition) IEQUIPTYPEDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e eQUIPTYPEDo) Select(conds ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e eQUIPTYPEDo) Where(conds ...gen.Condition) IEQUIPTYPEDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e eQUIPTYPEDo) Order(conds ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e eQUIPTYPEDo) Distinct(cols ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e eQUIPTYPEDo) Omit(cols ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e eQUIPTYPEDo) Join(table schema.Tabler, on ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e eQUIPTYPEDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e eQUIPTYPEDo) RightJoin(table schema.Tabler, on ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e eQUIPTYPEDo) Group(cols ...field.Expr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e eQUIPTYPEDo) Having(conds ...gen.Condition) IEQUIPTYPEDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e eQUIPTYPEDo) Limit(limit int) IEQUIPTYPEDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e eQUIPTYPEDo) Offset(offset int) IEQUIPTYPEDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e eQUIPTYPEDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEQUIPTYPEDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e eQUIPTYPEDo) Unscoped() IEQUIPTYPEDo {
	return e.withDO(e.DO.Unscoped())
}

func (e eQUIPTYPEDo) Create(values ...*model.EQUIPTYPE) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e eQUIPTYPEDo) CreateInBatches(values []*model.EQUIPTYPE, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e eQUIPTYPEDo) Save(values ...*model.EQUIPTYPE) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e eQUIPTYPEDo) First() (*model.EQUIPTYPE, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EQUIPTYPE), nil
	}
}

func (e eQUIPTYPEDo) Take() (*model.EQUIPTYPE, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EQUIPTYPE), nil
	}
}

func (e eQUIPTYPEDo) Last() (*model.EQUIPTYPE, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EQUIPTYPE), nil
	}
}

func (e eQUIPTYPEDo) Find() ([]*model.EQUIPTYPE, error) {
	result, err := e.DO.Find()
	return result.([]*model.EQUIPTYPE), err
}

func (e eQUIPTYPEDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EQUIPTYPE, err error) {
	buf := make([]*model.EQUIPTYPE, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e eQUIPTYPEDo) FindInBatches(result *[]*model.EQUIPTYPE, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e eQUIPTYPEDo) Attrs(attrs ...field.AssignExpr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e eQUIPTYPEDo) Assign(attrs ...field.AssignExpr) IEQUIPTYPEDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e eQUIPTYPEDo) Joins(fields ...field.RelationField) IEQUIPTYPEDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e eQUIPTYPEDo) Preload(fields ...field.RelationField) IEQUIPTYPEDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e eQUIPTYPEDo) FirstOrInit() (*model.EQUIPTYPE, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EQUIPTYPE), nil
	}
}

func (e eQUIPTYPEDo) FirstOrCreate() (*model.EQUIPTYPE, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EQUIPTYPE), nil
	}
}

func (e eQUIPTYPEDo) FindByPage(offset int, limit int) (result []*model.EQUIPTYPE, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e eQUIPTYPEDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e eQUIPTYPEDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e eQUIPTYPEDo) Delete(models ...*model.EQUIPTYPE) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *eQUIPTYPEDo) withDO(do gen.Dao) *eQUIPTYPEDo {
	e.DO = *do.(*gen.DO)
	return e
}