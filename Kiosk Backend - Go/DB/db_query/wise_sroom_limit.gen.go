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

func newWiseSroomLimit(db *gorm.DB, opts ...gen.DOOption) wiseSroomLimit {
	_wiseSroomLimit := wiseSroomLimit{}

	_wiseSroomLimit.wiseSroomLimitDo.UseDB(db, opts...)
	_wiseSroomLimit.wiseSroomLimitDo.UseModel(&model.WiseSroomLimit{})

	tableName := _wiseSroomLimit.wiseSroomLimitDo.TableName()
	_wiseSroomLimit.ALL = field.NewAsterisk(tableName)
	_wiseSroomLimit.SroomNo = field.NewInt64(tableName, "sroom_no")
	_wiseSroomLimit.SroomName = field.NewString(tableName, "sroom_name")
	_wiseSroomLimit.StartDate = field.NewString(tableName, "start_date")
	_wiseSroomLimit.EndDate = field.NewString(tableName, "end_date")
	_wiseSroomLimit.ComStime = field.NewString(tableName, "com_stime")
	_wiseSroomLimit.ComEtime = field.NewString(tableName, "com_etime")
	_wiseSroomLimit.SatStime = field.NewString(tableName, "sat_stime")
	_wiseSroomLimit.SatEtime = field.NewString(tableName, "sat_etime")
	_wiseSroomLimit.SatUseYn = field.NewString(tableName, "sat_use_yn")
	_wiseSroomLimit.ReserveDayCnt = field.NewInt64(tableName, "reserve_day_cnt")
	_wiseSroomLimit.InputTime = field.NewString(tableName, "input_time")
	_wiseSroomLimit.Bigo = field.NewString(tableName, "bigo")
	_wiseSroomLimit.SunTime = field.NewString(tableName, "sun_time")
	_wiseSroomLimit.SunEtime = field.NewString(tableName, "sun_etime")
	_wiseSroomLimit.SunUseYn = field.NewString(tableName, "sun_use_yn")
	_wiseSroomLimit.VacationStime = field.NewString(tableName, "vacation_stime")
	_wiseSroomLimit.VacationEtime = field.NewString(tableName, "vacation_etime")
	_wiseSroomLimit.PatLimit = field.NewString(tableName, "pat_limit")

	_wiseSroomLimit.fillFieldMap()

	return _wiseSroomLimit
}

type wiseSroomLimit struct {
	wiseSroomLimitDo

	ALL           field.Asterisk
	SroomNo       field.Int64
	SroomName     field.String
	StartDate     field.String
	EndDate       field.String
	ComStime      field.String
	ComEtime      field.String
	SatStime      field.String
	SatEtime      field.String
	SatUseYn      field.String
	ReserveDayCnt field.Int64
	InputTime     field.String
	Bigo          field.String
	SunTime       field.String
	SunEtime      field.String
	SunUseYn      field.String
	VacationStime field.String
	VacationEtime field.String
	PatLimit      field.String

	fieldMap map[string]field.Expr
}

func (w wiseSroomLimit) Table(newTableName string) *wiseSroomLimit {
	w.wiseSroomLimitDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wiseSroomLimit) As(alias string) *wiseSroomLimit {
	w.wiseSroomLimitDo.DO = *(w.wiseSroomLimitDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wiseSroomLimit) updateTableName(table string) *wiseSroomLimit {
	w.ALL = field.NewAsterisk(table)
	w.SroomNo = field.NewInt64(table, "sroom_no")
	w.SroomName = field.NewString(table, "sroom_name")
	w.StartDate = field.NewString(table, "start_date")
	w.EndDate = field.NewString(table, "end_date")
	w.ComStime = field.NewString(table, "com_stime")
	w.ComEtime = field.NewString(table, "com_etime")
	w.SatStime = field.NewString(table, "sat_stime")
	w.SatEtime = field.NewString(table, "sat_etime")
	w.SatUseYn = field.NewString(table, "sat_use_yn")
	w.ReserveDayCnt = field.NewInt64(table, "reserve_day_cnt")
	w.InputTime = field.NewString(table, "input_time")
	w.Bigo = field.NewString(table, "bigo")
	w.SunTime = field.NewString(table, "sun_time")
	w.SunEtime = field.NewString(table, "sun_etime")
	w.SunUseYn = field.NewString(table, "sun_use_yn")
	w.VacationStime = field.NewString(table, "vacation_stime")
	w.VacationEtime = field.NewString(table, "vacation_etime")
	w.PatLimit = field.NewString(table, "pat_limit")

	w.fillFieldMap()

	return w
}

func (w *wiseSroomLimit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wiseSroomLimit) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 18)
	w.fieldMap["sroom_no"] = w.SroomNo
	w.fieldMap["sroom_name"] = w.SroomName
	w.fieldMap["start_date"] = w.StartDate
	w.fieldMap["end_date"] = w.EndDate
	w.fieldMap["com_stime"] = w.ComStime
	w.fieldMap["com_etime"] = w.ComEtime
	w.fieldMap["sat_stime"] = w.SatStime
	w.fieldMap["sat_etime"] = w.SatEtime
	w.fieldMap["sat_use_yn"] = w.SatUseYn
	w.fieldMap["reserve_day_cnt"] = w.ReserveDayCnt
	w.fieldMap["input_time"] = w.InputTime
	w.fieldMap["bigo"] = w.Bigo
	w.fieldMap["sun_time"] = w.SunTime
	w.fieldMap["sun_etime"] = w.SunEtime
	w.fieldMap["sun_use_yn"] = w.SunUseYn
	w.fieldMap["vacation_stime"] = w.VacationStime
	w.fieldMap["vacation_etime"] = w.VacationEtime
	w.fieldMap["pat_limit"] = w.PatLimit
}

func (w wiseSroomLimit) clone(db *gorm.DB) wiseSroomLimit {
	w.wiseSroomLimitDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wiseSroomLimit) replaceDB(db *gorm.DB) wiseSroomLimit {
	w.wiseSroomLimitDo.ReplaceDB(db)
	return w
}

type wiseSroomLimitDo struct{ gen.DO }

type IWiseSroomLimitDo interface {
	gen.SubQuery
	Debug() IWiseSroomLimitDo
	WithContext(ctx context.Context) IWiseSroomLimitDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWiseSroomLimitDo
	WriteDB() IWiseSroomLimitDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWiseSroomLimitDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWiseSroomLimitDo
	Not(conds ...gen.Condition) IWiseSroomLimitDo
	Or(conds ...gen.Condition) IWiseSroomLimitDo
	Select(conds ...field.Expr) IWiseSroomLimitDo
	Where(conds ...gen.Condition) IWiseSroomLimitDo
	Order(conds ...field.Expr) IWiseSroomLimitDo
	Distinct(cols ...field.Expr) IWiseSroomLimitDo
	Omit(cols ...field.Expr) IWiseSroomLimitDo
	Join(table schema.Tabler, on ...field.Expr) IWiseSroomLimitDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWiseSroomLimitDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWiseSroomLimitDo
	Group(cols ...field.Expr) IWiseSroomLimitDo
	Having(conds ...gen.Condition) IWiseSroomLimitDo
	Limit(limit int) IWiseSroomLimitDo
	Offset(offset int) IWiseSroomLimitDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseSroomLimitDo
	Unscoped() IWiseSroomLimitDo
	Create(values ...*model.WiseSroomLimit) error
	CreateInBatches(values []*model.WiseSroomLimit, batchSize int) error
	Save(values ...*model.WiseSroomLimit) error
	First() (*model.WiseSroomLimit, error)
	Take() (*model.WiseSroomLimit, error)
	Last() (*model.WiseSroomLimit, error)
	Find() ([]*model.WiseSroomLimit, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseSroomLimit, err error)
	FindInBatches(result *[]*model.WiseSroomLimit, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WiseSroomLimit) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWiseSroomLimitDo
	Assign(attrs ...field.AssignExpr) IWiseSroomLimitDo
	Joins(fields ...field.RelationField) IWiseSroomLimitDo
	Preload(fields ...field.RelationField) IWiseSroomLimitDo
	FirstOrInit() (*model.WiseSroomLimit, error)
	FirstOrCreate() (*model.WiseSroomLimit, error)
	FindByPage(offset int, limit int) (result []*model.WiseSroomLimit, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWiseSroomLimitDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wiseSroomLimitDo) Debug() IWiseSroomLimitDo {
	return w.withDO(w.DO.Debug())
}

func (w wiseSroomLimitDo) WithContext(ctx context.Context) IWiseSroomLimitDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wiseSroomLimitDo) ReadDB() IWiseSroomLimitDo {
	return w.Clauses(dbresolver.Read)
}

func (w wiseSroomLimitDo) WriteDB() IWiseSroomLimitDo {
	return w.Clauses(dbresolver.Write)
}

func (w wiseSroomLimitDo) Session(config *gorm.Session) IWiseSroomLimitDo {
	return w.withDO(w.DO.Session(config))
}

func (w wiseSroomLimitDo) Clauses(conds ...clause.Expression) IWiseSroomLimitDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wiseSroomLimitDo) Returning(value interface{}, columns ...string) IWiseSroomLimitDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wiseSroomLimitDo) Not(conds ...gen.Condition) IWiseSroomLimitDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wiseSroomLimitDo) Or(conds ...gen.Condition) IWiseSroomLimitDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wiseSroomLimitDo) Select(conds ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wiseSroomLimitDo) Where(conds ...gen.Condition) IWiseSroomLimitDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wiseSroomLimitDo) Order(conds ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wiseSroomLimitDo) Distinct(cols ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wiseSroomLimitDo) Omit(cols ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wiseSroomLimitDo) Join(table schema.Tabler, on ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wiseSroomLimitDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wiseSroomLimitDo) RightJoin(table schema.Tabler, on ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wiseSroomLimitDo) Group(cols ...field.Expr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wiseSroomLimitDo) Having(conds ...gen.Condition) IWiseSroomLimitDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wiseSroomLimitDo) Limit(limit int) IWiseSroomLimitDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wiseSroomLimitDo) Offset(offset int) IWiseSroomLimitDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wiseSroomLimitDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWiseSroomLimitDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wiseSroomLimitDo) Unscoped() IWiseSroomLimitDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wiseSroomLimitDo) Create(values ...*model.WiseSroomLimit) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wiseSroomLimitDo) CreateInBatches(values []*model.WiseSroomLimit, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wiseSroomLimitDo) Save(values ...*model.WiseSroomLimit) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wiseSroomLimitDo) First() (*model.WiseSroomLimit, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseSroomLimit), nil
	}
}

func (w wiseSroomLimitDo) Take() (*model.WiseSroomLimit, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseSroomLimit), nil
	}
}

func (w wiseSroomLimitDo) Last() (*model.WiseSroomLimit, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseSroomLimit), nil
	}
}

func (w wiseSroomLimitDo) Find() ([]*model.WiseSroomLimit, error) {
	result, err := w.DO.Find()
	return result.([]*model.WiseSroomLimit), err
}

func (w wiseSroomLimitDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WiseSroomLimit, err error) {
	buf := make([]*model.WiseSroomLimit, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wiseSroomLimitDo) FindInBatches(result *[]*model.WiseSroomLimit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wiseSroomLimitDo) Attrs(attrs ...field.AssignExpr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wiseSroomLimitDo) Assign(attrs ...field.AssignExpr) IWiseSroomLimitDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wiseSroomLimitDo) Joins(fields ...field.RelationField) IWiseSroomLimitDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wiseSroomLimitDo) Preload(fields ...field.RelationField) IWiseSroomLimitDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wiseSroomLimitDo) FirstOrInit() (*model.WiseSroomLimit, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseSroomLimit), nil
	}
}

func (w wiseSroomLimitDo) FirstOrCreate() (*model.WiseSroomLimit, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WiseSroomLimit), nil
	}
}

func (w wiseSroomLimitDo) FindByPage(offset int, limit int) (result []*model.WiseSroomLimit, count int64, err error) {
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

func (w wiseSroomLimitDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wiseSroomLimitDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wiseSroomLimitDo) Delete(models ...*model.WiseSroomLimit) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wiseSroomLimitDo) withDO(do gen.Dao) *wiseSroomLimitDo {
	w.DO = *do.(*gen.DO)
	return w
}
