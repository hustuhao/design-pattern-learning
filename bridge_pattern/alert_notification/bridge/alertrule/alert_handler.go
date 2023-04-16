package alertrule

import (
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/api"
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/notify"
)

type AlertHandler interface {
	Do(info api.Info)
	SetRule(rule AlertRule)
	SetNotification(notis []notify.Notification)
}

type AlertHandlerImpl struct {
	rule  AlertRule
	notis []notify.Notification //多种通知方式
}

func (r *AlertHandlerImpl) Do(info api.Info) {
	if r.Check(info) {
		r.Notify()
	}
}

func (r *AlertHandlerImpl) Check(info api.Info) bool {
	return true
}

func (r *AlertHandlerImpl) SetRule(rule AlertRule) {
	r.rule = rule
}

func (r *AlertHandlerImpl) SetNotification(notis []notify.Notification) {
	r.notis = notis
}

func (r *AlertHandlerImpl) Notify() {
	for _, noti := range r.notis {
		noti.Notify(r.rule.Message, r.rule.Recipients)
	}
}
