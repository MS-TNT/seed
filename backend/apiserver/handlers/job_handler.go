package handlers

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/example/tracers"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/spf13/viper"
	"seed/backend/internal/log"
	"time"
)

type JobHandler struct {
	mServer *machinery.Server
}

func NewJobHandler() (*JobHandler, error) {
	queueName := viper.GetString("worker.queueName")
	resultsExpireInSec := viper.GetInt("worker.resultsExpireInSec")
	brokerAddr := viper.GetString("worker.brokerAddr")
	resultBackendAddr := viper.GetString("worker.resultBackendAddr")
	cnf := &config.Config{
		DefaultQueue:    queueName,
		ResultsExpireIn: resultsExpireInSec,
		Broker:          brokerAddr,
		ResultBackend:   resultBackendAddr,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}
	return &JobHandler{mServer: server}, nil
}

func (jh *JobHandler) GetSchemaJobSend(datasourceId int64, options string) {
	cleanup, err := tracers.SetupTracer("sender")
	if err != nil {
		log.Errorf("Unable to instantiate a tracer:", err)
		return
	}
	defer cleanup()

	getSchemaTask := tasks.Signature{
		UUID: fmt.Sprintf("dataSource:schema:%d", datasourceId),
		Name: "getSchema",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: datasourceId,
			},
			{
				Type:  "string",
				Value: options,
			},
		},
	}
	context := context.Background()
	asyncResult, err := jh.mServer.SendTaskWithContext(context, &getSchemaTask)
	if err != nil {
		log.Errorf("Could not send task: %s", err.Error())
		return
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		log.Errorf("Getting task result failed with error: %s", err.Error())
		return
	}
	log.Infof("1 + 1 = %v\n", tasks.HumanReadableResults(results))
	return
}
