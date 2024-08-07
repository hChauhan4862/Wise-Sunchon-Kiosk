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

func newSeatAdmin(db *gorm.DB, opts ...gen.DOOption) seatAdmin {
	_seatAdmin := seatAdmin{}

	_seatAdmin.seatAdminDo.UseDB(db, opts...)
	_seatAdmin.seatAdminDo.UseModel(&model.SeatAdmin{})

	tableName := _seatAdmin.seatAdminDo.TableName()
	_seatAdmin.ALL = field.NewAsterisk(tableName)
	_seatAdmin.ID = field.NewString(tableName, "id")
	_seatAdmin.PassWd = field.NewString(tableName, "pass_wd")
	_seatAdmin.Auth = field.NewString(tableName, "auth")
	_seatAdmin.RegDt = field.NewString(tableName, "reg_dt")
	_seatAdmin.UpdDt = field.NewString(tableName, "upd_dt")

	_seatAdmin.fillFieldMap()

	return _seatAdmin
}

type seatAdmin struct {
	seatAdminDo

	ALL    field.Asterisk
	ID     field.String
	PassWd field.String
	Auth   field.String
	RegDt  field.String
	UpdDt  field.String

	fieldMap map[string]field.Expr
}

func (s seatAdmin) Table(newTableName string) *seatAdmin {
	s.seatAdminDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatAdmin) As(alias string) *seatAdmin {
	s.seatAdminDo.DO = *(s.seatAdminDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatAdmin) updateTableName(table string) *seatAdmin {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.PassWd = field.NewString(table, "pass_wd")
	s.Auth = field.NewString(table, "auth")
	s.RegDt = field.NewString(table, "reg_dt")
	s.UpdDt = field.NewString(table, "upd_dt")

	s.fillFieldMap()

	return s
}

func (s *seatAdmin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatAdmin) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["pass_wd"] = s.PassWd
	s.fieldMap["auth"] = s.Auth
	s.fieldMap["reg_dt"] = s.RegDt
	s.fieldMap["upd_dt"] = s.UpdDt
}

func (s seatAdmin) clone(db *gorm.DB) seatAdmin {
	s.seatAdminDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatAdmin) replaceDB(db *gorm.DB) seatAdmin {
	s.seatAdminDo.ReplaceDB(db)
	return s
}

type seatAdminDo struct{ gen.DO }

type ISeatAdminDo interface {
	gen.SubQuery
	Debug() ISeatAdminDo
	WithContext(ctx context.Context) ISeatAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatAdminDo
	WriteDB() ISeatAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatAdminDo
	Not(conds ...gen.Condition) ISeatAdminDo
	Or(conds ...gen.Condition) ISeatAdminDo
	Select(conds ...field.Expr) ISeatAdminDo
	Where(conds ...gen.Condition) ISeatAdminDo
	Order(conds ...field.Expr) ISeatAdminDo
	Distinct(cols ...field.Expr) ISeatAdminDo
	Omit(cols ...field.Expr) ISeatAdminDo
	Join(table schema.Tabler, on ...field.Expr) ISeatAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatAdminDo
	Group(cols ...field.Expr) ISeatAdminDo
	Having(conds ...gen.Condition) ISeatAdminDo
	Limit(limit int) ISeatAdminDo
	Offset(offset int) ISeatAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatAdminDo
	Unscoped() ISeatAdminDo
	Create(values ...*model.SeatAdmin) error
	CreateInBatches(values []*model.SeatAdmin, batchSize int) error
	Save(values ...*model.SeatAdmin) error
	First() (*model.SeatAdmin, error)
	Take() (*model.SeatAdmin, error)
	Last() (*model.SeatAdmin, error)
	Find() ([]*model.SeatAdmin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatAdmin, err error)
	FindInBatches(result *[]*model.SeatAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatAdmin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatAdminDo
	Assign(attrs ...field.AssignExpr) ISeatAdminDo
	Joins(fields ...field.RelationField) ISeatAdminDo
	Preload(fields ...field.RelationField) ISeatAdminDo
	FirstOrInit() (*model.SeatAdmin, error)
	FirstOrCreate() (*model.SeatAdmin, error)
	FindByPage(offset int, limit int) (result []*model.SeatAdmin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatAdminDo) Debug() ISeatAdminDo {
	return s.withDO(s.DO.Debug())
}

func (s seatAdminDo) WithContext(ctx context.Context) ISeatAdminDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatAdminDo) ReadDB() ISeatAdminDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatAdminDo) WriteDB() ISeatAdminDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatAdminDo) Session(config *gorm.Session) ISeatAdminDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatAdminDo) Clauses(conds ...clause.Expression) ISeatAdminDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatAdminDo) Returning(value interface{}, columns ...string) ISeatAdminDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatAdminDo) Not(conds ...gen.Condition) ISeatAdminDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatAdminDo) Or(conds ...gen.Condition) ISeatAdminDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatAdminDo) Select(conds ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatAdminDo) Where(conds ...gen.Condition) ISeatAdminDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatAdminDo) Order(conds ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatAdminDo) Distinct(cols ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatAdminDo) Omit(cols ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatAdminDo) Join(table schema.Tabler, on ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatAdminDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatAdminDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatAdminDo) Group(cols ...field.Expr) ISeatAdminDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatAdminDo) Having(conds ...gen.Condition) ISeatAdminDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatAdminDo) Limit(limit int) ISeatAdminDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatAdminDo) Offset(offset int) ISeatAdminDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatAdminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatAdminDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatAdminDo) Unscoped() ISeatAdminDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatAdminDo) Create(values ...*model.SeatAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatAdminDo) CreateInBatches(values []*model.SeatAdmin, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatAdminDo) Save(values ...*model.SeatAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatAdminDo) First() (*model.SeatAdmin, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatAdmin), nil
	}
}

func (s seatAdminDo) Take() (*model.SeatAdmin, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatAdmin), nil
	}
}

func (s seatAdminDo) Last() (*model.SeatAdmin, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatAdmin), nil
	}
}

func (s seatAdminDo) Find() ([]*model.SeatAdmin, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatAdmin), err
}

func (s seatAdminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatAdmin, err error) {
	buf := make([]*model.SeatAdmin, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatAdminDo) FindInBatches(result *[]*model.SeatAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatAdminDo) Attrs(attrs ...field.AssignExpr) ISeatAdminDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatAdminDo) Assign(attrs ...field.AssignExpr) ISeatAdminDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatAdminDo) Joins(fields ...field.RelationField) ISeatAdminDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatAdminDo) Preload(fields ...field.RelationField) ISeatAdminDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatAdminDo) FirstOrInit() (*model.SeatAdmin, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatAdmin), nil
	}
}

func (s seatAdminDo) FirstOrCreate() (*model.SeatAdmin, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatAdmin), nil
	}
}

func (s seatAdminDo) FindByPage(offset int, limit int) (result []*model.SeatAdmin, count int64, err error) {
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

func (s seatAdminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatAdminDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatAdminDo) Delete(models ...*model.SeatAdmin) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatAdminDo) withDO(do gen.Dao) *seatAdminDo {
	s.DO = *do.(*gen.DO)
	return s
}
