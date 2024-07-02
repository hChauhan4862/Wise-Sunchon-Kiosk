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

func newWiseDroomReserve(db *gorm.DB, opts ...gen.DOOption) wiseDroomReserve {
	_wiseDroomReserve := wiseDroomReserve{}

	_wiseDroomReserve.wiseDroomReserveDo.UseDB(db, opts...)
	_wiseDroomReserve.wiseDroomReserveDo.UseModel(&model.WiseDroomReserve{})

	tableName := _wiseDroomReserve.wiseDroomReserveDo.TableName()
	_wiseDroomReserve.ALL = field.NewAsterisk(tableName)
	_wiseDroomReserve.ReserveNo = field.NewString(tableName, "reserve_no")
	_wiseDroomReserve.DroomNo = field.NewInt64(tableName, "droom_no")
	_wiseDroomReserve.DroomName = field.NewString(tableName, "droom_name")
	_wiseDroomReserve.InputTime = field.NewString(tableName, "input_time")
	_wiseDroomReserve.ReserveDate = field.NewString(tableName, "reserve_date")
	_wiseDroomReserve.StartTime = field.NewString(tableName, "start_time")
	_wiseDroomReserve.EndTime = field.NewString(tableName, "end_time")
	_wiseDroomReserve.UserID = field.NewString(tableName, "user_id")
	_wiseDroomReserve.UserName = field.NewString(tableName, "user_name")
	_wiseDroomReserve.UserComment = field.NewString(tableName, "user_comment")
	_wiseDroomReserve.UseMan = field.NewInt64(tableName, "use_man")
	_wiseDroomReserve.CouserIds = field.NewString(tableName, "couser_ids")
	_wiseDroomReserve.ReserveStat = field.NewInt64(tableName, "reserve_stat")
	_wiseDroomReserve.CanBigo = field.NewString(tableName, "can_bigo")
	_wiseDroomReserve.Bigo = field.NewString(tableName, "bigo")

	_wiseDroomReserve.fillFieldMap()

	return _wiseDroomReserve
}

type wiseDroomReserve struct {
	wiseDroomReserveDo

	ALL         field.Asterisk
	ReserveNo   field.String
	DroomNo     field.Int64
	DroomName   field.String
	InputTime   field.String
	ReserveDate field.String
	StartTime   field.String
	EndTime     field.String
	UserID      field.String
	UserName    field.String
	UserComment field.String
	UseMan      field.Int64
	CouserIds   field.String
	ReserveStat field.Int64
	CanBigo     field.String
	Bigo        field.String

	fieldMap map[string]field.Expr
}

func (w wiseDroomReserve) Table(newTableName string) *wiseDroomReserve {
	w.wiseDroomReserveDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wiseDroomReserve) As(alias string) *wiseDroomReserve {
	w.wiseDroomReserveDo.DO = *(w.wiseDroomReserveDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wiseDroomReserve) updateTableName(table string) *wiseDroomReserve {
	w.ALL = field.NewAsterisk(table)
	w.ReserveNo = field.NewString(table, "reserve_no")
	w.DroomNo = field.NewInt64(table, "droom_no")
	w.DroomName = field.NewString(table, "droom_name")
	w.InputTime = field.NewString(table, "input_time")
	w.ReserveDate = field.NewString(table, "reserve_date")
	w.StartTime = field.NewString(table, "start_time")
	w.EndTime = field.NewString(table, "end_time")
	w.UserID = field.NewString(table, "user_id")
	w.UserName = field.NewString(table, "user_name")
	w.UserComment = field.NewString(table, "user_comment")
	w.UseMan = field.NewInt64(table, "use_man")
	w.CouserIds = field.NewString(table, "couser_ids")
	w.ReserveStat = field.NewInt64(table, "reserve_stat")
	w.CanBigo = field.NewString(table, "can_bigo")
	w.Bigo = field.NewString(table, "bigo")

	w.fillFieldMap()

	return w
}

func (w *wiseDroomReserve) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wiseDroomReserve) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 15)
	w.fieldMap["reserve_no"] = w.ReserveNo
	w.fieldMap["droom_no"] = w.DroomNo
	w.fieldMap["droom_name"] = w.DroomName
	w.fieldMap["input_time"] = w.InputTime
	w.fieldMap["reserve_date"] = w.ReserveDate
	w.fieldMap["start_time"] = w.StartTime
	w.fieldMap["end_time"] = w.EndTime
	w.fieldMap["user_id"] = w.UserID
	w.fieldMap["user_name"] = w.UserName
	w.fieldMap["user_comment"] = w.UserComment
	w.fieldMap["use_man"] = w.UseMan
	w.fieldMap["couser_ids"] = w.CouserIds
	w.fieldMap["reserve_stat"] = w.ReserveStat
	w.fieldMap["can_bigo"] = w.CanBigo
	w.fieldMap["bigo"] = w.Bigo
}

func (w wiseDroomReserve) clone(db *gorm.DB) wiseDroomReserve {
	w.wiseDroomReserveDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wiseDroomReserve) replaceDB(db *gorm.DB) wiseDroomReserve {
	w.wiseDroomReserveDo.ReplaceDB(db)
	return w
}

type wiseDroomReserveDo struct{ gen.DO }

type IWiseDroomReserveDo interface {
	gen.SubQuery
	Debug() IWiseDroomReserveDo
	WithContext(ctx context.Context) IWiseDroomReserveDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWiseDroomReserveDo
	WriteDB() IWiseDroomReserveDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWiseDroomReserveDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWiseDroomReserveDo
	Not(conds ...gen.Condition) IWiseDroomReserveDo
	Or(conds ...gen.Condition) IWiseDroomReserveDo
	Select(conds ...field.Expr) IWiseDroomReserveDo
	Where(conds ...gen.Condition) IWiseDroomReserveDo
	Order(conds ...field.Expr) IWiseDroomReserveDo
	Distinct(cols ...field.Expr) IWiseDroomReserveDo
	Omit(cols ...field.Expr) IWiseDroomReserveDo
	Join(table schema.Tabler, on ...field.Expr) IWiseDroomReserveDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWiseDroomReserveDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWiseDroomReserveDo
	Group(cols ...field.Expr) IWiseDroomReserveDo
	Having(conds ...gen.Condition) IWiseDroomReserveDo
	Limit(limit int) IWiseDroomReserveDo
	Offset(offset int) IWiseDroomReserveDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseDroomReserveDo
	Unscoped() IWiseDroomReserveDo
	Create(values ...*model.WiseDroomReserve) error
	CreateInBatches(values []*model.WiseDroomReserve, batchSize int) error
	Save(values ...*model.WiseDroomReserve) error
	First() (*model.WiseDroomReserve, error)
	Take() (*model.WiseDroomReserve, error)
	Last() (*model.WiseDroomReserve, error)
	Find() ([]*model.WiseDroomReserve, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseDroomReserve, err error)
	FindInBatches(result *[]*model.WiseDroomReserve, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WiseDroomReserve) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWiseDroomReserveDo
	Assign(attrs ...field.AssignExpr) IWiseDroomReserveDo
	Joins(fields ...field.RelationField) IWiseDroomReserveDo
	Preload(fields ...field.RelationField) IWiseDroomReserveDo
	FirstOrInit() (*model.WiseDroomReserve, error)
	FirstOrCreate() (*model.WiseDroomReserve, error)
	FindByPage(offset int, limit int) (result []*model.WiseDroomReserve, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWiseDroomReserveDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wiseDroomReserveDo) Debug() IWiseDroomReserveDo {
	return w.withDO(w.DO.Debug())
}

func (w wiseDroomReserveDo) WithContext(ctx context.Context) IWiseDroomReserveDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wiseDroomReserveDo) ReadDB() IWiseDroomReserveDo {
	return w.Clauses(dbresolver.Read)
}

func (w wiseDroomReserveDo) WriteDB() IWiseDroomReserveDo {
	return w.Clauses(dbresolver.Write)
}

func (w wiseDroomReserveDo) Session(config *gorm.Session) IWiseDroomReserveDo {
	return w.withDO(w.DO.Session(config))
}

func (w wiseDroomReserveDo) Clauses(conds ...clause.Expression) IWiseDroomReserveDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wiseDroomReserveDo) Returning(value interface{}, columns ...string) IWiseDroomReserveDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wiseDroomReserveDo) Not(conds ...gen.Condition) IWiseDroomReserveDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wiseDroomReserveDo) Or(conds ...gen.Condition) IWiseDroomReserveDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wiseDroomReserveDo) Select(conds ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wiseDroomReserveDo) Where(conds ...gen.Condition) IWiseDroomReserveDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wiseDroomReserveDo) Order(conds ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wiseDroomReserveDo) Distinct(cols ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wiseDroomReserveDo) Omit(cols ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wiseDroomReserveDo) Join(table schema.Tabler, on ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wiseDroomReserveDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wiseDroomReserveDo) RightJoin(table schema.Tabler, on ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wiseDroomReserveDo) Group(cols ...field.Expr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wiseDroomReserveDo) Having(conds ...gen.Condition) IWiseDroomReserveDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wiseDroomReserveDo) Limit(limit int) IWiseDroomReserveDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wiseDroomReserveDo) Offset(offset int) IWiseDroomReserveDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wiseDroomReserveDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseDroomReserveDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wiseDroomReserveDo) Unscoped() IWiseDroomReserveDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wiseDroomReserveDo) Create(values ...*model.WiseDroomReserve) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wiseDroomReserveDo) CreateInBatches(values []*model.WiseDroomReserve, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wiseDroomReserveDo) Save(values ...*model.WiseDroomReserve) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wiseDroomReserveDo) First() (*model.WiseDroomReserve, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroomReserve), nil
	}
}

func (w wiseDroomReserveDo) Take() (*model.WiseDroomReserve, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroomReserve), nil
	}
}

func (w wiseDroomReserveDo) Last() (*model.WiseDroomReserve, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroomReserve), nil
	}
}

func (w wiseDroomReserveDo) Find() ([]*model.WiseDroomReserve, error) {
	result, err := w.DO.Find()
	return result.([]*model.WiseDroomReserve), err
}

func (w wiseDroomReserveDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseDroomReserve, err error) {
	buf := make([]*model.WiseDroomReserve, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wiseDroomReserveDo) FindInBatches(result *[]*model.WiseDroomReserve, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wiseDroomReserveDo) Attrs(attrs ...field.AssignExpr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wiseDroomReserveDo) Assign(attrs ...field.AssignExpr) IWiseDroomReserveDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wiseDroomReserveDo) Joins(fields ...field.RelationField) IWiseDroomReserveDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wiseDroomReserveDo) Preload(fields ...field.RelationField) IWiseDroomReserveDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wiseDroomReserveDo) FirstOrInit() (*model.WiseDroomReserve, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroomReserve), nil
	}
}

func (w wiseDroomReserveDo) FirstOrCreate() (*model.WiseDroomReserve, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseDroomReserve), nil
	}
}

func (w wiseDroomReserveDo) FindByPage(offset int, limit int) (result []*model.WiseDroomReserve, count int64, err error) {
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

func (w wiseDroomReserveDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wiseDroomReserveDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wiseDroomReserveDo) Delete(models ...*model.WiseDroomReserve) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wiseDroomReserveDo) withDO(do gen.Dao) *wiseDroomReserveDo {
	w.DO = *do.(*gen.DO)
	return w
}