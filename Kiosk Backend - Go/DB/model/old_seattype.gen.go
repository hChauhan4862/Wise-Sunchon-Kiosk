// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_old_seattype = "old_seattype"

// HCV_old_seattype mapped from table <old_seattype>
type HCV_old_seattype struct {
	TYPENO        int64   `gorm:"column:TYPENO;not null" json:"TYPENO"`
	NAME          *string `gorm:"column:NAME" json:"NAME"`
	ENNAME        *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	MEDIATYPES    *string `gorm:"column:MEDIATYPES" json:"MEDIATYPES"`
	EQTYPES       *string `gorm:"column:EQTYPES" json:"EQTYPES"`
	ISKIOSKASSIGN *int64  `gorm:"column:IS_KIOSKASSIGN" json:"IS_KIOSKASSIGN"`
	DAYMAXUSECNT  *int64  `gorm:"column:DAY_MAXUSECNT" json:"DAY_MAXUSECNT"`
	DAYMAXUSEMIN  *int64  `gorm:"column:DAY_MAXUSEMIN" json:"DAY_MAXUSEMIN"`
	MONMAXUSECNT  *int64  `gorm:"column:MON_MAXUSECNT" json:"MON_MAXUSECNT"`
	MONMAXUSEMIN  *int64  `gorm:"column:MON_MAXUSEMIN" json:"MON_MAXUSEMIN"`
	PCLOGINCHECK  *int64  `gorm:"column:PC_LOGIN_CHECK" json:"PC_LOGIN_CHECK"`
	AUTOPRINTCNT  *int64  `gorm:"column:AUTO_PRINT_CNT" json:"AUTO_PRINT_CNT"`
}

// TableName HCV_old_seattype's table name
func (*HCV_old_seattype) TableName() string {
	return TableNameHCV_old_seattype
}
