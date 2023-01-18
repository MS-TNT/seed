package models

import "time"

type QueryResult struct{
	Id  int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DataSourceId      int  `gorm:"column:data_source_id" json:"data_source_id"`
	DataSource          DataSource    `gorm:"ForeignKey:DataSourceId;AssociationForeignKey:Id"`
	QueryText string  `gorm:"column:query" json:"query_text"`
	CreatedAt          time.Time
}

func (qr *QueryResult) TableName() string {
	return "query_results"
}