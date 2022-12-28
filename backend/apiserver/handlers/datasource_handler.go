package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"seed/backend/apiserver/models"
	"seed/backend/apiserver/serializers"
	"seed/backend/cfg"
	"seed/backend/driver"
	"seed/backend/internal/log"
	"strconv"
)

type DataSourceParam struct {
	Nmae    string                 `json:"name" binding:"required"`
	DsType  string                 `json:"type" binding:"required"`
	Options map[string]interface{} `json:"options" binding:"required"`
}

type DataSourceHandler struct {
	pgDriver   *driver.PostgreSql
	jobHandler *JobHandler
}

func NewDataSourceHandler(pg *driver.PostgreSql, jh *JobHandler) *DataSourceHandler {
	return &DataSourceHandler{
		pgDriver:   pg,
		jobHandler: jh,
	}
}

func (dsh *DataSourceHandler) Add(context *gin.Context) {
	var dsParam DataSourceParam
	if err := context.ShouldBindJSON(&dsParam); err != nil {
		log.Warnf("request err:%v", err)
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasource param should be json format. please check datasource param")
		return
	}
	log.Infof("will add datasource,param: %+v", dsParam)
	optionsStr, err := json.Marshal(dsParam.Options)
	if err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasource param should be json format. please check datasource options")
		return
	}

	dataSource := models.DataSource{Name: dsParam.Nmae, DsType: dsParam.DsType, Options: string(optionsStr)}
	result := dsh.pgDriver.DB.Create(&dataSource)
	if result.Error != nil {
		log.Errorf("add a datasource error:%s", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.ServerError, "server error, please connect manager to check the system")
		return
	}
	data, err := serializers.DataSourceSerialize(&dataSource)
	if err != nil {
		log.Errorf("add a datasource error:%s", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.ServerError, "server error, please connect manager to check the system")
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    cfg.Success,
		"message": "",
		"data":    data,
	})
}

// Get one datasource
func (dsh *DataSourceHandler) Get(context *gin.Context) {
	var ds models.DataSource
	var id int
	var err error
	id, err = strconv.Atoi(context.Param("datasourceId"))
	if err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasourceId is illegal. please check param")
		return
	}
	log.Infof("get datasource id:%d", id)
	result := dsh.pgDriver.DB.First(&ds, id)
	if result.Error != nil {
		log.Warnf("get a datasource error:%v", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.ServerError, "server err, datasource query error")
		return
	}
	data, err := serializers.DataSourceSerialize(&ds)
	if err != nil {
		log.Warnf("datasource serialize error:%v", err)
		Abort(context, http.StatusInternalServerError, cfg.DataSourceSerializeError, "server err, datasource serialize error")
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    cfg.Success,
		"message": "",
		"data":    data,
	})
}

func (dsh *DataSourceHandler) GetAll(context *gin.Context) {
	dss := make([]models.DataSource, 0)
	result := dsh.pgDriver.DB.Find(&dss)
	if result.Error != nil {
		log.Warnf("get all datasource result:%v", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.ServerError, "server error, datasource query error")
		return
	}
	data, err := serializers.DataSourceBatchSerialize(dss)
	if err != nil {
		log.Warnf("datasource serialize error:%v", err)
		Abort(context, http.StatusInternalServerError, cfg.DataSourceSerializeError, "server err, datasource serialize error")
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    cfg.Success,
		"message": "",
		"data":    data,
	})
}

func (dsh *DataSourceHandler) Update(context *gin.Context) {
	dsMap := make(map[string]interface{})
	var ds models.DataSource
	var id int
	var err error
	id, err = strconv.Atoi(context.Param("datasourceId"))
	if err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasourceId is illegal. please check param")
		return
	}
	log.Infof("get datasource id:%d", id)
	ds.Id = id

	if err := context.ShouldBindJSON(&dsMap); err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasource param should be json format. please check datasource param")
		return
	}
	log.Infof("will add datasource,param: %+v", dsMap)
	//options可能不存在
	options, ok := dsMap["options"]
	if !ok {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    cfg.DataSourceParamError,
			"message": "datasource options should be exists. please check datasource options",
		})
		return
	}
	optionsStr, err := json.Marshal(options)
	if err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasource param should be json format. please check datasource options")
		return
	}
	dsMap["options"] = optionsStr
	result := dsh.pgDriver.DB.Model(&ds).Updates(dsMap)
	if result.Error != nil {
		log.Errorf("update datasource error:%s", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.ServerError, "server error, please connect manager to check the system")
		return
	}
	data, err := serializers.DataSourceSerialize(&ds)
	if err != nil {
		log.Errorf("add a datasource error:%s", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.ServerError, "server error, please connect manager to check the system")
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    cfg.Success,
		"message": "",
		"data":    data,
	})
}

//Delete a datasource
func (dsh *DataSourceHandler) Delete(context *gin.Context) {
	var ds models.DataSource
	var id int
	var err error
	id, err = strconv.Atoi(context.Param("datasourceId"))
	if err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasourceId is illegal. please check param")
		return
	}
	log.Infof("get datasource id:%d", id)
	//更新pg数据库对应datasource_id的数据
	ds.Id = id
	result := dsh.pgDriver.DB.Delete(&ds)
	if result.Error != nil {
		log.Warnf("update a datasource result:%v", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.DataSourceSerializeError, "server error, datasource update error")
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    cfg.Success,
		"message": "",
	})
}

func (dsh *DataSourceHandler) Schema(context *gin.Context) {
	var ds models.DataSource
	var id int
	var err error
	id, err = strconv.Atoi(context.Param("datasourceId"))
	if err != nil {
		Abort(context, http.StatusBadRequest, cfg.DataSourceParamError, "datasourceId is illegal. please check param")
		return
	}
	log.Infof("get datasource id:%d", id)
	//更新pg数据库对应datasource_id的数据
	ds.Id = id
	result := dsh.pgDriver.DB.First(&ds, id)
	if result.Error != nil {
		log.Warnf("get a datasource result:%v", result.Error)
		Abort(context, http.StatusInternalServerError, cfg.DataSourceSerializeError, "server error, datasource query error")
		return
	}
	//send a job to redis queue
	dsh.jobHandler.GetSchemaJobSend(int64(id), ds.Options)
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    cfg.Success,
		"message": "",
	})
}
