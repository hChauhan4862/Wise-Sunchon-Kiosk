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

func newIDCARDISSUE(db *gorm.DB, opts ...gen.DOOption) iDCARDISSUE {
	_iDCARDISSUE := iDCARDISSUE{}

	_iDCARDISSUE.iDCARDISSUEDo.UseDB(db, opts...)
	_iDCARDISSUE.iDCARDISSUEDo.UseModel(&model.IDCARDISSUE{})

	tableName := _iDCARDISSUE.iDCARDISSUEDo.TableName()
	_iDCARDISSUE.ALL = field.NewAsterisk(tableName)
	_iDCARDISSUE.PID = field.NewString(tableName, "PID")
	_iDCARDISSUE.ISSUECNT = field.NewInt64(tableName, "ISSUE_CNT")
	_iDCARDISSUE.ISSUEDATE = field.NewString(tableName, "ISSUE_DATE")
	_iDCARDISSUE.DEVICEGB = field.NewString(tableName, "DEVICE_GB")
	_iDCARDISSUE.TELNO = field.NewString(tableName, "TEL_NO")
	_iDCARDISSUE.USEYN = field.NewString(tableName, "USE_YN")
	_iDCARDISSUE.FSTREGDT = field.NewString(tableName, "FST_REG_DT")
	_iDCARDISSUE.DELDT = field.NewString(tableName, "DEL_DT")

	_iDCARDISSUE.fillFieldMap()

	return _iDCARDISSUE
}

type iDCARDISSUE struct {
	iDCARDISSUEDo

	ALL       field.Asterisk
	PID       field.String
	ISSUECNT  field.Int64
	ISSUEDATE field.String
	DEVICEGB  field.String
	TELNO     field.String
	USEYN     field.String
	FSTREGDT  field.String
	DELDT     field.String

	fieldMap map[string]field.Expr
}

func (i iDCARDISSUE) Table(newTableName string) *iDCARDISSUE {
	i.iDCARDISSUEDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i iDCARDISSUE) As(alias string) *iDCARDISSUE {
	i.iDCARDISSUEDo.DO = *(i.iDCARDISSUEDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *iDCARDISSUE) updateTableName(table string) *iDCARDISSUE {
	i.ALL = field.NewAsterisk(table)
	i.PID = field.NewString(table, "PID")
	i.ISSUECNT = field.NewInt64(table, "ISSUE_CNT")
	i.ISSUEDATE = field.NewString(table, "ISSUE_DATE")
	i.DEVICEGB = field.NewString(table, "DEVICE_GB")
	i.TELNO = field.NewString(table, "TEL_NO")
	i.USEYN = field.NewString(table, "USE_YN")
	i.FSTREGDT = field.NewString(table, "FST_REG_DT")
	i.DELDT = field.NewString(table, "DEL_DT")

	i.fillFieldMap()

	return i
}

func (i *iDCARDISSUE) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *iDCARDISSUE) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 8)
	i.fieldMap["PID"] = i.PID
	i.fieldMap["ISSUE_CNT"] = i.ISSUECNT
	i.fieldMap["ISSUE_DATE"] = i.ISSUEDATE
	i.fieldMap["DEVICE_GB"] = i.DEVICEGB
	i.fieldMap["TEL_NO"] = i.TELNO
	i.fieldMap["USE_YN"] = i.USEYN
	i.fieldMap["FST_REG_DT"] = i.FSTREGDT
	i.fieldMap["DEL_DT"] = i.DELDT
}

func (i iDCARDISSUE) clone(db *gorm.DB) iDCARDISSUE {
	i.iDCARDISSUEDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i iDCARDISSUE) replaceDB(db *gorm.DB) iDCARDISSUE {
	i.iDCARDISSUEDo.ReplaceDB(db)
	return i
}

type iDCARDISSUEDo struct{ gen.DO }

type IIDCARDISSUEDo interface {
	gen.SubQuery
	Debug() IIDCARDISSUEDo
	WithContext(ctx context.Context) IIDCARDISSUEDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIDCARDISSUEDo
	WriteDB() IIDCARDISSUEDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIDCARDISSUEDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIDCARDISSUEDo
	Not(conds ...gen.Condition) IIDCARDISSUEDo
	Or(conds ...gen.Condition) IIDCARDISSUEDo
	Select(conds ...field.Expr) IIDCARDISSUEDo
	Where(conds ...gen.Condition) IIDCARDISSUEDo
	Order(conds ...field.Expr) IIDCARDISSUEDo
	Distinct(cols ...field.Expr) IIDCARDISSUEDo
	Omit(cols ...field.Expr) IIDCARDISSUEDo
	Join(table schema.Tabler, on ...field.Expr) IIDCARDISSUEDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIDCARDISSUEDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIDCARDISSUEDo
	Group(cols ...field.Expr) IIDCARDISSUEDo
	Having(conds ...gen.Condition) IIDCARDISSUEDo
	Limit(limit int) IIDCARDISSUEDo
	Offset(offset int) IIDCARDISSUEDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIDCARDISSUEDo
	Unscoped() IIDCARDISSUEDo
	Create(values ...*model.IDCARDISSUE) error
	CreateInBatches(values []*model.IDCARDISSUE, batchSize int) error
	Save(values ...*model.IDCARDISSUE) error
	First() (*model.IDCARDISSUE, error)
	Take() (*model.IDCARDISSUE, error)
	Last() (*model.IDCARDISSUE, error)
	Find() ([]*model.IDCARDISSUE, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IDCARDISSUE, err error)
	FindInBatches(result *[]*model.IDCARDISSUE, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.IDCARDISSUE) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIDCARDISSUEDo
	Assign(attrs ...field.AssignExpr) IIDCARDISSUEDo
	Joins(fields ...field.RelationField) IIDCARDISSUEDo
	Preload(fields ...field.RelationField) IIDCARDISSUEDo
	FirstOrInit() (*model.IDCARDISSUE, error)
	FirstOrCreate() (*model.IDCARDISSUE, error)
	FindByPage(offset int, limit int) (result []*model.IDCARDISSUE, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIDCARDISSUEDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i iDCARDISSUEDo) Debug() IIDCARDISSUEDo {
	return i.withDO(i.DO.Debug())
}

func (i iDCARDISSUEDo) WithContext(ctx context.Context) IIDCARDISSUEDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i iDCARDISSUEDo) ReadDB() IIDCARDISSUEDo {
	return i.Clauses(dbresolver.Read)
}

func (i iDCARDISSUEDo) WriteDB() IIDCARDISSUEDo {
	return i.Clauses(dbresolver.Write)
}

func (i iDCARDISSUEDo) Session(config *gorm.Session) IIDCARDISSUEDo {
	return i.withDO(i.DO.Session(config))
}

func (i iDCARDISSUEDo) Clauses(conds ...clause.Expression) IIDCARDISSUEDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i iDCARDISSUEDo) Returning(value interface{}, columns ...string) IIDCARDISSUEDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i iDCARDISSUEDo) Not(conds ...gen.Condition) IIDCARDISSUEDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i iDCARDISSUEDo) Or(conds ...gen.Condition) IIDCARDISSUEDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i iDCARDISSUEDo) Select(conds ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i iDCARDISSUEDo) Where(conds ...gen.Condition) IIDCARDISSUEDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i iDCARDISSUEDo) Order(conds ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i iDCARDISSUEDo) Distinct(cols ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i iDCARDISSUEDo) Omit(cols ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i iDCARDISSUEDo) Join(table schema.Tabler, on ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i iDCARDISSUEDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i iDCARDISSUEDo) RightJoin(table schema.Tabler, on ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i iDCARDISSUEDo) Group(cols ...field.Expr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i iDCARDISSUEDo) Having(conds ...gen.Condition) IIDCARDISSUEDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i iDCARDISSUEDo) Limit(limit int) IIDCARDISSUEDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i iDCARDISSUEDo) Offset(offset int) IIDCARDISSUEDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i iDCARDISSUEDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIDCARDISSUEDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i iDCARDISSUEDo) Unscoped() IIDCARDISSUEDo {
	return i.withDO(i.DO.Unscoped())
}

func (i iDCARDISSUEDo) Create(values ...*model.IDCARDISSUE) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i iDCARDISSUEDo) CreateInBatches(values []*model.IDCARDISSUE, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i iDCARDISSUEDo) Save(values ...*model.IDCARDISSUE) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i iDCARDISSUEDo) First() (*model.IDCARDISSUE, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDISSUE), nil
	}
}

func (i iDCARDISSUEDo) Take() (*model.IDCARDISSUE, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDISSUE), nil
	}
}

func (i iDCARDISSUEDo) Last() (*model.IDCARDISSUE, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDISSUE), nil
	}
}

func (i iDCARDISSUEDo) Find() ([]*model.IDCARDISSUE, error) {
	result, err := i.DO.Find()
	return result.([]*model.IDCARDISSUE), err
}

func (i iDCARDISSUEDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IDCARDISSUE, err error) {
	buf := make([]*model.IDCARDISSUE, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i iDCARDISSUEDo) FindInBatches(result *[]*model.IDCARDISSUE, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i iDCARDISSUEDo) Attrs(attrs ...field.AssignExpr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i iDCARDISSUEDo) Assign(attrs ...field.AssignExpr) IIDCARDISSUEDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i iDCARDISSUEDo) Joins(fields ...field.RelationField) IIDCARDISSUEDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i iDCARDISSUEDo) Preload(fields ...field.RelationField) IIDCARDISSUEDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i iDCARDISSUEDo) FirstOrInit() (*model.IDCARDISSUE, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDISSUE), nil
	}
}

func (i iDCARDISSUEDo) FirstOrCreate() (*model.IDCARDISSUE, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDISSUE), nil
	}
}

func (i iDCARDISSUEDo) FindByPage(offset int, limit int) (result []*model.IDCARDISSUE, count int64, err error) {
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

func (i iDCARDISSUEDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i iDCARDISSUEDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i iDCARDISSUEDo) Delete(models ...*model.IDCARDISSUE) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *iDCARDISSUEDo) withDO(do gen.Dao) *iDCARDISSUEDo {
	i.DO = *do.(*gen.DO)
	return i
}
