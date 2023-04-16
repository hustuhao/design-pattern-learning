package simple

type NotificationEmergencyLevel int32

const (
	SEVERE  NotificationEmergencyLevel = 0 // 严重
	URGENCY NotificationEmergencyLevel = 1 // 紧急
	NORMAL  NotificationEmergencyLevel = 2 // 普通
	TRIVIAL NotificationEmergencyLevel = 3 // 无关紧要
)

// Notification 通知类
type Notification struct {
	Telephones     []string
	EmailAddresses []string
	WechatIds      []string
}

func (n *Notification) Notify(level NotificationEmergencyLevel, msg string) {
	if level == SEVERE {
		// 自动语音电话通知

	} else if level == URGENCY {
		// 短信通知

	} else if level == NORMAL {
		// 微信通知

	} else if level == TRIVIAL {
		// 发邮件

	}
}
