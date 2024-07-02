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

func newAUTHGRPMENU(db *gorm.DB, opts ...gen.DOOption) aUTHGRPMENU {
	_aUTHGRPMENU := aUTHGRPMENU{}

	_aUTHGRPMENU.aUTHGRPMENUDo.UseDB(db, opts...)
	_aUTHGRPMENU.aUTHGRPMENUDo.UseModel(&model.AUTHGRPMENU{})

	tableName := _aUTHGRPMENU.aUTHGRPMENUDo.TableName()
	_aUTHGRPMENU.ALL = field.NewAsterisk(tableName)
	_aUTHGRPMENU.AUTHGRPCD = field.NewInt64(tableName, "AUTH_GRP_CD")
	_aUTHGRPMENU.MENUID = field.NewString(tableName, "MENU_ID")
	_aUTHGRPMENU.REGDT = field.NewString(tableName, "REG_DT")
	_aUTHGRPMENU.REGID = field.NewString(tableName, "REG_ID")
	_aUTHGRPMENU.MODDT = field.NewString(tableName, "MOD_DT")
	_aUTHGRPMENU.MODID = field.NewString(tableName, "MOD_ID")

	_aUTHGRPMENU.fillFieldMap()

	return _aUTHGRPMENU
}

type aUTHGRPMENU struct {
	aUTHGRPMENUDo

	ALL       field.Asterisk
	AUTHGRPCD field.Int64
	MENUID    field.String
	REGDT     field.String
	REGID     field.String
	MODDT     field.String
	MODID     field.String

	fieldMap map[string]field.Expr
}

func (a aUTHGRPMENU) Table(newTableName string) *aUTHGRPMENU {
	a.aUTHGRPMENUDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aUTHGRPMENU) As(alias string) *aUTHGRPMENU {
	a.aUTHGRPMENUDo.DO = *(a.aUTHGRPMENUDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aUTHGRPMENU) updateTableName(table string) *aUTHGRPMENU {
	a.ALL = field.NewAsterisk(table)
	a.AUTHGRPCD = field.NewInt64(table, "AUTH_GRP_CD")
	a.MENUID = field.NewString(table, "MENU_ID")
	a.REGDT = field.NewString(table, "REG_DT")
	a.REGID = field.NewString(table, "REG_ID")
	a.MODDT = field.NewString(table, "MOD_DT")
	a.MODID = field.NewString(table, "MOD_ID")

	a.fillFieldMap()

	return a
}

func (a *aUTHGRPMENU) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aUTHGRPMENU) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["AUTH_GRP_CD"] = a.AUTHGRPCD
	a.fieldMap["MENU_ID"] = a.MENUID
	a.fieldMap["REG_DT"] = a.REGDT
	a.fieldMap["REG_ID"] = a.REGID
	a.fieldMap["MOD_DT"] = a.MODDT
	a.fieldMap["MOD_ID"] = a.MODID
}

func (a aUTHGRPMENU) clone(db *gorm.DB) aUTHGRPMENU {
	a.aUTHGRPMENUDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aUTHGRPMENU) replaceDB(db *gorm.DB) aUTHGRPMENU {
	a.aUTHGRPMENUDo.ReplaceDB(db)
	return a
}

type aUTHGRPMENUDo struct{ gen.DO }

type IAUTHGRPMENUDo interface {
	gen.SubQuery
	Debug() IAUTHGRPMENUDo
	WithContext(ctx context.Context) IAUTHGRPMENUDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAUTHGRPMENUDo
	WriteDB() IAUTHGRPMENUDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAUTHGRPMENUDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAUTHGRPMENUDo
	Not(conds ...gen.Condition) IAUTHGRPMENUDo
	Or(conds ...gen.Condition) IAUTHGRPMENUDo
	Select(conds ...field.Expr) IAUTHGRPMENUDo
	Where(conds ...gen.Condition) IAUTHGRPMENUDo
	Order(conds ...field.Expr) IAUTHGRPMENUDo
	Distinct(cols ...field.Expr) IAUTHGRPMENUDo
	Omit(cols ...field.Expr) IAUTHGRPMENUDo
	Join(table schema.Tabler, on ...field.Expr) IAUTHGRPMENUDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAUTHGRPMENUDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAUTHGRPMENUDo
	Group(cols ...field.Expr) IAUTHGRPMENUDo
	Having(conds ...gen.Condition) IAUTHGRPMENUDo
	Limit(limit int) IAUTHGRPMENUDo
	Offset(offset int) IAUTHGRPMENUDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAUTHGRPMENUDo
	Unscoped() IAUTHGRPMENUDo
	Create(values ...*model.AUTHGRPMENU) error
	CreateInBatches(values []*model.AUTHGRPMENU, batchSize int) error
	Save(values ...*model.AUTHGRPMENU) error
	First() (*model.AUTHGRPMENU, error)
	Take() (*model.AUTHGRPMENU, error)
	Last() (*model.AUTHGRPMENU, error)
	Find() ([]*model.AUTHGRPMENU, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AUTHGRPMENU, err error)
	FindInBatches(result *[]*model.AUTHGRPMENU, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AUTHGRPMENU) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAUTHGRPMENUDo
	Assign(attrs ...field.AssignExpr) IAUTHGRPMENUDo
	Joins(fields ...field.RelationField) IAUTHGRPMENUDo
	Preload(fields ...field.RelationField) IAUTHGRPMENUDo
	FirstOrInit() (*model.AUTHGRPMENU, error)
	FirstOrCreate() (*model.AUTHGRPMENU, error)
	FindByPage(offset int, limit int) (result []*model.AUTHGRPMENU, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAUTHGRPMENUDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aUTHGRPMENUDo) Debug() IAUTHGRPMENUDo {
	return a.withDO(a.DO.Debug())
}

func (a aUTHGRPMENUDo) WithContext(ctx context.Context) IAUTHGRPMENUDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aUTHGRPMENUDo) ReadDB() IAUTHGRPMENUDo {
	return a.Clauses(dbresolver.Read)
}

func (a aUTHGRPMENUDo) WriteDB() IAUTHGRPMENUDo {
	return a.Clauses(dbresolver.Write)
}

func (a aUTHGRPMENUDo) Session(config *gorm.Session) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Session(config))
}

func (a aUTHGRPMENUDo) Clauses(conds ...clause.Expression) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aUTHGRPMENUDo) Returning(value interface{}, columns ...string) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aUTHGRPMENUDo) Not(conds ...gen.Condition) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aUTHGRPMENUDo) Or(conds ...gen.Condition) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aUTHGRPMENUDo) Select(conds ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aUTHGRPMENUDo) Where(conds ...gen.Condition) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aUTHGRPMENUDo) Order(conds ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aUTHGRPMENUDo) Distinct(cols ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aUTHGRPMENUDo) Omit(cols ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aUTHGRPMENUDo) Join(table schema.Tabler, on ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aUTHGRPMENUDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aUTHGRPMENUDo) RightJoin(table schema.Tabler, on ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aUTHGRPMENUDo) Group(cols ...field.Expr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aUTHGRPMENUDo) Having(conds ...gen.Condition) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aUTHGRPMENUDo) Limit(limit int) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aUTHGRPMENUDo) Offset(offset int) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aUTHGRPMENUDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aUTHGRPMENUDo) Unscoped() IAUTHGRPMENUDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aUTHGRPMENUDo) Create(values ...*model.AUTHGRPMENU) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aUTHGRPMENUDo) CreateInBatches(values []*model.AUTHGRPMENU, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aUTHGRPMENUDo) Save(values ...*model.AUTHGRPMENU) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aUTHGRPMENUDo) First() (*model.AUTHGRPMENU, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AUTHGRPMENU), nil
	}
}

func (a aUTHGRPMENUDo) Take() (*model.AUTHGRPMENU, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AUTHGRPMENU), nil
	}
}

func (a aUTHGRPMENUDo) Last() (*model.AUTHGRPMENU, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AUTHGRPMENU), nil
	}
}

func (a aUTHGRPMENUDo) Find() ([]*model.AUTHGRPMENU, error) {
	result, err := a.DO.Find()
	return result.([]*model.AUTHGRPMENU), err
}

func (a aUTHGRPMENUDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AUTHGRPMENU, err error) {
	buf := make([]*model.AUTHGRPMENU, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aUTHGRPMENUDo) FindInBatches(result *[]*model.AUTHGRPMENU, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aUTHGRPMENUDo) Attrs(attrs ...field.AssignExpr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aUTHGRPMENUDo) Assign(attrs ...field.AssignExpr) IAUTHGRPMENUDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aUTHGRPMENUDo) Joins(fields ...field.RelationField) IAUTHGRPMENUDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aUTHGRPMENUDo) Preload(fields ...field.RelationField) IAUTHGRPMENUDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aUTHGRPMENUDo) FirstOrInit() (*model.AUTHGRPMENU, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AUTHGRPMENU), nil
	}
}

func (a aUTHGRPMENUDo) FirstOrCreate() (*model.AUTHGRPMENU, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AUTHGRPMENU), nil
	}
}

func (a aUTHGRPMENUDo) FindByPage(offset int, limit int) (result []*model.AUTHGRPMENU, count int64, err error) {
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

func (a aUTHGRPMENUDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aUTHGRPMENUDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aUTHGRPMENUDo) Delete(models ...*model.AUTHGRPMENU) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aUTHGRPMENUDo) withDO(do gen.Dao) *aUTHGRPMENUDo {
	a.DO = *do.(*gen.DO)
	return a
}
