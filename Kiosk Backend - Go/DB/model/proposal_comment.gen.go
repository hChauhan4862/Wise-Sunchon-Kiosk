// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProposalComment = "proposal_comment"

// ProposalComment mapped from table <proposal_comment>
type ProposalComment struct {
	PropID     int64   `gorm:"column:prop_id;not null" json:"prop_id"`
	AdminID    string  `gorm:"column:admin_id;not null" json:"admin_id"`
	ComContent string  `gorm:"column:com_content;not null" json:"com_content"`
	RegDt      string  `gorm:"column:reg_dt;not null" json:"reg_dt"`
	Seq        int64   `gorm:"column:seq;not null" json:"seq"`
	COMTITLE   *string `gorm:"column:COM_TITLE" json:"COM_TITLE"`
}

// TableName ProposalComment's table name
func (*ProposalComment) TableName() string {
	return TableNameProposalComment
}
