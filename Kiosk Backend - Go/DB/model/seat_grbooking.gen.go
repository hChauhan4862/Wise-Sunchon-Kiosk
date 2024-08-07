// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSeatGrbooking = "seat_grbooking"

// SeatGrbooking mapped from table <seat_grbooking>
type SeatGrbooking struct {
	BSEQNO    int64      `gorm:"column:BSEQNO;primaryKey" json:"BSEQNO"`
	MEMBERS   *string    `gorm:"column:MEMBERS" json:"MEMBERS"`
	PURPOSE   *string    `gorm:"column:PURPOSE" json:"PURPOSE"`
	MEMBERCNT *int64     `gorm:"column:MEMBERCNT" json:"MEMBERCNT"`
	EMAIL     *string    `gorm:"column:EMAIL" json:"EMAIL"`
	TEL       *string    `gorm:"column:TEL" json:"TEL"`
	USESTART  *time.Time `gorm:"column:USESTART" json:"USESTART"`
	USEEXPIRE *time.Time `gorm:"column:USEEXPIRE" json:"USEEXPIRE"`
	KEYSTATUS *int64     `gorm:"column:KEY_STATUS" json:"KEY_STATUS"`
	PURPOSENO *int64     `gorm:"column:PURPOSE_NO" json:"PURPOSE_NO"`
}

// TableName SeatGrbooking's table name
func (*SeatGrbooking) TableName() string {
	return TableNameSeatGrbooking
}
