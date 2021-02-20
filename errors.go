package common

import (
	"fmt"
)

type ACError struct {
	errCode int64
	errMsg  string
}

func (this *ACError) Code() int64 {
	return this.errCode
}

// WARNING:for ablecloud inner system error
func newError(code int64, msg string) *ACError {
	err := &ACError{errCode: code, errMsg: msg}
	if errorMap[code] != nil {
		str := fmt.Sprintf("find message code definition duplicated: %d", code)
		panic(str)
	}
	errorMap[code] = err
	return err
}

// WARNING:just for outer system error
func NewError(code int64, msg string) *ACError {
	return newError(code, msg)
}

func (this ACError) Error() string {
	return fmt.Sprintf("%s", this.errMsg)
}

type ACNewError struct {
	basic *ACError
	desc  string
}

// 获取基础错误。
// 返回基础错误、是否获取到。
func GetBasicError(err error) (*ACError, bool) {
	if e, ok := err.(*ACError); ok {
		return e, true
	}
	if e, ok := err.(*ACNewError); ok {
		return e.basic, true
	}
	return nil, false
}

func ACErrorf(err error, format string, a ...interface{}) *ACNewError {
	if e, ok := err.(*ACError); ok {
		return NewACNewError(e, fmt.Sprintf(format, a...))
	}
	return NewACNewError(ErrInternalError, fmt.Sprintf(format, a...))
}

func NewACNewError(basic *ACError, desc string) *ACNewError {
	return &ACNewError{basic, desc}
}

func (this *ACNewError) Basic() *ACError {
	return this.basic
}

func (this *ACNewError) Desc() string {
	return this.desc
}

func (this *ACNewError) Error() string {
	if this.basic != nil {
		return fmt.Sprintf("%d:%s(%s)", this.basic.errCode, this.basic.errMsg, this.desc)
	} else {
		return this.desc
	}
}

var errorMap = make(map[int64]*ACError)

///////////////////////////////////////////////////////////
// inner sepecial error
///////////////////////////////////////////////////////////
// if get error by code not find return this error
var ErrInvalidError = newError(2000, "invalid error")

// if not ACError return this error for resp.getACError()
var ErrInternalError = newError(3000, "internal error")

///////////////////////////////////////////////////////////
// 3001-3500 common error code
///////////////////////////////////////////////////////////
var (
	ErrInvalidHeader       = newError(3001, "invalid header")
	ErrInvalidParam        = newError(3002, "invalid param")
	ErrNotSupported        = newError(3003, "not supported")
	ErrNotAllowed          = newError(3004, "not allowed")
	ErrNoPrivilege         = newError(3005, "no privelige")
	ErrInvalidURI          = newError(3006, "invalid request uri")
	ErrMajorDomainNotExist = newError(3007, "major domain not exist")
	ErrSubDomainNotExist   = newError(3008, "sub domain not exist")
	ErrServiceNotExist     = newError(3009, "service not exist")
	ErrMsgNotSupported     = newError(3010, "message not supported")
	ErrServiceDown         = newError(3011, "service not available")
	ErrTimeout             = newError(3012, "request timeout")
	ErrNetworkError        = newError(3013, "network error")
	ErrSignExpires         = newError(3014, "signature expires")
	ErrInvalidSign         = newError(3015, "invalid signature")
	ErrNotImplemented      = newError(3016, "not implemented")
	ErrHttpError           = newError(3017, "http error")
	ErrServiceError        = newError(3018, "service error")
	ErrUdsException        = newError(3019, "uds throw exception")
	ErrServiceHandleFailed = newError(3020, "service handle msg failed")
	ErrInvalidRpcRespType  = newError(3021, "invalid rpc resp type(ack/error)")
	//<--- error codes for request context check - start from 3050
	ErrReqCtxCheckNullCtx      = newError(3050, "request context check - context is null")
	ErrReqCtxCheckType         = newError(3051, "request context check wrong checkType")
	ErrReqCtxCheckMajorDomain  = newError(3052, "request context check error majorDomain")
	ErrReqCtxCheckDeveloperId  = newError(3053, "request context check error developerId")
	ErrReqCtxCheckInnerService = newError(3054, "request context check error innerService")
	ErrReqCtxCheckDevOrInner   = newError(3055, "request context check error developerId_or_innerService")
	ErrReqCtxCheckSubDomain    = newError(3056, "request context check error subDomain")
	//--->
)

//////////////////////////////////////////////////////////
// 3501-3600 account related error code
///////////////////////////////////////////////////////////
var (
	ErrAccountNotExist          = newError(3501, "account not exist")
	ErrAccountExist             = newError(3502, "account already exist")
	ErrInvalidName              = newError(3503, "invalid name")
	ErrWrongPassword            = newError(3504, "password wrong")
	ErrInvalidVerifyCode        = newError(3505, "invalid verify code")
	ErrVerifyCodeExpires        = newError(3506, "verify code expires")
	ErrInvalidEmail             = newError(3507, "invalid email address")
	ErrInvalidPhone             = newError(3508, "invalid phone number")
	ErrInvalidAccountStatus     = newError(3509, "invalid account status")
	ErrAccountAlreadyBound      = newError(3510, "account already bound")
	ErrVerifyOauthFailed        = newError(3511, "verify oauth failed")
	ErrInvalidUserProfile       = newError(3512, "invalid user extend profile")
	ErrInvalidAccessToken       = newError(3513, "invalid access token")
	ErrInvalidRefreshToken      = newError(3514, "invalid refresh token")
	ErrAccessTokenExpire        = newError(3515, "access token expire")
	ErrRefreshTokenExpire       = newError(3516, "refresh token expire")
	ErrExtendColumnAlreadyExist = newError(3517, "extend column is already exist")
	ErrExtendColumnNotExist     = newError(3518, "extend column is not exist")
	ErrAccountNumExceedMax      = newError(3519, "account num exceed max")
	ErrProfileNameIsReserved    = newError(3520, "profile name is reserved")
)

//////////////////////////////////////////////////////////
// 3601-3800 group related error code
//////////////////////////////////////////////////////////
var (
	ErrGroupNotExist       = newError(3601, "group not exist")
	ErrGroupExist          = newError(3602, "group already exist")
	ErrInvalidGroupStatus  = newError(3603, "invalid group status")
	ErrMemberNotExist      = newError(3604, "member not exist")
	ErrMemberExist         = newError(3605, "member already exist")
	ErrInvalidMemberStatus = newError(3606, "invalid member status")
)

//////////////////////////////////////////////////////////
// 3801-3900 device related error code
//////////////////////////////////////////////////////////

var (
	ErrInvalidMsgCode           = newError(3801, "invalid message code")
	ErrDeviceNotExist           = newError(3802, "device not exist")
	ErrDeviceExist              = newError(3803, "device already exit")
	ErrInvalidDevice            = newError(3804, "invalid device")
	ErrBindCodeExpires          = newError(3805, "bind code expires")
	ErrInvalidBindCode          = newError(3806, "invalid bind code")
	ErrDeviceOffline            = newError(3807, "device offline")
	ErrMasterNotExist           = newError(3808, "master device not exist")
	ErrDeviceIsMaster           = newError(3809, "device is master")
	ErrDeviceIsSlave            = newError(3810, "device is slave")
	ErrAlreadyBound             = newError(3811, "device already bound")
	ErrNotBound                 = newError(3812, "device not bound")
	ErrInvalidDeviceStatus      = newError(3813, "invalid device status")
	ErrDeviceTimeout            = newError(3814, "device response timeout")
	ErrShareCodeNotExist        = newError(3815, "share code not exist")
	ErrInvalidShareCode         = newError(3816, "invalid share code")
	ErrShareCodeExpires         = newError(3817, "share code expires")
	ErrOutOfValidity            = newError(3818, "bind device out of validity")
	ErrOwnerNotExist            = newError(3819, "owner not exist")
	ErrGatewayNotMatch          = newError(3820, "gateway not match")
	ErrOwnerNotMatch            = newError(3821, "owner not match")
	ErrDeviceNotActivated       = newError(3822, "device not activated")
	ErrDeviceBusy               = newError(3823, "device busy")
	ErrDeviceResponse           = newError(3824, "device response error")
	ErrSessionIdMismatch        = newError(3825, "device connection's session id mismatch")
	ErrDeviceSubDomainMissMatch = newError(3826, "device's sub domain dose not match")
	//<--- device license related errors (3850 ~ 3865)
	ErrLicenseOutOfQuota            = newError(3850, "license out of quota")
	ErrLicenseNotEnabled            = newError(3851, "this product did not enable license")
	ErrLicenseProductQuotaDuplicate = newError(3852, "this product has already has quota info exists")

	ErrLicenseNotGenerated   = newError(3854, "use license which has not been generated")
	ErrLicenseWrongDomain    = newError(3855, "use license with wrong domain")
	ErrLicenseConflict       = newError(3856, "license has been associated to other device")
	ErrDeviceLicenseConflict = newError(3857, "device has been assigned other license")
	ErrDeviceLicenseError    = newError(3858, "cannot control device because auth failed")
	ErrDeviceLicenseNotNull  = newError(3859, "set license failed due to this device already has license which is not NULL")
	ErrDeviceAuthFailed      = newError(3860, "auth device failed defalut reason check warehouse for detail")
	ErrAuthInsertRecord      = newError(3861, "auth device failed for insert record table")
	//--->
	ErrGetProInfo             = newError(3862, "get product info error")
	ErrStockNumForModifyPro   = newError(3863, "can not modify product license for sotck num not 0")
	ErrModifyDeviceIdConflict = newError(3864, "can not modify device id for the new device id exist already")
	/// product attribute相关error
	ErrProductAttrCommon    = newError(3865, "product attribute operation error, check log for detail")
	ErrProductAttrConflict  = newError(3866, "product attribute exist already in db")
	ErrProductAttrDuplicate = newError(3867, "product attribute duplicate from import list")
	ErrGetProductKey        = newError(3868, "get product secret key error")
	// device fault相关error
	ErrFaultAttrNotExist    = newError(3870, "fault attribute not exist")
	// 设备数据流转发的相关错误
	ErrDeviceUdsNotSet      = newError(3871, "device uds not set")
	// parsing相关error
	ErrParsingDevDataflowFailed = newError(3875, "failed to parse device's dataflow")
	ErrBinaryParserNotExist     = newError(3876, "binary parser not exist")
	ErrBinaryParserDuplicate    = newError(3877, "duplicate (domainId,subDomainId,msgCode) for binary parser")
)

//////////////////////////////////////////////////////////
// 3901-4000 storage related error code
//////////////////////////////////////////////////////////
var (
	ErrFileNotExist         = newError(3901, "file not exist")
	ErrFileExist            = newError(3902, "file already exist")
	ErrInvalidFileStatus    = newError(3903, "invalid file state")
	ErrChecksumError        = newError(3904, "file checksum error")
	ErrInvalidFileContent   = newError(3905, "invalid file content")
	ErrClassNotExist        = newError(3920, "class not exist")
	ErrClassExist           = newError(3921, "class already exist")
	ErrCheckDataError       = newError(3922, "data error")
	ErrDataNotExist         = newError(3923, "data not exist")
	ErrDataExist            = newError(3924, "data already exist")
	ErrInvalidValueParam    = newError(3925, "invalid value param")
	ErrEndOfPartition       = newError(3926, "end of partition")
	ErrInvalidFilterParam   = newError(3927, "invalid filter param")
	ErrInvalidExprParam     = newError(3928, "invalid expr param")
	ErrColumnNotExist       = newError(3929, "column not exist")
	ErrInvalidPartitionKey  = newError(3930, "invalid partiton key")
	ErrInvalidStoreIndex    = newError(3931, "invalid store index")
	ErrInvalidPrimaryKey    = newError(3932, "invalid primary key")
	ErrInvalidColumnParam   = newError(3933, "invalid column param")
	ErrInvalidAggrParam     = newError(3934, "invalid aggregate param")
	ErrInvalidGroupByParam  = newError(3935, "invalid groupBy param")
	ErrInvalidOrderByParam  = newError(3936, "invalid orderBy param")
	ErrInvalidSelectParam   = newError(3937, "invalid select param")
	ErrInvalidDataType      = newError(3938, "invalid data type")
	ErrColumnIsAlreadyExist = newError(3939, "column is already exist")
	ErrInvalidGroupName     = newError(3940, "invalid group name")
	ErrInvalidClassName     = newError(3941, "invalid class name")
	ErrInvalidKeysParam     = newError(3942, "invalid keys param")
	ErrDatabaseExecuteFailed  = newError(3943, "database execute failed")
)

//////////////////////////////////////////////////////////
// 4001-4050 notification related error code
//////////////////////////////////////////////////////////
var (
	ErrInvalidTitle         = newError(4001, "title is emtpy")
	ErrInvalidContent       = newError(4002, "invalid content")
	ErrNoAvailIosDevice     = newError(4003, "no available ios device")
	ErrNoAvailAndroidDevice = newError(4004, "no available android device")
	ErrWhiteListNotSet      = newError(4005, "white list not set")
	ErrSignFailed           = newError(4006, "sign failed, pls check ak/sk")
	ErrNotifyInfoError      = newError(4007, "notify info error, pls check notification info on the platform")
	ErrReqServiceFailed     = newError(4008, "request umeng service failed")
	ErrAppKeyNotExist       = newError(4009, "appKey not exist, pls check it on umeng platform")
	ErrProductCertNotExist  = newError(4010, "product certificate not exist, pls upload it on umeng platform")
	ErrDevCertNotExist      = newError(4011, "developer certificate not exist, pls upload it on umeng platform")
	ErrUserNotExist         = newError(4012, "userList is empty or users not exist")
)

//////////////////////////////////////////////////////////
// 4051-4100 notification related error code
//////////////////////////////////////////////////////////
var (
	ErrSMSInfoNotExist              = newError(4050, "sms info not exist")
	ErrSMSTemplateNotExist          = newError(4051, "sms template not exist")
	ErrEmailTemplateModifyForbidden = newError(4052, "email template modify forbidden")
	ErrVerifyEmailSmtpDialErr       = newError(4053, "dial smtp server failed")
	ErrVerifyEmailAuthErr           = newError(4054, "email auth failed")
	ErrSendEmailFailed              = newError(4055, "send email failed")
	ErrEmailHasInvalidChar          = newError(4056, "email template name/subject/content has invalid characters")
	ErrSendEmailNotPassed           = newError(4057, "send email template which has not been granted")
	ErrEmailTemplateNotExist        = newError(4058, "email template not exist")
)

//////////////////////////////////////////////////////////
// 4101-4200 ota related error code
//////////////////////////////////////////////////////////
var (
	ErrOtaNoNewVersion    = newError(4101, "no new version")
	ErrOtaConfirmExpire   = newError(4102, "confirm device upgrade expired")
	ErrOtaDismissExpire   = newError(4103, "dismiss device upgrade expired")
	ErrOtaDownloadOtaFile = newError(4104, "download ota file failed")
)

//////////////////////////////////////////////////////////
// 4201-4300 group related error code
//////////////////////////////////////////////////////////
var (
	ErrAlreadyUserInGroup   = newError(4201, "user already in group")
	ErrAlreadyDeviceInGroup = newError(4202, "device already in group")
	ErrAlreadyGroupInGroup  = newError(4203, "group already in group")
	ErrDeviceNotInGroup     = newError(4204, "device not in group")
	ErrUserNotInGroup       = newError(4205, "user not in group")
	ErrGroupNotInGroup      = newError(4206, "group not in group")
	ErrGatewayNotInGroup    = newError(4207, "device gateway not in group")
)

//4401-4500 used for timer task
var (
	ErrTimeHasExpired          = newError(4401, "timer task has already expired")
	ErrTaskNotExist            = newError(4402, "timer task is not exist")
	ErrTaskStatus              = newError(4403, "timer task status is invalid")
	ErrInvalidTimeZone         = newError(4404, "time zone is invalid")
	ErrInvalidTimePoint        = newError(4405, "time point is invalid")
	ErrInvalidTimeCycle        = newError(4406, "time cycle is invalid")
	ErrTaskExist               = newError(4407, "task is already exist")
	ErrTaskAlreadyStart        = newError(4408, "task is already start")
	ErrTaskAlreadyStop         = newError(4409, "task is already stop")
	ErrIsNotCloudTask          = newError(4410, "this is not a cloud task")
	ErrTaskGroupIsExist        = newError(4411, "task group is already exist")
	ErrTaskGroupIsNotExist     = newError(4412, "task group is not exist")
	ErrInvalidTaskType         = newError(4413, "invalid timer task type")
	ErrInvalidTaskCommand      = newError(4414, "invalid timer task command")
	ErrUnsupportTaskCommand    = newError(4415, "unsupport timer task command")
	ErrTaskOwnerIsInconsistent = newError(4416, "task owner is inconsistent")
)

//4501-4550 used for feedback service
var (
	ErrInvalidFeedback        = newError(4501, "invalid feedback")
	ErrFeedbackColumnNotExist = newError(4502, "feedback column is not exist")
)

//4551-4600 used for ranking service
var (
	ErrRankingSetAlreadyExist = newError(4551, "ranking set is already exist")
	ErrRankingSetNotExist     = newError(4552, "ranking set is not exist")
)

//4601-4650 used for file manager service
var (
	ErrBucketAlreadyExist = newError(4601, "bucket is already exist")
	ErrBucketNotExist     = newError(4602, "bucket is not exist")
	ErrFileAlreadyExist   = newError(5403, "file is already exist")
	ErrBucketNotEmpty     = newError(5404, "bucket is not empty")
)

//4651-4750 used for access
var (
	ErrAccessCommon         = newError(4651, "access common failed")
	ErrRoleNameDuplicate    = newError(4652, "role name exist already")
	ErrRoleNotExist         = newError(4653, "role not exist")
	ErrRoleHasUser          = newError(4654, "role has user")
	ErrDeptNameDuplicate    = newError(4655, "dept name exist already")
	ErrDeptNotExist         = newError(4656, "dept not exist")
	ErrDeptHasUser          = newError(4657, "dept has user")
	ErrDeptHasChild         = newError(4658, "dept has child dept")
	ErrUserPhoneDuplicate   = newError(4659, "user phone exist already")
	ErrUserEmailDuplicate   = newError(4660, "user email exist already")
	ErrAccessUserNotExist   = newError(4661, "user not exist")
	ErrAccountDuplicate     = newError(4662, "account duplicate")
	ErrCompanyDuplicate     = newError(4663, "company exist already in this domain")
	ErrDataEnvNotExist      = newError(4664, "data env dir not exist")
	ErrDataProductNotExist  = newError(4665, "data product not exist")
	ErrParentFnNotExist     = newError(4666, "parent fn not exist")
	ErrFnNotExist           = newError(4667, "fn not exist")
	ErrInvalidCompanyStatus = newError(4668, "invalid company status")
	ErrCompanyNotExist      = newError(4669, "company not exist")
	ErrAccountActivated     = newError(4670, "user and company activated")
)

//4751-4800 used for portal
var (
	ErrPortalCommon        = newError(4751, "portal common failed")
	ErrNewsNotExist        = newError(4752, "news not exist")
	ErrProjectCaseNotExist = newError(4753, "project case not exist")
	ErrSubsciberEmailExist = newError(4754, "subsciber email exist")
)

//4801-4850 used for oauth2
var (
	ErrThirdNameInvalid = newError(4801, "third name invalid")
	ErrOAuthCommon      = newError(4802, "aouth common failed")
	ErrClientNotExist   = newError(4803, "client not exist")
	ErrTokenNotExist    = newError(4804, "token not exist")
)

//////////////////////////////////////////////////////////
// 5001-5200 platform related error code
//////////////////////////////////////////////////////////
var (
	ErrProjectNotExist           = newError(5001, "project not exist")
	ErrMajorDomainExist          = newError(5002, "major domain already exist")
	ErrInvalidDomain             = newError(5003, "domain format invalid")
	ErrTooManyKeyPair            = newError(5004, "too many key pairs")
	ErrServiceIsOnline           = newError(5005, "service is online")
	ErrVersionNotExist           = newError(5006, "version not exist")
	ErrVersionRollback           = newError(5007, "version rollback")
	ErrTooManyVersionPublished   = newError(5008, "too many version published")
	ErrInvalidRollback           = newError(5009, "invalid rollback")
	ErrVersionCompatible         = newError(5010, "version compatible")
	ErrVersionNotCompatible      = newError(5011, "version not compatible")
	ErrInvalidMajorVersion       = newError(5012, "invalid major version")
	ErrInvalidMinorVersion       = newError(5013, "invalid minor version")
	ErrInvalidPatchVersion       = newError(5014, "invalid patch version")
	ErrInstanceNotExist          = newError(5015, "instance not exist")
	ErrInstanceExist             = newError(5016, "instance already exist")
	ErrEmptyAppFileFromBlobStore = newError(5017, "empty APP file from blobstore")
	ErrPortInUse                 = newError(5018, "port already in use")
	ErrPortNotInUse              = newError(5019, "port not in use")
	ErrNotEnoughPort             = newError(5020, "port exhausted")
	ErrInvalidContainerName      = newError(5021, "invalid container name")
	ErrAgentAlreadyRegistered    = newError(5022, "agent already registered")
	ErrAgentNotRegistered        = newError(5023, "agent not registered yet")
	ErrAgentExist                = newError(5024, "agent already exist")
	ErrAgentNotExist             = newError(5025, "agent not exist")
	ErrAgentExceedCapacity       = newError(5026, "exceed agent capacity")
	ErrLogAllLevel               = newError(5027, "can not log at ALL level")
	ErrDisabledKeyPair           = newError(5028, "key pair has been disabled")
	ErrProjectHasNoService       = newError(5029, "project has no service")
	ErrInvalidDomainPair         = newError(5030, "invalid pair of major domain and sub domain")
	ErrInvalidCronScheduleRule   = newError(5031, "invalid crontab schedule rule")
	ErrKeyPairNotExist           = newError(5032, "key pair dose not exist")
	ErrVersionIsNotOffline       = newError(5033, "Version is not offline")
	ErrServiceIsNotEmpty         = newError(5034, "Service is not empty")
	ErrDeveloperNotExist         = newError(5035, "developer is not exist")
	ErrSubDomainExist            = newError(5036, "sub domain already exist")
)

//////////////////////////////////////////////////////////
// 6001-8000 only for ablecloud error code
//////////////////////////////////////////////////////////
var (
	ErrNullValue           = newError(6001, "null value")
	ErrInvalidConf         = newError(6002, "invalid config")
	ErrNotInit             = newError(6003, "not inited")
	ErrAlreadyInited       = newError(6004, "already inited")
	ErrEntryNotDir         = newError(6005, "entry not dir")
	ErrInvalidEncryptKey   = newError(6006, "invalid encrypt key")
	ErrEntryNotExist       = newError(6007, "entry not exist")
	ErrEntryExist          = newError(6008, "entry exist")
	ErrIteratorEnd         = newError(6009, "iterator end")
	ErrInvalidVersion      = newError(6011, "invalid version")
	ErrInvalidResult       = newError(6012, "invalid result")
	ErrEncodeFailed        = newError(6013, "encode error")
	ErrDecodeFailed        = newError(6014, "decode error")
	ErrDataTypeError       = newError(6015, "data type error")
	ErrDataBaseNotExist    = newError(6016, "database not exist")
	ErrPartitionNotExist   = newError(6017, "partition not exist")
	ErrInvalidAccess       = newError(6018, "no privelige send msg")
	ErrNoValidEndpoint     = newError(6019, "no valid endpoint")
	ErrInvalidEndpoint     = newError(6020, "endpoint not in whitelist")
	ErrDeviceConnError     = newError(6021, "device connection exception")
	ErrInvalidDeviceMsg    = newError(6022, "invalid device message")
	ErrInvalidInnerRequest = newError(6023, "invalid inner request")
	ErrPayloadLenError     = newError(6024, "check payload length failed")
	ErrDecryptMsg          = newError(6025, "decrypt message error")
	ErrEncryptMsg          = newError(6026, "encrypt message error")
	ErrInvalidFormat       = newError(6027, "invalid format")
	ErrInvalidMetaName     = newError(6028, "invalid meta name")
	ErrServerBusy          = newError(6029, "server busy")
	ErrDailyLimit          = newError(6030, "have exceeded your daily limit")
	ErrInvalidKey          = newError(6031, "invalid key")
	ErrUserRateLimit       = newError(6032, "have exceeded the requests per second per userlimit")
)

//////////////////////////////////////////////////////////
// 10000 ~ 20000 for bigdata error code
//////////////////////////////////////////////////////////
var (
	ErrInvalidJson                 = newError(10000, "invalid json format")
	ErrFieldExist                  = newError(10001, "field already exist")
	ErrSchemaNotExist              = newError(10002, "schema not exist")
	ErrParseEvent                  = newError(10003, "pasre event data error")
	ErrTimeframeNotExist           = newError(10004, "time frame not exist")
	ErrTimeframeExist              = newError(10005, "time frame exist")
	ErrInvalidTimeframe            = newError(10006, "invalid timeframe format")
	ErrGroupByNotExist             = newError(10007, "group by not exist")
	ErrIntervalNotExist            = newError(10008, "interval not exist")
	ErrInvalidInterval             = newError(10009, "invalid interval format")
	ErrIntervalNotSupported        = newError(10010, "interval not supported")
	ErrFiltersNotExist             = newError(10011, "filters not exist")
	ErrInvalidFilters              = newError(10012, "invalid filters")
	ErrOperatorNotSupported        = newError(10013, "filter operator not supported")
	ErrQueryParamInvalid           = newError(10014, "invalid query parameter")
	ErrStepsNotExist               = newError(10015, "steps not exist")
	ErrInvalidStep                 = newError(10016, "invalid step parameter")
	ErrCohortsNotExist             = newError(10017, "cohorts not exist")
	ErrInvalidCohort               = newError(10018, "invalid cohort")
	ErrCriteriaNotExist            = newError(10019, "criteria not exist in cohort analysis")
	ErrInvalidCriteria             = newError(10020, "invalid criteria")
	ErrEOR                         = newError(10021, "end of rows")
	ErrWriteEventToFile            = newError(10022, "write event data to local file error")
	ErrOrderByNotExist             = newError(10023, "order by not exist")
	ErrLimitNotExist               = newError(10024, "limit not exist")
	ErrIdNotExist                  = newError(10025, "id in profile not exist")
	ErrIdentifierNotExist          = newError(10026, "identifier property not exist")
	ErrInvalidField                = newError(10027, "field not exist")
	ErrTimezoneNotExist            = newError(10028, "timezone not exist")
	ErrInvalidHaving               = newError(10029, "invalid having")
	ErrInvalidArgument             = newError(10030, "invalid argument")
	ErrSchemaAlreadyExists         = newError(10031, "schema already exists")
	ErrTooManyConnections          = newError(10032, "too many db connections")
	ErrNeedRetry                   = newError(10033, "need retry again")
	ErrKeyNotExists                = newError(10034, "key not exists")
	ErrBookmarkIdNotExists         = newError(10035, "bookmark id not exists")
	ErrModelNotExists              = newError(10036, "model not exists")
	ErrTableNotExists              = newError(10037, "table not exists")
	ErrMatchCRAndCIDS              = newError(10038, "cohort_relation and cohort_ids doesnt match")
	ErrUnsupportedRelation         = newError(10039, "this relation is not supported")
	ErrInvalidCohortIds            = newError(10040, "wrong with cohort_ids")
	ErrUnsupportedGroupByType      = newError(10041, "this group by type is not supported")
	ErrNoSpacesFound               = newError(10042, "no suitable spaces available")
	ErrUnsupportedTableType        = newError(10043, "table type unsupported")
	ErrInvalidProfileType          = newError(10044, "profile type is invalid")
	ErrExtractProfileTypeFailed    = newError(10045, "extract profile_type from profile table name failed")
	ErrFormulaExists               = newError(10046, "formula already exists")
	ErrFormulaNotExists            = newError(10047, "formula not exists")
	ErrFormulaIdNotExists          = newError(10048, "formula id not exists")
	ErrFormulaNameNotExists        = newError(10049, "formula name not exists")
	ErrInvalidFormulaExpression    = newError(10050, "invalid formula expression")
	ErrFormulaAggregationNotExists = newError(10051, "formula aggregation not exists")
	ErrFormulaInvalidWeight        = newError(10052, "formula weight invalid")
	ErrFormulaInvalidOperandValue  = newError(10053, "formula operand invalid")
	ErrInvalidTimeFormat           = newError(10054, "invalid time format")
)

func Error(code int64) *ACError {
	if errorMap[code] != nil {
		return errorMap[code]
	}
	return ErrInvalidError
}
