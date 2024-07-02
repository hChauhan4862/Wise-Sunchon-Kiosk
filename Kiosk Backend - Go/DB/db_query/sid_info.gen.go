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

func newSIDINFO(db *gorm.DB, opts ...gen.DOOption) sIDINFO {
	_sIDINFO := sIDINFO{}

	_sIDINFO.sIDINFODo.UseDB(db, opts...)
	_sIDINFO.sIDINFODo.UseModel(&model.SIDINFO{})

	tableName := _sIDINFO.sIDINFODo.TableName()
	_sIDINFO.ALL = field.NewAsterisk(tableName)
	_sIDINFO.SID = field.NewString(tableName, "SID")
	_sIDINFO.ADMINID = field.NewString(tableName, "ADMIN_ID")
	_sIDINFO.ADMINNM = field.NewString(tableName, "ADMIN_NM")
	_sIDINFO.AUTHGRPCD = field.NewInt64(tableName, "AUTH_GRP_CD")
	_sIDINFO.AUTHTYPE = field.NewString(tableName, "AUTH_TYPE")
	_sIDINFO.REGDT = field.NewString(tableName, "REG_DT")

	_sIDINFO.fillFieldMap()

	return _sIDINFO
}

type sIDINFO struct {
	sIDINFODo

	ALL       field.Asterisk
	SID       field.String
	ADMINID   field.String
	ADMINNM   field.String
	AUTHGRPCD field.Int64
	AUTHTYPE  field.String
	REGDT     field.String

	fieldMap map[string]field.Expr
}

func (s sIDINFO) Table(newTableName string) *sIDINFO {
	s.sIDINFODo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sIDINFO) As(alias string) *sIDINFO {
	s.sIDINFODo.DO = *(s.sIDINFODo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sIDINFO) updateTableName(table string) *sIDINFO {
	s.ALL = field.NewAsterisk(table)
	s.SID = field.NewString(table, "SID")
	s.ADMINID = field.NewString(table, "ADMIN_ID")
	s.ADMINNM = field.NewString(table, "ADMIN_NM")
	s.AUTHGRPCD = field.NewInt64(table, "AUTH_GRP_CD")
	s.AUTHTYPE = field.NewString(table, "AUTH_TYPE")
	s.REGDT = field.NewString(table, "REG_DT")

	s.fillFieldMap()

	return s
}

func (s *sIDINFO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sIDINFO) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["SID"] = s.SID
	s.fieldMap["ADMIN_ID"] = s.ADMINID
	s.fieldMap["ADMIN_NM"] = s.ADMINNM
	s.fieldMap["AUTH_GRP_CD"] = s.AUTHGRPCD
	s.fieldMap["AUTH_TYPE"] = s.AUTHTYPE
	s.fieldMap["REG_DT"] = s.REGDT
}

func (s sIDINFO) clone(db *gorm.DB) sIDINFO {
	s.sIDINFODo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sIDINFO) replaceDB(db *gorm.DB) sIDINFO {
	s.sIDINFODo.ReplaceDB(db)
	return s
}

type sIDINFODo struct{ gen.DO }

type ISIDINFODo interface {
	gen.SubQuery
	Debug() ISIDINFODo
	WithContext(ctx context.Context) ISIDINFODo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISIDINFODo
	WriteDB() ISIDINFODo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISIDINFODo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISIDINFODo
	Not(conds ...gen.Condition) ISIDINFODo
	Or(conds ...gen.Condition) ISIDINFODo
	Select(conds ...field.Expr) ISIDINFODo
	Where(conds ...gen.Condition) ISIDINFODo
	Order(conds ...field.Expr) ISIDINFODo
	Distinct(cols ...field.Expr) ISIDINFODo
	Omit(cols ...field.Expr) ISIDINFODo
	Join(table schema.Tabler, on ...field.Expr) ISIDINFODo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISIDINFODo
	RightJoin(table schema.Tabler, on ...field.Expr) ISIDINFODo
	Group(cols ...field.Expr) ISIDINFODo
	Having(conds ...gen.Condition) ISIDINFODo
	Limit(limit int) ISIDINFODo
	Offset(offset int) ISIDINFODo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISIDINFODo
	Unscoped() ISIDINFODo
	Create(values ...*model.SIDINFO) error
	CreateInBatches(values []*model.SIDINFO, batchSize int) error
	Save(values ...*model.SIDINFO) error
	First() (*model.SIDINFO, error)
	Take() (*model.SIDINFO, error)
	Last() (*model.SIDINFO, error)
	Find() ([]*model.SIDINFO, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SIDINFO, err error)
	FindInBatches(result *[]*model.SIDINFO, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SIDINFO) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISIDINFODo
	Assign(attrs ...field.AssignExpr) ISIDINFODo
	Joins(fields ...field.RelationField) ISIDINFODo
	Preload(fields ...field.RelationField) ISIDINFODo
	FirstOrInit() (*model.SIDINFO, error)
	FirstOrCreate() (*model.SIDINFO, error)
	FindByPage(offset int, limit int) (result []*model.SIDINFO, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISIDINFODo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sIDINFODo) Debug() ISIDINFODo {
	return s.withDO(s.DO.Debug())
}

func (s sIDINFODo) WithContext(ctx context.Context) ISIDINFODo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sIDINFODo) ReadDB() ISIDINFODo {
	return s.Clauses(dbresolver.Read)
}

func (s sIDINFODo) WriteDB() ISIDINFODo {
	return s.Clauses(dbresolver.Write)
}

func (s sIDINFODo) Session(config *gorm.Session) ISIDINFODo {
	return s.withDO(s.DO.Session(config))
}

func (s sIDINFODo) Clauses(conds ...clause.Expression) ISIDINFODo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sIDINFODo) Returning(value interface{}, columns ...string) ISIDINFODo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sIDINFODo) Not(conds ...gen.Condition) ISIDINFODo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sIDINFODo) Or(conds ...gen.Condition) ISIDINFODo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sIDINFODo) Select(conds ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sIDINFODo) Where(conds ...gen.Condition) ISIDINFODo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sIDINFODo) Order(conds ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sIDINFODo) Distinct(cols ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sIDINFODo) Omit(cols ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sIDINFODo) Join(table schema.Tabler, on ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sIDINFODo) LeftJoin(table schema.Tabler, on ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sIDINFODo) RightJoin(table schema.Tabler, on ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sIDINFODo) Group(cols ...field.Expr) ISIDINFODo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sIDINFODo) Having(conds ...gen.Condition) ISIDINFODo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sIDINFODo) Limit(limit int) ISIDINFODo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sIDINFODo) Offset(offset int) ISIDINFODo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sIDINFODo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISIDINFODo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sIDINFODo) Unscoped() ISIDINFODo {
	return s.withDO(s.DO.Unscoped())
}

func (s sIDINFODo) Create(values ...*model.SIDINFO) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sIDINFODo) CreateInBatches(values []*model.SIDINFO, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sIDINFODo) Save(values ...*model.SIDINFO) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sIDINFODo) First() (*model.SIDINFO, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SIDINFO), nil
	}
}

func (s sIDINFODo) Take() (*model.SIDINFO, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SIDINFO), nil
	}
}

func (s sIDINFODo) Last() (*model.SIDINFO, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SIDINFO), nil
	}
}

func (s sIDINFODo) Find() ([]*model.SIDINFO, error) {
	result, err := s.DO.Find()
	return result.([]*model.SIDINFO), err
}

func (s sIDINFODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SIDINFO, err error) {
	buf := make([]*model.SIDINFO, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sIDINFODo) FindInBatches(result *[]*model.SIDINFO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sIDINFODo) Attrs(attrs ...field.AssignExpr) ISIDINFODo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sIDINFODo) Assign(attrs ...field.AssignExpr) ISIDINFODo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sIDINFODo) Joins(fields ...field.RelationField) ISIDINFODo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sIDINFODo) Preload(fields ...field.RelationField) ISIDINFODo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sIDINFODo) FirstOrInit() (*model.SIDINFO, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SIDINFO), nil
	}
}

func (s sIDINFODo) FirstOrCreate() (*model.SIDINFO, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SIDINFO), nil
	}
}

func (s sIDINFODo) FindByPage(offset int, limit int) (result []*model.SIDINFO, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sIDINFODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sIDINFODo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sIDINFODo) Delete(models ...*model.SIDINFO) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sIDINFODo) withDO(do gen.Dao) *sIDINFODo {
	s.DO = *do.(*gen.DO)
	return s
}
