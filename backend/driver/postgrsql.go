package driver

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"seed/backend/apiserver/models"
)

type PostgreSql struct {
	DB *gorm.DB
}

func NewPostgreSql(host, user, pwd, dbname string, port int32) *PostgreSql {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, pwd, port, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{}, &models.ResourceGroup{}, &models.DataSource{}, &models.Query{},
		&models.QueryResult{}, &models.Dashboard{}, &models.Visualization{}, &models.Widget{})
	if err != nil {
		panic(err)
	}
	return &PostgreSql{DB: db}
}
