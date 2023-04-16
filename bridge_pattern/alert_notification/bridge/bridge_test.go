package bridge

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/alertrule"
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/notify"
	"turato.com/design-pattern/bridge_pattern/alert_notification/bridge/service"
)

// 根据严重程度，有不同的通知方式（notify）,具体发送信息又有不同的实现（ msgsender）
//针对 Notification 的代码，我们将不同渠道的发送逻辑离出来，形成独立的消息发送类 （MsgSender 相关类）。
//其中，Notification 类相当于抽象，MsgSender 类相当于实现，两者可以独立开发，通过组合关系（也就是桥梁）任意组合在一起。
//所谓任意组合的意思就是，不同紧急程度的消息和发送渠道之间的对应关系，不是在代码中固定写死的，我们可以动态地去指定（比如，通过读取配置来获取对应关系）。

var config alertrule.Config

func LoadConfig() {
	// ioutil.ReadFile 会一次性将整个文件读入内存，这对于小型文件来说没有问题，但对于大型文件来说，可能会导致内存泄漏问题。
	//具体来说，如果您的程序读取大量大型文件，并使用 ioutil.ReadFile 函数来读取它们，那么每次调用该函数时，都会在内存中分配一个缓冲区来保存文件内容。由于这些缓冲区不会自动释放，因此它们可能会占用大量内存，从而导致内存泄漏问题。
	//相比之下，使用 os.Open 和 ioutil.ReadAll 可以一次只读取文件的一部分，这可以避免一次性将整个文件读入内存并导致内存泄漏问题。
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Failed to open YAML file: %v", err)
	}
	defer file.Close()

	yamlData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	fmt.Printf("%+v\n", config)
}

func BuildNotification() {

}

func TestName(t *testing.T) {
	LoadConfig() // 配置可以写入数据库
	// 初始化通知方式
	//notifications := []notify.Notification{ // 可以继续优化，email、sms、wechat 可以作为关键字
	//	"email":  &notify.EmailNotification{},
	//	"sms":    &notify.SMSNotification{},
	//	"wechat": &notify.WeChatNotification{}, // 还有其他没有支持的方式
	//}
	notifications := []notify.Notification{
		&notify.EmailNotification{},
		&notify.SMSNotification{},
		&notify.WeChatNotification{},
	}

	alertHandlers := []alertrule.AlertHandler{
		&alertrule.CpuAlertHandler{},
	}
	// 初始化告警服务
	alertService := service.NewAlertService(alertHandlers, notifications)

	// 触发告警
	for _, alert := range config.Alerts {
		alertService.Trigger(alert)
	}
	// OUTPUT:
	// Sending email notification: [critical][High CPU Usage] The CPU usage on the server has exceeded 90%. [admin@example.com op@example.com]
	// Sending SMS notification: [critical][High CPU Usage] The CPU usage on the server has exceeded 90%. [+1-555-1234 +1-666-1234]
	// Sending email notification: [warning][Low Disk Space] The available disk space on the server is less than 10%. [developer@example.com]
}
