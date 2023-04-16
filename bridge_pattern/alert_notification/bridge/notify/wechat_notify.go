package notify

import "fmt"

// 微信通知方式
type WeChatNotification struct{}

func (n *WeChatNotification) Notify(message string, recipients map[string][]string) {
	if rev, ok := recipients["wechat"]; ok {
		fmt.Println("Sending wechat notification:", message, rev)
	}
}
