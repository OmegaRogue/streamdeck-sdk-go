package sdk

const (
	CommonAction     = "action"
	CommonEvent      = "event"
	CommonContext    = "context"
	CommonPayload    = "payload"
	CommonDevice     = "device"
	CommonDeviceInfo = "deviceInfo"

	EventKeyDown                       = "keyDown"
	EventKeyUp                         = "keyUp"
	EventWillAppear                    = "willAppear"
	EventWillDisappear                 = "willDisappear"
	EventDeviceDidConnect              = "deviceDidConnect"
	EventDeviceDidDisconnect           = "deviceDidDisconnect"
	EventApplicationDidLaunch          = "applicationDidLaunch"
	EventApplicationDidTerminate       = "applicationDidTerminate"
	EventTitleParametersDidChange      = "titleParametersDidChange"
	EventDidReceiveSettings            = "didReceiveSettings"
	EventDidReceiveGlobalSettings      = "didReceiveGlobalSettings"
	EventPropertyInspectorDidAppear    = "propertyInspectorDidAppear"
	EventPropertyInspectorDidDisappear = "propertyInspectorDidDisappear"
	EventDialDown                      = "dialDown"
	EventDialUp                        = "dialUp"
	EventDialRotate                    = "dialRotate"

	EventSetTitle                = "setTitle"
	EventSetImage                = "setImage"
	EventShowAlert               = "showAlert"
	EventShowOK                  = "showOk"
	EventGetSettings             = "getSettings"
	EventSetSettings             = "setSettings"
	EventGetGlobalSettings       = "getGlobalSettings"
	EventSetGlobalSettings       = "setGlobalSettings"
	EventSetState                = "setState"
	EventSwitchToProfile         = "switchToProfile"
	EventSendToPropertyInspector = "sendToPropertyInspector"
	EventSendToPlugin            = "sendToPlugin"
	EventOpenURL                 = "openUrl"
	EventLogMessage              = "logMessage"

	TargetBoth     = 0
	TargetHardware = 1
	TargetSoftware = 2
)
