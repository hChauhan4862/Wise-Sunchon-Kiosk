// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_TV_DAY_USE = "TV_DAY_USE"

// HCV_TV_DAY_USE mapped from table <TV_DAY_USE>
type HCV_TV_DAY_USE struct {
	Libno    int64   `gorm:"column:libno;not null" json:"libno"`
	UseDate  *string `gorm:"column:use_date" json:"use_date"`
	Weekday0 *int64  `gorm:"column:weekday0" json:"weekday0"`
	UseCnt   *int64  `gorm:"column:use_cnt" json:"use_cnt"`
}

// TableName HCV_TV_DAY_USE's table name
func (*HCV_TV_DAY_USE) TableName() string {
	return TableNameHCV_TV_DAY_USE
}
