// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameHCV_TV_YESTERDAY_USE_TIME = "TV_YESTERDAY_USE_TIME"

// HCV_TV_YESTERDAY_USE_TIME mapped from table <TV_YESTERDAY_USE_TIME>
type HCV_TV_YESTERDAY_USE_TIME struct {
	Libno      int64      `gorm:"column:libno;not null" json:"libno"`
	UseDate    *string    `gorm:"column:use_date" json:"use_date"`
	StartTime  *string    `gorm:"column:start_time" json:"start_time"`
	EndTime    *string    `gorm:"column:end_time" json:"end_time"`
	STARTTIME  *time.Time `gorm:"column:STARTTIME" json:"STARTTIME"`
	RETURNTIME *time.Time `gorm:"column:RETURNTIME" json:"RETURNTIME"`
	STATUS     *int64     `gorm:"column:STATUS" json:"STATUS"`
	Gubun      int64      `gorm:"column:gubun;not null" json:"gubun"`
}

// TableName HCV_TV_YESTERDAY_USE_TIME's table name
func (*HCV_TV_YESTERDAY_USE_TIME) TableName() string {
	return TableNameHCV_TV_YESTERDAY_USE_TIME
}
