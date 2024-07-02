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

func newIDCARDUSER(db *gorm.DB, opts ...gen.DOOption) iDCARDUSER {
	_iDCARDUSER := iDCARDUSER{}

	_iDCARDUSER.iDCARDUSERDo.UseDB(db, opts...)
	_iDCARDUSER.iDCARDUSERDo.UseModel(&model.IDCARDUSER{})

	tableName := _iDCARDUSER.iDCARDUSERDo.TableName()
	_iDCARDUSER.ALL = field.NewAsterisk(tableName)
	_iDCARDUSER.PID = field.NewString(tableName, "PID")
	_iDCARDUSER.RID = field.NewString(tableName, "RID")
	_iDCARDUSER.NAME = field.NewString(tableName, "NAME")
	_iDCARDUSER.DEPTCODE = field.NewString(tableName, "DEPT_CODE")
	_iDCARDUSER.DEPTNAME = field.NewString(tableName, "DEPT_NAME")
	_iDCARDUSER.PATCODE = field.NewString(tableName, "PAT_CODE")
	_iDCARDUSER.PATNAME = field.NewString(tableName, "PAT_NAME")
	_iDCARDUSER.TELNO = field.NewString(tableName, "TEL_NO")
	_iDCARDUSER.DEVICEGB = field.NewString(tableName, "DEVICE_GB")
	_iDCARDUSER.PASSWD = field.NewString(tableName, "PASS_WD")
	_iDCARDUSER.PUSHKEY = field.NewString(tableName, "PUSH_KEY")
	_iDCARDUSER.PHOTOURL = field.NewString(tableName, "PHOTO_URL")
	_iDCARDUSER.DESCRIPT = field.NewString(tableName, "DESCRIPT")
	_iDCARDUSER.FSTREGDT = field.NewString(tableName, "FST_REG_DT")
	_iDCARDUSER.UPDDT = field.NewString(tableName, "UPD_DT")
	_iDCARDUSER.DELDT = field.NewString(tableName, "DEL_DT")
	_iDCARDUSER.AUTHKEY = field.NewString(tableName, "AUTH_KEY")

	_iDCARDUSER.fillFieldMap()

	return _iDCARDUSER
}

type iDCARDUSER struct {
	iDCARDUSERDo

	ALL      field.Asterisk
	PID      field.String
	RID      field.String
	NAME     field.String
	DEPTCODE field.String
	DEPTNAME field.String
	PATCODE  field.String
	PATNAME  field.String
	TELNO    field.String
	DEVICEGB field.String
	PASSWD   field.String
	PUSHKEY  field.String
	PHOTOURL field.String
	DESCRIPT field.String
	FSTREGDT field.String
	UPDDT    field.String
	DELDT    field.String
	AUTHKEY  field.String

	fieldMap map[string]field.Expr
}

func (i iDCARDUSER) Table(newTableName string) *iDCARDUSER {
	i.iDCARDUSERDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i iDCARDUSER) As(alias string) *iDCARDUSER {
	i.iDCARDUSERDo.DO = *(i.iDCARDUSERDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *iDCARDUSER) updateTableName(table string) *iDCARDUSER {
	i.ALL = field.NewAsterisk(table)
	i.PID = field.NewString(table, "PID")
	i.RID = field.NewString(table, "RID")
	i.NAME = field.NewString(table, "NAME")
	i.DEPTCODE = field.NewString(table, "DEPT_CODE")
	i.DEPTNAME = field.NewString(table, "DEPT_NAME")
	i.PATCODE = field.NewString(table, "PAT_CODE")
	i.PATNAME = field.NewString(table, "PAT_NAME")
	i.TELNO = field.NewString(table, "TEL_NO")
	i.DEVICEGB = field.NewString(table, "DEVICE_GB")
	i.PASSWD = field.NewString(table, "PASS_WD")
	i.PUSHKEY = field.NewString(table, "PUSH_KEY")
	i.PHOTOURL = field.NewString(table, "PHOTO_URL")
	i.DESCRIPT = field.NewString(table, "DESCRIPT")
	i.FSTREGDT = field.NewString(table, "FST_REG_DT")
	i.UPDDT = field.NewString(table, "UPD_DT")
	i.DELDT = field.NewString(table, "DEL_DT")
	i.AUTHKEY = field.NewString(table, "AUTH_KEY")

	i.fillFieldMap()

	return i
}

func (i *iDCARDUSER) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *iDCARDUSER) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 17)
	i.fieldMap["PID"] = i.PID
	i.fieldMap["RID"] = i.RID
	i.fieldMap["NAME"] = i.NAME
	i.fieldMap["DEPT_CODE"] = i.DEPTCODE
	i.fieldMap["DEPT_NAME"] = i.DEPTNAME
	i.fieldMap["PAT_CODE"] = i.PATCODE
	i.fieldMap["PAT_NAME"] = i.PATNAME
	i.fieldMap["TEL_NO"] = i.TELNO
	i.fieldMap["DEVICE_GB"] = i.DEVICEGB
	i.fieldMap["PASS_WD"] = i.PASSWD
	i.fieldMap["PUSH_KEY"] = i.PUSHKEY
	i.fieldMap["PHOTO_URL"] = i.PHOTOURL
	i.fieldMap["DESCRIPT"] = i.DESCRIPT
	i.fieldMap["FST_REG_DT"] = i.FSTREGDT
	i.fieldMap["UPD_DT"] = i.UPDDT
	i.fieldMap["DEL_DT"] = i.DELDT
	i.fieldMap["AUTH_KEY"] = i.AUTHKEY
}

func (i iDCARDUSER) clone(db *gorm.DB) iDCARDUSER {
	i.iDCARDUSERDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i iDCARDUSER) replaceDB(db *gorm.DB) iDCARDUSER {
	i.iDCARDUSERDo.ReplaceDB(db)
	return i
}

type iDCARDUSERDo struct{ gen.DO }

type IIDCARDUSERDo interface {
	gen.SubQuery
	Debug() IIDCARDUSERDo
	WithContext(ctx context.Context) IIDCARDUSERDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIDCARDUSERDo
	WriteDB() IIDCARDUSERDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIDCARDUSERDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIDCARDUSERDo
	Not(conds ...gen.Condition) IIDCARDUSERDo
	Or(conds ...gen.Condition) IIDCARDUSERDo
	Select(conds ...field.Expr) IIDCARDUSERDo
	Where(conds ...gen.Condition) IIDCARDUSERDo
	Order(conds ...field.Expr) IIDCARDUSERDo
	Distinct(cols ...field.Expr) IIDCARDUSERDo
	Omit(cols ...field.Expr) IIDCARDUSERDo
	Join(table schema.Tabler, on ...field.Expr) IIDCARDUSERDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIDCARDUSERDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIDCARDUSERDo
	Group(cols ...field.Expr) IIDCARDUSERDo
	Having(conds ...gen.Condition) IIDCARDUSERDo
	Limit(limit int) IIDCARDUSERDo
	Offset(offset int) IIDCARDUSERDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIDCARDUSERDo
	Unscoped() IIDCARDUSERDo
	Create(values ...*model.IDCARDUSER) error
	CreateInBatches(values []*model.IDCARDUSER, batchSize int) error
	Save(values ...*model.IDCARDUSER) error
	First() (*model.IDCARDUSER, error)
	Take() (*model.IDCARDUSER, error)
	Last() (*model.IDCARDUSER, error)
	Find() ([]*model.IDCARDUSER, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IDCARDUSER, err error)
	FindInBatches(result *[]*model.IDCARDUSER, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.IDCARDUSER) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIDCARDUSERDo
	Assign(attrs ...field.AssignExpr) IIDCARDUSERDo
	Joins(fields ...field.RelationField) IIDCARDUSERDo
	Preload(fields ...field.RelationField) IIDCARDUSERDo
	FirstOrInit() (*model.IDCARDUSER, error)
	FirstOrCreate() (*model.IDCARDUSER, error)
	FindByPage(offset int, limit int) (result []*model.IDCARDUSER, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIDCARDUSERDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i iDCARDUSERDo) Debug() IIDCARDUSERDo {
	return i.withDO(i.DO.Debug())
}

func (i iDCARDUSERDo) WithContext(ctx context.Context) IIDCARDUSERDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i iDCARDUSERDo) ReadDB() IIDCARDUSERDo {
	return i.Clauses(dbresolver.Read)
}

func (i iDCARDUSERDo) WriteDB() IIDCARDUSERDo {
	return i.Clauses(dbresolver.Write)
}

func (i iDCARDUSERDo) Session(config *gorm.Session) IIDCARDUSERDo {
	return i.withDO(i.DO.Session(config))
}

func (i iDCARDUSERDo) Clauses(conds ...clause.Expression) IIDCARDUSERDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i iDCARDUSERDo) Returning(value interface{}, columns ...string) IIDCARDUSERDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i iDCARDUSERDo) Not(conds ...gen.Condition) IIDCARDUSERDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i iDCARDUSERDo) Or(conds ...gen.Condition) IIDCARDUSERDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i iDCARDUSERDo) Select(conds ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i iDCARDUSERDo) Where(conds ...gen.Condition) IIDCARDUSERDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i iDCARDUSERDo) Order(conds ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i iDCARDUSERDo) Distinct(cols ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i iDCARDUSERDo) Omit(cols ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i iDCARDUSERDo) Join(table schema.Tabler, on ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i iDCARDUSERDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i iDCARDUSERDo) RightJoin(table schema.Tabler, on ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i iDCARDUSERDo) Group(cols ...field.Expr) IIDCARDUSERDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i iDCARDUSERDo) Having(conds ...gen.Condition) IIDCARDUSERDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i iDCARDUSERDo) Limit(limit int) IIDCARDUSERDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i iDCARDUSERDo) Offset(offset int) IIDCARDUSERDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i iDCARDUSERDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIDCARDUSERDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i iDCARDUSERDo) Unscoped() IIDCARDUSERDo {
	return i.withDO(i.DO.Unscoped())
}

func (i iDCARDUSERDo) Create(values ...*model.IDCARDUSER) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i iDCARDUSERDo) CreateInBatches(values []*model.IDCARDUSER, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i iDCARDUSERDo) Save(values ...*model.IDCARDUSER) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i iDCARDUSERDo) First() (*model.IDCARDUSER, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDUSER), nil
	}
}

func (i iDCARDUSERDo) Take() (*model.IDCARDUSER, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDUSER), nil
	}
}

func (i iDCARDUSERDo) Last() (*model.IDCARDUSER, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDUSER), nil
	}
}

func (i iDCARDUSERDo) Find() ([]*model.IDCARDUSER, error) {
	result, err := i.DO.Find()
	return result.([]*model.IDCARDUSER), err
}

func (i iDCARDUSERDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IDCARDUSER, err error) {
	buf := make([]*model.IDCARDUSER, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i iDCARDUSERDo) FindInBatches(result *[]*model.IDCARDUSER, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i iDCARDUSERDo) Attrs(attrs ...field.AssignExpr) IIDCARDUSERDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i iDCARDUSERDo) Assign(attrs ...field.AssignExpr) IIDCARDUSERDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i iDCARDUSERDo) Joins(fields ...field.RelationField) IIDCARDUSERDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i iDCARDUSERDo) Preload(fields ...field.RelationField) IIDCARDUSERDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i iDCARDUSERDo) FirstOrInit() (*model.IDCARDUSER, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDUSER), nil
	}
}

func (i iDCARDUSERDo) FirstOrCreate() (*model.IDCARDUSER, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDUSER), nil
	}
}

func (i iDCARDUSERDo) FindByPage(offset int, limit int) (result []*model.IDCARDUSER, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i iDCARDUSERDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i iDCARDUSERDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i iDCARDUSERDo) Delete(models ...*model.IDCARDUSER) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *iDCARDUSERDo) withDO(do gen.Dao) *iDCARDUSERDo {
	i.DO = *do.(*gen.DO)
	return i
}
