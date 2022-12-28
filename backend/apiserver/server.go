package apiserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"seed/backend/apiserver/auth"
	"seed/backend/apiserver/handlers"
	"seed/backend/driver"
	"seed/backend/internal/log"
)

type SeedServer struct {
	engine     *gin.Engine
	dataSource *handlers.DataSourceHandler
}

func NewSeedServer() *SeedServer {
	//gin
	//router := gin.New()
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	//postgresql
	pgHost := viper.GetString("server.postgresql.host")
	pgUser := viper.GetString("server.postgresql.user")
	pgPwd := viper.GetString("server.postgresql.password")
	pgDb := viper.GetString("server.postgresql.database")
	pgPort := viper.GetInt32("server.postgresql.port")
	log.Infof("init postgresql param: host:%s, user:%s, pwd:%s, db:%s, port:%d", pgHost, pgUser, pgPwd, pgDb, pgPort)
	pg := driver.NewPostgreSql(pgHost, pgUser, pgPwd, pgDb, pgPort)
	//sender
	jh, err := handlers.NewJobHandler()
	if err != nil {
		log.Errorf("[apiServer][NewSeedServer]NewSendJobHandler err:%s", err)
		return nil
	}
	//datasource
	ds := handlers.NewDataSourceHandler(pg, jh)
	seedServer := SeedServer{
		engine:     router,
		dataSource: ds,
	}
	seedServer.initRouter()
	return &seedServer
}

func (ss *SeedServer) Run() {
	err := ss.engine.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
	if err != nil {

	}
}

func (ss *SeedServer) initRouter() {
	ss.engine.POST("/login", auth.Login)

	api := ss.engine.Group("/api")
	{
		//data_source
		api.GET("/data_sources", ss.dataSource.GetAll)                      //get all data_sources
		api.POST("/data_sources", ss.dataSource.Add)                        //add a data_source
		api.GET("/data_sources/:dataSourceId", ss.dataSource.Get)           //get data_source_id's data_source detail
		api.POST("/data_sources/:dataSourceId", ss.dataSource.Update)       //update data_source_id's data_source
		api.DELETE("/data_sources/:dataSourceId", ss.dataSource.Delete)     //delete data_source_id's data_source
		api.POST("/data_sources/:dataSourceId/test")                        //test data_source_id's data_source connect
		api.GET("/data_sources/:dataSourceId/schema", ss.dataSource.Schema) //get data_source_id's data_source schema info
		//query
		api.GET("/queries")             //get all queries
		api.POST("/queries")            //add a query
		api.GET("/queries/:queryId")    //get query_id's query detail
		api.POST("/queries/:queryId")   //update query_id's query
		api.DELETE("/queries/:queryId") //delete query_id's query
		//query_result
		api.GET("/query_results/:queryResultId") //get query_result_id's result
		api.POST("/query_results")               //execute a unsaved query
		//job
		api.GET("/jobs/:jobId")    //get a job's result
		api.DELETE("/jobs/:jobId") //delete a job
		//visualization
		api.POST("/visualizations")                    //add a new visualization
		api.POST("/visualizations/:visualizationId")   //update a visualization
		api.DELETE("/visualizations/:visualizationId") //delete a visualization
		//widget
		api.POST("/widgets")             //add a widget
		api.POST("/widgets/:widgetId")   //update a widget
		api.POST("/widgets/batch")       // batch update widgets
		api.DELETE("/widgets/:widgetId") //delete a widget
		//dashboard
		api.POST("/dashboards")                //add a dashboard
		api.GET("/dashboards")                 //get all dashboards
		api.GET("/dashboards/:dashboardId")    //get a dashboard
		api.POST("/dashboards/:dashboardId")   //update a dashboard
		api.DELETE("/dashboards/:dashboardId") //delete a dashboard
	}
}
