package models

import "time"

type Dashboard struct{
	Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name               string `gorm:"column:name;size:255" json:"name"`
	ResourceGroupId int32         `gorm:"column:resource_group_id;default:0" json:"resource_group_id"`
	ResourceGroup   ResourceGroup `gorm:"ForeignKey:ResourceGroupId;AssociationForeignKey:Id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (d *Dashboard) TableName() string {
	return "Dashboards"
}