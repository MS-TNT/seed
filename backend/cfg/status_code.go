package cfg

//状态码
const (
	Success     = 0
	ServerError = 9999 //服务器内部错误
	//数据源
	DataSourceParamError     = 2201 //参数错误
	DataSourceNotExist       = 2202 //数据源不存在
	DataSourceSerializeError = 2203 //数据源序列化失败
)
