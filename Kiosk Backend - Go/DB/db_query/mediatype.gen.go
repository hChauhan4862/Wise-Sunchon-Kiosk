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

func newMEDIATYPE(db *gorm.DB, opts ...gen.DOOption) mEDIATYPE {
	_mEDIATYPE := mEDIATYPE{}

	_mEDIATYPE.mEDIATYPEDo.UseDB(db, opts...)
	_mEDIATYPE.mEDIATYPEDo.UseModel(&model.MEDIATYPE{})

	tableName := _mEDIATYPE.mEDIATYPEDo.TableName()
	_mEDIATYPE.ALL = field.NewAsterisk(tableName)
	_mEDIATYPE.MEDIATYPE = field.NewString(tableName, "MEDIA_TYPE")
	_mEDIATYPE.NAME = field.NewString(tableName, "NAME")
	_mEDIATYPE.ENNAME = field.NewString(tableName, "EN_NAME")
	_mEDIATYPE.CLASSID = field.NewString(tableName, "CLASSID")

	_mEDIATYPE.fillFieldMap()

	return _mEDIATYPE
}

type mEDIATYPE struct {
	mEDIATYPEDo

	ALL       field.Asterisk
	MEDIATYPE field.String
	NAME      field.String
	ENNAME    field.String
	CLASSID   field.String

	fieldMap map[string]field.Expr
}

func (m mEDIATYPE) Table(newTableName string) *mEDIATYPE {
	m.mEDIATYPEDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mEDIATYPE) As(alias string) *mEDIATYPE {
	m.mEDIATYPEDo.DO = *(m.mEDIATYPEDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mEDIATYPE) updateTableName(table string) *mEDIATYPE {
	m.ALL = field.NewAsterisk(table)
	m.MEDIATYPE = field.NewString(table, "MEDIA_TYPE")
	m.NAME = field.NewString(table, "NAME")
	m.ENNAME = field.NewString(table, "EN_NAME")
	m.CLASSID = field.NewString(table, "CLASSID")

	m.fillFieldMap()

	return m
}

func (m *mEDIATYPE) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mEDIATYPE) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["MEDIA_TYPE"] = m.MEDIATYPE
	m.fieldMap["NAME"] = m.NAME
	m.fieldMap["EN_NAME"] = m.ENNAME
	m.fieldMap["CLASSID"] = m.CLASSID
}

func (m mEDIATYPE) clone(db *gorm.DB) mEDIATYPE {
	m.mEDIATYPEDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mEDIATYPE) replaceDB(db *gorm.DB) mEDIATYPE {
	m.mEDIATYPEDo.ReplaceDB(db)
	return m
}

type mEDIATYPEDo struct{ gen.DO }

type IMEDIATYPEDo interface {
	gen.SubQuery
	Debug() IMEDIATYPEDo
	WithContext(ctx context.Context) IMEDIATYPEDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMEDIATYPEDo
	WriteDB() IMEDIATYPEDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMEDIATYPEDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMEDIATYPEDo
	Not(conds ...gen.Condition) IMEDIATYPEDo
	Or(conds ...gen.Condition) IMEDIATYPEDo
	Select(conds ...field.Expr) IMEDIATYPEDo
	Where(conds ...gen.Condition) IMEDIATYPEDo
	Order(conds ...field.Expr) IMEDIATYPEDo
	Distinct(cols ...field.Expr) IMEDIATYPEDo
	Omit(cols ...field.Expr) IMEDIATYPEDo
	Join(table schema.Tabler, on ...field.Expr) IMEDIATYPEDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMEDIATYPEDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMEDIATYPEDo
	Group(cols ...field.Expr) IMEDIATYPEDo
	Having(conds ...gen.Condition) IMEDIATYPEDo
	Limit(limit int) IMEDIATYPEDo
	Offset(offset int) IMEDIATYPEDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMEDIATYPEDo
	Unscoped() IMEDIATYPEDo
	Create(values ...*model.MEDIATYPE) error
	CreateInBatches(values []*model.MEDIATYPE, batchSize int) error
	Save(values ...*model.MEDIATYPE) error
	First() (*model.MEDIATYPE, error)
	Take() (*model.MEDIATYPE, error)
	Last() (*model.MEDIATYPE, error)
	Find() ([]*model.MEDIATYPE, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MEDIATYPE, err error)
	FindInBatches(result *[]*model.MEDIATYPE, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MEDIATYPE) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMEDIATYPEDo
	Assign(attrs ...field.AssignExpr) IMEDIATYPEDo
	Joins(fields ...field.RelationField) IMEDIATYPEDo
	Preload(fields ...field.RelationField) IMEDIATYPEDo
	FirstOrInit() (*model.MEDIATYPE, error)
	FirstOrCreate() (*model.MEDIATYPE, error)
	FindByPage(offset int, limit int) (result []*model.MEDIATYPE, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMEDIATYPEDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mEDIATYPEDo) Debug() IMEDIATYPEDo {
	return m.withDO(m.DO.Debug())
}

func (m mEDIATYPEDo) WithContext(ctx context.Context) IMEDIATYPEDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mEDIATYPEDo) ReadDB() IMEDIATYPEDo {
	return m.Clauses(dbresolver.Read)
}

func (m mEDIATYPEDo) WriteDB() IMEDIATYPEDo {
	return m.Clauses(dbresolver.Write)
}

func (m mEDIATYPEDo) Session(config *gorm.Session) IMEDIATYPEDo {
	return m.withDO(m.DO.Session(config))
}

func (m mEDIATYPEDo) Clauses(conds ...clause.Expression) IMEDIATYPEDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mEDIATYPEDo) Returning(value interface{}, columns ...string) IMEDIATYPEDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mEDIATYPEDo) Not(conds ...gen.Condition) IMEDIATYPEDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mEDIATYPEDo) Or(conds ...gen.Condition) IMEDIATYPEDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mEDIATYPEDo) Select(conds ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mEDIATYPEDo) Where(conds ...gen.Condition) IMEDIATYPEDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mEDIATYPEDo) Order(conds ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mEDIATYPEDo) Distinct(cols ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mEDIATYPEDo) Omit(cols ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mEDIATYPEDo) Join(table schema.Tabler, on ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mEDIATYPEDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mEDIATYPEDo) RightJoin(table schema.Tabler, on ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mEDIATYPEDo) Group(cols ...field.Expr) IMEDIATYPEDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mEDIATYPEDo) Having(conds ...gen.Condition) IMEDIATYPEDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mEDIATYPEDo) Limit(limit int) IMEDIATYPEDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mEDIATYPEDo) Offset(offset int) IMEDIATYPEDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mEDIATYPEDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMEDIATYPEDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mEDIATYPEDo) Unscoped() IMEDIATYPEDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mEDIATYPEDo) Create(values ...*model.MEDIATYPE) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mEDIATYPEDo) CreateInBatches(values []*model.MEDIATYPE, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mEDIATYPEDo) Save(values ...*model.MEDIATYPE) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mEDIATYPEDo) First() (*model.MEDIATYPE, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MEDIATYPE), nil
	}
}

func (m mEDIATYPEDo) Take() (*model.MEDIATYPE, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MEDIATYPE), nil
	}
}

func (m mEDIATYPEDo) Last() (*model.MEDIATYPE, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MEDIATYPE), nil
	}
}

func (m mEDIATYPEDo) Find() ([]*model.MEDIATYPE, error) {
	result, err := m.DO.Find()
	return result.([]*model.MEDIATYPE), err
}

func (m mEDIATYPEDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MEDIATYPE, err error) {
	buf := make([]*model.MEDIATYPE, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mEDIATYPEDo) FindInBatches(result *[]*model.MEDIATYPE, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mEDIATYPEDo) Attrs(attrs ...field.AssignExpr) IMEDIATYPEDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mEDIATYPEDo) Assign(attrs ...field.AssignExpr) IMEDIATYPEDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mEDIATYPEDo) Joins(fields ...field.RelationField) IMEDIATYPEDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mEDIATYPEDo) Preload(fields ...field.RelationField) IMEDIATYPEDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mEDIATYPEDo) FirstOrInit() (*model.MEDIATYPE, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MEDIATYPE), nil
	}
}

func (m mEDIATYPEDo) FirstOrCreate() (*model.MEDIATYPE, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MEDIATYPE), nil
	}
}

func (m mEDIATYPEDo) FindByPage(offset int, limit int) (result []*model.MEDIATYPE, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m mEDIATYPEDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mEDIATYPEDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mEDIATYPEDo) Delete(models ...*model.MEDIATYPE) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mEDIATYPEDo) withDO(do gen.Dao) *mEDIATYPEDo {
	m.DO = *do.(*gen.DO)
	return m
}