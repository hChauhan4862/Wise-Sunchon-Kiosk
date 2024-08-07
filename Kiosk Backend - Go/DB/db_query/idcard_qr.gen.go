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

func newIDCARDQR(db *gorm.DB, opts ...gen.DOOption) iDCARDQR {
	_iDCARDQR := iDCARDQR{}

	_iDCARDQR.iDCARDQRDo.UseDB(db, opts...)
	_iDCARDQR.iDCARDQRDo.UseModel(&model.IDCARDQR{})

	tableName := _iDCARDQR.iDCARDQRDo.TableName()
	_iDCARDQR.ALL = field.NewAsterisk(tableName)
	_iDCARDQR.PID = field.NewString(tableName, "PID")
	_iDCARDQR.DEVICEGB = field.NewString(tableName, "DEVICE_GB")
	_iDCARDQR.ISSUEDATE = field.NewString(tableName, "ISSUE_DATE")
	_iDCARDQR.QRCODE = field.NewString(tableName, "QR_CODE")
	_iDCARDQR.CHECKDIGIT = field.NewString(tableName, "CHECK_DIGIT")

	_iDCARDQR.fillFieldMap()

	return _iDCARDQR
}

type iDCARDQR struct {
	iDCARDQRDo

	ALL        field.Asterisk
	PID        field.String
	DEVICEGB   field.String
	ISSUEDATE  field.String
	QRCODE     field.String
	CHECKDIGIT field.String

	fieldMap map[string]field.Expr
}

func (i iDCARDQR) Table(newTableName string) *iDCARDQR {
	i.iDCARDQRDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i iDCARDQR) As(alias string) *iDCARDQR {
	i.iDCARDQRDo.DO = *(i.iDCARDQRDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *iDCARDQR) updateTableName(table string) *iDCARDQR {
	i.ALL = field.NewAsterisk(table)
	i.PID = field.NewString(table, "PID")
	i.DEVICEGB = field.NewString(table, "DEVICE_GB")
	i.ISSUEDATE = field.NewString(table, "ISSUE_DATE")
	i.QRCODE = field.NewString(table, "QR_CODE")
	i.CHECKDIGIT = field.NewString(table, "CHECK_DIGIT")

	i.fillFieldMap()

	return i
}

func (i *iDCARDQR) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *iDCARDQR) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 5)
	i.fieldMap["PID"] = i.PID
	i.fieldMap["DEVICE_GB"] = i.DEVICEGB
	i.fieldMap["ISSUE_DATE"] = i.ISSUEDATE
	i.fieldMap["QR_CODE"] = i.QRCODE
	i.fieldMap["CHECK_DIGIT"] = i.CHECKDIGIT
}

func (i iDCARDQR) clone(db *gorm.DB) iDCARDQR {
	i.iDCARDQRDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i iDCARDQR) replaceDB(db *gorm.DB) iDCARDQR {
	i.iDCARDQRDo.ReplaceDB(db)
	return i
}

type iDCARDQRDo struct{ gen.DO }

type IIDCARDQRDo interface {
	gen.SubQuery
	Debug() IIDCARDQRDo
	WithContext(ctx context.Context) IIDCARDQRDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIDCARDQRDo
	WriteDB() IIDCARDQRDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIDCARDQRDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIDCARDQRDo
	Not(conds ...gen.Condition) IIDCARDQRDo
	Or(conds ...gen.Condition) IIDCARDQRDo
	Select(conds ...field.Expr) IIDCARDQRDo
	Where(conds ...gen.Condition) IIDCARDQRDo
	Order(conds ...field.Expr) IIDCARDQRDo
	Distinct(cols ...field.Expr) IIDCARDQRDo
	Omit(cols ...field.Expr) IIDCARDQRDo
	Join(table schema.Tabler, on ...field.Expr) IIDCARDQRDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIDCARDQRDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIDCARDQRDo
	Group(cols ...field.Expr) IIDCARDQRDo
	Having(conds ...gen.Condition) IIDCARDQRDo
	Limit(limit int) IIDCARDQRDo
	Offset(offset int) IIDCARDQRDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIDCARDQRDo
	Unscoped() IIDCARDQRDo
	Create(values ...*model.IDCARDQR) error
	CreateInBatches(values []*model.IDCARDQR, batchSize int) error
	Save(values ...*model.IDCARDQR) error
	First() (*model.IDCARDQR, error)
	Take() (*model.IDCARDQR, error)
	Last() (*model.IDCARDQR, error)
	Find() ([]*model.IDCARDQR, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IDCARDQR, err error)
	FindInBatches(result *[]*model.IDCARDQR, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.IDCARDQR) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIDCARDQRDo
	Assign(attrs ...field.AssignExpr) IIDCARDQRDo
	Joins(fields ...field.RelationField) IIDCARDQRDo
	Preload(fields ...field.RelationField) IIDCARDQRDo
	FirstOrInit() (*model.IDCARDQR, error)
	FirstOrCreate() (*model.IDCARDQR, error)
	FindByPage(offset int, limit int) (result []*model.IDCARDQR, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIDCARDQRDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i iDCARDQRDo) Debug() IIDCARDQRDo {
	return i.withDO(i.DO.Debug())
}

func (i iDCARDQRDo) WithContext(ctx context.Context) IIDCARDQRDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i iDCARDQRDo) ReadDB() IIDCARDQRDo {
	return i.Clauses(dbresolver.Read)
}

func (i iDCARDQRDo) WriteDB() IIDCARDQRDo {
	return i.Clauses(dbresolver.Write)
}

func (i iDCARDQRDo) Session(config *gorm.Session) IIDCARDQRDo {
	return i.withDO(i.DO.Session(config))
}

func (i iDCARDQRDo) Clauses(conds ...clause.Expression) IIDCARDQRDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i iDCARDQRDo) Returning(value interface{}, columns ...string) IIDCARDQRDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i iDCARDQRDo) Not(conds ...gen.Condition) IIDCARDQRDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i iDCARDQRDo) Or(conds ...gen.Condition) IIDCARDQRDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i iDCARDQRDo) Select(conds ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i iDCARDQRDo) Where(conds ...gen.Condition) IIDCARDQRDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i iDCARDQRDo) Order(conds ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i iDCARDQRDo) Distinct(cols ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i iDCARDQRDo) Omit(cols ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i iDCARDQRDo) Join(table schema.Tabler, on ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i iDCARDQRDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i iDCARDQRDo) RightJoin(table schema.Tabler, on ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i iDCARDQRDo) Group(cols ...field.Expr) IIDCARDQRDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i iDCARDQRDo) Having(conds ...gen.Condition) IIDCARDQRDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i iDCARDQRDo) Limit(limit int) IIDCARDQRDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i iDCARDQRDo) Offset(offset int) IIDCARDQRDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i iDCARDQRDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIDCARDQRDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i iDCARDQRDo) Unscoped() IIDCARDQRDo {
	return i.withDO(i.DO.Unscoped())
}

func (i iDCARDQRDo) Create(values ...*model.IDCARDQR) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i iDCARDQRDo) CreateInBatches(values []*model.IDCARDQR, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i iDCARDQRDo) Save(values ...*model.IDCARDQR) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i iDCARDQRDo) First() (*model.IDCARDQR, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDQR), nil
	}
}

func (i iDCARDQRDo) Take() (*model.IDCARDQR, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDQR), nil
	}
}

func (i iDCARDQRDo) Last() (*model.IDCARDQR, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDQR), nil
	}
}

func (i iDCARDQRDo) Find() ([]*model.IDCARDQR, error) {
	result, err := i.DO.Find()
	return result.([]*model.IDCARDQR), err
}

func (i iDCARDQRDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IDCARDQR, err error) {
	buf := make([]*model.IDCARDQR, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i iDCARDQRDo) FindInBatches(result *[]*model.IDCARDQR, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i iDCARDQRDo) Attrs(attrs ...field.AssignExpr) IIDCARDQRDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i iDCARDQRDo) Assign(attrs ...field.AssignExpr) IIDCARDQRDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i iDCARDQRDo) Joins(fields ...field.RelationField) IIDCARDQRDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i iDCARDQRDo) Preload(fields ...field.RelationField) IIDCARDQRDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i iDCARDQRDo) FirstOrInit() (*model.IDCARDQR, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDQR), nil
	}
}

func (i iDCARDQRDo) FirstOrCreate() (*model.IDCARDQR, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.IDCARDQR), nil
	}
}

func (i iDCARDQRDo) FindByPage(offset int, limit int) (result []*model.IDCARDQR, count int64, err error) {
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

func (i iDCARDQRDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i iDCARDQRDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i iDCARDQRDo) Delete(models ...*model.IDCARDQR) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *iDCARDQRDo) withDO(do gen.Dao) *iDCARDQRDo {
	i.DO = *do.(*gen.DO)
	return i
}
