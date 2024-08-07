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

func newHCV_view_seat_seat(db *gorm.DB, opts ...gen.DOOption) hCV_view_seat_seat {
	_hCV_view_seat_seat := hCV_view_seat_seat{}

	_hCV_view_seat_seat.hCV_view_seat_seatDo.UseDB(db, opts...)
	_hCV_view_seat_seat.hCV_view_seat_seatDo.UseModel(&model.HCV_view_seat_seat{})

	tableName := _hCV_view_seat_seat.hCV_view_seat_seatDo.TableName()
	_hCV_view_seat_seat.ALL = field.NewAsterisk(tableName)
	_hCV_view_seat_seat.SEATNO = field.NewInt64(tableName, "SEATNO")
	_hCV_view_seat_seat.Name = field.NewString(tableName, "name")
	_hCV_view_seat_seat.EnName = field.NewString(tableName, "en_name")
	_hCV_view_seat_seat.Vname = field.NewString(tableName, "vname")
	_hCV_view_seat_seat.Status = field.NewInt64(tableName, "status")
	_hCV_view_seat_seat.StatusName = field.NewString(tableName, "status_name")
	_hCV_view_seat_seat.StatusEnName = field.NewString(tableName, "status_en_name")
	_hCV_view_seat_seat.LIBNO = field.NewInt64(tableName, "LIBNO")
	_hCV_view_seat_seat.LibName = field.NewString(tableName, "lib_name")
	_hCV_view_seat_seat.LibEnName = field.NewString(tableName, "lib_en_name")
	_hCV_view_seat_seat.FLOORNO = field.NewInt64(tableName, "FLOORNO")
	_hCV_view_seat_seat.FloorName = field.NewString(tableName, "floor_name")
	_hCV_view_seat_seat.FloorEnName = field.NewString(tableName, "floor_en_name")
	_hCV_view_seat_seat.Floor = field.NewInt64(tableName, "floor")
	_hCV_view_seat_seat.FLOORIMAGE = field.NewString(tableName, "FLOOR_IMAGE")
	_hCV_view_seat_seat.FLOORIMAGE2 = field.NewString(tableName, "FLOOR_IMAGE2")
	_hCV_view_seat_seat.ROOMNO = field.NewInt64(tableName, "ROOMNO")
	_hCV_view_seat_seat.RoomName = field.NewString(tableName, "room_name")
	_hCV_view_seat_seat.RoomEnName = field.NewString(tableName, "room_en_name")
	_hCV_view_seat_seat.SECTORNO = field.NewInt64(tableName, "SECTORNO")
	_hCV_view_seat_seat.SectorName = field.NewString(tableName, "sector_name")
	_hCV_view_seat_seat.SectorEnName = field.NewString(tableName, "sector_en_name")
	_hCV_view_seat_seat.BOOKINGYN = field.NewString(tableName, "BOOKING_YN")
	_hCV_view_seat_seat.MOBILEBOOKINGYN = field.NewString(tableName, "MOBILE_BOOKING_YN")
	_hCV_view_seat_seat.ASSIGNYN = field.NewString(tableName, "ASSIGN_YN")
	_hCV_view_seat_seat.MOBILEASSIGNYN = field.NewString(tableName, "MOBILE_ASSIGN_YN")
	_hCV_view_seat_seat.DAYFROM = field.NewInt64(tableName, "DAY_FROM")
	_hCV_view_seat_seat.DAYTO = field.NewInt64(tableName, "DAY_TO")
	_hCV_view_seat_seat.MEDIABOOKINGYN = field.NewString(tableName, "MEDIA_BOOKING_YN")
	_hCV_view_seat_seat.EQUIPBOOKINGYN = field.NewString(tableName, "EQUIP_BOOKING_YN")
	_hCV_view_seat_seat.SECTORIMAGE = field.NewString(tableName, "SECTOR_IMAGE")
	_hCV_view_seat_seat.SECTORIMAGE2 = field.NewString(tableName, "SECTOR_IMAGE2")
	_hCV_view_seat_seat.TYPENO = field.NewInt64(tableName, "TYPENO")
	_hCV_view_seat_seat.TypeName = field.NewString(tableName, "type_name")
	_hCV_view_seat_seat.TypeEnName = field.NewString(tableName, "type_en_name")
	_hCV_view_seat_seat.ICONTYPE = field.NewInt64(tableName, "ICONTYPE")
	_hCV_view_seat_seat.POSX = field.NewInt64(tableName, "POSX")
	_hCV_view_seat_seat.POSY = field.NewInt64(tableName, "POSY")
	_hCV_view_seat_seat.POSX2 = field.NewInt64(tableName, "POSX2")
	_hCV_view_seat_seat.POSY2 = field.NewInt64(tableName, "POSY2")
	_hCV_view_seat_seat.POSW = field.NewInt64(tableName, "POSW")
	_hCV_view_seat_seat.POSH = field.NewInt64(tableName, "POSH")
	_hCV_view_seat_seat.GRMIN = field.NewInt64(tableName, "GR_MIN")
	_hCV_view_seat_seat.GRMAX = field.NewInt64(tableName, "GR_MAX")
	_hCV_view_seat_seat.APPROVALYN = field.NewString(tableName, "APPROVAL_YN")
	_hCV_view_seat_seat.USEMIN = field.NewInt64(tableName, "USE_MIN")
	_hCV_view_seat_seat.CANCONTMIN = field.NewInt64(tableName, "CAN_CONT_MIN")
	_hCV_view_seat_seat.CONTMIN = field.NewInt64(tableName, "CONT_MIN")
	_hCV_view_seat_seat.MAXCONTCNT = field.NewInt64(tableName, "MAX_CONT_CNT")
	_hCV_view_seat_seat.NOSHOWYN = field.NewString(tableName, "NOSHOW_YN")
	_hCV_view_seat_seat.NOSHOWMIN = field.NewInt64(tableName, "NOSHOW_MIN")

	_hCV_view_seat_seat.fillFieldMap()

	return _hCV_view_seat_seat
}

type hCV_view_seat_seat struct {
	hCV_view_seat_seatDo

	ALL             field.Asterisk
	SEATNO          field.Int64
	Name            field.String
	EnName          field.String
	Vname           field.String
	Status          field.Int64
	StatusName      field.String
	StatusEnName    field.String
	LIBNO           field.Int64
	LibName         field.String
	LibEnName       field.String
	FLOORNO         field.Int64
	FloorName       field.String
	FloorEnName     field.String
	Floor           field.Int64
	FLOORIMAGE      field.String
	FLOORIMAGE2     field.String
	ROOMNO          field.Int64
	RoomName        field.String
	RoomEnName      field.String
	SECTORNO        field.Int64
	SectorName      field.String
	SectorEnName    field.String
	BOOKINGYN       field.String
	MOBILEBOOKINGYN field.String
	ASSIGNYN        field.String
	MOBILEASSIGNYN  field.String
	DAYFROM         field.Int64
	DAYTO           field.Int64
	MEDIABOOKINGYN  field.String
	EQUIPBOOKINGYN  field.String
	SECTORIMAGE     field.String
	SECTORIMAGE2    field.String
	TYPENO          field.Int64
	TypeName        field.String
	TypeEnName      field.String
	ICONTYPE        field.Int64
	POSX            field.Int64
	POSY            field.Int64
	POSX2           field.Int64
	POSY2           field.Int64
	POSW            field.Int64
	POSH            field.Int64
	GRMIN           field.Int64
	GRMAX           field.Int64
	APPROVALYN      field.String
	USEMIN          field.Int64
	CANCONTMIN      field.Int64
	CONTMIN         field.Int64
	MAXCONTCNT      field.Int64
	NOSHOWYN        field.String
	NOSHOWMIN       field.Int64

	fieldMap map[string]field.Expr
}

func (h hCV_view_seat_seat) Table(newTableName string) *hCV_view_seat_seat {
	h.hCV_view_seat_seatDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_view_seat_seat) As(alias string) *hCV_view_seat_seat {
	h.hCV_view_seat_seatDo.DO = *(h.hCV_view_seat_seatDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_view_seat_seat) updateTableName(table string) *hCV_view_seat_seat {
	h.ALL = field.NewAsterisk(table)
	h.SEATNO = field.NewInt64(table, "SEATNO")
	h.Name = field.NewString(table, "name")
	h.EnName = field.NewString(table, "en_name")
	h.Vname = field.NewString(table, "vname")
	h.Status = field.NewInt64(table, "status")
	h.StatusName = field.NewString(table, "status_name")
	h.StatusEnName = field.NewString(table, "status_en_name")
	h.LIBNO = field.NewInt64(table, "LIBNO")
	h.LibName = field.NewString(table, "lib_name")
	h.LibEnName = field.NewString(table, "lib_en_name")
	h.FLOORNO = field.NewInt64(table, "FLOORNO")
	h.FloorName = field.NewString(table, "floor_name")
	h.FloorEnName = field.NewString(table, "floor_en_name")
	h.Floor = field.NewInt64(table, "floor")
	h.FLOORIMAGE = field.NewString(table, "FLOOR_IMAGE")
	h.FLOORIMAGE2 = field.NewString(table, "FLOOR_IMAGE2")
	h.ROOMNO = field.NewInt64(table, "ROOMNO")
	h.RoomName = field.NewString(table, "room_name")
	h.RoomEnName = field.NewString(table, "room_en_name")
	h.SECTORNO = field.NewInt64(table, "SECTORNO")
	h.SectorName = field.NewString(table, "sector_name")
	h.SectorEnName = field.NewString(table, "sector_en_name")
	h.BOOKINGYN = field.NewString(table, "BOOKING_YN")
	h.MOBILEBOOKINGYN = field.NewString(table, "MOBILE_BOOKING_YN")
	h.ASSIGNYN = field.NewString(table, "ASSIGN_YN")
	h.MOBILEASSIGNYN = field.NewString(table, "MOBILE_ASSIGN_YN")
	h.DAYFROM = field.NewInt64(table, "DAY_FROM")
	h.DAYTO = field.NewInt64(table, "DAY_TO")
	h.MEDIABOOKINGYN = field.NewString(table, "MEDIA_BOOKING_YN")
	h.EQUIPBOOKINGYN = field.NewString(table, "EQUIP_BOOKING_YN")
	h.SECTORIMAGE = field.NewString(table, "SECTOR_IMAGE")
	h.SECTORIMAGE2 = field.NewString(table, "SECTOR_IMAGE2")
	h.TYPENO = field.NewInt64(table, "TYPENO")
	h.TypeName = field.NewString(table, "type_name")
	h.TypeEnName = field.NewString(table, "type_en_name")
	h.ICONTYPE = field.NewInt64(table, "ICONTYPE")
	h.POSX = field.NewInt64(table, "POSX")
	h.POSY = field.NewInt64(table, "POSY")
	h.POSX2 = field.NewInt64(table, "POSX2")
	h.POSY2 = field.NewInt64(table, "POSY2")
	h.POSW = field.NewInt64(table, "POSW")
	h.POSH = field.NewInt64(table, "POSH")
	h.GRMIN = field.NewInt64(table, "GR_MIN")
	h.GRMAX = field.NewInt64(table, "GR_MAX")
	h.APPROVALYN = field.NewString(table, "APPROVAL_YN")
	h.USEMIN = field.NewInt64(table, "USE_MIN")
	h.CANCONTMIN = field.NewInt64(table, "CAN_CONT_MIN")
	h.CONTMIN = field.NewInt64(table, "CONT_MIN")
	h.MAXCONTCNT = field.NewInt64(table, "MAX_CONT_CNT")
	h.NOSHOWYN = field.NewString(table, "NOSHOW_YN")
	h.NOSHOWMIN = field.NewInt64(table, "NOSHOW_MIN")

	h.fillFieldMap()

	return h
}

func (h *hCV_view_seat_seat) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_view_seat_seat) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 51)
	h.fieldMap["SEATNO"] = h.SEATNO
	h.fieldMap["name"] = h.Name
	h.fieldMap["en_name"] = h.EnName
	h.fieldMap["vname"] = h.Vname
	h.fieldMap["status"] = h.Status
	h.fieldMap["status_name"] = h.StatusName
	h.fieldMap["status_en_name"] = h.StatusEnName
	h.fieldMap["LIBNO"] = h.LIBNO
	h.fieldMap["lib_name"] = h.LibName
	h.fieldMap["lib_en_name"] = h.LibEnName
	h.fieldMap["FLOORNO"] = h.FLOORNO
	h.fieldMap["floor_name"] = h.FloorName
	h.fieldMap["floor_en_name"] = h.FloorEnName
	h.fieldMap["floor"] = h.Floor
	h.fieldMap["FLOOR_IMAGE"] = h.FLOORIMAGE
	h.fieldMap["FLOOR_IMAGE2"] = h.FLOORIMAGE2
	h.fieldMap["ROOMNO"] = h.ROOMNO
	h.fieldMap["room_name"] = h.RoomName
	h.fieldMap["room_en_name"] = h.RoomEnName
	h.fieldMap["SECTORNO"] = h.SECTORNO
	h.fieldMap["sector_name"] = h.SectorName
	h.fieldMap["sector_en_name"] = h.SectorEnName
	h.fieldMap["BOOKING_YN"] = h.BOOKINGYN
	h.fieldMap["MOBILE_BOOKING_YN"] = h.MOBILEBOOKINGYN
	h.fieldMap["ASSIGN_YN"] = h.ASSIGNYN
	h.fieldMap["MOBILE_ASSIGN_YN"] = h.MOBILEASSIGNYN
	h.fieldMap["DAY_FROM"] = h.DAYFROM
	h.fieldMap["DAY_TO"] = h.DAYTO
	h.fieldMap["MEDIA_BOOKING_YN"] = h.MEDIABOOKINGYN
	h.fieldMap["EQUIP_BOOKING_YN"] = h.EQUIPBOOKINGYN
	h.fieldMap["SECTOR_IMAGE"] = h.SECTORIMAGE
	h.fieldMap["SECTOR_IMAGE2"] = h.SECTORIMAGE2
	h.fieldMap["TYPENO"] = h.TYPENO
	h.fieldMap["type_name"] = h.TypeName
	h.fieldMap["type_en_name"] = h.TypeEnName
	h.fieldMap["ICONTYPE"] = h.ICONTYPE
	h.fieldMap["POSX"] = h.POSX
	h.fieldMap["POSY"] = h.POSY
	h.fieldMap["POSX2"] = h.POSX2
	h.fieldMap["POSY2"] = h.POSY2
	h.fieldMap["POSW"] = h.POSW
	h.fieldMap["POSH"] = h.POSH
	h.fieldMap["GR_MIN"] = h.GRMIN
	h.fieldMap["GR_MAX"] = h.GRMAX
	h.fieldMap["APPROVAL_YN"] = h.APPROVALYN
	h.fieldMap["USE_MIN"] = h.USEMIN
	h.fieldMap["CAN_CONT_MIN"] = h.CANCONTMIN
	h.fieldMap["CONT_MIN"] = h.CONTMIN
	h.fieldMap["MAX_CONT_CNT"] = h.MAXCONTCNT
	h.fieldMap["NOSHOW_YN"] = h.NOSHOWYN
	h.fieldMap["NOSHOW_MIN"] = h.NOSHOWMIN
}

func (h hCV_view_seat_seat) clone(db *gorm.DB) hCV_view_seat_seat {
	h.hCV_view_seat_seatDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_view_seat_seat) replaceDB(db *gorm.DB) hCV_view_seat_seat {
	h.hCV_view_seat_seatDo.ReplaceDB(db)
	return h
}

type hCV_view_seat_seatDo struct{ gen.DO }

type IHCV_view_seat_seatDo interface {
	gen.SubQuery
	Debug() IHCV_view_seat_seatDo
	WithContext(ctx context.Context) IHCV_view_seat_seatDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_view_seat_seatDo
	WriteDB() IHCV_view_seat_seatDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_view_seat_seatDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_view_seat_seatDo
	Not(conds ...gen.Condition) IHCV_view_seat_seatDo
	Or(conds ...gen.Condition) IHCV_view_seat_seatDo
	Select(conds ...field.Expr) IHCV_view_seat_seatDo
	Where(conds ...gen.Condition) IHCV_view_seat_seatDo
	Order(conds ...field.Expr) IHCV_view_seat_seatDo
	Distinct(cols ...field.Expr) IHCV_view_seat_seatDo
	Omit(cols ...field.Expr) IHCV_view_seat_seatDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_view_seat_seatDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_seatDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_seatDo
	Group(cols ...field.Expr) IHCV_view_seat_seatDo
	Having(conds ...gen.Condition) IHCV_view_seat_seatDo
	Limit(limit int) IHCV_view_seat_seatDo
	Offset(offset int) IHCV_view_seat_seatDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_seat_seatDo
	Unscoped() IHCV_view_seat_seatDo
	Create(values ...*model.HCV_view_seat_seat) error
	CreateInBatches(values []*model.HCV_view_seat_seat, batchSize int) error
	Save(values ...*model.HCV_view_seat_seat) error
	First() (*model.HCV_view_seat_seat, error)
	Take() (*model.HCV_view_seat_seat, error)
	Last() (*model.HCV_view_seat_seat, error)
	Find() ([]*model.HCV_view_seat_seat, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_seat_seat, err error)
	FindInBatches(result *[]*model.HCV_view_seat_seat, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_view_seat_seat) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_view_seat_seatDo
	Assign(attrs ...field.AssignExpr) IHCV_view_seat_seatDo
	Joins(fields ...field.RelationField) IHCV_view_seat_seatDo
	Preload(fields ...field.RelationField) IHCV_view_seat_seatDo
	FirstOrInit() (*model.HCV_view_seat_seat, error)
	FirstOrCreate() (*model.HCV_view_seat_seat, error)
	FindByPage(offset int, limit int) (result []*model.HCV_view_seat_seat, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_view_seat_seatDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_view_seat_seatDo) Debug() IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_view_seat_seatDo) WithContext(ctx context.Context) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_view_seat_seatDo) ReadDB() IHCV_view_seat_seatDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_view_seat_seatDo) WriteDB() IHCV_view_seat_seatDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_view_seat_seatDo) Session(config *gorm.Session) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_view_seat_seatDo) Clauses(conds ...clause.Expression) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_view_seat_seatDo) Returning(value interface{}, columns ...string) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_view_seat_seatDo) Not(conds ...gen.Condition) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_view_seat_seatDo) Or(conds ...gen.Condition) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_view_seat_seatDo) Select(conds ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_view_seat_seatDo) Where(conds ...gen.Condition) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_view_seat_seatDo) Order(conds ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_view_seat_seatDo) Distinct(cols ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_view_seat_seatDo) Omit(cols ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_view_seat_seatDo) Join(table schema.Tabler, on ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_view_seat_seatDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_view_seat_seatDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_view_seat_seatDo) Group(cols ...field.Expr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_view_seat_seatDo) Having(conds ...gen.Condition) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_view_seat_seatDo) Limit(limit int) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_view_seat_seatDo) Offset(offset int) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_view_seat_seatDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_view_seat_seatDo) Unscoped() IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_view_seat_seatDo) Create(values ...*model.HCV_view_seat_seat) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_view_seat_seatDo) CreateInBatches(values []*model.HCV_view_seat_seat, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_view_seat_seatDo) Save(values ...*model.HCV_view_seat_seat) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_view_seat_seatDo) First() (*model.HCV_view_seat_seat, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_seat), nil
	}
}

func (h hCV_view_seat_seatDo) Take() (*model.HCV_view_seat_seat, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_seat), nil
	}
}

func (h hCV_view_seat_seatDo) Last() (*model.HCV_view_seat_seat, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_seat), nil
	}
}

func (h hCV_view_seat_seatDo) Find() ([]*model.HCV_view_seat_seat, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_view_seat_seat), err
}

func (h hCV_view_seat_seatDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_seat_seat, err error) {
	buf := make([]*model.HCV_view_seat_seat, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_view_seat_seatDo) FindInBatches(result *[]*model.HCV_view_seat_seat, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_view_seat_seatDo) Attrs(attrs ...field.AssignExpr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_view_seat_seatDo) Assign(attrs ...field.AssignExpr) IHCV_view_seat_seatDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_view_seat_seatDo) Joins(fields ...field.RelationField) IHCV_view_seat_seatDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_view_seat_seatDo) Preload(fields ...field.RelationField) IHCV_view_seat_seatDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_view_seat_seatDo) FirstOrInit() (*model.HCV_view_seat_seat, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_seat), nil
	}
}

func (h hCV_view_seat_seatDo) FirstOrCreate() (*model.HCV_view_seat_seat, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_seat), nil
	}
}

func (h hCV_view_seat_seatDo) FindByPage(offset int, limit int) (result []*model.HCV_view_seat_seat, count int64, err error) {
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

func (h hCV_view_seat_seatDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_view_seat_seatDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_view_seat_seatDo) Delete(models ...*model.HCV_view_seat_seat) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_view_seat_seatDo) withDO(do gen.Dao) *hCV_view_seat_seatDo {
	h.DO = *do.(*gen.DO)
	return h
}
