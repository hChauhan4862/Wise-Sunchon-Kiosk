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

func newAPPLYROOM(db *gorm.DB, opts ...gen.DOOption) aPPLYROOM {
	_aPPLYROOM := aPPLYROOM{}

	_aPPLYROOM.aPPLYROOMDo.UseDB(db, opts...)
	_aPPLYROOM.aPPLYROOMDo.UseModel(&model.APPLYROOM{})

	tableName := _aPPLYROOM.aPPLYROOMDo.TableName()
	_aPPLYROOM.ALL = field.NewAsterisk(tableName)
	_aPPLYROOM.KIOSKNO = field.NewString(tableName, "KIOSK_NO")
	_aPPLYROOM.ROOMNAME = field.NewString(tableName, "ROOM_NAME")
	_aPPLYROOM.SEATNOS = field.NewString(tableName, "SEAT_NO_S")
	_aPPLYROOM.SEATNOE = field.NewString(tableName, "SEAT_NO_E")
	_aPPLYROOM.SEATNOCNT = field.NewString(tableName, "SEAT_NO_CNT")
	_aPPLYROOM.SEATFIX = field.NewString(tableName, "SEAT_FIX")
	_aPPLYROOM.USEYN = field.NewString(tableName, "USE_YN")
	_aPPLYROOM.INSERTDATE = field.NewString(tableName, "INSERT_DATE")
	_aPPLYROOM.UPDATEDATE = field.NewString(tableName, "UPDATE_DATE")
	_aPPLYROOM.BIGO = field.NewString(tableName, "BIGO")

	_aPPLYROOM.fillFieldMap()

	return _aPPLYROOM
}

type aPPLYROOM struct {
	aPPLYROOMDo

	ALL        field.Asterisk
	KIOSKNO    field.String
	ROOMNAME   field.String
	SEATNOS    field.String
	SEATNOE    field.String
	SEATNOCNT  field.String
	SEATFIX    field.String
	USEYN      field.String
	INSERTDATE field.String
	UPDATEDATE field.String
	BIGO       field.String

	fieldMap map[string]field.Expr
}

func (a aPPLYROOM) Table(newTableName string) *aPPLYROOM {
	a.aPPLYROOMDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aPPLYROOM) As(alias string) *aPPLYROOM {
	a.aPPLYROOMDo.DO = *(a.aPPLYROOMDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aPPLYROOM) updateTableName(table string) *aPPLYROOM {
	a.ALL = field.NewAsterisk(table)
	a.KIOSKNO = field.NewString(table, "KIOSK_NO")
	a.ROOMNAME = field.NewString(table, "ROOM_NAME")
	a.SEATNOS = field.NewString(table, "SEAT_NO_S")
	a.SEATNOE = field.NewString(table, "SEAT_NO_E")
	a.SEATNOCNT = field.NewString(table, "SEAT_NO_CNT")
	a.SEATFIX = field.NewString(table, "SEAT_FIX")
	a.USEYN = field.NewString(table, "USE_YN")
	a.INSERTDATE = field.NewString(table, "INSERT_DATE")
	a.UPDATEDATE = field.NewString(table, "UPDATE_DATE")
	a.BIGO = field.NewString(table, "BIGO")

	a.fillFieldMap()

	return a
}

func (a *aPPLYROOM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aPPLYROOM) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["KIOSK_NO"] = a.KIOSKNO
	a.fieldMap["ROOM_NAME"] = a.ROOMNAME
	a.fieldMap["SEAT_NO_S"] = a.SEATNOS
	a.fieldMap["SEAT_NO_E"] = a.SEATNOE
	a.fieldMap["SEAT_NO_CNT"] = a.SEATNOCNT
	a.fieldMap["SEAT_FIX"] = a.SEATFIX
	a.fieldMap["USE_YN"] = a.USEYN
	a.fieldMap["INSERT_DATE"] = a.INSERTDATE
	a.fieldMap["UPDATE_DATE"] = a.UPDATEDATE
	a.fieldMap["BIGO"] = a.BIGO
}

func (a aPPLYROOM) clone(db *gorm.DB) aPPLYROOM {
	a.aPPLYROOMDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aPPLYROOM) replaceDB(db *gorm.DB) aPPLYROOM {
	a.aPPLYROOMDo.ReplaceDB(db)
	return a
}

type aPPLYROOMDo struct{ gen.DO }

type IAPPLYROOMDo interface {
	gen.SubQuery
	Debug() IAPPLYROOMDo
	WithContext(ctx context.Context) IAPPLYROOMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAPPLYROOMDo
	WriteDB() IAPPLYROOMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAPPLYROOMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAPPLYROOMDo
	Not(conds ...gen.Condition) IAPPLYROOMDo
	Or(conds ...gen.Condition) IAPPLYROOMDo
	Select(conds ...field.Expr) IAPPLYROOMDo
	Where(conds ...gen.Condition) IAPPLYROOMDo
	Order(conds ...field.Expr) IAPPLYROOMDo
	Distinct(cols ...field.Expr) IAPPLYROOMDo
	Omit(cols ...field.Expr) IAPPLYROOMDo
	Join(table schema.Tabler, on ...field.Expr) IAPPLYROOMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAPPLYROOMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAPPLYROOMDo
	Group(cols ...field.Expr) IAPPLYROOMDo
	Having(conds ...gen.Condition) IAPPLYROOMDo
	Limit(limit int) IAPPLYROOMDo
	Offset(offset int) IAPPLYROOMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAPPLYROOMDo
	Unscoped() IAPPLYROOMDo
	Create(values ...*model.APPLYROOM) error
	CreateInBatches(values []*model.APPLYROOM, batchSize int) error
	Save(values ...*model.APPLYROOM) error
	First() (*model.APPLYROOM, error)
	Take() (*model.APPLYROOM, error)
	Last() (*model.APPLYROOM, error)
	Find() ([]*model.APPLYROOM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.APPLYROOM, err error)
	FindInBatches(result *[]*model.APPLYROOM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.APPLYROOM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAPPLYROOMDo
	Assign(attrs ...field.AssignExpr) IAPPLYROOMDo
	Joins(fields ...field.RelationField) IAPPLYROOMDo
	Preload(fields ...field.RelationField) IAPPLYROOMDo
	FirstOrInit() (*model.APPLYROOM, error)
	FirstOrCreate() (*model.APPLYROOM, error)
	FindByPage(offset int, limit int) (result []*model.APPLYROOM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAPPLYROOMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aPPLYROOMDo) Debug() IAPPLYROOMDo {
	return a.withDO(a.DO.Debug())
}

func (a aPPLYROOMDo) WithContext(ctx context.Context) IAPPLYROOMDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aPPLYROOMDo) ReadDB() IAPPLYROOMDo {
	return a.Clauses(dbresolver.Read)
}

func (a aPPLYROOMDo) WriteDB() IAPPLYROOMDo {
	return a.Clauses(dbresolver.Write)
}

func (a aPPLYROOMDo) Session(config *gorm.Session) IAPPLYROOMDo {
	return a.withDO(a.DO.Session(config))
}

func (a aPPLYROOMDo) Clauses(conds ...clause.Expression) IAPPLYROOMDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aPPLYROOMDo) Returning(value interface{}, columns ...string) IAPPLYROOMDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aPPLYROOMDo) Not(conds ...gen.Condition) IAPPLYROOMDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aPPLYROOMDo) Or(conds ...gen.Condition) IAPPLYROOMDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aPPLYROOMDo) Select(conds ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aPPLYROOMDo) Where(conds ...gen.Condition) IAPPLYROOMDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aPPLYROOMDo) Order(conds ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aPPLYROOMDo) Distinct(cols ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aPPLYROOMDo) Omit(cols ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aPPLYROOMDo) Join(table schema.Tabler, on ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aPPLYROOMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aPPLYROOMDo) RightJoin(table schema.Tabler, on ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aPPLYROOMDo) Group(cols ...field.Expr) IAPPLYROOMDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aPPLYROOMDo) Having(conds ...gen.Condition) IAPPLYROOMDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aPPLYROOMDo) Limit(limit int) IAPPLYROOMDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aPPLYROOMDo) Offset(offset int) IAPPLYROOMDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aPPLYROOMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAPPLYROOMDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aPPLYROOMDo) Unscoped() IAPPLYROOMDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aPPLYROOMDo) Create(values ...*model.APPLYROOM) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aPPLYROOMDo) CreateInBatches(values []*model.APPLYROOM, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aPPLYROOMDo) Save(values ...*model.APPLYROOM) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aPPLYROOMDo) First() (*model.APPLYROOM, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYROOM), nil
	}
}

func (a aPPLYROOMDo) Take() (*model.APPLYROOM, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYROOM), nil
	}
}

func (a aPPLYROOMDo) Last() (*model.APPLYROOM, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYROOM), nil
	}
}

func (a aPPLYROOMDo) Find() ([]*model.APPLYROOM, error) {
	result, err := a.DO.Find()
	return result.([]*model.APPLYROOM), err
}

func (a aPPLYROOMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.APPLYROOM, err error) {
	buf := make([]*model.APPLYROOM, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aPPLYROOMDo) FindInBatches(result *[]*model.APPLYROOM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aPPLYROOMDo) Attrs(attrs ...field.AssignExpr) IAPPLYROOMDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aPPLYROOMDo) Assign(attrs ...field.AssignExpr) IAPPLYROOMDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aPPLYROOMDo) Joins(fields ...field.RelationField) IAPPLYROOMDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aPPLYROOMDo) Preload(fields ...field.RelationField) IAPPLYROOMDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aPPLYROOMDo) FirstOrInit() (*model.APPLYROOM, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYROOM), nil
	}
}

func (a aPPLYROOMDo) FirstOrCreate() (*model.APPLYROOM, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.APPLYROOM), nil
	}
}

func (a aPPLYROOMDo) FindByPage(offset int, limit int) (result []*model.APPLYROOM, count int64, err error) {
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

func (a aPPLYROOMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aPPLYROOMDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aPPLYROOMDo) Delete(models ...*model.APPLYROOM) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aPPLYROOMDo) withDO(do gen.Dao) *aPPLYROOMDo {
	a.DO = *do.(*gen.DO)
	return a
}
