package notify

import "fmt"

// SMSNotification 短信通知方式
type SMSNotification struct{}

func (n *SMSNotification) Notify(message string, recipients map[string][]string) {
	if rev, ok := recipients["sms"]; ok {
		fmt.Println("Sending sms notification:", message, rev)
	}
}
