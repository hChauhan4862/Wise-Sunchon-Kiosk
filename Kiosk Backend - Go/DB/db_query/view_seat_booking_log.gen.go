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

func newHCV_view_seat_booking_log(db *gorm.DB, opts ...gen.DOOption) hCV_view_seat_booking_log {
	_hCV_view_seat_booking_log := hCV_view_seat_booking_log{}

	_hCV_view_seat_booking_log.hCV_view_seat_booking_logDo.UseDB(db, opts...)
	_hCV_view_seat_booking_log.hCV_view_seat_booking_logDo.UseModel(&model.HCV_view_seat_booking_log{})

	tableName := _hCV_view_seat_booking_log.hCV_view_seat_booking_logDo.TableName()
	_hCV_view_seat_booking_log.ALL = field.NewAsterisk(tableName)
	_hCV_view_seat_booking_log.LOGNO = field.NewInt64(tableName, "LOGNO")
	_hCV_view_seat_booking_log.BSEQNO = field.NewInt64(tableName, "BSEQNO")
	_hCV_view_seat_booking_log.SCHOOLNO = field.NewString(tableName, "SCHOOLNO")
	_hCV_view_seat_booking_log.SEATNO = field.NewInt64(tableName, "SEATNO")
	_hCV_view_seat_booking_log.STATUS = field.NewInt64(tableName, "STATUS")
	_hCV_view_seat_booking_log.StatusName = field.NewString(tableName, "status_name")
	_hCV_view_seat_booking_log.StatusEnName = field.NewString(tableName, "status_en_name")
	_hCV_view_seat_booking_log.USESTART = field.NewTime(tableName, "USESTART")
	_hCV_view_seat_booking_log.USEEXPIRE = field.NewTime(tableName, "USEEXPIRE")
	_hCV_view_seat_booking_log.STARTTIME = field.NewTime(tableName, "STARTTIME")
	_hCV_view_seat_booking_log.RETURNTIME = field.NewTime(tableName, "RETURNTIME")
	_hCV_view_seat_booking_log.NEWSEATNO = field.NewInt64(tableName, "NEW_SEATNO")
	_hCV_view_seat_booking_log.EXTENDMIN = field.NewInt64(tableName, "EXTEND_MIN")
	_hCV_view_seat_booking_log.EXTENDCNT = field.NewInt64(tableName, "EXTEND_CNT")
	_hCV_view_seat_booking_log.LOGTIME = field.NewTime(tableName, "LOGTIME")
	_hCV_view_seat_booking_log.ISSUEFROM = field.NewInt64(tableName, "ISSUEFROM")
	_hCV_view_seat_booking_log.ISSUEDETAIL = field.NewString(tableName, "ISSUEDETAIL")
	_hCV_view_seat_booking_log.ISSUETYPENO = field.NewInt64(tableName, "ISSUE_TYPE_NO")
	_hCV_view_seat_booking_log.LIBNO = field.NewInt64(tableName, "LIBNO")
	_hCV_view_seat_booking_log.LibName = field.NewString(tableName, "lib_name")
	_hCV_view_seat_booking_log.LibEnName = field.NewString(tableName, "lib_en_name")
	_hCV_view_seat_booking_log.FLOORNO = field.NewInt64(tableName, "FLOORNO")
	_hCV_view_seat_booking_log.FloorName = field.NewString(tableName, "floor_name")
	_hCV_view_seat_booking_log.FloorEnName = field.NewString(tableName, "floor_en_name")
	_hCV_view_seat_booking_log.Floor = field.NewInt64(tableName, "floor")
	_hCV_view_seat_booking_log.ROOMNO = field.NewInt64(tableName, "ROOMNO")
	_hCV_view_seat_booking_log.RoomName = field.NewString(tableName, "room_name")
	_hCV_view_seat_booking_log.RoomEnName = field.NewString(tableName, "room_en_name")
	_hCV_view_seat_booking_log.SECTORNO = field.NewInt64(tableName, "SECTORNO")
	_hCV_view_seat_booking_log.SectorName = field.NewString(tableName, "sector_name")
	_hCV_view_seat_booking_log.SectorEnName = field.NewString(tableName, "sector_en_name")
	_hCV_view_seat_booking_log.TYPENO = field.NewInt64(tableName, "TYPENO")
	_hCV_view_seat_booking_log.SeatTypeName = field.NewString(tableName, "seat_type_name")
	_hCV_view_seat_booking_log.SeatTypeEnName = field.NewString(tableName, "seat_type_en_name")
	_hCV_view_seat_booking_log.SeatName = field.NewString(tableName, "seat_name")
	_hCV_view_seat_booking_log.SeatEnName = field.NewString(tableName, "seat_en_name")
	_hCV_view_seat_booking_log.SeatStatus = field.NewInt64(tableName, "seat_status")
	_hCV_view_seat_booking_log.SeatStatusName = field.NewString(tableName, "seat_status_name")
	_hCV_view_seat_booking_log.SeatStatusEnName = field.NewString(tableName, "seat_status_en_name")
	_hCV_view_seat_booking_log.SeatVname = field.NewString(tableName, "seat_vname")
	_hCV_view_seat_booking_log.NewSeatName = field.NewString(tableName, "new_seat_name")
	_hCV_view_seat_booking_log.NewSeatEnName = field.NewString(tableName, "new_seat_en_name")
	_hCV_view_seat_booking_log.NewSeatStatus = field.NewInt64(tableName, "new_seat_status")
	_hCV_view_seat_booking_log.NewSeatVname = field.NewString(tableName, "new_seat_vname")

	_hCV_view_seat_booking_log.fillFieldMap()

	return _hCV_view_seat_booking_log
}

type hCV_view_seat_booking_log struct {
	hCV_view_seat_booking_logDo

	ALL              field.Asterisk
	LOGNO            field.Int64
	BSEQNO           field.Int64
	SCHOOLNO         field.String
	SEATNO           field.Int64
	STATUS           field.Int64
	StatusName       field.String
	StatusEnName     field.String
	USESTART         field.Time
	USEEXPIRE        field.Time
	STARTTIME        field.Time
	RETURNTIME       field.Time
	NEWSEATNO        field.Int64
	EXTENDMIN        field.Int64
	EXTENDCNT        field.Int64
	LOGTIME          field.Time
	ISSUEFROM        field.Int64
	ISSUEDETAIL      field.String
	ISSUETYPENO      field.Int64
	LIBNO            field.Int64
	LibName          field.String
	LibEnName        field.String
	FLOORNO          field.Int64
	FloorName        field.String
	FloorEnName      field.String
	Floor            field.Int64
	ROOMNO           field.Int64
	RoomName         field.String
	RoomEnName       field.String
	SECTORNO         field.Int64
	SectorName       field.String
	SectorEnName     field.String
	TYPENO           field.Int64
	SeatTypeName     field.String
	SeatTypeEnName   field.String
	SeatName         field.String
	SeatEnName       field.String
	SeatStatus       field.Int64
	SeatStatusName   field.String
	SeatStatusEnName field.String
	SeatVname        field.String
	NewSeatName      field.String
	NewSeatEnName    field.String
	NewSeatStatus    field.Int64
	NewSeatVname     field.String

	fieldMap map[string]field.Expr
}

func (h hCV_view_seat_booking_log) Table(newTableName string) *hCV_view_seat_booking_log {
	h.hCV_view_seat_booking_logDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_view_seat_booking_log) As(alias string) *hCV_view_seat_booking_log {
	h.hCV_view_seat_booking_logDo.DO = *(h.hCV_view_seat_booking_logDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_view_seat_booking_log) updateTableName(table string) *hCV_view_seat_booking_log {
	h.ALL = field.NewAsterisk(table)
	h.LOGNO = field.NewInt64(table, "LOGNO")
	h.BSEQNO = field.NewInt64(table, "BSEQNO")
	h.SCHOOLNO = field.NewString(table, "SCHOOLNO")
	h.SEATNO = field.NewInt64(table, "SEATNO")
	h.STATUS = field.NewInt64(table, "STATUS")
	h.StatusName = field.NewString(table, "status_name")
	h.StatusEnName = field.NewString(table, "status_en_name")
	h.USESTART = field.NewTime(table, "USESTART")
	h.USEEXPIRE = field.NewTime(table, "USEEXPIRE")
	h.STARTTIME = field.NewTime(table, "STARTTIME")
	h.RETURNTIME = field.NewTime(table, "RETURNTIME")
	h.NEWSEATNO = field.NewInt64(table, "NEW_SEATNO")
	h.EXTENDMIN = field.NewInt64(table, "EXTEND_MIN")
	h.EXTENDCNT = field.NewInt64(table, "EXTEND_CNT")
	h.LOGTIME = field.NewTime(table, "LOGTIME")
	h.ISSUEFROM = field.NewInt64(table, "ISSUEFROM")
	h.ISSUEDETAIL = field.NewString(table, "ISSUEDETAIL")
	h.ISSUETYPENO = field.NewInt64(table, "ISSUE_TYPE_NO")
	h.LIBNO = field.NewInt64(table, "LIBNO")
	h.LibName = field.NewString(table, "lib_name")
	h.LibEnName = field.NewString(table, "lib_en_name")
	h.FLOORNO = field.NewInt64(table, "FLOORNO")
	h.FloorName = field.NewString(table, "floor_name")
	h.FloorEnName = field.NewString(table, "floor_en_name")
	h.Floor = field.NewInt64(table, "floor")
	h.ROOMNO = field.NewInt64(table, "ROOMNO")
	h.RoomName = field.NewString(table, "room_name")
	h.RoomEnName = field.NewString(table, "room_en_name")
	h.SECTORNO = field.NewInt64(table, "SECTORNO")
	h.SectorName = field.NewString(table, "sector_name")
	h.SectorEnName = field.NewString(table, "sector_en_name")
	h.TYPENO = field.NewInt64(table, "TYPENO")
	h.SeatTypeName = field.NewString(table, "seat_type_name")
	h.SeatTypeEnName = field.NewString(table, "seat_type_en_name")
	h.SeatName = field.NewString(table, "seat_name")
	h.SeatEnName = field.NewString(table, "seat_en_name")
	h.SeatStatus = field.NewInt64(table, "seat_status")
	h.SeatStatusName = field.NewString(table, "seat_status_name")
	h.SeatStatusEnName = field.NewString(table, "seat_status_en_name")
	h.SeatVname = field.NewString(table, "seat_vname")
	h.NewSeatName = field.NewString(table, "new_seat_name")
	h.NewSeatEnName = field.NewString(table, "new_seat_en_name")
	h.NewSeatStatus = field.NewInt64(table, "new_seat_status")
	h.NewSeatVname = field.NewString(table, "new_seat_vname")

	h.fillFieldMap()

	return h
}

func (h *hCV_view_seat_booking_log) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_view_seat_booking_log) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 44)
	h.fieldMap["LOGNO"] = h.LOGNO
	h.fieldMap["BSEQNO"] = h.BSEQNO
	h.fieldMap["SCHOOLNO"] = h.SCHOOLNO
	h.fieldMap["SEATNO"] = h.SEATNO
	h.fieldMap["STATUS"] = h.STATUS
	h.fieldMap["status_name"] = h.StatusName
	h.fieldMap["status_en_name"] = h.StatusEnName
	h.fieldMap["USESTART"] = h.USESTART
	h.fieldMap["USEEXPIRE"] = h.USEEXPIRE
	h.fieldMap["STARTTIME"] = h.STARTTIME
	h.fieldMap["RETURNTIME"] = h.RETURNTIME
	h.fieldMap["NEW_SEATNO"] = h.NEWSEATNO
	h.fieldMap["EXTEND_MIN"] = h.EXTENDMIN
	h.fieldMap["EXTEND_CNT"] = h.EXTENDCNT
	h.fieldMap["LOGTIME"] = h.LOGTIME
	h.fieldMap["ISSUEFROM"] = h.ISSUEFROM
	h.fieldMap["ISSUEDETAIL"] = h.ISSUEDETAIL
	h.fieldMap["ISSUE_TYPE_NO"] = h.ISSUETYPENO
	h.fieldMap["LIBNO"] = h.LIBNO
	h.fieldMap["lib_name"] = h.LibName
	h.fieldMap["lib_en_name"] = h.LibEnName
	h.fieldMap["FLOORNO"] = h.FLOORNO
	h.fieldMap["floor_name"] = h.FloorName
	h.fieldMap["floor_en_name"] = h.FloorEnName
	h.fieldMap["floor"] = h.Floor
	h.fieldMap["ROOMNO"] = h.ROOMNO
	h.fieldMap["room_name"] = h.RoomName
	h.fieldMap["room_en_name"] = h.RoomEnName
	h.fieldMap["SECTORNO"] = h.SECTORNO
	h.fieldMap["sector_name"] = h.SectorName
	h.fieldMap["sector_en_name"] = h.SectorEnName
	h.fieldMap["TYPENO"] = h.TYPENO
	h.fieldMap["seat_type_name"] = h.SeatTypeName
	h.fieldMap["seat_type_en_name"] = h.SeatTypeEnName
	h.fieldMap["seat_name"] = h.SeatName
	h.fieldMap["seat_en_name"] = h.SeatEnName
	h.fieldMap["seat_status"] = h.SeatStatus
	h.fieldMap["seat_status_name"] = h.SeatStatusName
	h.fieldMap["seat_status_en_name"] = h.SeatStatusEnName
	h.fieldMap["seat_vname"] = h.SeatVname
	h.fieldMap["new_seat_name"] = h.NewSeatName
	h.fieldMap["new_seat_en_name"] = h.NewSeatEnName
	h.fieldMap["new_seat_status"] = h.NewSeatStatus
	h.fieldMap["new_seat_vname"] = h.NewSeatVname
}

func (h hCV_view_seat_booking_log) clone(db *gorm.DB) hCV_view_seat_booking_log {
	h.hCV_view_seat_booking_logDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_view_seat_booking_log) replaceDB(db *gorm.DB) hCV_view_seat_booking_log {
	h.hCV_view_seat_booking_logDo.ReplaceDB(db)
	return h
}

type hCV_view_seat_booking_logDo struct{ gen.DO }

type IHCV_view_seat_booking_logDo interface {
	gen.SubQuery
	Debug() IHCV_view_seat_booking_logDo
	WithContext(ctx context.Context) IHCV_view_seat_booking_logDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_view_seat_booking_logDo
	WriteDB() IHCV_view_seat_booking_logDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_view_seat_booking_logDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_view_seat_booking_logDo
	Not(conds ...gen.Condition) IHCV_view_seat_booking_logDo
	Or(conds ...gen.Condition) IHCV_view_seat_booking_logDo
	Select(conds ...field.Expr) IHCV_view_seat_booking_logDo
	Where(conds ...gen.Condition) IHCV_view_seat_booking_logDo
	Order(conds ...field.Expr) IHCV_view_seat_booking_logDo
	Distinct(cols ...field.Expr) IHCV_view_seat_booking_logDo
	Omit(cols ...field.Expr) IHCV_view_seat_booking_logDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_logDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_logDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_logDo
	Group(cols ...field.Expr) IHCV_view_seat_booking_logDo
	Having(conds ...gen.Condition) IHCV_view_seat_booking_logDo
	Limit(limit int) IHCV_view_seat_booking_logDo
	Offset(offset int) IHCV_view_seat_booking_logDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_seat_booking_logDo
	Unscoped() IHCV_view_seat_booking_logDo
	Create(values ...*model.HCV_view_seat_booking_log) error
	CreateInBatches(values []*model.HCV_view_seat_booking_log, batchSize int) error
	Save(values ...*model.HCV_view_seat_booking_log) error
	First() (*model.HCV_view_seat_booking_log, error)
	Take() (*model.HCV_view_seat_booking_log, error)
	Last() (*model.HCV_view_seat_booking_log, error)
	Find() ([]*model.HCV_view_seat_booking_log, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_seat_booking_log, err error)
	FindInBatches(result *[]*model.HCV_view_seat_booking_log, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_view_seat_booking_log) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_view_seat_booking_logDo
	Assign(attrs ...field.AssignExpr) IHCV_view_seat_booking_logDo
	Joins(fields ...field.RelationField) IHCV_view_seat_booking_logDo
	Preload(fields ...field.RelationField) IHCV_view_seat_booking_logDo
	FirstOrInit() (*model.HCV_view_seat_booking_log, error)
	FirstOrCreate() (*model.HCV_view_seat_booking_log, error)
	FindByPage(offset int, limit int) (result []*model.HCV_view_seat_booking_log, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_view_seat_booking_logDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_view_seat_booking_logDo) Debug() IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_view_seat_booking_logDo) WithContext(ctx context.Context) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_view_seat_booking_logDo) ReadDB() IHCV_view_seat_booking_logDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_view_seat_booking_logDo) WriteDB() IHCV_view_seat_booking_logDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_view_seat_booking_logDo) Session(config *gorm.Session) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_view_seat_booking_logDo) Clauses(conds ...clause.Expression) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_view_seat_booking_logDo) Returning(value interface{}, columns ...string) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_view_seat_booking_logDo) Not(conds ...gen.Condition) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_view_seat_booking_logDo) Or(conds ...gen.Condition) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_view_seat_booking_logDo) Select(conds ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_view_seat_booking_logDo) Where(conds ...gen.Condition) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_view_seat_booking_logDo) Order(conds ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_view_seat_booking_logDo) Distinct(cols ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_view_seat_booking_logDo) Omit(cols ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_view_seat_booking_logDo) Join(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_view_seat_booking_logDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_view_seat_booking_logDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_view_seat_booking_logDo) Group(cols ...field.Expr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_view_seat_booking_logDo) Having(conds ...gen.Condition) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_view_seat_booking_logDo) Limit(limit int) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_view_seat_booking_logDo) Offset(offset int) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_view_seat_booking_logDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_view_seat_booking_logDo) Unscoped() IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_view_seat_booking_logDo) Create(values ...*model.HCV_view_seat_booking_log) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_view_seat_booking_logDo) CreateInBatches(values []*model.HCV_view_seat_booking_log, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_view_seat_booking_logDo) Save(values ...*model.HCV_view_seat_booking_log) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_view_seat_booking_logDo) First() (*model.HCV_view_seat_booking_log, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_log), nil
	}
}

func (h hCV_view_seat_booking_logDo) Take() (*model.HCV_view_seat_booking_log, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_log), nil
	}
}

func (h hCV_view_seat_booking_logDo) Last() (*model.HCV_view_seat_booking_log, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_log), nil
	}
}

func (h hCV_view_seat_booking_logDo) Find() ([]*model.HCV_view_seat_booking_log, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_view_seat_booking_log), err
}

func (h hCV_view_seat_booking_logDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_seat_booking_log, err error) {
	buf := make([]*model.HCV_view_seat_booking_log, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_view_seat_booking_logDo) FindInBatches(result *[]*model.HCV_view_seat_booking_log, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_view_seat_booking_logDo) Attrs(attrs ...field.AssignExpr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_view_seat_booking_logDo) Assign(attrs ...field.AssignExpr) IHCV_view_seat_booking_logDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_view_seat_booking_logDo) Joins(fields ...field.RelationField) IHCV_view_seat_booking_logDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_view_seat_booking_logDo) Preload(fields ...field.RelationField) IHCV_view_seat_booking_logDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_view_seat_booking_logDo) FirstOrInit() (*model.HCV_view_seat_booking_log, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_log), nil
	}
}

func (h hCV_view_seat_booking_logDo) FirstOrCreate() (*model.HCV_view_seat_booking_log, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_log), nil
	}
}

func (h hCV_view_seat_booking_logDo) FindByPage(offset int, limit int) (result []*model.HCV_view_seat_booking_log, count int64, err error) {
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

func (h hCV_view_seat_booking_logDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_view_seat_booking_logDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_view_seat_booking_logDo) Delete(models ...*model.HCV_view_seat_booking_log) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_view_seat_booking_logDo) withDO(do gen.Dao) *hCV_view_seat_booking_logDo {
	h.DO = *do.(*gen.DO)
	return h
}