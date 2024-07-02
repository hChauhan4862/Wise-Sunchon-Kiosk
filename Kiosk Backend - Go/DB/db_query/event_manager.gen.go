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

func newEVENTMANAGER(db *gorm.DB, opts ...gen.DOOption) eVENTMANAGER {
	_eVENTMANAGER := eVENTMANAGER{}

	_eVENTMANAGER.eVENTMANAGERDo.UseDB(db, opts...)
	_eVENTMANAGER.eVENTMANAGERDo.UseModel(&model.EVENTMANAGER{})

	tableName := _eVENTMANAGER.eVENTMANAGERDo.TableName()
	_eVENTMANAGER.ALL = field.NewAsterisk(tableName)
	_eVENTMANAGER.EventID = field.NewInt64(tableName, "event_id")
	_eVENTMANAGER.EventTitle = field.NewString(tableName, "event_title")
	_eVENTMANAGER.Startdate = field.NewString(tableName, "startdate")
	_eVENTMANAGER.Enddate = field.NewString(tableName, "enddate")
	_eVENTMANAGER.UseYn = field.NewString(tableName, "use_yn")
	_eVENTMANAGER.RegDt = field.NewString(tableName, "reg_dt")

	_eVENTMANAGER.fillFieldMap()

	return _eVENTMANAGER
}

type eVENTMANAGER struct {
	eVENTMANAGERDo

	ALL        field.Asterisk
	EventID    field.Int64
	EventTitle field.String
	Startdate  field.String
	Enddate    field.String
	UseYn      field.String
	RegDt      field.String

	fieldMap map[string]field.Expr
}

func (e eVENTMANAGER) Table(newTableName string) *eVENTMANAGER {
	e.eVENTMANAGERDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e eVENTMANAGER) As(alias string) *eVENTMANAGER {
	e.eVENTMANAGERDo.DO = *(e.eVENTMANAGERDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *eVENTMANAGER) updateTableName(table string) *eVENTMANAGER {
	e.ALL = field.NewAsterisk(table)
	e.EventID = field.NewInt64(table, "event_id")
	e.EventTitle = field.NewString(table, "event_title")
	e.Startdate = field.NewString(table, "startdate")
	e.Enddate = field.NewString(table, "enddate")
	e.UseYn = field.NewString(table, "use_yn")
	e.RegDt = field.NewString(table, "reg_dt")

	e.fillFieldMap()

	return e
}

func (e *eVENTMANAGER) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *eVENTMANAGER) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["event_id"] = e.EventID
	e.fieldMap["event_title"] = e.EventTitle
	e.fieldMap["startdate"] = e.Startdate
	e.fieldMap["enddate"] = e.Enddate
	e.fieldMap["use_yn"] = e.UseYn
	e.fieldMap["reg_dt"] = e.RegDt
}

func (e eVENTMANAGER) clone(db *gorm.DB) eVENTMANAGER {
	e.eVENTMANAGERDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e eVENTMANAGER) replaceDB(db *gorm.DB) eVENTMANAGER {
	e.eVENTMANAGERDo.ReplaceDB(db)
	return e
}

type eVENTMANAGERDo struct{ gen.DO }

type IEVENTMANAGERDo interface {
	gen.SubQuery
	Debug() IEVENTMANAGERDo
	WithContext(ctx context.Context) IEVENTMANAGERDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEVENTMANAGERDo
	WriteDB() IEVENTMANAGERDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEVENTMANAGERDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEVENTMANAGERDo
	Not(conds ...gen.Condition) IEVENTMANAGERDo
	Or(conds ...gen.Condition) IEVENTMANAGERDo
	Select(conds ...field.Expr) IEVENTMANAGERDo
	Where(conds ...gen.Condition) IEVENTMANAGERDo
	Order(conds ...field.Expr) IEVENTMANAGERDo
	Distinct(cols ...field.Expr) IEVENTMANAGERDo
	Omit(cols ...field.Expr) IEVENTMANAGERDo
	Join(table schema.Tabler, on ...field.Expr) IEVENTMANAGERDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEVENTMANAGERDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEVENTMANAGERDo
	Group(cols ...field.Expr) IEVENTMANAGERDo
	Having(conds ...gen.Condition) IEVENTMANAGERDo
	Limit(limit int) IEVENTMANAGERDo
	Offset(offset int) IEVENTMANAGERDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEVENTMANAGERDo
	Unscoped() IEVENTMANAGERDo
	Create(values ...*model.EVENTMANAGER) error
	CreateInBatches(values []*model.EVENTMANAGER, batchSize int) error
	Save(values ...*model.EVENTMANAGER) error
	First() (*model.EVENTMANAGER, error)
	Take() (*model.EVENTMANAGER, error)
	Last() (*model.EVENTMANAGER, error)
	Find() ([]*model.EVENTMANAGER, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EVENTMANAGER, err error)
	FindInBatches(result *[]*model.EVENTMANAGER, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EVENTMANAGER) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEVENTMANAGERDo
	Assign(attrs ...field.AssignExpr) IEVENTMANAGERDo
	Joins(fields ...field.RelationField) IEVENTMANAGERDo
	Preload(fields ...field.RelationField) IEVENTMANAGERDo
	FirstOrInit() (*model.EVENTMANAGER, error)
	FirstOrCreate() (*model.EVENTMANAGER, error)
	FindByPage(offset int, limit int) (result []*model.EVENTMANAGER, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEVENTMANAGERDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e eVENTMANAGERDo) Debug() IEVENTMANAGERDo {
	return e.withDO(e.DO.Debug())
}

func (e eVENTMANAGERDo) WithContext(ctx context.Context) IEVENTMANAGERDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e eVENTMANAGERDo) ReadDB() IEVENTMANAGERDo {
	return e.Clauses(dbresolver.Read)
}

func (e eVENTMANAGERDo) WriteDB() IEVENTMANAGERDo {
	return e.Clauses(dbresolver.Write)
}

func (e eVENTMANAGERDo) Session(config *gorm.Session) IEVENTMANAGERDo {
	return e.withDO(e.DO.Session(config))
}

func (e eVENTMANAGERDo) Clauses(conds ...clause.Expression) IEVENTMANAGERDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e eVENTMANAGERDo) Returning(value interface{}, columns ...string) IEVENTMANAGERDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e eVENTMANAGERDo) Not(conds ...gen.Condition) IEVENTMANAGERDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e eVENTMANAGERDo) Or(conds ...gen.Condition) IEVENTMANAGERDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e eVENTMANAGERDo) Select(conds ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e eVENTMANAGERDo) Where(conds ...gen.Condition) IEVENTMANAGERDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e eVENTMANAGERDo) Order(conds ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e eVENTMANAGERDo) Distinct(cols ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e eVENTMANAGERDo) Omit(cols ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e eVENTMANAGERDo) Join(table schema.Tabler, on ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e eVENTMANAGERDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e eVENTMANAGERDo) RightJoin(table schema.Tabler, on ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e eVENTMANAGERDo) Group(cols ...field.Expr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e eVENTMANAGERDo) Having(conds ...gen.Condition) IEVENTMANAGERDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e eVENTMANAGERDo) Limit(limit int) IEVENTMANAGERDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e eVENTMANAGERDo) Offset(offset int) IEVENTMANAGERDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e eVENTMANAGERDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEVENTMANAGERDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e eVENTMANAGERDo) Unscoped() IEVENTMANAGERDo {
	return e.withDO(e.DO.Unscoped())
}

func (e eVENTMANAGERDo) Create(values ...*model.EVENTMANAGER) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e eVENTMANAGERDo) CreateInBatches(values []*model.EVENTMANAGER, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e eVENTMANAGERDo) Save(values ...*model.EVENTMANAGER) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e eVENTMANAGERDo) First() (*model.EVENTMANAGER, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EVENTMANAGER), nil
	}
}

func (e eVENTMANAGERDo) Take() (*model.EVENTMANAGER, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EVENTMANAGER), nil
	}
}

func (e eVENTMANAGERDo) Last() (*model.EVENTMANAGER, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EVENTMANAGER), nil
	}
}

func (e eVENTMANAGERDo) Find() ([]*model.EVENTMANAGER, error) {
	result, err := e.DO.Find()
	return result.([]*model.EVENTMANAGER), err
}

func (e eVENTMANAGERDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EVENTMANAGER, err error) {
	buf := make([]*model.EVENTMANAGER, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e eVENTMANAGERDo) FindInBatches(result *[]*model.EVENTMANAGER, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e eVENTMANAGERDo) Attrs(attrs ...field.AssignExpr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e eVENTMANAGERDo) Assign(attrs ...field.AssignExpr) IEVENTMANAGERDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e eVENTMANAGERDo) Joins(fields ...field.RelationField) IEVENTMANAGERDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e eVENTMANAGERDo) Preload(fields ...field.RelationField) IEVENTMANAGERDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e eVENTMANAGERDo) FirstOrInit() (*model.EVENTMANAGER, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EVENTMANAGER), nil
	}
}

func (e eVENTMANAGERDo) FirstOrCreate() (*model.EVENTMANAGER, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EVENTMANAGER), nil
	}
}

func (e eVENTMANAGERDo) FindByPage(offset int, limit int) (result []*model.EVENTMANAGER, count int64, err error) {
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

func (e eVENTMANAGERDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e eVENTMANAGERDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e eVENTMANAGERDo) Delete(models ...*model.EVENTMANAGER) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *eVENTMANAGERDo) withDO(do gen.Dao) *eVENTMANAGERDo {
	e.DO = *do.(*gen.DO)
	return e
}
