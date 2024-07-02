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

func newHCV_vw_seat_ip_setting(db *gorm.DB, opts ...gen.DOOption) hCV_vw_seat_ip_setting {
	_hCV_vw_seat_ip_setting := hCV_vw_seat_ip_setting{}

	_hCV_vw_seat_ip_setting.hCV_vw_seat_ip_settingDo.UseDB(db, opts...)
	_hCV_vw_seat_ip_setting.hCV_vw_seat_ip_settingDo.UseModel(&model.HCV_vw_seat_ip_setting{})

	tableName := _hCV_vw_seat_ip_setting.hCV_vw_seat_ip_settingDo.TableName()
	_hCV_vw_seat_ip_setting.ALL = field.NewAsterisk(tableName)
	_hCV_vw_seat_ip_setting.RoomNo = field.NewInt64(tableName, "room_no")
	_hCV_vw_seat_ip_setting.RoomName = field.NewString(tableName, "room_name")
	_hCV_vw_seat_ip_setting.SeatNo = field.NewInt64(tableName, "seat_no")
	_hCV_vw_seat_ip_setting.SeatName = field.NewString(tableName, "seat_name")
	_hCV_vw_seat_ip_setting.DcuNo = field.NewInt64(tableName, "dcu_no")
	_hCV_vw_seat_ip_setting.DtcNo = field.NewInt64(tableName, "dtc_no")
	_hCV_vw_seat_ip_setting.UnitNo = field.NewInt64(tableName, "unit_no")
	_hCV_vw_seat_ip_setting.SeatType = field.NewString(tableName, "seat_type")
	_hCV_vw_seat_ip_setting.LockerPwd = field.NewString(tableName, "locker_pwd")
	_hCV_vw_seat_ip_setting.LockPwd = field.NewString(tableName, "lock_pwd")
	_hCV_vw_seat_ip_setting.UseYn = field.NewString(tableName, "use_yn")
	_hCV_vw_seat_ip_setting.PcIP = field.NewString(tableName, "pc_ip")
	_hCV_vw_seat_ip_setting.PcMac = field.NewString(tableName, "pc_mac")
	_hCV_vw_seat_ip_setting.SelfYn = field.NewString(tableName, "self_yn")
	_hCV_vw_seat_ip_setting.RealUseCnt = field.NewInt64(tableName, "real_use_cnt")
	_hCV_vw_seat_ip_setting.MissLogCnt = field.NewInt64(tableName, "miss_log_cnt")
	_hCV_vw_seat_ip_setting.MissQueueCnt = field.NewInt64(tableName, "miss_queue_cnt")
	_hCV_vw_seat_ip_setting.ReserveStat = field.NewInt64(tableName, "reserve_stat")
	_hCV_vw_seat_ip_setting.SroomUseYn = field.NewString(tableName, "sroom_use_yn")

	_hCV_vw_seat_ip_setting.fillFieldMap()

	return _hCV_vw_seat_ip_setting
}

type hCV_vw_seat_ip_setting struct {
	hCV_vw_seat_ip_settingDo

	ALL          field.Asterisk
	RoomNo       field.Int64
	RoomName     field.String
	SeatNo       field.Int64
	SeatName     field.String
	DcuNo        field.Int64
	DtcNo        field.Int64
	UnitNo       field.Int64
	SeatType     field.String
	LockerPwd    field.String
	LockPwd      field.String
	UseYn        field.String
	PcIP         field.String
	PcMac        field.String
	SelfYn       field.String
	RealUseCnt   field.Int64
	MissLogCnt   field.Int64
	MissQueueCnt field.Int64
	ReserveStat  field.Int64
	SroomUseYn   field.String

	fieldMap map[string]field.Expr
}

func (h hCV_vw_seat_ip_setting) Table(newTableName string) *hCV_vw_seat_ip_setting {
	h.hCV_vw_seat_ip_settingDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_vw_seat_ip_setting) As(alias string) *hCV_vw_seat_ip_setting {
	h.hCV_vw_seat_ip_settingDo.DO = *(h.hCV_vw_seat_ip_settingDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_vw_seat_ip_setting) updateTableName(table string) *hCV_vw_seat_ip_setting {
	h.ALL = field.NewAsterisk(table)
	h.RoomNo = field.NewInt64(table, "room_no")
	h.RoomName = field.NewString(table, "room_name")
	h.SeatNo = field.NewInt64(table, "seat_no")
	h.SeatName = field.NewString(table, "seat_name")
	h.DcuNo = field.NewInt64(table, "dcu_no")
	h.DtcNo = field.NewInt64(table, "dtc_no")
	h.UnitNo = field.NewInt64(table, "unit_no")
	h.SeatType = field.NewString(table, "seat_type")
	h.LockerPwd = field.NewString(table, "locker_pwd")
	h.LockPwd = field.NewString(table, "lock_pwd")
	h.UseYn = field.NewString(table, "use_yn")
	h.PcIP = field.NewString(table, "pc_ip")
	h.PcMac = field.NewString(table, "pc_mac")
	h.SelfYn = field.NewString(table, "self_yn")
	h.RealUseCnt = field.NewInt64(table, "real_use_cnt")
	h.MissLogCnt = field.NewInt64(table, "miss_log_cnt")
	h.MissQueueCnt = field.NewInt64(table, "miss_queue_cnt")
	h.ReserveStat = field.NewInt64(table, "reserve_stat")
	h.SroomUseYn = field.NewString(table, "sroom_use_yn")

	h.fillFieldMap()

	return h
}

func (h *hCV_vw_seat_ip_setting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_vw_seat_ip_setting) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 19)
	h.fieldMap["room_no"] = h.RoomNo
	h.fieldMap["room_name"] = h.RoomName
	h.fieldMap["seat_no"] = h.SeatNo
	h.fieldMap["seat_name"] = h.SeatName
	h.fieldMap["dcu_no"] = h.DcuNo
	h.fieldMap["dtc_no"] = h.DtcNo
	h.fieldMap["unit_no"] = h.UnitNo
	h.fieldMap["seat_type"] = h.SeatType
	h.fieldMap["locker_pwd"] = h.LockerPwd
	h.fieldMap["lock_pwd"] = h.LockPwd
	h.fieldMap["use_yn"] = h.UseYn
	h.fieldMap["pc_ip"] = h.PcIP
	h.fieldMap["pc_mac"] = h.PcMac
	h.fieldMap["self_yn"] = h.SelfYn
	h.fieldMap["real_use_cnt"] = h.RealUseCnt
	h.fieldMap["miss_log_cnt"] = h.MissLogCnt
	h.fieldMap["miss_queue_cnt"] = h.MissQueueCnt
	h.fieldMap["reserve_stat"] = h.ReserveStat
	h.fieldMap["sroom_use_yn"] = h.SroomUseYn
}

func (h hCV_vw_seat_ip_setting) clone(db *gorm.DB) hCV_vw_seat_ip_setting {
	h.hCV_vw_seat_ip_settingDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_vw_seat_ip_setting) replaceDB(db *gorm.DB) hCV_vw_seat_ip_setting {
	h.hCV_vw_seat_ip_settingDo.ReplaceDB(db)
	return h
}

type hCV_vw_seat_ip_settingDo struct{ gen.DO }

type IHCV_vw_seat_ip_settingDo interface {
	gen.SubQuery
	Debug() IHCV_vw_seat_ip_settingDo
	WithContext(ctx context.Context) IHCV_vw_seat_ip_settingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_vw_seat_ip_settingDo
	WriteDB() IHCV_vw_seat_ip_settingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_vw_seat_ip_settingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_vw_seat_ip_settingDo
	Not(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo
	Or(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo
	Select(conds ...field.Expr) IHCV_vw_seat_ip_settingDo
	Where(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo
	Order(conds ...field.Expr) IHCV_vw_seat_ip_settingDo
	Distinct(cols ...field.Expr) IHCV_vw_seat_ip_settingDo
	Omit(cols ...field.Expr) IHCV_vw_seat_ip_settingDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_vw_seat_ip_settingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_seat_ip_settingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_seat_ip_settingDo
	Group(cols ...field.Expr) IHCV_vw_seat_ip_settingDo
	Having(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo
	Limit(limit int) IHCV_vw_seat_ip_settingDo
	Offset(offset int) IHCV_vw_seat_ip_settingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_seat_ip_settingDo
	Unscoped() IHCV_vw_seat_ip_settingDo
	Create(values ...*model.HCV_vw_seat_ip_setting) error
	CreateInBatches(values []*model.HCV_vw_seat_ip_setting, batchSize int) error
	Save(values ...*model.HCV_vw_seat_ip_setting) error
	First() (*model.HCV_vw_seat_ip_setting, error)
	Take() (*model.HCV_vw_seat_ip_setting, error)
	Last() (*model.HCV_vw_seat_ip_setting, error)
	Find() ([]*model.HCV_vw_seat_ip_setting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_seat_ip_setting, err error)
	FindInBatches(result *[]*model.HCV_vw_seat_ip_setting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_vw_seat_ip_setting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_vw_seat_ip_settingDo
	Assign(attrs ...field.AssignExpr) IHCV_vw_seat_ip_settingDo
	Joins(fields ...field.RelationField) IHCV_vw_seat_ip_settingDo
	Preload(fields ...field.RelationField) IHCV_vw_seat_ip_settingDo
	FirstOrInit() (*model.HCV_vw_seat_ip_setting, error)
	FirstOrCreate() (*model.HCV_vw_seat_ip_setting, error)
	FindByPage(offset int, limit int) (result []*model.HCV_vw_seat_ip_setting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_vw_seat_ip_settingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_vw_seat_ip_settingDo) Debug() IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_vw_seat_ip_settingDo) WithContext(ctx context.Context) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_vw_seat_ip_settingDo) ReadDB() IHCV_vw_seat_ip_settingDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_vw_seat_ip_settingDo) WriteDB() IHCV_vw_seat_ip_settingDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_vw_seat_ip_settingDo) Session(config *gorm.Session) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_vw_seat_ip_settingDo) Clauses(conds ...clause.Expression) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Returning(value interface{}, columns ...string) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_vw_seat_ip_settingDo) Not(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Or(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Select(conds ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Where(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Order(conds ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Distinct(cols ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_vw_seat_ip_settingDo) Omit(cols ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_vw_seat_ip_settingDo) Join(table schema.Tabler, on ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_vw_seat_ip_settingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_vw_seat_ip_settingDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_vw_seat_ip_settingDo) Group(cols ...field.Expr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_vw_seat_ip_settingDo) Having(conds ...gen.Condition) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_vw_seat_ip_settingDo) Limit(limit int) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_vw_seat_ip_settingDo) Offset(offset int) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_vw_seat_ip_settingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_vw_seat_ip_settingDo) Unscoped() IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_vw_seat_ip_settingDo) Create(values ...*model.HCV_vw_seat_ip_setting) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_vw_seat_ip_settingDo) CreateInBatches(values []*model.HCV_vw_seat_ip_setting, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_vw_seat_ip_settingDo) Save(values ...*model.HCV_vw_seat_ip_setting) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_vw_seat_ip_settingDo) First() (*model.HCV_vw_seat_ip_setting, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_seat_ip_setting), nil
	}
}

func (h hCV_vw_seat_ip_settingDo) Take() (*model.HCV_vw_seat_ip_setting, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_seat_ip_setting), nil
	}
}

func (h hCV_vw_seat_ip_settingDo) Last() (*model.HCV_vw_seat_ip_setting, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_seat_ip_setting), nil
	}
}

func (h hCV_vw_seat_ip_settingDo) Find() ([]*model.HCV_vw_seat_ip_setting, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_vw_seat_ip_setting), err
}

func (h hCV_vw_seat_ip_settingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_vw_seat_ip_setting, err error) {
	buf := make([]*model.HCV_vw_seat_ip_setting, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_vw_seat_ip_settingDo) FindInBatches(result *[]*model.HCV_vw_seat_ip_setting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_vw_seat_ip_settingDo) Attrs(attrs ...field.AssignExpr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_vw_seat_ip_settingDo) Assign(attrs ...field.AssignExpr) IHCV_vw_seat_ip_settingDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_vw_seat_ip_settingDo) Joins(fields ...field.RelationField) IHCV_vw_seat_ip_settingDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_vw_seat_ip_settingDo) Preload(fields ...field.RelationField) IHCV_vw_seat_ip_settingDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_vw_seat_ip_settingDo) FirstOrInit() (*model.HCV_vw_seat_ip_setting, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_seat_ip_setting), nil
	}
}

func (h hCV_vw_seat_ip_settingDo) FirstOrCreate() (*model.HCV_vw_seat_ip_setting, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_vw_seat_ip_setting), nil
	}
}

func (h hCV_vw_seat_ip_settingDo) FindByPage(offset int, limit int) (result []*model.HCV_vw_seat_ip_setting, count int64, err error) {
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

func (h hCV_vw_seat_ip_settingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_vw_seat_ip_settingDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_vw_seat_ip_settingDo) Delete(models ...*model.HCV_vw_seat_ip_setting) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_vw_seat_ip_settingDo) withDO(do gen.Dao) *hCV_vw_seat_ip_settingDo {
	h.DO = *do.(*gen.DO)
	return h
}
