// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatSeat = "seat_seat"

// SeatSeat mapped from table <seat_seat>
type SeatSeat struct {
	SEATNO   int64   `gorm:"column:SEATNO;primaryKey" json:"SEATNO"`
	NAME     *string `gorm:"column:NAME" json:"NAME"`
	ENNAME   *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	VNAME    *string `gorm:"column:VNAME" json:"VNAME"`
	SECTORNO *int64  `gorm:"column:SECTORNO" json:"SECTORNO"`
	STATUS   *int64  `gorm:"column:STATUS" json:"STATUS"`
	POSX     *int64  `gorm:"column:POSX" json:"POSX"`
	POSY     *int64  `gorm:"column:POSY" json:"POSY"`
	POSW     *int64  `gorm:"column:POSW" json:"POSW"`
	POSH     *int64  `gorm:"column:POSH" json:"POSH"`
	ICONTYPE *int64  `gorm:"column:ICONTYPE" json:"ICONTYPE"`
	GRMIN    *int64  `gorm:"column:GR_MIN" json:"GR_MIN"`
	GRMAX    *int64  `gorm:"column:GR_MAX" json:"GR_MAX"`
	POSX2    *int64  `gorm:"column:POSX2" json:"POSX2"`
	POSY2    *int64  `gorm:"column:POSY2" json:"POSY2"`
	MANX     *int64  `gorm:"column:MANX" json:"MANX"`
	MANY     *int64  `gorm:"column:MANY" json:"MANY"`
}

// TableName SeatSeat's table name
func (*SeatSeat) TableName() string {
	return TableNameSeatSeat
}
