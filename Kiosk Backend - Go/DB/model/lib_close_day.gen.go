// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLIBCLOSEDAY = "LIB_CLOSE_DAY"

// LIBCLOSEDAY mapped from table <LIB_CLOSE_DAY>
type LIBCLOSEDAY struct {
	SEQNO     int64     `gorm:"column:SEQ_NO;primaryKey" json:"SEQ_NO"`
	TITLE     string    `gorm:"column:TITLE;not null" json:"TITLE"`
	ENTITLE   *string   `gorm:"column:EN_TITLE" json:"EN_TITLE"`
	LIBNO     int64     `gorm:"column:LIBNO;not null" json:"LIBNO"`
	FLOOR     *int64    `gorm:"column:FLOOR" json:"FLOOR"`
	ROOMNO    *int64    `gorm:"column:ROOMNO" json:"ROOMNO"`
	CLOSEDTFR string    `gorm:"column:CLOSE_DT_FR;not null" json:"CLOSE_DT_FR"`
	CLOSETMFR string    `gorm:"column:CLOSE_TM_FR;not null" json:"CLOSE_TM_FR"`
	CLOSEDTTO string    `gorm:"column:CLOSE_DT_TO;not null" json:"CLOSE_DT_TO"`
	CLOSETMTO string    `gorm:"column:CLOSE_TM_TO;not null" json:"CLOSE_TM_TO"`
	REGDT     time.Time `gorm:"column:REG_DT;not null" json:"REG_DT"`
	REGID     string    `gorm:"column:REG_ID;not null" json:"REG_ID"`
	MODDT     time.Time `gorm:"column:MOD_DT;not null" json:"MOD_DT"`
	MODID     string    `gorm:"column:MOD_ID;not null" json:"MOD_ID"`
	REMARK    *string   `gorm:"column:REMARK" json:"REMARK"`
	ENREMARK  *string   `gorm:"column:EN_REMARK" json:"EN_REMARK"`
}

// TableName LIBCLOSEDAY's table name
func (*LIBCLOSEDAY) TableName() string {
	return TableNameLIBCLOSEDAY
}