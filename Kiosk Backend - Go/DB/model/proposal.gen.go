// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProposal = "proposal"

// Proposal mapped from table <proposal>
type Proposal struct {
	PropID      int64   `gorm:"column:prop_id;primaryKey" json:"prop_id"`
	UserID      string  `gorm:"column:user_id;not null" json:"user_id"`
	UserName    string  `gorm:"column:user_name;not null" json:"user_name"`
	PropTitle   string  `gorm:"column:prop_title;not null" json:"prop_title"`
	PropContent string  `gorm:"column:prop_content;not null" json:"prop_content"`
	DelYn       string  `gorm:"column:del_yn;not null" json:"del_yn"`
	RegDt       string  `gorm:"column:reg_dt;not null" json:"reg_dt"`
	ANSWERYN    *string `gorm:"column:ANSWER_YN;not null;default:N" json:"ANSWER_YN"`
}

// TableName Proposal's table name
func (*Proposal) TableName() string {
	return TableNameProposal
}
