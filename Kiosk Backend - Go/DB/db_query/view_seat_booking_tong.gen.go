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

func newHCV_view_seat_booking_tong(db *gorm.DB, opts ...gen.DOOption) hCV_view_seat_booking_tong {
	_hCV_view_seat_booking_tong := hCV_view_seat_booking_tong{}

	_hCV_view_seat_booking_tong.hCV_view_seat_booking_tongDo.UseDB(db, opts...)
	_hCV_view_seat_booking_tong.hCV_view_seat_booking_tongDo.UseModel(&model.HCV_view_seat_booking_tong{})

	tableName := _hCV_view_seat_booking_tong.hCV_view_seat_booking_tongDo.TableName()
	_hCV_view_seat_booking_tong.ALL = field.NewAsterisk(tableName)
	_hCV_view_seat_booking_tong.BSEQNO = field.NewInt64(tableName, "BSEQNO")
	_hCV_view_seat_booking_tong.STATUS = field.NewInt64(tableName, "STATUS")
	_hCV_view_seat_booking_tong.StatusName = field.NewString(tableName, "status_name")
	_hCV_view_seat_booking_tong.StatusEnName = field.NewString(tableName, "status_en_name")
	_hCV_view_seat_booking_tong.SCHOOLNO = field.NewString(tableName, "SCHOOLNO")
	_hCV_view_seat_booking_tong.USERID = field.NewString(tableName, "USERID")
	_hCV_view_seat_booking_tong.SEATNO = field.NewInt64(tableName, "SEATNO")
	_hCV_view_seat_booking_tong.SeatName = field.NewString(tableName, "seat_name")
	_hCV_view_seat_booking_tong.SeatEnName = field.NewString(tableName, "seat_en_name")
	_hCV_view_seat_booking_tong.SeatStatus = field.NewInt64(tableName, "seat_status")
	_hCV_view_seat_booking_tong.SeatStatusName = field.NewString(tableName, "seat_status_name")
	_hCV_view_seat_booking_tong.SeatStatusEnName = field.NewString(tableName, "seat_status_en_name")
	_hCV_view_seat_booking_tong.SeatVname = field.NewString(tableName, "seat_vname")
	_hCV_view_seat_booking_tong.ISCANCELED = field.NewInt64(tableName, "ISCANCELED")
	_hCV_view_seat_booking_tong.ISSUEFROM = field.NewInt64(tableName, "ISSUEFROM")
	_hCV_view_seat_booking_tong.ISSUEDETAIL = field.NewString(tableName, "ISSUEDETAIL")
	_hCV_view_seat_booking_tong.ISSUETYPENO = field.NewInt64(tableName, "ISSUE_TYPE_NO")
	_hCV_view_seat_booking_tong.USESTART = field.NewTime(tableName, "USESTART")
	_hCV_view_seat_booking_tong.USEEXPIRE = field.NewTime(tableName, "USEEXPIRE")
	_hCV_view_seat_booking_tong.EXPIREREASON = field.NewInt64(tableName, "EXPIREREASON")
	_hCV_view_seat_booking_tong.ExpirereasonName = field.NewString(tableName, "expirereason_name")
	_hCV_view_seat_booking_tong.ExpirereasonEnName = field.NewString(tableName, "expirereason_en_name")
	_hCV_view_seat_booking_tong.PRINTCNT = field.NewInt64(tableName, "PRINTCNT")
	_hCV_view_seat_booking_tong.REGTIME = field.NewTime(tableName, "REGTIME")
	_hCV_view_seat_booking_tong.EXTENDMIN = field.NewInt64(tableName, "EXTEND_MIN")
	_hCV_view_seat_booking_tong.EXTENDCNT = field.NewInt64(tableName, "EXTEND_CNT")
	_hCV_view_seat_booking_tong.STARTTIME = field.NewTime(tableName, "STARTTIME")
	_hCV_view_seat_booking_tong.RETURNTIME = field.NewTime(tableName, "RETURNTIME")
	_hCV_view_seat_booking_tong.USERNAME = field.NewString(tableName, "USER_NAME")
	_hCV_view_seat_booking_tong.USERPOSITION = field.NewString(tableName, "USER_POSITION")
	_hCV_view_seat_booking_tong.COMPANYCODE = field.NewString(tableName, "COMPANY_CODE")
	_hCV_view_seat_booking_tong.OrgDeptCode = field.NewString(tableName, "org_dept_code")
	_hCV_view_seat_booking_tong.GATEOUTTIME = field.NewTime(tableName, "GATEOUTTIME")
	_hCV_view_seat_booking_tong.ISADMINBOOKING = field.NewInt64(tableName, "ISADMINBOOKING")
	_hCV_view_seat_booking_tong.CHECKGATEINTIME = field.NewTime(tableName, "CHECK_GATEINTIME")
	_hCV_view_seat_booking_tong.CHECKGATEIN = field.NewInt64(tableName, "CHECK_GATEIN")
	_hCV_view_seat_booking_tong.MOBILE = field.NewString(tableName, "MOBILE")
	_hCV_view_seat_booking_tong.PCLOGINCHECKSTATUS = field.NewInt64(tableName, "PCLOGIN_CHECK_STATUS")
	_hCV_view_seat_booking_tong.EXTENDMSGSTATUS = field.NewInt64(tableName, "EXTEND_MSG_STATUS")
	_hCV_view_seat_booking_tong.BEFORESEATRETURNSTATUS = field.NewInt64(tableName, "BEFORE_SEAT_RETURN_STATUS")
	_hCV_view_seat_booking_tong.TYPENO = field.NewInt64(tableName, "TYPENO")
	_hCV_view_seat_booking_tong.SeatTypeName = field.NewString(tableName, "seat_type_name")
	_hCV_view_seat_booking_tong.SeatTypeEnName = field.NewString(tableName, "seat_type_en_name")
	_hCV_view_seat_booking_tong.LIBNO = field.NewInt64(tableName, "LIBNO")
	_hCV_view_seat_booking_tong.LibName = field.NewString(tableName, "lib_name")
	_hCV_view_seat_booking_tong.LibEnName = field.NewString(tableName, "lib_en_name")
	_hCV_view_seat_booking_tong.FLOORNO = field.NewInt64(tableName, "FLOORNO")
	_hCV_view_seat_booking_tong.FloorName = field.NewString(tableName, "floor_name")
	_hCV_view_seat_booking_tong.FloorEnName = field.NewString(tableName, "floor_en_name")
	_hCV_view_seat_booking_tong.FLOOR = field.NewInt64(tableName, "FLOOR")
	_hCV_view_seat_booking_tong.ROOMNO = field.NewInt64(tableName, "ROOMNO")
	_hCV_view_seat_booking_tong.RoomName = field.NewString(tableName, "room_name")
	_hCV_view_seat_booking_tong.RoomEnName = field.NewString(tableName, "room_en_name")
	_hCV_view_seat_booking_tong.SECTORNO = field.NewInt64(tableName, "SECTORNO")
	_hCV_view_seat_booking_tong.SectorName = field.NewString(tableName, "sector_name")
	_hCV_view_seat_booking_tong.SectorEnName = field.NewString(tableName, "sector_en_name")
	_hCV_view_seat_booking_tong.BOOKINGYN = field.NewString(tableName, "BOOKING_YN")
	_hCV_view_seat_booking_tong.MOBILEBOOKINGYN = field.NewString(tableName, "MOBILE_BOOKING_YN")
	_hCV_view_seat_booking_tong.ASSIGNYN = field.NewString(tableName, "ASSIGN_YN")
	_hCV_view_seat_booking_tong.MOBILEASSIGNYN = field.NewString(tableName, "MOBILE_ASSIGN_YN")
	_hCV_view_seat_booking_tong.DAYFROM = field.NewInt64(tableName, "DAY_FROM")
	_hCV_view_seat_booking_tong.DAYTO = field.NewInt64(tableName, "DAY_TO")
	_hCV_view_seat_booking_tong.MEDIABOOKINGYN = field.NewString(tableName, "MEDIA_BOOKING_YN")
	_hCV_view_seat_booking_tong.EQUIPBOOKINGYN = field.NewString(tableName, "EQUIP_BOOKING_YN")
	_hCV_view_seat_booking_tong.APPROVALYN = field.NewString(tableName, "APPROVAL_YN")
	_hCV_view_seat_booking_tong.USEMIN = field.NewInt64(tableName, "USE_MIN")
	_hCV_view_seat_booking_tong.CANCONTMIN = field.NewInt64(tableName, "CAN_CONT_MIN")
	_hCV_view_seat_booking_tong.CONTMIN = field.NewInt64(tableName, "CONT_MIN")
	_hCV_view_seat_booking_tong.MAXCONTCNT = field.NewInt64(tableName, "MAX_CONT_CNT")
	_hCV_view_seat_booking_tong.NOSHOWYN = field.NewString(tableName, "NOSHOW_YN")
	_hCV_view_seat_booking_tong.NOSHOWMIN = field.NewInt64(tableName, "NOSHOW_MIN")
	_hCV_view_seat_booking_tong.MEMBERS = field.NewString(tableName, "MEMBERS")
	_hCV_view_seat_booking_tong.PURPOSE = field.NewString(tableName, "PURPOSE")
	_hCV_view_seat_booking_tong.MEMBERCNT = field.NewInt64(tableName, "MEMBERCNT")
	_hCV_view_seat_booking_tong.EMAIL = field.NewString(tableName, "EMAIL")
	_hCV_view_seat_booking_tong.TEL = field.NewString(tableName, "TEL")
	_hCV_view_seat_booking_tong.KEYSTATUS = field.NewInt64(tableName, "KEY_STATUS")
	_hCV_view_seat_booking_tong.GRMIN = field.NewInt64(tableName, "GR_MIN")
	_hCV_view_seat_booking_tong.GRMAX = field.NewInt64(tableName, "GR_MAX")
	_hCV_view_seat_booking_tong.UseApproval = field.NewInt64(tableName, "use_approval")
	_hCV_view_seat_booking_tong.Reason = field.NewString(tableName, "reason")
	_hCV_view_seat_booking_tong.CSCHOOLNO = field.NewString(tableName, "CSCHOOLNO")
	_hCV_view_seat_booking_tong.CBIGO = field.NewString(tableName, "CBIGO")
	_hCV_view_seat_booking_tong.AltPid = field.NewString(tableName, "alt_pid")
	_hCV_view_seat_booking_tong.Username = field.NewString(tableName, "username")
	_hCV_view_seat_booking_tong.PatType = field.NewString(tableName, "pat_type")
	_hCV_view_seat_booking_tong.PatName = field.NewString(tableName, "pat_name")
	_hCV_view_seat_booking_tong.DeptCode = field.NewString(tableName, "dept_code")
	_hCV_view_seat_booking_tong.DeptName = field.NewString(tableName, "dept_name")
	_hCV_view_seat_booking_tong.CollegeCode = field.NewString(tableName, "college_code")
	_hCV_view_seat_booking_tong.CollegeName = field.NewString(tableName, "college_name")

	_hCV_view_seat_booking_tong.fillFieldMap()

	return _hCV_view_seat_booking_tong
}

type hCV_view_seat_booking_tong struct {
	hCV_view_seat_booking_tongDo

	ALL                    field.Asterisk
	BSEQNO                 field.Int64
	STATUS                 field.Int64
	StatusName             field.String
	StatusEnName           field.String
	SCHOOLNO               field.String
	USERID                 field.String
	SEATNO                 field.Int64
	SeatName               field.String
	SeatEnName             field.String
	SeatStatus             field.Int64
	SeatStatusName         field.String
	SeatStatusEnName       field.String
	SeatVname              field.String
	ISCANCELED             field.Int64
	ISSUEFROM              field.Int64
	ISSUEDETAIL            field.String
	ISSUETYPENO            field.Int64
	USESTART               field.Time
	USEEXPIRE              field.Time
	EXPIREREASON           field.Int64
	ExpirereasonName       field.String
	ExpirereasonEnName     field.String
	PRINTCNT               field.Int64
	REGTIME                field.Time
	EXTENDMIN              field.Int64
	EXTENDCNT              field.Int64
	STARTTIME              field.Time
	RETURNTIME             field.Time
	USERNAME               field.String
	USERPOSITION           field.String
	COMPANYCODE            field.String
	OrgDeptCode            field.String
	GATEOUTTIME            field.Time
	ISADMINBOOKING         field.Int64
	CHECKGATEINTIME        field.Time
	CHECKGATEIN            field.Int64
	MOBILE                 field.String
	PCLOGINCHECKSTATUS     field.Int64
	EXTENDMSGSTATUS        field.Int64
	BEFORESEATRETURNSTATUS field.Int64
	TYPENO                 field.Int64
	SeatTypeName           field.String
	SeatTypeEnName         field.String
	LIBNO                  field.Int64
	LibName                field.String
	LibEnName              field.String
	FLOORNO                field.Int64
	FloorName              field.String
	FloorEnName            field.String
	FLOOR                  field.Int64
	ROOMNO                 field.Int64
	RoomName               field.String
	RoomEnName             field.String
	SECTORNO               field.Int64
	SectorName             field.String
	SectorEnName           field.String
	BOOKINGYN              field.String
	MOBILEBOOKINGYN        field.String
	ASSIGNYN               field.String
	MOBILEASSIGNYN         field.String
	DAYFROM                field.Int64
	DAYTO                  field.Int64
	MEDIABOOKINGYN         field.String
	EQUIPBOOKINGYN         field.String
	APPROVALYN             field.String
	USEMIN                 field.Int64
	CANCONTMIN             field.Int64
	CONTMIN                field.Int64
	MAXCONTCNT             field.Int64
	NOSHOWYN               field.String
	NOSHOWMIN              field.Int64
	MEMBERS                field.String
	PURPOSE                field.String
	MEMBERCNT              field.Int64
	EMAIL                  field.String
	TEL                    field.String
	KEYSTATUS              field.Int64
	GRMIN                  field.Int64
	GRMAX                  field.Int64
	UseApproval            field.Int64
	Reason                 field.String
	CSCHOOLNO              field.String
	CBIGO                  field.String
	AltPid                 field.String
	Username               field.String
	PatType                field.String
	PatName                field.String
	DeptCode               field.String
	DeptName               field.String
	CollegeCode            field.String
	CollegeName            field.String

	fieldMap map[string]field.Expr
}

func (h hCV_view_seat_booking_tong) Table(newTableName string) *hCV_view_seat_booking_tong {
	h.hCV_view_seat_booking_tongDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hCV_view_seat_booking_tong) As(alias string) *hCV_view_seat_booking_tong {
	h.hCV_view_seat_booking_tongDo.DO = *(h.hCV_view_seat_booking_tongDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hCV_view_seat_booking_tong) updateTableName(table string) *hCV_view_seat_booking_tong {
	h.ALL = field.NewAsterisk(table)
	h.BSEQNO = field.NewInt64(table, "BSEQNO")
	h.STATUS = field.NewInt64(table, "STATUS")
	h.StatusName = field.NewString(table, "status_name")
	h.StatusEnName = field.NewString(table, "status_en_name")
	h.SCHOOLNO = field.NewString(table, "SCHOOLNO")
	h.USERID = field.NewString(table, "USERID")
	h.SEATNO = field.NewInt64(table, "SEATNO")
	h.SeatName = field.NewString(table, "seat_name")
	h.SeatEnName = field.NewString(table, "seat_en_name")
	h.SeatStatus = field.NewInt64(table, "seat_status")
	h.SeatStatusName = field.NewString(table, "seat_status_name")
	h.SeatStatusEnName = field.NewString(table, "seat_status_en_name")
	h.SeatVname = field.NewString(table, "seat_vname")
	h.ISCANCELED = field.NewInt64(table, "ISCANCELED")
	h.ISSUEFROM = field.NewInt64(table, "ISSUEFROM")
	h.ISSUEDETAIL = field.NewString(table, "ISSUEDETAIL")
	h.ISSUETYPENO = field.NewInt64(table, "ISSUE_TYPE_NO")
	h.USESTART = field.NewTime(table, "USESTART")
	h.USEEXPIRE = field.NewTime(table, "USEEXPIRE")
	h.EXPIREREASON = field.NewInt64(table, "EXPIREREASON")
	h.ExpirereasonName = field.NewString(table, "expirereason_name")
	h.ExpirereasonEnName = field.NewString(table, "expirereason_en_name")
	h.PRINTCNT = field.NewInt64(table, "PRINTCNT")
	h.REGTIME = field.NewTime(table, "REGTIME")
	h.EXTENDMIN = field.NewInt64(table, "EXTEND_MIN")
	h.EXTENDCNT = field.NewInt64(table, "EXTEND_CNT")
	h.STARTTIME = field.NewTime(table, "STARTTIME")
	h.RETURNTIME = field.NewTime(table, "RETURNTIME")
	h.USERNAME = field.NewString(table, "USER_NAME")
	h.USERPOSITION = field.NewString(table, "USER_POSITION")
	h.COMPANYCODE = field.NewString(table, "COMPANY_CODE")
	h.OrgDeptCode = field.NewString(table, "org_dept_code")
	h.GATEOUTTIME = field.NewTime(table, "GATEOUTTIME")
	h.ISADMINBOOKING = field.NewInt64(table, "ISADMINBOOKING")
	h.CHECKGATEINTIME = field.NewTime(table, "CHECK_GATEINTIME")
	h.CHECKGATEIN = field.NewInt64(table, "CHECK_GATEIN")
	h.MOBILE = field.NewString(table, "MOBILE")
	h.PCLOGINCHECKSTATUS = field.NewInt64(table, "PCLOGIN_CHECK_STATUS")
	h.EXTENDMSGSTATUS = field.NewInt64(table, "EXTEND_MSG_STATUS")
	h.BEFORESEATRETURNSTATUS = field.NewInt64(table, "BEFORE_SEAT_RETURN_STATUS")
	h.TYPENO = field.NewInt64(table, "TYPENO")
	h.SeatTypeName = field.NewString(table, "seat_type_name")
	h.SeatTypeEnName = field.NewString(table, "seat_type_en_name")
	h.LIBNO = field.NewInt64(table, "LIBNO")
	h.LibName = field.NewString(table, "lib_name")
	h.LibEnName = field.NewString(table, "lib_en_name")
	h.FLOORNO = field.NewInt64(table, "FLOORNO")
	h.FloorName = field.NewString(table, "floor_name")
	h.FloorEnName = field.NewString(table, "floor_en_name")
	h.FLOOR = field.NewInt64(table, "FLOOR")
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
	h.APPROVALYN = field.NewString(table, "APPROVAL_YN")
	h.USEMIN = field.NewInt64(table, "USE_MIN")
	h.CANCONTMIN = field.NewInt64(table, "CAN_CONT_MIN")
	h.CONTMIN = field.NewInt64(table, "CONT_MIN")
	h.MAXCONTCNT = field.NewInt64(table, "MAX_CONT_CNT")
	h.NOSHOWYN = field.NewString(table, "NOSHOW_YN")
	h.NOSHOWMIN = field.NewInt64(table, "NOSHOW_MIN")
	h.MEMBERS = field.NewString(table, "MEMBERS")
	h.PURPOSE = field.NewString(table, "PURPOSE")
	h.MEMBERCNT = field.NewInt64(table, "MEMBERCNT")
	h.EMAIL = field.NewString(table, "EMAIL")
	h.TEL = field.NewString(table, "TEL")
	h.KEYSTATUS = field.NewInt64(table, "KEY_STATUS")
	h.GRMIN = field.NewInt64(table, "GR_MIN")
	h.GRMAX = field.NewInt64(table, "GR_MAX")
	h.UseApproval = field.NewInt64(table, "use_approval")
	h.Reason = field.NewString(table, "reason")
	h.CSCHOOLNO = field.NewString(table, "CSCHOOLNO")
	h.CBIGO = field.NewString(table, "CBIGO")
	h.AltPid = field.NewString(table, "alt_pid")
	h.Username = field.NewString(table, "username")
	h.PatType = field.NewString(table, "pat_type")
	h.PatName = field.NewString(table, "pat_name")
	h.DeptCode = field.NewString(table, "dept_code")
	h.DeptName = field.NewString(table, "dept_name")
	h.CollegeCode = field.NewString(table, "college_code")
	h.CollegeName = field.NewString(table, "college_name")

	h.fillFieldMap()

	return h
}

func (h *hCV_view_seat_booking_tong) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hCV_view_seat_booking_tong) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 91)
	h.fieldMap["BSEQNO"] = h.BSEQNO
	h.fieldMap["STATUS"] = h.STATUS
	h.fieldMap["status_name"] = h.StatusName
	h.fieldMap["status_en_name"] = h.StatusEnName
	h.fieldMap["SCHOOLNO"] = h.SCHOOLNO
	h.fieldMap["USERID"] = h.USERID
	h.fieldMap["SEATNO"] = h.SEATNO
	h.fieldMap["seat_name"] = h.SeatName
	h.fieldMap["seat_en_name"] = h.SeatEnName
	h.fieldMap["seat_status"] = h.SeatStatus
	h.fieldMap["seat_status_name"] = h.SeatStatusName
	h.fieldMap["seat_status_en_name"] = h.SeatStatusEnName
	h.fieldMap["seat_vname"] = h.SeatVname
	h.fieldMap["ISCANCELED"] = h.ISCANCELED
	h.fieldMap["ISSUEFROM"] = h.ISSUEFROM
	h.fieldMap["ISSUEDETAIL"] = h.ISSUEDETAIL
	h.fieldMap["ISSUE_TYPE_NO"] = h.ISSUETYPENO
	h.fieldMap["USESTART"] = h.USESTART
	h.fieldMap["USEEXPIRE"] = h.USEEXPIRE
	h.fieldMap["EXPIREREASON"] = h.EXPIREREASON
	h.fieldMap["expirereason_name"] = h.ExpirereasonName
	h.fieldMap["expirereason_en_name"] = h.ExpirereasonEnName
	h.fieldMap["PRINTCNT"] = h.PRINTCNT
	h.fieldMap["REGTIME"] = h.REGTIME
	h.fieldMap["EXTEND_MIN"] = h.EXTENDMIN
	h.fieldMap["EXTEND_CNT"] = h.EXTENDCNT
	h.fieldMap["STARTTIME"] = h.STARTTIME
	h.fieldMap["RETURNTIME"] = h.RETURNTIME
	h.fieldMap["USER_NAME"] = h.USERNAME
	h.fieldMap["USER_POSITION"] = h.USERPOSITION
	h.fieldMap["COMPANY_CODE"] = h.COMPANYCODE
	h.fieldMap["org_dept_code"] = h.OrgDeptCode
	h.fieldMap["GATEOUTTIME"] = h.GATEOUTTIME
	h.fieldMap["ISADMINBOOKING"] = h.ISADMINBOOKING
	h.fieldMap["CHECK_GATEINTIME"] = h.CHECKGATEINTIME
	h.fieldMap["CHECK_GATEIN"] = h.CHECKGATEIN
	h.fieldMap["MOBILE"] = h.MOBILE
	h.fieldMap["PCLOGIN_CHECK_STATUS"] = h.PCLOGINCHECKSTATUS
	h.fieldMap["EXTEND_MSG_STATUS"] = h.EXTENDMSGSTATUS
	h.fieldMap["BEFORE_SEAT_RETURN_STATUS"] = h.BEFORESEATRETURNSTATUS
	h.fieldMap["TYPENO"] = h.TYPENO
	h.fieldMap["seat_type_name"] = h.SeatTypeName
	h.fieldMap["seat_type_en_name"] = h.SeatTypeEnName
	h.fieldMap["LIBNO"] = h.LIBNO
	h.fieldMap["lib_name"] = h.LibName
	h.fieldMap["lib_en_name"] = h.LibEnName
	h.fieldMap["FLOORNO"] = h.FLOORNO
	h.fieldMap["floor_name"] = h.FloorName
	h.fieldMap["floor_en_name"] = h.FloorEnName
	h.fieldMap["FLOOR"] = h.FLOOR
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
	h.fieldMap["APPROVAL_YN"] = h.APPROVALYN
	h.fieldMap["USE_MIN"] = h.USEMIN
	h.fieldMap["CAN_CONT_MIN"] = h.CANCONTMIN
	h.fieldMap["CONT_MIN"] = h.CONTMIN
	h.fieldMap["MAX_CONT_CNT"] = h.MAXCONTCNT
	h.fieldMap["NOSHOW_YN"] = h.NOSHOWYN
	h.fieldMap["NOSHOW_MIN"] = h.NOSHOWMIN
	h.fieldMap["MEMBERS"] = h.MEMBERS
	h.fieldMap["PURPOSE"] = h.PURPOSE
	h.fieldMap["MEMBERCNT"] = h.MEMBERCNT
	h.fieldMap["EMAIL"] = h.EMAIL
	h.fieldMap["TEL"] = h.TEL
	h.fieldMap["KEY_STATUS"] = h.KEYSTATUS
	h.fieldMap["GR_MIN"] = h.GRMIN
	h.fieldMap["GR_MAX"] = h.GRMAX
	h.fieldMap["use_approval"] = h.UseApproval
	h.fieldMap["reason"] = h.Reason
	h.fieldMap["CSCHOOLNO"] = h.CSCHOOLNO
	h.fieldMap["CBIGO"] = h.CBIGO
	h.fieldMap["alt_pid"] = h.AltPid
	h.fieldMap["username"] = h.Username
	h.fieldMap["pat_type"] = h.PatType
	h.fieldMap["pat_name"] = h.PatName
	h.fieldMap["dept_code"] = h.DeptCode
	h.fieldMap["dept_name"] = h.DeptName
	h.fieldMap["college_code"] = h.CollegeCode
	h.fieldMap["college_name"] = h.CollegeName
}

func (h hCV_view_seat_booking_tong) clone(db *gorm.DB) hCV_view_seat_booking_tong {
	h.hCV_view_seat_booking_tongDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hCV_view_seat_booking_tong) replaceDB(db *gorm.DB) hCV_view_seat_booking_tong {
	h.hCV_view_seat_booking_tongDo.ReplaceDB(db)
	return h
}

type hCV_view_seat_booking_tongDo struct{ gen.DO }

type IHCV_view_seat_booking_tongDo interface {
	gen.SubQuery
	Debug() IHCV_view_seat_booking_tongDo
	WithContext(ctx context.Context) IHCV_view_seat_booking_tongDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHCV_view_seat_booking_tongDo
	WriteDB() IHCV_view_seat_booking_tongDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHCV_view_seat_booking_tongDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHCV_view_seat_booking_tongDo
	Not(conds ...gen.Condition) IHCV_view_seat_booking_tongDo
	Or(conds ...gen.Condition) IHCV_view_seat_booking_tongDo
	Select(conds ...field.Expr) IHCV_view_seat_booking_tongDo
	Where(conds ...gen.Condition) IHCV_view_seat_booking_tongDo
	Order(conds ...field.Expr) IHCV_view_seat_booking_tongDo
	Distinct(cols ...field.Expr) IHCV_view_seat_booking_tongDo
	Omit(cols ...field.Expr) IHCV_view_seat_booking_tongDo
	Join(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_tongDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_tongDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_tongDo
	Group(cols ...field.Expr) IHCV_view_seat_booking_tongDo
	Having(conds ...gen.Condition) IHCV_view_seat_booking_tongDo
	Limit(limit int) IHCV_view_seat_booking_tongDo
	Offset(offset int) IHCV_view_seat_booking_tongDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_seat_booking_tongDo
	Unscoped() IHCV_view_seat_booking_tongDo
	Create(values ...*model.HCV_view_seat_booking_tong) error
	CreateInBatches(values []*model.HCV_view_seat_booking_tong, batchSize int) error
	Save(values ...*model.HCV_view_seat_booking_tong) error
	First() (*model.HCV_view_seat_booking_tong, error)
	Take() (*model.HCV_view_seat_booking_tong, error)
	Last() (*model.HCV_view_seat_booking_tong, error)
	Find() ([]*model.HCV_view_seat_booking_tong, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_seat_booking_tong, err error)
	FindInBatches(result *[]*model.HCV_view_seat_booking_tong, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HCV_view_seat_booking_tong) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHCV_view_seat_booking_tongDo
	Assign(attrs ...field.AssignExpr) IHCV_view_seat_booking_tongDo
	Joins(fields ...field.RelationField) IHCV_view_seat_booking_tongDo
	Preload(fields ...field.RelationField) IHCV_view_seat_booking_tongDo
	FirstOrInit() (*model.HCV_view_seat_booking_tong, error)
	FirstOrCreate() (*model.HCV_view_seat_booking_tong, error)
	FindByPage(offset int, limit int) (result []*model.HCV_view_seat_booking_tong, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHCV_view_seat_booking_tongDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hCV_view_seat_booking_tongDo) Debug() IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Debug())
}

func (h hCV_view_seat_booking_tongDo) WithContext(ctx context.Context) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hCV_view_seat_booking_tongDo) ReadDB() IHCV_view_seat_booking_tongDo {
	return h.Clauses(dbresolver.Read)
}

func (h hCV_view_seat_booking_tongDo) WriteDB() IHCV_view_seat_booking_tongDo {
	return h.Clauses(dbresolver.Write)
}

func (h hCV_view_seat_booking_tongDo) Session(config *gorm.Session) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Session(config))
}

func (h hCV_view_seat_booking_tongDo) Clauses(conds ...clause.Expression) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hCV_view_seat_booking_tongDo) Returning(value interface{}, columns ...string) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hCV_view_seat_booking_tongDo) Not(conds ...gen.Condition) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hCV_view_seat_booking_tongDo) Or(conds ...gen.Condition) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hCV_view_seat_booking_tongDo) Select(conds ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hCV_view_seat_booking_tongDo) Where(conds ...gen.Condition) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hCV_view_seat_booking_tongDo) Order(conds ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hCV_view_seat_booking_tongDo) Distinct(cols ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hCV_view_seat_booking_tongDo) Omit(cols ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hCV_view_seat_booking_tongDo) Join(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hCV_view_seat_booking_tongDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hCV_view_seat_booking_tongDo) RightJoin(table schema.Tabler, on ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hCV_view_seat_booking_tongDo) Group(cols ...field.Expr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hCV_view_seat_booking_tongDo) Having(conds ...gen.Condition) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hCV_view_seat_booking_tongDo) Limit(limit int) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hCV_view_seat_booking_tongDo) Offset(offset int) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hCV_view_seat_booking_tongDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hCV_view_seat_booking_tongDo) Unscoped() IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hCV_view_seat_booking_tongDo) Create(values ...*model.HCV_view_seat_booking_tong) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hCV_view_seat_booking_tongDo) CreateInBatches(values []*model.HCV_view_seat_booking_tong, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hCV_view_seat_booking_tongDo) Save(values ...*model.HCV_view_seat_booking_tong) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hCV_view_seat_booking_tongDo) First() (*model.HCV_view_seat_booking_tong, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_tong), nil
	}
}

func (h hCV_view_seat_booking_tongDo) Take() (*model.HCV_view_seat_booking_tong, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_tong), nil
	}
}

func (h hCV_view_seat_booking_tongDo) Last() (*model.HCV_view_seat_booking_tong, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_tong), nil
	}
}

func (h hCV_view_seat_booking_tongDo) Find() ([]*model.HCV_view_seat_booking_tong, error) {
	result, err := h.DO.Find()
	return result.([]*model.HCV_view_seat_booking_tong), err
}

func (h hCV_view_seat_booking_tongDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HCV_view_seat_booking_tong, err error) {
	buf := make([]*model.HCV_view_seat_booking_tong, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hCV_view_seat_booking_tongDo) FindInBatches(result *[]*model.HCV_view_seat_booking_tong, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hCV_view_seat_booking_tongDo) Attrs(attrs ...field.AssignExpr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hCV_view_seat_booking_tongDo) Assign(attrs ...field.AssignExpr) IHCV_view_seat_booking_tongDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hCV_view_seat_booking_tongDo) Joins(fields ...field.RelationField) IHCV_view_seat_booking_tongDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hCV_view_seat_booking_tongDo) Preload(fields ...field.RelationField) IHCV_view_seat_booking_tongDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hCV_view_seat_booking_tongDo) FirstOrInit() (*model.HCV_view_seat_booking_tong, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_tong), nil
	}
}

func (h hCV_view_seat_booking_tongDo) FirstOrCreate() (*model.HCV_view_seat_booking_tong, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HCV_view_seat_booking_tong), nil
	}
}

func (h hCV_view_seat_booking_tongDo) FindByPage(offset int, limit int) (result []*model.HCV_view_seat_booking_tong, count int64, err error) {
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

func (h hCV_view_seat_booking_tongDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hCV_view_seat_booking_tongDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hCV_view_seat_booking_tongDo) Delete(models ...*model.HCV_view_seat_booking_tong) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hCV_view_seat_booking_tongDo) withDO(do gen.Dao) *hCV_view_seat_booking_tongDo {
	h.DO = *do.(*gen.DO)
	return h
}
