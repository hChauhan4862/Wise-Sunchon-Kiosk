// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatConf = "seat_conf"

// SeatConf mapped from table <seat_conf>
type SeatConf struct {
	SeatCode  *string `gorm:"column:seat_code" json:"seat_code"`
	CodeValue *string `gorm:"column:code_value" json:"code_value"`
}

// TableName SeatConf's table name
func (*SeatConf) TableName() string {
	return TableNameSeatConf
}
