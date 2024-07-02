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

func newSeatActlog(db *gorm.DB, opts ...gen.DOOption) seatActlog {
	_seatActlog := seatActlog{}

	_seatActlog.seatActlogDo.UseDB(db, opts...)
	_seatActlog.seatActlogDo.UseModel(&model.SeatActlog{})

	tableName := _seatActlog.seatActlogDo.TableName()
	_seatActlog.ALL = field.NewAsterisk(tableName)
	_seatActlog.ActTime = field.NewString(tableName, "act_time")
	_seatActlog.UserID = field.NewString(tableName, "user_id")
	_seatActlog.ActGubun = field.NewInt64(tableName, "act_gubun")
	_seatActlog.PatType = field.NewString(tableName, "pat_type")
	_seatActlog.DeptCode = field.NewString(tableName, "dept_code")
	_seatActlog.KioskNo = field.NewInt64(tableName, "kiosk_no")
	_seatActlog.TimeZone = field.NewString(tableName, "time_zone")

	_seatActlog.fillFieldMap()

	return _seatActlog
}

type seatActlog struct {
	seatActlogDo

	ALL      field.Asterisk
	ActTime  field.String
	UserID   field.String
	ActGubun field.Int64
	PatType  field.String
	DeptCode field.String
	KioskNo  field.Int64
	TimeZone field.String

	fieldMap map[string]field.Expr
}

func (s seatActlog) Table(newTableName string) *seatActlog {
	s.seatActlogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatActlog) As(alias string) *seatActlog {
	s.seatActlogDo.DO = *(s.seatActlogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatActlog) updateTableName(table string) *seatActlog {
	s.ALL = field.NewAsterisk(table)
	s.ActTime = field.NewString(table, "act_time")
	s.UserID = field.NewString(table, "user_id")
	s.ActGubun = field.NewInt64(table, "act_gubun")
	s.PatType = field.NewString(table, "pat_type")
	s.DeptCode = field.NewString(table, "dept_code")
	s.KioskNo = field.NewInt64(table, "kiosk_no")
	s.TimeZone = field.NewString(table, "time_zone")

	s.fillFieldMap()

	return s
}

func (s *seatActlog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatActlog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["act_time"] = s.ActTime
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["act_gubun"] = s.ActGubun
	s.fieldMap["pat_type"] = s.PatType
	s.fieldMap["dept_code"] = s.DeptCode
	s.fieldMap["kiosk_no"] = s.KioskNo
	s.fieldMap["time_zone"] = s.TimeZone
}

func (s seatActlog) clone(db *gorm.DB) seatActlog {
	s.seatActlogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatActlog) replaceDB(db *gorm.DB) seatActlog {
	s.seatActlogDo.ReplaceDB(db)
	return s
}

type seatActlogDo struct{ gen.DO }

type ISeatActlogDo interface {
	gen.SubQuery
	Debug() ISeatActlogDo
	WithContext(ctx context.Context) ISeatActlogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatActlogDo
	WriteDB() ISeatActlogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatActlogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatActlogDo
	Not(conds ...gen.Condition) ISeatActlogDo
	Or(conds ...gen.Condition) ISeatActlogDo
	Select(conds ...field.Expr) ISeatActlogDo
	Where(conds ...gen.Condition) ISeatActlogDo
	Order(conds ...field.Expr) ISeatActlogDo
	Distinct(cols ...field.Expr) ISeatActlogDo
	Omit(cols ...field.Expr) ISeatActlogDo
	Join(table schema.Tabler, on ...field.Expr) ISeatActlogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatActlogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatActlogDo
	Group(cols ...field.Expr) ISeatActlogDo
	Having(conds ...gen.Condition) ISeatActlogDo
	Limit(limit int) ISeatActlogDo
	Offset(offset int) ISeatActlogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatActlogDo
	Unscoped() ISeatActlogDo
	Create(values ...*model.SeatActlog) error
	CreateInBatches(values []*model.SeatActlog, batchSize int) error
	Save(values ...*model.SeatActlog) error
	First() (*model.SeatActlog, error)
	Take() (*model.SeatActlog, error)
	Last() (*model.SeatActlog, error)
	Find() ([]*model.SeatActlog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatActlog, err error)
	FindInBatches(result *[]*model.SeatActlog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatActlog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatActlogDo
	Assign(attrs ...field.AssignExpr) ISeatActlogDo
	Joins(fields ...field.RelationField) ISeatActlogDo
	Preload(fields ...field.RelationField) ISeatActlogDo
	FirstOrInit() (*model.SeatActlog, error)
	FirstOrCreate() (*model.SeatActlog, error)
	FindByPage(offset int, limit int) (result []*model.SeatActlog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatActlogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatActlogDo) Debug() ISeatActlogDo {
	return s.withDO(s.DO.Debug())
}

func (s seatActlogDo) WithContext(ctx context.Context) ISeatActlogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatActlogDo) ReadDB() ISeatActlogDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatActlogDo) WriteDB() ISeatActlogDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatActlogDo) Session(config *gorm.Session) ISeatActlogDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatActlogDo) Clauses(conds ...clause.Expression) ISeatActlogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatActlogDo) Returning(value interface{}, columns ...string) ISeatActlogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatActlogDo) Not(conds ...gen.Condition) ISeatActlogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatActlogDo) Or(conds ...gen.Condition) ISeatActlogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatActlogDo) Select(conds ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatActlogDo) Where(conds ...gen.Condition) ISeatActlogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatActlogDo) Order(conds ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatActlogDo) Distinct(cols ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatActlogDo) Omit(cols ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatActlogDo) Join(table schema.Tabler, on ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatActlogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatActlogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatActlogDo) Group(cols ...field.Expr) ISeatActlogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatActlogDo) Having(conds ...gen.Condition) ISeatActlogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatActlogDo) Limit(limit int) ISeatActlogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatActlogDo) Offset(offset int) ISeatActlogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatActlogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatActlogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatActlogDo) Unscoped() ISeatActlogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatActlogDo) Create(values ...*model.SeatActlog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatActlogDo) CreateInBatches(values []*model.SeatActlog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatActlogDo) Save(values ...*model.SeatActlog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatActlogDo) First() (*model.SeatActlog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatActlog), nil
	}
}

func (s seatActlogDo) Take() (*model.SeatActlog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatActlog), nil
	}
}

func (s seatActlogDo) Last() (*model.SeatActlog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatActlog), nil
	}
}

func (s seatActlogDo) Find() ([]*model.SeatActlog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatActlog), err
}

func (s seatActlogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatActlog, err error) {
	buf := make([]*model.SeatActlog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatActlogDo) FindInBatches(result *[]*model.SeatActlog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatActlogDo) Attrs(attrs ...field.AssignExpr) ISeatActlogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatActlogDo) Assign(attrs ...field.AssignExpr) ISeatActlogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatActlogDo) Joins(fields ...field.RelationField) ISeatActlogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatActlogDo) Preload(fields ...field.RelationField) ISeatActlogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatActlogDo) FirstOrInit() (*model.SeatActlog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatActlog), nil
	}
}

func (s seatActlogDo) FirstOrCreate() (*model.SeatActlog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatActlog), nil
	}
}

func (s seatActlogDo) FindByPage(offset int, limit int) (result []*model.SeatActlog, count int64, err error) {
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

func (s seatActlogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatActlogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatActlogDo) Delete(models ...*model.SeatActlog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatActlogDo) withDO(do gen.Dao) *seatActlogDo {
	s.DO = *do.(*gen.DO)
	return s
}
