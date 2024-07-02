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

func newSeatSeatExt(db *gorm.DB, opts ...gen.DOOption) seatSeatExt {
	_seatSeatExt := seatSeatExt{}

	_seatSeatExt.seatSeatExtDo.UseDB(db, opts...)
	_seatSeatExt.seatSeatExtDo.UseModel(&model.SeatSeatExt{})

	tableName := _seatSeatExt.seatSeatExtDo.TableName()
	_seatSeatExt.ALL = field.NewAsterisk(tableName)
	_seatSeatExt.SEATNO = field.NewInt64(tableName, "SEATNO")
	_seatSeatExt.MONITORYN = field.NewString(tableName, "MONITOR_YN")
	_seatSeatExt.LANYN = field.NewString(tableName, "LAN_YN")
	_seatSeatExt.STANDYN = field.NewString(tableName, "STAND_YN")
	_seatSeatExt.POWERYN = field.NewString(tableName, "POWER_YN")
	_seatSeatExt.COMMENT = field.NewString(tableName, "COMMENT")
	_seatSeatExt.ENCOMMENT = field.NewString(tableName, "EN_COMMENT")
	_seatSeatExt.SEATIMAGE = field.NewString(tableName, "SEAT_IMAGE")

	_seatSeatExt.fillFieldMap()

	return _seatSeatExt
}

type seatSeatExt struct {
	seatSeatExtDo

	ALL       field.Asterisk
	SEATNO    field.Int64
	MONITORYN field.String
	LANYN     field.String
	STANDYN   field.String
	POWERYN   field.String
	COMMENT   field.String
	ENCOMMENT field.String
	SEATIMAGE field.String

	fieldMap map[string]field.Expr
}

func (s seatSeatExt) Table(newTableName string) *seatSeatExt {
	s.seatSeatExtDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatSeatExt) As(alias string) *seatSeatExt {
	s.seatSeatExtDo.DO = *(s.seatSeatExtDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatSeatExt) updateTableName(table string) *seatSeatExt {
	s.ALL = field.NewAsterisk(table)
	s.SEATNO = field.NewInt64(table, "SEATNO")
	s.MONITORYN = field.NewString(table, "MONITOR_YN")
	s.LANYN = field.NewString(table, "LAN_YN")
	s.STANDYN = field.NewString(table, "STAND_YN")
	s.POWERYN = field.NewString(table, "POWER_YN")
	s.COMMENT = field.NewString(table, "COMMENT")
	s.ENCOMMENT = field.NewString(table, "EN_COMMENT")
	s.SEATIMAGE = field.NewString(table, "SEAT_IMAGE")

	s.fillFieldMap()

	return s
}

func (s *seatSeatExt) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatSeatExt) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["SEATNO"] = s.SEATNO
	s.fieldMap["MONITOR_YN"] = s.MONITORYN
	s.fieldMap["LAN_YN"] = s.LANYN
	s.fieldMap["STAND_YN"] = s.STANDYN
	s.fieldMap["POWER_YN"] = s.POWERYN
	s.fieldMap["COMMENT"] = s.COMMENT
	s.fieldMap["EN_COMMENT"] = s.ENCOMMENT
	s.fieldMap["SEAT_IMAGE"] = s.SEATIMAGE
}

func (s seatSeatExt) clone(db *gorm.DB) seatSeatExt {
	s.seatSeatExtDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatSeatExt) replaceDB(db *gorm.DB) seatSeatExt {
	s.seatSeatExtDo.ReplaceDB(db)
	return s
}

type seatSeatExtDo struct{ gen.DO }

type ISeatSeatExtDo interface {
	gen.SubQuery
	Debug() ISeatSeatExtDo
	WithContext(ctx context.Context) ISeatSeatExtDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatSeatExtDo
	WriteDB() ISeatSeatExtDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatSeatExtDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatSeatExtDo
	Not(conds ...gen.Condition) ISeatSeatExtDo
	Or(conds ...gen.Condition) ISeatSeatExtDo
	Select(conds ...field.Expr) ISeatSeatExtDo
	Where(conds ...gen.Condition) ISeatSeatExtDo
	Order(conds ...field.Expr) ISeatSeatExtDo
	Distinct(cols ...field.Expr) ISeatSeatExtDo
	Omit(cols ...field.Expr) ISeatSeatExtDo
	Join(table schema.Tabler, on ...field.Expr) ISeatSeatExtDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatSeatExtDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatSeatExtDo
	Group(cols ...field.Expr) ISeatSeatExtDo
	Having(conds ...gen.Condition) ISeatSeatExtDo
	Limit(limit int) ISeatSeatExtDo
	Offset(offset int) ISeatSeatExtDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatSeatExtDo
	Unscoped() ISeatSeatExtDo
	Create(values ...*model.SeatSeatExt) error
	CreateInBatches(values []*model.SeatSeatExt, batchSize int) error
	Save(values ...*model.SeatSeatExt) error
	First() (*model.SeatSeatExt, error)
	Take() (*model.SeatSeatExt, error)
	Last() (*model.SeatSeatExt, error)
	Find() ([]*model.SeatSeatExt, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatSeatExt, err error)
	FindInBatches(result *[]*model.SeatSeatExt, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatSeatExt) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatSeatExtDo
	Assign(attrs ...field.AssignExpr) ISeatSeatExtDo
	Joins(fields ...field.RelationField) ISeatSeatExtDo
	Preload(fields ...field.RelationField) ISeatSeatExtDo
	FirstOrInit() (*model.SeatSeatExt, error)
	FirstOrCreate() (*model.SeatSeatExt, error)
	FindByPage(offset int, limit int) (result []*model.SeatSeatExt, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatSeatExtDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatSeatExtDo) Debug() ISeatSeatExtDo {
	return s.withDO(s.DO.Debug())
}

func (s seatSeatExtDo) WithContext(ctx context.Context) ISeatSeatExtDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatSeatExtDo) ReadDB() ISeatSeatExtDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatSeatExtDo) WriteDB() ISeatSeatExtDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatSeatExtDo) Session(config *gorm.Session) ISeatSeatExtDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatSeatExtDo) Clauses(conds ...clause.Expression) ISeatSeatExtDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatSeatExtDo) Returning(value interface{}, columns ...string) ISeatSeatExtDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatSeatExtDo) Not(conds ...gen.Condition) ISeatSeatExtDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatSeatExtDo) Or(conds ...gen.Condition) ISeatSeatExtDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatSeatExtDo) Select(conds ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatSeatExtDo) Where(conds ...gen.Condition) ISeatSeatExtDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatSeatExtDo) Order(conds ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatSeatExtDo) Distinct(cols ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatSeatExtDo) Omit(cols ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatSeatExtDo) Join(table schema.Tabler, on ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatSeatExtDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatSeatExtDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatSeatExtDo) Group(cols ...field.Expr) ISeatSeatExtDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatSeatExtDo) Having(conds ...gen.Condition) ISeatSeatExtDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatSeatExtDo) Limit(limit int) ISeatSeatExtDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatSeatExtDo) Offset(offset int) ISeatSeatExtDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatSeatExtDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatSeatExtDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatSeatExtDo) Unscoped() ISeatSeatExtDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatSeatExtDo) Create(values ...*model.SeatSeatExt) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatSeatExtDo) CreateInBatches(values []*model.SeatSeatExt, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatSeatExtDo) Save(values ...*model.SeatSeatExt) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatSeatExtDo) First() (*model.SeatSeatExt, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSeatExt), nil
	}
}

func (s seatSeatExtDo) Take() (*model.SeatSeatExt, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSeatExt), nil
	}
}

func (s seatSeatExtDo) Last() (*model.SeatSeatExt, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSeatExt), nil
	}
}

func (s seatSeatExtDo) Find() ([]*model.SeatSeatExt, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatSeatExt), err
}

func (s seatSeatExtDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatSeatExt, err error) {
	buf := make([]*model.SeatSeatExt, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatSeatExtDo) FindInBatches(result *[]*model.SeatSeatExt, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatSeatExtDo) Attrs(attrs ...field.AssignExpr) ISeatSeatExtDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatSeatExtDo) Assign(attrs ...field.AssignExpr) ISeatSeatExtDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatSeatExtDo) Joins(fields ...field.RelationField) ISeatSeatExtDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatSeatExtDo) Preload(fields ...field.RelationField) ISeatSeatExtDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatSeatExtDo) FirstOrInit() (*model.SeatSeatExt, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSeatExt), nil
	}
}

func (s seatSeatExtDo) FirstOrCreate() (*model.SeatSeatExt, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSeatExt), nil
	}
}

func (s seatSeatExtDo) FindByPage(offset int, limit int) (result []*model.SeatSeatExt, count int64, err error) {
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

func (s seatSeatExtDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatSeatExtDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatSeatExtDo) Delete(models ...*model.SeatSeatExt) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatSeatExtDo) withDO(do gen.Dao) *seatSeatExtDo {
	s.DO = *do.(*gen.DO)
	return s
}