// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameHCV_VIEW_GATE_HISTORY = "VIEW_GATE_HISTORY"

// HCV_VIEW_GATE_HISTORY mapped from table <VIEW_GATE_HISTORY>
type HCV_VIEW_GATE_HISTORY struct {
	SEQNO    int64      `gorm:"column:SEQNO;not null" json:"SEQNO"`
	SCHOOLNO *string    `gorm:"column:SCHOOLNO" json:"SCHOOLNO"`
	GATE     *int64     `gorm:"column:GATE" json:"GATE"`
	GATENO   *string    `gorm:"column:GATENO" json:"GATENO"`
	GUBUN    *string    `gorm:"column:GUBUN" json:"GUBUN"`
	LOGTIME  *time.Time `gorm:"column:LOGTIME" json:"LOGTIME"`
	NAME     *string    `gorm:"column:NAME" json:"NAME"`
	SEX      *string    `gorm:"column:SEX" json:"SEX"`
	SOSOKCD1 *string    `gorm:"column:SOSOK_CD1" json:"SOSOK_CD1"`
	SOSOKCD2 *string    `gorm:"column:SOSOK_CD2" json:"SOSOK_CD2"`
	SINBUN   *string    `gorm:"column:SINBUN" json:"SINBUN"`
	STATUSCD *string    `gorm:"column:STATUS_CD" json:"STATUS_CD"`
}

// TableName HCV_VIEW_GATE_HISTORY's table name
func (*HCV_VIEW_GATE_HISTORY) TableName() string {
	return TableNameHCV_VIEW_GATE_HISTORY
}
