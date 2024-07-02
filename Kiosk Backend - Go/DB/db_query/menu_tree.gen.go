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

func newMENUTREE(db *gorm.DB, opts ...gen.DOOption) mENUTREE {
	_mENUTREE := mENUTREE{}

	_mENUTREE.mENUTREEDo.UseDB(db, opts...)
	_mENUTREE.mENUTREEDo.UseModel(&model.MENUTREE{})

	tableName := _mENUTREE.mENUTREEDo.TableName()
	_mENUTREE.ALL = field.NewAsterisk(tableName)
	_mENUTREE.MENUID = field.NewString(tableName, "MENU_ID")
	_mENUTREE.MENUNM = field.NewString(tableName, "MENU_NM")
	_mENUTREE.DISPORDER = field.NewInt64(tableName, "DISP_ORDER")
	_mENUTREE.LVL = field.NewString(tableName, "LVL")
	_mENUTREE.UPPERMENUID = field.NewString(tableName, "UPPER_MENU_ID")
	_mENUTREE.MENUURL = field.NewString(tableName, "MENU_URL")
	_mENUTREE.USEYN = field.NewString(tableName, "USE_YN")
	_mENUTREE.REGDT = field.NewString(tableName, "REG_DT")
	_mENUTREE.REGID = field.NewString(tableName, "REG_ID")
	_mENUTREE.MODDT = field.NewString(tableName, "MOD_DT")
	_mENUTREE.MODID = field.NewString(tableName, "MOD_ID")

	_mENUTREE.fillFieldMap()

	return _mENUTREE
}

type mENUTREE struct {
	mENUTREEDo

	ALL         field.Asterisk
	MENUID      field.String
	MENUNM      field.String
	DISPORDER   field.Int64
	LVL         field.String
	UPPERMENUID field.String
	MENUURL     field.String
	USEYN       field.String
	REGDT       field.String
	REGID       field.String
	MODDT       field.String
	MODID       field.String

	fieldMap map[string]field.Expr
}

func (m mENUTREE) Table(newTableName string) *mENUTREE {
	m.mENUTREEDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mENUTREE) As(alias string) *mENUTREE {
	m.mENUTREEDo.DO = *(m.mENUTREEDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mENUTREE) updateTableName(table string) *mENUTREE {
	m.ALL = field.NewAsterisk(table)
	m.MENUID = field.NewString(table, "MENU_ID")
	m.MENUNM = field.NewString(table, "MENU_NM")
	m.DISPORDER = field.NewInt64(table, "DISP_ORDER")
	m.LVL = field.NewString(table, "LVL")
	m.UPPERMENUID = field.NewString(table, "UPPER_MENU_ID")
	m.MENUURL = field.NewString(table, "MENU_URL")
	m.USEYN = field.NewString(table, "USE_YN")
	m.REGDT = field.NewString(table, "REG_DT")
	m.REGID = field.NewString(table, "REG_ID")
	m.MODDT = field.NewString(table, "MOD_DT")
	m.MODID = field.NewString(table, "MOD_ID")

	m.fillFieldMap()

	return m
}

func (m *mENUTREE) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mENUTREE) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 11)
	m.fieldMap["MENU_ID"] = m.MENUID
	m.fieldMap["MENU_NM"] = m.MENUNM
	m.fieldMap["DISP_ORDER"] = m.DISPORDER
	m.fieldMap["LVL"] = m.LVL
	m.fieldMap["UPPER_MENU_ID"] = m.UPPERMENUID
	m.fieldMap["MENU_URL"] = m.MENUURL
	m.fieldMap["USE_YN"] = m.USEYN
	m.fieldMap["REG_DT"] = m.REGDT
	m.fieldMap["REG_ID"] = m.REGID
	m.fieldMap["MOD_DT"] = m.MODDT
	m.fieldMap["MOD_ID"] = m.MODID
}

func (m mENUTREE) clone(db *gorm.DB) mENUTREE {
	m.mENUTREEDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mENUTREE) replaceDB(db *gorm.DB) mENUTREE {
	m.mENUTREEDo.ReplaceDB(db)
	return m
}

type mENUTREEDo struct{ gen.DO }

type IMENUTREEDo interface {
	gen.SubQuery
	Debug() IMENUTREEDo
	WithContext(ctx context.Context) IMENUTREEDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMENUTREEDo
	WriteDB() IMENUTREEDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMENUTREEDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMENUTREEDo
	Not(conds ...gen.Condition) IMENUTREEDo
	Or(conds ...gen.Condition) IMENUTREEDo
	Select(conds ...field.Expr) IMENUTREEDo
	Where(conds ...gen.Condition) IMENUTREEDo
	Order(conds ...field.Expr) IMENUTREEDo
	Distinct(cols ...field.Expr) IMENUTREEDo
	Omit(cols ...field.Expr) IMENUTREEDo
	Join(table schema.Tabler, on ...field.Expr) IMENUTREEDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMENUTREEDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMENUTREEDo
	Group(cols ...field.Expr) IMENUTREEDo
	Having(conds ...gen.Condition) IMENUTREEDo
	Limit(limit int) IMENUTREEDo
	Offset(offset int) IMENUTREEDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMENUTREEDo
	Unscoped() IMENUTREEDo
	Create(values ...*model.MENUTREE) error
	CreateInBatches(values []*model.MENUTREE, batchSize int) error
	Save(values ...*model.MENUTREE) error
	First() (*model.MENUTREE, error)
	Take() (*model.MENUTREE, error)
	Last() (*model.MENUTREE, error)
	Find() ([]*model.MENUTREE, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MENUTREE, err error)
	FindInBatches(result *[]*model.MENUTREE, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MENUTREE) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMENUTREEDo
	Assign(attrs ...field.AssignExpr) IMENUTREEDo
	Joins(fields ...field.RelationField) IMENUTREEDo
	Preload(fields ...field.RelationField) IMENUTREEDo
	FirstOrInit() (*model.MENUTREE, error)
	FirstOrCreate() (*model.MENUTREE, error)
	FindByPage(offset int, limit int) (result []*model.MENUTREE, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMENUTREEDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mENUTREEDo) Debug() IMENUTREEDo {
	return m.withDO(m.DO.Debug())
}

func (m mENUTREEDo) WithContext(ctx context.Context) IMENUTREEDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mENUTREEDo) ReadDB() IMENUTREEDo {
	return m.Clauses(dbresolver.Read)
}

func (m mENUTREEDo) WriteDB() IMENUTREEDo {
	return m.Clauses(dbresolver.Write)
}

func (m mENUTREEDo) Session(config *gorm.Session) IMENUTREEDo {
	return m.withDO(m.DO.Session(config))
}

func (m mENUTREEDo) Clauses(conds ...clause.Expression) IMENUTREEDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mENUTREEDo) Returning(value interface{}, columns ...string) IMENUTREEDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mENUTREEDo) Not(conds ...gen.Condition) IMENUTREEDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mENUTREEDo) Or(conds ...gen.Condition) IMENUTREEDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mENUTREEDo) Select(conds ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mENUTREEDo) Where(conds ...gen.Condition) IMENUTREEDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mENUTREEDo) Order(conds ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mENUTREEDo) Distinct(cols ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mENUTREEDo) Omit(cols ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mENUTREEDo) Join(table schema.Tabler, on ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mENUTREEDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mENUTREEDo) RightJoin(table schema.Tabler, on ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mENUTREEDo) Group(cols ...field.Expr) IMENUTREEDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mENUTREEDo) Having(conds ...gen.Condition) IMENUTREEDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mENUTREEDo) Limit(limit int) IMENUTREEDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mENUTREEDo) Offset(offset int) IMENUTREEDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mENUTREEDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMENUTREEDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mENUTREEDo) Unscoped() IMENUTREEDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mENUTREEDo) Create(values ...*model.MENUTREE) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mENUTREEDo) CreateInBatches(values []*model.MENUTREE, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mENUTREEDo) Save(values ...*model.MENUTREE) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mENUTREEDo) First() (*model.MENUTREE, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MENUTREE), nil
	}
}

func (m mENUTREEDo) Take() (*model.MENUTREE, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MENUTREE), nil
	}
}

func (m mENUTREEDo) Last() (*model.MENUTREE, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MENUTREE), nil
	}
}

func (m mENUTREEDo) Find() ([]*model.MENUTREE, error) {
	result, err := m.DO.Find()
	return result.([]*model.MENUTREE), err
}

func (m mENUTREEDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MENUTREE, err error) {
	buf := make([]*model.MENUTREE, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mENUTREEDo) FindInBatches(result *[]*model.MENUTREE, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mENUTREEDo) Attrs(attrs ...field.AssignExpr) IMENUTREEDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mENUTREEDo) Assign(attrs ...field.AssignExpr) IMENUTREEDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mENUTREEDo) Joins(fields ...field.RelationField) IMENUTREEDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mENUTREEDo) Preload(fields ...field.RelationField) IMENUTREEDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mENUTREEDo) FirstOrInit() (*model.MENUTREE, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MENUTREE), nil
	}
}

func (m mENUTREEDo) FirstOrCreate() (*model.MENUTREE, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MENUTREE), nil
	}
}

func (m mENUTREEDo) FindByPage(offset int, limit int) (result []*model.MENUTREE, count int64, err error) {
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

func (m mENUTREEDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mENUTREEDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mENUTREEDo) Delete(models ...*model.MENUTREE) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mENUTREEDo) withDO(do gen.Dao) *mENUTREEDo {
	m.DO = *do.(*gen.DO)
	return m
}
