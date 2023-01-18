package models

import "time"

type ResourceGroup struct{
	Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name               string `gorm:"column:name;size:255" json:"name"`
	GroupType string         `gorm:"column:type;size:255" json:"type"`
	ParentGroupId   int `gorm:"column:id;default:0;NOT NULL" json:"parent_group_id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (rg *ResourceGroup) TableName() string {
	return "resource_groups"
}