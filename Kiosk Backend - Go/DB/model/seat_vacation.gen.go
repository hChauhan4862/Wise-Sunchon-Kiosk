// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatVacation = "seat_vacation"

// SeatVacation mapped from table <seat_vacation>
type SeatVacation struct {
	TITLE   string `gorm:"column:TITLE;primaryKey" json:"TITLE"`
	DAYFROM *int64 `gorm:"column:DAY_FROM" json:"DAY_FROM"`
	DAYTO   *int64 `gorm:"column:DAY_TO" json:"DAY_TO"`
}

// TableName SeatVacation's table name
func (*SeatVacation) TableName() string {
	return TableNameSeatVacation
}