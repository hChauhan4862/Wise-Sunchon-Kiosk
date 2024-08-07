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

func newWiseDroom(db *gorm.DB, opts ...gen.DOOption) wiseDroom {
	_wiseDroom := wiseDroom{}

	_wiseDroom.wiseDroomDo.UseDB(db, opts...)
	_wiseDroom.wiseDroomDo.UseModel(&model.WiseDroom{})

	tableName := _wiseDroom.wiseDroomDo.TableName()
	_wiseDroom.ALL = field.NewAsterisk(tableName)
	_wiseDroom.DroomNo = field.NewInt64(tableName, "droom_no")
	_wiseDroom.DroomName = field.NewString(tableName, "droom_name")
	_wiseDroom.DroomGubun = field.NewString(tableName, "droom_gubun")
	_wiseDroom.SeatCnt = field.NewInt64(tableName, "seat_cnt")
	_wiseDroom.UseYn = field.NewString(tableName, "use_yn")
	_wiseDroom.UseMin = field.NewInt64(tableName, "use_min")
	_wiseDroom.MinCnt = field.NewInt64(tableName, "min_cnt")
	_wiseDroom.ReserveDay = field.NewInt64(tableName, "reserve_day")
	_wiseDroom.ReserveMin = field.NewInt64(tableName, "reserve_min")
	_wiseDroom.CancelMin = field.NewInt64(tableName, "cancel_min")
	_wiseDroom.BlockYn = field.NewString(tableName, "block_yn")
	_wiseDroom.BlockMin = field.NewInt64(tableName, "block_min")
	_wiseDroom.ContYn = field.NewString(tableName, "cont_yn")
	_wiseDroom.ContMin = field.NewInt64(tableName, "cont_min")
	_wiseDroom.EtcUseYn = field.NewString(tableName, "etc_use_yn")
	_wiseDroom.EtcName = field.NewString(tableName, "etc_name")
	_wiseDroom.Bigo = field.NewString(tableName, "bigo")

	_wiseDroom.fillFieldMap()

	return _wiseDroom
}

type wiseDroom struct {
	wiseDroomDo

	ALL        field.Asterisk
	DroomNo    field.Int64
	DroomName  field.String
	DroomGubun field.String
	SeatCnt    field.Int64
	UseYn      field.String
	UseMin     field.Int64
	MinCnt     field.Int64
	ReserveDay field.Int64
	ReserveMin field.Int64
	CancelMin  field.Int64
	BlockYn    field.String
	BlockMin   field.Int64
	ContYn     field.String
	ContMin    field.Int64
	EtcUseYn   field.String
	EtcName    field.String
	Bigo       field.String

	fieldMap map[string]field.Expr
}

func (w wiseDroom) Table(newTableName string) *wiseDroom {
	w.wiseDroomDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wiseDroom) As(alias string) *wiseDroom {
	w.wiseDroomDo.DO = *(w.wiseDroomDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wiseDroom) updateTableName(table string) *wiseDroom {
	w.ALL = field.NewAsterisk(table)
	w.DroomNo = field.NewInt64(table, "droom_no")
	w.DroomName = field.NewString(table, "droom_name")
	w.DroomGubun = field.NewString(table, "droom_gubun")
	w.SeatCnt = field.NewInt64(table, "seat_cnt")
	w.UseYn = field.NewString(table, "use_yn")
	w.UseMin = field.NewInt64(table, "use_min")
	w.MinCnt = field.NewInt64(table, "min_cnt")
	w.ReserveDay = field.NewInt64(table, "reserve_day")
	w.ReserveMin = field.NewInt64(table, "reserve_min")
	w.CancelMin = field.NewInt64(table, "cancel_min")
	w.BlockYn = field.NewString(table, "block_yn")
	w.BlockMin = field.NewInt64(table, "block_min")
	w.ContYn = field.NewString(table, "cont_yn")
	w.ContMin = field.NewInt64(table, "cont_min")
	w.EtcUseYn = field.NewString(table, "etc_use_yn")
	w.EtcName = field.NewString(table, "etc_name")
	w.Bigo = field.NewString(table, "bigo")

	w.fillFieldMap()

	return w
}

func (w *wiseDroom) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wiseDroom) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 17)
	w.fieldMap["droom_no"] = w.DroomNo
	w.fieldMap["droom_name"] = w.DroomName
	w.fieldMap["droom_gubun"] = w.DroomGubun
	w.fieldMap["seat_cnt"] = w.SeatCnt
	w.fieldMap["use_yn"] = w.UseYn
	w.fieldMap["use_min"] = w.UseMin
	w.fieldMap["min_cnt"] = w.MinCnt
	w.fieldMap["reserve_day"] = w.ReserveDay
	w.fieldMap["reserve_min"] = w.ReserveMin
	w.fieldMap["cancel_min"] = w.CancelMin
	w.fieldMap["block_yn"] = w.BlockYn
	w.fieldMap["block_min"] = w.BlockMin
	w.fieldMap["cont_yn"] = w.ContYn
	w.fieldMap["cont_min"] = w.ContMin
	w.fieldMap["etc_use_yn"] = w.EtcUseYn
	w.fieldMap["etc_name"] = w.EtcName
	w.fieldMap["bigo"] = w.Bigo
}

func (w wiseDroom) clone(db *gorm.DB) wiseDroom {
	w.wiseDroomDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wiseDroom) replaceDB(db *gorm.DB) wiseDroom {
	w.wiseDroomDo.ReplaceDB(db)
	return w
}

type wiseDroomDo struct{ gen.DO }

type IWiseDroomDo interface {
	gen.SubQuery
	Debug() IWiseDroomDo
	WithContext(ctx context.Context) IWiseDroomDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWiseDroomDo
	WriteDB() IWiseDroomDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWiseDroomDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWiseDroomDo
	Not(conds ...gen.Condition) IWiseDroomDo
	Or(conds ...gen.Condition) IWiseDroomDo
	Select(conds ...field.Expr) IWiseDroomDo
	Where(conds ...gen.Condition) IWiseDroomDo
	Order(conds ...field.Expr) IWiseDroomDo
	Distinct(cols ...field.Expr) IWiseDroomDo
	Omit(cols ...field.Expr) IWiseDroomDo
	Join(table schema.Tabler, on ...field.Expr) IWiseDroomDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWiseDroomDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWiseDroomDo
	Group(cols ...field.Expr) IWiseDroomDo
	Having(conds ...gen.Condition) IWiseDroomDo
	Limit(limit int) IWiseDroomDo
	Offset(offset int) IWiseDroomDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseDroomDo
	Unscoped() IWiseDroomDo
	Create(values ...*model.WiseDroom) error
	CreateInBatches(values []*model.WiseDroom, batchSize int) error
	Save(values ...*model.WiseDroom) error
	First() (*model.WiseDroom, error)
	Take() (*model.WiseDroom, error)
	Last() (*model.WiseDroom, error)
	Find() ([]*model.WiseDroom, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseDroom, err error)
	FindInBatches(result *[]*model.WiseDroom, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WiseDroom) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWiseDroomDo
	Assign(attrs ...field.AssignExpr) IWiseDroomDo
	Joins(fields ...field.RelationField) IWiseDroomDo
	Preload(fields ...field.RelationField) IWiseDroomDo
	FirstOrInit() (*model.WiseDroom, error)
	FirstOrCreate() (*model.WiseDroom, error)
	FindByPage(offset int, limit int) (result []*model.WiseDroom, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWiseDroomDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wiseDroomDo) Debug() IWiseDroomDo {
	return w.withDO(w.DO.Debug())
}

func (w wiseDroomDo) WithContext(ctx context.Context) IWiseDroomDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wiseDroomDo) ReadDB() IWiseDroomDo {
	return w.Clauses(dbresolver.Read)
}

func (w wiseDroomDo) WriteDB() IWiseDroomDo {
	return w.Clauses(dbresolver.Write)
}

func (w wiseDroomDo) Session(config *gorm.Session) IWiseDroomDo {
	return w.withDO(w.DO.Session(config))
}

func (w wiseDroomDo) Clauses(conds ...clause.Expression) IWiseDroomDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wiseDroomDo) Returning(value interface{}, columns ...string) IWiseDroomDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wiseDroomDo) Not(conds ...gen.Condition) IWiseDroomDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wiseDroomDo) Or(conds ...gen.Condition) IWiseDroomDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wiseDroomDo) Select(conds ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wiseDroomDo) Where(conds ...gen.Condition) IWiseDroomDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wiseDroomDo) Order(conds ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wiseDroomDo) Distinct(cols ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wiseDroomDo) Omit(cols ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wiseDroomDo) Join(table schema.Tabler, on ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wiseDroomDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wiseDroomDo) RightJoin(table schema.Tabler, on ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wiseDroomDo) Group(cols ...field.Expr) IWiseDroomDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wiseDroomDo) Having(conds ...gen.Condition) IWiseDroomDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wiseDroomDo) Limit(limit int) IWiseDroomDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wiseDroomDo) Offset(offset int) IWiseDroomDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wiseDroomDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseDroomDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wiseDroomDo) Unscoped() IWiseDroomDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wiseDroomDo) Create(values ...*model.WiseDroom) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wiseDroomDo) CreateInBatches(values []*model.WiseDroom, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wiseDroomDo) Save(values ...*model.WiseDroom) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wiseDroomDo) First() (*model.WiseDroom, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroom), nil
	}
}

func (w wiseDroomDo) Take() (*model.WiseDroom, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroom), nil
	}
}

func (w wiseDroomDo) Last() (*model.WiseDroom, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroom), nil
	}
}

func (w wiseDroomDo) Find() ([]*model.WiseDroom, error) {
	result, err := w.DO.Find()
	return result.([]*model.WiseDroom), err
}

func (w wiseDroomDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseDroom, err error) {
	buf := make([]*model.WiseDroom, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wiseDroomDo) FindInBatches(result *[]*model.WiseDroom, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wiseDroomDo) Attrs(attrs ...field.AssignExpr) IWiseDroomDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wiseDroomDo) Assign(attrs ...field.AssignExpr) IWiseDroomDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wiseDroomDo) Joins(fields ...field.RelationField) IWiseDroomDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wiseDroomDo) Preload(fields ...field.RelationField) IWiseDroomDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wiseDroomDo) FirstOrInit() (*model.WiseDroom, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroom), nil
	}
}

func (w wiseDroomDo) FirstOrCreate() (*model.WiseDroom, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroom), nil
	}
}

func (w wiseDroomDo) FindByPage(offset int, limit int) (result []*model.WiseDroom, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wiseDroomDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wiseDroomDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wiseDroomDo) Delete(models ...*model.WiseDroom) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wiseDroomDo) withDO(do gen.Dao) *wiseDroomDo {
	w.DO = *do.(*gen.DO)
	return w
}
