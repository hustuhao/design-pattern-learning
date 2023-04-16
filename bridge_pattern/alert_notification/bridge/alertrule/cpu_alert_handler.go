package alertrule

import (
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/api"
)

type CpuAlertHandler struct {
	AlertHandlerImpl
}

func (c *CpuAlertHandler) Check(info api.Info) bool {
	// 根据 info 中的 CPU 使用情况，进行告警分级，判断是否符合
	return true
}
