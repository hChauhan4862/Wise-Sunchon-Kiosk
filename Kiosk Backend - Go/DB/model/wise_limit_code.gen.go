// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseLimitCode = "wise_limit_code"

// WiseLimitCode mapped from table <wise_limit_code>
type WiseLimitCode struct {
	LimitNo    *int64  `gorm:"column:limit_no" json:"limit_no"`
	LimitName  *string `gorm:"column:limit_name" json:"limit_name"`
	LimitDay   *int64  `gorm:"column:limit_day" json:"limit_day"`
	UseYn      *string `gorm:"column:use_yn" json:"use_yn"`
	LimitGubun *int64  `gorm:"column:limit_gubun" json:"limit_gubun"`
	Bigo       *string `gorm:"column:bigo" json:"bigo"`
}

// TableName WiseLimitCode's table name
func (*WiseLimitCode) TableName() string {
	return TableNameWiseLimitCode
}
