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

func newSeatStatcd(db *gorm.DB, opts ...gen.DOOption) seatStatcd {
	_seatStatcd := seatStatcd{}

	_seatStatcd.seatStatcdDo.UseDB(db, opts...)
	_seatStatcd.seatStatcdDo.UseModel(&model.SeatStatcd{})

	tableName := _seatStatcd.seatStatcdDo.TableName()
	_seatStatcd.ALL = field.NewAsterisk(tableName)
	_seatStatcd.STATUS = field.NewInt64(tableName, "STATUS")
	_seatStatcd.NAME = field.NewString(tableName, "NAME")
	_seatStatcd.ENNAME = field.NewString(tableName, "EN_NAME")
	_seatStatcd.BOOKABLE = field.NewInt64(tableName, "BOOKABLE")
	_seatStatcd.DISABLED = field.NewInt64(tableName, "DISABLED")

	_seatStatcd.fillFieldMap()

	return _seatStatcd
}

type seatStatcd struct {
	seatStatcdDo

	ALL      field.Asterisk
	STATUS   field.Int64
	NAME     field.String
	ENNAME   field.String
	BOOKABLE field.Int64
	DISABLED field.Int64

	fieldMap map[string]field.Expr
}

func (s seatStatcd) Table(newTableName string) *seatStatcd {
	s.seatStatcdDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatStatcd) As(alias string) *seatStatcd {
	s.seatStatcdDo.DO = *(s.seatStatcdDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatStatcd) updateTableName(table string) *seatStatcd {
	s.ALL = field.NewAsterisk(table)
	s.STATUS = field.NewInt64(table, "STATUS")
	s.NAME = field.NewString(table, "NAME")
	s.ENNAME = field.NewString(table, "EN_NAME")
	s.BOOKABLE = field.NewInt64(table, "BOOKABLE")
	s.DISABLED = field.NewInt64(table, "DISABLED")

	s.fillFieldMap()

	return s
}

func (s *seatStatcd) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatStatcd) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["STATUS"] = s.STATUS
	s.fieldMap["NAME"] = s.NAME
	s.fieldMap["EN_NAME"] = s.ENNAME
	s.fieldMap["BOOKABLE"] = s.BOOKABLE
	s.fieldMap["DISABLED"] = s.DISABLED
}

func (s seatStatcd) clone(db *gorm.DB) seatStatcd {
	s.seatStatcdDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatStatcd) replaceDB(db *gorm.DB) seatStatcd {
	s.seatStatcdDo.ReplaceDB(db)
	return s
}

type seatStatcdDo struct{ gen.DO }

type ISeatStatcdDo interface {
	gen.SubQuery
	Debug() ISeatStatcdDo
	WithContext(ctx context.Context) ISeatStatcdDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatStatcdDo
	WriteDB() ISeatStatcdDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatStatcdDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatStatcdDo
	Not(conds ...gen.Condition) ISeatStatcdDo
	Or(conds ...gen.Condition) ISeatStatcdDo
	Select(conds ...field.Expr) ISeatStatcdDo
	Where(conds ...gen.Condition) ISeatStatcdDo
	Order(conds ...field.Expr) ISeatStatcdDo
	Distinct(cols ...field.Expr) ISeatStatcdDo
	Omit(cols ...field.Expr) ISeatStatcdDo
	Join(table schema.Tabler, on ...field.Expr) ISeatStatcdDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatStatcdDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatStatcdDo
	Group(cols ...field.Expr) ISeatStatcdDo
	Having(conds ...gen.Condition) ISeatStatcdDo
	Limit(limit int) ISeatStatcdDo
	Offset(offset int) ISeatStatcdDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatStatcdDo
	Unscoped() ISeatStatcdDo
	Create(values ...*model.SeatStatcd) error
	CreateInBatches(values []*model.SeatStatcd, batchSize int) error
	Save(values ...*model.SeatStatcd) error
	First() (*model.SeatStatcd, error)
	Take() (*model.SeatStatcd, error)
	Last() (*model.SeatStatcd, error)
	Find() ([]*model.SeatStatcd, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatStatcd, err error)
	FindInBatches(result *[]*model.SeatStatcd, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatStatcd) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatStatcdDo
	Assign(attrs ...field.AssignExpr) ISeatStatcdDo
	Joins(fields ...field.RelationField) ISeatStatcdDo
	Preload(fields ...field.RelationField) ISeatStatcdDo
	FirstOrInit() (*model.SeatStatcd, error)
	FirstOrCreate() (*model.SeatStatcd, error)
	FindByPage(offset int, limit int) (result []*model.SeatStatcd, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatStatcdDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatStatcdDo) Debug() ISeatStatcdDo {
	return s.withDO(s.DO.Debug())
}

func (s seatStatcdDo) WithContext(ctx context.Context) ISeatStatcdDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatStatcdDo) ReadDB() ISeatStatcdDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatStatcdDo) WriteDB() ISeatStatcdDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatStatcdDo) Session(config *gorm.Session) ISeatStatcdDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatStatcdDo) Clauses(conds ...clause.Expression) ISeatStatcdDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatStatcdDo) Returning(value interface{}, columns ...string) ISeatStatcdDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatStatcdDo) Not(conds ...gen.Condition) ISeatStatcdDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatStatcdDo) Or(conds ...gen.Condition) ISeatStatcdDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatStatcdDo) Select(conds ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatStatcdDo) Where(conds ...gen.Condition) ISeatStatcdDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatStatcdDo) Order(conds ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatStatcdDo) Distinct(cols ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatStatcdDo) Omit(cols ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatStatcdDo) Join(table schema.Tabler, on ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatStatcdDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatStatcdDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatStatcdDo) Group(cols ...field.Expr) ISeatStatcdDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatStatcdDo) Having(conds ...gen.Condition) ISeatStatcdDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatStatcdDo) Limit(limit int) ISeatStatcdDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatStatcdDo) Offset(offset int) ISeatStatcdDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatStatcdDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatStatcdDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatStatcdDo) Unscoped() ISeatStatcdDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatStatcdDo) Create(values ...*model.SeatStatcd) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatStatcdDo) CreateInBatches(values []*model.SeatStatcd, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatStatcdDo) Save(values ...*model.SeatStatcd) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatStatcdDo) First() (*model.SeatStatcd, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatStatcd), nil
	}
}

func (s seatStatcdDo) Take() (*model.SeatStatcd, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatStatcd), nil
	}
}

func (s seatStatcdDo) Last() (*model.SeatStatcd, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatStatcd), nil
	}
}

func (s seatStatcdDo) Find() ([]*model.SeatStatcd, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatStatcd), err
}

func (s seatStatcdDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatStatcd, err error) {
	buf := make([]*model.SeatStatcd, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatStatcdDo) FindInBatches(result *[]*model.SeatStatcd, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatStatcdDo) Attrs(attrs ...field.AssignExpr) ISeatStatcdDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatStatcdDo) Assign(attrs ...field.AssignExpr) ISeatStatcdDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatStatcdDo) Joins(fields ...field.RelationField) ISeatStatcdDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatStatcdDo) Preload(fields ...field.RelationField) ISeatStatcdDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatStatcdDo) FirstOrInit() (*model.SeatStatcd, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatStatcd), nil
	}
}

func (s seatStatcdDo) FirstOrCreate() (*model.SeatStatcd, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatStatcd), nil
	}
}

func (s seatStatcdDo) FindByPage(offset int, limit int) (result []*model.SeatStatcd, count int64, err error) {
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

func (s seatStatcdDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatStatcdDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatStatcdDo) Delete(models ...*model.SeatStatcd) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatStatcdDo) withDO(do gen.Dao) *seatStatcdDo {
	s.DO = *do.(*gen.DO)
	return s
}