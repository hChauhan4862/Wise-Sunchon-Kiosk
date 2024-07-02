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

func newSeatMulti(db *gorm.DB, opts ...gen.DOOption) seatMulti {
	_seatMulti := seatMulti{}

	_seatMulti.seatMultiDo.UseDB(db, opts...)
	_seatMulti.seatMultiDo.UseModel(&model.SeatMulti{})

	tableName := _seatMulti.seatMultiDo.TableName()
	_seatMulti.ALL = field.NewAsterisk(tableName)
	_seatMulti.RoomNo = field.NewInt64(tableName, "room_no")
	_seatMulti.MultiLimitTime = field.NewString(tableName, "multi_limit_time")
	_seatMulti.MultiMassage = field.NewString(tableName, "multi_massage")
	_seatMulti.MultiMessagemain = field.NewString(tableName, "multi_messagemain")
	_seatMulti.MultiMessagerefund = field.NewString(tableName, "multi_messagerefund")
	_seatMulti.UseYn = field.NewString(tableName, "use_yn")
	_seatMulti.SystemEndYn = field.NewString(tableName, "system_end_yn")
	_seatMulti.SystemEndTime = field.NewString(tableName, "system_end_time")

	_seatMulti.fillFieldMap()

	return _seatMulti
}

type seatMulti struct {
	seatMultiDo

	ALL                field.Asterisk
	RoomNo             field.Int64
	MultiLimitTime     field.String
	MultiMassage       field.String
	MultiMessagemain   field.String
	MultiMessagerefund field.String
	UseYn              field.String
	SystemEndYn        field.String
	SystemEndTime      field.String

	fieldMap map[string]field.Expr
}

func (s seatMulti) Table(newTableName string) *seatMulti {
	s.seatMultiDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatMulti) As(alias string) *seatMulti {
	s.seatMultiDo.DO = *(s.seatMultiDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatMulti) updateTableName(table string) *seatMulti {
	s.ALL = field.NewAsterisk(table)
	s.RoomNo = field.NewInt64(table, "room_no")
	s.MultiLimitTime = field.NewString(table, "multi_limit_time")
	s.MultiMassage = field.NewString(table, "multi_massage")
	s.MultiMessagemain = field.NewString(table, "multi_messagemain")
	s.MultiMessagerefund = field.NewString(table, "multi_messagerefund")
	s.UseYn = field.NewString(table, "use_yn")
	s.SystemEndYn = field.NewString(table, "system_end_yn")
	s.SystemEndTime = field.NewString(table, "system_end_time")

	s.fillFieldMap()

	return s
}

func (s *seatMulti) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatMulti) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["room_no"] = s.RoomNo
	s.fieldMap["multi_limit_time"] = s.MultiLimitTime
	s.fieldMap["multi_massage"] = s.MultiMassage
	s.fieldMap["multi_messagemain"] = s.MultiMessagemain
	s.fieldMap["multi_messagerefund"] = s.MultiMessagerefund
	s.fieldMap["use_yn"] = s.UseYn
	s.fieldMap["system_end_yn"] = s.SystemEndYn
	s.fieldMap["system_end_time"] = s.SystemEndTime
}

func (s seatMulti) clone(db *gorm.DB) seatMulti {
	s.seatMultiDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatMulti) replaceDB(db *gorm.DB) seatMulti {
	s.seatMultiDo.ReplaceDB(db)
	return s
}

type seatMultiDo struct{ gen.DO }

type ISeatMultiDo interface {
	gen.SubQuery
	Debug() ISeatMultiDo
	WithContext(ctx context.Context) ISeatMultiDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatMultiDo
	WriteDB() ISeatMultiDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatMultiDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatMultiDo
	Not(conds ...gen.Condition) ISeatMultiDo
	Or(conds ...gen.Condition) ISeatMultiDo
	Select(conds ...field.Expr) ISeatMultiDo
	Where(conds ...gen.Condition) ISeatMultiDo
	Order(conds ...field.Expr) ISeatMultiDo
	Distinct(cols ...field.Expr) ISeatMultiDo
	Omit(cols ...field.Expr) ISeatMultiDo
	Join(table schema.Tabler, on ...field.Expr) ISeatMultiDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatMultiDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatMultiDo
	Group(cols ...field.Expr) ISeatMultiDo
	Having(conds ...gen.Condition) ISeatMultiDo
	Limit(limit int) ISeatMultiDo
	Offset(offset int) ISeatMultiDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatMultiDo
	Unscoped() ISeatMultiDo
	Create(values ...*model.SeatMulti) error
	CreateInBatches(values []*model.SeatMulti, batchSize int) error
	Save(values ...*model.SeatMulti) error
	First() (*model.SeatMulti, error)
	Take() (*model.SeatMulti, error)
	Last() (*model.SeatMulti, error)
	Find() ([]*model.SeatMulti, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatMulti, err error)
	FindInBatches(result *[]*model.SeatMulti, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatMulti) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatMultiDo
	Assign(attrs ...field.AssignExpr) ISeatMultiDo
	Joins(fields ...field.RelationField) ISeatMultiDo
	Preload(fields ...field.RelationField) ISeatMultiDo
	FirstOrInit() (*model.SeatMulti, error)
	FirstOrCreate() (*model.SeatMulti, error)
	FindByPage(offset int, limit int) (result []*model.SeatMulti, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatMultiDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatMultiDo) Debug() ISeatMultiDo {
	return s.withDO(s.DO.Debug())
}

func (s seatMultiDo) WithContext(ctx context.Context) ISeatMultiDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatMultiDo) ReadDB() ISeatMultiDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatMultiDo) WriteDB() ISeatMultiDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatMultiDo) Session(config *gorm.Session) ISeatMultiDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatMultiDo) Clauses(conds ...clause.Expression) ISeatMultiDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatMultiDo) Returning(value interface{}, columns ...string) ISeatMultiDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatMultiDo) Not(conds ...gen.Condition) ISeatMultiDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatMultiDo) Or(conds ...gen.Condition) ISeatMultiDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatMultiDo) Select(conds ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatMultiDo) Where(conds ...gen.Condition) ISeatMultiDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatMultiDo) Order(conds ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatMultiDo) Distinct(cols ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatMultiDo) Omit(cols ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatMultiDo) Join(table schema.Tabler, on ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatMultiDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatMultiDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatMultiDo) Group(cols ...field.Expr) ISeatMultiDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatMultiDo) Having(conds ...gen.Condition) ISeatMultiDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatMultiDo) Limit(limit int) ISeatMultiDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatMultiDo) Offset(offset int) ISeatMultiDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatMultiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatMultiDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatMultiDo) Unscoped() ISeatMultiDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatMultiDo) Create(values ...*model.SeatMulti) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatMultiDo) CreateInBatches(values []*model.SeatMulti, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatMultiDo) Save(values ...*model.SeatMulti) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatMultiDo) First() (*model.SeatMulti, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatMulti), nil
	}
}

func (s seatMultiDo) Take() (*model.SeatMulti, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatMulti), nil
	}
}

func (s seatMultiDo) Last() (*model.SeatMulti, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatMulti), nil
	}
}

func (s seatMultiDo) Find() ([]*model.SeatMulti, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatMulti), err
}

func (s seatMultiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatMulti, err error) {
	buf := make([]*model.SeatMulti, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatMultiDo) FindInBatches(result *[]*model.SeatMulti, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatMultiDo) Attrs(attrs ...field.AssignExpr) ISeatMultiDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatMultiDo) Assign(attrs ...field.AssignExpr) ISeatMultiDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatMultiDo) Joins(fields ...field.RelationField) ISeatMultiDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatMultiDo) Preload(fields ...field.RelationField) ISeatMultiDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatMultiDo) FirstOrInit() (*model.SeatMulti, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatMulti), nil
	}
}

func (s seatMultiDo) FirstOrCreate() (*model.SeatMulti, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatMulti), nil
	}
}

func (s seatMultiDo) FindByPage(offset int, limit int) (result []*model.SeatMulti, count int64, err error) {
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

func (s seatMultiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatMultiDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatMultiDo) Delete(models ...*model.SeatMulti) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatMultiDo) withDO(do gen.Dao) *seatMultiDo {
	s.DO = *do.(*gen.DO)
	return s
}
