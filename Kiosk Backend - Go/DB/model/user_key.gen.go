// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserKey = "user_key"

// UserKey mapped from table <user_key>
type UserKey struct {
	UserID    string `gorm:"column:user_id;not null" json:"user_id"`
	UserKey   string `gorm:"column:user_key;not null" json:"user_key"`
	DecodeKey string `gorm:"column:decode_key;not null" json:"decode_key"`
}

// TableName UserKey's table name
func (*UserKey) TableName() string {
	return TableNameUserKey
}
