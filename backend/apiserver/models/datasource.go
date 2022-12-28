package models

import "time"

type DataSource struct {
	//gorm.Model
	Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	//orgId int `gorm:"column:org_id,foreignKey:organizations.id"`
	Name               string `gorm:"column:name;size:255" json:"name"`
	DsType             string `gorm:"column:type;size:255" json:"type"`
	Options            string `gorm:"column:options;serializer:json" json:"options"`
	QueueName          string `gorm:"column:queue_name;size:255;default:queries"`
	ScheduledQueueName string `gorm:"column:scheduled_queue_name;size:255;default:scheduled_queries"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (ds *DataSource) TableName() string {
	return "data_sources"
}
