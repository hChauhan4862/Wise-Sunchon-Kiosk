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

func newSeatDept(db *gorm.DB, opts ...gen.DOOption) seatDept {
	_seatDept := seatDept{}

	_seatDept.seatDeptDo.UseDB(db, opts...)
	_seatDept.seatDeptDo.UseModel(&model.SeatDept{})

	tableName := _seatDept.seatDeptDo.TableName()
	_seatDept.ALL = field.NewAsterisk(tableName)
	_seatDept.DeptCode = field.NewString(tableName, "dept_code")
	_seatDept.DeptName = field.NewString(tableName, "dept_name")

	_seatDept.fillFieldMap()

	return _seatDept
}

type seatDept struct {
	seatDeptDo

	ALL      field.Asterisk
	DeptCode field.String
	DeptName field.String

	fieldMap map[string]field.Expr
}

func (s seatDept) Table(newTableName string) *seatDept {
	s.seatDeptDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatDept) As(alias string) *seatDept {
	s.seatDeptDo.DO = *(s.seatDeptDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatDept) updateTableName(table string) *seatDept {
	s.ALL = field.NewAsterisk(table)
	s.DeptCode = field.NewString(table, "dept_code")
	s.DeptName = field.NewString(table, "dept_name")

	s.fillFieldMap()

	return s
}

func (s *seatDept) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatDept) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["dept_code"] = s.DeptCode
	s.fieldMap["dept_name"] = s.DeptName
}

func (s seatDept) clone(db *gorm.DB) seatDept {
	s.seatDeptDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatDept) replaceDB(db *gorm.DB) seatDept {
	s.seatDeptDo.ReplaceDB(db)
	return s
}

type seatDeptDo struct{ gen.DO }

type ISeatDeptDo interface {
	gen.SubQuery
	Debug() ISeatDeptDo
	WithContext(ctx context.Context) ISeatDeptDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatDeptDo
	WriteDB() ISeatDeptDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatDeptDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatDeptDo
	Not(conds ...gen.Condition) ISeatDeptDo
	Or(conds ...gen.Condition) ISeatDeptDo
	Select(conds ...field.Expr) ISeatDeptDo
	Where(conds ...gen.Condition) ISeatDeptDo
	Order(conds ...field.Expr) ISeatDeptDo
	Distinct(cols ...field.Expr) ISeatDeptDo
	Omit(cols ...field.Expr) ISeatDeptDo
	Join(table schema.Tabler, on ...field.Expr) ISeatDeptDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatDeptDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatDeptDo
	Group(cols ...field.Expr) ISeatDeptDo
	Having(conds ...gen.Condition) ISeatDeptDo
	Limit(limit int) ISeatDeptDo
	Offset(offset int) ISeatDeptDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatDeptDo
	Unscoped() ISeatDeptDo
	Create(values ...*model.SeatDept) error
	CreateInBatches(values []*model.SeatDept, batchSize int) error
	Save(values ...*model.SeatDept) error
	First() (*model.SeatDept, error)
	Take() (*model.SeatDept, error)
	Last() (*model.SeatDept, error)
	Find() ([]*model.SeatDept, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatDept, err error)
	FindInBatches(result *[]*model.SeatDept, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatDept) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatDeptDo
	Assign(attrs ...field.AssignExpr) ISeatDeptDo
	Joins(fields ...field.RelationField) ISeatDeptDo
	Preload(fields ...field.RelationField) ISeatDeptDo
	FirstOrInit() (*model.SeatDept, error)
	FirstOrCreate() (*model.SeatDept, error)
	FindByPage(offset int, limit int) (result []*model.SeatDept, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatDeptDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatDeptDo) Debug() ISeatDeptDo {
	return s.withDO(s.DO.Debug())
}

func (s seatDeptDo) WithContext(ctx context.Context) ISeatDeptDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatDeptDo) ReadDB() ISeatDeptDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatDeptDo) WriteDB() ISeatDeptDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatDeptDo) Session(config *gorm.Session) ISeatDeptDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatDeptDo) Clauses(conds ...clause.Expression) ISeatDeptDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatDeptDo) Returning(value interface{}, columns ...string) ISeatDeptDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatDeptDo) Not(conds ...gen.Condition) ISeatDeptDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatDeptDo) Or(conds ...gen.Condition) ISeatDeptDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatDeptDo) Select(conds ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatDeptDo) Where(conds ...gen.Condition) ISeatDeptDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatDeptDo) Order(conds ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatDeptDo) Distinct(cols ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatDeptDo) Omit(cols ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatDeptDo) Join(table schema.Tabler, on ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatDeptDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatDeptDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatDeptDo) Group(cols ...field.Expr) ISeatDeptDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatDeptDo) Having(conds ...gen.Condition) ISeatDeptDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatDeptDo) Limit(limit int) ISeatDeptDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatDeptDo) Offset(offset int) ISeatDeptDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatDeptDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatDeptDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatDeptDo) Unscoped() ISeatDeptDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatDeptDo) Create(values ...*model.SeatDept) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatDeptDo) CreateInBatches(values []*model.SeatDept, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatDeptDo) Save(values ...*model.SeatDept) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatDeptDo) First() (*model.SeatDept, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatDept), nil
	}
}

func (s seatDeptDo) Take() (*model.SeatDept, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatDept), nil
	}
}

func (s seatDeptDo) Last() (*model.SeatDept, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatDept), nil
	}
}

func (s seatDeptDo) Find() ([]*model.SeatDept, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatDept), err
}

func (s seatDeptDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatDept, err error) {
	buf := make([]*model.SeatDept, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatDeptDo) FindInBatches(result *[]*model.SeatDept, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatDeptDo) Attrs(attrs ...field.AssignExpr) ISeatDeptDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatDeptDo) Assign(attrs ...field.AssignExpr) ISeatDeptDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatDeptDo) Joins(fields ...field.RelationField) ISeatDeptDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatDeptDo) Preload(fields ...field.RelationField) ISeatDeptDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatDeptDo) FirstOrInit() (*model.SeatDept, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatDept), nil
	}
}

func (s seatDeptDo) FirstOrCreate() (*model.SeatDept, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatDept), nil
	}
}

func (s seatDeptDo) FindByPage(offset int, limit int) (result []*model.SeatDept, count int64, err error) {
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

func (s seatDeptDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatDeptDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatDeptDo) Delete(models ...*model.SeatDept) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatDeptDo) withDO(do gen.Dao) *seatDeptDo {
	s.DO = *do.(*gen.DO)
	return s
}
