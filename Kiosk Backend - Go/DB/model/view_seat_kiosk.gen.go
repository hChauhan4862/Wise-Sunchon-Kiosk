// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameHCV_view_seat_kiosk = "view_seat_kiosk"

// HCV_view_seat_kiosk mapped from table <view_seat_kiosk>
type HCV_view_seat_kiosk struct {
	KioskNo           int64      `gorm:"column:KioskNo;not null" json:"KioskNo"`
	NAME              *string    `gorm:"column:NAME" json:"NAME"`
	ENNAME            *string    `gorm:"column:EN_NAME" json:"EN_NAME"`
	IPAddr            *string    `gorm:"column:IPAddr" json:"IPAddr"`
	MAC               *string    `gorm:"column:MAC" json:"MAC"`
	LibNo             *int64     `gorm:"column:LibNo" json:"LibNo"`
	FloorNo           *int64     `gorm:"column:FloorNo" json:"FloorNo"`
	Issuefrom         *int64     `gorm:"column:Issuefrom" json:"Issuefrom"`
	IssueDetail       *string    `gorm:"column:IssueDetail" json:"IssueDetail"`
	KioskStatus       *int64     `gorm:"column:KioskStatus" json:"KioskStatus"`
	AssignLibOnly     *int64     `gorm:"column:AssignLibOnly" json:"AssignLibOnly"`
	Assignable        *int64     `gorm:"column:Assignable" json:"Assignable"`
	Movable           *int64     `gorm:"column:Movable" json:"Movable"`
	Extendable        *int64     `gorm:"column:Extendable" json:"Extendable"`
	Returnable        *int64     `gorm:"column:Returnable" json:"Returnable"`
	StatusMemo        *string    `gorm:"column:StatusMemo" json:"StatusMemo"`
	InsertTime        *time.Time `gorm:"column:InsertTime" json:"InsertTime"`
	AdminID           *string    `gorm:"column:AdminID" json:"AdminID"`
	IsDelete          *int64     `gorm:"column:IsDelete" json:"IsDelete"`
	PaperAmount       *int64     `gorm:"column:PaperAmount" json:"PaperAmount"`
	PaperReplaceTime  *time.Time `gorm:"column:PaperReplaceTime" json:"PaperReplaceTime"`
	PaperReplaceAdmin *string    `gorm:"column:PaperReplaceAdmin" json:"PaperReplaceAdmin"`
	PrintErrorCode    *int64     `gorm:"column:PrintErrorCode" json:"PrintErrorCode"`
	PrintErrorTitle   *string    `gorm:"column:PrintErrorTitle" json:"PrintErrorTitle"`
	OnTime            *int64     `gorm:"column:OnTime" json:"OnTime"`
	OffTime           *int64     `gorm:"column:OffTime" json:"OffTime"`
	Floor             *int64     `gorm:"column:floor" json:"floor"`
	FloorName         *string    `gorm:"column:floor_name" json:"floor_name"`
	FloorEnName       *string    `gorm:"column:floor_en_name" json:"floor_en_name"`
	LibName           *string    `gorm:"column:lib_name" json:"lib_name"`
	LibEnName         *string    `gorm:"column:lib_en_name" json:"lib_en_name"`
}

// TableName HCV_view_seat_kiosk's table name
func (*HCV_view_seat_kiosk) TableName() string {
	return TableNameHCV_view_seat_kiosk
}
