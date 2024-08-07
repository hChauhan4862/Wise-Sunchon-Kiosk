// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_old_sectorcode = "old_sectorcode"

// HCV_old_sectorcode mapped from table <old_sectorcode>
type HCV_old_sectorcode struct {
	Sectorno         int64   `gorm:"column:sectorno;not null" json:"sectorno"`
	Name             *string `gorm:"column:name" json:"name"`
	EnName           *string `gorm:"column:en_name" json:"en_name"`
	Typeno           *int64  `gorm:"column:typeno" json:"typeno"`
	Roomno           *int64  `gorm:"column:roomno" json:"roomno"`
	Bookable         int64   `gorm:"column:bookable;not null" json:"bookable"`
	Mediabookable    int64   `gorm:"column:mediabookable;not null" json:"mediabookable"`
	Eqbookable       int64   `gorm:"column:eqbookable;not null" json:"eqbookable"`
	BookDayFrom      *int64  `gorm:"column:book_day_from" json:"book_day_from"`
	BookDayTo        *int64  `gorm:"column:book_day_to" json:"book_day_to"`
	UseApproval      int64   `gorm:"column:use_approval;not null" json:"use_approval"`
	UserPos          string  `gorm:"column:user_pos;not null" json:"user_pos"`
	UserStatus       string  `gorm:"column:user_status;not null" json:"user_status"`
	MobileBookable   int64   `gorm:"column:mobile_bookable;not null" json:"mobile_bookable"`
	MobileAssignable int64   `gorm:"column:mobile_assignable;not null" json:"mobile_assignable"`
}

// TableName HCV_old_sectorcode's table name
func (*HCV_old_sectorcode) TableName() string {
	return TableNameHCV_old_sectorcode
}
