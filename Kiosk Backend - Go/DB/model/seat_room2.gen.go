// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSeatRoom2 = "seat_room2"

// SeatRoom2 mapped from table <seat_room2>
type SeatRoom2 struct {
	ROOMNO     int64   `gorm:"column:ROOMNO;primaryKey" json:"ROOMNO"`
	NAME       *string `gorm:"column:NAME" json:"NAME"`
	ENNAME     *string `gorm:"column:EN_NAME" json:"EN_NAME"`
	FLOORNO    *int64  `gorm:"column:FLOORNO" json:"FLOORNO"`
	APPROVALYN *string `gorm:"column:APPROVAL_YN" json:"APPROVAL_YN"`
	BOOKINGYN  *string `gorm:"column:BOOKING_YN" json:"BOOKING_YN"`
	OPENTYPE   *int64  `gorm:"column:OPEN_TYPE" json:"OPEN_TYPE"`
	USEMIN     *int64  `gorm:"column:USE_MIN" json:"USE_MIN"`
	CANCONTMIN *int64  `gorm:"column:CAN_CONT_MIN" json:"CAN_CONT_MIN"`
	CONTMIN    *int64  `gorm:"column:CONT_MIN" json:"CONT_MIN"`
	MAXCONTCNT *int64  `gorm:"column:MAX_CONT_CNT" json:"MAX_CONT_CNT"`
	USERPOS    *string `gorm:"column:USER_POS" json:"USER_POS"`
	USERDEPT   *string `gorm:"column:USER_DEPT" json:"USER_DEPT"`
	USERSTATUS *string `gorm:"column:USER_STATUS" json:"USER_STATUS"`
	NOSHOWYN   *string `gorm:"column:NOSHOW_YN" json:"NOSHOW_YN"`
	NOSHOWMIN  *int64  `gorm:"column:NOSHOW_MIN" json:"NOSHOW_MIN"`
}

// TableName SeatRoom2's table name
func (*SeatRoom2) TableName() string {
	return TableNameSeatRoom2
}
