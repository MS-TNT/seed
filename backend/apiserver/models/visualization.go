package models

import "time"

type Visualization struct{
	Id   int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name              string `gorm:"column:name;size:255" json:"name"`
	VisualizationType string `gorm:"column:type;size:255" json:"type"` //视图类型
	QueryId           int    `gorm:"column:query_id;NOT NULL" json:"query_id"`
	Query             Query  `gorm:"ForeignKey:QueryId;AssociationForeignKey:Id"`
	Description       string `gorm:"column:description;size:4096" json:"description"` //视图描述
	Options           string `gorm:"column:options;serializer:json" json:"options"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (v *Visualization) TableName() string {
	return "visualizations"
}
