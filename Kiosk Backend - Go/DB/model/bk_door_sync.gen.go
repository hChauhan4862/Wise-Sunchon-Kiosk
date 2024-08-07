// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBKDOORSYNC = "BK_DOOR_SYNC"

// BKDOORSYNC mapped from table <BK_DOOR_SYNC>
type BKDOORSYNC struct {
	BSEQNO    int64      `gorm:"column:BSEQNO;primaryKey" json:"BSEQNO"`
	STATUS    *int64     `gorm:"column:STATUS" json:"STATUS"`
	SCHOOLNO  *string    `gorm:"column:SCHOOLNO" json:"SCHOOLNO"`
	USESTART  *time.Time `gorm:"column:USESTART" json:"USESTART"`
	USEEXPIRE *time.Time `gorm:"column:USEEXPIRE" json:"USEEXPIRE"`
	MEMBERS   *string    `gorm:"column:MEMBERS" json:"MEMBERS"`
	SEATNO    *int64     `gorm:"column:SEATNO" json:"SEATNO"`
	SYNCTIME  *time.Time `gorm:"column:SYNC_TIME" json:"SYNC_TIME"`
	SYNCCHECK *string    `gorm:"column:SYNC_CHECK;default:0" json:"SYNC_CHECK"`
}

// TableName BKDOORSYNC's table name
func (*BKDOORSYNC) TableName() string {
	return TableNameBKDOORSYNC
}
