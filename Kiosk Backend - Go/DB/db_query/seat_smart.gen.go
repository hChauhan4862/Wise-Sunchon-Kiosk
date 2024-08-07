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

func newSeatSmart(db *gorm.DB, opts ...gen.DOOption) seatSmart {
	_seatSmart := seatSmart{}

	_seatSmart.seatSmartDo.UseDB(db, opts...)
	_seatSmart.seatSmartDo.UseModel(&model.SeatSmart{})

	tableName := _seatSmart.seatSmartDo.TableName()
	_seatSmart.ALL = field.NewAsterisk(tableName)
	_seatSmart.RoomNo = field.NewInt64(tableName, "room_no")

	_seatSmart.fillFieldMap()

	return _seatSmart
}

type seatSmart struct {
	seatSmartDo

	ALL    field.Asterisk
	RoomNo field.Int64

	fieldMap map[string]field.Expr
}

func (s seatSmart) Table(newTableName string) *seatSmart {
	s.seatSmartDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatSmart) As(alias string) *seatSmart {
	s.seatSmartDo.DO = *(s.seatSmartDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatSmart) updateTableName(table string) *seatSmart {
	s.ALL = field.NewAsterisk(table)
	s.RoomNo = field.NewInt64(table, "room_no")

	s.fillFieldMap()

	return s
}

func (s *seatSmart) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatSmart) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 1)
	s.fieldMap["room_no"] = s.RoomNo
}

func (s seatSmart) clone(db *gorm.DB) seatSmart {
	s.seatSmartDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatSmart) replaceDB(db *gorm.DB) seatSmart {
	s.seatSmartDo.ReplaceDB(db)
	return s
}

type seatSmartDo struct{ gen.DO }

type ISeatSmartDo interface {
	gen.SubQuery
	Debug() ISeatSmartDo
	WithContext(ctx context.Context) ISeatSmartDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatSmartDo
	WriteDB() ISeatSmartDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatSmartDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatSmartDo
	Not(conds ...gen.Condition) ISeatSmartDo
	Or(conds ...gen.Condition) ISeatSmartDo
	Select(conds ...field.Expr) ISeatSmartDo
	Where(conds ...gen.Condition) ISeatSmartDo
	Order(conds ...field.Expr) ISeatSmartDo
	Distinct(cols ...field.Expr) ISeatSmartDo
	Omit(cols ...field.Expr) ISeatSmartDo
	Join(table schema.Tabler, on ...field.Expr) ISeatSmartDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatSmartDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatSmartDo
	Group(cols ...field.Expr) ISeatSmartDo
	Having(conds ...gen.Condition) ISeatSmartDo
	Limit(limit int) ISeatSmartDo
	Offset(offset int) ISeatSmartDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatSmartDo
	Unscoped() ISeatSmartDo
	Create(values ...*model.SeatSmart) error
	CreateInBatches(values []*model.SeatSmart, batchSize int) error
	Save(values ...*model.SeatSmart) error
	First() (*model.SeatSmart, error)
	Take() (*model.SeatSmart, error)
	Last() (*model.SeatSmart, error)
	Find() ([]*model.SeatSmart, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatSmart, err error)
	FindInBatches(result *[]*model.SeatSmart, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatSmart) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatSmartDo
	Assign(attrs ...field.AssignExpr) ISeatSmartDo
	Joins(fields ...field.RelationField) ISeatSmartDo
	Preload(fields ...field.RelationField) ISeatSmartDo
	FirstOrInit() (*model.SeatSmart, error)
	FirstOrCreate() (*model.SeatSmart, error)
	FindByPage(offset int, limit int) (result []*model.SeatSmart, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatSmartDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatSmartDo) Debug() ISeatSmartDo {
	return s.withDO(s.DO.Debug())
}

func (s seatSmartDo) WithContext(ctx context.Context) ISeatSmartDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatSmartDo) ReadDB() ISeatSmartDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatSmartDo) WriteDB() ISeatSmartDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatSmartDo) Session(config *gorm.Session) ISeatSmartDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatSmartDo) Clauses(conds ...clause.Expression) ISeatSmartDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatSmartDo) Returning(value interface{}, columns ...string) ISeatSmartDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatSmartDo) Not(conds ...gen.Condition) ISeatSmartDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatSmartDo) Or(conds ...gen.Condition) ISeatSmartDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatSmartDo) Select(conds ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatSmartDo) Where(conds ...gen.Condition) ISeatSmartDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatSmartDo) Order(conds ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatSmartDo) Distinct(cols ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatSmartDo) Omit(cols ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatSmartDo) Join(table schema.Tabler, on ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatSmartDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatSmartDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatSmartDo) Group(cols ...field.Expr) ISeatSmartDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatSmartDo) Having(conds ...gen.Condition) ISeatSmartDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatSmartDo) Limit(limit int) ISeatSmartDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatSmartDo) Offset(offset int) ISeatSmartDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatSmartDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatSmartDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatSmartDo) Unscoped() ISeatSmartDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatSmartDo) Create(values ...*model.SeatSmart) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatSmartDo) CreateInBatches(values []*model.SeatSmart, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatSmartDo) Save(values ...*model.SeatSmart) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatSmartDo) First() (*model.SeatSmart, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSmart), nil
	}
}

func (s seatSmartDo) Take() (*model.SeatSmart, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSmart), nil
	}
}

func (s seatSmartDo) Last() (*model.SeatSmart, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSmart), nil
	}
}

func (s seatSmartDo) Find() ([]*model.SeatSmart, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatSmart), err
}

func (s seatSmartDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatSmart, err error) {
	buf := make([]*model.SeatSmart, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatSmartDo) FindInBatches(result *[]*model.SeatSmart, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatSmartDo) Attrs(attrs ...field.AssignExpr) ISeatSmartDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatSmartDo) Assign(attrs ...field.AssignExpr) ISeatSmartDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatSmartDo) Joins(fields ...field.RelationField) ISeatSmartDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatSmartDo) Preload(fields ...field.RelationField) ISeatSmartDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatSmartDo) FirstOrInit() (*model.SeatSmart, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSmart), nil
	}
}

func (s seatSmartDo) FirstOrCreate() (*model.SeatSmart, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSmart), nil
	}
}

func (s seatSmartDo) FindByPage(offset int, limit int) (result []*model.SeatSmart, count int64, err error) {
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

func (s seatSmartDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatSmartDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatSmartDo) Delete(models ...*model.SeatSmart) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatSmartDo) withDO(do gen.Dao) *seatSmartDo {
	s.DO = *do.(*gen.DO)
	return s
}
