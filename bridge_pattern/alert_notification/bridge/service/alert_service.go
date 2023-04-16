package service

import (
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/alertrule"
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/api"
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/notify"
)

// 告警服务
type AlertService struct {
	AlertHandlers []alertrule.AlertHandler
	Notifications []notify.Notification
	ApiInfo       api.Info
}

func NewAlertService(alertHandlers []alertrule.AlertHandler, notifications []notify.Notification) *AlertService {
	return &AlertService{
		Notifications: notifications,
		AlertHandlers: alertHandlers,
	}
}

func (as *AlertService) Trigger(alert alertrule.AlertRule) {
	for _, handler := range as.AlertHandlers {
		handler.SetNotification(as.Notifications)
		handler.SetRule(alert)
		handler.Do(as.ApiInfo)
	}
	// 检查系统指标
	//for notifyType, recipients := range alert.Recipients {
	//	// 根据告警级别选择通知方式
	//	if notifier, ok := as.Notifications[notifyType]; ok {
	//
	//		notifier.Notify(fmt.Sprintf("[%s][%s] %s", alert.Level, alert.Name, alert.Message), recipients)
	//	}
	//}
}
