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

func newSeatRoom2Ext(db *gorm.DB, opts ...gen.DOOption) seatRoom2Ext {
	_seatRoom2Ext := seatRoom2Ext{}

	_seatRoom2Ext.seatRoom2ExtDo.UseDB(db, opts...)
	_seatRoom2Ext.seatRoom2ExtDo.UseModel(&model.SeatRoom2Ext{})

	tableName := _seatRoom2Ext.seatRoom2ExtDo.TableName()
	_seatRoom2Ext.ALL = field.NewAsterisk(tableName)
	_seatRoom2Ext.ROOMNO = field.NewInt64(tableName, "ROOMNO")
	_seatRoom2Ext.Sort = field.NewInt64(tableName, "sort")
	_seatRoom2Ext.TimeStart = field.NewString(tableName, "time_start")
	_seatRoom2Ext.TimeEnd = field.NewString(tableName, "time_end")
	_seatRoom2Ext.KioskYn = field.NewString(tableName, "kiosk_yn")
	_seatRoom2Ext.WebYn = field.NewString(tableName, "web_yn")
	_seatRoom2Ext.MobileYn = field.NewString(tableName, "mobile_yn")

	_seatRoom2Ext.fillFieldMap()

	return _seatRoom2Ext
}

type seatRoom2Ext struct {
	seatRoom2ExtDo

	ALL       field.Asterisk
	ROOMNO    field.Int64
	Sort      field.Int64
	TimeStart field.String
	TimeEnd   field.String
	KioskYn   field.String
	WebYn     field.String
	MobileYn  field.String

	fieldMap map[string]field.Expr
}

func (s seatRoom2Ext) Table(newTableName string) *seatRoom2Ext {
	s.seatRoom2ExtDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatRoom2Ext) As(alias string) *seatRoom2Ext {
	s.seatRoom2ExtDo.DO = *(s.seatRoom2ExtDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatRoom2Ext) updateTableName(table string) *seatRoom2Ext {
	s.ALL = field.NewAsterisk(table)
	s.ROOMNO = field.NewInt64(table, "ROOMNO")
	s.Sort = field.NewInt64(table, "sort")
	s.TimeStart = field.NewString(table, "time_start")
	s.TimeEnd = field.NewString(table, "time_end")
	s.KioskYn = field.NewString(table, "kiosk_yn")
	s.WebYn = field.NewString(table, "web_yn")
	s.MobileYn = field.NewString(table, "mobile_yn")

	s.fillFieldMap()

	return s
}

func (s *seatRoom2Ext) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatRoom2Ext) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["ROOMNO"] = s.ROOMNO
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["time_start"] = s.TimeStart
	s.fieldMap["time_end"] = s.TimeEnd
	s.fieldMap["kiosk_yn"] = s.KioskYn
	s.fieldMap["web_yn"] = s.WebYn
	s.fieldMap["mobile_yn"] = s.MobileYn
}

func (s seatRoom2Ext) clone(db *gorm.DB) seatRoom2Ext {
	s.seatRoom2ExtDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatRoom2Ext) replaceDB(db *gorm.DB) seatRoom2Ext {
	s.seatRoom2ExtDo.ReplaceDB(db)
	return s
}

type seatRoom2ExtDo struct{ gen.DO }

type ISeatRoom2ExtDo interface {
	gen.SubQuery
	Debug() ISeatRoom2ExtDo
	WithContext(ctx context.Context) ISeatRoom2ExtDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatRoom2ExtDo
	WriteDB() ISeatRoom2ExtDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatRoom2ExtDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatRoom2ExtDo
	Not(conds ...gen.Condition) ISeatRoom2ExtDo
	Or(conds ...gen.Condition) ISeatRoom2ExtDo
	Select(conds ...field.Expr) ISeatRoom2ExtDo
	Where(conds ...gen.Condition) ISeatRoom2ExtDo
	Order(conds ...field.Expr) ISeatRoom2ExtDo
	Distinct(cols ...field.Expr) ISeatRoom2ExtDo
	Omit(cols ...field.Expr) ISeatRoom2ExtDo
	Join(table schema.Tabler, on ...field.Expr) ISeatRoom2ExtDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatRoom2ExtDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatRoom2ExtDo
	Group(cols ...field.Expr) ISeatRoom2ExtDo
	Having(conds ...gen.Condition) ISeatRoom2ExtDo
	Limit(limit int) ISeatRoom2ExtDo
	Offset(offset int) ISeatRoom2ExtDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatRoom2ExtDo
	Unscoped() ISeatRoom2ExtDo
	Create(values ...*model.SeatRoom2Ext) error
	CreateInBatches(values []*model.SeatRoom2Ext, batchSize int) error
	Save(values ...*model.SeatRoom2Ext) error
	First() (*model.SeatRoom2Ext, error)
	Take() (*model.SeatRoom2Ext, error)
	Last() (*model.SeatRoom2Ext, error)
	Find() ([]*model.SeatRoom2Ext, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatRoom2Ext, err error)
	FindInBatches(result *[]*model.SeatRoom2Ext, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatRoom2Ext) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatRoom2ExtDo
	Assign(attrs ...field.AssignExpr) ISeatRoom2ExtDo
	Joins(fields ...field.RelationField) ISeatRoom2ExtDo
	Preload(fields ...field.RelationField) ISeatRoom2ExtDo
	FirstOrInit() (*model.SeatRoom2Ext, error)
	FirstOrCreate() (*model.SeatRoom2Ext, error)
	FindByPage(offset int, limit int) (result []*model.SeatRoom2Ext, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatRoom2ExtDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatRoom2ExtDo) Debug() ISeatRoom2ExtDo {
	return s.withDO(s.DO.Debug())
}

func (s seatRoom2ExtDo) WithContext(ctx context.Context) ISeatRoom2ExtDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatRoom2ExtDo) ReadDB() ISeatRoom2ExtDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatRoom2ExtDo) WriteDB() ISeatRoom2ExtDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatRoom2ExtDo) Session(config *gorm.Session) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatRoom2ExtDo) Clauses(conds ...clause.Expression) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatRoom2ExtDo) Returning(value interface{}, columns ...string) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatRoom2ExtDo) Not(conds ...gen.Condition) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatRoom2ExtDo) Or(conds ...gen.Condition) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatRoom2ExtDo) Select(conds ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatRoom2ExtDo) Where(conds ...gen.Condition) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatRoom2ExtDo) Order(conds ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatRoom2ExtDo) Distinct(cols ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatRoom2ExtDo) Omit(cols ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatRoom2ExtDo) Join(table schema.Tabler, on ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatRoom2ExtDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatRoom2ExtDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatRoom2ExtDo) Group(cols ...field.Expr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatRoom2ExtDo) Having(conds ...gen.Condition) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatRoom2ExtDo) Limit(limit int) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatRoom2ExtDo) Offset(offset int) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatRoom2ExtDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatRoom2ExtDo) Unscoped() ISeatRoom2ExtDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatRoom2ExtDo) Create(values ...*model.SeatRoom2Ext) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatRoom2ExtDo) CreateInBatches(values []*model.SeatRoom2Ext, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatRoom2ExtDo) Save(values ...*model.SeatRoom2Ext) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatRoom2ExtDo) First() (*model.SeatRoom2Ext, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoom2Ext), nil
	}
}

func (s seatRoom2ExtDo) Take() (*model.SeatRoom2Ext, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoom2Ext), nil
	}
}

func (s seatRoom2ExtDo) Last() (*model.SeatRoom2Ext, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoom2Ext), nil
	}
}

func (s seatRoom2ExtDo) Find() ([]*model.SeatRoom2Ext, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatRoom2Ext), err
}

func (s seatRoom2ExtDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatRoom2Ext, err error) {
	buf := make([]*model.SeatRoom2Ext, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatRoom2ExtDo) FindInBatches(result *[]*model.SeatRoom2Ext, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatRoom2ExtDo) Attrs(attrs ...field.AssignExpr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatRoom2ExtDo) Assign(attrs ...field.AssignExpr) ISeatRoom2ExtDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatRoom2ExtDo) Joins(fields ...field.RelationField) ISeatRoom2ExtDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatRoom2ExtDo) Preload(fields ...field.RelationField) ISeatRoom2ExtDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatRoom2ExtDo) FirstOrInit() (*model.SeatRoom2Ext, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoom2Ext), nil
	}
}

func (s seatRoom2ExtDo) FirstOrCreate() (*model.SeatRoom2Ext, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatRoom2Ext), nil
	}
}

func (s seatRoom2ExtDo) FindByPage(offset int, limit int) (result []*model.SeatRoom2Ext, count int64, err error) {
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

func (s seatRoom2ExtDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatRoom2ExtDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatRoom2ExtDo) Delete(models ...*model.SeatRoom2Ext) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatRoom2ExtDo) withDO(do gen.Dao) *seatRoom2ExtDo {
	s.DO = *do.(*gen.DO)
	return s
}