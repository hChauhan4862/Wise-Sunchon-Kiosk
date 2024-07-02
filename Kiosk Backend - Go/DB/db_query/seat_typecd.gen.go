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

func newSeatTypecd(db *gorm.DB, opts ...gen.DOOption) seatTypecd {
	_seatTypecd := seatTypecd{}

	_seatTypecd.seatTypecdDo.UseDB(db, opts...)
	_seatTypecd.seatTypecdDo.UseModel(&model.SeatTypecd{})

	tableName := _seatTypecd.seatTypecdDo.TableName()
	_seatTypecd.ALL = field.NewAsterisk(tableName)
	_seatTypecd.TYPENO = field.NewInt64(tableName, "TYPENO")
	_seatTypecd.NAME = field.NewString(tableName, "NAME")
	_seatTypecd.ENNAME = field.NewString(tableName, "EN_NAME")
	_seatTypecd.MEDIATYPES = field.NewString(tableName, "MEDIATYPES")
	_seatTypecd.EQTYPES = field.NewString(tableName, "EQTYPES")
	_seatTypecd.ISKIOSKASSIGN = field.NewInt64(tableName, "IS_KIOSKASSIGN")
	_seatTypecd.DAYMAXUSECNT = field.NewInt64(tableName, "DAY_MAXUSECNT")
	_seatTypecd.DAYMAXUSEMIN = field.NewInt64(tableName, "DAY_MAXUSEMIN")
	_seatTypecd.MONMAXUSECNT = field.NewInt64(tableName, "MON_MAXUSECNT")
	_seatTypecd.MONMAXUSEMIN = field.NewInt64(tableName, "MON_MAXUSEMIN")
	_seatTypecd.PCLOGINCHECK = field.NewInt64(tableName, "PC_LOGIN_CHECK")
	_seatTypecd.AUTOPRINTCNT = field.NewInt64(tableName, "AUTO_PRINT_CNT")

	_seatTypecd.fillFieldMap()

	return _seatTypecd
}

type seatTypecd struct {
	seatTypecdDo

	ALL           field.Asterisk
	TYPENO        field.Int64
	NAME          field.String
	ENNAME        field.String
	MEDIATYPES    field.String
	EQTYPES       field.String
	ISKIOSKASSIGN field.Int64
	DAYMAXUSECNT  field.Int64
	DAYMAXUSEMIN  field.Int64
	MONMAXUSECNT  field.Int64
	MONMAXUSEMIN  field.Int64
	PCLOGINCHECK  field.Int64
	AUTOPRINTCNT  field.Int64

	fieldMap map[string]field.Expr
}

func (s seatTypecd) Table(newTableName string) *seatTypecd {
	s.seatTypecdDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatTypecd) As(alias string) *seatTypecd {
	s.seatTypecdDo.DO = *(s.seatTypecdDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatTypecd) updateTableName(table string) *seatTypecd {
	s.ALL = field.NewAsterisk(table)
	s.TYPENO = field.NewInt64(table, "TYPENO")
	s.NAME = field.NewString(table, "NAME")
	s.ENNAME = field.NewString(table, "EN_NAME")
	s.MEDIATYPES = field.NewString(table, "MEDIATYPES")
	s.EQTYPES = field.NewString(table, "EQTYPES")
	s.ISKIOSKASSIGN = field.NewInt64(table, "IS_KIOSKASSIGN")
	s.DAYMAXUSECNT = field.NewInt64(table, "DAY_MAXUSECNT")
	s.DAYMAXUSEMIN = field.NewInt64(table, "DAY_MAXUSEMIN")
	s.MONMAXUSECNT = field.NewInt64(table, "MON_MAXUSECNT")
	s.MONMAXUSEMIN = field.NewInt64(table, "MON_MAXUSEMIN")
	s.PCLOGINCHECK = field.NewInt64(table, "PC_LOGIN_CHECK")
	s.AUTOPRINTCNT = field.NewInt64(table, "AUTO_PRINT_CNT")

	s.fillFieldMap()

	return s
}

func (s *seatTypecd) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatTypecd) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["TYPENO"] = s.TYPENO
	s.fieldMap["NAME"] = s.NAME
	s.fieldMap["EN_NAME"] = s.ENNAME
	s.fieldMap["MEDIATYPES"] = s.MEDIATYPES
	s.fieldMap["EQTYPES"] = s.EQTYPES
	s.fieldMap["IS_KIOSKASSIGN"] = s.ISKIOSKASSIGN
	s.fieldMap["DAY_MAXUSECNT"] = s.DAYMAXUSECNT
	s.fieldMap["DAY_MAXUSEMIN"] = s.DAYMAXUSEMIN
	s.fieldMap["MON_MAXUSECNT"] = s.MONMAXUSECNT
	s.fieldMap["MON_MAXUSEMIN"] = s.MONMAXUSEMIN
	s.fieldMap["PC_LOGIN_CHECK"] = s.PCLOGINCHECK
	s.fieldMap["AUTO_PRINT_CNT"] = s.AUTOPRINTCNT
}

func (s seatTypecd) clone(db *gorm.DB) seatTypecd {
	s.seatTypecdDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatTypecd) replaceDB(db *gorm.DB) seatTypecd {
	s.seatTypecdDo.ReplaceDB(db)
	return s
}

type seatTypecdDo struct{ gen.DO }

type ISeatTypecdDo interface {
	gen.SubQuery
	Debug() ISeatTypecdDo
	WithContext(ctx context.Context) ISeatTypecdDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatTypecdDo
	WriteDB() ISeatTypecdDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatTypecdDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatTypecdDo
	Not(conds ...gen.Condition) ISeatTypecdDo
	Or(conds ...gen.Condition) ISeatTypecdDo
	Select(conds ...field.Expr) ISeatTypecdDo
	Where(conds ...gen.Condition) ISeatTypecdDo
	Order(conds ...field.Expr) ISeatTypecdDo
	Distinct(cols ...field.Expr) ISeatTypecdDo
	Omit(cols ...field.Expr) ISeatTypecdDo
	Join(table schema.Tabler, on ...field.Expr) ISeatTypecdDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatTypecdDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatTypecdDo
	Group(cols ...field.Expr) ISeatTypecdDo
	Having(conds ...gen.Condition) ISeatTypecdDo
	Limit(limit int) ISeatTypecdDo
	Offset(offset int) ISeatTypecdDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatTypecdDo
	Unscoped() ISeatTypecdDo
	Create(values ...*model.SeatTypecd) error
	CreateInBatches(values []*model.SeatTypecd, batchSize int) error
	Save(values ...*model.SeatTypecd) error
	First() (*model.SeatTypecd, error)
	Take() (*model.SeatTypecd, error)
	Last() (*model.SeatTypecd, error)
	Find() ([]*model.SeatTypecd, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatTypecd, err error)
	FindInBatches(result *[]*model.SeatTypecd, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatTypecd) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatTypecdDo
	Assign(attrs ...field.AssignExpr) ISeatTypecdDo
	Joins(fields ...field.RelationField) ISeatTypecdDo
	Preload(fields ...field.RelationField) ISeatTypecdDo
	FirstOrInit() (*model.SeatTypecd, error)
	FirstOrCreate() (*model.SeatTypecd, error)
	FindByPage(offset int, limit int) (result []*model.SeatTypecd, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatTypecdDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatTypecdDo) Debug() ISeatTypecdDo {
	return s.withDO(s.DO.Debug())
}

func (s seatTypecdDo) WithContext(ctx context.Context) ISeatTypecdDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatTypecdDo) ReadDB() ISeatTypecdDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatTypecdDo) WriteDB() ISeatTypecdDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatTypecdDo) Session(config *gorm.Session) ISeatTypecdDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatTypecdDo) Clauses(conds ...clause.Expression) ISeatTypecdDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatTypecdDo) Returning(value interface{}, columns ...string) ISeatTypecdDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatTypecdDo) Not(conds ...gen.Condition) ISeatTypecdDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatTypecdDo) Or(conds ...gen.Condition) ISeatTypecdDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatTypecdDo) Select(conds ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatTypecdDo) Where(conds ...gen.Condition) ISeatTypecdDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatTypecdDo) Order(conds ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatTypecdDo) Distinct(cols ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatTypecdDo) Omit(cols ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatTypecdDo) Join(table schema.Tabler, on ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatTypecdDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatTypecdDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatTypecdDo) Group(cols ...field.Expr) ISeatTypecdDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatTypecdDo) Having(conds ...gen.Condition) ISeatTypecdDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatTypecdDo) Limit(limit int) ISeatTypecdDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatTypecdDo) Offset(offset int) ISeatTypecdDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatTypecdDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatTypecdDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatTypecdDo) Unscoped() ISeatTypecdDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatTypecdDo) Create(values ...*model.SeatTypecd) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatTypecdDo) CreateInBatches(values []*model.SeatTypecd, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatTypecdDo) Save(values ...*model.SeatTypecd) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatTypecdDo) First() (*model.SeatTypecd, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatTypecd), nil
	}
}

func (s seatTypecdDo) Take() (*model.SeatTypecd, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatTypecd), nil
	}
}

func (s seatTypecdDo) Last() (*model.SeatTypecd, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatTypecd), nil
	}
}

func (s seatTypecdDo) Find() ([]*model.SeatTypecd, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatTypecd), err
}

func (s seatTypecdDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatTypecd, err error) {
	buf := make([]*model.SeatTypecd, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatTypecdDo) FindInBatches(result *[]*model.SeatTypecd, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatTypecdDo) Attrs(attrs ...field.AssignExpr) ISeatTypecdDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatTypecdDo) Assign(attrs ...field.AssignExpr) ISeatTypecdDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatTypecdDo) Joins(fields ...field.RelationField) ISeatTypecdDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatTypecdDo) Preload(fields ...field.RelationField) ISeatTypecdDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatTypecdDo) FirstOrInit() (*model.SeatTypecd, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatTypecd), nil
	}
}

func (s seatTypecdDo) FirstOrCreate() (*model.SeatTypecd, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatTypecd), nil
	}
}

func (s seatTypecdDo) FindByPage(offset int, limit int) (result []*model.SeatTypecd, count int64, err error) {
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

func (s seatTypecdDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatTypecdDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatTypecdDo) Delete(models ...*model.SeatTypecd) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatTypecdDo) withDO(do gen.Dao) *seatTypecdDo {
	s.DO = *do.(*gen.DO)
	return s
}