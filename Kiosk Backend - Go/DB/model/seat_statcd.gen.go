// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatStatcd = "seat_statcd"

// SeatStatcd mapped from table <seat_statcd>
type SeatStatcd struct {
	STATUS   int64   `gorm:"column:STATUS;primaryKey" json:"STATUS"`
	NAME     *string `gorm:"column:NAME" json:"NAME"`
	ENNAME   *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	BOOKABLE *int64  `gorm:"column:BOOKABLE" json:"BOOKABLE"`
	DISABLED *int64  `gorm:"column:DISABLED" json:"DISABLED"`
}

// TableName SeatStatcd's table name
func (*SeatStatcd) TableName() string {
	return TableNameSeatStatcd
}
