package models

import "time"

type Widget struct{
	Id              int           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	VisualizationId string        `gorm:"column:visualization_id;size:255;NOT NULL" json:"visualization_id"`
	Visualization   Visualization `gorm:"ForeignKey:VisualizationId;AssociationForeignKey:Id"`
	DashboardId     int           `gorm:"column:dashboard_id;NOT NULL" json:"dashboard_id"`
	Dashboard       Dashboard     `gorm:"ForeignKey:DashboardId;AssociationForeignKey:Id"`
	Options         string        `gorm:"column:options;serializer:json" json:"options"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (w *Widget) TableName() string {
	return "widgets"
}