// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHCV_VIEW_POSITION = "VIEW_POSITION"

// HCV_VIEW_POSITION mapped from table <VIEW_POSITION>
type HCV_VIEW_POSITION struct {
	CODE      string `gorm:"column:CODE;not null" json:"CODE"`
	CODENM    string `gorm:"column:CODE_NM;not null" json:"CODE_NM"`
	DISPORDER *int64 `gorm:"column:DISP_ORDER" json:"DISP_ORDER"`
}

// TableName HCV_VIEW_POSITION's table name
func (*HCV_VIEW_POSITION) TableName() string {
	return TableNameHCV_VIEW_POSITION
}
