// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatSinbun = "seat_sinbun"

// SeatSinbun mapped from table <seat_sinbun>
type SeatSinbun struct {
	PatType *string `gorm:"column:pat_type" json:"pat_type"`
	PatName *string `gorm:"column:pat_name" json:"pat_name"`
}

// TableName SeatSinbun's table name
func (*SeatSinbun) TableName() string {
	return TableNameSeatSinbun
}