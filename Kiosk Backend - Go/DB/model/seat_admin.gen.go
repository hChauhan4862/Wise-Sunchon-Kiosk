// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatAdmin = "seat_admin"

// SeatAdmin mapped from table <seat_admin>
type SeatAdmin struct {
	ID     string  `gorm:"column:id;not null" json:"id"`
	PassWd string  `gorm:"column:pass_wd;not null" json:"pass_wd"`
	Auth   string  `gorm:"column:auth;not null" json:"auth"`
	RegDt  string  `gorm:"column:reg_dt;not null" json:"reg_dt"`
	UpdDt  *string `gorm:"column:upd_dt" json:"upd_dt"`
}

// TableName SeatAdmin's table name
func (*SeatAdmin) TableName() string {
	return TableNameSeatAdmin
}
