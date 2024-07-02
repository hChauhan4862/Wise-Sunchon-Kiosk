// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIDCARDQR = "IDCARD_QR"

// IDCARDQR mapped from table <IDCARD_QR>
type IDCARDQR struct {
	PID        string  `gorm:"column:PID;primaryKey" json:"PID"`
	DEVICEGB   string  `gorm:"column:DEVICE_GB;not null" json:"DEVICE_GB"`
	ISSUEDATE  string  `gorm:"column:ISSUE_DATE;primaryKey" json:"ISSUE_DATE"`
	QRCODE     string  `gorm:"column:QR_CODE;not null" json:"QR_CODE"`
	CHECKDIGIT *string `gorm:"column:CHECK_DIGIT" json:"CHECK_DIGIT"`
}

// TableName IDCARDQR's table name
func (*IDCARDQR) TableName() string {
	return TableNameIDCARDQR
}