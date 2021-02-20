package common

// 监控类型
type MonitorType int64
const (
	MONITOR_DOMAIN				MonitorType	= 1  	// 获取主域监控信息
	MONITOR_SUBDOMAIN           MonitorType = 2     // 获取子域监控信息
	MONITOR_SUBDOMAIN_BY_DOMAIN MonitorType = 3     // 获取主域下所有子域的监控信息,如果传入此类型，则scope参数必须填上主域id.可以多个.
	MONITOR_OTHER               MonitorType = 4    	// 其他没有主子域信息的监控,可扩展
	MONITOR_ALL                 MonitorType = 65535 // 所有类型的全部监控信息
)

// 监控数据源类型
type DataSourceType string
const (
	DATA_SOURCE_COUNTER 		DataSourceType = "COUNTER"  	// 一直递增的监控项数据,eg:网络流量,访问量.实际存入的值是(当前的值-上个周期的值)
	DATA_SOURCE_GAUGE 			DataSourceType = "GAUGE"  		// 原始监控数据
	DATA_SOURCE_DERIVE 			DataSourceType = "DERIVE"
	DATA_SOURCE_ABSOLUTE		DataSourceType = "ABSOLUTE"
	DATA_SOURCE_COMPUTE 		DataSourceType = "COMPUTE"
)

type MonitorResult struct {
	Metric      string 			`json:"metric"`               // 监控项名称
	Endpoint    string 			`json:"endpoint"`             // 主机名称,和open-falcon上模块绑定的机器名一致即可
	DSType 		DataSourceType  `json:"counterType"`		  // 数据源类型
	Value      	int64 			`json:"value"`   	          // 监控项的值
	Tags        string 			`json:"tags"`                 // 子域级别的监控tag信息里面包含其所属主域信息
}