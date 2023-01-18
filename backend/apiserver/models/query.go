package models

import "time"

type Query struct {
	//gorm.Model
	Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	//orgId int `gorm:"column:org_id,foreignKey:organizations.id"`
	//Version           int  `gorm:"column:version" json:"version"`
	DataSourceId      int  `gorm:"column:data_source_id" json:"data_source_id"`
	DataSource          DataSource    `gorm:"ForeignKey:DataSourceId;AssociationForeignKey:Id"`
	LatestQueryResultId int  `gorm:"column:latest_query_result_id" json:"latest_query_result_id"`
	QueryResult         QueryResult   `gorm:"ForeignKey:LatestQueryResultId;AssociationForeignKey:Id"`
	ResourceGroupId int32         `gorm:"column:resource_group_id;default:0" json:"resource_group_id"`
	ResourceGroup   ResourceGroup `gorm:"ForeignKey:ResourceGroupId;AssociationForeignKey:Id"`
	Name              string `gorm:"column:name;size:255"`
	Description       string `gorm:"column:description;size:255"`
	Query             string `gorm:"column:query;size:10000"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (q *Query) TableName() string {
	return "queries"
}
