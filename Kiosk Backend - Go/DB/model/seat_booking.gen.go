// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSeatBooking = "seat_booking"

// SeatBooking mapped from table <seat_booking>
type SeatBooking struct {
	BSEQNO                 int64      `gorm:"column:BSEQNO;primaryKey;autoIncrement:false" json:"BSEQNO"`
	STATUS                 *int64     `gorm:"column:STATUS" json:"STATUS"`
	SCHOOLNO               *string    `gorm:"column:SCHOOLNO" json:"SCHOOLNO"`
	USERID                 *string    `gorm:"column:USERID" json:"USERID"`
	SEATNO                 *int64     `gorm:"column:SEATNO" json:"SEATNO"`
	ISCANCELED             *int64     `gorm:"column:ISCANCELED" json:"ISCANCELED"`
	ISSUEFROM              *int64     `gorm:"column:ISSUEFROM" json:"ISSUEFROM"`
	ISSUEDETAIL            *string    `gorm:"column:ISSUEDETAIL" json:"ISSUEDETAIL"`
	USESTART               *time.Time `gorm:"column:USESTART" json:"USESTART"`
	USEEXPIRE              *time.Time `gorm:"column:USEEXPIRE" json:"USEEXPIRE"`
	EXPIREREASON           *int64     `gorm:"column:EXPIREREASON" json:"EXPIREREASON"`
	QUOTAMIN               *int64     `gorm:"column:QUOTAMIN" json:"QUOTAMIN"`
	PRINTCNT               *int64     `gorm:"column:PRINTCNT" json:"PRINTCNT"`
	REGTIME                *time.Time `gorm:"column:REGTIME" json:"REGTIME"`
	EXTENDMIN              *int64     `gorm:"column:EXTEND_MIN" json:"EXTEND_MIN"`
	EXTENDCNT              *int64     `gorm:"column:EXTEND_CNT" json:"EXTEND_CNT"`
	STARTTIME              *time.Time `gorm:"column:STARTTIME" json:"STARTTIME"`
	RETURNTIME             *time.Time `gorm:"column:RETURNTIME" json:"RETURNTIME"`
	USERNAME               *string    `gorm:"column:USER_NAME" json:"USER_NAME"`
	USERPOSITION           *string    `gorm:"column:USER_POSITION" json:"USER_POSITION"`
	STATUSCODE             *string    `gorm:"column:STATUS_CODE" json:"STATUS_CODE"`
	GRADECODE              *string    `gorm:"column:GRADE_CODE" json:"GRADE_CODE"`
	CAMPUSDIV              *string    `gorm:"column:CAMPUS_DIV" json:"CAMPUS_DIV"`
	COMPANYCODE            *string    `gorm:"column:COMPANY_CODE" json:"COMPANY_CODE"`
	DEPTCODE               *string    `gorm:"column:DEPT_CODE" json:"DEPT_CODE"`
	MAJORCODE              *string    `gorm:"column:MAJOR_CODE" json:"MAJOR_CODE"`
	GATEOUTTIME            *time.Time `gorm:"column:GATEOUTTIME" json:"GATEOUTTIME"`
	ISADMINBOOKING         *int64     `gorm:"column:ISADMINBOOKING" json:"ISADMINBOOKING"`
	CHECKGATEINTIME        *time.Time `gorm:"column:CHECK_GATEINTIME" json:"CHECK_GATEINTIME"`
	CHECKGATEIN            *int64     `gorm:"column:CHECK_GATEIN" json:"CHECK_GATEIN"`
	MOBILE                 *string    `gorm:"column:MOBILE" json:"MOBILE"`
	PCLOGINCHECKSTATUS     *int64     `gorm:"column:PCLOGIN_CHECK_STATUS" json:"PCLOGIN_CHECK_STATUS"`
	EXTENDMSGSTATUS        *int64     `gorm:"column:EXTEND_MSG_STATUS" json:"EXTEND_MSG_STATUS"`
	ISSUETYPENO            *int64     `gorm:"column:ISSUE_TYPE_NO" json:"ISSUE_TYPE_NO"`
	BEFORESEATRETURNSTATUS *int64     `gorm:"column:BEFORE_SEAT_RETURN_STATUS" json:"BEFORE_SEAT_RETURN_STATUS"`
}

// TableName SeatBooking's table name
func (*SeatBooking) TableName() string {
	return TableNameSeatBooking
}