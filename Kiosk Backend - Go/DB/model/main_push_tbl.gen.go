// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMainPushTbl = "main_push_tbl"

// MainPushTbl mapped from table <main_push_tbl>
type MainPushTbl struct {
	PushSeq     int64   `gorm:"column:push_seq;primaryKey" json:"push_seq"`
	PushGb      string  `gorm:"column:push_gb;not null" json:"push_gb"`
	SendDate    string  `gorm:"column:send_date;not null" json:"send_date"`
	ReceiveDate *string `gorm:"column:receive_date" json:"receive_date"`
	SendCnt     int64   `gorm:"column:send_cnt;not null" json:"send_cnt"`
	ReceiveChk  *string `gorm:"column:receive_chk" json:"receive_chk"`
	Msg         *string `gorm:"column:msg" json:"msg"`
	Pid         string  `gorm:"column:pid;not null" json:"pid"`
}

// TableName MainPushTbl's table name
func (*MainPushTbl) TableName() string {
	return TableNameMainPushTbl
}