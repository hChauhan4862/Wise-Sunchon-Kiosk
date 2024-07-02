// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGATEHISTORY = "GATE_HISTORY"

// GATEHISTORY mapped from table <GATE_HISTORY>
type GATEHISTORY struct {
	SEQNO    int64      `gorm:"column:SEQNO;primaryKey" json:"SEQNO"`
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

// TableName GATEHISTORY's table name
func (*GATEHISTORY) TableName() string {
	return TableNameGATEHISTORY
}