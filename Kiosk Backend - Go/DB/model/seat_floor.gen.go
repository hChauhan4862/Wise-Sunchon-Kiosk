// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatFloor = "seat_floor"

// SeatFloor mapped from table <seat_floor>
type SeatFloor struct {
	FLOORNO     int64   `gorm:"column:FLOORNO;primaryKey" json:"FLOORNO"`
	NAME        *string `gorm:"column:NAME" json:"NAME"`
	ENNAME      *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	FLOORIMAGE  *string `gorm:"column:FLOOR_IMAGE" json:"FLOOR_IMAGE"`
	LIBNO       int64   `gorm:"column:LIBNO;not null" json:"LIBNO"`
	FLOORIMAGE2 *string `gorm:"column:FLOOR_IMAGE2" json:"FLOOR_IMAGE2"`
	FLOOR       *int64  `gorm:"column:FLOOR" json:"FLOOR"`
}

// TableName SeatFloor's table name
func (*SeatFloor) TableName() string {
	return TableNameSeatFloor
}
