// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatSmart = "seat_smart"

// SeatSmart mapped from table <seat_smart>
type SeatSmart struct {
	RoomNo *int64 `gorm:"column:room_no" json:"room_no"`
}

// TableName SeatSmart's table name
func (*SeatSmart) TableName() string {
	return TableNameSeatSmart
}
