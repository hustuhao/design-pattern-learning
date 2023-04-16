package notify

import "fmt"

// EmailNotification 邮件通知方式
type EmailNotification struct{}

func (n *EmailNotification) Notify(message string, recipients map[string][]string) {
	if rev, ok := recipients["email"]; ok {
		fmt.Println("Sending email notification:", message, rev)
	}

}
