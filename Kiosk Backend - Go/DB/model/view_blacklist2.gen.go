// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_VIEW_BLACKLIST2 = "VIEW_BLACKLIST2"

// HCV_VIEW_BLACKLIST2 mapped from table <VIEW_BLACKLIST2>
type HCV_VIEW_BLACKLIST2 struct {
	BLSEQNO      int64   `gorm:"column:BLSEQNO;not null" json:"BLSEQNO"`
	SCHOOLNO     string  `gorm:"column:SCHOOLNO;not null" json:"SCHOOLNO"`
	USERNAME     *string `gorm:"column:USER_NAME" json:"USER_NAME"`
	REASON       *string `gorm:"column:REASON" json:"REASON"`
	REGTIME      *string `gorm:"column:REGTIME" json:"REGTIME"`
	BLOCKSTART   *string `gorm:"column:BLOCKSTART" json:"BLOCKSTART"`
	BLOCKEXPIRE  *string `gorm:"column:BLOCKEXPIRE" json:"BLOCKEXPIRE"`
	STATUS       *int64  `gorm:"column:STATUS" json:"STATUS"`
	SETFROM      string  `gorm:"column:SETFROM;not null" json:"SETFROM"`
	BLDAYBEFORE  *int64  `gorm:"column:BL_DAYBEFORE" json:"BL_DAYBEFORE"`
	BLLIMITCNT   *int64  `gorm:"column:BL_LIMITCNT" json:"BL_LIMITCNT"`
	BLBOOKINGDAY *int64  `gorm:"column:BL_BOOKINGDAY" json:"BL_BOOKINGDAY"`
}

// TableName HCV_VIEW_BLACKLIST2's table name
func (*HCV_VIEW_BLACKLIST2) TableName() string {
	return TableNameHCV_VIEW_BLACKLIST2
}
