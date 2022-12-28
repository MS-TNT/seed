package serializers

import (
	"encoding/json"
	"seed/backend/apiserver/models"
	"seed/backend/internal/log"
)

func DataSourceSerialize(dataSource *models.DataSource) (map[string]interface{}, error) {
	dsMap := make(map[string]interface{})
	dsMap["id"] = dataSource.Id
	dsMap["name"] = dataSource.Name
	dsMap["type"] = dataSource.DsType
	optionsMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(dataSource.Options), &optionsMap); err != nil {
		log.Errorf("datasource options unmarshal err: %s, options:%s", err.Error(), dataSource.Options)
		return nil, err
	}
	dsMap["options"] = optionsMap
	return dsMap, nil
}

func DataSourceBatchSerialize(dataSources []models.DataSource) ([]map[string]interface{}, error) {
	dssSlice := make([]map[string]interface{}, 0)
	for _, ds := range dataSources {
		dsMap, err := DataSourceSerialize(&ds)
		if err != nil {
			return nil, err
		}
		dssSlice = append(dssSlice, dsMap)
	}
	return dssSlice, nil
}
