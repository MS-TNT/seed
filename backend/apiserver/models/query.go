package models

import "time"

type Query struct {
	//gorm.Model
	Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	//orgId int `gorm:"column:org_id,foreignKey:organizations.id"`
	Version           int32  `gorm:"column:version" json:"version"`
	DataSourceId      int32  `gorm:"column:data_source_id" json:"data_source_id"`
	LatestQueryDataId int32  `gorm:"column:latest_query_data_id" json:"latest_query_data_id"`
	Name              string `gorm:"column:name;size:255"`
	Description       string `gorm:"column:description;size:255"`
	Query             string `gorm:"column:query;size:10000"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (q *Query) TableName() string {
	return "query"
}
