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

func newSeatSectorView(db *gorm.DB, opts ...gen.DOOption) seatSectorView {
	_seatSectorView := seatSectorView{}

	_seatSectorView.seatSectorViewDo.UseDB(db, opts...)
	_seatSectorView.seatSectorViewDo.UseModel(&model.SeatSectorView{})

	tableName := _seatSectorView.seatSectorViewDo.TableName()
	_seatSectorView.ALL = field.NewAsterisk(tableName)
	_seatSectorView.VIEWSECTORNO = field.NewInt64(tableName, "VIEWSECTORNO")
	_seatSectorView.NAME = field.NewString(tableName, "NAME")
	_seatSectorView.ENNAME = field.NewString(tableName, "EN_NAME")
	_seatSectorView.ROOMNO = field.NewInt64(tableName, "ROOMNO")
	_seatSectorView.IMAGE = field.NewString(tableName, "IMAGE")
	_seatSectorView.LINKPAGE = field.NewString(tableName, "LINKPAGE")
	_seatSectorView.MAPPOINT = field.NewString(tableName, "MAPPOINT")
	_seatSectorView.DEFSECTORNO = field.NewInt64(tableName, "DEFSECTORNO")
	_seatSectorView.MAPLABEL = field.NewString(tableName, "MAPLABEL")
	_seatSectorView.ENMAPLABEL = field.NewString(tableName, "EN_MAPLABEL")
	_seatSectorView.MAPPOINT2 = field.NewString(tableName, "MAPPOINT2")
	_seatSectorView.MAPLABEL2 = field.NewString(tableName, "MAPLABEL2")
	_seatSectorView.ENMAPLABEL2 = field.NewString(tableName, "EN_MAPLABEL2")

	_seatSectorView.fillFieldMap()

	return _seatSectorView
}

type seatSectorView struct {
	seatSectorViewDo

	ALL          field.Asterisk
	VIEWSECTORNO field.Int64
	NAME         field.String
	ENNAME       field.String
	ROOMNO       field.Int64
	IMAGE        field.String
	LINKPAGE     field.String
	MAPPOINT     field.String
	DEFSECTORNO  field.Int64
	MAPLABEL     field.String
	ENMAPLABEL   field.String
	MAPPOINT2    field.String
	MAPLABEL2    field.String
	ENMAPLABEL2  field.String

	fieldMap map[string]field.Expr
}

func (s seatSectorView) Table(newTableName string) *seatSectorView {
	s.seatSectorViewDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seatSectorView) As(alias string) *seatSectorView {
	s.seatSectorViewDo.DO = *(s.seatSectorViewDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seatSectorView) updateTableName(table string) *seatSectorView {
	s.ALL = field.NewAsterisk(table)
	s.VIEWSECTORNO = field.NewInt64(table, "VIEWSECTORNO")
	s.NAME = field.NewString(table, "NAME")
	s.ENNAME = field.NewString(table, "EN_NAME")
	s.ROOMNO = field.NewInt64(table, "ROOMNO")
	s.IMAGE = field.NewString(table, "IMAGE")
	s.LINKPAGE = field.NewString(table, "LINKPAGE")
	s.MAPPOINT = field.NewString(table, "MAPPOINT")
	s.DEFSECTORNO = field.NewInt64(table, "DEFSECTORNO")
	s.MAPLABEL = field.NewString(table, "MAPLABEL")
	s.ENMAPLABEL = field.NewString(table, "EN_MAPLABEL")
	s.MAPPOINT2 = field.NewString(table, "MAPPOINT2")
	s.MAPLABEL2 = field.NewString(table, "MAPLABEL2")
	s.ENMAPLABEL2 = field.NewString(table, "EN_MAPLABEL2")

	s.fillFieldMap()

	return s
}

func (s *seatSectorView) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seatSectorView) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["VIEWSECTORNO"] = s.VIEWSECTORNO
	s.fieldMap["NAME"] = s.NAME
	s.fieldMap["EN_NAME"] = s.ENNAME
	s.fieldMap["ROOMNO"] = s.ROOMNO
	s.fieldMap["IMAGE"] = s.IMAGE
	s.fieldMap["LINKPAGE"] = s.LINKPAGE
	s.fieldMap["MAPPOINT"] = s.MAPPOINT
	s.fieldMap["DEFSECTORNO"] = s.DEFSECTORNO
	s.fieldMap["MAPLABEL"] = s.MAPLABEL
	s.fieldMap["EN_MAPLABEL"] = s.ENMAPLABEL
	s.fieldMap["MAPPOINT2"] = s.MAPPOINT2
	s.fieldMap["MAPLABEL2"] = s.MAPLABEL2
	s.fieldMap["EN_MAPLABEL2"] = s.ENMAPLABEL2
}

func (s seatSectorView) clone(db *gorm.DB) seatSectorView {
	s.seatSectorViewDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seatSectorView) replaceDB(db *gorm.DB) seatSectorView {
	s.seatSectorViewDo.ReplaceDB(db)
	return s
}

type seatSectorViewDo struct{ gen.DO }

type ISeatSectorViewDo interface {
	gen.SubQuery
	Debug() ISeatSectorViewDo
	WithContext(ctx context.Context) ISeatSectorViewDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeatSectorViewDo
	WriteDB() ISeatSectorViewDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeatSectorViewDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeatSectorViewDo
	Not(conds ...gen.Condition) ISeatSectorViewDo
	Or(conds ...gen.Condition) ISeatSectorViewDo
	Select(conds ...field.Expr) ISeatSectorViewDo
	Where(conds ...gen.Condition) ISeatSectorViewDo
	Order(conds ...field.Expr) ISeatSectorViewDo
	Distinct(cols ...field.Expr) ISeatSectorViewDo
	Omit(cols ...field.Expr) ISeatSectorViewDo
	Join(table schema.Tabler, on ...field.Expr) ISeatSectorViewDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeatSectorViewDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeatSectorViewDo
	Group(cols ...field.Expr) ISeatSectorViewDo
	Having(conds ...gen.Condition) ISeatSectorViewDo
	Limit(limit int) ISeatSectorViewDo
	Offset(offset int) ISeatSectorViewDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatSectorViewDo
	Unscoped() ISeatSectorViewDo
	Create(values ...*model.SeatSectorView) error
	CreateInBatches(values []*model.SeatSectorView, batchSize int) error
	Save(values ...*model.SeatSectorView) error
	First() (*model.SeatSectorView, error)
	Take() (*model.SeatSectorView, error)
	Last() (*model.SeatSectorView, error)
	Find() ([]*model.SeatSectorView, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatSectorView, err error)
	FindInBatches(result *[]*model.SeatSectorView, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeatSectorView) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeatSectorViewDo
	Assign(attrs ...field.AssignExpr) ISeatSectorViewDo
	Joins(fields ...field.RelationField) ISeatSectorViewDo
	Preload(fields ...field.RelationField) ISeatSectorViewDo
	FirstOrInit() (*model.SeatSectorView, error)
	FirstOrCreate() (*model.SeatSectorView, error)
	FindByPage(offset int, limit int) (result []*model.SeatSectorView, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeatSectorViewDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seatSectorViewDo) Debug() ISeatSectorViewDo {
	return s.withDO(s.DO.Debug())
}

func (s seatSectorViewDo) WithContext(ctx context.Context) ISeatSectorViewDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seatSectorViewDo) ReadDB() ISeatSectorViewDo {
	return s.Clauses(dbresolver.Read)
}

func (s seatSectorViewDo) WriteDB() ISeatSectorViewDo {
	return s.Clauses(dbresolver.Write)
}

func (s seatSectorViewDo) Session(config *gorm.Session) ISeatSectorViewDo {
	return s.withDO(s.DO.Session(config))
}

func (s seatSectorViewDo) Clauses(conds ...clause.Expression) ISeatSectorViewDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seatSectorViewDo) Returning(value interface{}, columns ...string) ISeatSectorViewDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seatSectorViewDo) Not(conds ...gen.Condition) ISeatSectorViewDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seatSectorViewDo) Or(conds ...gen.Condition) ISeatSectorViewDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seatSectorViewDo) Select(conds ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seatSectorViewDo) Where(conds ...gen.Condition) ISeatSectorViewDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seatSectorViewDo) Order(conds ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seatSectorViewDo) Distinct(cols ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seatSectorViewDo) Omit(cols ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seatSectorViewDo) Join(table schema.Tabler, on ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seatSectorViewDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seatSectorViewDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seatSectorViewDo) Group(cols ...field.Expr) ISeatSectorViewDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seatSectorViewDo) Having(conds ...gen.Condition) ISeatSectorViewDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seatSectorViewDo) Limit(limit int) ISeatSectorViewDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seatSectorViewDo) Offset(offset int) ISeatSectorViewDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seatSectorViewDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeatSectorViewDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seatSectorViewDo) Unscoped() ISeatSectorViewDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seatSectorViewDo) Create(values ...*model.SeatSectorView) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seatSectorViewDo) CreateInBatches(values []*model.SeatSectorView, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seatSectorViewDo) Save(values ...*model.SeatSectorView) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seatSectorViewDo) First() (*model.SeatSectorView, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSectorView), nil
	}
}

func (s seatSectorViewDo) Take() (*model.SeatSectorView, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSectorView), nil
	}
}

func (s seatSectorViewDo) Last() (*model.SeatSectorView, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSectorView), nil
	}
}

func (s seatSectorViewDo) Find() ([]*model.SeatSectorView, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeatSectorView), err
}

func (s seatSectorViewDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeatSectorView, err error) {
	buf := make([]*model.SeatSectorView, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seatSectorViewDo) FindInBatches(result *[]*model.SeatSectorView, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seatSectorViewDo) Attrs(attrs ...field.AssignExpr) ISeatSectorViewDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seatSectorViewDo) Assign(attrs ...field.AssignExpr) ISeatSectorViewDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seatSectorViewDo) Joins(fields ...field.RelationField) ISeatSectorViewDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seatSectorViewDo) Preload(fields ...field.RelationField) ISeatSectorViewDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seatSectorViewDo) FirstOrInit() (*model.SeatSectorView, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSectorView), nil
	}
}

func (s seatSectorViewDo) FirstOrCreate() (*model.SeatSectorView, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeatSectorView), nil
	}
}

func (s seatSectorViewDo) FindByPage(offset int, limit int) (result []*model.SeatSectorView, count int64, err error) {
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

func (s seatSectorViewDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seatSectorViewDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seatSectorViewDo) Delete(models ...*model.SeatSectorView) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seatSectorViewDo) withDO(do gen.Dao) *seatSectorViewDo {
	s.DO = *do.(*gen.DO)
	return s
}
