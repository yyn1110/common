package common

import "time"

const (
	MAJOR_DOMAIN_ID = "majorDomainId"
	SUB_DOMAIN_ID   = "subDomainId"
	TIMESTAMP       = "timestamp"
)

// ac 内部使用的主域id,各个环境都一样
const (
	AC_MAJOR_DOMAIN_ID = int64(371)
)

const (
	EVENT_META = "eventMeta"
	EVENT_DATA = "eventData"
	EVENT_TYPE = "eventType"
)

const (
	OLD_TOKEN_VERSION = 0
	NEW_TOKEN_VERSION = 1
)

const (
	APPWEB_DEVICE   = 0
	PHYSICAL_DEVICE = 1
)

const (
	MILLISECONDS_PER_SECOND = int(time.Second / time.Millisecond)
	SECONDS_PER_MINUTE      = int(time.Minute / time.Second)
	SECONDS_PER_HOUR        = int(time.Hour / time.Second)
	MINUTES_PER_HOUR        = int(time.Hour / time.Minute)
)

//account json key
const (
	ACCOUNT_ATTRIBUTES       = "account"
	UID                      = "uid"
	UID_CRC                  = "uid_CRC"
	DID_CRC                  = "did_CRC"
	PID_CRC                  = "pid_CRC"
	PID                      = "pid"
	DID                      = "did"
	ID                       = "id"
	ID_CRC                   = "id_CRC"
	CRC_SUFFIX               = "CRC"
	REGISTER                 = "register"
	USER_REGISTER_COLLECTION = "user_register"
)

const (
	MAX_DOMAIN_ID     = int64((1 << 48) - 1) // at most 6 bytes for major domain ID
	MAX_SUB_DOMAIN_ID = int64((1 << 16) - 1) // at most 2 bytes for sub domain ID
)

// device json key
const (
	BIRTHDAY           = "birthday"
	USER_ID            = "userId"
	PHYSICAL_ID        = "physicalId"
	VENDOR             = "vendor"
	ACTIVATE_SEGMENT   = "activate"
	PRODUCT            = "product"
	FIRMWARE           = "firmware"
	IP                 = "ip"
	OP                 = "op"
	USER_TYPE          = "userType"
	DEVICE_BIND_CODE   = 0
	DEVICE_UNBIND_CODE = 1

	PREVIOUS     = "previous"
	CURRENT      = "current"
	STATUS       = "status"
	CHANNEL      = "channel"
	TYPE         = "type"
	SCENARIO     = "scenario"
	// 下面的是device debug log 用到的
	REQUEST      = "request"      //设备上报的消息内容
	REQUESTCODE  = "requestCode"  //设备上报时的msgCode,用于区分不同的上报消息类型
	REBOOTMSGLIST= "RebootMsgList" // 重启列表
	OFFLINEMSGLIST = "OfflineMsgList"
	DEVICE_TIMESTAMP = "Timestamp"
	REBOOT_REASON = "RebootReason"
	OFFLINE_REASON = "OfflineReason"
	EXTRAINFO    = "ExtraInfo"


	COMMAND      = "cmd"
	APP          = "app"
	OS           = "os"
	VERSION      = "version"
	OTA_TYPE     = "otaType"
	OTA          = "ota"
	NAME         = "name"
	SERIAL       = "serial"
	MCU_VERSION  = "mcu_version"
	WIFI_VERSION = "wifi_version"
	SEPARATOR    = "."
	LOCATION     = "location"
	COUNTRY      = "country"
	PROVINCE     = "province"
	CITY         = "city"
	// 增加device debug log request字段
	HEADER		 = "Header"
	DEBUG_TYPE   = "Typ"
	CONTENT      = "Content"
	MSG_TYPE	 = "Typ"
	WIFI_RSSI	 = "RSSI"
	MAC          = "MAC"
	LAN_IP		 = "LanIP"
	SSID		 = "SSID"

	// 增加故障类字段
	DEVICE_FAULT_TYPE  = "fault_type"
	DEVICE_FAULT_CODE  = "fault_code"
	DEVICE_FAULT_OCCUR = "fault_occur"

	DEVICE_FAULT_TYPE_OVERALL = "_overall"

	DEVICE_ACTIVATE_COLLECTION = "device_activate"
	DEVICE_BIND_COLLECTION     = "device_bind"
	DEVICE_OPERATE_COLLECTION  = "device_operate"
	DEVICE_OTA_COLLECTION      = "device_ota"
	DEVICE_FAULT_COLLECTION    = "device_fault" //设备故障
	DEVICE_WIFI_LOG_COLLECTION = "device_wifi_log"   //设备上报debug log
	DEVICE_REBOOT_OFFLINE_COLLECTION = "device_reboot_offline"   //设备上报debug log
)

var STD_TIME_FORMAT = "2006-01-02 15:04:05"
