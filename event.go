package common

import "fmt"

const (
	EVENT_DEVICE_STATUS     = "zc.deviceStatus"
	EVENT_DEVICE_MANAGE     = "zc.deviceManage"
	EVENT_DEVICE_MESSAGE    = "zc.deviceMessage" //废弃，拆成EVENT_DEVICE_REPORT和EVENT_DEVICE_CONTROL
	EVENT_DEVICE_REPORT     = "zc.deviceReport"
	EVENT_DEVICE_DEBUG_LOG  = "zc.deviceDebugLog" //设备调试
	EVENT_DEVICE_CONTROL    = "zc.deviceControl"
	EVENT_DEVICE_FAULT      = "zc.deviceFault"
	EVENT_DEVICE_OTA        = "zc.deviceOta"
	EVENT_DEVICE_TASK       = "zc.deviceTask"
	EVENT_DEVICE_RULE_ALERT = "zc.deviceRuleAlert" //　设备SaaS组件规则触发的报警
	EVENT_APP_APM           = "terminal.event"     // 目前，所有的移动端性能评估日志都暂定使用同一个topic
	EVENT_ACCOUNT_PROFILE   = "account.profile"
)

const (
	EVENT_STATUS_ONLINE     = "online"
	EVENT_STATUS_OFFLINE    = "offline"
	EVENT_STATUS_ACTIVE     = "activate"
	EVENT_STATUS_DEACTIVE   = "deactivate"
	EVENT_STATUS_HEARTBEAT  = "heartbeat"
	EVENT_STATUS_SWITCH_ON  = "switchOn"
	EVENT_STATUS_SWITCH_OFF = "switchOff"

	EVENT_MESSAGE_CONTROL = "control"
	EVENT_MESSAGE_REPORT  = "report"
	EVENT_MESSAGE_FAULT   = "fault"
	EVENT_MESSAGE_ALARM   = "alarm"

	EVENT_OTA_UPGRADE  = "upgrade"
	EVENT_OTA_DOWNLOAD = "downLoad"
	EVENT_OTA_CONFIRM  = "confirm"
	EVENT_OTA_FINISH   = "result"

	EVENT_TASK_CLOUD  = "cloudTask"
	EVENT_TASK_ADD    = "addDeviceTask"
	EVENT_TASK_DELETE = "deleteDeviceTask"
	EVENT_TASK_NTP    = "ntp"
	EVENT_TASK_DEVICE = "deviceTask"
)

const (
	MESSAGE_SOURCE_APP     = "app"
	MESSAGE_SOURCE_WEIXIN  = "weixin"
	MESSAGE_SOURCE_LIEBAO  = "liebao"
	MESSAGE_SOURCE_SUNNING = "suning"
	MESSAGE_SOURCE_SYSTEM  = "system"
	MESSAGE_SOURCE_PANEL   = "panel"
	MESSAGE_SOURCE_DEVICE  = "device"

	SCENE_LOCAL  = "local"
	SCENE_REMOTE = "remote"
)

const (
	DEVICE_OFFLINE_SELF_CLOSE     int64 = 1
	DEVICE_OFFLINE_LOST_HEARTBEAT int64 = 2
	DEVICE_OFFLINE_ERROR_MODULE   int64 = 3
	DEVICE_OFFLINE_DECRYPT_ERROR  int64 = 4
	DEVICE_OFFLINE_KILL_DUPLICATE int64 = 5
	DEVICE_OFFLINE_CLOUD_CLOSE    int64 = 6
	DEVICE_OFFLINE_UNKNOWN        int64 = 0xff
)

const (
	ACCOUNT_LOGIN                 int64 = 1
	ACCOUNT_REGISTER              int64 = 2
	ACCOUNT_PROFILE               int64 = 3
	ACCOUNT_UPDATE_ACCESS_TOKEN   int64 = 4
	ACCOUNT_BIND_WITH_OPEN_ID     int64 = 5
	ACCOUNT_UN_BIND_WITH_OPEN_ID  int64 = 6
	ACCOUNT_REGISTER_WITH_OPEN_ID int64 = 7
	ACCOUNT_SEND_VERIFY_CODE      int64 = 8
)

const (
	USER_ROLE_OWNER     int64 = 0
	USER_ROLE_USER      int64 = 1
	USER_ROLE_SYSTEM    int64 = 2
	USER_ROLE_DEVELOPER int64 = 3
	USER_ACT_BIND       int64 = 0
	USER_ACT_UNBIND     int64 = 1
)

type ProductInfo struct {
	Name          string `json:"type"`                    //产品名称
	Serial        string `json:"serial"`                  //产品型号
	Communication string `json:"communication,omitempty"` // communication protocol /wifi/bluetooth/ethernet/cellular
}

type Firmware struct {
	ModVersion string `json:"wifi_version,omitempty"`
	DevVersion string `json:"mcu_version,omitempty"`
}

type DeviceActive struct {
	Vendor   string      `json:"vendor"` //厂商
	Product  ProductInfo `json:"product"`
	Firmware Firmware    `json:"firmware"`
}

type DeviceOnline struct {
	Port     string   `json:"port"`
	Firmware Firmware `json:"firmware"`
}

type DeviceOffline struct {
	Reason int64 `json:"reason"`
}

type DeviceHeartbeat struct {
	Operation
}

type WifiInfo struct {
	Access int64  `json:"access,omitempty"` // 0:周围热点 1：接入热点
	Mac    string `json:"mac,omitempty"`
	Ssid   string `json:"ssid,omitempty"`
	Signal string `json:"signal,omitempty"`
}

type DeviceModBase struct {
	ModType string     `json:"modType,omitempty"`
	LocalIp string     `json:"LocalIp,omitempty"`
	Sdk     string     `json:"sdk,omitempty"` //SDK版本
	Wifis   []WifiInfo `json:"wifis,omitempty"`
}

////设备状态
type EventDeviceStatus struct {
	RemoteIp         string `json:"ip,omitempty"`
	GatewaySubDomain string `json:"gatewaySubDomain,omitempty"`
	GatewayDeviceId  string `json:"gatewayDeviceId,omitempty"`
	EventType        string `json:"eventType"` //1:设备上线(online) 2：设备激活(activate)
	//3：设备心跳(heartbeat) 4:设备下线(offline)
	Timestamp string      `json:"timestamp"`
	Online    interface{} `json:"online,omitempty"`    //DeviceOnline
	Activate  interface{} `json:"activate,omitempty"`  //DeviceActive
	Offline   interface{} `json:"offline,omitempty"`   //DeviceOffline
	Heartbeat interface{} `json:"heartbeat,omitempty"` //DeviceHeartbeat
}

type OpResult struct {
	Success int64  `json:"success"`          // 0:成功
	Reason  string `json:"reason,omitempty"` //失败原因
}

func GetOpResult(err error) OpResult {
	if err == nil {
		return OpResult{
			Success: 0,
		}
	} else if acErr, ok := err.(*ACError); ok {
		return OpResult{
			Success: acErr.Code(),
			Reason:  acErr.Error(),
		}
	} else if acNewErr, ok := err.(*ACNewError); ok {
		return OpResult{
			Success: acNewErr.Basic().Code(),
			Reason:  acNewErr.Desc(),
		}
	} else {
		// 返回内部错误
		return OpResult{
			Success: ErrInternalError.Code(),
			Reason:  fmt.Sprintf("%v", err),
		}
	}
}

////设备绑定解绑管理
type EventDeviceManage struct {
	Timestamp string      `json:"timestamp"`
	Role      int64       `json:"userType"`         //用户角色:0:管理员; 1:普通用户; 2:系统
	Action    int64       `json:"status"`           //动作:1:绑定; 2:解绑; 3:强制解绑
	Result    interface{} `json:"result,omitempty"` //OpResult 执行结果,成功，失败+失败原因
}

type Application struct {
	Os      string `json:"os,omitempty"` //android/ios/device/system
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Operation struct {
	AppType      string      `json:"type"`     //app/weixin/liebao/suning/panel/device
	Scenario     string      `json:"scenario"` //local/remote
	RequestCode  int64       `json:"requestCode,omitempty"`
	RequestId    int64       `json:"requestId,omitempty"`
	ResponseCode int64       `json:"responseCode,omitempty"`
	ResponseId   int64       `json:"responseId,omitempty"`
	Description  string      `json:"description,omitempty"`
	Result       interface{} `json:"result,omitempty"` //OpResult
	Request      interface{} `json:"request,omitempty"`
	Response     interface{} `json:"response,omitempty"`
}

////设备控制，上报
type EventDeviceMessage struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"eventType"` //1:设备控制(control) 2:设备上报(report)
	//3:设备故障(fault) 4:设备报警(alarm)
	App interface{} `json:"app,omitempty"` //Application
	Op  interface{} `json:"op,omitempty"`  //Operation
}

type FileMeta struct {
	FileType    string `json:"type"`
	Checksum    string `json:"checksum"`
	FileVersion string `json:"version"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
}

type DeviceUpgrade struct {
	CurrentVersion string     `json:"currentVersion"`
	TargetVersion  string     `json:"targetVersion"`
	fileMeta       []FileMeta `json:"fileMeta,omitempty"`
}

type FileDownLoad struct {
	CurrentVersion string `json:"currentVersion"`
	TargetVersion  string `json:"targetVersion"`
	Result         string `json:"result,omitempty"`
}

type UserConfirm struct {
	CurrentVersion string `json:"currentVersion"`
	TargetVersion  string `json:"targetVersion"`
}

type OtaFinish struct {
	PreviousVersion string `json:"previous"`
	CurrentVersion  string `json:"current"`
}

////设备OTA
type EventDeviceOta struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"eventType"` //1:设备升级(upgrade) 2:设备OTA文件下载完成(downLoad)
	// 3:设备升级确认(confirm) 4:设备升级成功(finish)
	OtaType  int64       `json:"otaType"`
	Upgrade  interface{} `json:"upgrade,omitempty"`  //DeviceUpgrade
	DownLoad interface{} `json:"downLoad,omitempty"` //FileDownLoad
	Confirm  interface{} `json:"confirm,omitempty"`  //UserConfirm
	Finish   interface{} `json:"result,omitempty"`   //OtaFinish
}

type CloudTimerTask struct {
	Time string      `json:"time"`
	Data interface{} `json:"data,omitempty"`
}

type UpdateTimerTask struct {
	Time         int64   `json:"time"` //云端下发定时任务的时间
	TaskId       int64   `json:"taskId"`
	Code         int64   `json:"code"`         //功能码
	RunCount     int64   `json:"runCount"`     //重复执行次数
	SubTaskCount int64   `json:"subTaskCount"` //子任务个数
	Cycle        int64   `json:"cycle"`        //任务周期
	Times        []int64 `json:"times"`        //各个子任务的执行时间
}

type DeviceNtp struct {
	DeviceLocalTime int64 `json:"deviceLocalTime,omitempty"` //本地维护的绝对秒
	DeviceNtpTime   int64 `json:"deviceNtpTime,omitempty"`   //获取到的NTP时间
}

type DeviceTimerTask struct {
	DeviceLocalTime   int64  `json:"deviceLocalTime,omitempty"`   //本地维护的绝对秒
	DeviceTaskTime    int64  `json:"deviceTaskTime,omitempty"`    //定时任务时间点
	DeviceSequenceId  int64  `json:"deviceSequenceId,omitempty"`  //设备上的组件号
	DeviceTimerAction string `json:"deviceTimerAction,omitempty"` //定时行为
	DeviceTimerCycle  string `json:"deviceTimerCycle,omitempty"`  //定时周期
}

////设备定时
type EventDeviceTimer struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"eventType"` //1:云端定时(cloudTask) 2:增加设备定时任务(addDeviceTask)
	//3:删除设备定时任务(deleteDeviceTask) 4:设备NTP(ntp) 5:设备定时成功(deviceTask)
	CloudTask  interface{} `json:"cloudTask,omitempty"`  //CloudTimerTask
	Update     interface{} `json:"update,omitempty"`     //UpdateTimerTask
	Ntp        interface{} `json:"ntp,omitempty"`        //DeviceNtp
	DeviceTask interface{} `json:"deviceTask,omitempty"` //DeviceTimerTask
}

////设备控制
type EventDeviceControl struct {
	Timestamp string      `json:"timestamp"`
	App       interface{} `json:"app,omitempty"` //Application
	Op        interface{} `json:"op,omitempty"`  //Operation
}

////设备上报
type EventDeviceReport struct {
	Timestamp string      `json:"timestamp"`
	App       interface{} `json:"app,omitempty"` //Application
	Op        interface{} `json:"op,omitempty"`  //Operation
}

// 设备故障事件
type EventDeviceFault struct {
	Timestamp     string `json:"timestamp"`   // 时间戳："2006-01-02 15:04:05.000000-0700"
	FaultCategory string `json:"fault_type"`  // 故障属性的 identification
	FaultValue    string `json:"fault_code"`  // 故障值（整型故障码=>字符串）
	FaultOccur    int64  `json:"fault_occur"` // 是否是首次出现（least significant bit 的比特位意义：LSB0 是否为首次故障; LSB1 是否是模块首次故障）
}

// 设备维度的规则触发的报警
// event topic: EVENT_DEVICE_RULE_ALERT
type EventDeviceRuleTriggeredAlert struct {
	Timestamp string `json:"timestamp"` // 时间戳："2006-01-02 15:04:05.000000-0700"
	Title     string `json:"title"`     // 报警标题
	Content   string `json:"content"`   // 报警内容
	RuleLHS   string `json:"ruleLHS"`   // 触发本次报警的规则
}

// 用户事件日志
type EventAccountProfile struct {
	Timestamp string      `json:"timestamp"` // 时间戳："2006-01-02 15:04:05.000000-0700"
	ExtTime   string      `json:"extTime"`
	EventType int64       `json:"eventType"`        // 1:注册 2：登录， 3：设置属性..
	Ip        string      `json:"ip"`               // ip地址
	Finish    interface{} `json:"result,omitempty"` // 成功
	Account   interface{} `json:"account"`          // 日志内容
	Phone     string      `json:"phone"`            // 手机号
	Email     string      `json:"email"`            // 邮箱
	OpenId    string      `json:"openId"`           // openId
}
