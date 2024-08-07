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

func newADMINMASTER(db *gorm.DB, opts ...gen.DOOption) aDMINMASTER {
	_aDMINMASTER := aDMINMASTER{}

	_aDMINMASTER.aDMINMASTERDo.UseDB(db, opts...)
	_aDMINMASTER.aDMINMASTERDo.UseModel(&model.ADMINMASTER{})

	tableName := _aDMINMASTER.aDMINMASTERDo.TableName()
	_aDMINMASTER.ALL = field.NewAsterisk(tableName)
	_aDMINMASTER.ADMINID = field.NewString(tableName, "ADMIN_ID")
	_aDMINMASTER.ADMINNM = field.NewString(tableName, "ADMIN_NM")
	_aDMINMASTER.PASSWD = field.NewString(tableName, "PASSWD")
	_aDMINMASTER.AUTHGRPCD = field.NewInt64(tableName, "AUTH_GRP_CD")
	_aDMINMASTER.TEL = field.NewString(tableName, "TEL")
	_aDMINMASTER.DEPT = field.NewString(tableName, "DEPT")
	_aDMINMASTER.JIKWIE = field.NewString(tableName, "JIKWIE")
	_aDMINMASTER.EMAIL = field.NewString(tableName, "EMAIL")
	_aDMINMASTER.HP = field.NewString(tableName, "HP")
	_aDMINMASTER.LOGINIP = field.NewString(tableName, "LOGIN_IP")
	_aDMINMASTER.REGDT = field.NewString(tableName, "REG_DT")
	_aDMINMASTER.REGID = field.NewString(tableName, "REG_ID")
	_aDMINMASTER.MODDT = field.NewString(tableName, "MOD_DT")
	_aDMINMASTER.MODID = field.NewString(tableName, "MOD_ID")

	_aDMINMASTER.fillFieldMap()

	return _aDMINMASTER
}

type aDMINMASTER struct {
	aDMINMASTERDo

	ALL       field.Asterisk
	ADMINID   field.String
	ADMINNM   field.String
	PASSWD    field.String
	AUTHGRPCD field.Int64
	TEL       field.String
	DEPT      field.String
	JIKWIE    field.String
	EMAIL     field.String
	HP        field.String
	LOGINIP   field.String
	REGDT     field.String
	REGID     field.String
	MODDT     field.String
	MODID     field.String

	fieldMap map[string]field.Expr
}

func (a aDMINMASTER) Table(newTableName string) *aDMINMASTER {
	a.aDMINMASTERDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aDMINMASTER) As(alias string) *aDMINMASTER {
	a.aDMINMASTERDo.DO = *(a.aDMINMASTERDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aDMINMASTER) updateTableName(table string) *aDMINMASTER {
	a.ALL = field.NewAsterisk(table)
	a.ADMINID = field.NewString(table, "ADMIN_ID")
	a.ADMINNM = field.NewString(table, "ADMIN_NM")
	a.PASSWD = field.NewString(table, "PASSWD")
	a.AUTHGRPCD = field.NewInt64(table, "AUTH_GRP_CD")
	a.TEL = field.NewString(table, "TEL")
	a.DEPT = field.NewString(table, "DEPT")
	a.JIKWIE = field.NewString(table, "JIKWIE")
	a.EMAIL = field.NewString(table, "EMAIL")
	a.HP = field.NewString(table, "HP")
	a.LOGINIP = field.NewString(table, "LOGIN_IP")
	a.REGDT = field.NewString(table, "REG_DT")
	a.REGID = field.NewString(table, "REG_ID")
	a.MODDT = field.NewString(table, "MOD_DT")
	a.MODID = field.NewString(table, "MOD_ID")

	a.fillFieldMap()

	return a
}

func (a *aDMINMASTER) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aDMINMASTER) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["ADMIN_ID"] = a.ADMINID
	a.fieldMap["ADMIN_NM"] = a.ADMINNM
	a.fieldMap["PASSWD"] = a.PASSWD
	a.fieldMap["AUTH_GRP_CD"] = a.AUTHGRPCD
	a.fieldMap["TEL"] = a.TEL
	a.fieldMap["DEPT"] = a.DEPT
	a.fieldMap["JIKWIE"] = a.JIKWIE
	a.fieldMap["EMAIL"] = a.EMAIL
	a.fieldMap["HP"] = a.HP
	a.fieldMap["LOGIN_IP"] = a.LOGINIP
	a.fieldMap["REG_DT"] = a.REGDT
	a.fieldMap["REG_ID"] = a.REGID
	a.fieldMap["MOD_DT"] = a.MODDT
	a.fieldMap["MOD_ID"] = a.MODID
}

func (a aDMINMASTER) clone(db *gorm.DB) aDMINMASTER {
	a.aDMINMASTERDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aDMINMASTER) replaceDB(db *gorm.DB) aDMINMASTER {
	a.aDMINMASTERDo.ReplaceDB(db)
	return a
}

type aDMINMASTERDo struct{ gen.DO }

type IADMINMASTERDo interface {
	gen.SubQuery
	Debug() IADMINMASTERDo
	WithContext(ctx context.Context) IADMINMASTERDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IADMINMASTERDo
	WriteDB() IADMINMASTERDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IADMINMASTERDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IADMINMASTERDo
	Not(conds ...gen.Condition) IADMINMASTERDo
	Or(conds ...gen.Condition) IADMINMASTERDo
	Select(conds ...field.Expr) IADMINMASTERDo
	Where(conds ...gen.Condition) IADMINMASTERDo
	Order(conds ...field.Expr) IADMINMASTERDo
	Distinct(cols ...field.Expr) IADMINMASTERDo
	Omit(cols ...field.Expr) IADMINMASTERDo
	Join(table schema.Tabler, on ...field.Expr) IADMINMASTERDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IADMINMASTERDo
	RightJoin(table schema.Tabler, on ...field.Expr) IADMINMASTERDo
	Group(cols ...field.Expr) IADMINMASTERDo
	Having(conds ...gen.Condition) IADMINMASTERDo
	Limit(limit int) IADMINMASTERDo
	Offset(offset int) IADMINMASTERDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IADMINMASTERDo
	Unscoped() IADMINMASTERDo
	Create(values ...*model.ADMINMASTER) error
	CreateInBatches(values []*model.ADMINMASTER, batchSize int) error
	Save(values ...*model.ADMINMASTER) error
	First() (*model.ADMINMASTER, error)
	Take() (*model.ADMINMASTER, error)
	Last() (*model.ADMINMASTER, error)
	Find() ([]*model.ADMINMASTER, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ADMINMASTER, err error)
	FindInBatches(result *[]*model.ADMINMASTER, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ADMINMASTER) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IADMINMASTERDo
	Assign(attrs ...field.AssignExpr) IADMINMASTERDo
	Joins(fields ...field.RelationField) IADMINMASTERDo
	Preload(fields ...field.RelationField) IADMINMASTERDo
	FirstOrInit() (*model.ADMINMASTER, error)
	FirstOrCreate() (*model.ADMINMASTER, error)
	FindByPage(offset int, limit int) (result []*model.ADMINMASTER, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IADMINMASTERDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aDMINMASTERDo) Debug() IADMINMASTERDo {
	return a.withDO(a.DO.Debug())
}

func (a aDMINMASTERDo) WithContext(ctx context.Context) IADMINMASTERDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aDMINMASTERDo) ReadDB() IADMINMASTERDo {
	return a.Clauses(dbresolver.Read)
}

func (a aDMINMASTERDo) WriteDB() IADMINMASTERDo {
	return a.Clauses(dbresolver.Write)
}

func (a aDMINMASTERDo) Session(config *gorm.Session) IADMINMASTERDo {
	return a.withDO(a.DO.Session(config))
}

func (a aDMINMASTERDo) Clauses(conds ...clause.Expression) IADMINMASTERDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aDMINMASTERDo) Returning(value interface{}, columns ...string) IADMINMASTERDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aDMINMASTERDo) Not(conds ...gen.Condition) IADMINMASTERDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aDMINMASTERDo) Or(conds ...gen.Condition) IADMINMASTERDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aDMINMASTERDo) Select(conds ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aDMINMASTERDo) Where(conds ...gen.Condition) IADMINMASTERDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aDMINMASTERDo) Order(conds ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aDMINMASTERDo) Distinct(cols ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aDMINMASTERDo) Omit(cols ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aDMINMASTERDo) Join(table schema.Tabler, on ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aDMINMASTERDo) LeftJoin(table schema.Tabler, on ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aDMINMASTERDo) RightJoin(table schema.Tabler, on ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aDMINMASTERDo) Group(cols ...field.Expr) IADMINMASTERDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aDMINMASTERDo) Having(conds ...gen.Condition) IADMINMASTERDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aDMINMASTERDo) Limit(limit int) IADMINMASTERDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aDMINMASTERDo) Offset(offset int) IADMINMASTERDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aDMINMASTERDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IADMINMASTERDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aDMINMASTERDo) Unscoped() IADMINMASTERDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aDMINMASTERDo) Create(values ...*model.ADMINMASTER) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aDMINMASTERDo) CreateInBatches(values []*model.ADMINMASTER, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aDMINMASTERDo) Save(values ...*model.ADMINMASTER) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aDMINMASTERDo) First() (*model.ADMINMASTER, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ADMINMASTER), nil
	}
}

func (a aDMINMASTERDo) Take() (*model.ADMINMASTER, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ADMINMASTER), nil
	}
}

func (a aDMINMASTERDo) Last() (*model.ADMINMASTER, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ADMINMASTER), nil
	}
}

func (a aDMINMASTERDo) Find() ([]*model.ADMINMASTER, error) {
	result, err := a.DO.Find()
	return result.([]*model.ADMINMASTER), err
}

func (a aDMINMASTERDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ADMINMASTER, err error) {
	buf := make([]*model.ADMINMASTER, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aDMINMASTERDo) FindInBatches(result *[]*model.ADMINMASTER, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aDMINMASTERDo) Attrs(attrs ...field.AssignExpr) IADMINMASTERDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aDMINMASTERDo) Assign(attrs ...field.AssignExpr) IADMINMASTERDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aDMINMASTERDo) Joins(fields ...field.RelationField) IADMINMASTERDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aDMINMASTERDo) Preload(fields ...field.RelationField) IADMINMASTERDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aDMINMASTERDo) FirstOrInit() (*model.ADMINMASTER, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ADMINMASTER), nil
	}
}

func (a aDMINMASTERDo) FirstOrCreate() (*model.ADMINMASTER, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ADMINMASTER), nil
	}
}

func (a aDMINMASTERDo) FindByPage(offset int, limit int) (result []*model.ADMINMASTER, count int64, err error) {
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

func (a aDMINMASTERDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aDMINMASTERDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aDMINMASTERDo) Delete(models ...*model.ADMINMASTER) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aDMINMASTERDo) withDO(do gen.Dao) *aDMINMASTERDo {
	a.DO = *do.(*gen.DO)
	return a
}
