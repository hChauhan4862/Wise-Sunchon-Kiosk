// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_old_libcode = "old_libcode"

// HCV_old_libcode mapped from table <old_libcode>
type HCV_old_libcode struct {
	Libno  int64   `gorm:"column:libno;not null" json:"libno"`
	Name   *string `gorm:"column:name" json:"name"`
	EnName *string `gorm:"column:en_name" json:"en_name"`
}

// TableName HCV_old_libcode's table name
func (*HCV_old_libcode) TableName() string {
	return TableNameHCV_old_libcode
}
