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

func newSeatLib(db *gorm.DB, opts ...gen.DOOption) seatLib {
	_seatLib := seatLib{}

	_seatLib.seatLibDo.UseDB(db, opts...)
	_seatLib.seatLibDo.UseModel(&model.SeatLib{})

	tableName := _seatLib.seatLibDo.TableName()
	_seatLib.ALL = field.NewAsterisk(tableName)
	_seatLib.LIBNO = field.NewInt64(tableName, "LIBNO")
	_seatLib.NAME = field.NewString(tableName, "NAME")
	_seatLib.ENNAME = field.NewString(tableName, "EN_NAME")
	_seatLib.SORT = field.NewInt64(tableName, "SORT")

	_seatLib.fillFieldMap()

	return _seatLib
}

type seatLib struct {
	seatLibDo

	ALL    field.Asterisk
	LIBNO  field.Int64
	NAME   field.String
	ENNAME field.String
	SORT   field.Int64

	fieldMap map[string]field.Expr
}

func (s seatLib) Table(newTableName string) *seatLib {
	s.seatLibDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatLib) As(alias string) *seatLib {
	s.seatLibDo.DO = *(s.seatLibDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatLib) updateTableName(table string) *seatLib {
	s.ALL = field.NewAsterisk(table)
	s.LIBNO = field.NewInt64(table, "LIBNO")
	s.NAME = field.NewString(table, "NAME")
	s.ENNAME = field.NewString(table, "EN_NAME")
	s.SORT = field.NewInt64(table, "SORT")

	s.fillFieldMap()

	return s
}

func (s *seatLib) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatLib) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["LIBNO"] = s.LIBNO
	s.fieldMap["NAME"] = s.NAME
	s.fieldMap["EN_NAME"] = s.ENNAME
	s.fieldMap["SORT"] = s.SORT
}

func (s seatLib) clone(db *gorm.DB) seatLib {
	s.seatLibDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatLib) replaceDB(db *gorm.DB) seatLib {
	s.seatLibDo.ReplaceDB(db)
	return s
}

type seatLibDo struct{ gen.DO }

type ISeatLibDo interface {
	gen.SubQuery
	Debug() ISeatLibDo
	WithContext(ctx context.Context) ISeatLibDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatLibDo
	WriteDB() ISeatLibDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatLibDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatLibDo
	Not(conds ...gen.Condition) ISeatLibDo
	Or(conds ...gen.Condition) ISeatLibDo
	Select(conds ...field.Expr) ISeatLibDo
	Where(conds ...gen.Condition) ISeatLibDo
	Order(conds ...field.Expr) ISeatLibDo
	Distinct(cols ...field.Expr) ISeatLibDo
	Omit(cols ...field.Expr) ISeatLibDo
	Join(table schema.Tabler, on ...field.Expr) ISeatLibDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatLibDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatLibDo
	Group(cols ...field.Expr) ISeatLibDo
	Having(conds ...gen.Condition) ISeatLibDo
	Limit(limit int) ISeatLibDo
	Offset(offset int) ISeatLibDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatLibDo
	Unscoped() ISeatLibDo
	Create(values ...*model.SeatLib) error
	CreateInBatches(values []*model.SeatLib, batchSize int) error
	Save(values ...*model.SeatLib) error
	First() (*model.SeatLib, error)
	Take() (*model.SeatLib, error)
	Last() (*model.SeatLib, error)
	Find() ([]*model.SeatLib, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatLib, err error)
	FindInBatches(result *[]*model.SeatLib, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatLib) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatLibDo
	Assign(attrs ...field.AssignExpr) ISeatLibDo
	Joins(fields ...field.RelationField) ISeatLibDo
	Preload(fields ...field.RelationField) ISeatLibDo
	FirstOrInit() (*model.SeatLib, error)
	FirstOrCreate() (*model.SeatLib, error)
	FindByPage(offset int, limit int) (result []*model.SeatLib, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatLibDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatLibDo) Debug() ISeatLibDo {
	return s.withDO(s.DO.Debug())
}

func (s seatLibDo) WithContext(ctx context.Context) ISeatLibDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatLibDo) ReadDB() ISeatLibDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatLibDo) WriteDB() ISeatLibDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatLibDo) Session(config *gorm.Session) ISeatLibDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatLibDo) Clauses(conds ...clause.Expression) ISeatLibDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatLibDo) Returning(value interface{}, columns ...string) ISeatLibDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatLibDo) Not(conds ...gen.Condition) ISeatLibDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatLibDo) Or(conds ...gen.Condition) ISeatLibDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatLibDo) Select(conds ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatLibDo) Where(conds ...gen.Condition) ISeatLibDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatLibDo) Order(conds ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatLibDo) Distinct(cols ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatLibDo) Omit(cols ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatLibDo) Join(table schema.Tabler, on ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatLibDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatLibDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatLibDo) Group(cols ...field.Expr) ISeatLibDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatLibDo) Having(conds ...gen.Condition) ISeatLibDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatLibDo) Limit(limit int) ISeatLibDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatLibDo) Offset(offset int) ISeatLibDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatLibDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatLibDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatLibDo) Unscoped() ISeatLibDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatLibDo) Create(values ...*model.SeatLib) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatLibDo) CreateInBatches(values []*model.SeatLib, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatLibDo) Save(values ...*model.SeatLib) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatLibDo) First() (*model.SeatLib, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatLib), nil
	}
}

func (s seatLibDo) Take() (*model.SeatLib, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatLib), nil
	}
}

func (s seatLibDo) Last() (*model.SeatLib, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatLib), nil
	}
}

func (s seatLibDo) Find() ([]*model.SeatLib, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatLib), err
}

func (s seatLibDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatLib, err error) {
	buf := make([]*model.SeatLib, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatLibDo) FindInBatches(result *[]*model.SeatLib, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatLibDo) Attrs(attrs ...field.AssignExpr) ISeatLibDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatLibDo) Assign(attrs ...field.AssignExpr) ISeatLibDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatLibDo) Joins(fields ...field.RelationField) ISeatLibDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatLibDo) Preload(fields ...field.RelationField) ISeatLibDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatLibDo) FirstOrInit() (*model.SeatLib, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatLib), nil
	}
}

func (s seatLibDo) FirstOrCreate() (*model.SeatLib, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatLib), nil
	}
}

func (s seatLibDo) FindByPage(offset int, limit int) (result []*model.SeatLib, count int64, err error) {
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

func (s seatLibDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatLibDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatLibDo) Delete(models ...*model.SeatLib) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatLibDo) withDO(do gen.Dao) *seatLibDo {
	s.DO = *do.(*gen.DO)
	return s
}
