// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWiseRroom = "wise_rroom"

// WiseRroom mapped from table <wise_rroom>
type WiseRroom struct {
	RroomNo    *int64  `gorm:"column:rroom_no;default:NULL" json:"rroom_no"`
	RroomName  *string `gorm:"column:rroom_name;default:NULL" json:"rroom_name"`
	RroomGubun *string `gorm:"column:rroom_gubun;default:NULL" json:"rroom_gubun"`
	RroomType  *int64  `gorm:"column:rroom_type;default:NULL" json:"rroom_type"`
	UseYn      *string `gorm:"column:use_yn;default:NULL" json:"use_yn"`
	BlockYn    *string `gorm:"column:block_yn;default:NULL" json:"block_yn"`
	BlockDay   *int64  `gorm:"column:block_day;default:NULL" json:"block_day"`
	Bigo       *string `gorm:"column:bigo;default:NULL" json:"bigo"`
}

// TableName WiseRroom's table name
func (*WiseRroom) TableName() string {
	return TableNameWiseRroom
}
