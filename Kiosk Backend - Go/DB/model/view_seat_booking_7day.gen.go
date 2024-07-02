// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameHCV_view_seat_booking_7day = "view_seat_booking_7day"

// HCV_view_seat_booking_7day mapped from table <view_seat_booking_7day>
type HCV_view_seat_booking_7day struct {
	BSEQNO                 int64      `gorm:"column:BSEQNO;not null" json:"BSEQNO"`
	STATUS                 *int64     `gorm:"column:STATUS" json:"STATUS"`
	StatusName             *string    `gorm:"column:status_name" json:"status_name"`
	StatusEnName           *string    `gorm:"column:status_en_name" json:"status_en_name"`
	SCHOOLNO               *string    `gorm:"column:SCHOOLNO" json:"SCHOOLNO"`
	USERID                 *string    `gorm:"column:USERID" json:"USERID"`
	SEATNO                 *int64     `gorm:"column:SEATNO" json:"SEATNO"`
	SeatName               *string    `gorm:"column:seat_name" json:"seat_name"`
	SeatEnName             *string    `gorm:"column:seat_en_name" json:"seat_en_name"`
	SeatStatus             *int64     `gorm:"column:seat_status" json:"seat_status"`
	SeatStatusName         *string    `gorm:"column:seat_status_name" json:"seat_status_name"`
	SeatStatusEnName       *string    `gorm:"column:seat_status_en_name" json:"seat_status_en_name"`
	SeatVname              *string    `gorm:"column:seat_vname" json:"seat_vname"`
	ISCANCELED             *int64     `gorm:"column:ISCANCELED" json:"ISCANCELED"`
	ISSUEFROM              *int64     `gorm:"column:ISSUEFROM" json:"ISSUEFROM"`
	ISSUEDETAIL            *string    `gorm:"column:ISSUEDETAIL" json:"ISSUEDETAIL"`
	ISSUETYPENO            *int64     `gorm:"column:ISSUE_TYPE_NO" json:"ISSUE_TYPE_NO"`
	USESTART               *time.Time `gorm:"column:USESTART" json:"USESTART"`
	USEEXPIRE              *time.Time `gorm:"column:USEEXPIRE" json:"USEEXPIRE"`
	EXPIREREASON           *int64     `gorm:"column:EXPIREREASON" json:"EXPIREREASON"`
	ExpirereasonName       *string    `gorm:"column:expirereason_name" json:"expirereason_name"`
	ExpirereasonEnName     *string    `gorm:"column:expirereason_en_name" json:"expirereason_en_name"`
	PRINTCNT               *int64     `gorm:"column:PRINTCNT" json:"PRINTCNT"`
	REGTIME                *time.Time `gorm:"column:REGTIME" json:"REGTIME"`
	EXTENDMIN              *int64     `gorm:"column:EXTEND_MIN" json:"EXTEND_MIN"`
	EXTENDCNT              *int64     `gorm:"column:EXTEND_CNT" json:"EXTEND_CNT"`
	STARTTIME              *time.Time `gorm:"column:STARTTIME" json:"STARTTIME"`
	RETURNTIME             *time.Time `gorm:"column:RETURNTIME" json:"RETURNTIME"`
	USERNAME               *string    `gorm:"column:USER_NAME" json:"USER_NAME"`
	USERPOSITION           *string    `gorm:"column:USER_POSITION" json:"USER_POSITION"`
	COMPANYCODE            *string    `gorm:"column:COMPANY_CODE" json:"COMPANY_CODE"`
	DEPTCODE               *string    `gorm:"column:DEPT_CODE" json:"DEPT_CODE"`
	GATEOUTTIME            *time.Time `gorm:"column:GATEOUTTIME" json:"GATEOUTTIME"`
	ISADMINBOOKING         *int64     `gorm:"column:ISADMINBOOKING" json:"ISADMINBOOKING"`
	CHECKGATEINTIME        *time.Time `gorm:"column:CHECK_GATEINTIME" json:"CHECK_GATEINTIME"`
	CHECKGATEIN            *int64     `gorm:"column:CHECK_GATEIN" json:"CHECK_GATEIN"`
	MOBILE                 *string    `gorm:"column:MOBILE" json:"MOBILE"`
	PCLOGINCHECKSTATUS     *int64     `gorm:"column:PCLOGIN_CHECK_STATUS" json:"PCLOGIN_CHECK_STATUS"`
	EXTENDMSGSTATUS        *int64     `gorm:"column:EXTEND_MSG_STATUS" json:"EXTEND_MSG_STATUS"`
	BEFORESEATRETURNSTATUS *int64     `gorm:"column:BEFORE_SEAT_RETURN_STATUS" json:"BEFORE_SEAT_RETURN_STATUS"`
	TYPENO                 *int64     `gorm:"column:TYPENO" json:"TYPENO"`
	SeatTypeName           *string    `gorm:"column:seat_type_name" json:"seat_type_name"`
	SeatTypeEnName         *string    `gorm:"column:seat_type_en_name" json:"seat_type_en_name"`
	LIBNO                  int64      `gorm:"column:LIBNO;not null" json:"LIBNO"`
	LibName                *string    `gorm:"column:lib_name" json:"lib_name"`
	LibEnName              *string    `gorm:"column:lib_en_name" json:"lib_en_name"`
	FLOORNO                int64      `gorm:"column:FLOORNO;not null" json:"FLOORNO"`
	FloorName              *string    `gorm:"column:floor_name" json:"floor_name"`
	FloorEnName            *string    `gorm:"column:floor_en_name" json:"floor_en_name"`
	FLOOR                  *int64     `gorm:"column:FLOOR" json:"FLOOR"`
	ROOMNO                 int64      `gorm:"column:ROOMNO;not null" json:"ROOMNO"`
	RoomName               *string    `gorm:"column:room_name" json:"room_name"`
	RoomEnName             *string    `gorm:"column:room_en_name" json:"room_en_name"`
	SECTORNO               int64      `gorm:"column:SECTORNO;not null" json:"SECTORNO"`
	SectorName             *string    `gorm:"column:sector_name" json:"sector_name"`
	SectorEnName           *string    `gorm:"column:sector_en_name" json:"sector_en_name"`
	BOOKINGYN              *string    `gorm:"column:BOOKING_YN" json:"BOOKING_YN"`
	MOBILEBOOKINGYN        *string    `gorm:"column:MOBILE_BOOKING_YN" json:"MOBILE_BOOKING_YN"`
	ASSIGNYN               *string    `gorm:"column:ASSIGN_YN" json:"ASSIGN_YN"`
	MOBILEASSIGNYN         *string    `gorm:"column:MOBILE_ASSIGN_YN" json:"MOBILE_ASSIGN_YN"`
	DAYFROM                *int64     `gorm:"column:DAY_FROM" json:"DAY_FROM"`
	DAYTO                  *int64     `gorm:"column:DAY_TO" json:"DAY_TO"`
	MEDIABOOKINGYN         *string    `gorm:"column:MEDIA_BOOKING_YN" json:"MEDIA_BOOKING_YN"`
	EQUIPBOOKINGYN         *string    `gorm:"column:EQUIP_BOOKING_YN" json:"EQUIP_BOOKING_YN"`
	APPROVALYN             *string    `gorm:"column:APPROVAL_YN" json:"APPROVAL_YN"`
	USEMIN                 *int64     `gorm:"column:USE_MIN" json:"USE_MIN"`
	CANCONTMIN             *int64     `gorm:"column:CAN_CONT_MIN" json:"CAN_CONT_MIN"`
	CONTMIN                *int64     `gorm:"column:CONT_MIN" json:"CONT_MIN"`
	MAXCONTCNT             *int64     `gorm:"column:MAX_CONT_CNT" json:"MAX_CONT_CNT"`
	NOSHOWYN               *string    `gorm:"column:NOSHOW_YN" json:"NOSHOW_YN"`
	NOSHOWMIN              *int64     `gorm:"column:NOSHOW_MIN" json:"NOSHOW_MIN"`
	MEMBERS                *string    `gorm:"column:MEMBERS" json:"MEMBERS"`
	PURPOSE                *string    `gorm:"column:PURPOSE" json:"PURPOSE"`
	MEMBERCNT              *int64     `gorm:"column:MEMBERCNT" json:"MEMBERCNT"`
	EMAIL                  *string    `gorm:"column:EMAIL" json:"EMAIL"`
	TEL                    *string    `gorm:"column:TEL" json:"TEL"`
	KEYSTATUS              *int64     `gorm:"column:KEY_STATUS" json:"KEY_STATUS"`
	GRMIN                  *int64     `gorm:"column:GR_MIN" json:"GR_MIN"`
	GRMAX                  *int64     `gorm:"column:GR_MAX" json:"GR_MAX"`
	UseApproval            *int64     `gorm:"column:use_approval" json:"use_approval"`
	Reason                 *string    `gorm:"column:reason" json:"reason"`
	CSCHOOLNO              *string    `gorm:"column:CSCHOOLNO" json:"CSCHOOLNO"`
	CBIGO                  *string    `gorm:"column:CBIGO" json:"CBIGO"`
}

// TableName HCV_view_seat_booking_7day's table name
func (*HCV_view_seat_booking_7day) TableName() string {
	return TableNameHCV_view_seat_booking_7day
}
