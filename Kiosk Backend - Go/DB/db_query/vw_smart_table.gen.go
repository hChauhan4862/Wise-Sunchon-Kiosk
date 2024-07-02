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

func newHCV_vw_smart_table(db *gorm.DB, opts ...gen.DOOption) hCV_vw_smart_table {
	_hCV_vw_smart_table := hCV_vw_smart_table{}

	_hCV_vw_smart_table.hCV_vw_smart_tableDo.UseDB(db, opts...)
	_hCV_vw_smart_table.hCV_vw_smart_tableDo.UseModel(&model.HCV_vw_smart_table{})

	tableName := _hCV_vw_smart_table.hCV_vw_smart_tableDo.TableName()
	_hCV_vw_smart_table.ALL = field.NewAsterisk(tableName)
	_hCV_vw_smart_table.RoomNo = field.NewInt64(tableName, "room_no")
	_hCV_vw_smart_table.RoomName = field.NewString(tableName, "room_name")
	_hCV_vw_smart_table.SeatNo = field.NewInt64(tableName, "seat_no")
	_hCV_vw_smart_table.DcuNo = field.NewInt64(tableName, "dcu_no")
	_hCV_vw_smart_table.DtcNo = field.NewInt64(tableName, "dtc_no")
	_hCV_vw_smart_table.UnitNo = field.NewInt64(tableName, "unit_no")
	_hCV_vw_smart_table.SeatType = field.NewString(tableName, "seat_type")
	_hCV_vw_smart_table.LockerPwd = field.NewString(tableName, "locker_pwd")
	_hCV_vw_smart_table.UseYn = field.NewString(tableName, "use_yn")
	_hCV_vw_smart_table.RealUseCnt = field.NewInt64(tableName, "real_use_cnt")
	_hCV_vw_smart_table.MissLogCnt = field.NewInt64(tableName, "miss_log_cnt")
	_hCV_vw_smart_table.MissQueueCnt = field.NewInt64(tableName, "miss_queue_cnt")

	_hCV_vw_smart_table.fillFieldMap()

	return _hCV_vw_smart_table
}

type hCV_vw_smart_table struct {
	hCV_vw_smart_tableDo

	ALL          field.Asterisk
	RoomNo       field.Int64
	RoomName     field.String
	SeatNo       field.Int64
	DcuNo        field.Int64
	DtcNo        field.Int64
	UnitNo       field.Int64
	SeatType     field.String
	LockerPwd    field.String
	UseYn        field.String
	RealUseCnt   field.Int64
	MissLogCnt   field.Int64
	MissQueueCnt field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_vw_smart_table) Table(newTableName string) *hCV_vw_smart_table {
	h.hCV_vw_smart_tableDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_vw_smart_table) As(alias string) *hCV_vw_smart_table {
	h.hCV_vw_smart_tableDo.DO = *(h.hCV_vw_smart_tableDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_vw_smart_table) updateTableName(table string) *hCV_vw_smart_table {
	h.ALL = field.NewAsterisk(table)
	h.RoomNo = field.NewInt64(table, "room_no")
	h.RoomName = field.NewString(table, "room_name")
	h.SeatNo = field.NewInt64(table, "seat_no")
	h.DcuNo = field.NewInt64(table, "dcu_no")
	h.DtcNo = field.NewInt64(table, "dtc_no")
	h.UnitNo = field.NewInt64(table, "unit_no")
	h.SeatType = field.NewString(table, "seat_type")
	h.LockerPwd = field.NewString(table, "locker_pwd")
	h.UseYn = field.NewString(table, "use_yn")
	h.RealUseCnt = field.NewInt64(table, "real_use_cnt")
	h.MissLogCnt = field.NewInt64(table, "miss_log_cnt")
	h.MissQueueCnt = field.NewInt64(table, "miss_queue_cnt")

	h.fillFieldMap()

	return h
}

func (h *hCV_vw_smart_table) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_vw_smart_table) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 12)
	h.fieldMap["room_no"] = h.RoomNo
	h.fieldMap["room_name"] = h.RoomName
	h.fieldMap["seat_no"] = h.SeatNo
	h.fieldMap["dcu_no"] = h.DcuNo
	h.fieldMap["dtc_no"] = h.DtcNo
	h.fieldMap["unit_no"] = h.UnitNo
	h.fieldMap["seat_type"] = h.SeatType
	h.fieldMap["locker_pwd"] = h.LockerPwd
	h.fieldMap["use_yn"] = h.UseYn
	h.fieldMap["real_use_cnt"] = h.RealUseCnt
	h.fieldMap["miss_log_cnt"] = h.MissLogCnt
	h.fieldMap["miss_queue_cnt"] = h.MissQueueCnt
}

func (h hCV_vw_smart_table) clone(db *gorm.DB) hCV_vw_smart_table {
	h.hCV_vw_smart_tableDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_vw_smart_table) replaceDB(db *gorm.DB) hCV_vw_smart_table {
	h.hCV_vw_smart_tableDo.ReplaceDB(db)
	return h
}

type hCV_vw_smart_tableDo struct{ gen.DO }

type IHCV_vw_smart_tableDo interface {
	gen.SubQuery
	Debug() IHCV_vw_smart_tableDo
	WithContext(ctx context.Context) IHCV_vw_smart_tableDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_vw_smart_tableDo
	WriteDB() IHCV_vw_smart_tableDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_vw_smart_tableDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_vw_smart_tableDo
	Not(conds ...gen.Condition) IHCV_vw_smart_tableDo
	Or(conds ...gen.Condition) IHCV_vw_smart_tableDo
	Select(conds ...field.Expr) IHCV_vw_smart_tableDo
	Where(conds ...gen.Condition) IHCV_vw_smart_tableDo
	Order(conds ...field.Expr) IHCV_vw_smart_tableDo
	Distinct(cols ...field.Expr) IHCV_vw_smart_tableDo
	Omit(cols ...field.Expr) IHCV_vw_smart_tableDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_vw_smart_tableDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_smart_tableDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_smart_tableDo
	Group(cols ...field.Expr) IHCV_vw_smart_tableDo
	Having(conds ...gen.Condition) IHCV_vw_smart_tableDo
	Limit(limit int) IHCV_vw_smart_tableDo
	Offset(offset int) IHCV_vw_smart_tableDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_smart_tableDo
	Unscoped() IHCV_vw_smart_tableDo
	Create(values ...*model.HCV_vw_smart_table) error
	CreateInBatches(values []*model.HCV_vw_smart_table, batchSize int) error
	Save(values ...*model.HCV_vw_smart_table) error
	First() (*model.HCV_vw_smart_table, error)
	Take() (*model.HCV_vw_smart_table, error)
	Last() (*model.HCV_vw_smart_table, error)
	Find() ([]*model.HCV_vw_smart_table, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_smart_table, err error)
	FindInBatches(result *[]*model.HCV_vw_smart_table, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_vw_smart_table) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_vw_smart_tableDo
	Assign(attrs ...field.AssignExpr) IHCV_vw_smart_tableDo
	Joins(fields ...field.RelationField) IHCV_vw_smart_tableDo
	Preload(fields ...field.RelationField) IHCV_vw_smart_tableDo
	FirstOrInit() (*model.HCV_vw_smart_table, error)
	FirstOrCreate() (*model.HCV_vw_smart_table, error)
	FindByPage(offset int, limit int) (result []*model.HCV_vw_smart_table, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_vw_smart_tableDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_vw_smart_tableDo) Debug() IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_vw_smart_tableDo) WithContext(ctx context.Context) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_vw_smart_tableDo) ReadDB() IHCV_vw_smart_tableDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_vw_smart_tableDo) WriteDB() IHCV_vw_smart_tableDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_vw_smart_tableDo) Session(config *gorm.Session) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_vw_smart_tableDo) Clauses(conds ...clause.Expression) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_vw_smart_tableDo) Returning(value interface{}, columns ...string) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_vw_smart_tableDo) Not(conds ...gen.Condition) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_vw_smart_tableDo) Or(conds ...gen.Condition) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_vw_smart_tableDo) Select(conds ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_vw_smart_tableDo) Where(conds ...gen.Condition) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_vw_smart_tableDo) Order(conds ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_vw_smart_tableDo) Distinct(cols ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_vw_smart_tableDo) Omit(cols ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_vw_smart_tableDo) Join(table schema.Tabler, on ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_vw_smart_tableDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_vw_smart_tableDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_vw_smart_tableDo) Group(cols ...field.Expr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_vw_smart_tableDo) Having(conds ...gen.Condition) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_vw_smart_tableDo) Limit(limit int) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_vw_smart_tableDo) Offset(offset int) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_vw_smart_tableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_vw_smart_tableDo) Unscoped() IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_vw_smart_tableDo) Create(values ...*model.HCV_vw_smart_table) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_vw_smart_tableDo) CreateInBatches(values []*model.HCV_vw_smart_table, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_vw_smart_tableDo) Save(values ...*model.HCV_vw_smart_table) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_vw_smart_tableDo) First() (*model.HCV_vw_smart_table, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_smart_table), nil
	}
}

func (h hCV_vw_smart_tableDo) Take() (*model.HCV_vw_smart_table, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_smart_table), nil
	}
}

func (h hCV_vw_smart_tableDo) Last() (*model.HCV_vw_smart_table, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_smart_table), nil
	}
}

func (h hCV_vw_smart_tableDo) Find() ([]*model.HCV_vw_smart_table, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_vw_smart_table), err
}

func (h hCV_vw_smart_tableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_smart_table, err error) {
	buf := make([]*model.HCV_vw_smart_table, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_vw_smart_tableDo) FindInBatches(result *[]*model.HCV_vw_smart_table, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_vw_smart_tableDo) Attrs(attrs ...field.AssignExpr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_vw_smart_tableDo) Assign(attrs ...field.AssignExpr) IHCV_vw_smart_tableDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_vw_smart_tableDo) Joins(fields ...field.RelationField) IHCV_vw_smart_tableDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_vw_smart_tableDo) Preload(fields ...field.RelationField) IHCV_vw_smart_tableDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_vw_smart_tableDo) FirstOrInit() (*model.HCV_vw_smart_table, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_smart_table), nil
	}
}

func (h hCV_vw_smart_tableDo) FirstOrCreate() (*model.HCV_vw_smart_table, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_smart_table), nil
	}
}

func (h hCV_vw_smart_tableDo) FindByPage(offset int, limit int) (result []*model.HCV_vw_smart_table, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hCV_vw_smart_tableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_vw_smart_tableDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_vw_smart_tableDo) Delete(models ...*model.HCV_vw_smart_table) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_vw_smart_tableDo) withDO(do gen.Dao) *hCV_vw_smart_tableDo {
	h.DO = *do.(*gen.DO)
	return h
}
